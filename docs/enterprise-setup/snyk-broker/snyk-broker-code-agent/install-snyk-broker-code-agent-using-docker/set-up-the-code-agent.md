# Set up the Code Agent

To set up the Code Agent using Docker:

* Download or update the Code Agent Docker image.
* Run the Code Agent container.
* If you are using a proxy server, set up the Code Agent to work with a proxy server.

## Download or update the Code Agent Docker image

### **Download the Code Agent Docker image**

Pull the Code Agent Docker image from Docker Hub. It is highly recommended to pull and use the latest Docker image version. Download the Snyk Code Agent Docker image to each machine that will run the Code Agent. Docker images are usually cached on the host machine.

To pull the Code Agent Docker image, in the terminal, enter:

```
docker pull snyk/code-agent
```

The download process starts for the Code Agent Docker image, for example:

<figure><img src="../../../../.gitbook/assets/Code Agent - Pull docker image - New.png" alt="Download process for Code Agent Docker image"><figcaption><p>Download process for Code Agent Docker image</p></figcaption></figure>

### Update the Code Agent Docker image

Pull the Code Agent Docker image again. If you are using the `latest` tag, the image is automatically updated. Otherwise, provide a new image tag:

```
docker pull snyk/code-agent:<image_tag>
```

Remove or stop the older Code Agent container.

Follow the steps in the next section start the new Code Agent container.

## Run the Code Agent container

Once the Code Agent image is stored on your machine, in the terminal, enter the following command to launch a container based on the Snyk Broker Code Agent image:

```
docker run --name <container_name> \
-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_ no.> \
-e PORT=<Code_Agent_container_port_no.> -e SNYK_TOKEN=<Snyk_API_token> --network <network_name> \
snyk/code-agent:<image_tag>
```

where:

* `--name <container_name>` is a new name for the Code Agent container. This name is used to define the `GIT_CLIENT_URL` parameter for the Broker Client that you run next, for example, `code-agent`.
* `-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_no.>` is the mapping of a physical open port in the host machine to a port in the Code Agent container. These port numbers on the host machine and container do not have to be the same.\
  Example: `3001:3000`.\
  The port number of the host machine must be unique.
* `-e PORT` is the port of the Code Agent container, where it accepts external connections. The default is `3000`. This port number must be the same as the `<Code_Agent_container_port_ no.>` in the `-p` preceding parameter.
* `-e SNYK_TOKEN` is your Snyk API token as it appears in your **Account Settings** page on the Snyk Web UI.
* `--network` is the name of the Docker bridge network that was previously created, for example, `mySnykBrokerNetwork`.
* `snyk/code-agent:<image_tag>` is the Docker image of the Code Agent container. Specify a tag if not using `latest`.

When the Code Agent setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"Application started", ... }`

<figure><img src="../../../../.gitbook/assets/Code Agent - Exmaple - success.png" alt="Success message when Code Agent setup is completed"><figcaption><p>Success message when Code Agent setup is completed</p></figcaption></figure>

### Verify the setup and details of the Code Agent container

Run the following:

```
docker ps
```

The output is similar to the following:

```
CONTAINER ID   IMAGE            COMMAND                 CREATED      STATUS      PORTS                    NAMES
eebd7d4f0568   snyk/code-agent "docker-entrypoint.s…"   9 days ago   Up 9 days   0.0.0.0:3000->3000/tcp   code-agent
```

### Example **for** running the Code Agent

In this example, the following command was entered in a terminal to launch a Code Agent container:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
snyk/code-agent
```

where:

* `--name` is the name of the new Code Agent container, `code-agent`.
* `-p` - port `3000` on the host machine is mapped to port `3000` on the Code Agent container.
* `-e PORT` is the port of the Code Agent container, where it accepts external connections, `3000`.
* `-e SNYK_TOKEN` is the Snyk API token, `fa7f….`
* `--network` is the name of the Docker bridge network, used for the communication with the Client Broker, `mySnykBrokerNetwork`.
* `snyk/code-agent` is the Docker image of the Code Agent container.

### **Connect to a Git instance with an internal certificate**

By default, the Code Agent establishes HTTPS connections to the Git instance. If your Git instatnce is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Code Agent.

For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the Docker container by mounting the folder and using the `CA_CERT` environment variable:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
-e CA_CERT=/private/ca.cert.pem \
-v /local/path/to/private:/private \
snyk/code-agent
```

## Set up the Code Agent to work with a proxy server

To use the Code Agent  in an infrastructure that uses a proxy, add the following environment variables to the `docker run` command of the Code Agent:

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

### **Custom certificates**

To use Code Agent with a proxy secured by a custom certificate (HTTPS), add the following environment variables to the `docker run` command of the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.> \
-e HTTPS_PROXY=https://my.proxy.address:<port_no.>
```

The following steps depend on the version of Code Agent you are running. If you are using the `latest` tag, to find your nearest versioned image:

* Compare the `digest` of your local image against [Docker Hub Code Agent Tags](https://hub.docker.com/r/snyk/code-agent/tags): `docker images snyk/code-agent --digest`
* Find the next image tag of the form `x.y.z` that was released _before_ your local image was built.

### **Version `1.18.0` and later**

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

### **Version `1.16.0` to `1.17.0`**

Follow the preceding steps and add the following argument to the `docker run` command of the Code Agent:

```
-e CODE_AGENT_GIT_CLI=true
```

### **Version `1.15.2` and earlier**

Code Agent `1.15.2` and below do not support trust of custom Certificate Authorities, and instead must run in a mode that trusts all certificates.

Add the following environment variable to the `docker run` command of the Code Agent:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```
