package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type awsProvider struct{}

func (p *awsProvider) Output(resourceName, stateID string) string {
	bucketTokens := hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: fmt.Appendf(nil, "aws_s3_bucket.%s.id", stateID)},
	}
	f := hclwrite.NewEmptyFile()

	rootBody := f.Body()

	bucketBody := rootBody.AppendNewBlock("resource", []string{"aws_s3_bucket", stateID}).Body()
	bucketBody.SetAttributeValue("bucket", cty.StringVal(resourceName))

	rootBody.AppendNewline()

	bucketVersioningBody := rootBody.AppendNewBlock(
		"resource", []string{"aws_s3_bucket_versioning", stateID},
	).Body()
	bucketVersioningBody.SetAttributeRaw("bucket", bucketTokens)
	bucketVersioningBody.AppendNewline()
	bucketVersioningConfigurationBody := bucketVersioningBody.AppendNewBlock("versioning_configuration", nil).Body()
	bucketVersioningConfigurationBody.SetAttributeValue("status", cty.StringVal("Enabled"))

	rootBody.AppendNewline()

	bucketPublicAccessBlockBody := rootBody.AppendNewBlock(
		"resource", []string{"aws_s3_bucket_public_access_block", stateID},
	).Body()
	bucketPublicAccessBlockBody.SetAttributeRaw("bucket", bucketTokens)
	bucketPublicAccessBlockBody.SetAttributeValue("block_public_acls", cty.BoolVal(true))
	bucketPublicAccessBlockBody.SetAttributeValue("block_public_policy", cty.BoolVal(true))
	bucketPublicAccessBlockBody.SetAttributeValue("ignore_public_acls", cty.BoolVal(true))
	bucketPublicAccessBlockBody.SetAttributeValue("restrict_public_buckets", cty.BoolVal(true))

	rootBody.AppendNewline()

	dynamodbTableBody := rootBody.AppendNewBlock("resource", []string{"aws_dynamodb_table", stateID}).Body()
	dynamodbTableBody.SetAttributeValue("name", cty.StringVal(resourceName))
	dynamodbTableBody.SetAttributeValue("billing_mode", cty.StringVal("PAY_PER_REQUEST"))
	dynamodbTableBody.SetAttributeValue("hash_key", cty.StringVal("LockID"))
	dynamodbTableBody.AppendNewline()

	dynamodbTableAttributeBody := dynamodbTableBody.AppendNewBlock("attribute", nil).Body()
	dynamodbTableAttributeBody.SetAttributeValue("name", cty.StringVal("LockID"))
	dynamodbTableAttributeBody.SetAttributeValue("type", cty.StringVal("S"))

	return string(f.Bytes())
}

func (p *awsProvider) CreateResourceNameSuffix(length int) string {
	return createGenericSuffix(length)
}

func (p *awsProvider) Import(ctx context.Context, resourceName, stateID string) error {
	if err := run(ctx, "terraform", "import", "aws_s3_bucket."+stateID, resourceName); err != nil {
		return err
	}

	return run(ctx, "terraform", "import", "aws_dynamodb_table."+stateID, resourceName)
}
