# Finding the Snyk ID and internal name of an Organization

When setting an Organization for the CLI tests, you can use either the Organization ID or the Organization internal name. These Organization identification details are generated automatically by Snyk for each Organization when it is created. The value you select to enter in the command will be shown as the **Organization** name in the test results. You can find the Snyk ID and internal name on the **Settings** page of the Organization on the Web UI.

Follow these steps to find an Organization ID and internal name:

1\. On the Snyk Web UI, open the Organization whose details you want to find:

<figure><img src="../../../../.gitbook/assets/snyk-org-switcher.png" alt="Open an Organization to find its details"><figcaption><p>Open an Organization to find its details</p></figcaption></figure>

2\. Once the selected Organization is open, click the **Org Settings.**

3\. On the **Settings** page of the Organization, select the **General** tab to find the following:

* **Internal name:** stated in the **Organization name** section.
  * You can change the display name of an Organization, but not the internal name.
  * The internal name also appears in the URL of the Organization.
  * When using the internal name for setting the Organization for the CLI tests, copy the name from the **Settings** page. The internal name is always written in lowercase letters.
* **ID:** appears in the **Organization ID** section. You can use the **Copy** button to copy the ID to the CLI.

<figure><img src="../../../../.gitbook/assets/snyk-org-info.png" alt="Organization name and ID"><figcaption><p>Organization name and ID</p></figcaption></figure>
