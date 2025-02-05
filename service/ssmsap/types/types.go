// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An SAP application registered with AWS Systems Manager for SAP.
type Application struct {

	// The Amazon Resource Name (ARN) of the Application Registry.
	AppRegistryArn *string

	// The Amazon Resource Name (ARN) of the application.
	Arn *string

	// The components of the application.
	Components []string

	// The ID of the application.
	Id *string

	// The time at which the application was last updated.
	LastUpdated *time.Time

	// The status of the application.
	Status ApplicationStatus

	// The status message.
	StatusMessage *string

	// The type of the application.
	Type ApplicationType

	noSmithyDocumentSerde
}

// The credentials of your SAP application.
type ApplicationCredential struct {

	// The type of the application credentials.
	//
	// This member is required.
	CredentialType CredentialType

	// The name of the SAP HANA database.
	//
	// This member is required.
	DatabaseName *string

	// The secret ID created in AWS Secrets Manager to store the credentials of the
	// SAP application.
	//
	// This member is required.
	SecretId *string

	noSmithyDocumentSerde
}

// The summary of the SAP application registered with AWS Systems Manager for SAP.
type ApplicationSummary struct {

	// The Amazon Resource Name (ARN) of the application.
	Arn *string

	// The ID of the application.
	Id *string

	// The tags on the application.
	Tags map[string]string

	// The type of the application.
	Type ApplicationType

	noSmithyDocumentSerde
}

// The SAP component of your application.
type Component struct {

	// The ID of the application.
	ApplicationId *string

	// The ID of the component.
	ComponentId *string

	// The type of the component.
	ComponentType ComponentType

	// The SAP HANA databases of the component.
	Databases []string

	// The hosts of the component.
	Hosts []Host

	// The time at which the component was last updated.
	LastUpdated *time.Time

	// The primary host of the component.
	PrimaryHost *string

	// The status of the component.
	Status ComponentStatus

	noSmithyDocumentSerde
}

// The summary of the component.
type ComponentSummary struct {

	// The ID of the application.
	ApplicationId *string

	// The ID of the component.
	ComponentId *string

	// The type of the component.
	ComponentType ComponentType

	// The tags of the component.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The SAP HANA database of the application registered with AWS Systems Manager
// for SAP.
type Database struct {

	// The ID of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the database.
	Arn *string

	// The ID of the component.
	ComponentId *string

	// The credentials of the database.
	Credentials []ApplicationCredential

	// The ID of the SAP HANA database.
	DatabaseId *string

	// The name of the database.
	DatabaseName *string

	// The type of the database.
	DatabaseType DatabaseType

	// The time at which the database was last updated.
	LastUpdated *time.Time

	// The primary host of the database.
	PrimaryHost *string

	// The SQL port of the database.
	SQLPort *int32

	// The status of the database.
	Status DatabaseStatus

	noSmithyDocumentSerde
}

// The summary of the database.
type DatabaseSummary struct {

	// The ID of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the database.
	Arn *string

	// The ID of the component.
	ComponentId *string

	// The ID of the database.
	DatabaseId *string

	// The type of the database.
	DatabaseType DatabaseType

	// The tags of the database.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A specific result obtained by specifying the name, value, and operator.
type Filter struct {

	// The name of the filter. Filter names are case-sensitive.
	//
	// This member is required.
	Name *string

	// The operator for the filter.
	//
	// This member is required.
	Operator FilterOperator

	// The filter values. Filter values are case-sensitive. If you specify multiple
	// values for a filter, the values are joined with an OR, and the request returns
	// all results that match any of the specified values
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Describes the properties of the Dedicated Host.
type Host struct {

	// The IP address of the Dedicated Host.
	HostIp *string

	// The name of the Dedicated Host.
	HostName *string

	// The role of the Dedicated Host.
	HostRole HostRole

	// The instance ID of the instance on the Dedicated Host.
	InstanceId *string

	noSmithyDocumentSerde
}

// The operations performed by AWS Systems Manager for SAP.
type Operation struct {

	// The end time of the operation.
	EndTime *time.Time

	// The ID of the operation.
	Id *string

	// The time at which the operation was last updated.
	LastUpdatedTime *time.Time

	// The properties of the operation.
	Properties map[string]*string

	// The Amazon Resource Name (ARN) of the operation.
	ResourceArn *string

	// The resource ID of the operation.
	ResourceId *string

	// The resource type of the operation.
	ResourceType *string

	// The start time of the operation.
	StartTime *time.Time

	// The status of the operation.
	Status OperationStatus

	// The status message of the operation.
	StatusMessage *string

	// The type of the operation.
	Type *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
