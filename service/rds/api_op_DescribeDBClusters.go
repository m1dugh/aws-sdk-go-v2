// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Returns information about Amazon Aurora DB clusters and Multi-AZ DB clusters.
// This API supports pagination. For more information on Amazon Aurora DB clusters,
// see  What is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ deployments with two readable standby DB instances
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide. This operation can also return information for
// Amazon Neptune DB instances and Amazon DocumentDB instances.
func (c *Client) DescribeDBClusters(ctx context.Context, params *DescribeDBClustersInput, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
	if params == nil {
		params = &DescribeDBClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusters", params, optFns, c.addOperationDescribeDBClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClustersInput struct {

	// The user-supplied DB cluster identifier or the Amazon Resource Name (ARN) of the
	// DB cluster. If this parameter is specified, information from only the specific
	// DB cluster is returned. This parameter isn't case-sensitive. Constraints:
	//
	// * If
	// supplied, must match an existing DBClusterIdentifier.
	DBClusterIdentifier *string

	// A filter that specifies one or more DB clusters to describe. Supported
	// filters:
	//
	// * clone-group-id - Accepts clone group identifiers. The results list
	// only includes information about the DB clusters associated with these clone
	// groups.
	//
	// * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	// Resource Names (ARNs). The results list only includes information about the DB
	// clusters identified by these ARNs.
	//
	// * domain - Accepts Active Directory
	// directory IDs. The results list only includes information about the DB clusters
	// associated with these domains.
	//
	// * engine - Accepts engine names. The results
	// list only includes information about the DB clusters for these engines.
	Filters []types.Filter

	// Optional Boolean parameter that specifies whether the output includes
	// information about clusters shared from other Amazon Web Services accounts.
	IncludeShared bool

	// An optional pagination token provided by a previous DescribeDBClusters request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the DescribeDBClusters action.
type DescribeDBClustersOutput struct {

	// Contains a list of DB clusters for the user.
	DBClusters []types.DBCluster

	// A pagination token that can be used in a later DescribeDBClusters request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusters{}, middleware.After)
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
	if err = addOpDescribeDBClustersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusters(options.Region), middleware.Before); err != nil {
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

// DescribeDBClustersAPIClient is a client that implements the DescribeDBClusters
// operation.
type DescribeDBClustersAPIClient interface {
	DescribeDBClusters(context.Context, *DescribeDBClustersInput, ...func(*Options)) (*DescribeDBClustersOutput, error)
}

var _ DescribeDBClustersAPIClient = (*Client)(nil)

// DescribeDBClustersPaginatorOptions is the paginator options for
// DescribeDBClusters
type DescribeDBClustersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClustersPaginator is a paginator for DescribeDBClusters
type DescribeDBClustersPaginator struct {
	options   DescribeDBClustersPaginatorOptions
	client    DescribeDBClustersAPIClient
	params    *DescribeDBClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClustersPaginator returns a new DescribeDBClustersPaginator
func NewDescribeDBClustersPaginator(client DescribeDBClustersAPIClient, params *DescribeDBClustersInput, optFns ...func(*DescribeDBClustersPaginatorOptions)) *DescribeDBClustersPaginator {
	if params == nil {
		params = &DescribeDBClustersInput{}
	}

	options := DescribeDBClustersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusters page.
func (p *DescribeDBClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBClusters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DBClusterAvailableWaiterOptions are waiter options for DBClusterAvailableWaiter
type DBClusterAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBClusterAvailableWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, DBClusterAvailableWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeDBClustersInput, *DescribeDBClustersOutput, error) (bool, error)
}

// DBClusterAvailableWaiter defines the waiters for DBClusterAvailable
type DBClusterAvailableWaiter struct {
	client DescribeDBClustersAPIClient

	options DBClusterAvailableWaiterOptions
}

// NewDBClusterAvailableWaiter constructs a DBClusterAvailableWaiter.
func NewDBClusterAvailableWaiter(client DescribeDBClustersAPIClient, optFns ...func(*DBClusterAvailableWaiterOptions)) *DBClusterAvailableWaiter {
	options := DBClusterAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBClusterAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBClusterAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBClusterAvailable waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *DBClusterAvailableWaiter) Wait(ctx context.Context, params *DescribeDBClustersInput, maxWaitDur time.Duration, optFns ...func(*DBClusterAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBClusterAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *DBClusterAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeDBClustersInput, maxWaitDur time.Duration, optFns ...func(*DBClusterAvailableWaiterOptions)) (*DescribeDBClustersOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeDBClusters(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for DBClusterAvailable waiter")
}

func dBClusterAvailableStateRetryable(ctx context.Context, input *DescribeDBClustersInput, output *DescribeDBClustersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-restore"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-parameters"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// DBClusterDeletedWaiterOptions are waiter options for DBClusterDeletedWaiter
type DBClusterDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// DBClusterDeletedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, DBClusterDeletedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeDBClustersInput, *DescribeDBClustersOutput, error) (bool, error)
}

// DBClusterDeletedWaiter defines the waiters for DBClusterDeleted
type DBClusterDeletedWaiter struct {
	client DescribeDBClustersAPIClient

	options DBClusterDeletedWaiterOptions
}

// NewDBClusterDeletedWaiter constructs a DBClusterDeletedWaiter.
func NewDBClusterDeletedWaiter(client DescribeDBClustersAPIClient, optFns ...func(*DBClusterDeletedWaiterOptions)) *DBClusterDeletedWaiter {
	options := DBClusterDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = dBClusterDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &DBClusterDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for DBClusterDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *DBClusterDeletedWaiter) Wait(ctx context.Context, params *DescribeDBClustersInput, maxWaitDur time.Duration, optFns ...func(*DBClusterDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for DBClusterDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *DBClusterDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeDBClustersInput, maxWaitDur time.Duration, optFns ...func(*DBClusterDeletedWaiterOptions)) (*DescribeDBClustersOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeDBClusters(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for DBClusterDeleted waiter")
}

func dBClusterDeletedStateRetryable(ctx context.Context, input *DescribeDBClustersInput, output *DescribeDBClustersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(DBClusters) == `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var errorType *types.DBClusterNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "creating"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "modifying"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "rebooting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("DBClusters[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "resetting-master-credentials"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeDBClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusters",
	}
}
