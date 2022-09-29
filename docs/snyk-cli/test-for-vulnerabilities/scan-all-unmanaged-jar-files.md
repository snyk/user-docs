# Scan all unmanaged JAR files

The Snyk CLI can scan unmanaged JAR files in [Java applications](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven.md). The CLI identifies the package name, version, and vulnerabilities only if the local JAR file hash matches the Maven Central JAR file hash.

Java apps typically have JAR files in a number of locations within an application. To avoid problems from scanning multiple JAR files in the same folder, scan individual JAR files, especially for old Java apps that use Ant.

**WAR file support**: You can scan individual WAR files that are published in Maven Central. To scan open-source dependency JARs directly, you must extract (unzip) all other WAR files or JAR files containing other JARs.

**Prerequisite:** Scanning individual files requires **Maven 3.1.0** or newer to be installed alongside the Snyk CLI, which requires **maven-dependency-plugin 2.2** or higher.

To view the versions of Maven and the maven-dependency-plugin run `mvn -v`.

![Screenshot of output from mvn -v](../../.gitbook/assets/untitled\_\_1\_.png)

Use the `snyk test --scan-all-unmanaged` command to scan all JAR files in a single folder to match any dependencies hosted on Maven central. Scanning all unmanaged files may lead to dependency conflicts if an application was not built using a package manager like Gradle or Maven. This is especially true for applications built using Ant. Therefore test each JAR file individually using the following command:\
`snyk test --scan-unmanaged --file=/path/to/file`

Testing each JAR file individually also has the benefit of showing the name of the JAR file that was scanned on the Snyk web UI; running a scan using --`scan-all-unmanaged` does not show the file. A simple wrapper is required in order to find and test JAR files in all sub-folders of an application. Results can then be grouped in the Snyk UI by using the --`remote-repo-url=AppName` option for `snyk test`.

The following is a Linux/Mac BASH script that iterates through all subfolders starting with the current folder and tests each individual JAR file. The **PROJECT\_NAME\_HERE** in --**remote-repo-url** is important; it combines multiple scan results under a single Snyk project in the UI.

`find . -type f -name '*.jar' | uniq | xargs -I {} snyk monitor --file={} --scan-unmanaged --remote-repo-url=PROJECT_NAME_HERE --project-name={}`

The following is a Windows batch script, run from a **scanjar.bat** file.

```
REM Usage:    
REM For example: scanjar.bat "C:\workspace\app" "myapp" 
SET workspace=%1 
SET appname=%2 
for /R %workspace% %%f in (*.jar) do cmd /c snyk monitor --scan-unmanaged --remote-repo-url=%appname% --file=%%f
```

The following shows the end result in the Snyk UI for an app with the --**remote-repo-url=econnect**

![Result of scanning unmanaged JAR files](<../../.gitbook/assets/untitled (2).png>)
