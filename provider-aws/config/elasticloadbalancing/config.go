/*
Copyright 2021 Upbound Inc.
*/

package elasticloadbalancing

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for elasticloadbalancing group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_lb", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		r.References = config.References{
			"security_groups": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupRefs",
				SelectorFieldName: "SecurityGroupSelector",
			},
			"subnets": {
				Type:              "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
				RefFieldName:      "SubnetRefs",
				SelectorFieldName: "SubnetSelector",
			},
			"access_logs.bucket": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket",
			},
			"subnet_mapping.subnet_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.Subnet",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_lb_listener", func(r *config.Resource) {
		r.References = config.References{
			"load_balancer_arn": {
				Type: "LB",
			},
			"default_action.target_group_arn": {
				Type: "LBTargetGroup",
			},
			"default_action.forward.target_group.arn": {
				Type: "LBTargetGroup",
			},
		}
	})

	p.AddResourceConfigurator("aws_lb_target_group", func(r *config.Resource) {
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "name_prefix")
		if s, ok := r.TerraformResource.Schema["name"]; ok {
			s.Optional = false
			s.ForceNew = true
			s.Computed = false
		}
	})
	p.AddResourceConfigurator("aws_lb_target_group_attachment", func(r *config.Resource) {
		r.References = config.References{
			"target_group_arn": {
				Type: "LBTargetGroup",
			},
		}
	})
}
