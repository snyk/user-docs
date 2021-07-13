# Snyk for Scala

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### Language & package manager support

* [ Language support summary](https://support.snyk.io/hc/en-us/articles/360020352437-Language-support-summary)
* [ Snyk for JavaScript](https://support.snyk.io/hc/en-us/articles/360004712477-Snyk-for-JavaScript)
* [ Snyk for Java \(Gradle, Maven\)](https://support.snyk.io/hc/en-us/articles/360003817357-Snyk-for-Java-Gradle-Maven-)
* [ Snyk for .NET](https://support.snyk.io/hc/en-us/articles/360004519138-Snyk-for-NET)
* [ Snyk for Python](https://support.snyk.io/hc/en-us/articles/360004699377-Snyk-for-Python)
* [ Snyk for Golang](https://support.snyk.io/hc/en-us/articles/360003817417-Snyk-for-Golang)
* [ Snyk for Swift and Objective-C \(CocoaPods\)](https://support.snyk.io/hc/en-us/articles/360004701658-Snyk-for-Swift-and-Objective-C-CocoaPods-)
* [ Snyk for Scala]()
* [ Snyk for Ruby](https://support.snyk.io/hc/en-us/articles/360003781298-Snyk-for-Ruby)
* [ Snyk for PHP](https://support.snyk.io/hc/en-us/articles/360003817397-Snyk-for-PHP)

 [See more](https://support.snyk.io/hc/en-us/sections/360001087857-Language-package-manager-support)

1.  [Docs Library \| Snyk](https://support.snyk.io/hc/en-us)
2.  [Snyk Open Source](https://support.snyk.io/hc/en-us/categories/360003049458-Snyk-Open-Source)
3.  [Language & package manager support](https://support.snyk.io/hc/en-us/sections/360001087857-Language-package-manager-support)

##  Snyk for Scala

Snyk supports testing Scala projects that have their dependencies managed by [sbt](https://www.scala-sbt.org/). Support is available via the Snyk UI and CLI.

### Note

To use the Snyk CLI with versions of sbt 1.2 and older, you will need to first [install the sbt-dependency-graph plugin](https://support.snyk.io/hc/en-us/articles/360004167317).

#### Testing Scala projects: how it works

We scan Scala projects by examining your build.sbt to compare the specific versions of every direct and deep dependency in your project against our [Maven vulnerability database](https://snyk.io/vuln?type=maven).

* 
