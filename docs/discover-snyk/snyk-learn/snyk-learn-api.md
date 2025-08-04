# Snyk Learn API

{% hint style="info" %}
**Release status**&#x20;

The Snyk Learn API endpoints are [Beta](../../snyk-api/rest-api/about-the-rest-api.md#versioning) endpoints.
{% endhint %}

The Snyk Learn API endpoints are part of the [Snyk REST API](../../snyk-api/rest-api/about-the-rest-api.md) and allow for programmatic interaction with Snyk Learn.

The Snyk REST API requires authentication. For information about authentication, see [Authentication for API](../../snyk-api/authentication-for-api/).

The following table shows the Snyk Learn API endpoints available for each plan. For more information about plans, see [Plans and pricing](https://snyk.io/plans/).

<table><thead><tr><th width="294">Plan</th><th>Features</th></tr></thead><tbody><tr><td>Free and Team</td><td><ul><li>View the <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/learn/catalog">content catalog</a>.</li></ul></td></tr><tr><td>Enterprise</td><td><ul><li>Track learning progress for <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/users">individuals</a> in an Organization.</li><li>Track overall learning progress for your <a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/progress/catalog">Organization</a>.</li></ul></td></tr><tr><td>Learning Management add-on</td><td><ul><li><a href="https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/learn/assignments">Manage assignments</a>.</li></ul></td></tr></tbody></table>
