// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/auth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sbauth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/stretchr/testify/require"
)

type fakeTokenCredential struct {
	azcore.TokenCredential
	expires time.Time
}

func (ftc *fakeTokenCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		ExpiresOn: ftc.expires,
	}, nil
}

func TestNamespaceNegotiateClaim(t *testing.T) {
	expires := time.Now().Add(24 * time.Hour)

	ns := &Namespace{
		RetryOptions:  retryOptionsOnlyOnce,
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{expires: expires}),
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++
		return nil
	}

	newAMQPClientCalled := 0

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		newAMQPClientCalled++
		return &amqpwrap.AMQPClientWrapper{}, nil
	}

	// fire off a basic negotiate claim. The renewal duration is so long that it won't run - that's a separate test.
	cancel, _, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"my entity path",
		cbsNegotiateClaim,
		func(expirationTimeParam, currentTime time.Time) time.Duration {
			require.EqualValues(t, expires, expirationTimeParam)
			// wiggle room, but just want to check that they're passing me the time.Now() value (silly)
			require.GreaterOrEqual(t, time.Minute, time.Since(currentTime))

			// we're going to cancel out pretty much immediately
			return 24 * time.Hour
		})
	defer cancel()

	require.NoError(t, err)
	cancel()

	require.EqualValues(t, newAMQPClientCalled, 1)
	require.EqualValues(t, 1, cbsNegotiateClaimCalled)
}

func TestNamespaceNegotiateClaimRenewal(t *testing.T) {
	expires := time.Now().Add(24 * time.Hour)

	ns := &Namespace{
		RetryOptions:  retryOptionsOnlyOnce,
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{expires: expires}),
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++
		return nil
	}

	var errorsLogged []error
	nextRefreshDurationChecks := 0

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return &amqpwrap.AMQPClientWrapper{Inner: &amqp.Client{}}, nil
	}

	cancel, _, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"my entity path",
		cbsNegotiateClaim, func(expirationTimeParam, currentTime time.Time) time.Duration {
			require.EqualValues(t, expires, expirationTimeParam)
			nextRefreshDurationChecks++

			if nextRefreshDurationChecks == 1 {
				return 0
			}

			return 24 * time.Hour // ie, we don't need to do it again.
		})
	defer cancel()

	require.NoError(t, err)
	time.Sleep(3 * time.Second) // make sure, even with variability, we get at least one renewal

	require.EqualValues(t, 2, nextRefreshDurationChecks)
	require.EqualValues(t, 2, cbsNegotiateClaimCalled)
	require.Empty(t, errorsLogged)

	cancel()
}

func TestNamespaceNegotiateClaimFailsToGetClient(t *testing.T) {
	ns := &Namespace{
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{expires: time.Now()}),
	}

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return nil, errors.New("Getting *amqp.Client failed")
	}

	cancel, _, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"entity path",
		func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
			return errors.New("NegotiateClaim amqp.Client failed")
		}, func(expirationTime, currentTime time.Time) time.Duration {
			// refresh immediately since we're in a unit test.
			return 0
		})

	require.EqualError(t, err, "Getting *amqp.Client failed")
	require.Nil(t, cancel)
}

func TestNamespaceNegotiateClaimNonRenewableToken(t *testing.T) {
	ns := &Namespace{
		RetryOptions: retryOptionsOnlyOnce,
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{
			// credentials that don't renew return a zero-initialized time.
			expires: time.Time{},
		}),
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++
		return nil
	}

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return &amqpwrap.AMQPClientWrapper{Inner: &amqp.Client{}}, nil
	}

	// since the token is non-renewable we will just do the single cbsNegotiateClaim call and never renew.
	_, done, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"my entity path",
		cbsNegotiateClaim,
		func(expirationTimeParam, currentTime time.Time) time.Duration {
			panic("Won't be called, no refreshing of claims will be done")
		})

	require.NoError(t, err)
	require.Equal(t, 1, cbsNegotiateClaimCalled)

	select {
	case <-done:
	default:
		require.Fail(t, "cancel() returns a channel that is already Done()")
	}
}

func TestNamespaceNegotiateClaimFails(t *testing.T) {
	ns := &Namespace{
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{expires: time.Now()}),
	}

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return &fakeAMQPClient{}, nil
	}

	cancel, _, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"entity path",
		func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
			return errors.New("NegotiateClaim amqp.Client failed")
		}, func(expirationTime, currentTime time.Time) time.Duration {
			// not even used.
			return 0
		})

	require.EqualError(t, err, "NegotiateClaim amqp.Client failed")
	require.Nil(t, cancel)
}

