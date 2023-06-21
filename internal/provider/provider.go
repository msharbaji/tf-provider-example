package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/msharbaji/grpc-go-example/pkg/client"
)

var (
	_ provider.Provider = &ExampleProvider{}
)

type ExampleProvider struct {
	version string
}

type ExampleProviderModel struct {
	Endpoint  string `tfsdk:"endpoint"`
	KeyID     string `tfsdk:"key_id"`
	SecretKey string `tfsdk:"secret_key"`
}

// New returns a new provider.Provider.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ExampleProvider{
			version: version,
		}
	}
}
func (e *ExampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "example"
	response.Version = e.version
}

func (e *ExampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The endpoint of the example provider",
				Description:         "The endpoint of the example provider",
				Required:            true,
			},
			"key_id": schema.StringAttribute{
				MarkdownDescription: "The key ID of the example provider",
				Description:         "The key ID of the example provider",
				Required:            true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "The secret key of the example provider",
				Description:         "The secret key of the example provider",
				Required:            true,
			},
		},
	}
}

func (e *ExampleProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	var data ExampleProviderModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	grpcClient, err := client.NewClient(data.Endpoint, data.KeyID, data.SecretKey)
	if err != nil {
		response.Diagnostics.AddError("failed to create _", "")
		return
	}

	response.DataSourceData = grpcClient
	response.ResourceData = grpcClient

}

func (e *ExampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUserDataSource,
		NewUsersDataSource,
	}
}

func (e *ExampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUserResource,
	}
}
