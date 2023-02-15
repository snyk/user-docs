# Setting the default Organization for all CLI tests

You can set a default Organization globally for all CLI tests via the CLI. This default Organization will override the Organization that is set as the [**Preferred Organization**](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/manage-snyk-organizations#setting-your-preferred-organization). When entering this command, you can use either the [ID or internal name](finding-the-snyk-id-and-internal-name-of-an-organization.md) of the new default Organization.

**Note**: Regardless of the Organization you set as a global default, you can [run specific CLI tests under a different Organization](setting-an-organization-for-a-specific-cli-test.md).

**To set a default Organization for all CLI tests:**

* In the terminal, enter

```
snyk config set org=<ORG_ID_or_ORG_INTERNAL_NAME>
```

You receive the following confirmation:

```
org updated
```

From now on, all your CLI tests will run under the specified Organization.

For example:

To set the Snyk Demo Org as the default Organization for the CLI tests, we use the Organization ID and enter:

```
snyk config set org=a7708807-3881-xxxx-xxxx-xxxxxxxxxxxx
```

![](<../../../../.gitbook/assets/Snyk Code - CLI - Org - Setting global default.png>)

From now on, all the CLI tests will run by default under the Snyk Demo Org Organization, and the Snyk Demo Org ID will appear in the test results:

![](<../../../../.gitbook/assets/Snyk Code - CLI - Organization - Global Settings - Results - 2.png>)
