/*
Copyright 2023.

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

package uffizzicluster

import (
	"github.com/UffizziCloud/uffizzi-cluster-operator/controllers/constants"
	"github.com/UffizziCloud/uffizzi-cluster-operator/controllers/helm/build/vcluster"
	"testing"

	"github.com/UffizziCloud/uffizzi-cluster-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBuildVClusterHelmReleaseName(t *testing.T) {
	uCluster := &v1alpha1.UffizziCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "test",
		},
	}

	helmReleaseName := vcluster.BuildVClusterHelmReleaseName(uCluster)
	expectedHelmReleaseName := constants.UCLUSTER_NAME_PREFIX + uCluster.Name

	if helmReleaseName != expectedHelmReleaseName {
		t.Errorf("Expected helmReleaseName to be uffizzi-test, got %s", helmReleaseName)
	}
}