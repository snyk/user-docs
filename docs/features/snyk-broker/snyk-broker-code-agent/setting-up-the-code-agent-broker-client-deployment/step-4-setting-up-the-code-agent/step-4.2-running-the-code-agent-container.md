# Step 4.2: Running the Code Agent container

Once the Code Agent image is stored on your machine, you need to use the `docker run` command in order to run the image, and launch a Code Agent container that is based on it.

**To run the Code Agent container:**

* In the terminal, enter the following command to launch a container based on the Snyk Code Agent image:

```
docker run --name <container_name> \
-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_ no.> \
-e PORT=<Code_Agent_container_port_no.> -e SNYK_TOKEN=<Snyk_API_token> --network <network_name> \
snyk/code-agent
```

Where:

* `--name <container_name>` - a new name for the Code Agent container. This name will be used to define the `GIT_CLIENT_URL` parameter for the Broker Client that you will run next. For example, `code-agent`.
* `-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_ no.>` - the mapping of a physical open port in the host machine to a port in the Code Agent container. These port numbers on the host machine and container do not have to be the same. For example, `3001:3000`.\
  **Note**: The port no. of the host machine must be unique.
* `-e PORT` - the port of the Code Agent container, where it  accepts external connections. The default is `3000`. This port no. must be the same as the `<Code_Agent_container_port_ no.>` in the `-p` parameter above.
* `-e SNYK_TOKEN` **** - your [Snyk API token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-snyk-api-token), as appears in your **Account Settings** page on the Snyk Web UI.&#x20;
* `--network` – the name of the [Docker bridge network](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-3-creating-a-network-for-the-broker-client-and-code-agent-communication) that was previously created. For example, `mySnykBrokerNetwork`.&#x20;
* `snyk/code-agent` - the Docker image of the Code Agent container.

Once the Code Agent setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"Application started", ... }`

<figure><img src="../../../../../.gitbook/assets/Code Agent - Exmaple - success.png" alt=""><figcaption></figcaption></figure>

**To verify the setup and details of the Code Agent container:**

* Run:

```
docker ps
```

The output should look similar to the following:

```
CONTAINER ID   IMAGE            COMMAND                 CREATED      STATUS      PORTS                    NAMES
eebd7d4f0568   snyk/code-agent "docker-entrypoint.s…"   9 days ago   Up 9 days   0.0.0.0:3000->3000/tcp   code-agent
```

### An example **for** running the Code Agent &#x20;

&#x20;For example, the following command was entered in a terminal, to launch a Code Agent container:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
snyk/code-agent
```

Where:

* `--name` - the name of the new Code Agent container is `code-agent`.
* `-p` - port `3000` on the host machine is mapped to port `3000` on the Code Agent container.
* `-e PORT` - the port of the Code Agent container, where is accepts external connections, is `3000`.
* `-e SNYK_TOKEN` - the Snyk API token is `fa7f….`
* `--network` - the name of the Docker bridge network, which will be used for the communication with the Client Broker, is `mySnykBrokerNetwork`.
* `snyk/code-agent` - the Docker image of the Code Agent container.

### **Setting up the Code Agent to work with a Proxy Server**

To use the Code Agent - Broker Client deployment in an infrastructure that uses a proxy,  add the following environment variables to the `docker run` command of the Code Agent:

```
-e HTTP_PROXY=http://my.proxy.address:<port_no.>
-e HTTPS_PROXY=http://my.proxy.address:<port_no.>
```

In addition, you will need to add these environment variables to the Broker Client component and a command to bypass the Code Agent, as described on [Setting up the Broker Client to work with a Proxy server](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display/setting-up-the-broker-client-to-work-with-a-proxy-server).

**Note**: For more information on using Docker containers with a proxy, see [Configure Docker to use a proxy server](https://docs.docker.com/network/proxy/).

&#x20;

To disable a certificate verification, for example, in the case of self-signed certificates, add the following parameter to the Code Agent `docker run` command:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```

&#x20;
