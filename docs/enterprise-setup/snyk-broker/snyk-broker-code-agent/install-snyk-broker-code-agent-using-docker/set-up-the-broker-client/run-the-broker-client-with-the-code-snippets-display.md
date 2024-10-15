# Run the Broker client with the code snippets display

After the Broker client image is stored on your machine, use the `docker run` command to run the image and launch the Broker client.

## Docker run command to set up the Broker client to display code snippets

The setup commands for running the Broker client described in this section include the common commands used for all SCMs. Some SCMs require additional parameters for the Broker client setup, and those parameters are indicated in this section, but when you are setting up a Broker client for a specific SCM, see also the instructions for that SCM in the section [Install and configure Snyk Broker](../../../install-and-configure-snyk-broker/).

The following explains how to set up the Broker client in a way that displays code snippets of the Snyk Code results in the Web UI:

<figure><img src="../../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="Broker Client run with display of code snippets"><figcaption><p>Broker Client run with display of code snippets</p></figcaption></figure>

When you are setting up the Broker client for displaying code snippets, use the same parameters as for the basic Broker client setup, and add the following file and parameters:

Download a pre-defined `accept.json` file to your machine. This `accept.json` file is customized for each SCM, and it contains the required rules for displaying the code snippets.

Add these parameters to the setup commands to mount the `accept.json` file to the Broker Client:

`-e ACCEPT=/<folder_name>/accept.json`

`-v / local/path/to/<folder_name>:/<folder_name>`

## **Download the accept.json file**

To display the code snippets of the Snyk Code results in the Web UI, first download a pre-defined `accept.json` file. When downloading the accept.json file, select the file that is customized for your integrated SCM and save it in an accessible location.

1. Download the `accept.json` file for your integrated SCM:\
   On each Integration Setup page, locate the pre-defined `accept.json` file for this integration, and download it.
   * [GitHub](../../../install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md)
   * [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)
   * [Gitlab](../../../install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md)
   * [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md)
   * [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md)
2. Verify that the downloaded file is called `accept.json`. If during the download process, the file name has changed, rename the file to its original name.
3. Save the file in a secure and separate folder such as `./private/accept.json`.

## **Run the Broker client with the accept.json file**

In the terminal, enter the following command to launch the Broker client:

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

where:

* `-- restart=always` is a Docker command that determines that the Broker client container will always restart regardless of the exit status.
* `-p <host_machine_port_no._mapped to>:<Broker_Client_container_port_ no.>` is the mapping of a physical open port in the host machine to a port in the Broker client container. These port numbers on the host machine and container do not have to be the same, for example, `8001:8000`. The port number of the host machine must be unique.
* `-e BROKER_TOKEN` is the Broker token that is associated with the specific Organization and the specific integrated SCM.
* `-e <SCM_TOKEN>` is the SCM token for the specific integrated SCM.
* `-e <SCM_domain>=` is your SCM domain name without `http/https`, for example, `snyk.git.com`. For each SCM,. Use the following parameter:
  * **GitHub** - the `-e <SCM_domain>` parameter is NOT required.
  * **GitHub Enterprise**: `-e GITHUB`\
    For [GitHub Enterprise](../../../install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md) add the following parameters also:\
    `-e GITHUB_API=<your.ghe.domain.com/api/v3_(without_http/s)> \`\
    `-e GITHUB_GRAPHQL=<your.ghe.domain.com/api_(without_http/s)> \`
  * **Azure Repos**: `-e AZURE_REPOS_HOST`\
    For [Azure Repos](../../../install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md) add the following parameter also:\
    `-e AZURE_REPOS_ORG=<azure_repo_org_name> \`
  * **Bitbucket Server/Data Center**: `-e BITBUCKET`\
    For [Bitbucket Server/Data Center](../../../install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md) add the following parameter also:\
    `-e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0_(without http/s)> \`
  * **GitLab**: `-e GITLAB`
* \[Optional] `-e BROKER_CLIENT_URL=` is the URL to the host machine of the Broker client. The URL can include an IP address or a DNS with the port number of the host machine, for example, `http://localhost:8001`.\
  Add this parameter only if you are using the same Broker client for other Snyk products, and you want to enable for them the Automatic PR Checks feature. Since the Automatic PR Checks feature is not supported for the Code Agent, you do not have to use this parameter for the Code Agent.
* `-e PORT` is the port number of the Broker Client container, where it accepts external connections. The default is `8000`. This port number must be the same as the `<Broker_Client_container_port_ no.>` in the `-p` parameter above.
* `-e GIT_CLIENT_URL=` is a URL to the port of the running Code Agent container. The URL includes the name of the Code Agent container with its port number, for example, `http://code-agent:3000`.
* `-e ACCEPT=` is the name of the folder that stores the downloaded `accept.json` file with the name of the file, for example, `/private/accept.json`.
* `-v /local/path/to/<folder_name>:/<folder_name>` is the path to the folder that stores the accept.json file with the name of the folder, for example, `/private:/private`.
* `--network` is the name of the [Docker bridge network](../create-network-for-broker-client-and-code-agent-communication.md), used for communication with the Code Agent.
* `snyk/broker:<SCM_tag>` is the [Docker image of the Broker Client](download-or-update-the-snyk-broker-client-docker-image.md) for the specific integrated SCM.

