# API Docs Generator

This utility generates API documentation from the openapi specification files. See `config.yml` for the configuration options.

By default the utility will:
1. Fetch the latest rest specification from `api.snyk.io` and update the copy in the docs repo
2. Use the v1 specification in the docs repo as a source of truth for the API documentation
3. Generate the API documentation in the `docs/snyk-api/reference` directory based on the v1 and REST specs

## Usage

To generate the API documentation locally, run:

```bash
make run
```

## Development

To test the utility, run:
```bash
make test
```