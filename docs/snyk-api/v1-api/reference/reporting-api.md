# Reporting API

**Note:** The endpoints in this category only support Snyk legacy reporting, not the latest release. As such, they are **not** available in regions such as **single-tenant** or **multi-tenant EU/AU** and you can instead use [the Issues REST API](https://apidocs.snyk.io/#tag--Issues).

The reporting API powers our reports section.

With it you can find answers to questions like how many issues your organisation has, or how many tests have been conducted in a given time frame.

Current rate limit is up to 70 requests per minute, per user. All requests above the limit will get a response with status code `429` - `Too many requests` until requests stop for the duration of the rate-limiting interval (currently a minute).

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/issues/latest" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/issues/" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/counts/issues/latest" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/counts/issues" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/counts/projects/latest" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/counts/projects" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}

{% swagger src="../../../.gitbook/assets/spec.yaml" path="/reporting/counts/tests" method="post" %}
[spec.yaml](../../../.gitbook/assets/spec.yaml)
{% endswagger %}
