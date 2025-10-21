# View licenses

The **Licenses** tab displays all licenses detected for your Projects, with summaries of all dependencies in your Projects and all of your Projects using the license. This allows you to see which Projects and dependencies have a license. An example follows:

<div align="left"><figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 15.38.44.png" alt="Licenses tab"><figcaption><p>Licenses tab</p></figcaption></figure></div>

## **Field details**

<table><thead><tr><th width="162">Field</th><th>Description</th></tr></thead><tbody><tr><td><strong>Severity</strong></td><td>The assessed severity level. You can set the severity level in the license policy; see <a href="../../policies/license-policies/create-a-license-policy-and-rules.md">Create a license policy and rules</a>. The default Snyk license policy also defines the Severities for a few licenses.</td></tr><tr><td><strong>License</strong></td><td>The full name of the license. SPDX licenses are linked to the <a href="https://spdx.org/licenses/">SPDX</a> site, and non-SPDX licenses are linked to their respective sites.<br>If the license is marked as <strong>Unknown</strong>, Snyk could not find a license for this package.</td></tr><tr><td><strong>Dependencies</strong></td><td>The total number of dependent packages with this license in your Projects, linked to a side panel that displays the full list of affected dependencies in the same layout as the <strong>Dependencies</strong> tab.</td></tr><tr><td><strong>Projects</strong></td><td>The number of Projects with this license, linked to a side panel that displays the full list of your affected Projects, with the same layout as the <strong>Dependencies</strong> tab.</td></tr></tbody></table>

## **Licenses tab actions**

The actions appear at the top of the tab.

<div align="left"><figure><img src="../../../.gitbook/assets/Screenshot 2023-05-11 at 15.50.22.png" alt="Licenses tab actions"><figcaption><p>Licenses tab actions</p></figcaption></figure></div>

* **Search for Licenses**: Enter free text and begins searching with the first character you type. You can also select multiple packages from the dropdown list that opens when you click the field. You can also click the **Select All** or **Deselect All** links that appear dynamically in the dropdown list.
* **Export as CSV**: Export issue data in CSV file format.
