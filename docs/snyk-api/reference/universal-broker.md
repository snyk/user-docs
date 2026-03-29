# Universal Broker

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/credentials/{credential_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/contexts
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/deployments/{deployment_id}/connections/{connection_id}/bulk_migration
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integrations/{integration_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integrations/{integration_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integration" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/contexts/{context_id}/integration -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/installs/{install_id}/connections/{connection_id}/contexts" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/installs/{install_id}/connections/{connection_id}/contexts
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/deployments" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/deployments
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integrations/{integration_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integrations/{integration_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integration" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/connections/{connection_id}/orgs/{org_id}/integration -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/tenants/{tenant_id}/brokers/connections/{connection_id}/integrations" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/tenants/{tenant_id}/brokers/connections/{connection_id}/integrations
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/brokers/connections" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/brokers/connections
```
