# View cloud issues in the Snyk Web UI

You can view cloud issues for an Organization through the Snyk Web UI.

To view your cloud issues in the Snyk Web UI, navigate to the Organization, and on the menu, select **Cloud.**

<figure><img src="../../../../.gitbook/assets/snyk-cloud-access-issues-page-3.png" alt=""><figcaption><p>Select Cloud to view the issues</p></figcaption></figure>

By default, Snyk displays open issues across all cloud environments in an Organization. The issues are initially [grouped by the security rule they failed](view-cloud-issues-in-the-snyk-web-ui.md#group-issues-by-rule), but you can choose to [group them by resource](view-cloud-issues-in-the-snyk-web-ui.md#group-issues-by-resource).

## Group issues by rule

Issues are grouped by rule by default.

If you have grouped them by resource instead, you can group them by rule again by selecting the **Group by Resource** drop-down menu next to the search bar and selecting **Rule**:

<figure><img src="../../../../.gitbook/assets/snyk-cloud-how-to-group-by-rule-3.png" alt=""><figcaption><p>Group by rule</p></figcaption></figure>

The rules are sorted by severity and then by number of issues, both from highest to lowest.

## Rule details

Each rule shows the following information:

* Severity
  * C: Critical
  * H: High
  * M: Medium
  * L: Low
* Rule title
* Rule ID
* Number of open issues

<figure><img src="../../../../.gitbook/assets/snyk-cloud-grouped-by-rule-3.png" alt=""><figcaption><p>Issues grouped by rule</p></figcaption></figure>

By default, when you expand a rule you see all the resources with an open issue associated with that rule. Each resource shows the following information:

<figure><img src="../../../../.gitbook/assets/snyk-cloud-grouped-by-rule-resource-4.png" alt=""><figcaption><p>A rule expanded to show resources</p></figcaption></figure>

* Resource name
* Environment kind (provider)
* Environment name
* Environment native ID (Amazon Web Services account ID, Google project ID, Azure subscription, and so on)
* Resource tags
* Age of issue

## Group issues by resource

If issues are currently grouped by rule, you can group them by resource instead by selecting the **Group by Rule** drop-down menu next to the search bar and selecting **Resource**:

<figure><img src="../../../../.gitbook/assets/snyk-cloud-how-to-group-by-resource-3.png" alt=""><figcaption><p>Group by resource</p></figcaption></figure>

The resources are sorted by severity of issue and then by number of issues, both from highest to lowest.

## Resource details

Each resource shows the following information:

* Resource name
* Resource type
* Environment kind
* Environment name
* Environment native ID
* Resource tags
* Severity of open issues
* Number of issues for each severity

<figure><img src="../../../../.gitbook/assets/snyk-cloud-grouped-by-resource-3.png" alt=""><figcaption><p>Issues grouped by resource</p></figcaption></figure>

By default, when you expand a resource you see all the rules with an open issue associated with that resource.

<figure><img src="../../../../.gitbook/assets/snyk-cloud-grouped-by-resource-rule-4.png" alt=""><figcaption><p>A resource expanded to show rules</p></figcaption></figure>

Each rule shows the following information:

* Severity
* Rule title
* Rule category
* Rule ID
* Age of issue

## View issue details on issue summary panel

To view an issue, expand an issue then select a corresponding resource or rule. Snyk displays a panel with the cloud issue details.

The cloud issue summary panel displays the **Overview** by default, which contains the following information:

* Severity
* Status (open or closed)
* Resource details:
  * Native ID (Amazon Resource Name); AWS only
  * Name
  * Type
  * Location (AWS region, Google project, Azure subscription, Terraform file, and so on)
  * Input type (`cloud_scan`, `tf_hcl`, and so on)
  * Platform (`aws`, `google`, `azure`, and so on)
  * ID
* Environment details:
  * Name
  * Native ID (AWS account ID, Google project, Azure subscription, `:cli`, and so on)
  * Kind (AWS, Google, Azure, CLI, and so on)
  * ID
* Rule details:
  * Title
  * Category
  * Severity
  * ID, with a link to fix advice
  * Rule description (**Why to fix it** section)

<figure><img src="../../../../.gitbook/assets/snyk-cloud-issue-panel-overview-2.png" alt=""><figcaption><p>Issue summary panel, Overview tab</p></figcaption></figure>

Select the **Resource** tab to view the scanned resource's attributes:

<figure><img src="../../../../.gitbook/assets/snyk-cloud-issue-panel-resource-2.png" alt=""><figcaption><p>Issue summary panel, Resource tab</p></figcaption></figure>

## Filter issues

To filter which cloud issues are shown:

1. Select the **Filter** drop-down menu. The name of the menu shows how many filters are selected, for example, "1 Filter".
2. Select the parameter you want to filter by, for example "Severity".
3. Check the box for the values you want to show, for example "High".

By default, the **Status** filter is set to include only open issues.

You can also filter issues by compliance control, resource type, environment name, and more.

<figure><img src="../../../../.gitbook/assets/snyk-cloud-issue-filters-4.png" alt=""><figcaption><p>Filters on the issues page</p></figcaption></figure>

## Search issues

As you select filters from the **Filters** drop-down menu, the search bar is populated with key-value pairs.

To search cloud issues by key-value pair:

1. Select the search bar.
2. Select a key from the drop-down list, or type in a key from the list of available filters.
3. Select a value from the drop-down list, or type in a value.

To remove a key-value pair, select the **X** or press the **Delete** key on your keyboard.
