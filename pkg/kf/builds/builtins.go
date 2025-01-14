// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builds

import build "github.com/knative/build/pkg/apis/build/v1alpha1"

const (
	clusterBuildTemplate = "ClusterBuildTemplate"
	buildTemplate        = "BuildTemplate"
)

// BuildpackTemplate gets the template spec for the bulidpack template.
func BuildpackTemplate() build.TemplateInstantiationSpec {
	return build.TemplateInstantiationSpec{
		Name: "buildpack",
		Kind: clusterBuildTemplate,
	}
}

// clusterBuiltins returns a list of all ClusterBuildTemplates
func clusterBuiltins() []build.TemplateInstantiationSpec {
	return []build.TemplateInstantiationSpec{
		BuildpackTemplate(),
	}
}
