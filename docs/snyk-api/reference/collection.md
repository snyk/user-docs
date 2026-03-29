# Collection

{% hint style="info" %}
This document uses the REST API. For more details, see the [Authentication for API](../authentication-for-api/) page.
{% endhint %}

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}" method="patch" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id} -X PATCH \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id}
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id} -X DELETE
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}/relationships/projects" method="post" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id}/relationships/projects -X POST \
  --input body.json
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}/relationships/projects" method="get" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id}/relationships/projects
```

{% openapi src="../../.gitbook/assets/rest-spec.json" path="/orgs/{org_id}/collections/{collection_id}/relationships/projects" method="delete" %}
[rest-spec.json](../../.gitbook/assets/rest-spec.json)
{% endopenapi %}

```bash
snyk api /rest/orgs/{org_id}/collections/{collection_id}/relationships/projects -X DELETE \
  --input body.json
```
