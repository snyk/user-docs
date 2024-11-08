# Snyk language support details

&#x20;This page is being built. For languages not yet described on this page, see the pages on the language for the details of what is supported.

## Apex

**Snyk Code support for Apex:**

**Frameworks and libraries supported**:  Apex Standard Library - Comprehensive\
**Import your app through SCM**: Available\
**Test or monitor your app through CLI and IDE**: Available\
**Features**:\
Support for Interfile analysis\
Support for `.trigger, .tgr` and `.cls` files\
Custom rules\
Reports\
Interfile analysis

## Bazel

Refer to the [Bazel pages](bazel/) for complete information on Snyk for Bazel.\


## C/C++

<table><thead><tr><th>Snyk Code support for C/C++</th><th>Snyk Open Source support for C/C++</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported:</strong> argparse parser - Comprehensive<br>Asio Library - Comprehensive<br>Boost Library - Partial<br>Botan LIbrary - Comprehensive<br>C Standard Library - Comprehensive<br>C++ Standard Library - Comprehensive<br>Curl library - Comprehensive<br>fstream framework - Comprehensive HTTPlib framework - Comprehensive JsonCpp library - Comprehensive<br>liboai framework - Comprehensive<br>libpq library - Comprehensive<br>libpqxx framework - Comprehensive libsodium library - Comprehensive LibTomCrypt framework - Comprehensive libxml2 framework - Comprehensive<br>MySQL framework - Comprehensive OpenSSL framework- Comprehensive<br>POSIX LIbrary - Comprehensive<br>pugixml library - Comprehensive<br>SQLite library - Comprehensive<br>WinHTTP framework - Comprehensive<br>Xerces libraries - Comprehensive</td><td><strong>Package manager</strong>: NA<br><strong>Package registry</strong>: No single registry, multiple sources <br><strong>Import your app through SCM</strong>: NA<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:generic</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:generic</code><br><strong>Features</strong>:<br>License scanning<br>Reports<br><strong>Package manager versions</strong>: NA</td><td></td></tr><tr><td><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>: Reports<br><strong>Operating systems:</strong> Linux, Windows (limited)<br><strong>Embedded systems:</strong> Linux<br><strong>IDE:</strong> No additional options are required. The Snyk plugin has views within the IDE for displaying results.<br><strong>Feature:</strong> Interfile analysis</td><td></td><td></td></tr></tbody></table>

## Dart and Flutter

**Snyk Open Source Support for Dart and Flutter:**

