# Issues tab

The Issues tab displays all known vulnerability and license discrepancies across your projects, indicating details about each issue and which of your projects is affected. In addition, this area also displays information about which issues can be fixed, and how.

By default, issues are grouped together per issue, similar to the following example:

![](../../.gitbook/assets/mceclip0-31-.png)

Get an overview of the general health of your organization on the whole by inspecting the number of projects affected per issue and the number and kinds of issues affecting your projects.

Alternatively, you can ungroup data, viewing a separate line for each project in which the issue occurs—meaning the same issue can appear multiple times if it affects multiple projects. Click the **View issues ungrouped** or **View issues grouped** link to toggle between views.

This view is beneficial in providing additional details about:

* Each affected project
* The recommended course of remediation

This view can be seen in the Elements column below, where all the elements of both views are described.

{% hint style="info" %}
**Note**  
Data in each of the four tabs is displayed based on the filters you've applied from the top of the Reports area, as well as the group or organization that you're viewing from.
{% endhint %}

## Issues tab elements

Just under the **Export** button, click the **View issues ungrouped** or **View issues grouped** link to toggle between views.

The Issues area appears similar to this image when viewing issues ungrouped:

![](../../.gitbook/assets/uuid-e02b41ac-13ff-99e8-1028-69cd24310313-en.png)

The following table describes the different parts of the Issues area as displayed when viewing issues either grouped or ungrouped:

<table>
  <thead>
    <tr>
      <th style="text-align:left">Element</th>
      <th style="text-align:left">Description</th>
      <th style="text-align:left"><b>Possible values</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Issue tab badge</td>
      <td style="text-align:left">The number of issues that you&#x2019;re currently viewing, based on the
        filters and searches you used, are displayed on the issues tab.</td>
      <td
      style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Severity</td>
      <td style="text-align:left">The icon of the associated severity for this issue.</td>
      <td style="text-align:left">
        <p>
          <img src="../../.gitbook/assets/mceclip2-7-.png" alt="mceclip2.png" />Critical</p>
        <p>
          <img src="../../.gitbook/assets/mceclip3-3-.png" alt="mceclip3.png" />High</p>
        <p>
          <img src="../../.gitbook/assets/mceclip4 (2) (1) (1) (4) (2).png" alt="mceclip4.png"
          />Medium</p>
        <p>
          <img src="../../.gitbook/assets/mceclip6 (4) (5).png" alt="mceclip6.png"
          />Low</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Priority Score</td>
      <td style="text-align:left">
        <p>Snyk processes a number of issue characteristics in a proprietary algorithm
          that weights each, producing a score that can be used to guide the order
          in which you fix issues.</p>
        <p><a href="https://support.snyk.io/hc/en-us/articles/360009884837">Find out more about Priority Score</a>
        </p>
      </td>
      <td style="text-align:left">0-1000</td>
    </tr>
    <tr>
      <td style="text-align:left">Issue</td>
      <td style="text-align:left">The full officially recognized name of the issue and a list of all affected
        packages that are contained by your projects. The issue is linked to the <b>Package</b> page.</td>
      <td
      style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Identifiers</td>
      <td style="text-align:left">All associated CVE identifiers. Each identifier is linked individually
        to the full official CVE or CWE vulnerability details as relevant.</td>
      <td
      style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Projects</td>
      <td style="text-align:left">
        <p>When viewing issues ungrouped, this is a complete list of all projects
          affected by the specific issue, and an indicator of the source of the projects.</p>
        <p>If viewing the issues grouped, this column displays the number of projects
          affected by the issue.</p>
        <p>Click on the total of projects to open a panel where a list of all affected
          projects in that grouping is displayed. Details in this view include:</p>
        <ul>
          <li>Project</li>
          <li>Status</li>
          <li>Introduced-date the issue was detected in the project</li>
          <li>Fixable&#x2014;whether the issue can be eliminated with an upgrade or
            patch</li>
        </ul>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Exploit maturity</td>
      <td style="text-align:left">
        <p>These values are available:</p>
        <ul>
          <li><b>Mature:</b> a published code exploit that can easily be used for this
            vulnerability is available.</li>
          <li><b>Proof of concept:</b> a published, theoretical proof-of-concept or detailed
            explanation that demonstrates how to exploit this vulnerability is available.</li>
          <li><b>No known exploit:</b> neither a proof-of-concept code nor an exploit
            were found for this vulnerability, or are not publicly available.</li>
          <li><b>No data</b>: this value indicates one of the following:
            <ul>
              <li>The issue is not a vulnerability (but rather, a license issue);</li>
              <li>The ecosystem is not currently supported by Snyk (Linux); or</li>
              <li>The project was imported prior to the release of this feature</li>
              <li>Reimport the project in order to scan for this data</li>
            </ul>
          </li>
        </ul>
      </td>
      <td style="text-align:left"></td>
    </tr>
  </tbody>
