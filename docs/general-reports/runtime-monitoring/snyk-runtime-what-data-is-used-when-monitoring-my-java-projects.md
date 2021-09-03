# Snyk runtime: what data is used when monitoring my Java projects?

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

Once you have installed the Snyk agent to monitor your projects at runtime, Snyk takes the following data from your deployment in order to properly monitor and offer remediation.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Category</th>
      <th style="text-align:left">Parameter</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Snyk project ID</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">The unique Snyk ID for the project that is monitored.</td>
    </tr>
    <tr>
      <td style="text-align:left">Unique runtime agent ID</td>
      <td style="text-align:left"></td>
      <td style="text-align:left">An ID generated at runtime to identify the Snyk agent.</td>
    </tr>
    <tr>
      <td style="text-align:left">System info</td>
      <td style="text-align:left">Runtime agent version</td>
      <td style="text-align:left">The Snyk agent version.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Hostname</td>
      <td style="text-align:left">The name of the machine on which the Snyk agent is installed.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">OS Details</td>
      <td style="text-align:left">Including info such as operating system type and version</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">JVM details</td>
      <td style="text-align:left">
        <p>VM name</p>
        <p>Vendor</p>
        <p>Version</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">List of third-party package vulnerable monitored functions</td>
      <td style="text-align:left">Details about the third-party packages with vulnerable functions that
        are observed by Snyk at runtime.</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">JAR name</td>
      <td style="text-align:left">Name of the JAR that was seen.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Relative file path</td>
      <td style="text-align:left">Of the vulnerable function within the JAR.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Function name</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Artifact ID</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Class names</td>
      <td style="text-align:left">Class names for the function.</td>
    </tr>
    <tr>
      <td style="text-align:left">List of third-party package vulnerable functions, called at runtime</td>
      <td
      style="text-align:left">Details about the third-party packages with vulnerable functions that
        are called at runtime.</td>
        <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">JAR name</td>
      <td style="text-align:left">Name of the JAR that was seen.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">File path</td>
      <td style="text-align:left">Of the vulnerable function within the JAR.</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Function name</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Artifact ID</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Class names</td>
      <td style="text-align:left">
        <p>Class names for the function.</p>
        <p>Java only</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">JAR crc32</td>
      <td style="text-align:left">crc32 JAR checksum</td>
    </tr>
    <tr>
      <td style="text-align:left">List of all JAR files loaded at runtime</td>
      <td style="text-align:left">JAR file path</td>
      <td style="text-align:left">The file path</td>
    </tr>
    <tr>
      <td style="text-align:left"></td>
      <td style="text-align:left">Artifact ID</td>
      <td style="text-align:left"></td>
    </tr>
  </tbody>
</table>

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}