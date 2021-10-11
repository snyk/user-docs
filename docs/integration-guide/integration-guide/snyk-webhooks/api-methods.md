# API Methods

{% hint style="warning" %}
**Official** Snyk API is available at [https://snyk.docs.apiary.io/#](https://snyk.docs.apiary.io/#)
{% endhint %}

{% swagger baseUrl="https://snyk.io/api/v1/org/:orgId" path="/webhooks" method="post" summary="Create a webhook" %}
{% swagger-description %}
This endpoint allows you to create a webhook.
{% endswagger-description %}

{% swagger-parameter in="path" name="orgId" type="string" %}
Organization ID that uniquely identifies your organization
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
token API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" %}
application/json; charset=utf-8
{% endswagger-parameter %}

{% swagger-parameter in="body" name="url" type="string" %}
Webhook handler endpoint
{% endswagger-parameter %}

{% swagger-parameter in="body" name="secret" type="string" %}
Webhook handler token
{% endswagger-parameter %}

{% swagger-response status="200" description="Webhook successfully created." %}
```javascript
{
    "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​,
    "url": ​"https://my.app.com/webhook-handler/snyk123"​
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://snyk.io/api/v1/org/:orgId" path="/webhooks" method="get" summary="List webhooks" %}
{% swagger-description %}
This endpoint allows you to get a list of all webhooks
{% endswagger-description %}

{% swagger-parameter in="path" name="orgId" type="string" %}
Organization ID that uniquely identifies your organization
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
token API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" %}
application/json; charset=utf-8
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
    "results": [
        {
            "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​, 
            "url": ​"https://my.app.com/webhook-handler/snyk123",​
        },
    ],
    "total": ​1
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId" method="get" summary="Retrieve a webhook" %}
{% swagger-description %}
This endpoint allows you to get details for a specific webhook
{% endswagger-description %}

{% swagger-parameter in="path" name="webhookId" type="string" %}
Snyk webhook ID
{% endswagger-parameter %}

{% swagger-parameter in="path" name="orgId" type="string" %}
Organization ID that uniquely identifies your organization
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
token API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" %}
application/json; charset=utf-8
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```javascript
{
    "id": ​"d3cf26b3-2d77-497b-bce2-23b33cc15362"​,
    "url": ​"https://my.app.com/webhook-handler/snyk123"​
}
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId/ping" method="post" summary="Ping a webhook" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="webhookId" type="string" %}
Snyk webhook ID
{% endswagger-parameter %}

{% swagger-parameter in="path" name="orgId" type="string" %}
Organization ID that uniquely identifies your organization
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
token API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" %}
application/json; charset=utf-8
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
Ok
```
{% endswagger-response %}
{% endswagger %}

{% swagger baseUrl="https://snyk.io/api/v1/org/:orgId" path="/webhooks/:webhookId" method="delete" summary="Delete a webhook" %}
{% swagger-description %}

{% endswagger-description %}

{% swagger-parameter in="path" name="webhookId" type="string" %}
Snyk webhook ID
{% endswagger-parameter %}

{% swagger-parameter in="path" name="orgId" type="string" %}
Organization ID that uniquely identifies your organization
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Authorization" type="string" %}
token API_KEY
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Content-Type" type="string" %}
application/json; charset=utf-8
{% endswagger-parameter %}

{% swagger-response status="200" description="" %}
```
Ok
```
{% endswagger-response %}
{% endswagger %}
