# Server returned HTTP response code 403 for URL

Check the endpoint URL and the authentication information.

Most likely the API URL or the token is not ok. They need to use [https://app.snyk.io/api](https://app.snyk.io/api) or [https://api.snyk.io](https://api.snyk.io/) without `/api` . For other tenants, the situation is similar: the path `/api` is not supported on the `api` subdomain. It only works on the `app` subdomain.

**Even if the URL may have worked before, the URL is invalid and must be changed.**

\