</table>

Following are the project level details that appear only when viewing issues ungrouped:

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Element</b>
      </th>
      <th style="text-align:left"><b>Description</b>
      </th>
      <th style="text-align:left"><b>Possible values</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Fixable</td>
      <td style="text-align:left">Indicates an upgrade or patch is available to fix the vulnerability.
        <br
        />Also see <a href="https://support.snyk.io/hc/en-us/articles/4405034808209-Fixed-in-version-vs-fixable-attribute-in-vulnerabilities">Fixed in version vs. fixable attribute in vulnerabilities</a>.</td>
      <td
      style="text-align:left">
        <p><b>Patch</b>: Issues that can be fixed with a Snyk patch.</p>
        <p><b>Upgrade</b>: Issues that can be resolved by upgrading the affected
          package.</p>
        <p><b>Pin</b>: Issues that can be fixed by making a transitive dependency
          into a direct dependency.
          <br /><b>Note:</b> this is currently Python-specific.</p>
        <p><b>No</b>: Issues for which there is currently no known fix.</p>
        </td>
    </tr>
    <tr>
      <td style="text-align:left">Introduced</td>
      <td style="text-align:left">The date the issue was introduced in the project</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Status</td>
      <td style="text-align:left">The current status of a vulnerability</td>
      <td style="text-align:left">
        <p>Open: Issues that have not been handled</p>
        <p>Fixed: Issues for which Fix PRs have been submitted (automatically by
          Snyk)</p>
        <p>Patched: Issues that have been fixed with Snyk patches</p>
        <p>Ignored: Issues to which the Ignore policy applies</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Reachability</td>
      <td style="text-align:left">
        <p>Indicates whether or not vulnerabilities have a path from the code to
          the vulnerable function</p>
        <p><a href="https://support.snyk.io/hc/en-us/articles/360010554837-Reachable-Vulnerabilities-">Find out more about Reachable Vulnerabilities</a>
        </p>
      </td>
      <td style="text-align:left">
        <p>Reachable: There is a clear path from the app&#x2019;s code down to the
          vulnerable function</p>
        <p>Potentially reachable: There are indications that you might be exposed
          to the vulnerability</p>
        <p>No path found: We don&#x2019;t have enough information to decide whether
          the vulnerability is reachable or not</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Jira issue</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">When a Jira integration is configured for the project and a Jira issue
        has been filed against the issue in Snyk, this column displays the Jira
        key and links to that same issue within Jira.</td>
    </tr>
  </tbody>
</table>

## Issues tab actions

These controls appear above the table:

![Reports\_Issues\_filters - REPLACE THIS IMAGE - ZENDESK IMAGE - UPDATE ME!](https://support.snyk.io/hc/article_attachments/360007031338/uuid-ef7a494d-8b10-9b28-dc63-3f9224519070-en.png)

**Search issues**—this search capability allows users to search based on CVE, CWE or identifier name \(i.e. DDoS\). When searching by CVE or CWE, you must provide an exact value \(for example CVE-1234\) whereas, for identifier name, typing in a piece of the word will return results.

**Issue filters**—mark the issues to be displayed by selecting a range of priority scores, specific issue types, exploit maturity, status, fixable values and reachability.

![](../../.gitbook/assets/screenshot_2020-07-30_at_11.30.19_am.png)

**Export**—click the button to choose which format you’d like to export issue data in:

* CSV
* Print/generate a preview from the Print dialog box in your local environment. Once you click the initial **Print** button, in the backend we load all the issues into one page. This process can take a few seconds. Make sure you don’t close your tab/browser or refresh the page while we prepare the document. Once ready, you get a new dialog box informing you that the document is ready to print. Within the print dialog, you select to save as PDF or print directly.

{% hint style="info" %}
**Note**  
Only 2000 issues can be generated at once. The number of issues you’re currently viewing is displayed on the Issues tab, as described in the Elements table.
{% endhint %}

**View issues ungrouped/View issues grouped**—toggle between views. When viewing “grouped”, each issue appears once only, grouping together all affected projects; when viewing “ungrouped”, each issue appears once per project that is affected by it.

