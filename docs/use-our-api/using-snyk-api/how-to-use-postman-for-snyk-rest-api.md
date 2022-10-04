# How to use Postman for Snyk REST API



You can use Postman to see the various endpoints in the Snyk REST API to learn how the API works. Follow these steps to set up Postman for use with the Snyk REST API.

1. Download Postman: [https://www.postman.com/downloads/](https://www.postman.com/downloads/)
2. Go to the Snyk REST API: [https://apidocs.snyk.io/?version=2022-06-08%7Ebeta#overview](https://apidocs.snyk.io/?version=2022-06-08%7Ebeta#overview). Be sure to select the drop down list for the latest version.
   1. Click **Download OpenAPI Spec**.
   2. Save this file locally to your machine.
3. Open Postman: [https://learning.postman.com/docs/getting-started/importing-and-exporting-data/](https://learning.postman.com/docs/getting-started/importing-and-exporting-data/) and follow these steps:
   1. Click on **Collection**
   2. Click the **Import** button
   3. Select the OpenAPI spec file you downloaded from the Snyk REST API.
   4. Accept the default and import.
4. Click on the top level of the Collection you imported and follow these steps:
   1. Under Authorization tab - Type = API Key, Key= `Authorization,` and Value = `token {{snyk_token}}`
   2. Under the Variable tab, add new variables:
      1. Variable = `snyk_api_version`, Initial Value and Current Value = `2022-06-08~beta` - this should be the version for which you downloaded the OpenAPI spec.
      2. Variable = `snyk_token`, Initial Value and Current Value = (enter your service account API token)
   3. Now click **Find and Replace** at the bottom left.
      1. Search for `version=<string>` - it should find all the url with that information.
      2. Replace with `version={{snyk_api_version}}` and this should update all your APIs with the right versions
