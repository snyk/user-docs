# View licenses

The **Licenses** area displays all licenses that have been detected for your Projects, as well as a summary of all dependencies in your Projects, and a summary of all of your Projects using the license. This allows you to understand which Projects and dependencies use a license.

For example:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 15.38.44.png" alt="Licenses tab"><figcaption><p>Licenses tab</p></figcaption></figure>

</div>

## **Field details**

<table><thead><tr><th width="162">Field</th><th>Description</th></tr></thead><tbody><tr><td>Severity</td><td>The assessed severity level. You can set the severity level in the license policy (see <a href="../policies/license-policies/create-a-license-policy-and-rules.md">Create a license policy and rules</a>). Snyk's default license policy also defines the Severities for a few licenses.</td></tr><tr><td>License</td><td>The full official name of the license. SPDX licenses are linked to the <a href="https://spdx.org">SPDX</a> site, and non-SPDX licenses are linked to their respective sites.<br>If the license is marked as <strong>Unknown</strong>, Snyk could not find a license for this package.</td></tr><tr><td>Dependencies</td><td>The total number of dependent packages with this license in your projects, linked to a side panel that displays the full list of affected dependencies in the same layout as the <strong>Dependencies</strong> area.</td></tr><tr><td>Projects</td><td>The total number of your projects using this license, linked to a side panel that displays the full list of your affected projects, with the same layout as the <strong>Dependencies</strong> area.</td></tr></tbody></table>

## **Controls and filters**

These controls appear at the top of the window:

<div align="left">

<figure><img src="../../.gitbook/assets/Screenshot 2023-05-11 at 15.50.22.png" alt="Licenses tab actions"><figcaption><p>Licenses tab actions</p></figcaption></figure>

</div>

* **Search for Licenses**—the dynamic search field enables you to enter free text and begins searching with the first character you type; alternatively, select multiple packages from the dropdown list that opens when you click in the field. In addition, click the Select All or Deselect All links that dynamically appear in the upper right-hand corner of the dropdown list.
* **Export as CSV**—export issue data in CSV file format.
