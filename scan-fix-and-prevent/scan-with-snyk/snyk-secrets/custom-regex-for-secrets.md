# Custom RegEx for Secrets

Snyk Secrets lets you define custom regular expressions (regex) to detect secrets that are specific to your Organization, such as database passwords. You can scope a custom regex rule to a **Group** or one or more **Organizations**.

## Prerequisite

To create a custom regex rule, you need the necessary permissions:

| Permission                   | Name                   | What it allows                                                          |
| ---------------------------- | ---------------------- | ----------------------------------------------------------------------- |
| group.rule\_extension.read   | View Rule Extensions   | See the rule extensions (including Secrets custom regexes) in the group |
| group.rule\_extension.create | Create Rule Extensions | Add new rule extensions in the group                                    |
| group.rule\_extension.edit   | Edit Rule Extensions   | Change existing rule extensions in the group                            |
| group.rule\_extension.delete | Delete Rule Extensions | Permanently remove rule extensions from the group                       |

Refer to [user role management](https://docs.snyk.io/platform-administration/user-management/user-role-management) for further information on configuring custom roles.

## Create a custom regex rule

1. Navigate to **Group settings** > **Snyk Secrets**. The **Custom RegEx Rules** section appears.
2. Click **+ New Rule**.
3. In the **RegEx pattern** field, enter your regex pattern.
4. Optionally, add sample strings to the **Test strings** box to validate the pattern.
5. Optionally, add a description in the **Description** box.
6. Click **Save Draft**.
7. Select the Group or one or more Organizations for the rule scope.
8. Click **Save and activate**.

{% hint style="warning" %}
You cannot edit a custom regex rule after you publish it. You must delete and recreate the rule.
{% endhint %}
