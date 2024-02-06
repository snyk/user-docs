# Supported languages, frameworks, and feature availability overview

Get an overview of supported languages and package managers across Snyk environments, including feature availability for open source and licensing (Snyk Open Source) and code analysis (Snyk Code).

## Open source and licensing (Snyk Open Source)

Below, you can find information about the programming languages, fully supported package managers, and features for Snyk Open Source.

{% hint style="info" %}
Before scanning your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Which Projects must be built before testing with CLI?](https://support.snyk.io/hc/en-us/articles/360015552617-Which-projects-must-be-built-before-testing-with-CLI-)
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

### [C/C++](c-c++.md)

**Package manager**: NA

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:generic`

**Test your app's packages**: Available, `pkg:generic`

**Features**:

* License scanning
* Reports

**Package manager versions**: NA

### [Dart and Flutter](../../scan-using-snyk/supported-languages-and-frameworks/dart-and-flutter.md)&#x20;

**Package manager**: Pub

**Import your app through SCM**: NA

**Test or monitor your app through CLI and IDE**: NA

**Test your app's SBOM**: NA

**Test your app's packages**: Available, `pkg:pub`

**Features**: NA

**Package manager versions**: NA

### [Elixir](../../scan-using-snyk/supported-languages-and-frameworks/elixir.md)&#x20;

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

### [Java and Kotlin](java-and-kotlin.md)&#x20;

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

### [JavaScript](javascript.md)&#x20;

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
    For details, see the [Snyk Javascript ](javascript.md#npm)page.
* Yarn
  * `Yarn 1, Yarn 2, Yarn 3`. For more information, see the [Snyk Javascript ](javascript.md#yarn)page.

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

### [Rust](../../scan-using-snyk/supported-languages-and-frameworks/rust.md)

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

The following table lists the programming languages and fully supported frameworks and features for Snyk Code.

<table data-full-width="false"><thead><tr><th width="232">Language and framework</th><th width="131" align="center">Git integration (SCM)</th><th width="207" align="center">Snyk CLI, plugins (IDE), CI/CD</th><th>Features</th></tr></thead><tbody><tr><td><strong>Apex</strong></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li><li>Custom rules</li></ul></td></tr><tr><td><p><strong>C#</strong></p><ul><li><strong>.NET</strong></li><li><strong>ASP.NET</strong></li><li><strong>.NET Core</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="c-c++.md"><strong>C/C++</strong></a></p><ul><li><strong>C++ Standard Library</strong></li><li><strong>POSIX</strong></li><li><strong>Win32</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td>Reports</td></tr><tr><td><strong>Go</strong></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="java-and-kotlin.md"><strong>Java</strong></a></p><ul><li><strong>Apache Camel</strong></li><li><strong>Apache Struts</strong></li><li><strong>Spring MVC</strong></li><li><strong>Spring JDBC</strong></li><li><strong>Jakarta XML Services</strong></li><li><strong>Dropwizard</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="javascript.md"><strong>JavaScript</strong></a></p><ul><li><strong>React</strong></li><li><strong>Vue.js</strong></li><li><strong>Express</strong></li><li><strong>JQuery</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><a href="java-and-kotlin.md"><strong>Kotlin</strong></a></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td>Reports</td></tr><tr><td><p><strong>PHP</strong></p><ul><li><strong>Symfony</strong></li><li><strong>Laravel</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="python.md"><strong>Python</strong></a></p><ul><li><strong>Django</strong></li><li><strong>Flask</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="ruby.md"><strong>Ruby</strong></a></p><ul><li><strong>Ruby On Rails</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><p><a href="scala.md"><strong>Scala</strong></a></p><ul><li><strong>Play</strong></li><li><strong>Akka</strong></li><li><strong>HTTP4S</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td>Reports</td></tr><tr><td><p><a href="swift-and-objective-c.md"><strong>Swift</strong></a></p><ul><li><strong>AlamoFire</strong></li><li><strong>Pathos</strong></li><li><strong>SQLite</strong></li><li><strong>CryptoKit</strong></li></ul></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td>Reports</td></tr><tr><td><a href="../../scan-using-snyk/supported-languages-and-frameworks/typescript.md"><strong>TypeScript</strong></a></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td><ul><li>Reports</li></ul><ul><li>Custom rules</li></ul></td></tr><tr><td><a href="../../scan-using-snyk/supported-languages-and-frameworks/vb.net.md"><strong>VB.NET</strong></a></td><td align="center"><strong>✔︎</strong></td><td align="center"><strong>✔︎</strong></td><td>Reports</td></tr></tbody></table>
