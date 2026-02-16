# CI/CD integrations FAQ

Find answers to the most common questions about integrating Snyk API & Web into your CI/CD pipeline, all at a glance.

<span class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&:nth-child(n+2)]:hidden lg:[&:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9"><span class="text-lg leading-6">C</span></span>

<span class="text-body-secondary-color"></span>

# **Which CI/CD providers do you support?**

You can integrate Snyk API & Web with any CI/CD provider. We support two primary methods:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Snyk API & Web CLI:** This is the recommended method. The CLI is a convenient wrapper for our API that is easy to use in shell scripts for tasks like starting a scan or adding a target.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Direct API:** For more complex integrations or to access the full range of features with fine-grained control, you can interact with the Snyk API & Web API directly.

  </div>

# **Is it possible to block the pipeline if vulnerabilities are detected?**

Yes. We provide blocking mode code examples for [GitLab](https://help.probely.com/en/articles/12692938-integrate-snyk-api-web-with-gitlab-ci-cd), [Bitbucket](https://help.probely.com/en/articles/12692848-integrate-snyk-api-web-with-bitbucket-pipelines), [GitHub Actions](https://help.probely.com/en/articles/8608589-integrate-snyk-api-web-with-github-actions) in their respective guides and our GitHub repo. These examples contain a script that checks the scan results for high-severity findings and blocks the pipeline (or causes the deployment to fail) if any are found.

If you use a different CI/CD tool, you can adapt the logic from these examples to fit your pipeline.

# **My scan is taking too long. What can I do?**

If your pipeline scans are too slow, you have several options to optimize their speed. We recommend a balanced approach: use less comprehensive scans earlier in the development cycle (on feature branches) and more comprehensive scans in pre-production.

Here are a few ways to reduce scan time:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Adjust the scan profile:** You can increase scan speed by choosing a less comprehensive [scan profile](https://help.probely.com/en/articles/8524283-how-to-customize-a-scan-profile).

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Enable [incremental scans](https://help.probely.com/en/articles/5721020-what-are-partial-scans-and-why-are-they-relevant):** This focuses the scan only on new or updated parts of your application, which is significantly faster than a full scan.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Leverage [partial scopes](https://help.probely.com/en/articles/5721020-what-are-partial-scans-and-why-are-they-relevant):** You can configure the scan to check only the specific endpoints or URLs that have been impacted by recent code changes.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Optimize tests:** You can create a custom scan profile and select only a specific list of vulnerabilities to test for. This is faster than running the default comprehensive profile, as it avoids running tests for vulnerability types that may not be a priority for a specific pipeline.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Vary scan strategies by environment:** A scan should be faster the closer it is to the developer. We recommend reducing the comprehensiveness of scans in dev or feature-branch environments to provide fast feedback. You must then run more comprehensive scans in your pre-production environment. You can also set different severity policies for each environment. For example, you might block the pipeline only for critical findings in dev, but block the pipeline for both critical and high findings in QA.

  </div>

# **How do I specify a scan profile in my CI/CD pipeline?**

You can specify a scan profile in two ways:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **To set a default:** You can update the target's settings in the Snyk API & Web UI or via an API call. This profile is used for all scans on that target unless overridden.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **To override for one scan:** For a one-time scan in a CI/CD pipeline, you can use the Snyk API & Web CLI to specify the profile in your YAML configuration file or make a direct API call to the /scan-now/ endpoint.

  </div>

For technical details, see the <a href="https://developers.probely.com/api/overview-api-documentation" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web API documentation</a>.

# **What is the difference between running SAST and DAST in a CI/CD pipeline?**

You typically run a SAST (Static Application Security Testing) scan on your source code *before* the application is built.

A DAST (Dynamic Application Security Testing) scan, like Snyk API & Web, requires a running application to test. Therefore, you must run DAST scans at a later stage in your pipeline, after your application has been deployed to a test or staging environment. DAST scans also typically take more time than SAST scans.

# **How can I dynamically manage target settings in my CI/CD pipeline?**

We recommend using the Snyk API & Web CLI for all CI/CD configurations. It is flexible and allows you to script any action, such as creating, updating, or deleting targets, and initiating scans.

For detailed commands, see the Snyk API & Web <a href="https://developers.probely.com/cli/reference/targets" rel="nofollow noopener noreferrer" target="_blank">CLI Targets Reference</a> and <a href="https://developers.probely.com/cli/reference/scans" rel="nofollow noopener noreferrer" target="_blank">CLI Scans Reference</a>.

# **How can I label findings with metadata like a commit hash or repo name?**

The recommended way to add CI/CD metadata is to set labels on specific findings after a scan completes. This allows you to trace vulnerabilities back to their source, such as the exact commit, branch, or repository.

The workflow from your pipeline script is:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Start the scan and wait for it to finish.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Get the list of new findings via the Snyk API & Web API.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Loop through the new findings and make an API call to add your desired labels (for example, `repo:my-project`, `commit:a1b2c3d4`) to each one.

    </div>

You can find the specific endpoints for this in our <a href="https://developers.probely.com/api/reference/finding-labels/" rel="nofollow noopener noreferrer" target="_blank">Finding Labels Reference</a>.

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F12702107-ci-cd-integrations-faq&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

