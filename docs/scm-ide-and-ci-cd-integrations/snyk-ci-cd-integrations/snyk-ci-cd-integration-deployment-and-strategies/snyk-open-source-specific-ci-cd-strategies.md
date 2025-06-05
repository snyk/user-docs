# Snyk Open Source-specific CI/CD strategies

These strategies are useful to teams using the Snyk SCA ([Software Composition Analysis](https://snyk.io/blog/what-is-software-composition-analysis-sca-and-does-my-company-need-it/)) testing features.

Use CLI options like `--fail-on` and `--severity-threshold` to customize the failure status for the build task. For more advanced testing, you can use the `--json` option to generate a JSON file containing the full vulnerability report and set your own build failure status based on the JSON data.

## Gradle and Scala

* For multi-project configurations, test all sub-projects. Use this option with the `monitor` or `test` command: `--all-sub-projects`.
* To scan specific configurations, select certain values of configuration attributes to resolve the dependencies. Use this option with the `test` or `monitor` command: `--configuration-attributes=`.

## Python

* Snyk uses Python to scan and find your dependencies. Snyk needs the Python version to start scanning and defaults to `python`. If you are using multiple Python versions, use the `--command=` option with the `test` or `monitor` command to specify the correct Python command for execution. An example follows:\
  `snyk test --command=python3`
* The `setup.py` file must be targeted. Use the command `snyk test --file=setup.py`
* If you scan a Pip Project and use the `--file=` option because your manifest file is not the standard `requirements.txt`, you must use the following option to specify Pip as the package manager `--package-manager=pip.`

## .NET

If you use a `.sln` file, you can specify the path to the file, and Snyk scans all the sub-projects that are part of the repository, for example:

```
snyk test --file=sln/.sln
```

## Yarn

For Yarn workspaces, use the `--yarn-workspaces` option to `test` and `monitor` your packages. The root lockfile will be referenced for scans of all the packages. Use the `--detection-depth` option to find sub-folders that are not auto-discovered by default.

{% hint style="info" %}
Support for Yarn workspaces is available for the `snyk test` and `snyk monitor` commands only.
{% endhint %}

An example command follows to scan only the packages that belong to any discovered workspaces in the current directory and five sub-directories deep.

```
snyk test --yarn-workspaces --detection-depth=6
```

You can use a common [`.snyk` policy file](../../../manage-risk/policies/the-.snyk-file.md) if you maintain ignores and patches in one place to be applied for all detected workspaces by providing the policy path as follows:

```
snyk test --yarn-workspaces --policy-path=src/.snyk
```

## Monorepo

Some customers have complex Projects, with multiple languages, package managers, and Projects in a single repository. To facilitate this, you can take different approaches:

*   As you build each Project and language, add a directive to run `snyk test` and target a specific Project file, for example:

    ```
    snyk test --file=package.json
    ```
* After you install the dependencies of each Project, make a similar call pointing to the specific artifact, such as `pom.xml`. This is fast and efficient, but can be difficult to scale, especially if you are not familiar with the Project.
* For most Gradle Projects, using `--all-projects` works, as it invokes Gradle-specific options behind the scenes in the form of `snyk test --file=build.gradle --all-sub-projects` when it finds the build file as part of the `--all-projects` search.
* Gradle may require additional configuration parameters. If so, target the other artifacts by using `--file=` for each manifest in the other languages and package managers. You must then use `--all-sub-projects` and potentially `--configuration-matching` to scan a complex Gradle Project.

See [Java and Kotlin](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/) for more information.
