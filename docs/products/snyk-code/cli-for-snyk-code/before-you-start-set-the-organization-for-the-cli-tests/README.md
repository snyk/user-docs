# Before you start – Set the organization for the CLI tests

If you have several organizations in your Snyk Account, before starting to use the CLI, you should define which organization will be used for the test count when running tests via the CLI.

**Note**: You can find information about your test count in the organization **Settings** page -> **Usage** tab -> **Test Usage** section -> **Snyk Code** field:

![](<../../../../.gitbook/assets/snyk Code - CLI - Test Count.png>)

By default, the CLI will run tests under your **Preferred Organization**, as defined in your **Account settings:**

![](<../../../../.gitbook/assets/snyk Code - CLI - Organization - Preferred Org.png>)

You can [change your **Preferred Organization**](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/manage-snyk-organizations#setting-your-preferred-organization), or set another organization for the CLI tests via the CLI.

When setting an organization for the CLI tests you have two options:

* [Set a default organization globally for all CLI tests](setting-the-default-organization-for-all-cli-tests.md).
* [Set an organization locally for a specific CLI test](setting-an-organization-for-a-specific-cli-test.md).

When setting an organization for the CLI tests, both globally and locally, you can use either the organization ID or the organization internal name. You can [find these details in the **Settings** page of the required organization on the Snyk Web U](finding-the-snyk-id-and-internal-name-of-an-organization.md)I.

