# Getting started with the CLI

* [ Getting started with the CLI](/hc/en-us/articles/360003812458-Getting-started-with-the-CLI)
* [ How to find vulnerabilities using your CLI](/hc/en-us/articles/360005743037-How-to-find-vulnerabilities-using-your-CLI)
* [ CLI reference](/hc/en-us/articles/360003812578-CLI-reference)
* [ Set severity thresholds for CLI tests](/hc/en-us/articles/360003851337-Set-severity-thresholds-for-CLI-tests)
* [ Using a .snyk file in a different directory from a manifest file](/hc/en-us/articles/360014760717-Using-a-snyk-file-in-a-different-directory-from-a-manifest-file)

##  Getting started with the CLI

The Snyk Command Line Interface \(CLI\) helps you find and fix known vulnerabilities in your dependencies, both manually and as part of your Continuous Integration \(CI\) build system.

See [Language Support](https://support.snyk.io/hc/en-us/articles/360000911957-Language-support) for details about package managers and languages that the CLI supports. 

### Install and authenticate the Snyk CLI

1. Install the Snyk CLI using npm, Homebrew, Scoop or a manual installer from Snyk’s GitHub. See [Install the Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812538-Install-the-Snyk-CLI). 
2. To associate your Snyk account with the CLI, authenticate your account, either through your browser or using your API token. See [Authenticate the CLI with your account](https://support.snyk.io/hc/en-us/articles/360004008258-Authenticate-the-CLI-with-your-account).

### Build your project

Before testing for vulnerabilities you must build your project, unless it has one of the following lock files:

* `package-lock.json` 
* `yarn.lock`  
* `Gemfile.lock`  
* `paket.lock`

Examples of how to build/install your project:

`npm install`  
`mvn install`  
`gradle build`  
`dotnet restore`  
`dep ensure`

### Test for vulnerabilities

Once installed and authenticated, change directory into a folder containing a supported package manifest file \(package.json, pom.xml, composer.lock, etc.\) and run:

`cd /my/project/`

`snyk test`

All vulnerabilities identified are listed, including their path, and remediation guidance. To monitor your project regularly and receive ongoing notifications when new vulnerabilities are introduced, see [Monitor your projects at regular intervals](/hc/articles/360003851297#UUID-0de07f93-0f4e-3665-7f4b-466fff3b327a).

Snyk uses monitoring to regularly test your code and notify you when new vulnerabilities are introduced. To set up your project to be monitored:

1. Open a terminal from your project directory and run `snyk monitor`.

   This command takes a snapshot of your current dependencies so Snyk can regularly scan your code, and notify you about newly disclosed vulnerabilities as they are introduced, or when a previously unavailable patch or upgrade path is created.

2. Log in to the Snyk app and navigate to the [Projects page](https://app.snyk.io/projects) to see the latest snapshot and scan results.

