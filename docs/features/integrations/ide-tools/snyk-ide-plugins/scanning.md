# Scanning with IDE plugins

## Manifest files <a href="#6f65ebbb-6b2b-47aa-99b9-93cac28849a8" id="6f65ebbb-6b2b-47aa-99b9-93cac28849a8"></a>

Snyk CLI supports a wide range of different programming languages and package managers for open source dependencies. For a complete list of supported manifest files see [Open source language and package manager support](../../../../products/snyk-open-source/language-and-package-manager-support/).

{% hint style="info" %}
The CLI defaults to scanning the first supported manifest file it detects. This means that when a project has multiple manifest files, you must provide the name of the manifest file to scan using the `--file` option.
{% endhint %}

Note that the Snyk CLI has multiple commands, but for the purpose of scanning you use the `test` command with the `--json` option to have the output converted to a machine readable format you are able to parse: `snyk test --file=<manifest file> --json`

The test command expects the dependencies to be installed already (that is, scan after `mvn install` or `npm install` have been executed) and for package lock files to be present if relevant. If you try to run the scan before the dependencies are installed you can expect to see an error like this:

```
~/git/goof $ snyk test
Missing node_modules folder: we can't test without dependencies.
Please run 'npm install' first.
```

For successful runs, you will see a long JSON file similar to the one that follows. To make sense of the data that is returned see the [Data mapping](data-mapping.md) page.

```javascript
{
  "vulnerabilities": [
    {
      "CVSSv3": "CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
      "alternativeIds": [],
      "creationTime": "2019-10-22T12:22:54.665794Z",
      "credit": [
        "Roman Burunkov"
      ],
      "cvssScore": 9.8,
      "disclosureTime": "2019-10-18T11:17:09Z",
      "exploit": "Not Defined",
      "fixedIn": [
        "1.1.6-alpha.6"
      ],
      "functions": [],
      "functions_new": [],
      "id": "SNYK-JS-EXPRESSFILEUPLOAD-473997",
      "identifiers": {
        "CVE": [],
        "CWE": [
          "CWE-79"
        ],
        "NSP": [
          1216
        ]
      },
      "language": "js",
      "modificationTime": "2019-11-20T09:48:38.528931Z",
      "moduleName": "express-fileupload",
      "packageManager": "npm",
      "packageName": "express-fileupload",
      "patches": [],
      "publicationTime": "2019-10-22T15:08:40Z",
      "references": [
        {
          "title": "GitHub PR",
          "url": "https://github.com/richardgirges/express-fileupload/pull/171"
        }
      ],
```

{% hint style="info" %}
All paying Snyk customers have the license scanning feature included in their plan. The `snyk test` command returns both vulnerability and open source license violations in the test results. To distinguish a license violation look for: `"type": "license"` inside the record fields.
{% endhint %}

## When to rerun scans <a href="#607b2cd8-2fb5-49ee-8473-319a42b8c421" id="607b2cd8-2fb5-49ee-8473-319a42b8c421"></a>

To make sure you are presenting up-to-date scan results to users, you must rerun the snyk scan any time one of the following happens:

* Any of the manifest files have been edited.
* A day has passed since the last scan. This is needed as the vulnerability database powering the Snyk CLI constantly collects new vulnerabilities which are relevant for the user to see even if they have not edited the manifest file.
* Any time the user explicitly reinstalls the dependencies (for example, calls `mvn install`) as semvers have flexibility and a reinstall may fetch different dependency versions than were previously installed even without a manifest file change.
