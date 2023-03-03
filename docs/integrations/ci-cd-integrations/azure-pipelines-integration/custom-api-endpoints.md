# Custom API endpoints

By default, the task uses the [https://snyk.io/api](https://snyk.io/api) endpoint. It is possible to configure Snyk to use a different endpoint by setting a `SNYK_API` environment variable in the pipeline, for example, `https://app.eu.io.api`:

<figure><img src="../../../.gitbook/assets/Screenshot 2022-07-22 at 17.36.54.png" alt="API environment variable"><figcaption><p>API environment variable</p></figcaption></figure>

Refer to [Configuration to connect to the Snyk API](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli#configuration-to-connect-to-the-snyk-api) in the CLI docs for more information about environment configuration.
