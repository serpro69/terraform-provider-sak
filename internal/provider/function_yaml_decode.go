package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = YamlDecodeFunction{}

func NewYamlDecodeFunction() function.Function {
	return &YamlDecodeFunction{}
}

type YamlDecodeFunction struct{}

func (f YamlDecodeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "yaml_decode"
}

func (f YamlDecodeFunction) Definition(_ context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Decode a YAML file containing one or multiple documents",
		MarkdownDescription: "Given a YAML text file containing one or multiple documents, will decode the file and return a tuple of object representations for each document.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "document",
				MarkdownDescription: "The YAML plaintext for a document",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (f YamlDecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var doc string

	resp.Error = req.Arguments.Get(ctx, &doc)
	if resp.Error != nil {
		return
	}

	tv, diags := decode(ctx, doc)
	if diags.HasError() {
		resp.Error = function.FuncErrorFromDiags(ctx, diags)
		return
	}

	dynamResp := types.DynamicValue(tv)
	resp.Error = resp.Result.Set(ctx, &dynamResp)
}
