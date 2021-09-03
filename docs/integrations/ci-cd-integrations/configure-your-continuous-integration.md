# Configure your Continuous Integration

To continuously avoid known vulnerabilities in your dependencies, integrate Snyk into your continuous integration \(a.k.a. build\) system. In addition to the documentation here, you're also invited to check out our [integration configuration examples](https://github.com/snyk-samples/snyk-cicd-integration-examples) in our GitHub repository.

### Set up automatic monitoring

If you monitor a project with Snyk, you’ll get notified if your project’s dependencies are affected by newly disclosed vulnerabilities. To make sure the list of dependencies we have for your project is up to date, refresh it continuously by running Snyk monitor in your deployment process. Configure your environment to include the SNYK\_TOKEN environment variable. You can find your API token on the dashboard after logging in.

### API token configuration

Make sure you don’t check your API token into source control, to avoid exposing it to others. Instead, use your CI environment variables to configure it.

See guidance for how to do this on:

* [Travis](https://docs.travis-ci.com/user/environment-variables/)
* [Circle](https://circleci.com/docs/environment-variables/)
* [Codeship](https://codeship.com/documentation/continuous-integration/set-environment-variables/)

You can find others through an easy [Google search](https://www.google.co.uk/search?q=setting+up+env+variables+in+CI).  
Learn more about [CI/CD](https://snyk.io/learn/what-is-ci-cd-pipeline-and-tools-explained/).

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}