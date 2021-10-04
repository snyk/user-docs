# Linking

## Creating a link to Snyk information

If you want to provide a way for customers to link back to see more information about a vulnerability inside Snyk use the following guidelines:

To create a deep link to a customer project in Snyk use this pattern:

* [https://app.snyk.io/org/](https://app.snyk.io/org/)project/

  `Org id` - is returned from the snyk orgs api

  `Project id` - is returned from the snyk projects api call

To create a link to a specific vulnerability use this pattern:

* [https://snyk.io/vuln/](https://snyk.io/vuln)

  Snyk `vuln id` is returned form the snyk issues api, or in the json output of the snyk cli