Once the Broker client setup is completed successfully, the following message appears in the terminal:

`{ ..., "msg":"successfully established a websocket connection to the broker server", ... }`

<figure><img src="../../../../../.gitbook/assets/Broker Client - Setup success message (1).png" alt="Confirmation message for Broker Client setup"><figcaption><p>Confirmation message for Broker Client setup</p></figcaption></figure>

**To verify the setup and details of the Broker client container**, run:

```
docker ps
```

The output is similar to the following:

```
CONTAINER ID   IMAGE                     COMMAND                  CREATED             STATUS             PORTS                    NAMES
1cef6e3e3290   snyk/broker:github-com    "broker --verbose"       About an hour ago   Up About an hour   0.0.0.0:8000->8000/tcp   nifty_cori  
6216e27b8d28   snyk/code-agent           "docker-entrypoint.s…"   2 hours ago         Up 2 hours         0.0.0.0:3000->3000/tcp   code-agent
```

## Examples for running the Broker client with the code snippets display

### **Run the Broker client for an integrated GitHub server**

The following command was entered to run the Broker client for an integrated GitHub server:

```
docker run --restart=always \
   -p 8000:8000 \
   -e BROKER_TOKEN=b1dcxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxx \
   -e BROKER_CLIENT_URL=http://localhost:8000 \
   -e PORT=8000 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   -e ACCEPT=/private/accept.json \
   -v /private:/private \
   --network mySnykBrokerNetwork \
snyk/broker:github-com
```

where:

* `-p 8000:8000`is port number `8000` on the host machine, mapped to port number `8000` on the Broker Client container. This is used for the communication between the Broker Client container and the Broker Server and Code Agent.
* `-e BROKER_TOKEN` is the Broker token associated with the specific Organization and GitHub.
* `-e GITHUB_TOKEN` is the GitHub token for accessing the GitHub repositories.
* `-e BROKER_CLIENT_URL` is the URL to the host machine of the Broker Client, `http://localhost:8000`.
* `-e PORT` is the local port, where the Broker client container accepts connections, `8000`.
* `-e GIT_CLIENT_URL=http://code-agent:3000` is the URL to the port of the running Code Agent container. The URL includes the name of the Code Agent container – `code-agent` - with its port no. - `3000`.
* `-e ACCEPT=/private/accept-github.json` is the name of the folder that stores the downloaded `accept.json` file, `private`, and the name of the file for GitHub, `accept.json`.
* `-v /private:/private` is the path to the folder that stores the `accept-github.json` file, private, and the name of the folder, private.
* `--network` is the name of the Docker bridge network, used for communication with the Broker, client `mySnykBrokerNetwork`.
* `snyk/broker:github-com` is the Docker image of the Broker client for GitHub.

### **Run the Broker client for an integrated Azure Repos server**

The following command was entered to run the Broker client for an integrated Azure Repos server:

```
docker run --restart=always \
   -p 8001:8001 \
   -e BROKER_TOKEN=fe5bxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
   -e AZURE_REPOS_TOKEN=brmyxxxxxxxxxxxxxxxx \
   -e AZURE_REPOS_ORG=snyktest \
   -e AZURE_REPOS_HOST=dev.azure.com \
   -e BROKER_CLIENT_URL=http://localhost:8001 \
   -e PORT=8001 \
   -e GIT_CLIENT_URL=http://code-agent:3000 \
   -e ACCEPT=/private/accept.json \
   -v ./private:/private \
   --network mySnykBrokerNetwork \
snyk/broker:azure-repos
```

where:

* `-p 8001:8001` is port number `8001` on the host machine, mapped to port number `8001` on the Broker Client container, for communication between the Broker Client container and the Broker Server and Code Agent.
* `-e BROKER_TOKEN` is the Broker token associated with the specific Organization and Azure Repos.
* `-e AZURE_REPOS_TOKEN` is the Azure Repo token for accessing the Azure Repos repositories.
* `-e AZURE_REPOS_ORG` is the name of the Azure Repos organization, `snyktest`.
* `-e AZURE_REPOS_HOST` is the domain name of the Azure Repos Server, `dev.azure`.com.
* `-e BROKER_CLIENT_URL` is the URL of the host machine of the Broker client, `http://localhost:8001`.
* `-e PORT` is the local port, where the Broker Client container accepts connections, `8001`.
* `-e GIT_CLIENT_URL=http://code-agent:3000` is the URL to the port of the running Code Agent container. The URL includes the name of the Code Agent container, `code-agent`, with its port number, `3000`.
* `-e ACCEPT=/private/accept.json` is the name of the folder that stores the downloaded `accept.json` file, private, and the name of the file for GitHub, `accept.json`.
* `-v ./private:/private` is the path to the folder that stores the `accept.json` file, `private`, and the name of the folder. `private`.
* `--network` is the name of the Docker bridge network, used for communication with the Broker client, `mySnykBrokerNetwork`.
* `snyk/broker:azure-repos` is the Docker image of the Broker client for Azure Repos.
