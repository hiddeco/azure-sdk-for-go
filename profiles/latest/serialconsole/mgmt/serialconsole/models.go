//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package serialconsole

import original "github.com/Azure/azure-sdk-for-go/services/serialconsole/mgmt/2018-05-01/serialconsole"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DisableSerialConsoleResult = original.DisableSerialConsoleResult
type EnableSerialConsoleResult = original.EnableSerialConsoleResult
type GetSerialConsoleSubscriptionNotFound = original.GetSerialConsoleSubscriptionNotFound
type Operations = original.Operations
type OperationsValueItem = original.OperationsValueItem
type OperationsValueItemDisplay = original.OperationsValueItemDisplay
type SetObject = original.SetObject
type Status = original.Status

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
