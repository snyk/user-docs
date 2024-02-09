# Step 4.2: Running the Code Agent container

Once the Code Agent image is stored on your machine, use the `docker run` command to run the image and launch a Code Agent container that is based on it.

{% hint style="info" %}
Environment variables (provided with -e) are case sensitive. Ensure that they are provided as defined on this page.
{% endhint %}

## Running the Code Agent container

In the terminal, enter the following command to launch a container based on the Snyk Code Agent image:

```
docker run --name <container_name> \
-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_ no.> \
-e PORT=<Code_Agent_container_port_no.> -e SNYK_TOKEN=<Snyk_API_token> --network <network_name> \
snyk/code-agent:<image_tag>
```

where:

* `--name <container_name>` is a new name for the Code Agent container. This name is used to define the `GIT_CLIENT_URL` parameter for the Broker Client that you run next. Example, `code-agent`.
* `-p <host_machine_port_no._mapped to>:<Code_Agent_container_port_no.>` is the mapping of a physical open port in the host machine to a port in the Code Agent container. These port numbers on the host machine and container do not have to be the same. Example: `3001:3000`.\
  The port number of the host machine must be unique.
* `-e PORT` is the port of the Code Agent container, where it accepts external connections. The default is `3000`. This port number must be the same as the `<Code_Agent_container_port_ no.>` in the `-p` parameter above.
* `-e SNYK_TOKEN` is your [Snyk API token](../../../../../getting-started/how-to-obtain-and-authenticate-with-your-snyk-api-token.md) as appears in your **Account Settings** page on the Snyk Web UI.
* `--network` is the name of the [Docker bridge network](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-3-creating-a-network-for-the-broker-client-and-code-agent-communication) that was previously created, for example, `mySnykBrokerNetwork`.
* `snyk/code-agent:<image_tag>` is the Docker image of the Code Agent container. Specify a tag if not using `latest`.

When the Code Agent setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"Application started", ... }`

<figure><img src="../../../../../.gitbook/assets/Code Agent - Exmaple - success.png" alt="Success message when Code Agent setup is completed"><figcaption><p>Success message when Code Agent setup is completed</p></figcaption></figure>

## Verifying the setup and details of the Code Agent container

Run the following:

```
docker ps
```

The output is similar to the following:

```
CONTAINER ID   IMAGE            COMMAND                 CREATED      STATUS      PORTS                    NAMES
eebd7d4f0568   snyk/code-agent "docker-entrypoint.s…"   9 days ago   Up 9 days   0.0.0.0:3000->3000/tcp   code-agent
```

## Example **for** running the Code Agent

In this example the following command was entered in a terminal to launch a Code Agent container:

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

## **Connecting to a Git with an internal certificate**

By default, the Code Agent establishes HTTPS connections to the Git. If your Git is serving an internal certificate (signed by your own CA), you can provide the CA certificate to the Code Agent.

For example, if your CA certificate is at `./private/ca.cert.pem`, provide it to the Docker container by mounting the folder and using the `CA_CERT` environment variable:

```
docker run --name code-agent \
-p 3000:3000 \
-e PORT=3000 -e SNYK_TOKEN=fa7fxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx --network mySnykBrokerNetwork \
-e CA_CERT=/private/ca.cert.pem \
-v /local/path/to/private:/private \
snyk/code-agent
```
