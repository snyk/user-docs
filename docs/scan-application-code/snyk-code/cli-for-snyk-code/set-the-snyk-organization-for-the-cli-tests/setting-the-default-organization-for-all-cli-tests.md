# Setting the default Organization for all CLI tests

You can set a default Organization globally for all CLI tests via the CLI. This default Organization will override the Organization set as the [preferred Organization](../../../../snyk-admin/manage-groups-and-organizations/manage-organizations.md). When entering this command, you can use either the [ID or internal name](finding-the-snyk-id-and-internal-name-of-an-organization.md) of the new default Organization.

Regardless of the Organization you set as a global default, you can [run specific CLI tests under a different Organization](setting-an-organization-for-a-specific-cli-test.md).

To set a default Organization for all CLI tests, In the terminal, enter

```
snyk config set org=<ORG_ID_or_ORG_INTERNAL_NAME>
```

You receive the following confirmation:

```
org updated
```

From now on, all your CLI tests will run under the specified Organization.

For example, to set the Snyk Demo Org as the default Organization for the CLI tests, use the Organization ID and enter:

```
snyk config set org=a7708807-3881-xxxx-xxxx-xxxxxxxxxxxx
```

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - Org - Setting global default.png" alt="Command to set Organization and confirmation"><figcaption><p>Command to set Organization and confirmation</p></figcaption></figure>

From now on, all the CLI tests will run by default under the Snyk Demo Org Organization, and the Snyk Demo Org ID will appear in the test results:

<figure><img src="../../../../.gitbook/assets/Snyk Code - CLI - Organization - Global Settings - Results - 2.png" alt="Snyk Demo Org ID in test results"><figcaption><p>Snyk Demo Org ID in test results</p></figcaption></figure>
