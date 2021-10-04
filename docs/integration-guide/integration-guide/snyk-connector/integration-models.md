# Integration Models

## What integration models are available?

A Snyk Connector can implement one of the following approaches. We will discuss each approach in the next section.

* Periodic Polling via Snyk API. This is the preferred method
* Ad Hoc data export via Snyk CLI.

## Which integration model should I use?

The recommended integration method is to use periodic polling via Snyk API. Below we captured some guidance to help you understand the difference.

Use Periodic Polling via Snyk API when:

* You have a backend system that can periodically make API calls.
* Customer is already importing projects into the Snyk UI \(most common\).
* You prefer customer not having to do any integration work to support extracting data out of Snyk \(best experience\)

Use Ad Hoc data export via CLI when:

* You do not have a backend system that can periodically make API calls
* Customer is not importing projects into Snyk UI \(not common\)
* You want to focus only on exporting data from a CI pipeline into your platform
* You don’t mind having the customer do a little integration work to import Snyk scan results into your platform.

### Periodic Polling via the Snyk API

Snyk users create integrations to various sources of their development toolchain. This includes Source Control Management \(SCM\) systems like Github, GitLab, and Bitbucket. Continuous Integration \(CI\) tools like Jenkins, Circle CI, etc. Container registries to scan their container images like Docker Hub, ECR, ACR, and GCR.

Once these integrations are established, Snyk users import projects from these various sources. A _Snyk project_ is a package that is tracked \(See Data Model below for complete details\). For example, if you import a GitHub project, Snyk will import the package manager file and track issues \(license and vulnerabilities\) associated with the application dependencies. Snyk regularly scans the projects once per day.

Snyk’s API allows you to fetch these **existing** Snyk scan results from any of the projects irrespective of how the customer imported them. Since these scan results already exist, retrieving this data is fairly quick. Snyk’s API generally allows to _trigger_ a scan, however, this functionality should **not** be used to build a connector.

#### Steps for **Periodic Polling**

The Snyk connector requires an `API_TOKEN` to be issued with each API call.

* Collect customer’s `API_TOKEN`. See section on _Snyk API Token_ below for details.
* Periodically run the following: **\(recommended is daily, but no less than weekly\)**
  * Call the groups API to retrieve list of groups
  * For every group, call organization API to retrieve list of organizations
  * For every organization, call the projects API
  * For every project, call the issues API
  * Map the metadata returned for each issue into your own data structure using the data mapping below

### Ad Hoc data export via the Snyk CLI

The Snyk CLI can be used to perform an ad hoc export of data using the `—-json` switch on the command line. The results of the Snyk CLI can be saved to a local file and transformed and imported.

#### Steps for Synk CLI

* Customer runs `snyk test --json > scan.json` inside an application folder
* Customer runs a conversion script from the snyk json format to an acceptable format for your platform: `convert_from_snyk_to_my_platform scan.json > converted_output`
* Customer uploads `converted_output` to your platform via your web experience or via APIs based on how your platform works

