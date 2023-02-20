# A .snyk policy file in a different directory from the manifest file

When you scan a project with the CLI, the `.snyk` policy file may be in a different directory from the manifest file, either because of the structure of the project or because the project has multiple manifest files using the same policy file.

Use the `--policy-path` option to specify a policy path manually for use in scanning or monitoring the project.
