# Getting started with the REST API

Follow these steps to make a simple call to the REST API using `curl` in the command line.

1. Log in to [Snyk](https://snyk.io/).
2. In your account, use the left navigation to find an **Organization** where you have Projects you can list.
3. Navigate to your **Organization Settings**, and on the **General** settings page, find your **Organization ID** and copy the value.
4. Navigate to your personal [General Account Settings](https://app.snyk.io/account/) and copy your **API Token**. For instructions, see [Authentication for API](authentication-for-api/).
5. Use a `curl` command to make your request. Replace the `{orgId}` and API\_TOKEN with your **Organization ID** and **API Token**, respectively. For the `version` parameter, Snyk advises using the current day's date. An example follows.

```sh
curl --request GET \
--url "https://api.snyk.io/rest/orgs/{orgId}/projects?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

{% hint style="info" %}
The API URL to use when calling an API is different for different regions. For a complete list, see [API URLs](about-the-rest-api.md#api-urls).

As an example, the`SNYK-US-02`region API URLs are the following:

* **API V1:** https://api.us.snyk.io/v1/
* **REST API:** https://api.us.snyk.io/rest/
{% endhint %}

Note that if you use the parameter `target-reference`, you must URL-encode it.

If you have any problems or questions, contact [Snyk support](https://support.snyk.io).
