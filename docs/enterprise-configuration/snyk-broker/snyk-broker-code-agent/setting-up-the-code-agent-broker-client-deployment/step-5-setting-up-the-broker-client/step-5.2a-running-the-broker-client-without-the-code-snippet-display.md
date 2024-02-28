# Step 5.2a – Running the Broker Client without the code snippet display

{% hint style="info" %}
The setup commands for running the Broker Client described in this section include the common commands used for all SCMs. However, some SCMs require additional parameters for the Broker Client setup.

When such additional parameters are required, they are indicated in this section, but when setting up a Broker Client for a specific SCM, use also the section that is specific to that SCM. For more information see [Install and configure Snyk Broker](../../../install-and-configure-snyk-broker/).
{% endhint %}

After the Broker Client image is stored on your machine, use the docker run command in order to run the image and launch a Broker Client container that is based on it.

The following explains how to set up the Broker Client in a way that does NOT display the code snippets of the Snyk Code results on the Web UI:

<figure><img src="../../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4) (1).png" alt="Broker Client run with no display of code snippets"><figcaption><p>Broker Client run with no display of code snippets</p></figcaption></figure>

To display the code snippets, see [Running the Broker Client with the code snippets display](step-5.2b-running-the-broker-client-with-the-code-snippets-display.md).

**To run the Broker Client container**, in the terminal enter the following command to launch a container based on the Snyk Broker Client image:

```
docker run --restart=always \
   -p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.> \
   -e BROKER_TOKEN=<Broker_token> \
   -e <SCM>_TOKEN=<SCM_token> \
   -e <SCM_domain>=<my.SCM.domain.com_(without_http/s)> \  
   -e BROKER_CLIENT_URL=<http://my.broker.client:<host_machine_port_no.> \
   -e PORT=<Broker_Client_container_port_no.> \
   -e GIT_CLIENT_URL=http://<Code_Agent_container_name:Code_Agent_port_no.> \
   --network mySnykBrokerNetwork \
   snyk/broker:<SCM_tag>
```

where:

* `-- restart=always` is a Docker command that determines that the Broker Client container will always restart regardless of the exit status.
* `-p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.>` is the mapping of a physical open port in the host machine to a port in the Broker Client container. These port numbers on the host machine and container do not have to be the same, for example, `8001:8000`.\
  The port number of the host machine must be unique.
* `-e BROKER_TOKEN` is the [Broker token](../step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token.md) that is associated with the specific Organization and the specific integrated SCM.
* `-e <SCM_TOKEN>` is the [SCM token](../step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token.md) for the specific integrated SCM.
* `-e <SCM_domain>=` is your SCM domain name, without http/https, for example, `snyk.git.com`. For each SCM. Use the parameter for your SCM:
  * **GitHub** - the `-e <SCM_domain>` parameter is NOT required.
  * **GitHub Enterprise**: `-e GITHUB`\
    For [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-github-enterprise.md) add the following parameters also:\
    `-e GITHUB_API=<your.ghe.domain.com/api/v3_(without_http/s)> \`\
    `-e GITHUB_GRAPHQL=<your.ghe.domain.com/api_(without_http/s)> \`
  * **Azure Repos**: `-e AZURE_REPOS_HOST`\
    For [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md) add the following parameter also:\
    `-e AZURE_REPOS_ORG=<azure_repo_org_name> \`
  * **Bitbucket Server/Data Center**: `-e BITBUCKET`\
    For [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md) add the following parameter also:\
    `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(without http/s)> \`
  * **GitLab**: `-e GITLAB`
* \[Optional] `-e BROKER_CLIENT_URL=` is the URL to the host machine of the Broker Client. The URL can include an IP address or a DNS with the port number of the host machine, for example, `http://localhost:8000`.\
  Add this parameter only if you are using the same Broker Client for other Snyk products, and you want to enable for them the Automatic PR Checks feature. Since currently the Automatic PR Checks feature is not supported by the Code Agent, you do not have to use this parameter for the Code Agent.
* `-e PORT` is the port number of the Broker Client container, where it accepts external connections. The default is `8000`. This port number must be the same as the `<Broker_Client_container_port_ no.>` in the `-p` parameter above.
* `-e GIT_CLIENT_URLis`a URL to the port of the running Code Agent container. The URL should include the name of the Code Agent container with its port number, for example, `http://code-agent:3000`.
* `--network` is the name of the [Docker bridge network](../step-3-creating-a-network-for-the-broker-client-and-code-agent-communication.md), which will be used for the communication with the Code Agent.
* `snyk/broker:<SCM_tag>` is the [Docker image of the Broker Client](step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image.md) for the specific integrated SCM.

When the Broker Client setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../.gitbook/assets/Broker Client - Setup success message.png" alt="Confirmation message for Broker Client setup"><figcaption><p>Confirmation message for Broker Client setup</p></figcaption></figure>

**To verify the setup and details of the Broker Client container**, run:

```
docker ps
```

The output is similar to the following:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
7d9a410e7eaa   snyk/broker:azure-repos   "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   sweet_williams
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```
