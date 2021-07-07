# Continuous Integration: language support

####  [Submit a Support ticket](https://support.snyk.io/hc/en-us/requests/new) [Sign in to Support](https://support.snyk.io/hc/en-us/signin) [Sign up for free](https://snyk.io/login?cta=sign-up&loc=nav&page=support_docs_page)

###  [ ]() <a id="category-name"></a>

#### CI/CD integrations

* [ Configure your Continuous Integration](/hc/en-us/articles/360004002258-Configure-your-Continuous-Integration)
* [ Continuous Integration: language support](/hc/en-us/articles/360004032157-Continuous-Integration-language-support)
* [ AWS CodePipeline integration](/hc/en-us/articles/4402158184081-AWS-CodePipeline-integration)
* [ Azure Pipelines integration](/hc/en-us/articles/360004127677-Azure-Pipelines-integration)
* [ Bitbucket Pipelines integration overview](/hc/en-us/articles/360004032177-Bitbucket-Pipelines-integration-overview)
* [ CircleCI integration overview](/hc/en-us/articles/360004002278-CircleCI-integration-overview)
* [ Configure your CircleCI integration](/hc/en-us/articles/360004002298-Configure-your-CircleCI-integration)
* [ Getting Snyk Orb details from the CircleCI registry](/hc/en-us/articles/360004032197-Getting-Snyk-Orb-details-from-the-CircleCI-registry)
* [ GitHub Actions integration](/hc/en-us/articles/360019718618-GitHub-Actions-integration)
* [ Bitbucket Pipelines integration](/hc/en-us/articles/360000921778-Bitbucket-Pipelines-integration)

 [See more](/hc/en-us/sections/360001152577-CI-CD-integrations)

1.  [Docs Library \| Snyk](/hc/en-us)
2.  [Integrations](/hc/en-us/categories/360000598398-Integrations)
3.  [CI/CD integrations](/hc/en-us/sections/360001152577-CI-CD-integrations)

##  Continuous Integration: language support

Configure your CI for different languages as follows.

####  For Node.js

1. Install the Snyk utility using `npm install -g snyk`.
2. Run \#`snyk wizard` in the directory of your project following the prompts which will also generate a `.snyk` policy file.
3. Ensure the `.snyk` file you generated was added to your source control \(`git add .snyk`\).
4. If you selected to, Snyk will include `snyk test` as part of your `npm test` command, so if there are new vulnerabilities in the future, your CI will fail, protecting you from introducing vulnerabilities to production. Alternatively, you can add `snyk test` to any other CI test platform you use.
5. Configure your CI environment to include the `SNYK_TOKEN` environment variable. You can find your API token in your [account settings on snyk.io](https://app.snyk.io/account/)

####  For Ruby, Python, Java, Go, and .NET

1. Install the Snyk utility using `npm install -g snyk`.
2. Add `snyk test` to your CI test platform.
3. Configure your environment to include the `SNYK_TOKEN` environment variable. You can find your API token in your account settings on snyk.io.

####  For Scala

1. Install the Snyk utility using `npm install -g snyk`.
2. Install the [sbt-dependency-graph plugin](https://github.com/jrudolph/sbt-dependency-graph).
3. Add `snyk test` to your CI test platform.

* 
