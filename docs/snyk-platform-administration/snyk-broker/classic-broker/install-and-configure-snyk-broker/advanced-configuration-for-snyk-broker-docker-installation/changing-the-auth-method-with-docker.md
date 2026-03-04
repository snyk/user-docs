# Changing the auth method with Docker

Each integration has an auth method set by default, with the exact method varying by service.

BitBucket Server and Data Center, for example, uses Basic Auth with a username and password in the `accept.json` file:

```json
{
  "private": [
    {
      ...,
      "auth": {
         "scheme": "basic",
         "username": "${BITBUCKET_USERNAME}",
         "password": "${BITBUCKET_PASSWORD}"
      }
    },
    ...
  ]
}
```

For Artifactory, the auth method is configured in the `.env` file by default:

```shell
# The URL to your artifactory
# If not using basic auth this will only be "<yourdomain.artifactory.com>/artifactory"
ARTIFACTORY_URL=<username>:<password>@<yourdomain.artifactory.com>/artifactory
```

For GitHub, the auth meth is part of the `origin` field in the `accept.json` file:

```json
{
  "private": [
    {
      ...,
      "origin": "https://${GITHUB_TOKEN}@${GITHUB_API}"
    },
    ...
  ]
}
```

You can override the authentication method. Valid values for `scheme` are `bearer`, `token`, and `basic`, which set the Authorization header to `Bearer`, `Token`, and `Basic` respectively. If a bearer token is preferred, the `accept.json` file can be configured as such:

```json
{
  "private": [
    {
      ...,
      "auth": {
        "scheme": "bearer",
        "token": "${BEARER_TOKEN}"
      }
    },
    ...
  ]
}
```

Note that you must set this for every individual object in the `private` array.

If `scheme` is `bearer` or `token`, you must provide a `token`, and if `scheme` is `basic`, you must provide a `username` and `password`.

This overrides any other configured authentication method, for example, setting the token in the `origin` field, or in the `.env` file.

{% hint style="info" %}
Snyk Broker does not support authentication with mTLS method. &#x20;
{% endhint %}
