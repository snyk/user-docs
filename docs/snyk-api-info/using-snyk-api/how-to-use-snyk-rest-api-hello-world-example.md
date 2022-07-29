# How to use Snyk REST API hello world example

Follow these steps to use the hello world example.

1. Verify that your permissions are set correctly for you to use Snyk API.
2. Log in to [Snyk](https://snyk.io/).
3. Navigate to the Snyk REST API and authenticate. For instructions, see See [Authentication for API](../authentication-for-api.md).
4. Observe that the Snyk REST API opens to the most recent beta version.
5. Navigate to the **Examples** endpoints and start with **POST** `/examples/hello_world` (Create a single result from the hello\_world example).
6. Copy the version string from the URL, for example, `2022-06-08~beta` and paste the version string into the **QUERY-STRING PARAMETERS** `version` field (required).
7. Click the **REQUEST BODY EXAMPLE** tab. Do the following to fill the text area.
   1. Scroll to the **RESPONSE SCHEMA** and click the **EXAMPLE** tab.
   2. Copy the response example and paste it into the **REQUEST BODY EXAMPLE**.
8. Click **TRY**.
9. Copy an `id` from the RESPONSE and continue with **GET** `/examples/hello_world/{id}` (Get a single result from the hello\_world example).
10. Paste the `id` into the the **PATH PARAMETERS** `id` field (required).
11. Copy the version string from the URL, for example, `2022-06-08~beta` and paste the version string into the **QUERY-STRING PARAMETERS** `version` field (also required).
12. Click **TRY**.
13. Scroll to the **RESPONSE**, which is a 200 response.

If you have any problems or questions, contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).
