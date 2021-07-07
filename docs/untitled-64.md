# My broker is not working

\(See the [broker documentation](https://support.snyk.io/hc/en-us/articles/360004032397) for more details\)

## Overview

Issues with the broker can typically be classified into one of three categories:

* Connectivity issues between your Broker Client and the Broker Server at Snyk
* Connectivity issues between your Broker Client and the Code repo or Jira; or
* Authentication issues with your Code repo or Jira

## Diagnosing issues with the Broker

There are a few diagnostic tools available to you to help identify the issues with your broker. The broker client exposes a diagnostics endpoint to check on its internal status. It is available at `/healthcheck`, and can be customised with the `BROKER_HEALTHCHECK_PATH` environment variable.

Poking this endpoint responds with some metadata about the broker client’s configuration, as well as a report on the status of the WebSocket connection between the broker client and the broker server. The response status code is 200 in case the WebSocket connection is established, 500 otherwise.

### Expected responses

For a properly configured broker client and indicates that there is a successful connection with Snyk via [https://broker.snyk.io](https://broker.snyk.io/)

[http://localhost:8000/healthcheck](http://localhost:8000/healthcheck)

```text
{ 
	ok: true, 
	websocketConnectionOpen: true, 
	brokerServerUrl: "<https://broker.snyk.io>", 
	version: "4.36.0"
}
```

When the WebSocket connection cannot be established:

```text
{ 
	ok: false, 
	websocketConnectionOpen: false, 
	brokerServerUrl: "<https://maybe.a.bad.server.name>", 
	version: "4.36.0"
}
```

In case the WebSocket connection is not established, Broker client logs will provide further info on the root cause of the issue.

If nothing is returned, verify that your broker is running on the Broker client by running the following command to see what Docker images are running:

```text
➜ ~ docker ps
```

You should see something like the following, and if not then you may need to restart your broker:

```text
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
6c5b979ea531 snyk/broker:github-com "docker-entrypoint.s…" 2 months ago
```

## Using SystemCheck to verify Connectivity to your SCM

Similar to `/healthcheck` above you can access`/systemcheck` to verify connectivity to your SCM from the broker. An example output of a Broker having an issue connecting to your SCM could be:

[http://localhost:8000/systemcheck](http://localhost:8000/systemcheck)

```text
{"brokerClientValidationUrl":"<server>/api/v3/user","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"unable to get local issuer certificate"}
```

Or

```text
{"brokerClientValidationUrl":"https://my-scm-server.com","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"self signed certificate in certificate chain"}
```

Often there will be a good hint as to what the problem is in these errors. You may want to review [the Broker documentation](https://github.com/snyk/broker) for guidance on how to resolve common HTTPS or Self-signed certificate issues with the broker.

You can verify connectivity to your SCM using a simple curl command, for example:

```text
curl http://your-scm-hostname
```

or

```text
curl http://broker/systemcheck
```

## Elevating log level for diagnostics

Add the following line to your Broker docker startup script to enable a more verbose log level:

```text
-e LOG_LEVEL=info
```

NOTE: You must restart your broker container to be able to make these changes take place.

## Checking Connectivity from Broker to the SCM

* Let's say you get connectivity issues and Broker can't connect to the SCM/Jira.
  * If the /systemcheck message appears to be something similar to: **"error":"getaddrinfo ENOTFOUND** [int.SCMdomain.com](http://int.scmdomain.com/) [int.SCMdomain.com](http://int.scmdomain.com/)[:443](http://stash.caseware.com:443/)" it means that the broker is unable to resolve the domain, likely because it's in internal domain.
  * Step 1: Curl or Ping from the machine the broker is on to the SCM. If that succeeds, go to step 2, otherwise they need to figure out network or connectivity from broker machine to the SCM.
    * Note that if the Broker image you're using does not have the tools that you need, you can always create your own. Snyk uses the "slim" flavour but if you rebuild it with the standard one, there's more tools available.
  * Step 2: Next we need to determine if the docker container itself can communicate with the SCM.
    * Try to connect directly from broker to the SCM using the following commands
      * '**docker exec -it containerID bash**' gives you a shell inside the running container \(Remember to run "**docker container ls**" to get the id of the container\). The container must be running, otherwise won't work. Note that usually very few tools are available. Curl might or might not be there. Vi usually is, nano is not, etc...
        * For example, if you have a broker that can't connect to SCM, use this, then run a ping to that machine.
          * ping or curl their-scm-hostname
      * If this fails, you need to enable DNS resolution from the container. In the following example this particular user-edited /etc/resolv.conf file
        * Replace upper case items:

          ~$ sudo docker container run -d --dns IPOFNAMESERVER --name jovial snyk/broker:bitbucket-server

          YOURIMAGEID

          ~$ sudo docker exec -it jovial cat /etc/resolv.conf
      * A common user mistake is to update a container or create a new one, for whatever reason, and leave the old one running. The consequence is that it may appear that the updated settings may not be updated or live. Make sure to kill all instances and make sure you're running the one with the changes.

## GitHub Enterprise, HTTPS, SSL, and OpenSSL

```text
error while sending websocket request over HTTP connection
```

Check the `/systemcheck` ​ and `/healthcheck` endpoints and review the data output there. If there's an error regarding openssl/key size \(`ess_ske_dhe:dh key too small`\)--the problem likely lies in OpenSSL's key requirements based on an update that requires a higher strength key. 

To see if this is the problem, comment out the following line in our docker file:

```text
RUN apt-get update && apt-get install -y ca-certificates
```

**Note:** the tactic above is not meant to be used in production as removing this line prevents Snyk from providing both regular and security updates and should only be used to troubleshoot.

OpenSSL key creation and documentation can be found in their article on [Command Line Utilities](https://wiki.openssl.org/index.php/Command_Line_Utilities).

