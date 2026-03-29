# Inventory Assets

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/{asset_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/{asset_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/{asset_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/{asset_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/{asset_id}/relationships/targets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/{asset_id}/relationships/targets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/{asset_id}/relationships/projects" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/{asset_id}/relationships/projects
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/searches" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/searches -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/searches/{search_id}/results" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/searches/{search_id}/results
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/groups" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/groups
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/groups/{group_field_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/groups/{group_field_id}/values
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/filters" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/filters
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/inventory/assets/filters/{filter_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/inventory/assets/filters/{filter_id}/values
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/{asset_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/{asset_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/{asset_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/{asset_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/{asset_id}/relationships/targets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/{asset_id}/relationships/targets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/{asset_id}/relationships/projects" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/{asset_id}/relationships/projects
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/searches" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/searches -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/searches/{search_id}/results" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/searches/{search_id}/results
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/groups" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/groups
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/groups/{group_field_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/groups/{group_field_id}/values
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/filters" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/filters
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/inventory/assets/filters/{filter_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/inventory/assets/filters/{filter_id}/values
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/{asset_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/{asset_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/{asset_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/{asset_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/{asset_id}/relationships/targets" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/{asset_id}/relationships/targets
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/{asset_id}/relationships/projects" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/{asset_id}/relationships/projects
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/searches" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/searches -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/searches/{search_id}/results" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/searches/{search_id}/results
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/groups" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/groups
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/groups/{group_field_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/groups/{group_field_id}/values
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/filters" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/filters
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/inventory/assets/filters/{filter_id}/values" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/inventory/assets/filters/{filter_id}/values
```
