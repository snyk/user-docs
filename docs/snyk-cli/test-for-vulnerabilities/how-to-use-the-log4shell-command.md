# Log4shell command use

## Introduction

`snyk log4shell` is a Snyk CLI command, that helps find traces of the **log4j** library that are affected by the **Log4Shell** vulnerability ([CVE-2021-44228](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44228)), even if this library is not declared in the manifest files (such as pom.xml or build.gradle).

This command tests your built project and third-party applications, and it is complementary to the `snyk test` and `snyk test --scan-all-unmanaged` commands.

{% hint style="info" %}
Read more about the Log4Shell vulnerability in the Snyk [VulnDB entry](https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720).
{% endhint %}

To test Java projects using their package manager manifest files, see [Snyk for Java (Gradle, Maven](../../products/snyk-open-source/language-and-package-manager-support/snyk-for-java-gradle-maven.md). To learn more about the `snyk test --scan-all-unmanaged`, see the [Maven options section of the CLI reference](https://docs.snyk.io/snyk-cli/cli-reference#options-for-maven-projects).

## Usage

Use `snyk log4shell` to scan a Java project, to see if the project includes:

* Any `.jar` or `.war` files with a vulnerable version of log4j.
* Any files that are known to be present in a vulnerable version of the log4j library. Such a finding indicates that it is possible the whole log4j is included.

## How to run

1. Install the latest version of the Snyk CLI - see [Install the Snyk CLI](../install-the-snyk-cli.md).
2. Make sure you have built the project.
3. Navigate to the project directory to scan.
4. Enter `snyk log4shell`.\
   **Note**: this command does not require (nor support) any additional arguments.

## Scan results

Results appear after the scan finishes.

For example:

```bash
$ snyk log4shell
Please note this command is for already built artifacts. To test source code use `snyk test`.

Results:
A vulnerable version of log4j was detected: 
         demo-0.0.1-SNAPSHOT/WEB-INF/lib/log4j-core-2.14.1.jar
         demo-0.0.1-SNAPSHOT.war/WEB-INF/lib/log4j-core-2.14.1.jar
         demo-0.0.1-SNAPSHOT.war.original/WEB-INF/lib/log4j-core-2.14.1.jar

 We highly recommend fixing this vulnerability. If it cannot be fixed by upgrading, see mitigation information here:
        - https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720
        - https://snyk.io/blog/log4shell-remediation-cheat-sheet/
```

If no traces of a vulnerable log4j library are found, the results show this:

```
$ snyk log4shell
Please note this command is for already built artifacts. To test source code use `snyk test`.

Results:
No known vulnerable version of log4j was detected
```

## Fix advice

For more details about fixing the affected packages, see the Snyk [Log4Shell fix cheatsheet](https://snyk.io/blog/log4shell-remediation-cheat-sheet).

## Limitations

* The Snyk CLI compares file signatures to a database of known files. If the Log4Shell vulnerability is fixed in a different way from updating the log4j library, the library is still reported in the results.
* If the source code of the log4j library has been modified, it is detected.
