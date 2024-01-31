# Best practices for C/C++

Use this guide to apply Snyk effectively in your technology stack.

## Language and package manager considerations

### Code analysis

* Snyk does not compile or require a build to perform analysis.
* Snyk Code analyzes source code directly.
* If you precompile components, make the source available during the scan.

### Open source and licensing

In the case of package managers like npm or maven, it traditionally uses the managed open source capabilities of `snyk test` and `snyk monitor`. In the case of C/C++, Snyk supports unmanaged dependencies by adding `--unmanaged`.

{% hint style="info" %}
Snyk does not hook into a build nor rely on a build to perform scanning. Snyk performs analysis from source code.
{% endhint %}

* Open Source source code must be present.
* Snyk fingerprints files and compares them to the Snyk database to identify packages, versions, licenses, and vulnerabilities.

## Snyk Integrations and common usage patterns

### IDE

#### With Snyk Code

No additional options are required. The Snyk plugin has views within the IDE for displaying results.

**With Snyk Open Source**&#x20;

Under **Additional Parameters** in the IDE settings, enter the `--unmanaged` option to scan for C/C++ open source dependencies.

<div align="left">

<figure><img src="https://lh6.googleusercontent.com/1j-2sJjuVejBJ6nARpaAx2uhdhqT7G3XyNCGZqFxBXJV9ujqRHBYiwInr_mFT7SH-fnhG6iUysKxzYKluPG1f3xUKyb2q-JycA_0QevtaS3hdm4I7-QT7M5benqzWkIe5N-7L3czV-F84_xUR5yl7k0" alt="Scan for dependencies"><figcaption><p>Scan for dependencies</p></figcaption></figure>

</div>

### CLI Tips and tricks

:link: [CLI cheat sheet](https://snyk.io/blog/snyk-cli-cheat-sheet/)

#### What to test

Use the `--help` option in the CLI for details of Snyk CLI commands.

:link: [CLI commands and options summary](../../../snyk-cli/cli-commands-and-options-summary.md)

#### Codebase

Snyk does not rely on a build to perform analysis. Only the source code is required.

Open the directory of the source code in the terminal and run the following command:

```
snyk code test
```

{% hint style="info" %}
If you precompile components, the source code should still be present to get the best results and coverage.
{% endhint %}

For reporting, you can generate reports using the [snyk-to-html](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) plugin to generate reporting artifacts. Additionally, there are JSON and SARIF export capabilities for programmatic access to results, using **--json** and **--sarif**, respectively. See [Exporting the test results to a JSON or SARIF file](broken-reference).

#### **Open Source libraries**

For C/C++ open source, use the **--unmanaged** option to analyze license compliance issues and known security issues associated with open source. See [Snyk for C/C++](../../../scan-applications/supported-languages-and-frameworks/c-c++.md) for details.

* To test, make sure the open source source code is present, and it may be placed in a vendor folder.&#x20;
* If you precompile open source, the open source code must still be present. For Snyk to make an accurate comparison with its existing knowledge base, the open source code must remain present.

```
snyk test --unmanaged
```

Similarly, for monitoring and sharing reporting:

```
snyk monitor --unmanaged --org=<org-id>
```

Where **org-id** is found under your Organization settings in the Snyk web interface, although the Organization id is not required, it's strongly suggested. Like Snyk Code, you can generate reports using the [snyk-to-html](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) plugin to generate reporting artifacts.&#x20;

* For individual or personal scans, use the CLI or IDE and use the **snyk monitor --unmanaged** command to upload results, but the recommendation is you send these results to your personal folder and disable the scheduled scanning in the Project settings to ensure an individual scan does not cause noise. This provides license/policy information in a viewable state.
* For automated scans, such as CI/CD, use **snyk monitor --unmanaged** and send results to the Organization of your choice. This provides license/policy information in a viewable state.

#### **Dependency lists**

Use **--print-deps** when performing open source scans to obtain a detailed list of discovered dependencies in your codebase and their origin source.

In C/C++, this has the additional benefit of identifying the confidence level of a given match. If there is a significant drop (< 90% confidence), it's likely the file has been modified and may not be the original source. Consider investigating if that's the case.

```
snyk test --unmanaged --print-deps
```

The list is printed before the issues list, as shown below:

<figure><img src="https://lh5.googleusercontent.com/x4y1uIQ2fCFX956f1eP4664i6VKEgK6eOOddlAZ4p4WnQWJu1t_ugSOpL394KEnuzSIPRs08gNAsmjvPa-GAV0C-975esRdy0EPDY7WImG1-SXSOFO0TIAVfh_Jp2DLYc6bm7iZu55UbE3Boh4TNk_I" alt="View dependency lists"><figcaption><p>View dependency lists</p></figcaption></figure>

#### **License policy text during the Beta phase**

{% hint style="info" %}
[License Compliance](../../snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md) allows a company to create a license policy for your Open Source, indicating what licenses are not approved for use. To access [License Compliance](../../snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md), you must be on a Snyk Team or Enterprise [plan](https://snyk.io/plans). Snyk detects and alerts when a match is found. The alert contains the name of the license and license policy text.&#x20;

**License policy text** is the text associated with the issue by your administrators that provides custom direction on what to do or why it's contrary to the policy, if it's found in your application.
{% endhint %}

For the Beta, while the issues are displayed, the license policy text associated with the issues is not currently displayed within the IDE or CLI. This is planned to be addressed by GA. To see the policy text, use the Snyk Web UI.&#x20;

Below, you can see the license policy text example at the bottom of the screen, giving you directions on what to do if the license is found.

<div align="left">

<figure><img src="https://lh4.googleusercontent.com/lIn5JFEyaZaTNMVenBoeGIgTpC6YHxpmAjK947z5ISPlHV1rlOvPNCLyzXxsGNj65AAlGn6ff9dF4lHVsVFYMaKXWC939tasD91k98xcDv_Ske6Dz7goMXl5lByyqg6ptvvqaK0UEqLSdzUU9GKrW4U" alt="License policy text example"><figcaption><p>License policy text example</p></figcaption></figure>

</div>

#### **Alternate testing options**

If you develop advanced dependency management strategies, you might not use the standard/well-traveled package managers. For that reason, Snyk has provided test APIs. In the case of C++, if you know the OS packages/versions that are included in the application but don't have the source code, the purl API can be used to do this analysis.

:link: PURL See [Purl issues](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org\_id-/packages/-purl-/issues)&#x20;

### **Options and plugins**

* To help generate reports locally or at build time, see [snyk-to-html plugin](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md).
* See `--json` and `--sarif` options for generating output that can be programmatically accessed.
* For advanced filtering options, see[ snyk-filter](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md).
