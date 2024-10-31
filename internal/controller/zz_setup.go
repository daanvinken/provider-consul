// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	policy "github.com/daanvinken/provider-consul/internal/controller/acl/policy"
	role "github.com/daanvinken/provider-consul/internal/controller/acl/role"
	token "github.com/daanvinken/provider-consul/internal/controller/acl/token"
	node "github.com/daanvinken/provider-consul/internal/controller/consul/node"
	service "github.com/daanvinken/provider-consul/internal/controller/consul/service"
	query "github.com/daanvinken/provider-consul/internal/controller/prepared/query"
	providerconfig "github.com/daanvinken/provider-consul/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		role.Setup,
		token.Setup,
		node.Setup,
		service.Setup,
		query.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
