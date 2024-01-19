/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Plan.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Plan) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rule); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

				"v1beta1", "Vault", "VaultList")
			if err !=
				nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule[i3].TargetVaultName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Rule[i3].TargetVaultNameRef,
				Selector:     mg.Spec.ForProvider.Rule[i3].TargetVaultNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Rule[i3].TargetVaultName")
		}
		mg.Spec.ForProvider.Rule[i3].TargetVaultName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Rule[i3].TargetVaultNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Rule); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

				"v1beta1", "Vault", "VaultList")
			if err !=
				nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Rule[i3].TargetVaultName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Rule[i3].TargetVaultNameRef,
				Selector:     mg.Spec.InitProvider.Rule[i3].TargetVaultNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Rule[i3].TargetVaultName")
		}
		mg.Spec.InitProvider.Rule[i3].TargetVaultName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Rule[i3].TargetVaultNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Selection.
func (mg *Selection) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IAMRoleArnRef,
			Selector:     mg.Spec.ForProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoleArn")
	}
	mg.Spec.ForProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Plan", "PlanList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlanID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.PlanIDRef,
			Selector:     mg.Spec.ForProvider.PlanIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlanID")
	}
	mg.Spec.ForProvider.PlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlanIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1", "Role", "RoleList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IAMRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.IAMRoleArnRef,
			Selector:     mg.Spec.InitProvider.IAMRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IAMRoleArn")
	}
	mg.Spec.InitProvider.IAMRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IAMRoleArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Plan", "PlanList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PlanID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.PlanIDRef,
			Selector:     mg.Spec.InitProvider.PlanIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlanID")
	}
	mg.Spec.InitProvider.PlanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlanIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Vault.
func (mg *Vault) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

			"v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
			Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io",

			"v1beta1", "Key", "KeyList")
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.KMSKeyArnRef,
			Selector:     mg.Spec.InitProvider.KMSKeyArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyArn")
	}
	mg.Spec.InitProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultLockConfiguration.
func (mg *VaultLockConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
			Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BackupVaultNameRef,
			Selector:     mg.Spec.InitProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BackupVaultName")
	}
	mg.Spec.InitProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackupVaultNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultNotifications.
func (mg *VaultNotifications) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
			Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
			Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BackupVaultNameRef,
			Selector:     mg.Spec.InitProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BackupVaultName")
	}
	mg.Spec.InitProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackupVaultNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

			"v1beta1", "Topic", "TopicList")
		if err !=

			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SnsTopicArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.SnsTopicArnRef,
			Selector:     mg.Spec.InitProvider.SnsTopicArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SnsTopicArn")
	}
	mg.Spec.InitProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SnsTopicArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VaultPolicy.
func (mg *VaultPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.BackupVaultNameRef,
			Selector:     mg.Spec.ForProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BackupVaultName")
	}
	mg.Spec.ForProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BackupVaultNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("backup.aws.upbound.io",

			"v1beta1", "Vault", "VaultList")
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BackupVaultName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.BackupVaultNameRef,
			Selector:     mg.Spec.InitProvider.BackupVaultNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.BackupVaultName")
	}
	mg.Spec.InitProvider.BackupVaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BackupVaultNameRef = rsp.ResolvedReference

	return nil
}
