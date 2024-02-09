# How to use Postman for Snyk REST API

You can use Postman to see the various endpoints in the Snyk REST API to learn how the API works. Follow these steps to set up Postman for use with the Snyk REST API.

1. Download Postman: [https://www.postman.com/downloads/](https://www.postman.com/downloads/)
2. Go to the Snyk REST API: [https://apidocs.snyk.io/](https://apidocs.snyk.io/).
   * If you need a specific version, for example, a beta release, select the version from the dropdown at the top of the page.
   * Click **Download OpenAPI Spec**.
   * Save this file locally to your machine.
3. Open Postman: [https://learning.postman.com/docs/getting-started/importing-and-exporting-data/](https://learning.postman.com/docs/getting-started/importing-and-exporting-data/) and follow these steps:
   * Click on **Collection**
   * Click the **Import** button
   * Select the OpenAPI spec file you downloaded from the Snyk REST API.
   * Accept the default and import.
4. Click on the top level of the Collection you imported and follow these steps:
   * Under Authorization tab - Type = API Key, Key= `Authorization,` and Value = `token {{snyk_token}}`
   * Under the Variable tab, add new variables:
     * Variable = `snyk_api_version`, Initial Value and Current Value = `yyyy-mm-dd`. This should be the version for which you downloaded the OpenAPI spec, or today's date, which defaults to the latest spec. If you are using a beta release, ensure you append `~beta` to the date.
     * Variable = `snyk_token`, Initial Value and Current Value = `your service account API token`.
   * Now click **Find and Replace** at the bottom left.
     * Search for `version=<string>` - it should find all the URLs with that information.
     * Replace with `version={{snyk_api_version}}` and this should update all your APIs with the right versions.
