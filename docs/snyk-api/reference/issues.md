# Issues

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/packages/{purl}/issues" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/packages/{purl}/issues
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/packages/issues" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/packages/issues -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/issues" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/issues
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/issues/{issue_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/issues/{issue_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/issues" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/issues
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/issues/{issue_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/issues/{issue_id}
```
