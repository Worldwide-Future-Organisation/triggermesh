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

package routing

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// URLPath returns a URL path to route requests for the given object.
func URLPath(o metav1.Object) string {
	return "/" + o.GetNamespace() + "/" + o.GetName()
}
