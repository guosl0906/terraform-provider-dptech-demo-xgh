package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ExampleResource{}
var _ resource.ResourceWithImportState = &ExampleResource{}

func NewExampleResource() resource.Resource {
	return &ExampleResource{}
}

// ExampleResource defines the resource implementation.
type ExampleResource struct {
	client *Client
}

// ExampleResourceModel describes the resource data model.
type ExampleResourceModel struct {
	// Uuid_count types.String `tfsdk:"uuid_count"`
	Rsinfo RealServiceParameter `tfsdk:"rsinfo"`
}

type RealServiceParameter struct {
	Name                types.String `tfsdk:"name" json:"name"`
	Address             types.String `tfsdk:"address" json:"address"`
	Port                types.String `tfsdk:"port" json:"port"`
	Weight              types.String `tfsdk:"weight" json:"weight"`
	ConnectionLimit     types.String `tfsdk:"connection_limit" json:"connectionLimit"`
	ConnectionRateLimit types.String `tfsdk:"connection_rate_limit" json:"connectionRateLimit"`
	RecoveryTime        types.String `tfsdk:"recovery_time" json:"recoveryTime"`
	WarmTime            types.String `tfsdk:"warm_time" json:"warmTime"`
	Monitor             types.String `tfsdk:"monitor" json:"monitor"`
	MonitorList         types.String `tfsdk:"monitor_list" json:"monitorList"`
	LeastNumber         types.String `tfsdk:"least_number" json:"leastNumber"`
	Priority            types.String `tfsdk:"priority" json:"priority"`
	MonitorLog          types.String `tfsdk:"monitor_log" json:"monitorLog"`
	SimulTunnelsLimit   types.String `tfsdk:"simul_tunnels_limit" json:"simulTunnelsLimit"`
	CpuWeight           types.String `tfsdk:"cpu_weight" json:"cpuWeight"`
	MemoryWeight        types.String `tfsdk:"memory_weight" json:"memoryWeight"`
	State               types.String `tfsdk:"state" json:"state"`
	VsysName            types.String `tfsdk:"vsys_name" json:"vsysName"`
}

func (r *ExampleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	// resp.TypeName = req.ProviderTypeName + "_example"
	resp.TypeName = "dptech-demo_RealService"
}

func (r *ExampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"rsinfo": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"address": schema.StringAttribute{
						Required: true,
					},
					"port": schema.StringAttribute{
						Required: true,
					},
					"weight": schema.StringAttribute{
						Optional: true,
					},
					"connection_limit": schema.StringAttribute{
						Optional: true,
					},
					"connection_rate_limit": schema.StringAttribute{
						Optional: true,
					},
					"recovery_time": schema.StringAttribute{
						Optional: true,
					},
					"warm_time": schema.StringAttribute{
						Optional: true,
					},
					"monitor": schema.StringAttribute{
						Optional: true,
					},
					"monitor_list": schema.StringAttribute{
						Optional: true,
					},
					"least_number": schema.StringAttribute{
						Optional: true,
					},
					"priority": schema.StringAttribute{
						Optional: true,
					},
					"monitor_log": schema.StringAttribute{
						Optional: true,
					},
					"simul_tunnels_limit": schema.StringAttribute{
						Optional: true,
					},
					"cpu_weight": schema.StringAttribute{
						Optional: true,
					},
					"memory_weight": schema.StringAttribute{
						Optional: true,
					},
					"state": schema.StringAttribute{
						Optional: true,
					},
					"vsys_name": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r *ExampleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)

	if req.ProviderData == nil {
		return
	}
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ExampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ExampleResourceModel
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Trace(ctx, "creating a resource ")

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ExampleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " read Start")
	sendToweb_main(ctx, r.client, data.Rsinfo)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ExampleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " Update Start ************")
	sendToweb_main(ctx, r.client, data.Rsinfo)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ExampleResourceModel
	tflog.Info(ctx, " Delete Start")
	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *ExampleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
func sendToweb_main(ctx context.Context, c *Client, Rsinfo RealServiceParameter) {

	body, _ := json.Marshal(Rsinfo)
	targetUrl := c.HostURL + "/func/web_main/api/slb/adx_slb/adx_slb_rs/rsinfo"
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	respn, err := http.DefaultClient.Do(req)
	if err != nil {
		tflog.Info(ctx, " read Error"+err.Error())
	}
	defer respn.Body.Close()
}
