# Snyk API reports 404 Error: unsupported url

##  Snyk API reports 404 Error: unsupported url

The Snyk API help will be particularly helpful in troubleshooting this error and can be found here: [https://snyk.docs.apiary.io/\#](https://snyk.docs.apiary.io/#)

If you receive a response code 404 with error **unsupported url** it can be due to some missing parameters that are required for the API call to function properly.

For example, if using the Test Maven File API call you may need to verify that you've formed your POST URL correctly. The two parameters required for this call are org and repository, and without them specified in your API call you may receive this 404 error.

Source: [https://snyk.docs.apiary.io/\#reference/test/maven/test-maven-file](https://snyk.docs.apiary.io/#reference/test/maven/test-maven-file)

An example of a well-formed query is included here:

```text
https://snyk.io/api/v1/test/maven?org=9695cbb1-3a87-4d6f-8ae1-61a1c37ee9f7&repository=https%3A%2F%2Frepo1.maven.org%2Fmaven2
```

 **Note**: This link does not work by itself- please substitute the org id with your own and ensure your headers are correctly set.  