func TestNamespaceNegotiateClaimFatalErrors(t *testing.T) {
	ns := &Namespace{
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{expires: time.Now()}),
	}

	cbsNegotiateClaimCalled := 0

	cbsNegotiateClaim := func(ctx context.Context, audience string, conn amqpwrap.AMQPClient, provider auth.TokenProvider) error {
		cbsNegotiateClaimCalled++

		// work the first time, fail on renewals.
		if cbsNegotiateClaimCalled > 1 {
			return errNonRetriable{Message: "non retriable error message"}
		}

		return nil
	}

	endCapture := test.CaptureLogsForTest()
	defer endCapture()

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return &amqpwrap.AMQPClientWrapper{Inner: &amqp.Client{}}, nil
	}

	_, done, err := ns.startNegotiateClaimRenewer(
		context.Background(),
		"entity path",
		cbsNegotiateClaim, func(expirationTime, currentTime time.Time) time.Duration {
			// instant renewals.
			return 0
		})

	require.NoError(t, err)

	select {
	case <-done:
		logs := endCapture()
		// check the log messages - we should have one telling us why we stopped the claims loop
		require.Contains(t, logs, "[azsb.Auth] [entity path] fatal error, stopping token refresh loop: non retriable error message")
	case <-time.After(3 * time.Second):
		// was locked! Should have been closed.
		require.Fail(t, "claim renewal was automatically cancelled because of a non-retriable error")
	}
}

func TestNamespaceNextClaimRefreshDuration(t *testing.T) {
	now := time.Now()

	clockDrift := 10 * time.Minute
	lessThanMin := now.Add(119 * time.Second).Add(clockDrift)
	greaterThanMax := now.Add(49*24*time.Hour + time.Second).Add(clockDrift)

	require.EqualValues(t, 2*time.Minute, nextClaimRefreshDuration(lessThanMin, now),
		"Just under the min refresh time, so we get the min instead")

	require.EqualValues(t, 49*24*time.Hour, nextClaimRefreshDuration(greaterThanMax, now),
		"Just over the max refresh time, so we just get the max instead")

	require.EqualValues(t, 3*time.Minute, nextClaimRefreshDuration(now.Add(3*time.Minute+clockDrift), now))
}

func TestNamespaceStaleConnection(t *testing.T) {
	ns := &Namespace{
		RetryOptions: retryOptionsOnlyOnce,
		TokenProvider: sbauth.NewTokenProvider(&fakeTokenCredential{
			// credentials that don't renew return a zero-initialized time.
			expires: time.Time{},
		}),
	}

	fakeClient := &fakeAMQPClient{}

	ns.client = fakeClient
	ns.connID = 101

	require.NoError(t, ns.Close(context.Background(), false))
	require.Equal(t, 1, fakeClient.closeCalled)
	require.Nil(t, ns.client)

	ns.newClientFn = func(ctx context.Context) (amqpwrap.AMQPClient, error) {
		return &fakeAMQPClient{}, nil
	}

	client, clientID, err := ns.GetAMQPClientImpl(context.Background())
	require.NoError(t, err)
	require.NotSame(t, fakeClient, client, "A new client should be created")
	require.Equal(t, uint64(101+1), clientID, "Client ID is incremented since we had to recreate it")
	require.NotNil(t, client)
}

func TestNamespaceUpdateClientWithoutLock(t *testing.T) {
	newClient := 0
	var clientToReturn amqpwrap.AMQPClient
	var err error

	ns := &Namespace{
		newClientFn: func(ctx context.Context) (amqpwrap.AMQPClient, error) {
			newClient++
			return clientToReturn, err
		},
		connID: 101,
	}

	err = errors.New("client error")

	client, clientID, err := ns.updateClientWithoutLock(context.Background())
	require.Error(t, err, "client error")
	require.Equal(t, uint64(0), clientID)
	require.Nil(t, client)

	// when they create a new client they'll get this one.
	clientToReturn = &fakeAMQPClient{}
	err = nil

	client, clientID, err = ns.updateClientWithoutLock(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(101+1), clientID)
	require.Same(t, clientToReturn, client)

	// change out the returned client (it won't get used because we return the cached one in ns.client)
	origClient := client
	clientToReturn = &fakeAMQPClient{}

	client, clientID, err = ns.updateClientWithoutLock(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(101+1), clientID)
	require.Same(t, origClient, client)
}

type fakeAMQPClient struct {
	amqpwrap.AMQPClient
	closeCalled int
}

func (f *fakeAMQPClient) Close() error {
	f.closeCalled++
	return nil
}
