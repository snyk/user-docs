# Snyk Security Scan task parameters and values

The following describes the Snyk task configuration fields on the configuration panel in Azure Pipelines, the associated parameters for Azure Pipelines integration, and the valid values.

## **Field: Snyk API token service**

**Parameter:** serviceConnectionEndpoint\
**Required:** Yes\
**Default:** none\
**Type:** String / Azure Service Connection Endpoint of type SnykAuth / Snyk Authentication

**Description:** The Azure DevOps service connection endpoint where your Snyk API token is defined. Your admin defines this within your Azure DevOps project settings, assigning it using a unique string in order to differentiate between different connections.

The configuration panel displays all available Snyk service connections from a dropdown list like the following:

<figure><img src="broken-reference" alt="Snyk service connections"><figcaption><p>Snyk service connections</p></figcaption></figure>

If multiple Snyk service connections are available from the dropdown list, ask your administrator which to use for the pipeline you are working with.

## **Field: What do you want to test**

**Parameter:** testType\
**Required:** Yes\
**Default:** application\
**Type:** string: `app` or `container` or `code`

**Description:** Determines the type of testing to perform, which also selects the dynamic fields to display as described on the rest of this page.

* Specifying `app` will perform analysis of your application's open-source dependencies (SCA)
* Specifying `container` will analyze container packages and detected application open-source packages (SCA)
* Specifying `code` will perform static analysis (SAST) on source code.

## **Field:** Container Image Name

**Parameter:** dockerImageName\
**Required:** Yes\
**Default:** none\
**Type:** String

**Description:** The name of the container image to test. This dynamic field appears when **What do you want to test** is set to **Container Image**. Set to **Yes** if you are performing a container image test.

## **Field:** Path to Dockerfile

**Parameter:** dockerfilePath\
**Required:** Yes\
**Default:** none\
**Type:** string

**Description:** The path to the Dockerfile corresponding to the `dockerImageName`. This dynamic field appears when **What do you want to test** is set to **Container Image**. Set to **Yes** if you are performing a container image test.

## **Field:** Custom path to manifest file to test

**Parameter:** targetFile\
**Required:** No\
**Default:** none\
**Type:** string

**Description:** Applicable to application-type tests only. The path to the manifest file to be used by Snyk. Should be provided only if non-standard. This dynamic field appears when **What do you want to test** is set to **Application**.

## **Field: Testing severity threshold**

**Parameter:** severityThreshold\
**Required:** No\
**Default:** "low"\
**Type:** string: "low" or "medium" or "high" or "critical"

**Description:** The severity threshold to use when testing. By default, issues of all severity types are found. If not configured, the default severity is set to **low**.

## **Field: Code testing severity threshold**

**Parameter:** codeSeverityThreshold\
**Required:** No\
**Default:** "low"\
**Type:** string: "low" or "medium" or "high"

**Description:** The Code severity threshold to use when testing. By default, issues of all severity types are found. If not configured, the default severity is set to `low`.

## **Field: When to run Snyk Monitor**

**Parameter:** monitorWhen\
**Required:** No\
**Default:** "always"\
**Type:** string: string: `always`, `onIssuesFound`, or `never`

**Description:** When to run `snyk monitor` to capture the dependency tree of the application or container image and monitor it within Snyk.

## **Field:** Fail build if Snyk finds issues

**Parameter:** failOnIssues\
**Required:** Yes\
**Default:** true\
**Type:** Boolean

**Description:** Specifies whether pipeline jobs should be failed or continued based on issues found by Snyk.

## **Field:** Project name in Snyk

**Parameter:** projectName\
**Required:** No\
**Default:** none\
**Type:** string

**Description:** A custom name for the Snyk Project to be created on snyk.io.

## **Field:** Organization ID in Snyk

**Parameter:** organization\
**Required:** No\
**Default:** none\
**Type:** string

**Description:** ID of the Snyk Organization under which this Project should be tested and monitored.

## **Field:** Test (Working) Directory

**Parameter:** testDirectory\
**Required:** No\
**Default:** none\
**Type:** string

**Description:** Alternate working directory. For example, if you want to test a manifest file in a directory other than the root of your repository, you would put in a relative path to that directory.

## **Field:** Additional command-line args for Snyk CLI (advanced)

**Parameter:** additionalArguments\
**Required:** No\
**Default:** none\
**Type:** string

**Description:** Additional Snyk CLI arguments to be passed in. See the [CLI commands and options summary](../../snyk-cli/cli-commands-and-options-summary.md) for details. Add `--all-projects` as good practice (for example, for .NET), if no Project has been found.

## **Field:** Trust unknown Certificate Authorities (advanced)

**Parameter:** ignoreUnknownCA\
**Required:** No\
**Default:** false\
**Type:** boolean

**Description:** Use to ignore unknown or self-signed certificates during certificate validation, so self-signed certificates are automatically trusted.
