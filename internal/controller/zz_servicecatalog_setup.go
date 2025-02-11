// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	budgetresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/budgetresourceassociation"
	constraint "github.com/upbound/provider-aws/internal/controller/servicecatalog/constraint"
	portfolio "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolio"
	portfolioshare "github.com/upbound/provider-aws/internal/controller/servicecatalog/portfolioshare"
	principalportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/principalportfolioassociation"
	product "github.com/upbound/provider-aws/internal/controller/servicecatalog/product"
	productportfolioassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/productportfolioassociation"
	provisioningartifact "github.com/upbound/provider-aws/internal/controller/servicecatalog/provisioningartifact"
	serviceaction "github.com/upbound/provider-aws/internal/controller/servicecatalog/serviceaction"
	tagoption "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoption"
	tagoptionresourceassociation "github.com/upbound/provider-aws/internal/controller/servicecatalog/tagoptionresourceassociation"
)

// Setup_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		budgetresourceassociation.Setup,
		constraint.Setup,
		portfolio.Setup,
		portfolioshare.Setup,
		principalportfolioassociation.Setup,
		product.Setup,
		productportfolioassociation.Setup,
		provisioningartifact.Setup,
		serviceaction.Setup,
		tagoption.Setup,
		tagoptionresourceassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
