# SDK GO Maintenance 

## Introduction

The GO SDK is an automatically generated library using a highly customized version of the [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator). It allows developers to interact with the public MongoDB [Atlas V2 API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/).

Developers should be aware that the SDK is **generated** and will be automatically updated based on changes to the OpenAPI definition. 
Manual changes to the generated code should be avoided, as they will be overwritten during the next generation process.

The [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) is a powerful tool that takes an OpenAPI Specification (OAS) file (formerly known as Swagger Specification). 
By using the OpenAPI Generator, the Atlas GO SDK can quickly adapt to changes in the MongoDB [Atlas V2 API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/) and provide an up-to-date and consistent interface for Atlas users.

## SDK Input

The SDK fetches the public MongoDB Atlas V2 API specification to generate the client library. This specification, written in the OpenAPI Specification format, serves as the blueprint for generating the SDK code. It defines all the available endpoints, request and response structures, and other relevant details required to interact with the MongoDB Atlas API.

Developers can update the SDK by running:

- Generate SDK (Generates an official SDK update PR)
- Generate Dev SDK (Generates SDK containing feature preview changes)

See GitHub workflows for more information:

https://github.com/mongodb/atlas-sdk-go/actions

> NOTE: When running the job manually please leave the required branch name argument unchanged during manual runs.

## SDK Output

The output of the SDK is the Golang source code, which is stored in the `./admin` folder. The generated code provides a set of functions and data structures that correspond to the API endpoints and data models defined in the MongoDB Atlas API specification.

The SDK output should not be modified directly. Instead, developers should make any necessary changes or customizations in:

- OpenAPI
- SDK Specific Mustache templates

For more information please refer to the [SDK Tools Documentation](./tools)

## Performing SDK Manual Changes

Although manual changes to the SDK are discouraged for generated content, sdk contains also a manually written codebase.
For example `./auth` API is currently maintained by the team. 

## SDK Tools

The OpenAPI Generator configuration and templates used for SDK generation are available under the `./tools` folder. 
These tools include various configuration files and templates that influence the code generation process.
SDK tools also contain an OpenAPI transformation engine that simplifies our OpenAPI file for SDK generation purposes.

For more information please refer to the [SDK Tools Documentation](./tools) 

## SDK Releases

The SDK is automatically released for each pull request containing generated code. 
This means that any updates or enhancements made to the MongoDB Atlas API will be reflected in the SDK with each new release automatically.
Each pull request containing generated code also includes review instructions. 

> NOTE: SDK update PRs require team reviews before the merge

See [./tools/releaser](./tools/releaser) for more information.

## Examples

To help developers get started quickly with the SDK, a number of runnable examples are included in the `./examples` folder. These examples demonstrate various use cases and interactions with the MongoDB Atlas API, showcasing how to perform common tasks using the SDK.

> NOTE: Examples are run manually and not covered by the e2e tests we have.

## Documentation

The documentation for the SDK is stored under the docs folder. It is divided into two types: manual documentation and automatically generated documentation.

### Manual Documentation

Manual documentation is stored in `./docs/doc_{number}` files. This documentation is written manually and covers aspects of the SDK that cannot be automatically generated from the OpenAPI Specification.

The `./docs/doc_{number}*.md` files should be used to provide detailed explanations, usage instructions, and best practices related to the SDK. Changes made to the manual documentation require review and approval by the docs team, ensuring that the documentation remains accurate and up-to-date.

### Automatically Generated Documentation

The `./docs/docs` folder contains documentation that is automatically generated by the OpenAPI Generator. This documentation is directly derived from the MongoDB Atlas API specification and reflects the available endpoints, request parameters, and response structures along with the API examples.
The automatically generated documentation serves as a reference for developers, providing insights into the available API operations and their expected inputs and outputs.

## Tracking Upstream Issues

To maintain traceability and streamline issue management, each upstream issue automatically generates an internal Jira.

## Tracking PR Reviews

Currently, apix-1 GitHub team is notified for each auto-update PR of the SDK. Each PR requires a team review.