package main

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type gcpProvider struct {
	bucketLocation string
	project        string
}

func (p *gcpProvider) Output(resourceName, stateID string) string {
	f := hclwrite.NewEmptyFile()
	root := f.Body()

	bucketBlock := root.AppendNewBlock("resource", []string{"google_storage_bucket", stateID}).Body()
	bucketBlock.SetAttributeValue("name", cty.StringVal(resourceName))

	if p.project != "" {
		bucketBlock.SetAttributeValue("project", cty.StringVal(p.project))
	}

	bucketBlock.SetAttributeValue("location", cty.StringVal(p.bucketLocation))
	bucketBlock.SetAttributeValue("storage_class", cty.StringVal("STANDARD"))
	bucketBlock.AppendNewline()
	bucketBlock.SetAttributeValue("uniform_bucket_level_access", cty.BoolVal(true))
	bucketBlock.SetAttributeValue("force_destroy", cty.BoolVal(false))
	bucketBlock.AppendNewline()

	versioning := bucketBlock.AppendNewBlock("versioning", nil).Body()
	versioning.SetAttributeValue("enabled", cty.BoolVal(true))

	return string(f.Bytes())
}

func (p *gcpProvider) CreateResourceNameSuffix(length int) string {
	return createGenericSuffix(length)
}

func (p *gcpProvider) Import(ctx context.Context, resourceName, stateID string) error {
	return run(ctx, "terraform", "import", "google_storage_bucket."+stateID, resourceName)
}
