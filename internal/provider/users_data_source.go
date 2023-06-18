package provider

import (
	"context"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/msharbaji/grpc-go-example/pkg/client"
)

var _ datasource.DataSource = &UsersDataSource{}

type UsersDataSource struct {
	client client.Client
}

type UsersDataSourceModel struct {
	ID    types.String          `tfsdk:"id"`
	Users []UserDataSourceModel `tfsdk:"users"`
}

func (u *UsersDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_users"
}

func (u *UsersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "The example provider schema",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "to avoid error: \"id\" conflicts with the existing attribute \"id\" in this block",
				Computed:            true,
				Required:            false,
				Optional:            false,
			},
			"users": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The ID of the user",
							Computed:            true,
							Required:            false,
							Optional:            false,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "The username of the user",
							Computed:            true,
							Required:            false,
							Optional:            false,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "The email of the user",
							Computed:            true,
							Required:            false,
							Optional:            false,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "The created_at of the user",
							Computed:            true,
							Required:            false,
							Optional:            false,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: "The updated_at of the user",
							Computed:            true,
							Required:            false,
							Optional:            false,
						},
					},
				},
			},
		},
	}
}

func (u *UsersDataSource) Read(ctx context.Context, _ datasource.ReadRequest, response *datasource.ReadResponse) {
	var state UsersDataSourceModel

	users, err := u.client.ListUsers(ctx)

	if err != nil {
		response.Diagnostics.AddError("failed to list Users", err.Error())
		return
	}

	for _, user := range users.GetUsers() {
		userState := UserDataSourceModel{
			ID:        types.StringValue(user.Id),
			Username:  types.StringValue(user.Username),
			Email:     types.StringValue(user.Email),
			CreatedAt: timestampToString(user.CreatedAt),
			UpdatedAt: timestampToString(user.UpdatedAt),
		}
		state.Users = append(state.Users, userState)
	}

	state.ID = types.StringValue(string(rune(rand.Int())))

	diags := response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (u *UsersDataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	grpcClient, ok := request.ProviderData.(client.Client)

	if !ok {
		response.Diagnostics.AddError("invalid _", "expected _")
		return
	}

	u.client = grpcClient
}

func NewUsersDataSource() datasource.DataSource {
	return &UsersDataSource{}
}
