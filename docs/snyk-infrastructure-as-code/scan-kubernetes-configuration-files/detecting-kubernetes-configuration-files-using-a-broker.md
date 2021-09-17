# Detecting Kubernetes configuration files using a broker

If you are using a privately hosted Git repository, Snyk Broker can connect it with Snyk products. See the [broker documentation](../../integrations/snyk-broker/broker-introduction.md/) for details.

This document describes the additional configuration required for the Kubernetes configuration and Helm files, for Snyk IaC.

## Writing the configuration

The Kubernetes scanning features need access to the YAML or JSON files from the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using.

1. Find and download the appropriate accept.json sample file for your source control system [from the Broker repository](https://github.com/snyk/broker/tree/master/client-templates).
2. Rename it to `accept.json` and add the below rules, appropriate to your SCM, to the **private** array in the JSON file.
3. Follow the [Configuring the broker](detecting-kubernetes-configuration-files-using-a-broker.md#configuring-the-broker/) instructions.

## GitHub & GitHub Enterprise rules

```text
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yaml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.yml",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.json",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*/*.tpl",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/repos/:name/:repo/contents/:path*%2F*.tpl",
  "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
},
```

## Bitbucket rules

```text
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.yaml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.yml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.yml",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.json",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.json",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*/*.tpl",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/projects/:project/repos/:repo/browse*%2F*.tpl",
  "origin": "https://${BITBUCKET_API}",
  "auth": {
    "scheme": "basic",
    "username": "${BITBUCKET_USERNAME}",
    "password": "${BITBUCKET_PASSWORD}"
  }
},
```

## GitLab rules

```text
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yaml",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.yml",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.json",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.json",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*/*.tpl",
  "origin": "https://${GITLAB}"
},
{
  "//": "used to determine Infrastructure as Code issues",
  "method": "GET",
  "path": "/api/v4/projects/:project/repository/files*%2F*.tpl",
  "origin": "https://${GITLAB}"
},
```

## Configuring the broker

The broker takes the path to the accept.json file \(with the rules above added\) in the ACCEPT environment variable. You can see an example of passing that to the GitHub broker below.

```text
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

Note that this gives Snyk the ability to query for any `.yaml`, `.yml` or `.json` files. If you would prefer to be stricter you can alter the paths in the examples above to be more restrictive to certain projects or file layouts.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

