# Ignores (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/project/{projectId}/ignores" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/project/{projectId}/ignores
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/project/{projectId}/ignore/{issueId}" method="put" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/project/{projectId}/ignore/{issueId} -X PUT \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/project/{projectId}/ignore/{issueId}" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/project/{projectId}/ignore/{issueId} -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/project/{projectId}/ignore/{issueId}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/project/{projectId}/ignore/{issueId}
```

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/org/{orgId}/project/{projectId}/ignore/{issueId}" method="delete" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

```bash
snyk api /v1/org/{orgId}/project/{projectId}/ignore/{issueId} -X DELETE
```
