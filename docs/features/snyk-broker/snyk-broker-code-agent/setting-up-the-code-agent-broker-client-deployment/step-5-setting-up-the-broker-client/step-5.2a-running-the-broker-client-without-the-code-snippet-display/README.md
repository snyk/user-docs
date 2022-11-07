# Step 5.2a – Running the Broker Client without the code snippet display

{% hint style="info" %}
**Important!** The setup commands for running the Broker Client that are described in this section include the common commands that should be used for all SCMs. However, some SCMs require additional parameters for the Broker Client setup. When such additional parameters are required, they are indicated in this section, but when setting up a Broker Client for a specific SCM, it is recommended to also use the section that is dedicated for that SCM. For more information, see [Snyk Broker - Integration Setups](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples).
{% endhint %}

Once the Broker Client image is stored on your machine, you need to use the docker run command in order to run the image and launch a Broker Client container that is based on it.

The following instructions describe how to set up the Broker Client in a way that will NOT display the code snippets of the Snyk Code results on the Web UI:

<figure><img src="../../../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3).png" alt=""><figcaption></figcaption></figure>

**Note**: To display the code snippets, see [Running the Broker Client with the code snippets display](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2b-running-the-broker-client-with-the-code-snippets-display).

\_\_ **To run the Broker Client container:**

* In the terminal, enter the following command to launch a container based on the Snyk Broker Client image:

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

Where:

* `-- restart=always` - a Docker command that determines that the Broker Client container will always restart regardless of the exit status.
* `-p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.>` - the mapping of a physical open port in the host machine to a port in the Broker Client container. These port numbers on the host machine and container do not have to be the same. For example, `8001:8000`.\
  **Note**: The port no. of the host machine must be unique.
* `-e BROKER_TOKEN` - [the Broker token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token) that is associated with the specific Organization and the specific integrated SCM.
* `-e <SCM_TOKEN>` - [the SCM token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token) for the specific integrated SCM.
* `-e <SCM_domain>=` - your SCM domain name, without http/https. For example, `snyk.git.com`. For each SCM, use the following parameter:
  * **GitHub** - the `-e <SCM_domain>` parameter is NOT required.
  * **GitHub Enterprise**: `-e GITHUB`\
    **Note**: For [GitHub Enterprise](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/setup-broker-with-github-enterprise) you also need to add the following parameters:\
    `-e GITHUB_API=<your.ghe.domain.com/api/v3_(without_http/s)> \`\
    `-e GITHUB_GRAPHQL=<your.ghe.domain.com/api_(without_http/s)> \`
  * **Azure Repos**: `-e AZURE_REPOS_HOST`\
    **Note**: For [Azure Repos](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/setup-broker-with-azure-repos) you also need to add the following parameter:\
    `-e AZURE_REPOS_ORG=<azure_repo_org_name> \`
  * **Bitbucket Server/Data Center**: `-e BITBUCKET`\
    **Note**: For [Bitbucket Server/Data Center](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/data-center) you also need to add the following parameter:\
    `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(without http/s)> \`
  * **GitLab**: `-e GITLAB`
* \[Optional] `-e BROKER_CLIENT_URL=` - the URL to the host machine of the Broker Client. The URL can include an IP address or a DNS with the port no. of the host machine. For example, `http://localhost:8000`.\
  Add this parameter only if you are using the same Broker Client for other Snyk products, and you want to enable for them the Automatic PR Checks feature. Since currently the Automatic PR Checks feature is not supported by the Code Agent, you do not have to use this parameter for the Code Agent.
* `-e PORT` – the port no. of the Broker Client container, where it accepts external connections. The default is `8000`. This port no. must be the same as the `<Broker_Client_container_port_ no.>` in the `-p` parameter above.
* `-e GIT_CLIENT_URL=` a URL to the port of the running Code Agent container. The URL should include the name of the Code Agent container with its port number. For example, `http://code-agent:3000`.
* `--network` – the name of [the Docker bridge network](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-3-creating-a-network-for-the-broker-client-and-code-agent-communication), which will be used for the communication with the Code Agent.
* `snyk/broker:<SCM_tag>` - [the Docker image of the Broker Client](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image) for the specific integrated SCM.

Once the Broker Client setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../../.gitbook/assets/Broker Client - Setup success message (1).png" alt=""><figcaption></figcaption></figure>

**To verify the setup and details of the Broker Client container:**

* Run:

```
docker ps
```

The output should look similar to the following:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
7d9a410e7eaa   snyk/broker:azure-repos   "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   sweet_williams
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```
