# API - Maven test using a different repository

By default, API requests to test Maven projects use [https://repo1.maven.org/maven2/](https://repo1.maven.org/maven2/) as the lookup store. 

When testing a Maven project, you can specify the repository store in the URL of the API request.

This can be done in Apiary for this command \(for example\):

https://snyk.io/api/v1/test/maven/

* repository  - `https://repository.jboss.org/nexus/content/repositories`
* groupId     - `org.jboss.resteasy`
* artifactId   - `resteasy-yaml-provider`
* version      - `3.9.3.Final`

 With the resulting URL looking like this:

```text
https://snyk.io/api/v1/test/maven/org.jboss.resteasy/resteasy-yaml-provider/3.9.3.Final?repository=https%3A%2F%2Frepository.jboss.org%2Fnexus%2Fcontent%2Frepositories
```

**Note**: don't forget to set your API headers and Snyk token 

