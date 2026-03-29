# Tests

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/tests" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/tests -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/tests/{test_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/tests/{test_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/test_jobs/{job_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/test_jobs/{job_id}
```
