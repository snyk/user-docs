# Snyk runtime: deployment metrics

{% hint style="info" %}
This feature is deprecated.
{% endhint %}

The following table describes Snyk agent deployment metrics based on different scenarios when using Snyk runtime monitoring.

<table>
  <thead>
    <tr>
      <th style="text-align:left">Language</th>
      <th style="text-align:left">During run</th>
      <th style="text-align:left">On startup</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Node.js</td>
      <td style="text-align:left">
        <p>~20 ms delay for each called function per minute</p>
        <p>~12ms delay for each loaded JavaScript file that contains vulnerable functions:</p>
        <ul>
          <li>When the package is loaded for the first time</li>
          <li>Periodically, as newer snapshots are introduced</li>
        </ul>
      </td>
      <td style="text-align:left">Medium-sized project: ~120 ms</td>
    </tr>
    <tr>
      <td style="text-align:left">Java</td>
      <td style="text-align:left">Negligible - not measurable</td>
      <td style="text-align:left">
        <ul>
          <li>Small and medium projects: Negligible</li>
          <li>Large projects: ~800 ms</li>
        </ul>
      </td>
    </tr>
  </tbody>
</table>

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}