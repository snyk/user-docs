# Snyk for Swift and Objective-C \(CocoaPods\)

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your CocoaPods projects:

### Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| ![i\_icon\_cocoapods.png](https://support.snyk.io/hc/article_attachments/360007259038/uuid-6de05da9-de7e-11cc-4316-8459517aaf57-en.png) | Cocoapods | ✔︎ | ✔︎  | ✔︎  |  |  |

#### Snyk CLI tool for CocoaPods projects

**CLI parameters for Swift and Objective-C**

When working with Swift and Objective-C projects from our CLI, you can prevent testing any lockfiles that are out-of-sync, as follows:

<table>
  <thead>
    <tr>
      <th style="text-align:left">Option</th>
      <th style="text-align:left">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left"><code>--strict-out-of-sync=</code>
      </td>
      <td style="text-align:left">
        <p>Prevent testing out-of-sync lockfiles.</p>
        <p>Defaults to <b>true</b>.</p>
      </td>
    </tr>
  </tbody>
</table>

#### Git services for CocoaPods projects

Swift and Objective-C projects managed by CocoaPods can be imported from any of the Git repositories we support. In order to test your projects, we analyze your Podfile and Podfile.lock files.

