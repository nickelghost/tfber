# tfber

Utility that allows quick adding of production-ready state storage to Terraform projects.

## Commands

- `generate` - create Terraform code and save it into a file
- `output` - create Terraform code and output it into standard output
- `import` - add already existing state storage as generated resources

## Supported providers

Currently only two providers are supported:

- aws
- gcp

The project is open to adding more providers.

## Importing

This functionality is meant for importing existing state resources into code. Please remember to provide the full `resource-name`.
