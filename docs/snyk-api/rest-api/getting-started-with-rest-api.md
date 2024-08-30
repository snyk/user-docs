---
description: Learn how to make calls to the REST API.
---

# Getting started with REST API

Follow these steps to make a simple call to the REST API using `curl` in the command line.

1. Log in to [Snyk](https://snyk.io/).
2. Navigate to the **Org Settings** (gear icon) for an organization where you have projects you can list.
3. Find the **Organization ID** and copy the value.
4. Navigate to your [General Account Settings](https://app.snyk.io/account/) and copy your **API Token**. For instructions, see See [How to obtain and authenticate with your API token.](../../getting-started/how-to-obtain-and-use-your-snyk-api-token.md)
5. Copy the version string from the URL, for example, `2022-06-08~beta` and paste the version string into the **QUERY-STRING PARAMETERS** `version` field (also required).
6. Use the `curl` command to make your request. Replace the `{orgId}` and API\_TOKEN with your **Organization ID** and **API Token**, respectively. For the `version` parameter, it is advised to use the current day's date.

```sh
curl --request GET \
--url "https://api.snyk.io/rest/orgs/{orgId}/projects?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

Note that if you use the parameter `target-reference`, you must URL-encode it. If you have any problems or questions, contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).
