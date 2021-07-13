# Snyk for PHP

* [ Language support summary](https://support.snyk.io/hc/en-us/articles/360020352437-Language-support-summary)
* [ Snyk for JavaScript](https://support.snyk.io/hc/en-us/articles/360004712477-Snyk-for-JavaScript)
* [ Snyk for Java \(Gradle, Maven\)](https://support.snyk.io/hc/en-us/articles/360003817357-Snyk-for-Java-Gradle-Maven-)
* [ Snyk for .NET](https://support.snyk.io/hc/en-us/articles/360004519138-Snyk-for-NET)
* [ Snyk for Python](https://support.snyk.io/hc/en-us/articles/360004699377-Snyk-for-Python)
* [ Snyk for Golang](https://support.snyk.io/hc/en-us/articles/360003817417-Snyk-for-Golang)
* [ Snyk for Swift and Objective-C \(CocoaPods\)](https://support.snyk.io/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-)
* [ Snyk for Scala](https://support.snyk.io/hc/en-us/articles/360003781318-Snyk-for-Scala)
* [ Snyk for Ruby](https://support.snyk.io/hc/en-us/articles/360003781298-Snyk-for-Ruby)
* [ Snyk for PHP]()

 [See more](https://support.snyk.io/hc/en-us/sections/360001087857-Language-package-manager-support)

##  Snyk for PHP

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your PHP projects:

### Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|  | [Composer](https://getcomposer.org/) | ✔︎ | ✔︎ | ✔︎ |  |  |

#### Snyk CLI tool for PHP projects

The way by which Snyk analyzes and builds the tree varies depending on the language and package manager of the project.

In order to build the dependency tree Snyk analyzes the `composer.json` and `composer.lock` files that it finds to analyze the dependencies and their versions.

#### Git services for PHP projects

PHP projects can be imported from any of the Git services we support. Once imported, Snyk analyzes your projects based on their supported manifest files.

Once you select a project for import, we build the dependency tree based on these manifest files: 

* Composer.json
* composer.lock

By default, Snyk scans your production dependencies. From the Snyk UI you can configure whether to include your development dependencies \(`require_dev`\) in the scan for vulnerabilities.

**To update language preferences:** 

1. 2. Click on settings ![cog\_icon.png](https://support.snyk.io/hc/article_attachments/4402908592145/cog_icon.png) &gt;  **Languages**.
3. Click **Edit settings** for PHP and select Scan dev dependencies to set for your PHP projects in the specific organization to include both development and production dependencies.
4. Click Update settings.

   These settings will then be applied to all newly imported projects, and once re-tested, to all existing projects.

#### Troubleshooting for your PHP projects

The following error messages may appear for you when working with your PHP projects:

* composer.json or composer.lock not found in path
* Manifest file not found in path
* Lockfile missing packages property
* Lockfile or manifest file is not a valid JSON

If you run across any of these, or other issues, please send the following files to us at [support@snyk.io](mailto:support@snyk.io) and we'll help you out:

* `composer.json`
* `composer.lock`

