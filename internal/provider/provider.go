// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure PowerBIProvider satisfies various provider interfaces.
var _ provider.Provider = &PowerBIProvider{}

// PowerBIProvider defines the provider implementation.
type PowerBIProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// PowerBIProviderModel describes the provider data model.
type PowerBIProviderModel struct {
	BaseURL types.String `tfsdk:"base_url"`
}

func (p *PowerBIProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "powerbi"
	resp.Version = p.version
}

func (p *PowerBIProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				MarkdownDescription: "The base url for the Power BI API. Default to \"https://api.powerbi.com\"",
				Description:         "The base url for the Power BI API. Default to \"https://api.powerbi.com\"",
				Optional:            true,
			},
		},
	}
}

func (p *PowerBIProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data PowerBIProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new instance of the powerbiapi.Client with the specified base URL.
	client, err := getClient(data.BaseURL.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to create client", err.Error())
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *PowerBIProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPowerBIWorkspaceResource,
	}
}

func (p *PowerBIProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewPowerBIWorkspaceDataSource,
		NewPowerBIWorkspacePermissionsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PowerBIProvider{
			version: version,
		}
	}
}
