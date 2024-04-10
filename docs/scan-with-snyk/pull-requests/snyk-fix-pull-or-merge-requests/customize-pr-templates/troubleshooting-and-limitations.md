# Troubleshooting and limitations

## Troubleshooting Custom PR templates

If you encounter an error, check your template and follow these instructions:

### Check template syntax

Ensure that the template is written in valid `JSON` syntax for API and `mustache` syntax for the YAML file. Any syntax errors can lead to unexpected behavior. Check the syntax and make any necessary changes. &#x20;

### Verify variable names

Review the variables used in the template and ensure they match the supported variables in the [Variables list and description](variables-list-and-description.md). If a variable is misspelled or not recognized, it may cause unexpected output or a fallback to default values.&#x20;

{% hint style="info" %}
As the feature progresses, Snyk will provide more comprehensive error reporting and validation of the template.
{% endhint %}

## Limitations of Custom PR templates

### Jira

You can add Jira tickets to your PR, but Snyk is limited to those created using Snyk workflows, either UI or API.

### Fail to open PR

Currently, Snyk does not identify errors when a PR cannot be opened. This will be resolved in an upcoming update.

### No line between variables

When you are customizing the YAML file, it is possible that a new line will not be added between variables when two lines start with a variable `{{}}`.

