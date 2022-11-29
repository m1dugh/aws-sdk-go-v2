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

// Updates the settings for the sensitivity inspection template for an account.
func (c *Client) UpdateSensitivityInspectionTemplate(ctx context.Context, params *UpdateSensitivityInspectionTemplateInput, optFns ...func(*Options)) (*UpdateSensitivityInspectionTemplateOutput, error) {
	if params == nil {
		params = &UpdateSensitivityInspectionTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSensitivityInspectionTemplate", params, optFns, c.addOperationUpdateSensitivityInspectionTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSensitivityInspectionTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSensitivityInspectionTemplateInput struct {

	// The unique identifier for the Amazon Macie resource that the request applies to.
	//
	// This member is required.
	Id *string

	// A custom description of the template.
	Description *string

	// The managed data identifiers to explicitly exclude (not use) when analyzing
	// data. To exclude an allow list or custom data identifier that's currently
	// included by the template, update the values for the
	// SensitivityInspectionTemplateIncludes.allowListIds and
	// SensitivityInspectionTemplateIncludes.customDataIdentifierIds properties,
	// respectively.
	Excludes *types.SensitivityInspectionTemplateExcludes

	// The allow lists, custom data identifiers, and managed data identifiers to
	// include (use) when analyzing data.
	Includes *types.SensitivityInspectionTemplateIncludes

	noSmithyDocumentSerde
}

type UpdateSensitivityInspectionTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSensitivityInspectionTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSensitivityInspectionTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSensitivityInspectionTemplate{}, middleware.After)
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
	if err = addOpUpdateSensitivityInspectionTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSensitivityInspectionTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSensitivityInspectionTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateSensitivityInspectionTemplate",
	}
}
