# Troubleshooting and limitations

## Troubleshooting

If you encounter an error, double-check your template and apply the below instructions:

### Check template syntax

Ensure that the template is written in valid mustache syntax. Any syntax errors can lead to unexpected behavior. Double-check the syntax and make any necessary changes. &#x20;

### Verify variable names

Review the variables used in the template and ensure they match the supported variables provided in the [Variables list](troubleshooting-and-limitations.md#variables-list-and-description). If a variable is misspelled or not recognized, it may result in unexpected output or fallback to default values.&#x20;

{% hint style="info" %}
As the feature progresses, Snyk will provide more comprehensive error reporting and validation of the template.
{% endhint %}

## Limitations

#### Branch names&#x20;

Although you can still customize your branch name, to avoid any potential issues, Snyk automatically appends a hash to the end of the branch name you choose. This hash derives from the PR contents and prevents opening duplicate branches, as well as naming conflicts with existing or future branches.

#### Jira

You can add Jira tickets to your PR, but Snyk is limited to those created via Snyk workflows (either UI or API).

#### Fail to open PR

Currently, Snyk does not surface errors when a PR cannot be opened. This will be resolved in an upcoming update.

#### No line between variables

It is possible that a new line is not added between variables when two lines start with a variable `{{}}`.

