# Relationships in Snyk REST API

For some endpoints the schema includes relationships. The following example is for GET /orgs/{org\_id}/targets (Get targets by org ID):

```
relationships: {
org: {
data*: {
id*: uuid
type*: string
}
```

This is defined as “Type of the related resource.”

Relationships describe what a given resource is related to in the API. For example, a project might relate to its target.

See [relationships](https://jsonapi.org/format/#document-resource-object-relationships) in the JSON:API specification for an explanation of the structure.
