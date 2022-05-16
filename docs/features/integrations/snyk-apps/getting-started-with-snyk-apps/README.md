# Getting started with Snyk Apps

**Prerequisites**

To create a Snyk App, you need access to the Snyk API. To get started, follow the steps for [authentication for API](https://docs.snyk.io/snyk-api-info/authentication-for-api).

You also need to retrieve the ID of the Snyk organization you intend the App to be owned by (your **orgId**). You can get this from the organization settings in the Snyk Web UI:

```
https://snyk.io/org/{your-org-name}/manage/settings
```

or by using the `https://snyk.io/api/v1/orgs` API endpoint (with the Authorization/API token from above in the Authorization header).

{% hint style="warning" %}
Snyk Apps have first class access to the API, regardless of whether users installing the App have paid for access or not. To take advantage of this feature, Apps must use API endpoints with the domain **`https://api.snyk.io/`** rather than the conventional **`https://snyk.io/api/`**, when accessing the API within the App.
{% endhint %}
