# Use Snyk Open Source from the CLI

To test your project for known vulnerabilities, browse to your project’s folder and run `snyk test`:

```text
cd ~/projects/myproj/
```

`snyk test` takes stock of all the local dependencies and queries the Snyk service for known vulnerabilities. It displays the found issues along with additional information. For Node.js, Ruby, Java projects, it also suggests remediation steps.

## How it works

Snyk analyzes only your manifest files, based on which we then create a hierarchical tree that represents the structure represented in the manifest file, both its direct and indirect \(transitive\) dependencies and the points at which the different packages are introduced. Once we’ve built the tree, we can use our vulnerability database to find vulnerabilities in any of the packages anywhere in the dependency tree. The entire project can be analyzed more easily, and you can quickly identify the point at which any given vulnerable package was actually introduced - then fixing it from its source.

When `snyk test` runs, it tries to autodetect your project type by looking for the following files, in this order, and then analyzing the first of the files that it finds:

{% hint style="info" %}
**NOTE**  
To analyze multiple manifest files, manually specify the file that Snyk should inspect for package information, as described later in this article.
{% endhint %}

1. yarn.lock
2. package-lock.json
3. package.json
4. Gemfile.lock
5. pom.xml
6. build.gradle
7. build.sbt
8. Pipfile
9. requirements.txt
10. Gopkg.lock
11. vendor/vendor.json
12. obj/project.assets.json
13. packages.config
14. composer.lock
15. build.gradle.kts
16. go.mod

The way by which Snyk analyzes and builds the tree then varies depending on the language and package manager of the project \(based in the manifest file type\), as well as the location of your project:

* analysis with our CLI tool 
* import from our app

{% hint style="info" %}
**NOTE**  
Additional and more specific details are provided per language, in [Language support](https://support.snyk.io/hc/en-us/categories/360000456257-Language-package-manager-support).
{% endhint %}

_Monorepos and projects with multiple manifest files_

For projects that have multiple manifest files, you can specify the file that Snyk should inspect for package information using the file flag:

```text
$ snyk test --file=package.json
```

_Manifest files with custom names_

If your manifest files are from a supported language but have a custom name you can let Snyk know about it by using a combination of the file flag and the package-manager flag:

```text
$ snyk test --file=req.txt --package-manager=pip
```

## **Test dev dependencies**

Many package managers allow calling out separately dependencies which are to be used only in a development/test context \(i.e don't get eventually shipped to production\). By default Snyk does not scan these dependencies. If you want your dev dependencies to be included in the scan use the dev flag:

```text
$ snyk test --dev
```

{% hint style="info" %}
**NOTE**  
Additional and more specific details are provided per language, in [Language support](https://support.snyk.io/hc/en-us/categories/360000456257-Language-package-manager-support).
{% endhint %}

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page/)
{% endhint %}

