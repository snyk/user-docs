# Consistent Ignores for Snyk Code

Snyk Code Consistent Ignores helps your teams focus on important tasks by filtering out distractions. It ensures that once an ignore is created, it is consistently respected regardless of how and where the test is run and what branch is being tested.&#x20;

By filtering out false positives, inapplicable threats, and accepted risks, your security teams can prioritize fixing real problems, and developers can code without interruptions.&#x20;

## Enable Snyk Code Consistent Ignores

Enable Snyk Code Consistent Ignores for your Group or Organization in the Snyk Web UI by navigating to **Group/Organization** > **Settings** > **General** > **Ignores across the repository for Snyk Code**.&#x20;

## Disable Snyk Code Consistent Ignores

Any ignores created or converted with the feature enabled will not be automatically converted back to Project-based ignores. You can recreate them manually after disabling the feature.

## User roles

To create, edit and remove ignores, you need to have a user role assigned with Ignore management permissions. Only Group Admins can set these permissions (see [User role management](../../../../snyk-admin/user-roles/user-role-management.md)).&#x20;

1. Log in to the Snyk Web UI and navigate to your Group and Organization.
2. Navigate to **Members** > **Manage Roles** and select one or more permissions.

<table><thead><tr><th width="203">Ignore management </th><th>Description</th></tr></thead><tbody><tr><td>View Ignores</td><td>View Ignore information.</td></tr><tr><td>Create Ignores</td><td>Create new Ignores.</td></tr><tr><td>Edit Ignores</td><td>Configure Ignores.</td></tr><tr><td>Remove Ignores</td><td>Permanently remove Ignores.</td></tr></tbody></table>

## Manage ignores at the Group level through security policies

You can manage ignores proactively using group-level Snyk Code security policies. As a general rule, you can apply ignore policies when you identify a recurring need to apply similar individual ignores.&#x20;

To manage the ignores through security policies, Snyk Code Consistent Ignores need to be enabled at the Group level. You do not require [conversion](convert-project-scoped-ignores-to-asset-scoped-ignores.md) for any previously applied Group level policy ignores. If you want to activate Snyk Code Security Policies at the Group level, contact your Snyk account team or Snyk Support.

Policies configured to ignore-based Project attributes do not result in ignores being applied in Snyk CLI and IDE settings where a Snyk Project is not available.

| Criteria | Description                                                                                                                                                                                                                                                      |
| -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CWE      | You can provide any CWE value. Snyk Code will ignore any subsequent tests that find associated findings or issues. You can find valid CWEs on [MITRE’s website](https://cwe.mitre.org/data/published/cwe_latest.pdf) or anywhere you can view Snyk Code results. |
| Rule IDs | You can provide any Snyk Code rule ID. You can find Snyk Code rule IDs as part of the SARIF output that you exported using [Snyk CLI](../../../../cli-ide-and-ci-cd-integrations/snyk-cli/).                                                                     |

## Manage ignores in Snyk Projects

You can take action from Project issues, but Snyk will apply any ignores to the underlying asset-scoped findings that can span across Snyk Projects, integrations, and branches.

When you create, modify, or delete an ignore, you must [retest the Project](../../../../scan-with-snyk/snyk-code/manage-code-vulnerabilities/#retesting-code-repository) to update the issue status.&#x20;

An indicator at the top of the Project page will notify you if a retest is needed to capture policy or ignore updates.&#x20;

Project retests typically occur on a nightly or weekly basis, but you can also retest manually.

### Create ignore

1. Log in to the Snyk Web UI and navigate to your Group and Organization.
2.  Open a Project and find an issue card.

    <figure><img src="../../../../.gitbook/assets/issue-card-ignore.png" alt=""><figcaption><p>Issue card overview</p></figcaption></figure>
3. Select **Ignore across repository** on an issue card to create an ignore.
4. Fill in the ignore information and confirm its creation. The issue will be updated and moved from **Open** to **Ignored**.\
   \
   If anyone loads the page before a retest, an indicator will appear and encourage retesting to capture policy or ignore changes.

### Modify ignore

#### Delete ignore

1. Log in to the Snyk Web UI and navigate to your Group and Organization.
2. Open a Project and find an issue card.
3. Select **Unignore** to set all future tests to show the associated finding or issue as open.

#### Edit ignore details

1. Log in to the Snyk Web UI and navigate to your Group and Organization.
2. Open a Project and find an issue card.
3. Select **Edit Ignore**, change the values, and then **Confirm**.
