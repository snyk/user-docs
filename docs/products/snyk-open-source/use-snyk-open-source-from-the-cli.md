# Use Snyk Open Source from the CLI

To test your project for known vulnerabilities, browse to the folder containing your project (`cd ~/projects/myproj/`) and run `$ snyk test`.

The `snyk test` command identifies all the local dependencies and queries the Snyk service for known vulnerabilities. `snyk test` displays the issues found along with additional information. For Node.js, Ruby, and Java projects, `snyk test` also suggests steps to fix.

## How Snyk works

Snyk analyzes only your manifest files. Based on the analysis, Snyk creates a hierarchical tree of the structure represented in the manifest file: both its direct and indirect (transitive) dependencies and the points at which the different packages are introduced. Once the tree has been built, Snyk uses its vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree. Using Snyk makes it easier to analyze the project than fixing the project from its source. You can quickly identify the point at which any given vulnerable package was introduced.

When `snyk test` runs, it tries to autodetect your project type by looking for the following files and analyzing the first of the files that it finds. To analyze multiple manifest files see [Monorepos and projects with multiple manifest files](use-snyk-open-source-from-the-cli.md#monorepos-and-projects-with-multiple-manifest-files) on this page. Files that Snyk uses to autodetect the project type include, but are not limited to, the following:

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

The way in which Snyk analyzes the file and builds the tree varies depending on the language and package manager of the project (based on the manifest file type), as well as the method of analysis:

* Using the Snyk CLI
* Importing from a Snyk SCM (Source Control Management) integration

See [Open Source language and package manager support](../../products/snyk-open-source/language-and-package-manager-support/) for more information.

### Monorepos and projects with multiple manifest files

For projects that have multiple manifest files, specify the file that you want Snyk to inspect for package information by using the `--file` option. To identify all of the files, use the `--all-projects` option, for example:

`$ snyk test --file=package.json`; `$ snyk test --all-projects`

### Manifest files with custom names

If your manifest files are from a supported language but have a custom name, you can pass the custom name to Snyk by using a combination of the file option and the package-manager option:

`$ snyk test --file=req.txt --package-manager=pip`

## **Test dev dependencies**

Many package managers allow calling out separately dependencies which are to be used only in a development or test context (that is, not eventually shipped to production). By default Snyk does not scan these dependencies. If you want your dev dependencies to be included in the scan use the `--dev` option:

`$ snyk test --dev`

See [Open Source language and package manager support](../../products/snyk-open-source/language-and-package-manager-support/) for more information concerning supported languages.
