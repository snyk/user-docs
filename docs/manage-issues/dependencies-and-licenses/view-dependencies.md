# View dependencies

The **Dependencies** tab acts as a Bill Of Materials (BOM) for all the direct dependencies in all of the Projects in the selected OrganizationProjects. This allows you to quickly and easily identify which Projects have a specific version of a dependency.

Dependency reports show details about the packages included in your Projects, such as their full names, the version of the package currently used in your Projects, and the Projects in which they are used, as well as a summary of the issues they contain:

<div align="left">

<figure><img src="../../.gitbook/assets/image (2) (2) (2).png" alt="Dependencies tab"><figcaption><p>Dependencies tab</p></figcaption></figure>

</div>

This assists in determining dependency health—or in other words, the health of your packages, by displaying important details per package.

## Field details

<table><thead><tr><th width="179">Element</th><th>Description</th></tr></thead><tbody><tr><td><strong>Dependency</strong></td><td>The full package names of the dependencies contained in at least one of your Projects. Click the link to view the detailed Package pag<em>e.</em></td></tr><tr><td>Warning icon</td><td><p>A warning icon appears next to the package name if a package is deprecated.<br></p><p><strong>Note:</strong> Only currently supported for npm.</p></td></tr><tr><td><strong>Version</strong></td><td>The version of this package used in your Projects. Use this and the <strong>Latest Version</strong> to see the difference between your current package version and the most recent package version available.</td></tr><tr><td><strong>Latest version</strong> and <strong>Last publish</strong></td><td><p>The most recent version updated by a maintainer for this package in its repository, and the last time a new version of the package was published by a maintainer. With these dates, you can more easily determine the maturity of the package, as well as activity frequency</p><p><br><strong>Note:</strong> Only currently supported for npm and Maven.</p></td></tr><tr><td><strong>Vulnerabilities</strong></td><td>The icon of the associated severity (critical/high/medium/low) for this issue.</td></tr><tr><td><strong>License</strong></td><td>The license or licenses used by this package. These can be:<br>- Known license: Snyk identified the package and its associated license type, and this information is shown in the list.<br>- Unknown license: Snyk identified the package but could not identify its associated license type, and this appears as <strong>Unknown</strong> in the list. <br>- Blank: Snyk could not identify the package and so has no information about it, and this appears as blank in the list.</td></tr><tr><td><strong>Copyrights</strong></td><td>Copyright information for the license. <br><br><strong>Note:</strong> Only currently supported for npm, PyPI and Maven.</td></tr><tr><td><strong>Projects</strong></td><td>The projects in which you use this package.</td></tr><tr><td><strong>Dependencies with issues</strong></td><td>A link to the dependencies in the package that have issues, with details about those issues.</td></tr></tbody></table>

## Controls and filters

These controls appear at the top of the window:

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 13.19.55.png" alt="Dependencies tab actions"><figcaption><p>Dependencies tab actions</p></figcaption></figure>

* **Search for Dependencies**—start typing to search for a package. To view the results of multiple packages, select them from the dropdown list that opens when you click in the field. In addition, you can also click the Select All or Deselect All links in the upper right-hand corner of the dropdown list.
*   **Dependency filters**—mark the packages to be displayed by selecting specific project types as well as by dependency health status. Only issues matching all selected criteria are displayed.

    When Deprecated is selected, only packages that are marked as deprecated will be displayed.
* **Hidden fields**—remove any of the default columns from the display in order to focus on details that are important to your current tasks.
* **Export as CSV**—export issue data in CSV file format.
