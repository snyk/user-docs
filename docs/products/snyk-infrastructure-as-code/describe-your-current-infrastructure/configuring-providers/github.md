# Github

## Authentication

To use describe, we need credentials to make authenticated requests to GitHub. Just like the terraform provider, we retrieve config from [environment variables](https://registry.terraform.io/providers/integrations/github/latest/docs#argument-reference).

```
$ GITHUB_TOKEN=14758f1afd44c09b7992073ccf00b43d\
  GITHUB_ORGANIZATION=my-org\
  snyk iac describe --to=github+tf
```

### Least privileged policy[​](https://docs.driftctl.com/0.22.0/providers/github/authentication#least-privileged-policy) <a href="#least-privileged-policy" id="least-privileged-policy"></a>

Below you can find the minimal scope required for describe to be able to scan every GitHub supported resources.

```
# Required to enumerate private repos
repo

# Required to list your organization teams
# and other organization related resources
read:org
```

**REPOSITORY PERMISSIONS**

Beware that if you don't set permission `repo` for your token, you won't see any errors for repositories listing. **Thus, all private repositories will appear as deleted from remote.**
