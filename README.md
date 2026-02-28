# tfber

Utility that allows quick adding of production-ready state storage to Terraform projects.

This generates Terraform code for resources that are required to store state in a specific cloud provider. The goal is to make it easier to set up TF state storage without having to add modules or really think about the storage's configuration - just add code to your project.

## Installation

You can simply install it via go's built in install command.

```bash
go install github.com/nickelghost/tfber@latest
```

## Quick start

```bash
# Generate an AWS state backend into state.tf
tfber generate -provider aws -resource-name my-tf-state

# Same thing, but print to stdout instead
tfber output -provider aws -resource-name my-tf-state

# GCP requires a bucket location
tfber generate -provider gcp -resource-name my-tf-state -gcp-bucket-location US

# Import existing resources into your Terraform state
tfber import -provider aws -resource-name my-tf-state -state-id my_tf_state
```

## Commands

| Command    | Description                                                  |
| ---------- | ------------------------------------------------------------ |
| `generate` | Create Terraform code and save it into a file                |
| `output`   | Create Terraform code and print it to standard output        |
| `import`   | Import already existing state storage as generated resources |

## Flags

### Common (all commands)

| Flag             | Description                               | Required      |
| ---------------- | ----------------------------------------- | ------------- |
| `-provider`      | State provider (`aws` or `gcp`)           | Yes           |
| `-resource-name` | Name of the resources inside the provider | Yes           |
| `-state-id`      | Terraform resource path ID                | `import` only |

### Generate / Output

| Flag                           | Default | Description                                 |
| ------------------------------ | ------- | ------------------------------------------- |
| `-resource-name-suffix`        | `true`  | Append a random suffix to the resource name |
| `-resource-name-suffix-length` | `6`     | Length of the random suffix                 |

### Generate only

| Flag         | Default    | Description                             |
| ------------ | ---------- | --------------------------------------- |
| `-file-name` | `state.tf` | Name of the generated file              |
| `-force`     | `false`    | Overwrite the file if it already exists |

### GCP (generate / output)

| Flag                   | Description                                     | Required |
| ---------------------- | ----------------------------------------------- | -------- |
| `-gcp-bucket-location` | Where the GCP bucket should live                | Yes      |
| `-gcp-project`         | What project the resources should be created in | No       |

## Supported providers

- **aws** — S3 bucket (with versioning and public access block) + DynamoDB lock table
- **gcp** — Cloud Storage bucket (with versioning and uniform access)

The project is open to adding more providers.

## Importing

This functionality is meant for importing existing state resources into code. The `-state-id` flag is required and should match the Terraform resource identifier you want to use. Remember to provide the full `-resource-name` as it exists in your cloud provider.
