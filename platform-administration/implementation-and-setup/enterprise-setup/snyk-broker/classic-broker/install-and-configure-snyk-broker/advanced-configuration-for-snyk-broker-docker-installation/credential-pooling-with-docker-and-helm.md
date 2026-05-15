# Credential pooling with Docker and Helm

Under some circumstances it can be desirable to create a "pool" of credentials, for example, to work around rate-limiting issues. You can do this by creating an environment variable ending in `_POOL`, separating each credential with a comma. The Broker Client will then, when doing variable replacement, look to see if the variable in use has a variant with a `_POOL` suffix, and if so, use the next item in that pool. For example, if you have set the environment variable `GITHUB_TOKEN`, but want to provide multiple tokens, you would do this instead:

```
GITHUB_TOKEN_POOL=token1, token2, token3
```

You can add this as env var + value in the Helm Chart.

Then the Broker Client would, any time it needed `GITHUB_TOKEN`, instead take an item from the `GITHUB_TOKEN_POOL`.

Credentials are taken in a round-robin fashion, so the first, the second, the third, and so on, until the Broker Client reaches the end and then takes the first credential again.

Calling the `/systemcheck` endpoint will validate all credentials, in order, and will return an array where the first item is the first credential, and so on. For example, if you were running the GitHub Client and have this:

```
GITHUB_TOKEN_POOL=good_token, bad_token
```

The `/systemcheck` endpoint would return the following, where the first object is for `good_token` and the second for `bad_token`:

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

The credentials are masked. However, note that if your credentials contain six (6) or fewer characters, they will be completely replaced with the mask.

## **Limitations of credential pooling**

Credential validity is not checked before using a credential, nor are invalid credentials removed from the pool, so it is _strongly_ recommended that credentials be used exclusively by the Broker Client to avoid credentials reaching rate limits at different times, and that the `/systemcheck` endpoint be called before use.

Some providers, such as GitHub, do rate-limiting on a per-user basis, not a per-token or per-credential basis, and in those cases you will need to create multiple accounts with one credential per account.

## **Credentials matrix**

Generating a Matrix of credentials is not supported.

A "Matrix" in this case is defined as taking two (or more) `_POOL`s of length `x` and `y`, and producing one final pool of length `x * y`. For example, given an input like:

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
