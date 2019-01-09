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

package containerregistry

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2018-02-01/containerregistry"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type BaseImageDependencyType = original.BaseImageDependencyType

const (
	BuildTime BaseImageDependencyType = original.BuildTime
	RunTime   BaseImageDependencyType = original.RunTime
)

type BaseImageTriggerType = original.BaseImageTriggerType

const (
	All     BaseImageTriggerType = original.All
	None    BaseImageTriggerType = original.None
	Runtime BaseImageTriggerType = original.Runtime
)

type BuildStatus = original.BuildStatus

const (
	Canceled  BuildStatus = original.Canceled
	Error     BuildStatus = original.Error
	Failed    BuildStatus = original.Failed
	Queued    BuildStatus = original.Queued
	Running   BuildStatus = original.Running
	Started   BuildStatus = original.Started
	Succeeded BuildStatus = original.Succeeded
	Timeout   BuildStatus = original.Timeout
)

type BuildTaskStatus = original.BuildTaskStatus

const (
	Disabled BuildTaskStatus = original.Disabled
	Enabled  BuildTaskStatus = original.Enabled
)

type BuildType = original.BuildType

const (
	AutoBuild  BuildType = original.AutoBuild
	QuickBuild BuildType = original.QuickBuild
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type ImportMode = original.ImportMode

const (
	Force   ImportMode = original.Force
	NoForce ImportMode = original.NoForce
)

type OsType = original.OsType

const (
	Linux   OsType = original.Linux
	Windows OsType = original.Windows
)

type PasswordName = original.PasswordName

const (
	Password  PasswordName = original.Password
	Password2 PasswordName = original.Password2
)

type PolicyStatus = original.PolicyStatus

const (
	PolicyStatusDisabled PolicyStatus = original.PolicyStatusDisabled
	PolicyStatusEnabled  PolicyStatus = original.PolicyStatusEnabled
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type RegistryUsageUnit = original.RegistryUsageUnit

const (
	Bytes RegistryUsageUnit = original.Bytes
	Count RegistryUsageUnit = original.Count
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Classic  SkuName = original.Classic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierClassic  SkuTier = original.SkuTierClassic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type SourceControlType = original.SourceControlType

const (
	Github                  SourceControlType = original.Github
	VisualStudioTeamService SourceControlType = original.VisualStudioTeamService
)

type TokenType = original.TokenType

const (
	OAuth TokenType = original.OAuth
	PAT   TokenType = original.PAT
)

type TrustPolicyType = original.TrustPolicyType

const (
	Notary TrustPolicyType = original.Notary
)

type Type = original.Type

const (
	TypeBuildStepProperties Type = original.TypeBuildStepProperties
	TypeDocker              Type = original.TypeDocker
)

type TypeBasicBuildStepPropertiesUpdateParameters = original.TypeBasicBuildStepPropertiesUpdateParameters

const (
	TypeBasicBuildStepPropertiesUpdateParametersTypeBuildStepPropertiesUpdateParameters TypeBasicBuildStepPropertiesUpdateParameters = original.TypeBasicBuildStepPropertiesUpdateParametersTypeBuildStepPropertiesUpdateParameters
	TypeBasicBuildStepPropertiesUpdateParametersTypeDocker                              TypeBasicBuildStepPropertiesUpdateParameters = original.TypeBasicBuildStepPropertiesUpdateParametersTypeDocker
)

type TypeBasicQueueBuildRequest = original.TypeBasicQueueBuildRequest

const (
	TypeBuildTask         TypeBasicQueueBuildRequest = original.TypeBuildTask
	TypeQueueBuildRequest TypeBasicQueueBuildRequest = original.TypeQueueBuildRequest
	TypeQuickBuild        TypeBasicQueueBuildRequest = original.TypeQuickBuild
)

type WebhookAction = original.WebhookAction

const (
	ChartDelete WebhookAction = original.ChartDelete
	ChartPush   WebhookAction = original.ChartPush
	Delete      WebhookAction = original.Delete
	Push        WebhookAction = original.Push
	Quarantine  WebhookAction = original.Quarantine
)

type WebhookStatus = original.WebhookStatus

const (
	WebhookStatusDisabled WebhookStatus = original.WebhookStatusDisabled
	WebhookStatusEnabled  WebhookStatus = original.WebhookStatusEnabled
)

type Actor = original.Actor
type BaseClient = original.BaseClient
type BaseImageDependency = original.BaseImageDependency
type BasicBuildStepProperties = original.BasicBuildStepProperties
type BasicBuildStepPropertiesUpdateParameters = original.BasicBuildStepPropertiesUpdateParameters
type BasicQueueBuildRequest = original.BasicQueueBuildRequest
type Build = original.Build
type BuildArgument = original.BuildArgument
type BuildArgumentList = original.BuildArgumentList
type BuildArgumentListIterator = original.BuildArgumentListIterator
type BuildArgumentListPage = original.BuildArgumentListPage
type BuildFilter = original.BuildFilter
type BuildGetLogResult = original.BuildGetLogResult
type BuildListResult = original.BuildListResult
type BuildListResultIterator = original.BuildListResultIterator
type BuildListResultPage = original.BuildListResultPage
type BuildProperties = original.BuildProperties
type BuildStep = original.BuildStep
type BuildStepList = original.BuildStepList
type BuildStepListIterator = original.BuildStepListIterator
type BuildStepListPage = original.BuildStepListPage
type BuildStepProperties = original.BuildStepProperties
type BuildStepPropertiesUpdateParameters = original.BuildStepPropertiesUpdateParameters
type BuildStepUpdateParameters = original.BuildStepUpdateParameters
type BuildStepsClient = original.BuildStepsClient
type BuildStepsCreateFuture = original.BuildStepsCreateFuture
type BuildStepsDeleteFuture = original.BuildStepsDeleteFuture
type BuildStepsUpdateFuture = original.BuildStepsUpdateFuture
type BuildTask = original.BuildTask
type BuildTaskBuildRequest = original.BuildTaskBuildRequest
type BuildTaskFilter = original.BuildTaskFilter
type BuildTaskListResult = original.BuildTaskListResult
type BuildTaskListResultIterator = original.BuildTaskListResultIterator
type BuildTaskListResultPage = original.BuildTaskListResultPage
type BuildTaskProperties = original.BuildTaskProperties
type BuildTaskPropertiesUpdateParameters = original.BuildTaskPropertiesUpdateParameters
type BuildTaskUpdateParameters = original.BuildTaskUpdateParameters
type BuildTasksClient = original.BuildTasksClient
type BuildTasksCreateFuture = original.BuildTasksCreateFuture
type BuildTasksDeleteFuture = original.BuildTasksDeleteFuture
type BuildTasksUpdateFuture = original.BuildTasksUpdateFuture
type BuildUpdateParameters = original.BuildUpdateParameters
type BuildsCancelFuture = original.BuildsCancelFuture
type BuildsClient = original.BuildsClient
type BuildsUpdateFuture = original.BuildsUpdateFuture
type CallbackConfig = original.CallbackConfig
type DockerBuildStep = original.DockerBuildStep
type DockerBuildStepUpdateParameters = original.DockerBuildStepUpdateParameters
type Event = original.Event
type EventContent = original.EventContent
type EventInfo = original.EventInfo
type EventListResult = original.EventListResult
type EventListResultIterator = original.EventListResultIterator
type EventListResultPage = original.EventListResultPage
type EventRequestMessage = original.EventRequestMessage
type EventResponseMessage = original.EventResponseMessage
type GitCommitTrigger = original.GitCommitTrigger
type IPRule = original.IPRule
type ImageDescriptor = original.ImageDescriptor
type ImageUpdateTrigger = original.ImageUpdateTrigger
type ImportImageParameters = original.ImportImageParameters
type ImportSource = original.ImportSource
type ImportSourceCredentials = original.ImportSourceCredentials
type NetworkRuleSet = original.NetworkRuleSet
type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationMetricSpecificationDefinition = original.OperationMetricSpecificationDefinition
type OperationPropertiesDefinition = original.OperationPropertiesDefinition
type OperationServiceSpecificationDefinition = original.OperationServiceSpecificationDefinition
type OperationsClient = original.OperationsClient
type PlatformProperties = original.PlatformProperties
type ProxyResource = original.ProxyResource
type QuarantinePolicy = original.QuarantinePolicy
type QueueBuildRequest = original.QueueBuildRequest
type QuickBuildRequest = original.QuickBuildRequest
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type RegistriesClient = original.RegistriesClient
type RegistriesCreateFuture = original.RegistriesCreateFuture
type RegistriesDeleteFuture = original.RegistriesDeleteFuture
type RegistriesImportImageFuture = original.RegistriesImportImageFuture
type RegistriesQueueBuildFuture = original.RegistriesQueueBuildFuture
type RegistriesUpdateFuture = original.RegistriesUpdateFuture
type RegistriesUpdatePoliciesFuture = original.RegistriesUpdatePoliciesFuture
type Registry = original.Registry
type RegistryIdentity = original.RegistryIdentity
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryListResultIterator = original.RegistryListResultIterator
type RegistryListResultPage = original.RegistryListResultPage
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryPolicies = original.RegistryPolicies
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type RegistryUsage = original.RegistryUsage
type RegistryUsageListResult = original.RegistryUsageListResult
type Replication = original.Replication
type ReplicationListResult = original.ReplicationListResult
type ReplicationListResultIterator = original.ReplicationListResultIterator
type ReplicationListResultPage = original.ReplicationListResultPage
type ReplicationProperties = original.ReplicationProperties
type ReplicationUpdateParameters = original.ReplicationUpdateParameters
type ReplicationsClient = original.ReplicationsClient
type ReplicationsCreateFuture = original.ReplicationsCreateFuture
type ReplicationsDeleteFuture = original.ReplicationsDeleteFuture
type ReplicationsUpdateFuture = original.ReplicationsUpdateFuture
type Request = original.Request
type Resource = original.Resource
type Sku = original.Sku
type Source = original.Source
type SourceControlAuthInfo = original.SourceControlAuthInfo
type SourceRepositoryProperties = original.SourceRepositoryProperties
type SourceRepositoryUpdateParameters = original.SourceRepositoryUpdateParameters
type SourceUploadDefinition = original.SourceUploadDefinition
type Status = original.Status
type StorageAccountProperties = original.StorageAccountProperties
type Target = original.Target
type TrustPolicy = original.TrustPolicy
type VirtualNetworkRule = original.VirtualNetworkRule
type Webhook = original.Webhook
type WebhookCreateParameters = original.WebhookCreateParameters
type WebhookListResult = original.WebhookListResult
type WebhookListResultIterator = original.WebhookListResultIterator
type WebhookListResultPage = original.WebhookListResultPage
type WebhookProperties = original.WebhookProperties
type WebhookPropertiesCreateParameters = original.WebhookPropertiesCreateParameters
type WebhookPropertiesUpdateParameters = original.WebhookPropertiesUpdateParameters
type WebhookUpdateParameters = original.WebhookUpdateParameters
type WebhooksClient = original.WebhooksClient
type WebhooksCreateFuture = original.WebhooksCreateFuture
type WebhooksDeleteFuture = original.WebhooksDeleteFuture
type WebhooksUpdateFuture = original.WebhooksUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBuildArgumentListIterator(page BuildArgumentListPage) BuildArgumentListIterator {
	return original.NewBuildArgumentListIterator(page)
}
func NewBuildArgumentListPage(getNextPage func(context.Context, BuildArgumentList) (BuildArgumentList, error)) BuildArgumentListPage {
	return original.NewBuildArgumentListPage(getNextPage)
}
func NewBuildListResultIterator(page BuildListResultPage) BuildListResultIterator {
	return original.NewBuildListResultIterator(page)
}
func NewBuildListResultPage(getNextPage func(context.Context, BuildListResult) (BuildListResult, error)) BuildListResultPage {
	return original.NewBuildListResultPage(getNextPage)
}
func NewBuildStepListIterator(page BuildStepListPage) BuildStepListIterator {
	return original.NewBuildStepListIterator(page)
}
func NewBuildStepListPage(getNextPage func(context.Context, BuildStepList) (BuildStepList, error)) BuildStepListPage {
	return original.NewBuildStepListPage(getNextPage)
}
func NewBuildStepsClient(subscriptionID string) BuildStepsClient {
	return original.NewBuildStepsClient(subscriptionID)
}
func NewBuildStepsClientWithBaseURI(baseURI string, subscriptionID string) BuildStepsClient {
	return original.NewBuildStepsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBuildTaskListResultIterator(page BuildTaskListResultPage) BuildTaskListResultIterator {
	return original.NewBuildTaskListResultIterator(page)
}
func NewBuildTaskListResultPage(getNextPage func(context.Context, BuildTaskListResult) (BuildTaskListResult, error)) BuildTaskListResultPage {
	return original.NewBuildTaskListResultPage(getNextPage)
}
func NewBuildTasksClient(subscriptionID string) BuildTasksClient {
	return original.NewBuildTasksClient(subscriptionID)
}
func NewBuildTasksClientWithBaseURI(baseURI string, subscriptionID string) BuildTasksClient {
	return original.NewBuildTasksClientWithBaseURI(baseURI, subscriptionID)
}
func NewBuildsClient(subscriptionID string) BuildsClient {
	return original.NewBuildsClient(subscriptionID)
}
func NewBuildsClientWithBaseURI(baseURI string, subscriptionID string) BuildsClient {
	return original.NewBuildsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventListResultIterator(page EventListResultPage) EventListResultIterator {
	return original.NewEventListResultIterator(page)
}
func NewEventListResultPage(getNextPage func(context.Context, EventListResult) (EventListResult, error)) EventListResultPage {
	return original.NewEventListResultPage(getNextPage)
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
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistryListResultIterator(page RegistryListResultPage) RegistryListResultIterator {
	return original.NewRegistryListResultIterator(page)
}
func NewRegistryListResultPage(getNextPage func(context.Context, RegistryListResult) (RegistryListResult, error)) RegistryListResultPage {
	return original.NewRegistryListResultPage(getNextPage)
}
func NewReplicationListResultIterator(page ReplicationListResultPage) ReplicationListResultIterator {
	return original.NewReplicationListResultIterator(page)
}
func NewReplicationListResultPage(getNextPage func(context.Context, ReplicationListResult) (ReplicationListResult, error)) ReplicationListResultPage {
	return original.NewReplicationListResultPage(getNextPage)
}
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClient(subscriptionID)
}
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebhookListResultIterator(page WebhookListResultPage) WebhookListResultIterator {
	return original.NewWebhookListResultIterator(page)
}
func NewWebhookListResultPage(getNextPage func(context.Context, WebhookListResult) (WebhookListResult, error)) WebhookListResultPage {
	return original.NewWebhookListResultPage(getNextPage)
}
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return original.NewWebhooksClient(subscriptionID)
}
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return original.NewWebhooksClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return original.PossibleBaseImageDependencyTypeValues()
}
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return original.PossibleBaseImageTriggerTypeValues()
}
func PossibleBuildStatusValues() []BuildStatus {
	return original.PossibleBuildStatusValues()
}
func PossibleBuildTaskStatusValues() []BuildTaskStatus {
	return original.PossibleBuildTaskStatusValues()
}
func PossibleBuildTypeValues() []BuildType {
	return original.PossibleBuildTypeValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleImportModeValues() []ImportMode {
	return original.PossibleImportModeValues()
}
func PossibleOsTypeValues() []OsType {
	return original.PossibleOsTypeValues()
}
func PossiblePasswordNameValues() []PasswordName {
	return original.PossiblePasswordNameValues()
}
func PossiblePolicyStatusValues() []PolicyStatus {
	return original.PossiblePolicyStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return original.PossibleRegistryUsageUnitValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSourceControlTypeValues() []SourceControlType {
	return original.PossibleSourceControlTypeValues()
}
func PossibleTokenTypeValues() []TokenType {
	return original.PossibleTokenTypeValues()
}
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return original.PossibleTrustPolicyTypeValues()
}
func PossibleTypeBasicBuildStepPropertiesUpdateParametersValues() []TypeBasicBuildStepPropertiesUpdateParameters {
	return original.PossibleTypeBasicBuildStepPropertiesUpdateParametersValues()
}
func PossibleTypeBasicQueueBuildRequestValues() []TypeBasicQueueBuildRequest {
	return original.PossibleTypeBasicQueueBuildRequestValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleWebhookActionValues() []WebhookAction {
	return original.PossibleWebhookActionValues()
}
func PossibleWebhookStatusValues() []WebhookStatus {
	return original.PossibleWebhookStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
