// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infrastructure

import (
	"github.com/gardener/gardener/extensions/pkg/controller/common"
	"github.com/gardener/gardener/extensions/pkg/controller/infrastructure"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type actuator struct {
	common.RESTConfigContext
	disableProjectedTokenMount bool
}

func (a *actuator) getOnmetalClientAndNamespaceFromInfraConfig(infra *extensionsv1alpha1.Infrastructure) (client.Client, string, error) {
	return nil, "", nil
}

// NewActuator creates a new infrastructure.Actuator.
func NewActuator(disableProjectedTokenMount bool) infrastructure.Actuator {
	return &actuator{
		disableProjectedTokenMount: disableProjectedTokenMount,
	}
}
