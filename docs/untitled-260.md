# Snyk reports an invalid Gradle dependency or dependency version in CLI

##  Snyk reports an invalid Gradle dependency or dependency version in CLI

**Problem**:

Snyk is reporting an incorrect dependency or dependency version when scanning a Gradle project via the CLI

**Discussion**:

When using the Snyk CLI, it obtains the list of dependencies from the gradle wrapper locally on the system where the CLI is being run. You can emulate what Snyk does to fetch these dependencies and compare them to the Snyk Output by running:

`gradle dependencies -q​`​ or `gradlew dependencies -q​`​

Snyk takes the results above and creates a configuration that gathers all of the dependencies from all configurations called `snykMergedDepsConf`.

From that output we create the dependency tree \(depTree\) to scan for vulnerabilities.

**Resolution**:

To resolve this issue, it's possible that Gradle itself is reporting the wrong dependency tree or dependency version. If this is the case, you'll need to troubleshoot your Gradle application to ensure that the version you expect is reported using the commands above.

If you're overriding or forcing a dependency version, you may be doing this in a way that is no longer supported by Gradle. For example, from the Gradle docs:

```text
Forcing dependencies via ExternalDependency.setForce(boolean) is deprecated and no longer recommended: forced dependencies suffer an ordering issue which can be hard to diagnose and will not work well together with other rich version constraints. You should prefer strict versions instead. If you are authoring and publishing a library, you also need to be aware that force is not published.
```

It might be worth trying to use strict versions as suggested [here](https://docs.gradle.org/current/userguide/dependency_downgrade_and_exclude.html#sec:enforcing_dependency_version) to see if this helps.

If you're still certain that this is a Snyk issue, please make sure that you include the output of the following command in your support ticket:

`gradle dependencies -q​`​ or `gradlew dependencies -q​`​

Where possible, you can also share a copy of your **build.gradle** to facilitate reproduction and troubleshooting.

