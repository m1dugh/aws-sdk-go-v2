// Code generated by smithy-go-codegen DO NOT EDIT.

// Package cloud9 provides the API client, operations, and parameter types for AWS
// Cloud9.
//
// Cloud9 Cloud9 is a collection of tools that you can use to code, build, run,
// test, debug, and release software in the cloud. For more information about
// Cloud9, see the Cloud9 User Guide (https://docs.aws.amazon.com/cloud9/latest/user-guide)
// . Cloud9 supports these operations:
//   - CreateEnvironmentEC2 : Creates an Cloud9 development environment, launches
//     an Amazon EC2 instance, and then connects from the instance to the environment.
//   - CreateEnvironmentSSH : Creates an Cloud9 SSH development environment.
//   - CreateEnvironmentMembership : Adds an environment member to an environment.
//   - DeleteEnvironment : Deletes an environment. If an Amazon EC2 instance is
//     connected to the environment, also terminates the instance.
//   - DeleteEnvironmentMembership : Deletes an environment member from an
//     environment.
//   - DescribeEnvironmentMemberships : Gets information about environment members
//     for an environment.
//   - DescribeEnvironments : Gets information about environments.
//   - DescribeEnvironmentStatus : Gets status information for an environment.
//   - ListEnvironments : Gets a list of environment identifiers.
//   - ListTagsForResource : Gets the tags for an environment.
//   - TagResource : Adds tags to an environment.
//   - UntagResource : Removes tags from an environment.
//   - UpdateEnvironment : Changes the settings of an existing environment.
//   - UpdateEnvironmentMembership : Changes the settings of an existing
//     environment member for an environment.
package cloud9
