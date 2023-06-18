package provider

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/msharbaji/grpc-go-example/api/pb"
	"github.com/msharbaji/grpc-go-example/pkg/client"
)

var _ resource.Resource = &UserResource{}

var _ resource.ResourceWithConfigure = &UserResource{}

type UserResource struct {
	client client.Client
}

type UserResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Username  types.String `tfsdk:"username"`
	Email     types.String `tfsdk:"email"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (u *UserResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user"
}

func (u *UserResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "User resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The id of the user",
				Computed:    true,
				Required:    false,
				Optional:    false,
			},
			"username": schema.StringAttribute{
				Description: "The username of the user",
				Computed:    false,
				Required:    true,
				Optional:    false,
			},
			"email": schema.StringAttribute{
				Description: "The email of the user",
				Computed:    false,
				Required:    true,
				Optional:    false,
			},
			"created_at": schema.StringAttribute{
				Description: "The created_at of the user",
				Computed:    true,
				Required:    false,
				Optional:    false,
			},
			"updated_at": schema.StringAttribute{
				Description: "The updated_at of the user",
				Computed:    true,
				Required:    false,
				Optional:    false,
			},
		},
	}
}

func (u *UserResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data *UserResourceModel

	// Read terraform plan data
	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Create user
	user, err := u.client.CreateUser(ctx, data.Username.ValueString(), data.Email.ValueString())
	if err != nil {
		response.Diagnostics.AddError("failed to create user", err.Error())
		return
	}

	// Set user data
	data.ID = types.StringValue(user.Id)
	data.Username = types.StringValue(user.Username)
	data.Email = types.StringValue(user.Email)
	data.CreatedAt = timestampToString(user.CreatedAt)
	data.UpdatedAt = timestampToString(user.UpdatedAt)

	// save user data
	response.Diagnostics.Append(response.State.Set(ctx, data)...)

}

func (u *UserResource) Read(ctx context.Context, _ resource.ReadRequest, response *resource.ReadResponse) {
	var data *UserResourceModel

	response.Diagnostics.Append(response.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

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

	// Set user data
	data.ID = types.StringValue(user.Id)
	data.Username = types.StringValue(user.Username)
	data.Email = types.StringValue(user.Email)
	data.CreatedAt = timestampToString(user.CreatedAt)
	data.UpdatedAt = timestampToString(user.UpdatedAt)

	// save user data
	response.Diagnostics.Append(response.State.Set(ctx, data)...)

}

func (u *UserResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var data UserResourceModel

	// Read terraform plan data
	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	user := &pb.User{
		Id:        data.ID.ValueString(),
		Username:  data.Username.ValueString(),
		Email:     data.Email.ValueString(),
		UpdatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}

	// Update user
	updatedUser, err := u.client.UpdateUser(ctx, user)
	if err != nil {
		response.Diagnostics.AddError("failed to update user", err.Error())
		return
	}

	// Set user data
	data.ID = types.StringValue(updatedUser.Id)
	data.Username = types.StringValue(updatedUser.Username)
	data.Email = types.StringValue(updatedUser.Email)
	data.CreatedAt = timestampToString(updatedUser.CreatedAt)
	data.UpdatedAt = timestampToString(updatedUser.UpdatedAt)

	// Save user data
	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (u *UserResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data *UserResourceModel

	// Read terraform plan data
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	// Delete user
	var err error

	switch {
	case data.ID != types.StringNull():
		_, err = u.client.DeleteUser(ctx, data.ID.ValueString(), "id")
	case data.Email != types.StringNull():
		_, err = u.client.DeleteUser(ctx, data.Email.ValueString(), "email")
	case data.Username != types.StringNull():
		_, err = u.client.DeleteUser(ctx, data.Username.ValueString(), "username")
	default:
		response.Diagnostics.AddError("id, email or username must be provided", "")
	}

	if err != nil {
		response.Diagnostics.AddError("failed to delete user", err.Error())
		return
	}
}

func (u *UserResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	grpcClient, ok := request.ProviderData.(client.Client)
	if !ok {
		tflog.Error(ctx, "invalid provider data")
		response.Diagnostics.AddError("invalid provider data", "")
		return
	}

	u.client = grpcClient
}

func NewUserResource() resource.Resource {
	return &UserResource{}
}
