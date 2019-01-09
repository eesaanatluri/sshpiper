// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package costmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2018-05-31/costmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type FormatType = original.FormatType

const (
	Csv FormatType = original.Csv
)

type GranularityType = original.GranularityType

const (
	Daily GranularityType = original.Daily
)

type RecurrenceType = original.RecurrenceType

const (
	RecurrenceTypeAnnually RecurrenceType = original.RecurrenceTypeAnnually
	RecurrenceTypeDaily    RecurrenceType = original.RecurrenceTypeDaily
	RecurrenceTypeMonthly  RecurrenceType = original.RecurrenceTypeMonthly
	RecurrenceTypeWeekly   RecurrenceType = original.RecurrenceTypeWeekly
)

type ReportConfigColumnType = original.ReportConfigColumnType

const (
	ReportConfigColumnTypeDimension ReportConfigColumnType = original.ReportConfigColumnTypeDimension
	ReportConfigColumnTypeTag       ReportConfigColumnType = original.ReportConfigColumnTypeTag
)

type StatusType = original.StatusType

const (
	Active   StatusType = original.Active
	Inactive StatusType = original.Inactive
)

type TimeframeType = original.TimeframeType

const (
	Custom      TimeframeType = original.Custom
	MonthToDate TimeframeType = original.MonthToDate
	WeekToDate  TimeframeType = original.WeekToDate
	YearToDate  TimeframeType = original.YearToDate
)

type BaseClient = original.BaseClient
type BillingAccountDimensionsClient = original.BillingAccountDimensionsClient
type Dimension = original.Dimension
type DimensionProperties = original.DimensionProperties
type DimensionsListResult = original.DimensionsListResult
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Query = original.Query
type QueryColumn = original.QueryColumn
type QueryProperties = original.QueryProperties
type QueryResult = original.QueryResult
type ReportConfig = original.ReportConfig
type ReportConfigAggregation = original.ReportConfigAggregation
type ReportConfigClient = original.ReportConfigClient
type ReportConfigComparisonExpression = original.ReportConfigComparisonExpression
type ReportConfigDataset = original.ReportConfigDataset
type ReportConfigDatasetConfiguration = original.ReportConfigDatasetConfiguration
type ReportConfigDefinition = original.ReportConfigDefinition
type ReportConfigDeliveryDestination = original.ReportConfigDeliveryDestination
type ReportConfigDeliveryInfo = original.ReportConfigDeliveryInfo
type ReportConfigFilter = original.ReportConfigFilter
type ReportConfigGrouping = original.ReportConfigGrouping
type ReportConfigListResult = original.ReportConfigListResult
type ReportConfigProperties = original.ReportConfigProperties
type ReportConfigRecurrencePeriod = original.ReportConfigRecurrencePeriod
type ReportConfigSchedule = original.ReportConfigSchedule
type ReportConfigTimePeriod = original.ReportConfigTimePeriod
type Resource = original.Resource
type ResourceGroupDimensionsClient = original.ResourceGroupDimensionsClient
type SubscriptionDimensionsClient = original.SubscriptionDimensionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBillingAccountDimensionsClient(subscriptionID string) BillingAccountDimensionsClient {
	return original.NewBillingAccountDimensionsClient(subscriptionID)
}
func NewBillingAccountDimensionsClientWithBaseURI(baseURI string, subscriptionID string) BillingAccountDimensionsClient {
	return original.NewBillingAccountDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReportConfigClient(subscriptionID string) ReportConfigClient {
	return original.NewReportConfigClient(subscriptionID)
}
func NewReportConfigClientWithBaseURI(baseURI string, subscriptionID string) ReportConfigClient {
	return original.NewReportConfigClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceGroupDimensionsClient(subscriptionID string) ResourceGroupDimensionsClient {
	return original.NewResourceGroupDimensionsClient(subscriptionID)
}
func NewResourceGroupDimensionsClientWithBaseURI(baseURI string, subscriptionID string) ResourceGroupDimensionsClient {
	return original.NewResourceGroupDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionDimensionsClient(subscriptionID string) SubscriptionDimensionsClient {
	return original.NewSubscriptionDimensionsClient(subscriptionID)
}
func NewSubscriptionDimensionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionDimensionsClient {
	return original.NewSubscriptionDimensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleFormatTypeValues() []FormatType {
	return original.PossibleFormatTypeValues()
}
func PossibleGranularityTypeValues() []GranularityType {
	return original.PossibleGranularityTypeValues()
}
func PossibleRecurrenceTypeValues() []RecurrenceType {
	return original.PossibleRecurrenceTypeValues()
}
func PossibleReportConfigColumnTypeValues() []ReportConfigColumnType {
	return original.PossibleReportConfigColumnTypeValues()
}
func PossibleStatusTypeValues() []StatusType {
	return original.PossibleStatusTypeValues()
}
func PossibleTimeframeTypeValues() []TimeframeType {
	return original.PossibleTimeframeTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