**Package manager**: Pub\
**Package registry**: [pub.dev](https://pub.dev/)\
**Test your app's SBOM**: Available, `pkg:pub`\
**Test your app's packages**: Available, `pkg:pub`

## Elixir

**Snyk Open Source Support for Elixir:**

**Package manager**: [Mix](https://hexdocs.pm/mix/Mix.html)/[Hex](https://hex.pm)\
**Package registry:**  [hex.pm](https://hex.pm/)\
**Import your app through SCM**: NA\
**Test or monitor your app through CLI and IDE**: Available\
**Test your app's SBOM**: Available, `pkg:hex`\
**Test your app's packages**: Available, `pkg:hex`\
**Feature:** Reports

## Go

<table><thead><tr><th>Snyk Code support for Go</th><th>Snyk Open Source support for Go</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported</strong>: <br>Azure/azure-sdk-for-go/sdk/ai/azopenai - Comprehensive<br>gage-technologies/mistral-go - Comprehensive<br>Gin - Partial<br>Go Standard Library - Comprehensive<br>google/generative-ai-go/genai - Comprehensive<br>GORM library - Partial<br>labstack/echo - Partial<br>sashabaranov/go-openai - Comprehensive<br>spf13/pflag - Comprehensive</td><td><strong>Package managers</strong>: Go Modules, dep<br><strong>Package registry</strong>: No single registry, multiple sources<br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:golang</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:golang</code><br><strong>Features</strong>:<br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>Go versions</strong>: Versions up to go1.16<br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Custom rules<br>Interfile analysis</td><td></td><td></td></tr></tbody></table>

## Java and Kotlin

<table><thead><tr><th>Snyk Code support for Java and Kotlin</th><th>Snyk Open Source support for Java and Kotlin</th><th data-hidden></th></tr></thead><tbody><tr><td><p><strong>Frameworks and libraries supported  for Java:</strong> <br>Amazon AWS SDK - Comprehensive<br>Android Standard Library - Partial<br>Apache Commons - Comprehensive<br>Apache Tomcat - Partial<br>Apache XML - Comprehensive<br>apache.mahou - Comprehensive<br>bouncycastle - Comprehensive<br>com.azure.ai.openai - Comprehensive<br>com.google.ai.client.generativeai - Comprehensive<br>com.google.cloud.vertexai.generativeai - Comprehensive<br>com.google.re2j - Comprehensive<br>com.google.gwt - Partial<br>Dropwizard - Comprehensive<br>elasticsearch - Partial<br>FasterXML Jackson - Comprehensive<br>Google Guava - Comprehensive<br>hibernate - Comprehensive<br>http4k - Comprehensive<br>io.jsonwebtoken - Comprehensive<br>Jakarta EE - Partial<br>Jakarta XML Services - Partial<br>Java EE - Partial<br>Java Servlet - Comprehensive<br>Java Servlet (javax) - Comprehensive<br>Java Server Pages - Partial<br>Java Standard Edition - Comprehensive<br>javalin - Partial<br>jooq - Comprehensive<br>Kyro - Comprehensive<br>Micronaut - Comprehensive<br>mongo-java-driver - Comprehensive<br>Netty - Comprehensive<br>okhttp3 - Comprehensive<br>org.apache.hc.client5 - Comprehensive<br>org.apache.http.client - Comprehensive<br>org.apache.sling - Partial<br>org.apache.tools.zip - Comprehensive<br>org.codehaus.plexus - Comprehensive<br>org.dom4j.io - Comprehensive<br>Playframework - Comprehensive<br>rxhttp - Comprehensive<br>Seam logger - Comprehensive<br>SnakeYaml - Comprehensive<br>Spongycastle - Comprehensive<br>Spring boot - Partial<br>Spring Web, MVC and JDBC - Comprehensive<br>Struts - Partial<br>Vaadin - Comprehensive<br>XStream - Comprehensive</p><p><strong>Frameworks and libraries supported for Kotlin:</strong> <br><strong>Note:</strong> Kotlin is fully interoperable with Java, as such the Java frameworks mentioned above (where applicable) are also supported for Kotlin.<br>Android Standard Library - Partial<br>com.aallam.openai - Comprehensive<br>com.expediagroup.graphql.server - Comprehensive<br>Javalin - Partial<br>Ktor - Comprehensive<br>Kotlin Standard Library - Comprehensive<br>khttp - Comprehensive</p></td><td><strong>Build tools</strong>: Maven, Gradle<br><strong>Build tool versions</strong>:<br>Maven: <code>3.*.</code> For more information, see the <a href="https://github.com/snyk/snyk-mvn-plugin#support">Snyk Maven plugin readme</a>.<br>Gradle: <code>4.*</code>, <code>5.*</code>, <code>6.*</code>, <code>7.*, 8.*.</code> For more information, see the <a href="https://github.com/snyk/snyk-gradle-plugin#support">Snyk Gradle plugin readme</a>.<br><strong>Package registry</strong>: <a href="https://maven.org/">maven.org</a><br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:maven</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:maven</code><br><strong>Features</strong>:<br>Fix PRs (Maven)<br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>Java versions</strong>: Versions up to Java SE 17<br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Report<br>Custom rules<br>Interfile analysis</td><td></td><td></td></tr><tr><td><strong>Kotlin Mobile Apps</strong>. Android is partially supported.<br><strong>Kotlin Web Applications</strong> are not supported.<br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Interfile analysis is supported for Kotlin<br>Android is partially supported</td><td></td><td></td></tr></tbody></table>

## JavaScript

<table><thead><tr><th>Snyk Code support for JavaScript</th><th>Snyk Open Source support for JavaScript</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported</strong>:<br>@Google Drive/generative-ai - Comprehensive<br>@anthropic-ai/sdk - Comprehensive<br>@huggingface/inference - Comprehensive<br>@mistralai/mistralai - Comprehensive<br>axios - Comprehensive<br>Angular - Partial<br>apollo-server - Partial<br>bcrypt-nodejs - Comprehensive<br>cross-spawn - Comprehensive<br>crypto-js - Comprehensive<br>date-fns - Comprehensive<br>dayjs - Comprehensive<br>dompurify  - Comprehensive<br>electron - Partial<br>ejs - Partial<br>execa - Comprehensive<br>express - Comprehensive<br>express-graphql - Partial<br>express-jwt - Partial<br>fs - Comprehensive<br>fs-extra - Comprehensive<br>fs-plus - Comprehensive<br>graceful-fs - Comprehensive<br>graphql-js - Partial<br>jQuery - Comprehensive<br>js-yaml - Comprehensive<br>jzip - Comprehensive<br>koa - Comprehensive<br>koa-graphql - Comprehensive<br>libxml - Comprehensive<br>libxmljs - Comprehensive<br>lodash - Comprehensive<br>luxon - Comprehensive<br>minimongo - Comprehensive<br>minimist - Comprehensive<br>mongodb - Comprehensive<br>Mongoose - Comprehensive<br>mercurius - Partial<br>Nestjs - Partial<br>Node Crypto - Comprehensive<br>node-buffer - Partial<br>node-cmd - Comprehensive<br>Node Crypto - Comprehensive<br>node-dir - Comprehensive<br>node-forge - Comprehensive<br>node-pty - Comprehensive<br>node-serialize - Comprehensive<br>octokit - Comprehensive<br>openai - Comprehensive<br>pg - Comprehensive<br>pg-promise - Comprehensive<br>React - Partial<br>request-promise - Comprehensive<br>restler - Partial<br>rimraf - Comprehensive<br>sanitize-html - Comprehensive<br>shelljs - Comprehensive<br>Stanford JS Crypto - Comprehensive<br>superagent - Comprehensive<br>tar-stream - Comprehensive<br>unirest - Comprehensive<br>unzip - Comprehensive<br>underscore - Comprehensive<br>url - Comprehensive<br>vm - Comprehensive<br>webstomp-client - Partial<br>WebCryptoAPI - Comprehensive<br>xpath - Comprehensive<br>yargs - Comprehensive</td><td><strong>Package managers</strong>: npm, pnpm, Yarn<br><strong>Package manager versions</strong>:<br>npm: <code>Lockfile 1</code>, <code>Lockfile 2</code>, <code>Lockfile 3, 7.*</code><br>pnpm: <code>pnpm 7</code>, <code>pnpm 8</code>, <code>pnpm 9</code><br>Yarn: <code>Yarn 1</code>, <code>Yarn 2</code>, <code>Yarn 3</code><br><strong>Package registry</strong>: <a href="https://www.npmjs.org">npmjs.org</a><br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:npm</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:npm</code><br><strong>Features</strong>: <br>Fix PRs<br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Custom rules<br>Interfile analysis</td><td></td><td></td></tr></tbody></table>

## .NET

<table><thead><tr><th>Snyk Code support for .NET</th><th>Snyk Open Source support for .NET</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported</strong>:<br>.NET 6 - Comprehensive<br>.NET Core - Comprehensive<br>.NET Framework 4.6-4.8.x - Comprehensive<br>Anthropic.SDK - Comprehensive<br>ASP.NET 6.x (C# only) - Comprehensive<br>Azure.AI.OpenAI - Comprehensive<br>Dapper - Comprehensive<br>fastJSON - Comprehensive<br>Google_GenerativeAI - Comprehensive<br>Microsoft.CodeAnalysis.VisualBasic - Comprehensive<br>Mistral.SDK - Comprehensive<br>System.CodeDom.Compiler - Comprehensive<br>Windows Forms - Partial</td><td><strong>Package managers</strong>: NuGet, Paket<br><strong>Package registry</strong>: <a href="https://www.nuget.org/">nuget.org</a><br><strong>Import your app through SCM</strong>: NuGet<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:nuget</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:nuget</code><br><strong>Features</strong>:<br>Fix PRs (NuGet)<br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Custom rules<br>Interfile analysis</td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table>

## PHP

<table><thead><tr><th>Snyk Code support for PHP</th><th>Snyk Open Source support for PHP</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported</strong>: <br>Laravel - Partial<br>llphant - Comprehensive<br>openai-php/client - Comprehensive<br>orhanerday/open-ai - Comprehensive<br>Pclzip - Comprehensive<br>Symfony - Partial<br>theodo-group/llphant - Comprehensive</td><td><strong>Package manager</strong>: Composer<br><strong>Package registry</strong>: <a href="https://packagist.org/">packagist.org</a><br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:composer</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:composer</code><br><strong>Features</strong>:<br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>PHP versions</strong>: Versions 5.2 up to 8.0<br><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Custom rules<br>Interfile analysis</td><td></td><td></td></tr></tbody></table>

## Python

<table><thead><tr><th>Snyk Code support for Python</th><th>Snyk Open Source support for Python</th><th data-hidden></th></tr></thead><tbody><tr><td><strong>Frameworks and libraries supported</strong>:<br>AioHTTP - Comprehensive<br>aiopg - Comprehensive<br>argparse - Comprehensive<br>anthropic - Comprehensive<br>bottle - Comprehensive<br>CherryPy - Comprehensive<br>Django - Comprehensive<br>defusedxml - Comprehensive<br>fastapi - Partial<br>flask - Comprehensive<br>flask_pymongo - Comprehensive<br>google.cloud.bigquery - Comprehensive<br>google_generativeai - Comprehensive<br>huggingface_hub - Comprehensive<br>httpx - Comprehensive<br>ldap3 - Comprehensive<br>libxml - Comprehensive<br>lxml - Comprehensive<br>mistralai - Comprehensive<br>mongoengine - Comprehensive<br>openai - Comprehensive<br>pandas - Partial<br>paramiko - Comprehensive<br>peewee - Comprehensive<br>pickle - Comprehensive<br>pilyaml - Comprehensive<br>pyca/cryptography - Comprehensive<br>pymongo - Comprehensive<br>pymssql - Comprehensive<br>pyramid - Comprehensive<br>psycopg - Comprehensive<br>python-ldap - Comprehensive<br>Python Standard Library - Comprehensive<br>requests - Comprehensive<br>sqlite3 (or pysqlite2) - Comprehensive<br>sqlalchemy - Comprehensive<br>turboGears - Comprehensive<br>urllib - Comprehensive<br>werkzeug - Comprehensive</td><td><strong>Package manager</strong>: Pip, Poetry, pipenv, setup.py<br><strong>Package manager versions</strong>: Suitable with <code>Python 2 -> 2.7.16</code>, and <code>Python 3 -> 3.7.4</code>.<br><strong>Package registry</strong>: <a href="https://pypi.org">pypi.org</a><br><strong>Import your app through SCM</strong>: Available for Pip, pipenv, and Poetry<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Test your app's SBOM</strong>: Available, <code>pkg:pypi</code><br><strong>Test your app's packages</strong>: Available, <code>pkg:pypi</code><br><strong>Features</strong>: <br>Fix PRs <br>License scanning<br>Reports</td><td></td></tr><tr><td><strong>Import your app through SCM</strong>: Available<br><strong>Test or monitor your app through CLI and IDE</strong>: Available<br><strong>Features</strong>:<br>Reports<br>Customer rules<br>Interfile analysis</td><td></td><td></td></tr></tbody></table>

## Rust

**Snyk Open Source support for Rust:**

**Package manager**: Cargo\
**Package registry**: [crates.io](https://crates.io/)\
**Test your app's SBOM**: Available, `pkg:cargo`\
**Test your app's packages**: Available, `pkg:cargo`\
See the Rust page for information about using the API to test Rust applications.

## Scala

<table><thead><tr><th>Snyk Code support for Scala</th><th>Snyk Open Source support for Sca;a</th><th data-hidden></th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table>

## Swift and Objective-C

<table><thead><tr><th width="353">Snyk Code support for Swift</th><th>Snyk Open Source support for Swift and Objective-C</th><th data-hidden></th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table>

## TypeScript

<table><thead><tr><th>Snyk Code support for TypeScript</th><th>Snyk Open Source support for Typescript</th><th data-hidden></th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table>

## VB.NET

**Snyk Code support for VB NET:**

**Frameworks and libraries supported**:\
.NET Core - Comprehensive\
.NET Framework 4.6-4.8.x - Comprehensive\
Anthropic.SDK - Comprehensive\
Azure.AI.OpenAI - Comprehensive\
Google\_GenerativeAI - Comprehensive\
Mistral.SDK - Comprehensive

**Import your app through SCM**: Available\
**Test or monitor your app through CLI and IDE**: Available\
**Feature**: Reports\
**Framework version**: Version 7\
**Feature:** Interfile analysis
