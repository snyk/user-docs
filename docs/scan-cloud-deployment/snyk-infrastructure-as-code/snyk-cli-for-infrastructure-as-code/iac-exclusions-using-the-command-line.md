# IaC exclusions using the command line

When you scan directories or a large collection of IaC files using the Snyk CLI `iac test` command, it is easy to include unwanted files or directories in your scan by mistake. When this happens, use your command line tools to exclude specific files or directories from the scan. This page describes some solutions to the most typical use cases.

{% hint style="info" %}
The following examples use command line tools like `find` and `xargs` that are widespread in UNIX-like environments. Make sure that the tools are available on your platform.
{% endhint %}

## Exclude files of the wrong type

Sometimes your Projects contain different kinds of files, and you want to scan only files with a specific extension, excluding everything else. The following command scans only files with a `.tf` extension contained in the current working directory and its subdirectories. Files that do not have a `.tf` extension will not be scanned.

```
find . -type file -name '*.tf' | xargs snyk iac test
```

## Exclude directories by name

In very large projects, it is common to save files that have different purposes in distinct directories. For example, you might save IaC files for your development, staging, and production environments in different directories. When scanning such Projects, you may want to exclude some of these directories. This command scans every file with a `.tf` extension in the current directory and its subdirectories, except those in the `prod` subdirectory.

```
find . -name '*.tf' -not -path './prod/*' | xargs snyk iac test
```
