# ServiceAccounts

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts/{serviceaccount_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts/{serviceaccount_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts/{serviceaccount_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts/{serviceaccount_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts/{serviceaccount_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts/{serviceaccount_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts/{serviceaccount_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts/{serviceaccount_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts/{serviceaccount_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts/{serviceaccount_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts/{serviceaccount_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts/{serviceaccount_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets -X POST \
  --input body.json
```
