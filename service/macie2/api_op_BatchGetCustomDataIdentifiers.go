// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about one or more custom data identifiers.
func (c *Client) BatchGetCustomDataIdentifiers(ctx context.Context, params *BatchGetCustomDataIdentifiersInput, optFns ...func(*Options)) (*BatchGetCustomDataIdentifiersOutput, error) {
	if params == nil {
		params = &BatchGetCustomDataIdentifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetCustomDataIdentifiers", params, optFns, c.addOperationBatchGetCustomDataIdentifiersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetCustomDataIdentifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetCustomDataIdentifiersInput struct {

	// An array of custom data identifier IDs, one for each custom data identifier to
	// retrieve information about.
	Ids []string

	noSmithyDocumentSerde
}

type BatchGetCustomDataIdentifiersOutput struct {

	// An array of objects, one for each custom data identifier that matches the
	// criteria specified in the request.
	CustomDataIdentifiers []types.BatchGetCustomDataIdentifierSummary

	// An array of custom data identifier IDs, one for each custom data identifier that
	// was specified in the request but doesn't correlate to an existing custom data
	// identifier.
	NotFoundIdentifierIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetCustomDataIdentifiersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetCustomDataIdentifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetCustomDataIdentifiers{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetCustomDataIdentifiers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opBatchGetCustomDataIdentifiers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "BatchGetCustomDataIdentifiers",
	}
}
