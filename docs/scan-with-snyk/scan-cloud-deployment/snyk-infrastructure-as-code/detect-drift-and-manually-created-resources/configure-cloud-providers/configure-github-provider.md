# Configure GitHub provider

## Authentication

To use `iac describe`, set up credentials to make authenticated requests to GitHub. Snyk retrieves configuration information from [environment variables](https://registry.terraform.io/providers/integrations/github/latest/docs#argument-reference).&#x20;

GitHub tokens can be created from [this GitHub page](https://github.com/settings/tokens/).&#x20;

```
$ GITHUB_TOKEN=14758f1afd44c09b7992073ccf00b43d \
  GITHUB_ORGANIZATION=my-org \
  snyk iac describe --to="github+tf"
```

## Least privileged policy​ <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The following GitHub token scopes are the minimum required for `iac describe` to  scan every GitHub-supported resource.

```
# Required to enumerate private repos
repo

# Required to list your organization teams
# and other organization related resources
read:org
```

## **Repository permissions**

To see errors for the repositories listing, you must set the permission `repo` for your token. Without this permission, all private repositories appear as deleted from remote.
