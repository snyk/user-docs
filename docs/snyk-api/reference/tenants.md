# Tenants

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/memberships" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/memberships
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/memberships/{membership_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/memberships/{membership_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/memberships/{membership_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/memberships/{membership_id} -X DELETE
```
