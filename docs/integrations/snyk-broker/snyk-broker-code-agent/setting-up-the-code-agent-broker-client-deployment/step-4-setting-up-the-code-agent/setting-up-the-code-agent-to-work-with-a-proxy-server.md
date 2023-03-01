# Setting up the Code Agent to work with a Proxy Server

{% hint style="info" %}
Environment variables (provided with -e) are case sensitive. Ensure that they are provided as defined on this page.
{% endhint %}

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy, add the following environment variables to the `docker run` command of the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.> \
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
```

If your proxy requires username and password authentication, add the following additional environment variable:

```
-e PROXY_AUTH=userID:userPass
```

In addition, add these environment variables to the Broker Client component and a command to bypass the Code Agent.

For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).

## **Custom certificates**

To use Code Agent with a proxy secured by a custom certificate (HTTPS), add the following environment variables to the `docker run` command of the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.> \
-e HTTPS_PROXY=https://my.proxy.address:<port_no.>
```

The following steps depend on the version of Code Agent you are running. If you are using the `latest` tag, to find your nearest versioned image:

* Compare the `digest` of your local image against [Docker Hub Code Agent Tags](https://hub.docker.com/r/snyk/code-agent/tags): `docker images snyk/code-agent --digest`
* Find the next image tag of the form `x.y.z` that was released _before_ your local image was built.

## **Version `1.18.0` and above**

To trust a custom Certificate Authority, you must have either:

* A single Certificate Authority (encoded as a `PEM`), or
* A directory containing multiple Certificate Authorities (encoded as `PEM`)

To trust a single certificate, add the following arguments to the `docker run` command of the Code Agent:

```
-v local/path/to/ca.pem:/etc/certs/ca.pem \
-e GIT_SSL_CAINFO='/etc/certs/ca.pem'
```

To trust a directory of certificates, add the following arguments to the `docker run` command of the Code Agent:

```
-v local/path/to/certdirectory:/etc/certs
-e GIT_SSL_CAPATH='/etc/certs'
```

## **Version `1.16.0` to `1.17.0`**

Follow the preceding steps and add the following argument to the `docker run` command of the Code Agent:

```
-e CODE_AGENT_GIT_CLI=true
```

## **Version `1.15.2` and below**

Code Agent `1.15.2` and below do not support trust of custom Certificate Authorities, and instead must run in a mode that trusts all certificates.

Add the following environment variable to the `docker run` command of the Code Agent:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```
