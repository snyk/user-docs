# Dependencies (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

Dependencies are packages or modules that your Projects depend on.

The rate limit is up to **150 requests per minute, per user**.

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/dependencies" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}
