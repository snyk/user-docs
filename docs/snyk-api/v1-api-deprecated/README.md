# V1 API (deprecated)

The Snyk API is available to customers on [Business and Enterprise plans](https://snyk.io/plans) and allows you to programmatically integrate with Snyk.

### REST API

We are in the process of building a new, improved API (`https://api.snyk.io/rest`) built using the OpenAPI and JSON API standards. We welcome you to try it out as we shape and release endpoints until it, ultimately, becomes a full replacement for our current API.

Looking for our REST API docs? Please head over to [https://apidocs.snyk.io](https://apidocs.snyk.io)

### API vs CLI vs Snyk integration

The API detailed below has the ability to test a package for issues, as they are defined by Snyk. It is important to note that for many package managers, using this API will be less accurate than running the [Snyk CLI](https://snyk.io/docs/using-snyk) as part of your build pipe, or just using it locally on your package. The reason for this is that more than one package version fit the requirements given in manifest files. Running the CLI locally tests the actual deployed code, and has an accurate snapshot of the dependency versions in use, while the API can only infer it, with inferior accuracy. It should be noted that the Snyk CLI has the ability to output machine-readable JSON output (with the `--json` flag to `snyk test`).

A third option, is to allow Snyk access to your development flow via the existing [Snyk integrations](https://snyk.io/docs/). The advantage to this approach is having Snyk monitor every new pull request, and suggest fixes by opening new pull requests. This can be achieved either by integrating Snyk directly to your source code management (SCM) tool, or via a broker to allow greater security and auditability.

If those are not viable options, this API is your best choice.

### API url

The base URL for APIv1 endpoints is [https://api.snyk.io/v1/](https://api.snyk.io/v1/) ([https://api.eu.snyk.io/v1/](https://api.eu.snyk.io/v1/) for EU-MT, [https://api.au.snyk.io/v1/](https://api.au.snyk.io/v1/) for AU-MT)
