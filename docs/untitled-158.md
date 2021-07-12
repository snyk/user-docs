# Snyk for Swift and Objective-C \(CocoaPods\)

* [ Language support summary](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360020352437-Language-support-summary/README.md)
* [ Snyk for JavaScript](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004712477-Snyk-for-JavaScript/README.md)
* [ Snyk for Java \(Gradle, Maven\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817357-Snyk-for-Java-Gradle-Maven-/README.md)
* [ Snyk for .NET](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004519138-Snyk-for-NET/README.md)
* [ Snyk for Python](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004699377-Snyk-for-Python/README.md)
* [ Snyk for Golang](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817417-Snyk-for-Golang/README.md)
* [ Snyk for Swift and Objective-C \(CocoaPods\)](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-/README.md)
* [ Snyk for Scala](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003781318-Snyk-for-Scala/README.md)
* [ Snyk for Ruby](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003781298-Snyk-for-Ruby/README.md)
* [ Snyk for PHP](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/articles/360003817397-Snyk-for-PHP/README.md)

  [See more](https://github.com/snyk/user-docs/tree/58f91d848e16ddf2ffcca3711d6b8852412be402/hc/en-us/sections/360001087857-Language-package-manager-support/README.md)

## Snyk for Swift and Objective-C \(CocoaPods\)

Snyk offers security scanning to test your projects for vulnerabilities, both through your CLI and through different integrations from our UI.

The following describes how to use Snyk to scan your CocoaPods projects:

### Note

Features might not be available, depending on your subscription plan.

|  | Package managers/Features | CLI support | Git support | License scanning | Remediation | Runtime monitoring |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|  | Cocoapods | ✔︎ | ✔︎ | ✔︎ |  |  |

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
      <td style="text-align:left"><code>--strict-out-of-sync=&lt;true|false&gt;</code>
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

