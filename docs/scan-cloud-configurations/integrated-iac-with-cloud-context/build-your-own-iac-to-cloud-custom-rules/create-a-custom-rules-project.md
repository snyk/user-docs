# Create a Custom Rules Project

## What is a Custom Rules Project?

With the Snyk CLI, you can create a Custom Rules Project, which is a directory tree for your custom rules and rule specs. The CLI relies on its structure for testing, building, bundling, and uploading your custom rules.

To create a Custom Rules Project, you can use the CLI to create an empty Project or clone a pre-populated one, such as the [snyk-labs/iac-to-cloud-example-custom-rules](https://github.com/snyk-labs/iac-to-cloud-example-custom-rules) repository.

### Option 1: Create an empty Project using the Snyk CLI

Use the following command:

<pre><code><strong>snyk iac rules init
</strong></code></pre>

This command introduces a series of interactive prompts and allows you to set up your Project, rule, rule spec (for testing), and resource relations. ​​By selecting `project` and responding to the prompts that follow, you will create a directory tree.

**Example output from interactive prompts:**

```
What do you want to initialize? project
Project name: project1
```

### Option 2: Get the IaC to cloud example Custom Rules repository

```
git clone https://github.com/snyk-labs/iac-to-cloud-example-custom-rules.git
```
