# Detecting Kubernetes configuration files using Snyk Broker (Custom)

## **Kubernetes configuration files in Snyk Broker**

By default, some file types used by Infrastructure-as-Code (IaC) are not enabled. To grant the Broker access to IaC files in your repository, for example, Kubernetes configuration files, you can add an environment variable `ACCEPT_IAC` with any combination of `tf,yaml,yml,json,tpl.`

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

Otherwise you can edit your `accept.json`, add the relevant IaC specific rules, and load the customized accept file into the container. Note that if a custom accept file (from a separate folder) is used (using `ACCEPT` environment variable), the `ACCEPT_IAC` mechanism cannot be used.

These are the instructions if you require a custom allow-list and want to add Kubernetes configuration files into the files Snyk can scan for.

## Writing the configuration

You must grant the Broker access to particular files in the repository. This requires specific API permissions. These API permissions are slightly different depending on which source control system you are using. The configuration that follows is for the file extensions `.yaml`, `.yml`, and `.json`. This allows the Broker to access potential Kubernetes and CloudFormation files, but you may adapt the configuration as necessary. For example, you may wish to add configurations for `.tf` files in order to scan Terraform HCL files.

1. Find the appropriate `accept.json` sample file for your source control system and download it [from the Broker repository](https://github.com/snyk/broker/tree/master/client-templates).
2. Rename the file to `accept.json` and to the **private** array in the JSON file, add the following rules as appropriate to your SCM.
3. Follow the instructions for [Configuring Broker](detecting-kubernetes-configuration-files-using-a-broker.md#configuring-broker).

### GitHub and GitHub Enterprise rules

```
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

### Bitbucket rules

```
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

### GitLab rules

```
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

### Azure Repo rules

```
{
  "public": [
    {
      "//": "used for pushing up webhooks from Azure",
      "method": "POST",
      "path": "/webhook/azure-repos/:webhookId"
    }
  ],
  "private": [
    {
      "//": "get list of projects for given organization",
      "method": "GET",
      "path": "/_apis/projects",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get specific repository for given organization",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get list of repositories for given organization",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get list of refs",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/refs",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "search through repositories of given organization",
      "method": "GET",
      "path": "_apis/git/repositories",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "create hook",
      "method": "POST",
      "path": "/_apis/hooks/subscriptions",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "delete hook",
      "method": "DELETE",
      "path": "/_apis/hooks/subscriptions/:subscriptionId",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get file content. restrict by file types",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/items",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "valid": [
        {
          "queryParam": "path",
          "values": [
            "**/package.json",
            "**%2Fpackage.json",
            "**/yarn.lock",
            "**%2Fyarn.lock",
            "**/package-lock.json",
            "**%2Fpackage-lock.json",
            "**/Gemfile",
            "**%2FGemfile",
            "**/Gemfile.lock",
            "**%2FGemfile.lock",
            "**/pom.xml",
            "**%2Fpom.xml",
            "**/*req*.txt",
            "**%2F*req*.txt",
            "**/requirements/*.txt",
            "**%2Frequirements%2F*.txt",
            "**/build.gradle",
            "**%2Fbuild.gradle",
            "**/gradle.lockfile",
            "**%2Fgradle.lockfile",
            "**/build.sbt",
            "**%2Fbuild.sbt",
            "**/.snyk",
            "**%2F.snyk",
            "**/packages.config",
            "**%2Fpackages.config",
            "**/*.csproj",
            "**%2F*.csproj",
            "**/*.vbproj",
            "**%2F*.vbproj",
            "**/*.fsproj",
            "**%2F*.fsproj",
            "**/project.json",
            "**%2Fproject.json",
            "**/Gopkg.toml",
            "**%2FGopkg.toml",
            "**/Gopkg.lock",
            "**%2FGopkg.lock",
            "**/vendor.json",
            "**%2Fvendor.json",
            "**/composer.lock",
            "**%2Fcomposer.lock",
            "**/composer.json",
            "**%2Fcomposer.json",
            "**/project.assets.json",
            "**%2Fproject.assets.json",
            "**/Podfile",
            "**%2FPodfile",
            "**/Podfile.lock",
            "**%2FPodfile.lock",
            "**/go.mod",
            "**%2Fgo.mod",
            "**/go.sum",
            "**%2Fgo.sum",
            "**/Dockerfile",
            "**%2FDockerfile"
          ]
        },
        {
          "queryParam": "recursionLevel",
          "values": ["none"]
        },
        {
          "queryParam": "download",
          "values": ["true"]
        },
        {
          "queryParam": "includeContent",
          "values": ["true"]
        }
      ],
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get list of files for given repository",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/items",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "valid": [
        {
          "queryParam": "recursionLevel",
          "values": ["full"]
        },
        {
          "queryParam": "download",
          "values": ["false"]
        },
        {
          "queryParam": "includeContent",
          "values": ["false"]
        }
      ],
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "get list of commits for given repository",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/commits",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "update status of given commit",
      "method": "POST",
      "path": "/:owner/_apis/git/repositories/:repo/commits/:commitId/statuses",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "update status of given pull request",
      "method": "POST",
      "path": "/:owner/_apis/git/repositories/:repo/pullRequests/:pullRef/statuses",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "find PR for given repository",
      "method": "GET",
      "path": "/:owner/_apis/git/repositories/:repo/pullrequests",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "create new PR in given repository",
      "method": "POST",
      "path": "/:owner/_apis/git/repositories/:repo/pullrequests",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "update existing PR in given repository",
      "method": "PATCH",
      "path": "/:owner/_apis/git/repositories/:repo/pullrequests/:pullRef",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "push new commit in given repository",
      "method": "POST",
      "path": "/:owner/_apis/git/repositories/:repo/pushes",
      "origin": "https://${AZURE_REPOS_HOST}/${AZURE_REPOS_ORG}",
      "auth": {
        "scheme": "basic",
        "token": "${BROKER_CLIENT_VALIDATION_BASIC_AUTH}"
      }
    },
    {
      "//": "used to redirect requests to snyk git client",
      "method": "any",
      "path": "/snykgit/*",
      "origin": "${GIT_CLIENT_URL}"
    }
  ]
}
```

## Configuring Broker

Broker takes the path to the accept.json file, with the applicable rules added, in the ACCEPT environment variable. The following provides an example of passing that variable to the GitHub Broker.

```
docker run --restart=always \
  -p 8000:8000 \
  -e BROKER_TOKEN=secret-broker-token \
  -e GITHUB_TOKEN=secret-github-token \
  -e PORT=8000 \
  -e BROKER_CLIENT_URL=https://my.broker.client:8000 \
  -e ACCEPT=/private/accept.json \
  -v /local/path/to/private:/private \
  snyk/broker:github-com
```

Note that this gives Snyk the ability to query for any `.yaml`, `.yml` or `.json` files. If you would prefer to be stricter, you can alter the paths in the preceding examples to be more restrictive for certain projects or file layouts.
