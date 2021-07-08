# Apiary - API server took too long to respond

Some API requests may take longer than the default timeout in Apiary.

In a case like this you could use a command like this to get your results:

```text
curl -s -N -X GET --header "content-type:application/json; charset=utf-8" --header "authorization:token << your Snyk API token >>" https://snyk.io/api/v1/test/maven/jboss/jboss/3.2.1
```

