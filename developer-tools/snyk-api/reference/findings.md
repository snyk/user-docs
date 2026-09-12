---
description: Snyk API reference for the Findings endpoints, including request parameters and response schemas
---

# Findings

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/tests/{test_id}/findings" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/tests/{test_id}/components/{component_id}/findings" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/findings/documents" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/documents/{document_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/document_jobs/{job_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}
