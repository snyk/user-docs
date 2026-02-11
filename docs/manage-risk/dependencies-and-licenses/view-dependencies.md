# View dependencies

The **Dependencies** tab acts as a Bill Of Materials (BOM) for all dependencies in all Projects in the selected Organization. This allows you to quickly and easily identify which Projects have a specific version of a dependency.

Dependency reports show details about the packages included in your Projects, including their full names, the version of the package currently used, the Projects in which they are used, and a summary of the issues they contain:

<div align="left"><figure><img src="../../.gitbook/assets/image (20) (1).png" alt="Dependencies tab"><figcaption><p>Dependencies tab</p></figcaption></figure></div>

Examining the details for each package can help you determine the dependency health of your packages.

## Field details

<table><thead><tr><th width="179">Element</th><th>Description</th></tr></thead><tbody><tr><td><strong>Dependencies</strong></td><td><p>The full package names of the dependencies contained in at least one of your Projects. Click the link to view the detailed Package pag<em>e.</em></p><p>For npm only, a <strong>warning icon</strong> appears next to the package name if a package is deprecated.</p></td></tr><tr><td><strong>Version</strong></td><td>The version of the package used in your Projects. Use this and the <strong>Latest Version</strong> to see the difference between your current package version and the most recent package version available.</td></tr><tr><td><strong>Latest version</strong> and <strong>Last publish</strong></td><td><p>The most recent version updated by a maintainer for this package in its repository, and the last time a new version of the package was published by a maintainer. Look at these dates to help you determine the maturity of the package and activity frequency</p><p><br><strong>Latest version</strong> and <strong>Last publish</strong> are supported for npm and Maven only.</p></td></tr><tr><td><strong>Vulnerabilities</strong></td><td>Each row shows the icon of the associated severity for the issues, <strong>C</strong>ritical, <strong>H</strong>igh, <strong>M</strong>edium, <strong>L</strong>ow. The vulnerabilities shown in the table are the total number of vulnerable paths associated with that dependency and version across all Snyk Projects. For example, if there is a dependency with one critical vulnerability in two Snyk Projects, but in one of those Projects the dependency (vulnerability) is brought in on two paths, Snyk shows three critical vulnerabilities associated with the dependency on the Dependencies tab.</td></tr><tr><td><strong>License</strong></td><td><p>The license or licenses used by this package. These can be:</p><p><strong>Known license name</strong>: Snyk identified the package and its associated license type, and this information is shown in the list.<br><strong>Unknown</strong>: Snyk identified the package but could not identify its associated license type.<br><strong>Blank</strong>: Snyk could not identify the package and therefore has no license or other information for the package.</p></td></tr><tr><td><strong>Copyrights</strong> (until January 8, 2024)</td><td>For npm, PyPI, and Maven, <strong>copyright information</strong> for the license.</td></tr><tr><td><strong>Projects</strong></td><td>The Projects in which this package is used by your Organization.</td></tr><tr><td><strong>Dependencies with issues</strong></td><td>A link to the dependencies in the package that have issues, with details about those issues.</td></tr></tbody></table>

## Dependencies tab actions

The actions appear at the top of the tab.

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 13.19.55.png" alt="Dependencies tab actions"><figcaption><p>Dependencies tab actions</p></figcaption></figure>

* **Search for Dependencies**: Start typing to search for a package. To view the results for multiple packages, select them from the dropdown list that opens when you click the field. You can also click the **Select All** or **Deselect All** links that appear dynamically in the dropdown list.
*   **Dependency filters**: Select specific Project types and dependency health status to mark the packages to be displayed. Only issues matching all selected criteria are displayed.

    When you select **Deprecated**, only packages marked as deprecated are displayed.
* **Hidden fields**: Remove any of the default columns from the display to focus on details that are important to your current tasks.
* **Export as CSV**: Export issue data in CSV file format.
