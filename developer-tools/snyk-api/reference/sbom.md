# SBOM

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/sbom_tests" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/sbom_tests/{job_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/sbom_tests/{job_id}/results" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/projects/{project_id}/sbom" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}
