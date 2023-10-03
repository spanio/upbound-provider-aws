package qldb

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the qldb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_qldb_stream", func(r *config.Resource) {
		r.References["kinesis_configuration.stream_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream",
			Extractor: common.PathTerraformIDExtractor,
		}
		r.References["ledger_name"] = config.Reference{
			Type:      "Ledger",
			Extractor: common.PathTerraformIDExtractor,
		}
	})
	p.AddResourceConfigurator("aws_qldb_ledger", func(r *config.Resource) {
		r.References["kms_key"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
	})
}
