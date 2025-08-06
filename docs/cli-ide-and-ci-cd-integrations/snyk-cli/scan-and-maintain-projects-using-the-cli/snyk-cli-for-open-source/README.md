# Snyk CLI for Open Source

Snyk Open Source scans your manifest files. Based on the scan, Snyk creates a hierarchical tree of the structure represented in the manifest file: both its direct and indirect (transitive) dependencies and the points at which the different packages are introduced.

After this tree is built, Snyk uses its vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree. Using Snyk makes it easier to analyze the Project than fix the Project from its source. You can quickly identify the point at which any vulnerable package was introduced.

## Test an Open Source Project

To test your Project for known vulnerabilities:

* Navigate to the folder containing your Project (`cd ~/projects/myproj/`)
* Run `$ snyk test`.

The `snyk test` command identifies all the local dependencies and queries the Snyk service for known vulnerabilities. `snyk test` displays the issues found along with additional information. For information about the snyk test results, see [Review the Snyk Open Source CLI results](review-the-snyk-open-source-cli-results.md).

{% hint style="info" %}
For Node.js, Ruby, and Java Projects, `snyk test` also suggests fix steps.
{% endhint %}

## Files Snyk uses to detect the Project type

When `snyk test` runs, it tries to autodetect your Project type by looking for a manifest file and scanning the first of these files that it finds. Files that Snyk uses to autodetect the Project type include, but are not limited to, the following:

* yarn.lock
* package-lock.json
* package.json
* Gemfile.lock
* pom.xml
* build.gradle
* build.sbt
* Pipfile
* requirements.txt
* Gopkg.lock
* vendor/vendor.json
* obj/project.assets.json
* packages.config
* composer.lock
* build.gradle.kts
* go.mod

{% hint style="info" %}
&#x20;To analyze multiple manifest files, see [Scan multiple manifest files](use-options-to-customize-the-snyk-test-command.md#scan-multiple-manifest-files).
{% endhint %}

The way Snyk analyzes the file and builds the tree varies depending on the following:

* The [language and package manager](../../../../supported-languages-package-managers-and-frameworks/) you use, as determined by the manifest file type
* The method of scanning, using the [Snyk CLI](../../), or importing a Project using a [Git repository integration](../../../../scm-integrations/organization-level-integrations/)

See [Use options to customize the snyk test command](use-options-to-customize-the-snyk-test-command.md) for tips on running `snyk test` with commonly used options.

See [Supported languages, package managers, and frameworks](../../../../supported-languages-package-managers-and-frameworks/) for more information about supported languages.
