# Users

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/self" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/self
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/users/{id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/users/{id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/users/{id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/users/{id} -X PATCH \
  --input body.json
```
