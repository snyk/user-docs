# Reporting API (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% hint style="warning" %}
The endpoints in this category support only Snyk legacy reporting, not the latest release. They are not available in regions such as single-tenant or multi-tenant US02, EU, and AU; instead use the [Issues REST API](issues.md).
{% endhint %}
Using the Reporting API, you can find answers to questions like how many issues your Organization has, or how many tests have been conducted in a given time period.

The rate limit is up to **70 requests per minute, per user**. All requests above the limit will get a response with the status code `429 - Too many requests` until requests stop for the duration of the rate-limiting interval (one minute).

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/issues" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/issues/latest" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/counts/tests" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/counts/projects" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/counts/projects/latest" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/counts/issues" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/reporting/counts/issues/latest" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}
