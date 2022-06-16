# Setting the default organization for all CLI tests

You can set a default organization globally for all CLI tests via the CLI. This default organization will override the organization that is set as the [**Preferred Organization**](https://docs.snyk.io/features/user-and-group-management/managing-groups-and-organizations/manage-snyk-organizations#setting-your-preferred-organization). When entering this command, you can use either the [ID or internal name](finding-the-snyk-id-and-internal-name-of-an-organization.md) of the new default organization.

**Note**: Regardless of the organization you set as a global default, you can [run specific CLI tests under a different organization](setting-an-organization-for-a-specific-cli-test.md).

**To set a default organization for all CLI tests:**

* In the terminal, enter

```
snyk config set org=<ORG_ID_or_ORG_INTERNAL_NAME>
```

You receive the following confirmation:

```
org updated
```

From now on, all your CLI tests will run under the specified organization.



For example:

To set the Snyk Demo Org as the default organization for the CLI tests, we use the organization ID and enter:

```
snyk config set org=a7708807-3881-xxxx-xxxx-xxxxxxxxxxxx
```

![](<../../../../.gitbook/assets/Snyk Code - CLI - Org - Setting global default.png>)

From now on, all the CLI tests will run by default under the Snyk Demo Org organization, and the Snyk Demo Org ID will appear in the test results:

![](<../../../../.gitbook/assets/Snyk Code - CLI - Organization - Global Settings - Results.png>)

&#x20;&#x20;
