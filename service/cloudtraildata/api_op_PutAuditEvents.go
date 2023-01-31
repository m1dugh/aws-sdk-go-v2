// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtraildata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtraildata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Ingests your application events into CloudTrail Lake. A required parameter,
// auditEvents, accepts the JSON records (also called payload) of events that you
// want CloudTrail to ingest. You can add up to 100 of these events (or up to 1 MB)
// per PutAuditEvents request.
func (c *Client) PutAuditEvents(ctx context.Context, params *PutAuditEventsInput, optFns ...func(*Options)) (*PutAuditEventsOutput, error) {
	if params == nil {
		params = &PutAuditEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAuditEvents", params, optFns, c.addOperationPutAuditEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAuditEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAuditEventsInput struct {

	// The JSON payload of events that you want to ingest. You can also point to the
	// JSON event payload in a file.
	//
	// This member is required.
	AuditEvents []types.AuditEvent

	// The ARN or ID (the ARN suffix) of a channel.
	//
	// This member is required.
	ChannelArn *string

	// A unique identifier that is conditionally required when the channel's resource
	// policy includes an external ID. This value can be any string, such as a
	// passphrase or account number.
	ExternalId *string

	noSmithyDocumentSerde
}

type PutAuditEventsOutput struct {

	// Lists events in the provided event payload that could not be ingested into
	// CloudTrail, and includes the error code and error message returned for events
	// that could not be ingested.
	//
	// This member is required.
	Failed []types.ResultErrorEntry

	// Lists events in the provided event payload that were successfully ingested into
	// CloudTrail.
	//
	// This member is required.
	Successful []types.AuditEventResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAuditEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAuditEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAuditEvents{}, middleware.After)
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
	if err = addOpPutAuditEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAuditEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAuditEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail-data",
		OperationName: "PutAuditEvents",
	}
}
