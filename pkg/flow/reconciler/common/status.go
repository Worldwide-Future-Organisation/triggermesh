/*
Copyright 2022 TriggerMesh Inc.

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

package common

import duckv1 "knative.dev/pkg/apis/duck/v1"

// CreateCloudEventAttributes returns CloudEvent attributes for the event types
// supported by the target.
func CreateCloudEventAttributes(source string, eventTypes []string) []duckv1.CloudEventAttributes {
	ceAttributes := make([]duckv1.CloudEventAttributes, len(eventTypes))

	for i, typ := range eventTypes {
		ceAttributes[i] = duckv1.CloudEventAttributes{
			Type:   typ,
			Source: source,
		}
	}

	return ceAttributes
}
