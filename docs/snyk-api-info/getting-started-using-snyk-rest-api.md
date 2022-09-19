# Getting started using Snyk REST API

Follow these steps to make a simple call to the Snyk REST API.

1. Verify that your permissions are set correctly for you to use Snyk API.
2. Log in to [Snyk](https://snyk.io/).
3. Navigate to the **Org Settings** (gear icon) for an organization where you have projects you can list.
4. Find the **Organization ID** so you can copy the value when you make the API call.
5. Navigate to the Snyk REST API and authenticate. For instructions, see See [Authentication for API](authentication-for-api.md).
6. Observe that the Snyk REST API opens to the most recent beta version.
7. Look for the **Projects** endpoint and navigate to the **GET** `/orgs/{org_id}/projects` (Get project by org ID) call.
8. Copy your **Organization ID** from your **Org settings** and paste the ID into the **PATH PARAMETERS** `org_id` field (asterisk designates a required request parameter).
9. Copy the version string from the URL, for example, `2022-06-08~beta` and paste the version string into the **QUERY-STRING PARAMETERS** `version` field (also required).
10. Click **TRY**.
11. Scroll to the **RESPONSE**, which is a 200 response (provided that your permissions and your **Organization ID** are correct), showing a list of projects, along with RESPONSE HEADERS and CURL to copy . If you get a different status code, do what is needed to resolve the error.
12. Observe the optional **QUERY-STRING PARAMETERS** to see the range of data you can retrieve from the Snyk REST API.

If you have any problems or questions, contact [Snyk support](https://support.snyk.io/hc/en-us/requests/new).
