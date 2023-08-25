# Setting an Organization for a specific CLI test

You can run a specific CLI test under a different Organization from the default. When using this option, the specified Organization will override the default Organization in a specific CLI test. You can use either the [ID or internal name of the Organization](finding-the-snyk-id-and-internal-name-of-an-organization.md) to run a command.

To set an Organization for a specific CLI test in the terminal, after the `test` command enter:

```
--org=<ORG_ID_or_ORG_INTERNAL_NAME>
```

For example, to set the Snyk Test Org as the Organization for a specific CLI test, we use the Organization internal name and enter:

```
snyk code test --org=snyk-xxxx-xxx 
```

The internal name of the Snyk Test Org Organization appears in the results of this test:

![Organization internal name in test results](<../../../../.gitbook/assets/Snyk Code - CLI - Organization - Specific test Settings - Results - 2.png>)
