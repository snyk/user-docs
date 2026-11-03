---
description: How to make your first Snyk REST API call with curl, including finding your Organization ID and authentication token
nav_context: agnostic
---

# Getting started with the REST API

Follow these steps to make a simple call to the REST API using `curl` in the command line.

You need `curl` and access to a Snyk Organization that has Projects.

1. Log in to [Snyk](https://snyk.io/).
2. In your account, use the left navigation to find an **Organization** where you have Projects you can list.
3. Navigate to your **Organization Settings**, and on the **General** settings page, find your **Organization ID** and copy the value.
4. Navigate to your personal [General Account Settings](https://app.snyk.io/account/) and copy your **API Token**. For instructions, visit [Authentication for API](https://docs.snyk.io/developer-tools/snyk-api/authentication-for-api).
5. Use a `curl` command to make your request. Replace `{orgId}` and `API_TOKEN` with your **Organization ID** and **API Token**, respectively. Snyk recommends using 2024-10-15 for the version number unless you are using an earlier version for a specific reason. Using the current day's date calls the most recent version of the API. An example follows:

```sh
curl --request GET \
--url "https://api.snyk.io/rest/orgs/{orgId}/projects?version=2024-10-15" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

A successful call returns the `200` status code with a `data` array listing the Projects in the Organization.

{% hint style="info" %}
The API URL to use when calling an API is different for different regions. For a complete list, visit [API URLs](https://docs.snyk.io/developer-tools/snyk-api/rest-api/about-the-rest-api#api-urls).

For example, the `SNYK-US-02` region API URLs are the following:

* API V1: https://api.us.snyk.io/v1/
* REST API: https://api.us.snyk.io/rest/
{% endhint %}

If you use the `target-reference` parameter, you must URL-encode it.

If you have any problems or questions, contact [Snyk support](https://support.snyk.io).
