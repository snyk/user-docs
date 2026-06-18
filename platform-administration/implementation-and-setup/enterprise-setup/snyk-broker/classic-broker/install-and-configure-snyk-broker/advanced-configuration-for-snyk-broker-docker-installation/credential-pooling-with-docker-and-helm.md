# Credential pooling with Docker and Helm

Under some circumstances, you might want to create a "pool" of credentials, for example, to work around rate-limiting issues. To do this, create an environment variable ending in `_POOL`, separating each credential with a comma. When doing variable replacement, the Broker Client looks to see whether the variable in use has a variant with a `_POOL` suffix, and if so, uses the next item in that pool. For example, if you set the environment variable `GITHUB_TOKEN` but want to provide multiple tokens, do this instead:

```
GITHUB_TOKEN_POOL=token1, token2, token3
```

You can add this as env var + value in the Helm Chart.

Then, any time the Broker Client needs `GITHUB_TOKEN`, it instead takes an item from the `GITHUB_TOKEN_POOL`.

Credentials are taken in a round-robin fashion, so the first, the second, the third, and so on, until the Broker Client reaches the end and then takes the first credential again.

Calling the `/systemcheck` endpoint validates all credentials, in order, and returns an array where the first item is the first credential, and so on. For example, if you run the GitHub Client and have this:

```
GITHUB_TOKEN_POOL=good_token, bad_token
```

The `/systemcheck` endpoint returns the following, where the first object is for `good_token` and the second for `bad_token`:

```
[
  {
    "brokerClientValidationUrl": "https://api.github.com/user",
    "brokerClientValidationMethod": "GET",
    "brokerClientValidationTimeoutMs": 5000,
    "brokerClientValidationUrlStatusCode": 200,
    "ok": true,
    "maskedCredentials": "goo***ken"
  },
  {
    "brokerClientValidationUrl": "https://api.github.com/user",
    "brokerClientValidationMethod": "GET",
    "brokerClientValidationTimeoutMs": 5000,
    "ok": false,
    "error": "401 - {\"message\":\"Bad credentials\",\"documentation_url\":\"https://docs.github.com/rest\"}",
    "maskedCredentials": "bad***ken"
  }
]
```

The credentials are masked. However, note that if your credentials contain six or fewer characters, the mask completely replaces them.

## **Limitations of credential pooling**

The Broker Client does not check credential validity before using a credential, nor does it remove invalid credentials from the pool. Snyk _strongly_ recommends that only the Broker Client use the credentials, to avoid credentials reaching rate limits at different times, and that you call the `/systemcheck` endpoint before use.

Some providers, such as GitHub, do rate-limiting on a per-user basis, not a per-token or per-credential basis, and in those cases you must create multiple accounts with one credential per account.

## **Credentials matrix**

Generating a Matrix of credentials is not supported.

A "Matrix" in this case means taking two or more `_POOL`s of length `x` and `y`, and producing one final pool of length `x * y`. For example, given an input like:

```
USERNAME_POOL=u1, u2, u3
PASSWORD_POOL=p1, p2, p3
CREDENTIALS_POOL=$USERNAME:$PASSWORD
```

Matrix support would generate this internally:

```
CREDENTIALS_POOL=u1:p1,u1:p2,u1:p3,u2:p1,u2:p2,u2:p3,u3:p1,u3:p2,u3:p3
```

Instead, the Broker Client would generate this internally, using only the first pool it finds:

```
CREDENTIALS_POOL=u1:$PASSWORD,u2:$PASSWORD,u3:$PASSWORD
```
