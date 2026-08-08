---
description: How Snyk Container and Snyk Open Source detect application vulnerabilities in images
nav_context: agnostic
---

# Application vulnerabilities in Snyk Container and Snyk Open Source

Snyk Container detects application vulnerabilities in your containers and overlaps Snyk Open Source capabilities.\
\
The results from a Snyk Container application vulnerability scan and a Snyk Open Source scan are generally the same, especially if Snyk is building a dependency graph from the same manifest files.\
\
However, results can vary significantly depending on the ecosystem and how the developer builds the application. An application in a container is a compiled application. So, in some ecosystems, Snyk Open Source can scan a more detailed manifest and thus build a more accurate dependency graph:

* `golang` Projects for Snyk Containers: Snyk does not have access to the list of dependencies as in Snyk Open Source. Therefore, Snyk Container reverse parses binaries, and the result differs slightly from Snyk Open Source. Snyk Container also reports vulnerabilities in the Go standard library, identified from the Go version recorded in the binary (see [Go standard library vulnerabilities](#go-standard-library-vulnerabilities)).
* `npm` packages as Snyk Containers: Snyk can access the list of dependencies. The result is generally the same as that of Snyk Open Source. For details, see [Support for npm](https://app.gitbook.com/s/L7HyJj9FsK1W4pNt8Gzl/supported-languages/supported-languages-list/javascript#support-for-npm).
* `java` applications for Snyk Containers: In Open Source, it is possible to include unmanaged jars (see [Scan all unmanaged jar files](https://app.gitbook.com/s/IEEjSXQQu36y0vmFV8zf/snyk-cli/snyk-cli/scan-and-maintain-projects-using-the-cli/scan-all-unmanaged-jar-files)). Thus the result is different from Snyk Container.
  * With Snyk Container, the scan traverses all the jars Snyk finds in the image (see [Detect application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md)). In addition, there are multiple ways to build a jar, and the method used affects how Snyk Container finds the dependencies.
  * In Snyk Open Source, if there are multiple potential versions of a dependency, the package manager dependency resolution logic ensures that only one version is selected. However, in Snyk Container, unpacked jars may contain other versions of dependencies, and because they all exist in the container, they are all reported.

## Go standard library vulnerabilities

Snyk Container reports vulnerabilities from the Go standard library in addition to third-party modules.

### How it works

Snyk derives the Go version from the toolchain version compiled into the binary. Scans report all standard library package vulnerabilities that affect that Go version (for example, from `fmt` or `net/http`), regardless of which libraries your application's code imports or calls. Snyk does not perform any reachability analysis for Go standard library vulnerabilities.

### What this means for your results

Expect to see vulnerabilities from Go standard library packages on your container Projects. These are valid, although not all of them may be reachable by your application's code. Snyk Container has not always reported these vulnerabilities, so expect to see an increase in vulnerabilities on your existing Projects. To address this, you can:

* **Remediate:** Rebuild the binary with a Go toolchain version that fixes the reported vulnerabilities.
* **Reduce noise:** [Ignore a reported standard library vulnerability](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/README.md) if you determine it is not relevant to your application.
