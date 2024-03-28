# Detecting Terraform configuration files using Snyk Broker (Custom)

## **Terraform in Snyk Broker**

By default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Terraform, you can add an environment variable, `ACCEPT_IAC`, with any combination of `tf,yaml,yml,json,tpl.`

Example:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl
       snyk/broker:github-com
```

Otherwise, you can edit your `accept.json`, add the relevant IaC specific rules, and load the customized accept file into the container. Note that if a custom accept file (from a separate folder) is used (using `ACCEPT` environment variable), the `ACCEPT_IAC` mechanism cannot be used.

These are the instructions if you require a custom allow-list and want to add Terraform files into the files Snyk can scan for.

## Writing the configuration

The Terraform scanning features need access to the `.tf` files from the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using.

1. Find the appropriate accept.json sample file for your source control system and download it [from the Broker repository](https://github.com/snyk/broker/tree/master/client-templates).
2. Rename the file to `accept.json` and to the **private** array in the JSON file, add the rules that follow as appropriate to your SCM.
3. Follow the instructions for [Configuring Broker](detecting-terraform-configuration-files-using-a-broker.md#configuring-broker).

### GitHub rules

```
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.tf",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.tf",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
```

### Bitbucket rules

```
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.tf",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.tf",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
```

### GitLab rules

```
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.tf",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Terraform issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.tf",
  "origin": "https://${GITLAB}"
},
```

### Azure Repos

Copy the following list of file extensions:

```
            "**/*.yaml",
            "**%2F*.yaml",
            "**/*.yml",
            "**%2F*.yml",
            "**/*.json",
            "**%2F*.json",
            "**/*.tf",
            "**%2F*.tf",
```

Add the extensions to the `values` array in two places in the [accept.json](https://github.com/snyk/broker/blob/master/client-templates/azure-repos/accept.json.sample):

* `"//": "get file content. restrict by file types"`
* `"//": "check file existence. restrict by file types"`

## Configuring Broker

Broker takes the path to the accept.json file, with the applicable rules added, in the ACCEPT environment variable. The following provides an example of passing that variable to the GitHub Broker.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e GITHUB_TOKEN=secret-github-token \
  -e PORT=8000 \
  -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
  -e ACCEPT=/private/accept.json
  -v /local/path/to/private:/private \
  snyk/broker:github-com
```

Note that this gives Snyk the ability to query for any `.tf` files. If you would prefer to be stricter, you can alter the paths in the preceding examples to be more restrictive for certain projects or file layouts.
