# Step 5.2b – Running the Broker Client with the code snippets display

{% hint style="info" %}
**Important!** The setup commands for running the Broker Client that are described in this section include the common commands that should be used for all SCMs. However, some SCMs require additional parameters for the Broker Client setup. When such additional parameters are required, they are indicated in this section, but when setting up a Broker Client for a specific SCM, it is recommended to also use the section that is dedicated for that SCM. For more information, see [Snyk Broker - Integration Setups](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples).
{% endhint %}

The following instructions describe how to set up the Broker Client in a way that will display the code snippets of the Snyk Code results on the Web UI:

<figure><img src="../../../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

**Note:** To set up the Broker Client without the code snippets, see [Step 5.2a – Running the Broker Client without the code snippet display](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display).

When setting up the Broker Client for displaying code snippets, you need to use the same parameters as for the basic Broker Client setup, while adding the following file and parameters:&#x20;

1\.  A pre-defined `accept.json` file needs to be downloaded to your machine. This `accept.json` file is customized for each SCM, and it contains the required rules for displaying the code snippets. &#x20;

2\.  Two additional parameters should be added to the setup commands, in order to mount the `accept.json` file to the Broker Client:

* `-e ACCEPT=/<folder_name>/accept.json`
* `-v / local/path/to/<folder_name>:/<folder_name>`



### **Downloading the accept.json file**

To display the code snippets of the Snyk Code results on the Web UI, first you need to download a pre-defined accept.json file. When downloading the accept.json file, select the file that is customized for your integrated SCM, and save it on an accessible location.

**To download the accept.json file:**

1\.  Download the `accept.json` file for your integrated SCM:

**Note**: On each Integration Setup page, locate the pre-defined `accept.json` file for this integration, and download it.

* [GitHub](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/broker-example-how-to-setup-broker-with-jira)
* [GitHub Enterprise](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/setup-broker-with-github-enterprise)
* [Gitlab](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/setup-broker-with-gitlab)
* [Bitbucket Server/Data Center](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/data-center)
* [Azure Repos](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/setup-broker-with-azure-repos)

2\.  Verify that the downloaded file is called “**`accept.json`**”. If during the download process the file name has changed, rename it to its original name.

3\.  Save the file in a secure and separate folder, such as `./private/accept.json`.

&#x20;The next step is running the Broker Client with the `accept.json` file.

&#x20;&#x20;

### **Running the Broker Client with the accept.json file**

**To run the Broker Client container with the `accept.json` file:**

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
   -e ACCEPT=/<folder_name>/accept.json \
   -v /local/path/to/<folder_name>:/<folder_name> \
   --network mySnykBrokerNetwork \
   snyk/broker:<SCM_tag>
```

Where:

* `-- restart=always` - a Docker command that determines that the Broker Client container will always restart regardless of the exit status.
* `-p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.>` - the mapping of a physical open port in the host machine to a port in the Broker Client container. These port numbers on the host machine and container do not have to be the same. For example, `8001:8000`.\
  **Note**: The port no. of the host machine must be unique.
* `-e BROKER_TOKEN` - [the Broker token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-broker-token) that is associated with the specific Organization and the specific integrated SCM.&#x20;
* `-e <SCM_TOKEN>` - [the SCM token](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-1-obtaining-the-required-tokens-for-the-setup-procedure/obtaining-your-scm-token) for the specific integrated SCM.&#x20;
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
    **Note**: For [Bitbucket Server/Data Center](https://docs.snyk.io/features/snyk-broker/snyk-broker-set-up-examples/data-center) you also need to add the following parameter:  \
    `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(without http/s)> \`
  * **GitLab**: `-e GITLAB`
* \[Optional] `-e BROKER_CLIENT_URL=` - the URL to the host machine of the Broker Client. The URL can include an IP address or a DNS with the port no. of the host machine. For example, `http://localhost:8001`.  \
  Add this parameter only if you are using the same Broker Client for other Snyk products, and you want to enable for them the Automatic PR Checks feature. Since currently the Automatic PR Checks feature is not supported by the Code Agent, you do not have to use this parameter for the Code Agent.
* `-e PORT` – the port no. of the Broker Client container, where it accepts external connections. The default is `8000`. This port no. must be the same as the `<Broker_Client_container_port_ no.>` in the `-p` parameter above.
* `-e GIT_CLIENT_URL=` a URL to the port of the running Code Agent container. The URL should include the name of the Code Agent container with its port number. For example, `http://code-agent:3000`.
* `-e ACCEPT=` - the name of the folder that stores the downloaded `accept.json` file with the name of the file. For example, `/private/accept.json`.
* `-v /local/path/to/<folder_name>:/<folder_name>` - the path to the folder that stores the accept.json file with the name of the folder. For example, `/private:/private`.
* `--network` - the name of [the Docker bridge network](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-3-creating-a-network-for-the-broker-client-and-code-agent-communication), which will be used for the communication with the Code Agent.&#x20;
* `snyk/broker:<SCM_tag>` - [the Docker image of the Broker Client](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.1-downloading-or-updating-the-snyk-broker-client-docker-image) for the specific integrated SCM.

{% hint style="info" %}
To use the Broker Client with a proxy server, see [Setting up the Broker Client to work with a Proxy server](https://docs.snyk.io/features/snyk-broker/snyk-broker-code-agent/setting-up-the-code-agent-broker-client-deployment/step-5-setting-up-the-broker-client/step-5.2a-running-the-broker-client-without-the-code-snippet-display/setting-up-the-broker-client-to-work-with-a-proxy-server).
{% endhint %}

Once the Broker Client setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../../.gitbook/assets/Broker Client - Setup success message (1).png" alt=""><figcaption></figcaption></figure>

****

**To verify the setup and details of the Broker Client container:**

* Run:

```
docker ps
```

The output should look similar to the following:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
1cef6e3e3290   snyk/broker:github-com    "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   nifty_cori  
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```

&#x20;
