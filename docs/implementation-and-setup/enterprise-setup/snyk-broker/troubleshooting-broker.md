# Troubleshooting Broker

{% hint style="info" %}
**Multi-tenant settings for regions**\
When installing, you must add a command in your script to set the Broker server URL for the region where your data is hosted. For the commands and URLs to use, see [Broker URLs](../../../snyk-data-and-governance/regional-hosting-and-data-residency.md#broker-server-urls).
{% endhint %}

This page has information and instructions for the following:

* [Logging with the Broker Client](troubleshooting-broker.md#logging-with-the-broker-client)
* Basic troubleshooting with the monitoring features, [Healthcheck](troubleshooting-broker.md#monitoring-healthcheck) and [Systemcheck](troubleshooting-broker.md#monitoring-systemcheck)
* [Troubleshooting Standalone Broker](troubleshooting-broker.md#troubleshooting-standalone-broker)
* [Support of big manifest files (> 1Mb) for GitHub and GitHub Enterprise](troubleshooting-broker.md#support-of-big-manifest-files-greater-than-1mb-for-github-and-github-enterprise)

For more comprehensive troubleshooting information, see [Broker Troubleshooting FAQs](https://support.snyk.io/s/article/Broker-Troubleshooting).

## Logging with the Broker Client

By default, the log level of the Broker is set to INFO. All SCM responses, regardless of HTTP status code, are logged by the Broker Client. Set the following environment variables to alter the logging behavior:

| Key               | Default | Notes                                                          |
| ----------------- | ------- | -------------------------------------------------------------- |
| LOG\_LEVEL        | info    | Set to "debug" for all logs.                                   |
| LOG\_ENABLE\_BODY | false   | Set to "true" to include the response body in the Client logs. |

To keep the logs concise in normal operation, Snyk produces minimal information on the INFO level, tracking the requests coming from Snyk into the client as well as the downstream request made to the targeted system, Github, Gitlab, JIRA, and so on, and logging the url hit and the response code received.\\

When you set `LOG_INFO_VERBOSE="true"`, the environment variable will add the headers in these log lines without requiring that you use debug.

{% hint style="warning" %}
If you override the default logging, some logs may be provided by other processes such as API requests, and may list credentials. Before you send any Broker logs with increased logging enabled, check for any passwords or tokens and redact them in bulk.
{% endhint %}

## Monitoring: Healthcheck

Broker exposes an endpoint at `/healthcheck`, which can be used to monitor the health of the running application. This endpoint responds with status code `200 OK` when the internal request is successful and returns `{ ok: true }` in the response body.

In the case of the Broker Client, this endpoint also reports on the status of the Broker websocket connection. If the websocket connection is not open, this endpoint responds with status code `500 Internal Server Error` and `{ ok: false }` in the response body.

This status can be tested by connecting to the Broker and running http://localhost:8000/healthcheck with the default settings.

To change the location of the healthcheck endpoint, you can specify an alternative path in an environment variable:

```dockerfile
ENV BROKER_HEALTHCHECK_PATH /path/to/healthcheck
```

## Monitoring: Systemcheck

The Broker Client exposes an endpoint at `/systemcheck`, which can be used to validate the connectivity and credentials of the brokered service, Git or another SCM, or the brokered container registry . This endpoint causes the Broker Client to make a request to a preconfigured URL and report on the success of the request. The supported configuration is:

* `BROKER_CLIENT_VALIDATION_URL` - the URL to which the request will be made.
* `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER` - \[optional] the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_BASIC_AUTH`.
* `BROKER_CLIENT_VALIDATION_BASIC_AUTH` - \[optional] the basic auth credentials (`username:password`) to be base64 encoded and placed in the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER`.
* `BROKER_CLIENT_VALIDATION_METHOD` - \[optional] the HTTP method of the request (default is `GET`).
* `BROKER_CLIENT_VALIDATION_TIMEOUT_MS` - \[optional] the request timeout in milliseconds (default is 5000 ms).

This endpoint responds with status code `200 OK` when the internal request is successful, and returns `{ ok: true }` in the response body, one object in the array per credential; see [Credential pooling](classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md). If the internal request fails, this endpoint responds with status code `500 Internal Server Error` and `{ ok: false }` in the response body.

This status can be tested by connecting to the Broker and running [http://localhost:8000/systemcheck](http://localhost:8000/systemcheck) with the default settings.

Example that enables the `/systemcheck` capability to verify connectivity between broker and Nexus:\
`-e BROKER_CLIENT_VALIDATION_URL=https://[username:password]@acme.com/service/rest/v1/status[/check] /`\
`snyk/broker:nexus`

To change the location of the systemcheck endpoint, you can specify an alternative path in an environment variable:

```dockerfile
ENV BROKER_SYSTEMCHECK_PATH /path/to/systemcheck
```

{% hint style="info" %}
Snyk Broker does not support authentication with mTLS method.
{% endhint %}

## Troubleshooting Standalone Broker

If after running the Broker there is still an error connecting to the on-premise Git, use the following troubleshooting steps.

1. Ensure there are relevant logs in the Broker container. To do this, attempt to connect to the on-premise Git. Do this by navigating to the Integrations tab and attempt to import. This generates a log in the Broker logs.
2. Review the logs of the container. This can be done on Docker by running `docker logs <container id>`
3. Review the logs to see where the problem is occurring.

### Common problems with Standalone Broker

* If there is no log after performing the preceding steps, ensure that the customer has the correct Broker token. If so, ensure that the websocket has been established. Some firewalls will block this
* Review the HTTP code in the request to the on-premise Git.
  * 404 - Not found - Ensure correct information in the Docker run command.
  * 401/403 - Check credentials.
  * If there is any reference to SSL, this can be caused by a self-signed certificate. Ensure you have either mounted the correct certificate or use the flag `-e NODE_TLS_REJECT_UNAUTHORIZED=0.`

### Testing connectivity for Standalone Broker

The Broker and the agents do not have `curl` in their image. To test connectivity to an agent or an endpoint like a Container registry or SCM, you can use the following commands:

```
#start node
node

#test a url with http
http = require("http")
http.get("<URL_HERE>", res => {console.log(`statusCode: ${res.statusCode}`)})

#test a url with https
https = require("https")
https.get('<URL_HERE>', res => {console.log(`statusCode: ${res.statusCode}`)})
```

## **Support of big manifest files (> 1Mb) for GitHub and GitHub Enterprise**

Open Fix/Upgrade PRs or PR/recurring tests may fail because of failure in fetching big manifest files (> 1Mb). To address this issue, follow either the [Docker](classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup.md) or [Helm](classic-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/snyk-open-source-scans-sca-of-large-manifest-files-helm-setup.md) instructions to allow large manifest files. Containers go down when you log out of the host

If your containers go down, along with the Broker ecosystem, when you detach from their host, run the following to ensure the containers stay online when you log out:

`loginctl enable-linger $(whoami)`
