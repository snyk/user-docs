# Examples for running the Broker Client with the code snippets display

### **Running the Broker Client for an integrated GitHub Server**

The following command was entered to run the Broker Client for an integrated GitHub Server:

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

Where:

* `-p 8000:8000` - port number `8000` on the host machine is mapped to port number `8000` on the Broker Client container, for the communication between the Broker Client container and the Broker Server and Code Agent.
* `-e BROKER_TOKEN` - the Broker token that is associated with the specific Organization and GitHub.
* `-e GITHUB_TOKEN` – the GitHub token for accessing the GitHub repositories.
* `-e BROKER_CLIENT_URL` - the URL to the host machine of the Broker Client is `http://localhost:8000`.
* `-e PORT` - the local port, where the Broker Client container accepts connections, is `8000`.
* `-e GIT_CLIENT_URL=http://code-agent:3000` - the URL to the port of the running Code Agent container. The URL includes the name of the Code Agent container – `code-agent` - with its port no. - `3000`.
* `-e ACCEPT=/private/accept-github.json` - the name of the folder that stores the downloaded `accept.json` file is `private`, and the name of the file for GitHub is `accept.json`.
* `-v /private:/private` - the path to the folder that stores the `accept-github.json` file is private and the name of the folder is private.
* `--network` – the name of the Docker bridge network, which will be used for the communication with the Client Broker, is `mySnykBrokerNetwork`.
* `snyk/broker:github-com` - the Docker image of the Broker Client for GitHub.

&#x20;

### **Running the Broker Client for an integrated Azure Repos Server**

The following command was entered to run the Broker Client for an integrated Azure Repos Server:

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

Where:

* `-p 8001:8001` - port number `8001` on the host machine is mapped to port number `8001` on the Broker Client container, for the communication between the Broker Client container and the Broker Server and Code Agent.
* `-e BROKER_TOKEN` - the Broker token that is associated with the specific Organization and Azure Repos.
* `-e AZURE_REPOS_TOKEN` – the Azure Repo token for accessing the Azure Repos repositories.
* `-e AZURE_REPOS_ORG` - the name of the Azure Repos organization is `snyktest`.
* `-e AZURE_REPOS_HOST` - the domain name of the Azure Repos Server is `dev.azure`.com.
* `-e BROKER_CLIENT_URL` - the URL to the host machine of the Broker Client is `http://localhost:8001`.
* `-e PORT` - the local port, where the Broker Client container accepts connections, is `8001`.
* `-e GIT_CLIENT_URL=http://code-agent:3000` - the URL to the port of the running Code Agent container. The URL includes the name of the Code Agent container – `code-agent` - with its port no. - `3000`.
* `-e ACCEPT=/private/accept.json` - the name of the folder that stores the downloaded `accept.json` file is private, and the name of the file for GitHub is `accept.json`.
* `-v ./private:/private` - the path to the folder that stores the `accept.json` file is `private` and the name of the folder is `private`.
* `--network` – the name of the Docker bridge network, which will be used for the communication with the Client Broker, is `mySnykBrokerNetwork`.
* `snyk/broker:azure-repos` - the Docker image of the Broker Client for Azure Repos.

&#x20;
