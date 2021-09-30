/*
Copyright 2021 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/typed/sources/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSourcesV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSourcesV1alpha1) AWSCloudWatchLogsSources(namespace string) v1alpha1.AWSCloudWatchLogsSourceInterface {
	return &FakeAWSCloudWatchLogsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCloudWatchSources(namespace string) v1alpha1.AWSCloudWatchSourceInterface {
	return &FakeAWSCloudWatchSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCodeCommitSources(namespace string) v1alpha1.AWSCodeCommitSourceInterface {
	return &FakeAWSCodeCommitSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCognitoIdentitySources(namespace string) v1alpha1.AWSCognitoIdentitySourceInterface {
	return &FakeAWSCognitoIdentitySources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSCognitoUserPoolSources(namespace string) v1alpha1.AWSCognitoUserPoolSourceInterface {
	return &FakeAWSCognitoUserPoolSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSDynamoDBSources(namespace string) v1alpha1.AWSDynamoDBSourceInterface {
	return &FakeAWSDynamoDBSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSKinesisSources(namespace string) v1alpha1.AWSKinesisSourceInterface {
	return &FakeAWSKinesisSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSPerformanceInsightsSources(namespace string) v1alpha1.AWSPerformanceInsightsSourceInterface {
	return &FakeAWSPerformanceInsightsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSS3Sources(namespace string) v1alpha1.AWSS3SourceInterface {
	return &FakeAWSS3Sources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSSNSSources(namespace string) v1alpha1.AWSSNSSourceInterface {
	return &FakeAWSSNSSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AWSSQSSources(namespace string) v1alpha1.AWSSQSSourceInterface {
	return &FakeAWSSQSSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AzureActivityLogsSources(namespace string) v1alpha1.AzureActivityLogsSourceInterface {
	return &FakeAzureActivityLogsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AzureBlobStorageSources(namespace string) v1alpha1.AzureBlobStorageSourceInterface {
	return &FakeAzureBlobStorageSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AzureEventGridSources(namespace string) v1alpha1.AzureEventGridSourceInterface {
	return &FakeAzureEventGridSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AzureEventHubSources(namespace string) v1alpha1.AzureEventHubSourceInterface {
	return &FakeAzureEventHubSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) AzureQueueStorageSources(namespace string) v1alpha1.AzureQueueStorageSourceInterface {
	return &FakeAzureQueueStorageSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) GoogleCloudAuditLogsSources(namespace string) v1alpha1.GoogleCloudAuditLogsSourceInterface {
	return &FakeGoogleCloudAuditLogsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) GoogleCloudBillingSources(namespace string) v1alpha1.GoogleCloudBillingSourceInterface {
	return &FakeGoogleCloudBillingSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) GoogleCloudPubSubSources(namespace string) v1alpha1.GoogleCloudPubSubSourceInterface {
	return &FakeGoogleCloudPubSubSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) GoogleCloudRepositoriesSources(namespace string) v1alpha1.GoogleCloudRepositoriesSourceInterface {
	return &FakeGoogleCloudRepositoriesSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) GoogleCloudStorageSources(namespace string) v1alpha1.GoogleCloudStorageSourceInterface {
	return &FakeGoogleCloudStorageSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) HTTPPollerSources(namespace string) v1alpha1.HTTPPollerSourceInterface {
	return &FakeHTTPPollerSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) OCIMetricsSources(namespace string) v1alpha1.OCIMetricsSourceInterface {
	return &FakeOCIMetricsSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) SalesforceSources(namespace string) v1alpha1.SalesforceSourceInterface {
	return &FakeSalesforceSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) SlackSources(namespace string) v1alpha1.SlackSourceInterface {
	return &FakeSlackSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) TwilioSources(namespace string) v1alpha1.TwilioSourceInterface {
	return &FakeTwilioSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) WebhookSources(namespace string) v1alpha1.WebhookSourceInterface {
	return &FakeWebhookSources{c, namespace}
}

func (c *FakeSourcesV1alpha1) ZendeskSources(namespace string) v1alpha1.ZendeskSourceInterface {
	return &FakeZendeskSources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSourcesV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
