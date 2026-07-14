# Test (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/yarn" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/sbt" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/sbt/{groupId}/{artifactId}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/rubygems" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/rubygems/{gemName}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/pip" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/pip/{packageName}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/npm" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/npm/{packageName}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/maven" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/maven/{groupId}/{artifactId}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/gradle" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/gradle/{group}/{name}/{version}" method="get" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/govendor" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/golangdep" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/dep-graph" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/v1-api-spec.yaml" path="/test/composer" method="post" %}
[v1-api-spec.yaml](../../.gitbook/assets/v1-api-spec.yaml)
{% endopenapi %}
