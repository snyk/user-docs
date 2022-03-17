# Troubleshooting Broker

{% hint style="info" %}
For more comprehensive troubleshooting information, see [Broker Troubleshooting FAQs](https://support.snyk.io/hc/en-us/articles/4404288846353-Broker-Troubleshooting).
{% endhint %}

## Basic Broker Troubleshooting

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

Can be tested by connecting to the broker and running [http://localhost:8000/systemcheck](http://localhost:8000/systemcheck) (default settings).



## Advanced Broker Troubleshooting

### Standalone Broker

If after running the broker, there is still an error connecting to the on-premise Git use the following troubleshooting steps.

1. Ensure there are relevant logs in the broker container. To do this, you need to attempt to connect to the on-premise Git. Do this by navigating to the Integrations tab and attempt to import. This will generate a log in the broker logs.
2. Review the logs of the container. This can be done on docker by running `docker logs <container id>`
3. Review the logs to see where the problem is occurring.

#### Common problems:&#x20;

* If there is no log after performing the above steps, ensure that the customer has the correct broker token. If so, ensure that the websocket has been established. Some firewalls will block this&#x20;
* Review the HTTP code in the request to the on-premise Git.&#x20;
  * **404 - Not found** - Ensure correct information in docker run command&#x20;
  * **401/403** - Check credentials&#x20;
  * If there is any reference to SSL - This can be caused by a self-signed certificate. Ensure you have either mounted the correct certificate, or use the flag `-e NODE_TLS_REJECT_UNAUTHORIZED`

### Broker with Code Agent

![](https://lh3.googleusercontent.com/r\_qtONpOOEW35gdyoBcWDAiC6j04M76q8mh922SHor4bdNZdt83sj2kP7d5hbzYcWVXp4Q2hZEiCeAVOmcj4Bu1yFPdnyp3rK7kKeBK8DZEd9S133Xn3YdjddclVf5maEbP23Jor)

The best way to troubleshoot the broker with the code agent is to understand the communication flow. Traffic travels from Snyk > Broker Client > Code Agent > On-premise Git > Code Agent > Snyk.

The vast majority of problems with the code agent are due to traffic being interrupted at one of these points.

Like described Standalone Broker, in order to troubleshoot the code agent, you need to generate logs. Do this by attempting to import a repository.

1. Ensure that the broker is functioning correctly and you can list the repositories. If this does not work, please review the Standalone Broker troubleshooting steps.
2. If after attempting to import a repo, you see an error message Bundle Creation Failed, review the logs of the containers.
3. Start with the broker container. Run `docker logs <container id>`
4. Look for the string `snykgit` . This is the API call from the broker container to the code agent container. If you get anything other than a 200 code, then there is some problem with the communication between the broker and the code agent. Ensure you have the proper flags set in the docker run command. Also ensure you have set up the docker network
5. Review the logs of the code agent by running `docker logs <container id>`

#### Common problems with the code agent:

* Communication with the on-premise Git is not functioning. There will be a 404 error on the attempt to clone the code If there is any reference to SSL - This can be caused by a self-signed certificate. Ensure you have either mounted the correct certificate, or use the flag `-e NODE_TLS_REJECT_UNAUTHORIZED`
* If you see the message: `“Uploaded Repo”`, the code agent and broker are configured correctly. If there are still errors on the import log, contact Snyk Support.
