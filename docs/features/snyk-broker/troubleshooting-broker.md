# Troubleshooting Broker

{% hint style="info" %}
For more comprehensive troubleshooting information, see [Broker Troubleshooting FAQs](https://support.snyk.io/hc/en-us/articles/4404288846353-Broker-Troubleshooting).
{% endhint %}

### Monitoring

#### Healthcheck

The Broker exposes an endpoint at `/healthcheck`, which can be used to monitor the health of the running application. This endpoint responds with status code `200 OK` when the internal request is successful, and returns `{ ok: true }` in the response body.

In the case of the Broker client, this endpoint also reports on the status of the Broker websocket connection. If the websocket connection is not open, this endpoint responds with status code `500 Internal Server Error` and `{ ok: false }` in the response body.

Can be tested by connecting to the broker and running [http://localhost:8000/healthcheck](http://localhost:8000/healthcheck) (default settings)

**Systemcheck**

The Broker client exposes an endpoint at `/systemcheck`, which can be used to validate the brokered service (Git or the like) connectivity and credentials. This endpoint causes the Broker client to make a request to a preconfigured URL, and report on the success of the request. The supported configuration is:

* `BROKER_CLIENT_VALIDATION_URL` - the URL to which the request will be made.
* `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER` - \[optional] the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_BASIC_AUTH`.
* `BROKER_CLIENT_VALIDATION_BASIC_AUTH` - \[optional] the basic auth credentials (`username:password`) to be base64 encoded and placed in the `Authorization` header value of the request. Mutually exclusive with `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER`.
* `BROKER_CLIENT_VALIDATION_METHOD` - \[optional] the HTTP method of the request (default is `GET`).
* `BROKER_CLIENT_VALIDATION_TIMEOUT_MS` - \[optional] the request timeout in milliseconds (default is 5000 ms).

This endpoint responds with status code `200 OK` when the internal request is successful, and returns `{ ok: true }` in the response body. If the internal request fails, this endpoint responds with status code `500 Internal Server Error` and `{ ok: false }` in the response body.

Can be tested by connecting to the broker and running  [http://localhost:8000/systemcheck](http://localhost:8000/systemcheck) (default settings).



