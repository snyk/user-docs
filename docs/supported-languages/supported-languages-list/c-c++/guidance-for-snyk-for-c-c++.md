# Guidance for Snyk for C/C++

This page reviews considerations about languages and package managers, to help you apply Snyk effectively in your technology stack.

## Code analysis

* Snyk does not compile or require a build to perform analysis.
* Snyk Code analyzes source code directly.
* If you precompile components, make the source available during the scan.

## Open source and licensing

In the case of package managers like npm or Maven, Snyk traditionally uses the managed open-source capabilities of `snyk test` and `snyk monitor`. In the case of C/C++, Snyk supports unmanaged dependencies by adding `--unmanaged`.

{% hint style="info" %}
Snyk does not hook into a build nor rely on a build to perform scanning. Snyk performs analysis from source code.
{% endhint %}

* Open Source source code must be present.
* Snyk fingerprints files and compares them to the Snyk database to identify packages, versions, licenses, and vulnerabilities.

## Snyk integrations and common usage patterns

### IDE

#### With Snyk Code

No additional options are required. The Snyk plugin has views within the IDE for displaying results.

**With Snyk Open Source**

Under **Additional Parameters** in the IDE settings, enter the `--unmanaged` option to scan for C/C++ open source dependencies.

<div align="left"><figure><img src="https://lh6.googleusercontent.com/1j-2sJjuVejBJ6nARpaAx2uhdhqT7G3XyNCGZqFxBXJV9ujqRHBYiwInr_mFT7SH-fnhG6iUysKxzYKluPG1f3xUKyb2q-JycA_0QevtaS3hdm4I7-QT7M5benqzWkIe5N-7L3czV-F84_xUR5yl7k0" alt="Scan for dependencies"><figcaption><p>Scan for dependencies</p></figcaption></figure></div>

### CLI tips and tricks

#### Codebase

Snyk does not rely on a build to perform analysis. Only the source code is required.

Open the directory of the source code in the terminal and run the following command:

```
snyk code test
```

{% hint style="info" %}
If you precompile components, the source code should still be present to get the best results and coverage.
{% endhint %}

For reporting, you can generate reports using the [snyk-to-html](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) plugin to generate reporting artifacts. Additionally, there are JSON and SARIF export capabilities for programmatic access to results, using the options `--json` and `--sarif`, respectively. For more information, see [Exporting the test results to a JSON or SARIF file](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results).

#### **Open Source libraries**

For C/C++ open source, use the `--unmanaged` option to analyze license compliance issues and known security issues associated with open source.

```
snyk test --unmanaged
```

See [Snyk for C/C++](./) for details.

* To test, ensure the open source source code is present; it may be placed in a vendor folder.
* If you precompile open source, the open source code must still be present. For Snyk to make an accurate comparison with its existing knowledge base, the open source code must remain present.

Similarly, for monitoring and sharing reporting, use the following command:

```
snyk monitor --unmanaged --org=<org-id>
```

where `org-id` is found under your Organization settings in the Snyk web interface, Although the Organization ID is not required, it is strongly recommended that you use it. As with Snyk Code, you can generate reports using the [snyk-to-html](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) plugin to generate reporting artifacts.

* For individual or personal scans, use the CLI or IDE, and use the `snyk monitor --unmanaged` command to upload results.
  * However, Snyk recommends that you send these results to your personal folder and disable the scheduled scanning in the Project settings to ensure an individual scan does not cause noise.
  * This provides license and policy information in a viewable state.
* For automated scans, such as CI/CD, use `snyk monitor --unmanaged` and send results to the Organization of your choice. This provides license and policy information in a viewable state.

#### **Dependency lists**

Use `--print-deps` when performing open source scans to obtain a detailed list of discovered dependencies in your codebase and their origin source.

In C/C++, this has the additional benefit of identifying the confidence level of a given match. If there is a significant drop (< 90% confidence), it is likely the file has been modified and may not be the original source. Consider investigating if that is the case.

```
snyk test --unmanaged --print-deps
```

The list is printed before the issues list, as shown in the following image:

<figure><img src="https://lh5.googleusercontent.com/x4y1uIQ2fCFX956f1eP4664i6VKEgK6eOOddlAZ4p4WnQWJu1t_ugSOpL394KEnuzSIPRs08gNAsmjvPa-GAV0C-975esRdy0EPDY7WImG1-SXSOFO0TIAVfh_Jp2DLYc6bm7iZu55UbE3Boh4TNk_I" alt="View dependency lists"><figcaption><p>View dependency lists</p></figcaption></figure>

#### **License policy**

This feature allows a company to create a license policy for Open Source applications, indicating what licenses are not approved for use. When Snyk detects a match for a license that is not approved, Snyk sends an alert. This alert contains the name of the license and license policy text.

License policy text is associated with the license issue by your administrators. This text provides custom direction on what to do or why the license issue is contrary to the policy.

The following shows a license policy text example at the bottom of the screen, giving you directions on what to do if the license is found.

<div align="left"><figure><img src="https://lh4.googleusercontent.com/lIn5JFEyaZaTNMVenBoeGIgTpC6YHxpmAjK947z5ISPlHV1rlOvPNCLyzXxsGNj65AAlGn6ff9dF4lHVsVFYMaKXWC939tasD91k98xcDv_Ske6Dz7goMXl5lByyqg6ptvvqaK0UEqLSdzUU9GKrW4U" alt="License policy text example"><figcaption><p>License policy text example</p></figcaption></figure></div>

#### **Alternate testing options**

If you develop advanced dependency management strategies, you might not use the standard and frequently used package managers. For that reason, Snyk has provided test APIs. In the case of C++, if you know the open-source packages and versions that are included in the application but do not have the source code, you can use the endpoint [List issues for a package](../../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues) to do the analysis.

### **Options and plugins**

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../developer-tools/snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).
