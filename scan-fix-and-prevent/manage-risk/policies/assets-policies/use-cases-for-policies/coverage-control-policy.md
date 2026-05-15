# Coverage control policy

You can use the **Set Coverage Control** action from the Policies view to identify your application assets, monitor them, and prioritize the risks. You can select one or multiple security products and also specify a timeframe for when the last scan should have taken place.

Identifying and setting coverage policies allows your team to define where certain security controls are expected to be in place.

The following example filters out assets that should have Snyk Open Source and Snyk Code security controls in place and then sets the coverage policies.

<figure><img src="../../../../.gitbook/assets/image (126).png" alt="AppRisk - Setting up a Coverage Control policy"><figcaption><p>Assets Policy - Setting up a Coverage Control policy</p></figcaption></figure>

To follow the example, these are the filters you need to apply:

* **Filter 1**: include only assets that are repositories.
* **Filter 2**: include assets that have coding language tags relevant to your application (Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C).

Next, you need to define two actions, one for Snyk Open Source, and one for Snyk Code, since the example is focused on these two security controls.

*   **Action 1**: add a Set Coverage Control Policy action with Snyk Open Source selected.

    In Snyk, by default, Open Source manifest files that are imported using the SCM integration are scanned on a daily basis. The Coverage Control Policy needs to check that a Snyk Open Source scan occurred for that repository in the last two days.
*   **Action 2**: add a Set Coverage Control Policy action with Snyk Code selected.

    For Snyk Code, scans happen by default once a week, or when changes have been pushed to the monitored branch. The Coverage Control Policy needs to check that a Snyk Code scan occurred for that repository in the last week.

In the Inventory view, any coverage gap is indicated with strikes through the control coverage icon. See more details about each icon on the [Inventory capabilities](../../../../manage-assets/assets-inventory-components.md) page.
