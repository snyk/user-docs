# More information about Java support

## Maven Bill of Materials (BOM)

Maven supports [bill of materials (BOM) POM files](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html#bill-of-materials-bom-poms) to centralize dependency versions known to work together.

## Contents and use of BOM files

A BOM file includes:

* a `pom` packaging type: `<packaging>pom</packaging>`.
* a `dependencyManagement` section.

Third-party Projects can provide BOM files to make dependency management easier. Here are some common examples:

* [spring-data-bom](https://github.com/spring-projects/spring-data-bom) - The Spring team provides a BOM for their Spring Data Project.
* [jackson-bom](https://github.com/FasterXML/jackson-bom) - The Jackson Project provides a BOM for Jackson dependencies.

Here is an example of a BOM file:

{% code title="Example 1 - BOM file" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <groupId>snyk</groupId>
    <artifactId>snyk-bom</artifactId>
    <version>1.0</version>
    <packaging>pom</packaging>
    <name>Snyk Bill Of Materials</name>
    
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>log4j</groupId>
                <artifactId>log4j</artifactId>
                <version>1.2.12</version>
            </dependency>
            <dependency>
                <groupId>commons-logging</groupId>
                <artifactId>commons-logging</artifactId>
                <version>1.1.1</version>
            </dependency>
        </dependencies>
    </dependencyManagement>
</project>
```
{% endcode %}

The `dependencyManagement` section contains dependency elements. Each dependency is a lookup reference for Maven to determine the version to select for transitive (and direct) dependencies.

Defining a dependency in the `dependencyManagement` section does not add it to the dependency tree of the Project. It is used only for lookup reference.

You can run `mvn dependency:tree` on the previous BOM example to show that Maven does not treat the contents as dependencies of the file itself.

This BOM can be imported into a Project POM as a parent. The `log4j` version does not need to be specified; it inherits it from the BOM:

{% code title="Example 2 - Project POM" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>snyk</groupId>
        <artifactId>snyk-bom</artifactId>
        <version>1.0</version>
    </parent>
    
    <groupId>snyk</groupId>
    <artifactId>snyk-project</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <packaging>jar</packaging>
    <name>Snyk Project</name>
    
    <dependency>
        <groupId>log4j</groupId>
        <artifactId>log4j</artifactId>
    </dependency>
</project>
```
{% endcode %}

## How Snyk handles BOMs

### Dependency resolution

Snyk applies the versions in the BOM `dependencyManagement` lookup to any dependencies declared in Project POMs that import it as a `parent`.

When Snyk scans the BOM files, the `dependencyManagement` contents are not considered dependencies of that file. These are only lookups.

The following explains how Snyk analyzes and treats each of the previous example files.

* BOM file - Snyk does not create a Snyk Project for this file because it has no dependencies.
* Project POM - Snyk creates a Project with a single dependency of `log4j,` with `v1.2.12.` Snyk applies the rules from the parent BOM to identify the correct version for `log4j`. The dependency `commons-logging` is not included, as it is not directly declared in the Project POM.

{% hint style="info" %}
If a BOM has direct dependencies outside`dependencyManagement`, then Snyk creates a Project for that BOM.
{% endhint %}

### Fix PRs

Snyk offers fix advice, including recommendations for upgrading vulnerable packages with the Fix PR feature.

Fix PRs can only be created for dependencies whose versions are managed in the POM file where the issue is reported.

If the version or dependency is managed in a parent BOM, then even though Snyk sees that it could fix the vulnerable path by changing the version, it cannot apply the fix.

## Package Registry Integrations (Artifactory/Nexus) - Maven

Artifactory and Nexus Package Registry integrations are available to Snyk Enterprise plan users.

* Snyk Open Source uses Artifactory or Nexus to resolve transitive dependencies through private packages.
* Snyk can be connected to a publicly available instance using username and password or a private server on your network using the Snyk Broker.
* Snyk Open Source provides integrations with Artifactory and Nexus both as local gatekeeper, and interacting with the registry for security testing. See [Nexus Repository Manager setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) and [Artifactory Registry setup](../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/).

{% hint style="info" %}
If you are not a Snyk Enterprise user and you use Artifactory or Nexus, analysis is best performed using CLI as the build system will retrieve the dependencies and be present locally.
{% endhint %}
