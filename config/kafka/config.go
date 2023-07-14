package kafka

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for kafka group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
		r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group",
		}
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["bootstrap_brokers"].(string); ok {
				conn["bootstrap_brokers"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_sasl_iam"].(string); ok {
				conn["bootstrap_brokers_sasl_iam"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_sasl_scram"].(string); ok {
				conn["bootstrap_brokers_sasl_scram"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_tls"].(string); ok {
				conn["bootstrap_brokers_tls"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_public"].(string); ok {
				conn["bootstrap_brokers_public"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_public_sasl_iam"].(string); ok {
				conn["bootstrap_brokers_public_sasl_iam"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_public_sasl_scram"].(string); ok {
				conn["bootstrap_brokers_public_sasl_scram"] = []byte(a)
			}
			if a, ok := attr["bootstrap_brokers_public_tls"].(string); ok {
				conn["bootstrap_brokers_public_tls"] = []byte(a)
			}
			if a, ok := attr["zookeeper_connect_string"].(string); ok {
				conn["zookeeper_connect_string"] = []byte(a)
			}
			if a, ok := attr["zookeeper_connect_string_tls"].(string); ok {
				conn["zookeeper_connect_string_tls"] = []byte(a)
			}
			return conn, nil
		}
	})
}
