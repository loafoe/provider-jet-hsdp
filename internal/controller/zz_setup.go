/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	application "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/application"
	client "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/client"
	group "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/group"
	org "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/org"
	passwordpolicy "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/passwordpolicy"
	proposition "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/proposition"
	role "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/role"
	service "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/iam/service"
	providerconfig "github.com/crossplane-contrib/provider-jet-hsdp/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		client.Setup,
		group.Setup,
		org.Setup,
		passwordpolicy.Setup,
		proposition.Setup,
		role.Setup,
		service.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
