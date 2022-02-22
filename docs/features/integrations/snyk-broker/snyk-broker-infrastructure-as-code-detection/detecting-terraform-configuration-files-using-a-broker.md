# Detecting Terraform configuration files using a broker

If you are using a privately hosted Git repository then you can use the Snyk Broker to connect Snyk to it. See the [full broker documentation for setup](../set-up-snyk-broker/). The following details additional configuration required for the Terraform files.

## Writing the configuration

The Terraform scanning features need access to the `.tf` files from the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using.

1. Find and download the appropriate accept.json sample file for your source control system [from the Broker repository](https://github.com/snyk/broker/tree/master/client-templates).
2. Rename it to `accept.json` and add the below rules, appropriate to your SCM, to the **private** array in the JSON file.
3. Follow the [Configuring the broker](detecting-terraform-configuration-files-using-a-broker.md#configuring-the-broker) instructions.

## GitHub rules

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

## Bitbucket rules

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

## GitLab rules

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

## Configuring the broker

The broker takes the path to the accept.json file (with the rules above added) in the ACCEPT environment variable. You can see an example of passing that to the GitHub broker below.

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

Note that this gives Snyk the ability to query for any `.tf` files. If you would prefer to be stricter you can alter the paths in the examples above to be more restrictive to certain projects or file layouts.
