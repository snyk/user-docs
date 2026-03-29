# Groups (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/tags" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/tags
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/tags/delete" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/tags/delete -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/settings" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/settings -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/settings" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/settings
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/roles" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/roles
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/orgs" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/orgs
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/org/{orgId}/members" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/org/{orgId}/members -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/group/{groupId}/members" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/group/{groupId}/members
```
