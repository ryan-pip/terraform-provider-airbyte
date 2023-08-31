// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk"
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/ryan-pip/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DestinationLangchainDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationLangchainDataSource{}

func NewDestinationLangchainDataSource() datasource.DataSource {
	return &DestinationLangchainDataSource{}
}

// DestinationLangchainDataSource is the data source implementation.
type DestinationLangchainDataSource struct {
	client *sdk.SDK
}

// DestinationLangchainDataSourceModel describes the data model.
type DestinationLangchainDataSourceModel struct {
	Configuration DestinationLangchain `tfsdk:"configuration"`
	DestinationID types.String         `tfsdk:"destination_id"`
	Name          types.String         `tfsdk:"name"`
	WorkspaceID   types.String         `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationLangchainDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_langchain"
}

// Schema defines the schema for the data source.
func (r *DestinationLangchainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationLangchain DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"langchain",
							),
						},
						Description: `must be one of ["langchain"]`,
					},
					"embedding": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_langchain_embedding_fake": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"fake",
											),
										},
										Description: `must be one of ["fake"]`,
									},
								},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
							},
							"destination_langchain_embedding_open_ai": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"openai",
											),
										},
										Description: `must be one of ["openai"]`,
									},
									"openai_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
							"destination_langchain_update_embedding_fake": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"fake",
											),
										},
										Description: `must be one of ["fake"]`,
									},
								},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
							},
							"destination_langchain_update_embedding_open_ai": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"openai",
											),
										},
										Description: `must be one of ["openai"]`,
									},
									"openai_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Embedding configuration`,
					},
					"indexing": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_langchain_indexing_chroma_local_persistance": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"collection_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the collection to use.`,
									},
									"destination_path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to the directory where chroma files will be written. The files will be placed inside that local mount.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"chroma_local",
											),
										},
										Description: `must be one of ["chroma_local"]`,
									},
								},
								Description: `Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync.`,
							},
							"destination_langchain_indexing_doc_array_hnsw_search": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"destination_path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"DocArrayHnswSearch",
											),
										},
										Description: `must be one of ["DocArrayHnswSearch"]`,
									},
								},
								Description: `DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite.`,
							},
							"destination_langchain_indexing_pinecone": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"index": schema.StringAttribute{
										Computed:    true,
										Description: `Pinecone index to use`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pinecone",
											),
										},
										Description: `must be one of ["pinecone"]`,
									},
									"pinecone_environment": schema.StringAttribute{
										Computed:    true,
										Description: `Pinecone environment to use`,
									},
									"pinecone_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain.`,
							},
							"destination_langchain_update_indexing_chroma_local_persistance": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"collection_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the collection to use.`,
									},
									"destination_path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to the directory where chroma files will be written. The files will be placed inside that local mount.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"chroma_local",
											),
										},
										Description: `must be one of ["chroma_local"]`,
									},
								},
								Description: `Chroma is a popular vector store that can be used to store and retrieve embeddings. It will build its index in memory and persist it to disk by the end of the sync.`,
							},
							"destination_langchain_update_indexing_doc_array_hnsw_search": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"destination_path": schema.StringAttribute{
										Computed:    true,
										Description: `Path to the directory where hnswlib and meta data files will be written. The files will be placed inside that local mount. All files in the specified destination directory will be deleted on each run.`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"DocArrayHnswSearch",
											),
										},
										Description: `must be one of ["DocArrayHnswSearch"]`,
									},
								},
								Description: `DocArrayHnswSearch is a lightweight Document Index implementation provided by Docarray that runs fully locally and is best suited for small- to medium-sized datasets. It stores vectors on disk in hnswlib, and stores all other data in SQLite.`,
							},
							"destination_langchain_update_indexing_pinecone": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"index": schema.StringAttribute{
										Computed:    true,
										Description: `Pinecone index to use`,
									},
									"mode": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pinecone",
											),
										},
										Description: `must be one of ["pinecone"]`,
									},
									"pinecone_environment": schema.StringAttribute{
										Computed:    true,
										Description: `Pinecone environment to use`,
									},
									"pinecone_key": schema.StringAttribute{
										Computed: true,
									},
								},
								Description: `Pinecone is a popular vector store that can be used to store and retrieve embeddings. It is a managed service and can also be queried from outside of langchain.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Indexing configuration`,
					},
					"processing": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"chunk_overlap": schema.Int64Attribute{
								Computed:    true,
								Description: `Size of overlap between chunks in tokens to store in vector store to better capture relevant context`,
							},
							"chunk_size": schema.Int64Attribute{
								Computed:    true,
								Description: `Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)`,
							},
							"text_fields": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be used to calculate the embedding. All other fields are passed along as meta fields. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array.`,
							},
						},
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationLangchainDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationLangchainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationLangchainDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationLangchainRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationLangchain(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
