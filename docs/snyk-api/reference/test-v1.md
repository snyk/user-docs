# Test (v1)

{% hint style="info" %}
This document uses the v1 API, which will eventually be deprecated, as further Snyk developments are now focused on the REST API. For more details, see the [v1 API](../v1-api.md).
{% endhint %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/yarn" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/sbt" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/sbt/{groupId}/{artifactId}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/rubygems" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/rubygems/{gemName}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/pip" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/pip/{packageName}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/npm" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/npm/{packageName}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/maven" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/maven/{groupId}/{artifactId}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/gradle" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/gradle/{group}/{name}/{version}" method="get" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/govendor" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/golangdep" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/dep-graph" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}

{% openapi src="../../.gitbook/assets/spec.yaml" path="/test/composer" method="post" %}
[spec.yaml](../../.gitbook/assets/spec.yaml)
{% endopenapi %}
