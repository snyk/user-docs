# TeamCity configuration parameters

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Parameters</b>
      </th>
      <th style="text-align:left"><b>Description and values</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><b>Snyk settings</b>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Severity threshold</td>
      <td style="text-align:left">
        <p>Default: low</p>
        <p>For the first vulnerability found in your build with the threshold as
          configured, the build fails.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Monitor project on build</td>
      <td style="text-align:left">
        <p>Default: ON</p>
        <p>Snyk runs the snyk monitor command during the build, sending a project
          snapshot to the Snyk app and continuing to monitor the project for vulnerabilities
          even after this build.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">File</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>If the manifest file is not on the root level, enter the relative path
          to that file.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Organization</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>The ID of the Snyk organization to which this project should be associated
          when imported to the UI.</p>
        <p>Copy the Organization ID from the Snyk UI in the Settings area.</p>
        <p>
          <img src="../../.gitbook/assets/uuid-dfede20b-acb5-fc08-8d1d-59e8476240a5-en.png"
          alt="image6.png" />
        </p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Project name</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>Enter any unique name for this project to recognize it when viewing from
          the Snyk UI.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Additional parameters</td>
      <td style="text-align:left">
        <p>Optional.</p>
        <p>Enter additional CLI arguments as necessary. See our CLI documentation
          and cheat sheet for additional information.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left"><b>Snyk tool settings</b>
      </td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Snyk API token</td>
      <td style="text-align:left">
        <p>From the Settings area in the Snyk UI, copy the Org or Personal API token
          or create a service account. This is the token used to authenticate your
          Snyk account when connecting to TeamCity.</p>
        <p>
          <img src="../../.gitbook/assets/uuid-c27d25fc-00a7-f0f4-261c-d0d9f8653d1d-en.png"
          alt="image7.png" />
        </p>
        <p>
          <img src="../../.gitbook/assets/uuid-be0e9602-023b-99a4-f08c-eded5ea77dac-en.png"
          alt="image8.png" />
        </p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Snyk version</td>
      <td style="text-align:left">
        <p>Default: the most recent version</p>
        <p>Select the plugin version to be used in your build if you would like an
          older Snyk CLI version to support the plugin.</p>
        <p>We recommend configuring automatic upgrades and using the most recent
          version.</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">Use custom build tool path</td>
      <td style="text-align:left">
        <p>Specify which tool instance in your local environment is to be used for
          this build by Snyk.</p>
        <p>Otherwise Snyk auto-detects the tool and locates it in your environment
          based on project type.</p>
      </td>
    </tr>
  </tbody>
</table>

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}