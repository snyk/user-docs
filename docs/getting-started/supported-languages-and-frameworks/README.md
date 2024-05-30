# Supported languages and frameworks

This page provides an overview of supported languages and package managers across Snyk environments, including feature availability for open source and licensing (Snyk Open Source) and code analysis (Snyk Code).

## Open source and licensing (Snyk Open Source)

Here you can find information about the programming languages, fully supported package managers, and features for Snyk Open Source.

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

### [.NET](.net/)

**Package manager**: NuGet, Paket

**Import your app through SCM**: NuGet

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:nuget`

**Test your app's packages**: Available, `pkg:nuget`

**Features**:&#x20;

* Fix PRs (NuGet)
* License scanning
* Reports

**Package manager versions**: NA

### [C/C++](c-c++/)

**Package manager**: NA

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:generic`

**Test your app's packages**: Available, `pkg:generic`

**Features**:

* License scanning
* Reports

**Package manager versions**: NA

### [Dart and Flutter](dart-and-flutter.md)&#x20;

**Package manager**: Pub

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: NA

**Test your app's SBOM**: NA

**Test your app's packages**: Available, `pkg:pub`

**Features**: NA

**Package manager versions**: NA

### [Elixir](elixir.md)&#x20;

**Package manager**: Hex

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:hex`

**Test your app's packages**: Available, `pkg:hex`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

### [Go](go.md)&#x20;

**Package manager**: Go Modules, dep

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:golang`

**Test your app's packages**: Available, `pkg:golang`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

### [Java and Kotlin](java-and-kotlin/)&#x20;

**Package manager**: Maven, Gradle

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:maven`

**Test your app's packages**: Available, `pkg:maven`

**Features**:&#x20;

* Fix PRs (Maven)
* License scanning
* Reports

**Package manager versions**:&#x20;

* Maven
  * `3.*` For details, see the [Snyk Maven plugin readme](https://github.com/snyk/snyk-mvn-plugin#support).
* Gradle
  * `4.*`, `5.*`, `6.*`, `7.*`\
    For more information, see the [Snyk Gradle plugin readme](https://github.com/snyk/snyk-gradle-plugin#support).

### [JavaScript](javascript/)&#x20;

**Package manager**: npm, Yarn

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:npm`

**Test your app's packages**: Available, `pkg:npm`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**:&#x20;

* npm
  * `Lockfile 1, Lockfile 2, Lockfile 3, 7.*`\
    For details, see the [Snyk Javascript ](javascript/#npm)page.
* Yarn
  * `Yarn 1, Yarn 2, Yarn 3`. For more information, see the [Snyk Javascript ](javascript/#yarn)page.

### [PHP](php.md)

**Package manager**: Composer

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:composer`

**Test your app's packages**: Available, `pkg:composer`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

### [Python](python.md)&#x20;

**Package manager**: Pip, Poetry, pipenv, setup.py

**Import your app through SCM**: Available for Pip, pipenv and Poetry

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:pypi`

**Test your app's packages**: Available, `pkg:pypi`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**: Suitable with `Python 2 -> 2.7.16`, and `Python 3 -> 3.7.4`.

### [Ruby](ruby.md)&#x20;

**Package manager**: Bundler

**Import your app through SCM**: Available&#x20;

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:gem`

**Test your app's packages**: Available, `pkg:gem`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**: All Gemfile and Gemfile.lock are compatible with the [Snyk supported Ruby versions](ruby.md#supported-ruby-versions).

### [Rust](rust.md)

**Package manager**: Cargo

**Import your app through SCM**: NA&#x20;

**Test or monitor your app through CLI and IDE**: NA

**Test your app's SBOM**: Available, `pkg:cargo`

**Test your app's packages**: Available, `pkg:cargo`

**Features**: NA

**Package manager versions**: NA

### [Scala](scala.md)&#x20;

**Package manager**: sbt

**Import your app through SCM**: Available&#x20;

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:maven`

**Test your app's packages**: Available, `pkg:maven`

**Features**:&#x20;

* License scanning
* Reports

**Package manager versions**: NA

### [Swift and Objective-C](swift-and-objective-c.md)

**Package manager**: CocoaPods, Swift Package Manager

**Import your app through SCM**: Available for CocoaPods

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:swift`, `pkg:cocoapods`

**Test your app's packages**: Available, `pkg:swift`, `pkg:cocoapods`

**Features**:&#x20;

* License scanning (CocoaPods)
* Reports

**Package manager versions**: CocoaPods, Swift Package Manager, Swift v3.0 or higher.

## Code analysis (Snyk Code)

Here you can find information about the programming languages, fully supported frameworks, and features for Snyk Code.

### [Apex](apex.md)

**Frameworks**: APEX&#x20;

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules
* Interfile is supported
* `.trigger` and `.cls` files are supported

### [.Net](.net/)

**Frameworks**:&#x20;

* .NET
* ASP.NET
* .NET Core

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules

**Framework versions**:&#x20;

* .NET Framework 4.6 - 4.8.x
* ASP.NET 6.x
* .NET 6

### [C/C++](c-c++/)

**Frameworks**: C++ Standard Library

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**: Reports

### [Go](go.md)

**Frameworks**: GORM

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules

**Framework versions**: Versions up to go1.16

### [Java](java-and-kotlin/)

**Frameworks:**

* Apache Camel
* Apache Struts
* Spring MVC
* Spring JDBC
* Jakarta XML Services
* Dropwizard

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Report&#x20;
* Custom rules

**Framework versions**: Versions up to Java SE 17

### [JavaScript](javascript/)

**Frameworks**:&#x20;

* React
* Vue.js
* Express
* jQuery
* Angular

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules

**Framework versions**: Versions up to ECMAScript 2020

### [Kotlin](java-and-kotlin/)

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Interfile is supported
* Android is partially supported

### [PHP](php.md)

**Frameworks**:&#x20;

* Symfony
* Laravel
* PHP Standards

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports&#x20;
* Custom rules

**Framework versions**: Versions 5.2 up to 8.0

### [Python](python.md)

**Frameworks**:&#x20;

* Django
* Flask
* Jinja2
* PyYAML
* Requests
* urllib3

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Customer rules

**Framework versions**: Versions up to 3.8.x

### [Ruby](ruby.md)

**Frameworks**: Ruby On Rails

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules

### [Scala](scala.md)

**Frameworks**:&#x20;

* Play Framework
* Akka
* HTTP4S

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Interfile is supported

**Framework versions**: Version 2.x

### [Swift](swift-and-objective-c.md)

**Frameworks**:&#x20;

* AlamoFire
* Pathos
* SQLite
* CryptoKit

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports

**Framework versions**: Versions up to 5.7.x

### [TypeScript](typescript.md)

**Frameworks**: npm

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports
* Custom rules

**Framework versions**: Versions up to 4.2

### [VB.NET](vb.net.md)

**Frameworks**: NuGet

**Import your app through SCM**: Available

**Test or monitor your app through CLI and IDE**: Available

**Features**:&#x20;

* Reports

**Framework versions**: Version 7
