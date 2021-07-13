# Snyk for Ruby

* [ Language support summary](https://support.snyk.io/hc/en-us/articles/360020352437-Language-support-summary)
* [ Snyk for JavaScript](https://support.snyk.io/hc/en-us/articles/360004712477-Snyk-for-JavaScript)
* [ Snyk for Java \(Gradle, Maven\)](https://support.snyk.io/hc/en-us/articles/360003817357-Snyk-for-Java-Gradle-Maven-)
* [ Snyk for .NET](https://support.snyk.io/hc/en-us/articles/360004519138-Snyk-for-NET)
* [ Snyk for Python](https://support.snyk.io/hc/en-us/articles/360004699377-Snyk-for-Python)
* [ Snyk for Golang](https://support.snyk.io/hc/en-us/articles/360003817417-Snyk-for-Golang)
* [ Snyk for Swift and Objective-C \(CocoaPods\)](https://support.snyk.io/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-)
* [ Snyk for Scala](https://support.snyk.io/hc/en-us/articles/360003781318-Snyk-for-Scala)
* [ Snyk for Ruby]()
* [ Snyk for PHP](https://support.snyk.io/hc/en-us/articles/360003817397-Snyk-for-PHP)

 [See more](https://support.snyk.io/hc/en-us/sections/360001087857-Language-package-manager-support)

##  Snyk for Ruby

Snyk supports testing, monitoring and fixing Ruby projects in the CLI and Git integrations that have their dependencies managed by [Bundler](https://bundler.io/), and comparing the specific dependency versions against the [Ruby vulnerability database](https://snyk.io/vuln?type=rubygems). 

Snyk tests all Bundler groups, and currently it is not possible to exclude certain groups \(such as test or development groups\).

If your Gemfile needs access to private Gem sources please [see below.]()

The following manifest files are supported:

### Note

Snyk requires both files to be present in order to correctly test, monitor & fix Ruby projects.

#### Fixing vulnerabilities in your Ruby projects

Snyk can fix vulnerabilities by updating vulnerable gems, using bundle update, after modifying your Gemfile \(sticking to the rules you have specified there as far as possible\).

This means that in some scenarios we won’t be able to upgrade all dependencies to non-vulnerable versions. In this case, you should consider updating the rules in your Gemfile.

In future releases, we are planning to provide suggestions to make this easier.

