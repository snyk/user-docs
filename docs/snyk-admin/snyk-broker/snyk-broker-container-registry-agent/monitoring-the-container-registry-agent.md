# Monitoring the Container Registry Agent

## **Healthcheck**

The Broker exposes an endpoint at `/healthcheck`, which can be used to monitor the health of the running application. This endpoint responds with status code `200 OK` when the internal request is successful, and returns `{ ok: true }` in the response body.

In the case of the Broker client, this endpoint also reports on the status of the Broker websocket connection. If the websocket connection is not open, this endpoint responds with status code `500 Internal Server Error` and `{ ok: false }` in the response body.

To change the location of the healthcheck endpoint, you can specify an alternative path via an environment variable:

```dockerfile
ENV BROKER_HEALTHCHECK_PATH /path/to/healthcheck
```

## **Systemcheck**

The Broker client exposes an endpoint at `/systemcheck`, which can be used to validate the connectivity and credentials of the brokered container registry. This endpoint causes the Broker client to make a request to a preconfigured URL, and report on the success of the request. The supported configuration is:

* `BROKER_CLIENT_VALIDATION_URL` - the URL to which the request will be made.
* `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER` - \[optional] the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_BASIC_AUTH`.
* `BROKER_CLIENT_VALIDATION_BASIC_AUTH` - \[optional] the basic auth credentials (`username:password`) to be base64 encoded and placed in the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER`.
* `BROKER_CLIENT_VALIDATION_METHOD` - \[optional] the HTTP method of the request (default is `GET`).
* `BROKER_CLIENT_VALIDATION_TIMEOUT_MS` - \[optional] the request timeout in milliseconds (default is 5000 ms).

This endpoint responds with status code `200 OK` when the internal request is successful, and returns `[{ ok: true, ... }]` in the response body (one object in the array per credential, see Credential Pooling). If the internal request fails, this endpoint responds with status code `500 Internal Server Error` and `[{ ok: false }, ...]` in the response body.

To change the location of the systemcheck endpoint, you can specify an alternative path via an environment variable:

```dockerfile
ENV BROKER_SYSTEMCHECK_PATH /path/to/systemcheck
```

## **Logging**

By default the log level of the Broker is set to INFO. All SCM responses regardless of HTTP status code will be logged by the Broker client. The following settings can be set in your environment variables to alter the logging behaviour:

| Key               | Default | Notes                                                         |
| ----------------- | ------- | ------------------------------------------------------------- |
| LOG\_LEVEL        | info    | Set to "debug" for all logs                                   |
| LOG\_ENABLE\_BODY | false   | Set to "true" to include the response body in the Client logs |
