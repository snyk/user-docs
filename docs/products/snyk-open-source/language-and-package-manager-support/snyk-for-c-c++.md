# Snyk for C / C++

{% hint style="info" %}
This feature is currently in Beta; contact Snyk for more details.
{% endhint %}

You can use Snyk to scan C / C++ projects.

### Features

{% hint style="info" %}
**NOTE**\
Features might not be available, depending on your subscription plan.
{% endhint %}

| Package managers / Features | CLI support | Git support | License scanning | Fixing | Runtime monitoring |
| --------------------------- | ----------- | ----------- | ---------------- | ------ | ------------------ |
| C/C++                       | ✔︎          |             |                  |        |                    |

### How it works

Scans are powered by an open source database, periodically updated with the latest source code from different online sources.

{% hint style="info" %}
Currently, we use and link to the US [National Vulnerability Database](https://nvd.nist.gov) (NVD). In future, we plan to also integrate the [Snyk Vuln DB](https://snyk.io/vuln).
{% endhint %}

When you run the `snyk unmanaged test` command, Snyk:

1. Converts all files down from your current folder into a list of hashes.
2. Sends the hashes to Snyk scan server.
3. Queries the database to find a list of potentially matching dependencies.
4. Links the dependencies to the known vulnerabilities.
5. Displays the results.

{% hint style="info" %}
To scan the project, the dependencies must be available as source code in the scanned directory. If the dependencies are in a different location, that location must be scanned.
{% endhint %}

#### Scanning archives

Snyk CLI extracts archives to analyze the source code inside.&#x20;

The extraction is not recursive. To enable recursive extraction, specify the depth of the extraction using the `--max-depth` parameter. To disable the extraction, set the the `--max-depth` to 0.

The supported archives formats are:

* zip-like archives
* tar
* tar with gzip compression algorithm

### Constraints and limitations

{% hint style="info" %}
The following constraints and limitations are by design. While we may work on improvements in the future, they are not considered an issue. Issues that are planned to be addressed are in the [Known Issues ](snyk-for-c-c++.md#known-issues)section.
{% endhint %}

**Dependencies source code needs to be available**

For Snyk CLI to be able to find any dependencies in your source code, the full source code of the dependencies needs to be present in the scanned folder. The following is a typical directory structure Snyk can scan (abbreviated):

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

Having a large percentage of files in their original (unchanged) form is critical to accurately identify the dependencies and so report the correct set of vulnerabilities. If you modify many of the files (or, for example, include only header files), this reduces the confidence of the scanning engine, leading to either dependencies not being identified, or being identified incorrectly (as a different version, or even a different package).

#### Data collection note

When you scan C++ projects, the following data is collected and may be stored for troubleshooting purposes:

| Category                        | Description                                                                                                                                                                                           |
| ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Hashes of the scanned files     | All files are converted to a list of irreversible hashes.                                                                                                                                             |
| Relative paths to scanned files | <p>The paths to files relative to the directory being scanned are included for better identification and matching.<br><br>Example: <br><code>./project-name/vendor/bzip2-1.0.6/blocksort.c</code></p> |

### Snyk CLI for C / C++ projects

#### Install the Snyk CLI

C/C++ scanning is available in [Snyk CLI](../../../features/snyk-cli/). See [Install the CLI](../../../features/snyk-cli/install-the-snyk-cli/) for details.

{% hint style="info" %}
The minimum version of Snyk CLI with C/C++ scanning is 1.713.
{% endhint %}

#### Run the test

To test your project for vulnerabilities, run:

```
$ snyk unmanaged test
```

{% hint style="warning" %}
If you scan a Linux project on Windows, make sure the repository is cloned with Linux line endings. See the [Known Issues](snyk-for-c-c++.md#known-issues) section for more details.
{% endhint %}

#### Displaying dependencies

To display dependencies, use the `--print-deps` command:

```bash
$ snyk unmanaged test --print-deps

Dependencies:

  cpython|https://github.com/python/cpython/archive/v3.7.2.zip@3.7.2
  confidence: 1.000
  
  zip|http://ftp.debian.org/debian/pool/main/z/zip/zip_3.0.orig.tar.gz@3.0
  confidence: 0.993
```

To learn what files contributed to each dependency being identified, use the `--print-dep-paths` argument:

```bash
$ snyk unmanaged test --print-dep-paths

Dependencies:

  curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0
  confidence: 1.000
  matching files:
    - c-example/deps/curl-7.58.0/CHANGES
    - c-example/deps/curl-7.58.0/CMake/CMakeConfigurableFile.in
    - c-example/deps/curl-7.58.0/CMake/CurlSymbolHiding.cmake
    ... and 2857 more files
```

#### Understanding the confidence level

You may need to change the source code of the dependencies that you use in your software. As Snyk uses file signatures to find the closest possible match to an open source library, your changes may decrease the accuracy of the identification of the actual library.

To learn how confident Snyk is about the identified dependency and its version, use the `--print-deps` or `--print-dep-paths` command line argument:

```
curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0
confidence: 0.993
```

This confidence level shows how confident Snyk is about the actual identification of the dependency. The number can be between 0 and 1 and the higher it is, the more accurate the identification is. So a confidence of **1** means that all the files in the source tree fully matched all the expected files in our database.

#### JSON output

To get a machine-readable output in JSON, use the `--json` argument:

```
$ snyk unmanaged test --json
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
      "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0": {
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

### Command line options

The following `snyk`command line options are supported with `snyk unmanaged`:

#### ORG\_NAME

`--org=ORG_NAME`

Specify the ORG\_NAME to run Snyk commands tied to a specific organization. This defines where new projects are created after running the **monitor** command, some features have availability and private tests limits. If you have multiple organizations, you can set a default from the CLI using:

```
snyk config set org=ORG_NAME
```

Setting a default ensures all newly monitored projects are created under your default organization. To override the default, use the **--org=ORG\_NAME** argument.

Default: uses the ORG\_NAME set as default in your Account settings.

#### json

`--json`

Prints results in JSON format.

#### OUTPUT\_FILE\_PATH

`--json-file-output=OUTPUT_FILE_PATH`

(only in test command) Save test output in JSON format directly to the specified file, regardless of whether or not you use the **--json** option.

This is useful to display the human-readable test output via **stdout** and at the same time save the JSON format output to a file.

**target-dir**

`--target-dir <directory>`

Scan the path specified in the argument instead of the current directory.

{% hint style="info" %}
Alternatively, you can run just `snyk unmanaged test <directory>`
{% endhint %}

**max-depth**

`--max-depth=1`

Specify the maximum level of archive extraction. Use `0` to disable archive extraction completely.

### Import scan results in Snyk App

To import the test results (issues and dependencies) in Snyk App, run the `snyk unmanaged monitor` command:

```
$ snyk unmanaged monitor
Monitoring /c-example (c-example)...

Explore this snapshot at https://app.snyk.io/org/example-org/project/8ac0e233-d0f9-403e-b422-5970e7a37443/history/5de4616d-3967-485f-bf21-bbbe91068029

Notifications about newly disclosed issues related to these dependencies will be emailed to you.
```

This creates a snapshot of dependencies and vulnerabilities, and imports them in Snyk App, where you can review the issues and see them included in your reports.

Importing a project with unmanaged dependencies creates a new project in Snyk App:

![Project with unmanaged dependencies](../../../.gitbook/assets/kuva.png)

{% hint style="info" %}
Automated regular testing and re-scanning from the Snyk App is not currently supported. To run a new scan and import its updated results, manually run the `snyk unmanaged monitor` command again.
{% endhint %}

### Known issues

**Scanning on Windows**

Many open source projects in git use Unix line endings. By default, git on Windows converts Unix line endings to Windows line endings and only converts them back for the actual commits. Our database contains source code signatures with the original line endings (as defined in the individual projects), so when you scan on Windows, the signatures generated for the files with Windows line endings are different than the signatures in our database. In such case, it is very likely no dependencies will be found.

To scan a project with Unix line endings on Windows, disable git line endings conversion. To configure this globally, run:

```shell
git config --global core.autocrlf false
```

### Frequently asked questions

#### **Is my source code sent to Snyk servers?**

No. The files are converted to a list of hashes before they are sent for scanning.

#### **Why did Snyk not find any dependencies?**

We store the official releases of many of open source components in our database but it is possible that the source code you scanned is not there or is just simply not found. Let us know and we can help you find out what happened and potentially improve our scanning algorithms.

Here are a few things that you can check on your own:

* The source code of the dependencies you scanned is actually available as source code (unpacked) in the folder that you scanned. If you use a package manager, such as Conan, the source code is likely to be in the Conan cache, along with the source code of other dependencies of your other projects. To scan dependencies managed by a package manager, we recommend that you do that in a clean environment (for example during a build).
* The source code of the dependencies is not from an official release of the OSS component, and we do not have it in the database
* The source code of the OSS has been modified too much, so Snyk cannot detect it. If there are too few files and you modify most of them, Snyk cannot match them to a component from our database.\
  Examples of common modifications are whitespace formatting, adding license or copyright headers.
* You are on Windows, and git converted line endings to Windows line endings. Currently we can recognize files that have retained their original line endings.
* The source code of the OSS components is too new. Our database is refreshed monthly but it takes time for the latest releases to get processed.
