# Use options to customize the snyk test command

The Snyk CLI has many commands that enable you to perform various tasks and many options that enable you to customize the commands. For details, see the [CLI help](../../commands/). This page explains how to customize snyk test to accomplish tasks you may want to do in testing Open Source Projects.

## Scan multiple manifest files

For Projects that have multiple manifest files, specify the file that you want Snyk to inspect for package information by using the `--file` option:

`$ snyk test --file=package.json`

To inspect all of the files, use the `--all-projects` option:

`$ snyk test --all-projects`

## Scan manifest files with custom names

If your manifest files are from a supported language but have a custom name, you can pass the custom name to Snyk by using a combination of the `--file` option and the `--package-manager` option:

`$ snyk test --file=req.txt --package-manager=pip`

## **Test dev dependencies**

Many package managers allow for separately calling out dependencies that are to be used only in a development or test context, that is, not eventually shipped to production. By default, Snyk does not scan these dependencies. If you want your `dev` dependencies to be included in the scan, use the `--dev` option:

`$ snyk test --dev`
