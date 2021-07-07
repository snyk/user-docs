# What counts as a test?

##  What counts as a test?

### For billing purposes what counts as a test, and counts towards the test limit?

Some Snyk accounts have a limit of the number of tests that they can run per month, so it can be important to know what counts as a test in terms of licensing.

Tests can be run in multiple ways:

**SCM integrations**  
The SCM integrations will run a test whenever we get changes in pull request/merge requests or inspect dependency changes you push.

**Recurring tests**  
Snyk will  periodically check if your repo is affected by newly disclosed vulnerabilities \(daily by default\). This is a setting that can be controlled in **Settings** &gt; **usage** or in the project settings page.

**CLI**  
With our [CLI](https://snyk.io/docs/using-snyk), you control exactly when and why you run a test. A test is counted for each run of `snyk test` or `snyk monitor` 

Tests on private projects count towards your test limit, tests on Open Source projects do not count and are free.

**Web based tests**  
A test is run when you add a new project or click the re-test button.

**APIs**  
Tests are counted when calls are made to the [https://snyk.io/api/v1/test](https://snyk.io/api/v1/test) endpoint.

