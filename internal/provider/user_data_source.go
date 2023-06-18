package provider

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/msharbaji/grpc-go-example/api/pb"
	"github.com/msharbaji/grpc-go-example/pkg/client"
)

var _ datasource.DataSourceWithValidateConfig = &UserDataSource{}

type UserDataSource struct {
	client client.Client
}

type UserDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	Email     types.String `tfsdk:"email"`
	Username  types.String `tfsdk:"username"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (u *UserDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user"
}
func (u *UserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "User data source",
		Description:         "User data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the user",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username of the user",
				Optional:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "The email of the user",
				Optional:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "The created_at of the user",
				Optional:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "The updated_at of the user",
				Optional:            true,
			},
		},
	}
}

func (u *UserDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data UserDataSourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	var user *pb.User
	var err error

	switch {
	case data.ID != types.StringNull():
		user, err = u.client.GetUser(ctx, data.ID.ValueString(), "id")
	case data.Email != types.StringNull():
		user, err = u.client.GetUser(ctx, data.Email.ValueString(), "email")
	case data.Username != types.StringNull():
		user, err = u.client.GetUser(ctx, data.Username.ValueString(), "username")
	default:
		response.Diagnostics.AddError("id, email or username must be provided", "")
	}

	if err != nil {
		response.Diagnostics.AddError("failed to get user", err.Error())
		return
	}

	// Set data
	data.ID = types.StringValue(user.Id)
	data.Username = types.StringValue(user.Username)
	data.Email = types.StringValue(user.Email)
	data.CreatedAt = timestampToString(user.CreatedAt)
	data.UpdatedAt = timestampToString(user.UpdatedAt)

	tflog.Trace(ctx, "read a data source")
	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}

func (u *UserDataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	grpcClient, ok := request.ProviderData.(client.Client)

	if !ok {
		response.Diagnostics.AddError("invalid provider data", "")
		return
	}

	u.client = grpcClient
}

func (u *UserDataSource) ValidateConfig(ctx context.Context, request datasource.ValidateConfigRequest, response *datasource.ValidateConfigResponse) {

	var data UserDataSourceModel
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if data == (UserDataSourceModel{}) {
		response.Diagnostics.AddError("data is not provided", "")
		return
	}

	if response.Diagnostics.HasError() {
		return
	}

	if data.ID == types.StringNull() && data.Email == types.StringNull() && data.Username == types.StringNull() {
		response.Diagnostics.AddError("id, email or username must be provided", "")
		return
	}
}

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

func timestampToString(timestamp *timestamp.Timestamp) types.String {
	if timestamp == nil {
		return types.StringNull()
	}

	timeValue := timestamp.AsTime()
	timeString := timeValue.Format(time.RFC3339)
	return types.StringValue(timeString)
}
