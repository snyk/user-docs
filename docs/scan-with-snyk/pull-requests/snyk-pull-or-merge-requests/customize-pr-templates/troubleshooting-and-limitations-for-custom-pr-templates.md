# Troubleshooting and limitations for custom PR templates

## Troubleshooting custom PR templates

If you encounter an error, check your template and follow these instructions.

### Check template syntax

Ensure that the template is written in valid `JSON` syntax for API and `mustache` syntax for the YAML file. Any syntax errors can lead to unexpected behavior. Check the syntax and make any necessary changes. &#x20;

### Verify variable names

Review the variables used in the template and ensure they match the supported variables in the [Variables list and description](variables-list-and-description.md). If a variable is misspelled or not recognized, it may cause unexpected output or a fallback to default values.&#x20;

### Combination of YAML and API template

If you open a custom PR template with values from both the YAML upload template and the API template, this is the expected behavior. Refer to the [order of precedence for processing custom PRs](./). If you do not want to see this combination, depending on your preference,  either remove the YAML  template or delete your API template.&#x20;

## Limitations of custom PR templates

### Jira tickets in your PR

You can add Jira tickets to your PR, but Snyk is limited to those created using Snyk workflows, either UI or API.

### Failure to open PR

If a PR cannot be opened, Snyk cannot identify errors when a PR cannot be opened.&#x20;

### No line between variables

When you are customizing the YAML file, it is possible that a new line will not be added between variables when two lines start with a variable `{{}}`.

