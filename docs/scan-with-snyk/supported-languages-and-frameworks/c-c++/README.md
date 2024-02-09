# C/C++

## Supported frameworks and package managers

### Code analysis

Snyk Code supports the following:

* Operating systems
  * Linux
  * Windows (limited)
* Embedded systems: Linux
* GUI Framework&#x20;
* Libraries: POSIX, C++ Standard Library, Boost, QT

### Open source and licensing

#### Open source policy

To manage licenses from your developer workflows through policy, see the following topics:

* [Defining a secure open source policy](https://snyk.io/series/open-source-security/open-source-policy/)
* [Use Snyk security policies to prioritize fixes more efficiently](https://snyk.io/blog/snyk-security-policies/)

<table><thead><tr><th width="267">Package managers / Features</th><th>CLI support</th><th>Git support</th><th>License scanning</th><th>Fix PRs</th></tr></thead><tbody><tr><td>C/C++</td><td>✔︎</td><td></td><td>✔︎</td><td></td></tr></tbody></table>

#### Open source license compliance

To check compliance for open source licenses, see [Getting Started with Snyk License Compliance Management](https://docs.snyk.io/scan-application-code/snyk-open-source/licenses/getting-started-snyk-licensing-compliance).

## Getting started with Snyk for C/C++ across environments

Scans are powered by an open source database, periodically updated with the latest source code from online sources.

{% hint style="info" %}
To navigate through the vulnerabilities for C/C++, use the [Snyk Vuln DB](https://security.snyk.io).
{% endhint %}

When you run the [`snyk test --unmanaged`](../../../snyk-cli/commands/test.md#unmanaged) command, Snyk does the following:

1. Converts all files from your current folder into a list of hashes.
2. Sends hashes to the Snyk scan server to compute the dependencies list.
3. Queries the database to find a list of potentially matching dependencies.
4. Links the dependencies to the known vulnerabilities.
5. Displays the results.

{% hint style="info" %}
For Snyk to scan the Project, the dependencies must be available as source code in the scanned directory. If the dependencies are in a different location, that location must be scanned.
{% endhint %}

### Scanning archives

By default, archives are not scanned. However, Snyk CLI can recursively extract archives to analyze the source code inside.

To enable archive extraction, specify the depth of the extraction using the `--max-depth` option.

The supported archive formats are:

* zip-like archives
* tar archives
* tar with gzip compression algorithm

### Constraints and limitations

{% hint style="info" %}
The following constraints and limitations are by design. While Snyk may work on improvements in the future, they are not considered an issue.&#x20;
{% endhint %}

### **Source code dependencies need to be available in the scanned folder**

For Snyk CLI to be able to find dependencies in your source code, enough of the full dependencies source code needs to be present in the scanned folder.

Having a large percentage of files in their original (unchanged) form is critical to accurately identifying dependencies and reporting the correct set of vulnerabilities back. Modifying that source code reduces the confidence of the scanning engine, resulting in less accurate results. Other potential issues could include dependencies not being identified or being identified incorrectly, as a different version or even a different package.

The example that follows shows a typical package with dependencies listed:

```
c-example
├── deps
│   ├── curl-7.58.0
│   │   ├── include
│   │   │   ├── Makefile.am
│   │   │   ├── Makefile.in
│   │   │   ├── README
│   │   │   └── curl
│   │   ├── install-sh
│   │   ├── lib
│   │   │   ├── asyn.h
│   │   │   ├── base64.c
│   │   │   ├── checksrc.pl
│   │   │   ├── config-amigaos.h
│   │   │   ├── conncache.c
│   │   │   ├── conncache.h
│   │   ├── src
│   │   │   ├── tool_binmode.c
│   │   │   ├── tool_binmode.h
│   │   │   ├── tool_bname.c
│   │   │   ├── tool_xattr.c
...
```

### Data collection note

When you scan C++ Projects, the following data is collected and may be stored for troubleshooting purposes:

**Hashes of the scanned files:** All files are converted to a list of irreversible hashes.

**Relative paths to scanned files:** The paths to files relative to the directory being scanned are included for better identification and matching.\
\
Example:\
`./project-name/vendor/bzip2-1.0.6/blocksort.c`

### Snyk CLI&#x20;

#### Prerequisites

1. [Create a Snyk account](../../../getting-started/quickstart/create-or-log-in-to-a-snyk-account.md)
2. [Install Snyk CLI and authenticate your machine](../../../snyk-cli/getting-started-with-the-snyk-cli.md#install-the-snyk-cli-and-authenticate-your-machine)
3. [Set the default Organization for all Snyk tests](broken-reference) (code analysis)

#### Code analysis

To start testing your code using Snyk Code open your repository in a terminal and run the following  command:

```javascript
snyk code test
```

To customize test options, run other commands, exclude directories and files, and explore the results in different formats, see the following:

* [Snyk CLI commands](../../../snyk-cli/commands/#available-commands)
* [Exclude directories and files from the Snyk tests](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
* [Explore test results in a JSON or SARIF format in the terminal ](broken-reference)
* [Exporting the test results to a JSON or SARIF file](broken-reference)

#### Open source and licensing

**Run the test**

To test your Project for vulnerabilities, run the following:

```
$ snyk test --unmanaged
```

**Displaying dependencies**

To display dependencies, use the `--print-deps` option:

```bash
$ snyk test --unmanaged --print-deps

Testing /Users/user/src/foo...


Dependencies:

  https://curl.se|curl@7.29.0
  purl: pkg:generic/curl@7.29.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Farcheology%2Fcurl-7.29.0.tar.gz
  confidence: 1.000

  https://github.com|nih-at/libzip@1.8.0
  purl: pkg:generic/libzip@1.8.0?download_url=https:%2F%2Fgithub.com%2Fnih-at%2Flibzip%2Farchive%2Fv1.8.0.tar.gz
  confidence: 1.000

  https://github.com|madler/zlib@1.2.11
  purl: pkg:generic/zlib@1.2.11?download_url=https:%2F%2Fzlib.net%2Ffossils%2Fzlib-1.2.11.tar.gz
  confidence: 1.000
```

To learn what files contributed to each dependency being identified, use the `--print-dep-paths` option:

```bash
$ snyk test --unmanaged --print-dep-paths

Testing /Users/user/src/foo...


Dependencies:

  https://curl.se|curl@7.29.0
  purl: pkg:generic/curl@7.29.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Farcheology%2Fcurl-7.29.0.tar.gz
  confidence: 1.000
  matching files:
    - curl-7.29.0/Android.mk
    - curl-7.29.0/CHANGES
    - curl-7.29.0/CMake/CMakeConfigurableFile.in
    ... and 1766 more files

  https://github.com|nih-at/libzip@1.8.0
  purl: pkg:generic/libzip@1.8.0?download_url=https:%2F%2Fgithub.com%2Fnih-at%2Flibzip%2Farchive%2Fv1.8.0.tar.gz
  confidence: 1.000
  matching files:
    - libzip-1.8.0/API-CHANGES.md
    - libzip-1.8.0/AUTHORS
    - libzip-1.8.0/CMakeLists.txt
    ... and 780 more files

  https://github.com|madler/zlib@1.2.11
  purl: pkg:generic/zlib@1.2.11?download_url=https:%2F%2Fzlib.net%2Ffossils%2Fzlib-1.2.11.tar.gz
  confidence: 1.000
  matching files:
    - zlib-1.2.11/CMakeLists.txt
    - zlib-1.2.11/ChangeLog
    - zlib-1.2.11/FAQ
    ... and 249 more files
```

**Understanding the confidence level**

You may need to change the source code of the dependencies that you use in your software. As Snyk uses file signatures to find the closest possible match to an open-source library, your changes may decrease the accuracy of the identification of the actual library.

To learn how confident Snyk is about the identified dependency and its version, use the `--print-deps` or `--print-dep-paths` command line option:

```
curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0
confidence: 0.993
```

This confidence level shows how confident Snyk is about the actual identification of the dependency. The number can be between **0** and **1** and the higher it is, the more accurate the identification is. Thus a confidence of **1** means that all the files in the source tree fully matched all the expected files in the Snyk database.

**JSON output**

To get a machine-readable output in JSON, use the `--json` option:

```
$ snyk test --unmanaged --json
[
  {
    "issues": [
      {
        "pkgName": "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz",
        "pkgVersion": "7.58.0",
        "issueId": "CVE-2019-5481",
        "fixInfo": {
          "isPatchable": false,
          "isPinnable": false
        }
      }
    ],
    "issuesData": {
      "CVE-2019-5481": {
        "severity": "high",
        "CVSSv3": "",
        "originalSeverity": "high",
        "severityWithCritical": "high",
        "type": "vuln",
        "alternativeIds": [
          ""
        ],
        "creationTime": "2019-09-16T19:15:00.000Z",
        "disclosureTime": "2019-09-16T19:15:00.000Z",
        "modificationTime": "2020-10-20T22:15:00.000Z",
        "publicationTime": "2019-09-16T19:15:00.000Z",
        "credit": [
          ""
        ],
        "id": "CVE-2019-5481",
        "packageManager": "cpp",
        "packageName": "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz",
        "language": "cpp",
        "fixedIn": [
          ""
        ],
        "patches": [],
        "exploit": "No Data",
        "functions": [
          ""
        ],
        "semver": {
          "vulnerable": [
            "7.58.0"
          ],
          "vulnerableHashes": [
            ""
          ],
          "vulnerableByDistro": {}
        },
        "references": [
          {
            "title": "https://curl.haxx.se/docs/CVE-2019-5481.html",
            "url": "https://curl.haxx.se/docs/CVE-2019-5481.html"
          },
        ],
        "internal": {},
        "identifiers": {
          "CVE": [
            "CVE-2019-5481"
          ],
          "CWE": [],
          "ALTERNATIVE": [
            ""
          ]
        },
        "title": "CVE-2019-5481",
        "description": "",
        "license": "",
        "proprietary": true,
        "nearestFixedInVersion": ""
      }
    },
    "fileSignaturesDetails": {
      "https://curl.se|curl@7.58.0": {
        "artifact": "curl",
        "version": "7.58.0",
        "author": "curl",
        "path": "curl-7.58.0",
        "id": "59d80da8ba341aaff828662700000000",
        "url": "https://curl.se/download/curl-7.58.0.tar.gz",
        "purl": "pkg:generic/curl@7.58.0?download_url=https:%2F%2Fcurl.se%2Fdownload%2Fcurl-7.58.0.tar.gz",
        "score": 1,
        "confidence": 1,
        "filePaths": [
          "deps/curl-7.58.0/CHANGES",
          "c-example/deps/curl-7.58.0/CMake/CMakeConfigurableFile.in",
          "c-example/deps/curl-7.58.0/CMake/CurlSymbolHiding.cmake"
        ],
        "confidence": 1
      }
    }
  }
]
```

**Command line options**

The following `snyk` command line options are supported with the `snyk test --unmanaged` and `snyk monitor --unmanaged` commands:

`--org=<ORG_ID>`\
`--json`\
`--json-file-output=<OUTPUT_FILE_PATH>` (`snyk test` only)\
`--remote-repo-url=<URL>`\
`--severity-threshold=<low|medium|high|critical>` (`snyk test` only)\
`--max-depth`\
`--print-dep-paths`\
`--target-reference=<TARGET_REFERENCE>` (`snyk monitor` only)\
`--project-name=<c-project>` (`snyk monitor` only)

For more information about command line options, see the Snyk help docs: [Options for scanning with `snyk test --unmanaged`](https://docs.snyk.io/snyk-cli/commands/test#options-for-scanning-using-unmanaged) or [`snyk monitor --unmanaged`](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-scanning-using-unmanaged).

To import the test results (issues and dependencies) in the Snyk CLI, run the `snyk monitor --unmanaged` command:

```
$ snyk monitor --unmanaged
Monitoring /c-example (c-example)...

Explore this snapshot at https://app.snyk.io/org/example-org/project/8ac0e233-d0f9-403e-b422-5970e7a37443/history/5de4616d-3967-485f-bf21-bbbe91068029

Notifications about newly disclosed issues related to these dependencies will be emailed to you.
```

This creates a snapshot of dependencies and vulnerabilities and imports them into the Snyk Web UI, where you can review the issues and see them included in your reports.

Importing a Project with unmanaged dependencies creates a new Project:

<figure><img src="../../../.gitbook/assets/Screenshot 2023-08-04 at 11.18.10.png" alt=""><figcaption><p>Project with unmanaged dependencies</p></figcaption></figure>

{% hint style="info" %}
Snyk Web UI supports only code analysis, using Snyk Code.
{% endhint %}

### Snyk integrations

#### Snyk Code

No additional options are required. The Snyk plugin has views within the IDE for displaying results.

**Snyk Open Source**&#x20;

Under **Additional Parameters** in the IDE settings, enter the **--unmanaged** option to scan for C/C++ open source dependencies.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-08-04 at 12.18.01.png" alt=""><figcaption><p>Scan for dependencies</p></figcaption></figure>

## Help

:link: See [best practices](best-practices-for-c-c++.md) and [troubleshooting](troubleshooting-c-c++.md).
