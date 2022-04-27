# FAQ

## **How do I scan multi-module Java projects?** <a href="#de2589f9-2a2d-47c5-acc4-e6e13b163b71" id="de2589f9-2a2d-47c5-acc4-e6e13b163b71"></a>

Some Java projects may be broken up into multiple `pom.xml` files. To scan all of these at once, use the `—-all-projects` option of the CLI.

## Does Snyk offer support for local modules? <a href="#25465972-e24e-4b43-ab73-c69b27535427" id="25465972-e24e-4b43-ab73-c69b27535427"></a>

Many customers reference proprietary modules the company builds, which are typically hosted in a local repository, for example, JFrog Artifactory. As the SSL certificate for these repositories is often not valid, Snyk returns an error when trying to connect to the repository. In such cases you can use the `--insecure` option of the CLI to ignore invalid certificates. As this is not good practice, but sometimes a necessity, Snyk offers this feature only as an explicit opt-in by the user.

## **Does Snyk have a test quota?** <a href="#39572c2f-bb72-4b12-a6cd-8fd1a2fc1650" id="39572c2f-bb72-4b12-a6cd-8fd1a2fc1650"></a>

Free Snyk accounts have a 200 tests per month test quota. All paid Snyk plans have unlimited scanning.

## **Are the Snyk test quotas enforced?** <a href="#83b6c566-7413-4ff8-adc1-fb27d29f5954" id="83b6c566-7413-4ff8-adc1-fb27d29f5954"></a>

No, currently the test quota is not enforced. The Snyk CLI only prints a warning similar to the following when the quota is reached:

{% hint style="danger" %}
_"You have reached your monthly limit of 200 private tests for your \[ORG] org. To learn more about our plans and increase your tests limit visit_ [_https://snyk.io/plans_](https://snyk.io/plans)_."_
{% endhint %}

However, the `--json` output does not currently include any indication of passing the quota.

## How do I include `dev dependencies` in the scan results? <a href="#f604610f-b039-43ee-8dfe-5609f0cabfd1" id="f604610f-b039-43ee-8dfe-5609f0cabfd1"></a>

Some package managers have the notion of dependencies which are to be used only during development time. These are typically referred to as `dev dependencies` and Snyk does not scan these by default. To include the dev dependencies in the scan results use the `--dev` option when running the `test` command:

`snyk test --dev --file=<manifest file> --json`
