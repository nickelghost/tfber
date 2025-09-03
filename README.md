# tfber

Utility that allows quick adding of production-ready state storage to Terraform projects.

This generates Terraform code for resources that are required to store state in a specific cloud provider. The goal is to make it easier to set up TF state storage without having to add modules or really think about the storage's configuration - just add code to your project.

## Installation

You can simply install it via go's built in install command.

```bash
go install github.com/nickelghost/tfber@latest
```

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
