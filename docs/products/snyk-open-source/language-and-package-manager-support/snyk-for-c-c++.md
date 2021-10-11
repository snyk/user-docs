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

|       | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| ----- | ----------- | ----------- | ---------------- | ----------- | ------------------ |
| C/C++ | ✔︎          |             |                  |             |                    |

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

#### Data collection note

When you scan C++ projects, the following data is collected and may be stored for troubleshooting purposes:

| Category                    | Description                                                                                 |
| --------------------------- | ------------------------------------------------------------------------------------------- |
| Hashes of the scanned files | All files are converted to a list of irreversible hashes.                                   |
| Full paths to scanned files | The paths to files on your local drive are included for better identification and matching. |

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

#### Displaying dependencies

To display dependencies, use the `--print-deps` command:

```bash
$ snyk unmanaged test --print-deps

Dependencies:

  cpython|https://github.com/python/cpython/archive/v3.7.2.zip@3.7.2
  zip|http://ftp.debian.org/debian/pool/main/z/zip/zip_3.0.orig.tar.gz@3.0
```

To learn what files contributed to each dependency being identified, use the `--print-dep-paths` argument:

```bash
$ snyk unmanaged test --print-dep-paths

Dependencies:

  curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0
  matching files:
    - c-example/deps/curl-7.58.0/CHANGES
    - c-example/deps/curl-7.58.0/CMake/CMakeConfigurableFile.in
    - c-example/deps/curl-7.58.0/CMake/CurlSymbolHiding.cmake
    ... and 2857 more files
```

#### JSON output

To get a machine-readable output in JSON, use the `--json` argument:

```
$ snyk unmanaged test --json
[{
	"issues": [{
		"pkgName": "curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz",
		"pkgVersion": "7.58.0",
		"issueId": "CVE-2019-5481",
		"fixInfo": {
			"isPatchable": false,
			"isPinnable": false
		}
	}],
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
			"references": [{
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
	"depsFilePaths": {
		"curl|https://github.com/curl/curl/releases/download/curl-7_58_0/curl-7.58.0.tar.xz@7.58.0": [
			"c-example/deps/curl-7.58.0/CHANGES",
			"c-example/deps/curl-7.58.0/CMake/CMakeConfigurableFile.in",
			"c-example/deps/curl-7.58.0/CMake/CurlSymbolHiding.cmake"
		]
	}
}]
```

### Command line options

The following `snyk`command line options are supported with `snyk unmanaged`:

#### ORG_NAME

`--org=ORG_NAME` 

Specify the ORG_NAME to run Snyk commands tied to a specific organization. This defines where new projects are created after running the **monitor** command, some features have availability and private tests limits. If you have multiple organizations, you can set a default from the CLI using:

```
snyk config set org=ORG_NAME
```

Setting a default ensures all newly monitored projects are created under your default organization. To override the default, use the **--org=ORG_NAME** argument.

Default: uses the ORG_NAME set as default in your Account settings.

#### json

`--json` 

Prints results in JSON format.

#### OUTPUT_FILE_PATH

`--json-file-output=OUTPUT_FILE_PATH`

(only in test command) Save test output in JSON format directly to the specified file, regardless of whether or not you use the **--json** option. 

This is  useful to display the human-readable test output via **stdout** and at the same time save the JSON format output to a file.

### Known issues

#### Files in hidden directories are ignored

When scanning a directory, all files that are in a hidden directory (such as **.conan** or **.git**) are ignored. Dependencies stored in such directories will not be detected.

### Frequently asked questions

#### **Is my source code sent to Snyk servers?**

No. The files are converted to a list of hashes before they are sent for scanning.

#### **Why Snyk did not find any dependencies?**

We store the official releases of many of open source components in our database but it is possible that the source code you scanned is not there or is just simply not found. Let us know and we can help you find out what happened and potentially improve our scanning algorithms.

Here are a few things that you can check on your own:

* The source code of the dependencies you scanned is actually available as source code (unpacked) in the folder that you scanned. If you use a package manager, such as Conan, the source code is likely to be in the Conan cache, along with the source code of other dependencies of your other projects. To scan dependencies managed by a package manager, we recommend that you do that in a clean environment (for example during a build).
* The source code of the dependencies is not from an official release of the OSS component, and we do not have it in the database
* The source code of the OSS has been modified too much, so Snyk cannot detect it. If there are too few files and you modify most of them, Snyk cannot match them to a component from our database. \
  Examples of common modifications are whitespace formatting, adding license or copyright headers.
* You are on Windows, and git converted line endings to Windows line endings. Currently we can recognize files that have retained their original line endings.
* The source code of the OSS components is too new. Our database is refreshed regularly but it takes time for the latest releases to get processed.

#### **What is coming next?**

Our plan is to show more information on how and why certain components were detected in our source code (show files that were detected to be a part of the component) and allow you to bring the information in the App (using the `snyk unmanaged monitor` command) so you can see the dependencies there.
