# Use Snyk Open Source from the CLI

To test your project for known vulnerabilities:

* Go to the folder containing your project (`cd ~/projects/myproj/`)
* Run `$ snyk test`.

The `snyk test` command identifies all the local dependencies and queries the Snyk service for known vulnerabilities. `snyk test` displays the issues found along with additional information.

{% hint style="info" %}
For Node.js, Ruby, and Java projects, `snyk test` also suggests steps to fix.
{% endhint %}

### How it works

Snyk Open Source scans your manifest files. Based on the scan, Snyk creates a hierarchical tree of the structure represented in the manifest file: both its direct and indirect (transitive) dependencies and the points at which the different packages are introduced.

After this tree is built, Snyk uses its vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree. Using Snyk makes it easier to analyze the project than fixing the project from its source. You can quickly identify the point at which any given vulnerable package was introduced.

#### Snyk Open Source file types

When `snyk test` runs, it tries to autodetect your project type by looking for the following files and scanning the first of the files that it finds (to analyze multiple manifest files, see [scan multiple manifest files](./#scan-multiple-manifest-files)).

Files that Snyk uses to autodetect the project type include, but are not limited to, the following:

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

The way in which Snyk analyzes the file and builds the tree varies depends on:

* The [language and package manager](../snyk-open-source-supported-languages-and-package-managers/) you use (as determined by the manifest file type)
* The method of scanning (using the [Snyk CLI](../../../snyk-cli/), or by importing from a Snyk [Git repository integration](../../../integrations/git-repository-scm-integrations/)).

### Scan multiple manifest files

For projects that have multiple manifest files, specify the file that you want Snyk to inspect for package information by using the `--file` option:

`$ snyk test --file=package.json`

To identify all of the files, use the `--all-projects` option:

`$ snyk test --all-projects`

### Manifest files with custom names

If your manifest files are from a supported language but have a custom name, you can pass the custom name to Snyk by using a combination of the file option and the package-manager option:

`$ snyk test --file=req.txt --package-manager=pip`

### **Test dev dependencies**

Many package managers allow calling out separately dependencies which are to be used only in a development or test context (that is, not eventually shipped to production). By default Snyk does not scan these dependencies. If you want your dev dependencies to be included in the scan use the `--dev` option:

`$ snyk test --dev`

See [Open Source language and package manager support](../snyk-open-source-supported-languages-and-package-managers/) for more information concerning supported languages.
