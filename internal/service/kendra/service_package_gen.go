// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceExperience,
			TypeName: "aws_kendra_experience",
			Name:     "Experience",
		},
		{
			Factory:  DataSourceFaq,
			TypeName: "aws_kendra_faq",
			Name:     "FAQ",
		},
		{
			Factory:  DataSourceIndex,
			TypeName: "aws_kendra_index",
			Name:     "Index",
		},
		{
			Factory:  DataSourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
			Name:     "Query Suggestions Block List",
		},
		{
			Factory:  DataSourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
			Name:     "Thesaurus",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_kendra_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceExperience,
			TypeName: "aws_kendra_experience",
			Name:     "Experience",
		},
		{
			Factory:  ResourceFaq,
			TypeName: "aws_kendra_faq",
			Name:     "FAQ",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceIndex,
			TypeName: "aws_kendra_index",
			Name:     "Index",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
			Name:     "Query Suggestions Block List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
			Name:     "Thesaurus",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kendra
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kendra.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*kendra.Options){
		kendra.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return kendra.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*kendra.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*kendra.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *kendra.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*kendra.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
