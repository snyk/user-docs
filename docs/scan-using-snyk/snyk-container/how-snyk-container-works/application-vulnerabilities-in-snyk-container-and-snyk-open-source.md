# Application vulnerabilities in Snyk Container and Snyk Open Source

Snyk Container detects application vulnerabilities in your container and overlaps Snyk Open Source capabilities.\
\
The results from the Snyk Container application vulnerability feature and Snyk Open Source are generally the same, especially if Snyk is building a dependency graph from the same manifest files.\
\
However, results can vary significantly depending on the ecosystem and how the developer builds the application. An application in a container is a compiled application. So, in some ecosystems, Snyk Open Source can scan a more detailed manifest and thus build a more accurate dependency graph:

* **`golang` Projects for Snyk Containers**: Snyk does not have access to the list of dependencies as in Snyk Open Source. Therefore, Snyk Container reverse parses binaries, and the result differs slightly from Snyk Open Source.
* **`npm` packages as Snyk Containers**: Snyk can access the list of dependencies. The result is generally the same as in Snyk Open Source. For details, see [Open Source and licensing](../../../supported-languages-and-frameworks/javascript/#open-source-and-licensing).
* **`java` applications for Snyk Containers**: In Open Source, it is possible to include unmanaged jars (see [Scan all unmanaged jar files](../../../snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files.md)). Thus the result is different from Snyk Container.
  * With Snyk Container, the scan traverses all the jars Snyk finds in the image (see [Detecting application vulnerabilities in container images](../../../scan-with-snyk/snyk-container/use-snyk-container-from-the-web-ui/detect-application-vulnerabilities-in-container-images.md)). In addition, there are multiple ways to build a jar, affecting how Snyk Container finds the dependencies.&#x20;
  * In Snyk Open Source, if there are multiple potential versions of a dependency, the package manager dependency resolution logic ensures that only one version is selected; however, in Snyk Container, unpacked jars may contain other versions of dependencies and because they all exist in the container, they are all reported.&#x20;
