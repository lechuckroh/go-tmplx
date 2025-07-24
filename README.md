# go-tmplx

**go-tmplx** is a lightweight CLI tool for rendering Go template files with variables loaded from JSON, YAML, or environment variables.

It is designed for developers and DevOps engineers who need to dynamically generate configuration files, scripts, or any text files.

## Features

- Render templates using Go's powerful `text/template` syntax.
- Load variables from multiple sources: JSON, YAML, and environment variables.
- Merge variables from different sources with priority rules.
- Simple and fast CLI interface, ideal for automation in CI/CD pipelines.

## Example

```bash
tmplx -t config.tmpl -v values.yaml -o config.yaml
```
