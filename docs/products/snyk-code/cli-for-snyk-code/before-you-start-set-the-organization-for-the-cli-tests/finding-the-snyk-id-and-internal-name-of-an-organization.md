# Finding the Snyk ID and internal name of an organization

When setting an organization for the CLI tests, you can use either the organization ID or the organization internal name. These organization identification details are generated automatically by Snyk for each organization upon its creation. The value you select to enter into the command, will be shown as the **Organization** name in the test results. You can find the Snyk ID and internal name in the **Settings** page of the organization on the Web UI.

**To find an organization ID and internal name:**

1\. On the Snyk Web UI, open the organization whose details you want to find:

![](<../../../../.gitbook/assets/Snyk Code - CLI - Org - Selecting from UI.png>)

2\. Once the selected organization is open, click the **Org Settings** button <img src="../../../../.gitbook/assets/Snyk Code - CLI - Org Settings button - Icon.png" alt="" data-size="line"> on the top menu:

![](<../../../../.gitbook/assets/Snyk Code - CLI - Org Settings button.png>)

3\. On the **Settings** page of the organization, select the **General** tab on the left. There, you can find the following organization details:

* **Internal name** – appears in the **Organization name** section with a gray background, in the paragraph above the display name.\
  **Notes**:
  * You can change the display name of an organization, but not the internal name.
  * The internal name also appears in the URL of the organization.
  * When using the internal name for setting the organization for the CLI tests, copy the name from the **Settings** page. The internal name is always written in lower case letters.
* **ID** - appears in the **Organization ID** section. You can use the **Copy** button to copy the ID to the CLI.

![](<../../../../.gitbook/assets/Snyk Code - CLI - Org - Details.png>)
