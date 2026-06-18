# Server returned HTTP response code 403 for URL

Check the endpoint URL and the authentication information.

Most likely the API URL or the token is incorrect. They need to use [https://api.snyk.io](https://api.snyk.io) without `/api` . For other tenants, the situation is similar: the path `/api` is not supported on the `api` subdomain.

Even if the URL worked before, the URL is invalid and you must change it.
