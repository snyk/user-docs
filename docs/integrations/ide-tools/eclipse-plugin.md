# Eclipse plugin

Install our **Snyk Vuln Scanner** in your Eclipse workflow to catch vulnerabilities and license issues directly from within your IDE \(integrated development environment\), before they are merged into your codebase.

Once installed and configured, every time you run the plugin, Snyk scans your project’s manifest files and:

* analyzes and delivers actionable vulnerability and license issue details
* records results per package
* displays results directly from the Eclipse UI

The Snyk plugin enables you to track and identify issues that risk your application’s security and avoid ever adding those issues to your shared repo.

{% hint style="info" %}
**Note**  
Both Snyk Broker and on-prem solutions support the Eclipse plugin.
{% endhint %}

### Supported languages and repos

Snyk supports all languages that are supported by both Eclipse and Snyk. Additionally, the Snyk plugin can also be implemented with our Broker and on-prem solutions.

### Installing the Eclipse Snyk plugin

1. Navigate to the Marketplace from within your running Eclipse instance.
2. Search for Snyk and click **Install**. 
3. When prompted accept the license agreement and the **Snyk Security** certificate to complete the installation.
4. Restart the Eclipse instance and navigate to **Eclipse Preferences** to ensure **Snyk Vuln Scanner** now appears in the list: 

![](../../.gitbook/assets/uuid-01198b42-f020-2cc5-c20f-93817eeb44a4-en.png)

![](../../.gitbook/assets/uuid-928012b7-8e49-fe6f-4965-77c5db026784-en.png)

### Use the Snyk plugin to secure your Eclipse projects

From the Snyk results click ![](../../.gitbook/assets/uuid-aa090aa8-d4fe-eb5d-2505-54a0b1555be9-en.png) whenever you are ready to scan your projects. It shouldn’t take too long for the results to appear—but no worries! You can continue to work as usual in the meantime anyway.

If for any reason you need to stop the scan before the build ends, click: ![](../../.gitbook/assets/uuid-29be01e6-6913-25f8-15ed-a8cf47230fa0-en.png) If you only want to scan a single project in your workspace, navigate to the Package Explorer panel, right-click the root of the project you want to test, and then choose **Snyk test**.

When the scan ends, results and any relevant error messages as well, are displayed from the **Snyk results**, grouped by project similar to the following:

![](../../.gitbook/assets/uuid-e868f739-eb55-9bd6-be33-acbb230ec1fa-en.png)

Work with Snyk results from Eclipse as follows:

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Column</b>
      </th>
      <th style="text-align:left"></th>
      <th style="text-align:left"><b>Description</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>Context menu</b>
      </td>
      <td style="text-align:left">Right-click menu</td>
      <td style="text-align:left">
        <p>Options include:</p>
        <p>Ignore issue&#x2014;Hover over the specific issue that you want to ignore
          for the next 30 days and then access the context menu.</p>
        <p>Snyk test&#x2014;Run the Snyk test for the entire workspace.</p>
        <p>Preferences&#x2014;Access and update Snyk Vuln Scanner preferences directly
          from the right-click menu.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>When collapsed</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Title</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">The name of the project.</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Dependency</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">A summary of vulnerabilities and the number of affected paths found per
        project.</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>When expanded</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Title</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">The full name of the vulnerability affecting your project, linked to a
        description and complete details of the vulnerability in our database,
        to assist you in resolving the issue.</td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Dependency</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">The name of the direct dependency package in your project (the package
        you explicitly installed) that is affected by the vulnerability, either
        directly or indirectly.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
      <td style="text-align:left">
        <p>All details appear on a single row and the Dependency (the name of the
          package explicitly used in the code) and Package (the name of the package
          that actually contains the vulnerability) columns both display the name
          of the same package:</p>
        <p>
          <img src="../../.gitbook/assets/uuid-e7accdc1-7495-e7a5-7a64-2403b066cb03-en.png"
          alt="image13.png" />
        </p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">
        <p><b>When your project is affected by an indirect vulnerability:</b>
        </p>
        <p><b>Collapsed mode</b>
        </p>
      </td>
      <td style="text-align:left">
        <p>An arrow appears on the row, grouping together all relevant details, similar
          to the following examples:</p>
        <p>
          <img src="../../.gitbook/assets/uuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png"
          alt="image14.png" />
        </p>
        <p>For example:</p>
        <p>Package X uses Package Y, which in turn uses Package Z.</p>
        <p>Package Z contains a Cross-Site Scripting (XSS) vulnerability, indirectly
          affecting your project.</p>
        <p>The Dependency (the name of the package explicitly used in the code) is
          Package X; the Package field displays Package Z (the name of the package
          that actually contains the vulnerability).</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left"><b>Expanded mode</b>
      </td>
      <td style="text-align:left">
        <p>Click the arrow on the row to expand and view the full path from the direct
          dependency to the actual vulnerable package.</p>
        <p>
          <img src="../../.gitbook/assets/uuid-35658aaf-3359-80c2-c094-41a34c7863cc-en.png"
          alt="image15.png" />
        </p>
        <p>In the example above, the full path would appear as:</p>
        <p>[Name of Package X]--&gt;[Name of Package Y]--&gt;[Name of Package Z]</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Package</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">
        <p>The name of the package in your project that is directly affected by the
          vulnerability.</p>
        <p>In the example above:</p>
        <ul>
          <li>the Dependency is indicated as Package X&#x2014;this is the package the
            developer explicitly uses in the code</li>
          <li>the Package field displays Package Z, which is the package that actually
            contains the vulnerability.</li>
        </ul>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Fix</b>
      </td>
      <td style="text-align:left"></td>
      <td style="text-align:left">The name of the package, if such exists, and the version that it can be
        upgraded to in order to resolve the issue.</td>
    </tr>
  </tbody>
</table>

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}