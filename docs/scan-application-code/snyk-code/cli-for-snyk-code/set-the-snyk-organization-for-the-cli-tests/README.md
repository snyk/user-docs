# Set the Snyk Organization for the CLI tests

If you have several Organizations in your Snyk account, before you test your code using the CLI, specify which Snyk Organization will be used for the test count.

You can find your available CLI test count on the Organization **Settings** page -> **Usage** tab -> **Test Usage** section -> **Snyk Code** field:

<figure><img src="../../../../.gitbook/assets/snyk-code-usage.png" alt="Snyk Code allowed test usage"><figcaption><p>Snyk Code allowed test usage</p></figcaption></figure>

By default, the CLI runs tests under your **Preferred Organization**, as defined in your **Account settings:**

<figure><img src="../../../../.gitbook/assets/snyk-pref-org.png" alt="Preferred Organization in Snyk Account settings"><figcaption><p>Preferred Organization in Snyk Account settings</p></figcaption></figure>

You can [change your **Preferred Organization**](../../../../snyk-admin/manage-groups-and-organizations/manage-organizations.md) or set another Organization for the CLI tests via the CLI.

When setting an Organization for the CLI tests, you have two options:

* [Set a default Organization globally for all CLI tests](setting-the-default-organization-for-all-cli-tests.md).
* [Set an Organization locally for a specific CLI test](setting-an-organization-for-a-specific-cli-test.md).

When setting an Organization for the CLI tests both globally and locally, you can use either the Organization ID or the Organization internal name. You can [find these on the **Settings** page of the Organization](finding-the-snyk-id-and-internal-name-of-an-organization.md).
