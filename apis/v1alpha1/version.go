// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VersionSpec defines the desired state of Version.
//

type VersionSpec struct {

// Only publish a version if the hash value matches the value that's specified.
// Use this option to avoid publishing a version if the function code has changed
// since you last updated it. You can get the hash for the version that you
// uploaded from the output of UpdateFunctionCode.
CodeSHA256 *string `json:"codeSHA256,omitempty"`
// A description for the version to override the description in the function
// configuration.
Description *string `json:"description,omitempty"`
FunctionEventInvokeConfig *PutFunctionEventInvokeConfigInput `json:"functionEventInvokeConfig,omitempty"`
// The name or ARN of the Lambda function.
// 
// Name formats
// 
//    * Function name - MyFunction.
// 
//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
// 
//    * Partial ARN - 123456789012:function:MyFunction.
// 
// The length constraint applies only to the full ARN. If you specify only the
// function name, it is limited to 64 characters in length.
FunctionName *string `json:"functionName,omitempty"`
FunctionRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"functionRef,omitempty"`
ProvisionedConcurrencyConfig *PutProvisionedConcurrencyConfigInput `json:"provisionedConcurrencyConfig,omitempty"`
// Only update the function if the revision ID matches the ID that's specified.
// Use this option to avoid publishing a version if the function configuration
// has changed since you last updated it.
RevisionID *string `json:"revisionID,omitempty"`
}

// VersionStatus defines the observed state of Version
type VersionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The instruction set architecture that the function supports. Architecture
// is a string array with one of the valid values. The default architecture
// value is x86_64.
	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty"`
	// The size of the function's deployment package, in bytes.
	// +kubebuilder:validation:Optional
	CodeSize *int64 `json:"codeSize,omitempty"`
	// The function's dead letter queue.
	// +kubebuilder:validation:Optional
	DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
	// The function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).
// Omitted from CloudTrail logs.
	// +kubebuilder:validation:Optional
	Environment *EnvironmentResponse `json:"environment,omitempty"`
	// The size of the function's /tmp directory in MB. The default value is 512,
// but can be any whole number between 512 and 10,240 MB. For more information,
// see Configuring ephemeral storage (console) (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-ephemeral-storage).
	// +kubebuilder:validation:Optional
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`
	// Connection settings for an Amazon EFS file system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html).
	// +kubebuilder:validation:Optional
	FileSystemConfigs []*FileSystemConfig `json:"fileSystemConfigs,omitempty"`
	// The function's Amazon Resource Name (ARN).
	// +kubebuilder:validation:Optional
	FunctionARN *string `json:"functionARN,omitempty"`
	// The function that Lambda calls to begin running your function.
	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty"`
	// The function's image configuration values.
	// +kubebuilder:validation:Optional
	ImageConfigResponse *ImageConfigResponse `json:"imageConfigResponse,omitempty"`
	// The ARN of the Key Management Service (KMS) customer managed key that's used
// to encrypt the following resources:
// 
//    * The function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption).
// 
//    * The function's Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html)
//    snapshots.
// 
//    * When used with SourceKMSKeyArn, the unzipped version of the .zip deployment
//    package that's used for function invocations. For more information, see
//    Specifying a customer managed key for Lambda (https://docs.aws.amazon.com/lambda/latest/dg/encrypt-zip-package.html#enable-zip-custom-encryption).
// 
//    * The optimized version of the container image that's used for function
//    invocations. Note that this is not the same key that's used to protect
//    your container image in the Amazon Elastic Container Registry (Amazon
//    ECR). For more information, see Function lifecycle (https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-lifecycle).
// 
// If you don't provide a customer managed key, Lambda uses an Amazon Web Services
// owned key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk)
// or an Amazon Web Services managed key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
	// +kubebuilder:validation:Optional
	KMSKeyARN *string `json:"kmsKeyARN,omitempty"`
	// The date and time that the function was last updated, in ISO-8601 format
// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	// +kubebuilder:validation:Optional
	LastModified *string `json:"lastModified,omitempty"`
	// The status of the last update that was performed on the function. This is
// first set to Successful after function creation completes.
	// +kubebuilder:validation:Optional
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty"`
	// The reason for the last update that was performed on the function.
	// +kubebuilder:validation:Optional
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty"`
	// The reason code for the last update that was performed on the function.
	// +kubebuilder:validation:Optional
	LastUpdateStatusReasonCode *string `json:"lastUpdateStatusReasonCode,omitempty"`
	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	// +kubebuilder:validation:Optional
	Layers []*Layer `json:"layers,omitempty"`
	// For Lambda@Edge functions, the ARN of the main function.
	// +kubebuilder:validation:Optional
	MasterARN *string `json:"masterARN,omitempty"`
	// The amount of memory available to the function at runtime.
	// +kubebuilder:validation:Optional
	MemorySize *int64 `json:"memorySize,omitempty"`
	// The type of deployment package. Set to Image for container image and set
// Zip for .zip file archive.
	// +kubebuilder:validation:Optional
	PackageType *string `json:"packageType,omitempty"`
	// The version of the Lambda function.
	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty"`
	// The function's execution role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty"`
	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
// Runtime is required if the deployment package is a .zip file archive. Specifying
// a runtime results in an error if you're deploying a function using a container
// image.
// 
// The following list includes deprecated runtimes. Lambda blocks creating new
// functions and updating existing functions shortly after each runtime is deprecated.
// For more information, see Runtime use after deprecation (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-deprecation-levels).
// 
// For a list of all currently supported runtimes, see Supported runtimes (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtimes-supported).
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty"`
	// The ARN of the signing job.
	// +kubebuilder:validation:Optional
	SigningJobARN *string `json:"signingJobARN,omitempty"`
	// The ARN of the signing profile version.
	// +kubebuilder:validation:Optional
	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
	// Set ApplyOn to PublishedVersions to create a snapshot of the initialized
// execution environment when you publish a function version. For more information,
// see Improving startup performance with Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html).
	// +kubebuilder:validation:Optional
	SnapStart *SnapStartResponse `json:"snapStart,omitempty"`
	// The current state of the function. When the state is Inactive, you can reactivate
// the function by invoking it.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
	// The reason for the function's current state.
	// +kubebuilder:validation:Optional
	StateReason *string `json:"stateReason,omitempty"`
	// The reason code for the function's current state. When the code is Creating,
// you can't invoke or modify the function.
	// +kubebuilder:validation:Optional
	StateReasonCode *string `json:"stateReasonCode,omitempty"`
	// The amount of time in seconds that Lambda allows a function to run before
// stopping it.
	// +kubebuilder:validation:Optional
	Timeout *int64 `json:"timeout,omitempty"`
	// The function's X-Ray tracing configuration.
	// +kubebuilder:validation:Optional
	TracingConfig *TracingConfigResponse `json:"tracingConfig,omitempty"`
	// The version of the Lambda function.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty"`
	// The function's networking configuration.
	// +kubebuilder:validation:Optional
	VPCConfig *VPCConfigResponse `json:"vpcConfig,omitempty"`
}

// Version is the Schema for the Versions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   VersionSpec   `json:"spec,omitempty"`
	Status VersionStatus `json:"status,omitempty"`
}

// VersionList contains a list of Version
// +kubebuilder:object:root=true
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Version `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Version{}, &VersionList{})
}