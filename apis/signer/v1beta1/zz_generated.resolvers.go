/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *SigningJob) ResolveReferences( // ResolveReferences of this SigningJob.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta1", "SigningProfile", "SigningProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProfileName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ProfileNameRef,
			Selector:     mg.Spec.ForProvider.ProfileNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProfileName")
	}
	mg.Spec.ForProvider.ProfileName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProfileNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta1", "SigningProfile", "SigningProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProfileName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ProfileNameRef,
			Selector:     mg.Spec.InitProvider.ProfileNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProfileName")
	}
	mg.Spec.InitProvider.ProfileName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProfileNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SigningProfilePermission.
func (mg *SigningProfilePermission) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta1", "SigningProfile", "SigningProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProfileName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ProfileNameRef,
			Selector:     mg.Spec.ForProvider.ProfileNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProfileName")
	}
	mg.Spec.ForProvider.ProfileName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProfileNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta1", "SigningProfile", "SigningProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProfileVersion),
			Extract:      resource.ExtractParamPath("version", true),
			Reference:    mg.Spec.ForProvider.ProfileVersionRef,
			Selector:     mg.Spec.ForProvider.ProfileVersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProfileVersion")
	}
	mg.Spec.ForProvider.ProfileVersion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProfileVersionRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("signer.aws.upbound.io", "v1beta1", "SigningProfile", "SigningProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ProfileVersion),
			Extract:      resource.ExtractParamPath("version", true),
			Reference:    mg.Spec.InitProvider.ProfileVersionRef,
			Selector:     mg.Spec.InitProvider.ProfileVersionSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ProfileVersion")
	}
	mg.Spec.InitProvider.ProfileVersion = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProfileVersionRef = rsp.ResolvedReference

	return nil
}
