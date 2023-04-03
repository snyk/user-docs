# Reports: Licenses tab

The **Licenses** area displays all licenses that are currently used in your Project, as well as a summary of all dependencies in your Projects and a summary of all of your Projects using the license.

<figure><img src="../../../.gitbook/assets/image (346).png" alt="Licenses overview in a dependency project."><figcaption><p>Licenses overview in a dependency project</p></figcaption></figure>

**Licenses tab elements**

The following table describes the different parts of the **Licenses** area:

| **Element**  | **Description**                                                                                                                                                                                                                                                      |
| ------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| License      | The full official name of the license is linked to additional details and information. SPDX licenses are linked to the [SPDX](https://spdx.org/) site, while non-SPDX licenses are linked to their respective site.                                                  |
| Dependencies | The total number of dependent packages with this license in your Projects is linked to a side panel that displays the full list of affected dependencies in the same layout as the [**Dependencies** area](https://snyk.io/?post\_type=docs\&p=12382\&preview=true). |
| Projects     | The total number of your Projects using this license is linked to a side panel that displays the full list of your affected Projects, with the same layout as the **Dependencies** tab.                                                                              |

**Licenses tab actions**

These controls appear at the top of the window:

<figure><img src="../../../.gitbook/assets/uuid-8399334e-74b7-0649-d55c-e0ddecb54272-en.png" alt="Licensing setting toolbar."><figcaption><p>Licensing setting toolbar</p></figcaption></figure>

* **Search for Licenses**—the dynamic search field enables you to enter free text and begins searching with the first character you type; alternatively, select multiple packages from the dropdown list that opens when you click in the field. In addition, click the Select All or Deselect All links that dynamically appear in the upper right-hand corner of the dropdown list.
* **Search for Projects**—the dynamic search field enables you to enter free text and begins searching with the first character you type; alternatively, select multiple packages from the dropdown list that opens when you click in the field. In addition, click the Select All or Deselect All links that dynamically appear in the upper right-hand corner of the dropdown list.
* **License filters**—mark the packages to be displayed by selecting specific Project types. Only issues matching all selected criteria are displayed.

<figure><img src="../../../.gitbook/assets/uuid-53b0da21-ca9b-a04c-354a-97219ae7c05b-en-1-.png" alt="Licensing filtering by type."><figcaption><p>Licensing filtering</p></figcaption></figure>

* **Export as CSV**—export issue data in CSV file format.
