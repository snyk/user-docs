# Github

## Authentication

To use `iac describe`, set up credentials to make authenticated requests to GitHub. Just as for the terraform provider, Snyk retrieves configuration information from [environment variables](https://registry.terraform.io/providers/integrations/github/latest/docs#argument-reference).

```
$ GITHUB_TOKEN=14758f1afd44c09b7992073ccf00b43d\
  GITHUB_ORGANIZATION=my-org\
  snyk iac describe --to=github+tf
```

## Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/github/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

The following code shows the minimal scope required for the `iac describe` command to be able to scan every GitHub supported resource.

```
# Required to enumerate private repos
repo

# Required to list your organization teams
# and other organization related resources
read:org
```

## **Repository permissions**

Note that if you don't set the `repo` permission for your token, you will see no errors for the repositories listing. **Thus, all private repositories will appear as deleted from remote.**
