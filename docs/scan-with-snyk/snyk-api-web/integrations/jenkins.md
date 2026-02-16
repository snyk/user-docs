# Integrate Snyk API & Web with Jenkins

Configure Jenkins CI/CD pipeline to scan your application for security issues.

<img src="../../../.gitbook/assets/snyk-api-web/0581f948f3d016b0996055b218f51b7485035d04.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Tiago Mendo avatar" />

<span class="text-body-secondary-color"></span>

With the Snyk API & Web plugin you can automatically start a scan every time your Jenkins pipeline is executed.\
​\
Jenkins allows you to have an arbitrary number of build and test scenarios, but a common pattern is as follows:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Build step**: compiles the application or creates the Docker containers

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Deploy step**: sends the compiled code or the containers to test server and execute them

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  **Test step**: execute tests on the running application

  </div>

The Snyk API & Web plugin is a build task that should run after the application is built and deployed.\
It is recommended that the security tests run after the integration/functional tests pass, to ensure the application is working properly. A broken application may cause security tests to miss vulnerabilities because a particular feature is not working.

# Installation and setup

Installing and setting up the plugin will take you less than 5 minutes.

1\. Open Jenkins and click on **Manage Jenkins**

<a href="https://downloads.intercomcdn.com/i/o/109267814/7bbe4c112c039de72c9c71dc/install_plugin_1.png?expires=1769985000&amp;signature=0f13798dcfe4bad50bd21b5d8ef36c657bba7e01827905d3f8b80b4c8c93cc2a&amp;req=dSAuFM95lYBbFb4f3HP0gLNgvMy9Up6e%2Fi05PnonKUpsCt9niWKui%2B9Vs%2B5k%0AR0mft20leqGQ6wppGw%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/6e1aa9a5b431bb2f58c01804e4fecf8160d225e3.png" width="2035" height="930" /></a>

2\. Click on **Manage Plugins\
​**

<a href="https://downloads.intercomcdn.com/i/o/109267813/cf80a245429c1c84daaeee71/install_plugin_2.png?expires=1769985000&amp;signature=f3aa223b1b527347ef507aa8605093447e41e12852717f878a1c404694186cb4&amp;req=dSAuFM95lYBcFb4f3HP0gGIF6hoUlGchnXYS2QvZ6MGZv6E0DjB8CZrdt0S4%0A%2FtBLHbA9h4rv0c8STA%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/e531dd7a599eac449637b6fbe87bbe9378e79cc7.png" width="2012" height="1050" /></a>

3\. Click on **Available** tab

<a href="https://downloads.intercomcdn.com/i/o/109267812/1e18d87e91382dcfe73c13fb/install_plugin_3.png?expires=1769985000&amp;signature=97e8758ddf68b1b6729bfa09cb4a3f31ea5bab4330808cc72558f2888fa163e3&amp;req=dSAuFM95lYBdFb4f3HP0gD7AiNZKfdGZmMZ76DwbW62aPu8f3QMyAp6I4KR9%0As33ovBM3y3jQtpyQVQ%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/cacd6ab090378d242f564cbcc82d6b66fa02d07d.png" width="2038" height="610" /></a>

4\. On the **Filter** search box, enter **probely**\
5. Select the **Probely Security Scanner** plugin\
6. Click on **Download now and install after restart**\
7. After Jenkins restarts, the plugin will be installed. Continue reading to set up the required API key from Snyk API & Web.

Generating an API key\
======================

Before using the plugin, you must [generate an API Key](https://help.probely.com/en/articles/8592281-how-to-generate-an-api-key) for Jenkins to be able to start a scan with Snyk API & Web.

Once the API key is created, take note of its value, as it will be required to configure the Plugin credentials later on, and it will not be displayed again. You will also need the ID of the target you want to scan; you can obtain this ID from the target page.

# Configuring the plugin

The plugin can be used in both Freestyle and Pipeline projects, and this post will provide an example for each one. You can learn more about these two project types and their differences <a href="https://jenkins.io/pipeline/getting-started-pipelines/" target="_blank" rel="nofollow noopener noreferrer">here</a>.

## Configuring credentials

1\. Click on **Credentials\
​**2. Click on the down arrow near (**global**) to enable the dropdown menu and choose **Add credentials**

<a href="https://downloads.intercomcdn.com/i/o/109268804/4b0da937cb37c7a08c1ec569/credentials_1.png?expires=1769985000&amp;signature=670071e62ccdaebe44588f67fbdf706fdad88d9483b34fde4836e79670755fb9&amp;req=dSAuFM92lYFbFb4f3HP0gOJlEcBIC3phlHrV1ouSnvxnpYB%2FCX%2FN2YnKo5nI%0ALShGugSXM6t1voBN%2BQ%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/b7e7637addbde29c47427751ea48f441677222b8.png" width="2039" height="899" /></a>

1\. On the Kind dropdown menu, choose **Secret text**\
2. Enter the API key in the **Secret** textbox\
3. Enter a value for the credentials in the **ID** textbox, for example **probely-test-site**\
4. Enter an optional Description and click **OK**

<a href="https://downloads.intercomcdn.com/i/o/109268815/83ff67bb444fe2baf7a951fb/credentials_2.png?expires=1769985000&amp;signature=f7cb30594e033f665d2471a29e75372a7cab8133e39c98c940b1e94034ff0329&amp;req=dSAuFM92lYBaFb4f3HP0gHb%2BTEudV%2BsmaLBqzo2Xe%2FcnXVUg5nIkNz2lRvWw%0A0GjCyHN9Paw3xZCh2A%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/5d6d3b2b09c7d8ab5a6c4d358394570167300eba.png" width="2036" height="645" /></a>

## Using the plugin in a Freestyle project

Creating a freestyle is the simplest way to have a repeatable process to build and test your application, especially for simple applications with just a few jobs.\
If you already have a freestyle project you only need to configure the plugin: in the project listing page, click the Configure in the drop-down menu next to the project name. If you are creating a new project, follow the next steps:

1\. Click **New Item.**

<a href="https://downloads.intercomcdn.com/i/o/109268971/7905ca32ffe82cc90230499b/new_item.png?expires=1769985000&amp;signature=2472f871c7d715ea1ac73e844f5498d4a31069af14d8d51576fa5525f54ecbcd&amp;req=dSAuFM92lIZeFb4f3HP0gNocjL5%2FjozUMHBlE9r4MmbzXpvATL8Cu40xXhR9%0A8yYiUMUVJFWq%2Bf8kAg%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/91ea0428838887a8184bf0f70d7a94b065ca9414.png" width="2550" height="1030" /></a>

2\. Enter your project name, choose **Freestyle Project** and click **OK**

<a href="https://downloads.intercomcdn.com/i/o/109268985/b81d28756fc47342eb0fc72e/freestyle_0.png?expires=1769985000&amp;signature=c465ea4b2b8a3d0837317881d529b6fdd1ebf9a28ff44de05cfedcd6502a8d91&amp;req=dSAuFM92lIlaFb4f3HP0gDzlyftGEyVDSYCfIJRTeX%2B56TZ%2Bc%2BHbfesIPwot%0A3w0%2FQ6%2BATE3o%2Fd7K3w%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/9a9bf8dd9eca389993222615ff933cc10f51a861.png" width="2014" height="632" /></a>

3\. In the **Build** section add **Probely scan step**.

We assume you have configured the other project options properly, such as checking out from your SCM, building the code and deploying it.

In the just added **Probely** **Security Scanner** section:

1\. Add the **Target ID** of the target you want to scan.\
2. Select the right credentials, which were configured in **Configuring credentials**. If the connection to Snyk API & Web's API is working correctly, and the credentials are valid, you should see the message "Credentials verified successfully".\
3. When all steps are properly configured, click on **Save**.

<a href="https://downloads.intercomcdn.com/i/o/109269575/1a0725b700b5af448c4f5cd1/freestyle_1.png?expires=1769985000&amp;signature=3d146e3c4e9e74c1a8f7dfe0ac07dcb9259a29f8fa9b007184bbbe0c75f8ed5e&amp;req=dSAuFM93mIZaFb4f3HP0gJIWzGKbXyLEQBgn1MKdDOXzkzCHfY%2BrhyeGuDox%0Aik1HQAl6gVUQZJBcQw%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/93a53e2f409dc89d9f2bf2e6cc6b00fe6cb154de.png" width="1530" height="765" /></a>

The next time the build job for this project runs, Snyk API & Web will test the security of the configured target and will send you an email with the scan results in the end.

## Using the plugin in a Pipeline project

Pipeline projects are the most flexible and powerful way of creating CI/CD pipelines with Jenkins.

The projects need a configuration file, a **Jenkinsfile**. The one in our example uses the more modern declarative syntax, instead of the imperative one.

1\. Click on **New Item**

<a href="https://downloads.intercomcdn.com/i/o/109269788/2079a908a40d8af95b738dc5/new_item.png?expires=1769985000&amp;signature=0ef181036bb853c2f44e64f4ebd770a45029adf631bc3ea74d509dd209aeb606&amp;req=dSAuFM93molXFb4f3HP0gHfT0C1myx2iUAzgDvuqxe%2BAsFgM4Ln3fIowt2t%2F%0AF5k89XJnwIld4kQxsw%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/91ea0428838887a8184bf0f70d7a94b065ca9414.png" width="2550" height="1030" /></a>

2\. Enter your project name, choose **Pipeline Project** and click **OK**

<a href="https://downloads.intercomcdn.com/i/o/109269797/17bc642503f68c149c31203c/pipeline_0.png?expires=1769985000&amp;signature=ea49d5040e1916b937455b53864f4e3ed4d99a5cf12d5cc527f2cf4a8c93f931&amp;req=dSAuFM93mohYFb4f3HP0gA8c3UOkkM90tWAwWU9N5uNxFIV4g7keUrCSCzN2%0AlwlmJdO96o9AKZevLQ%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/5e416319f7413c7c8c86eb6ae9d8f0df16b44962.png" width="2522" height="1372" /></a>

3\. Create a Jenkinsfile

    pipeline {
        agent {
            docker {
                image 'maven:3-alpine' 
            }
        }
        stages {
            stage('Unit tests') { 
                steps {
                    sh './gradlew check'
                }
            }
            stage('Scan with Probely') {
                steps {
                    probelyScan targetId: '9nl6yy0TWWKv', credentialsId: 'probely-test-site' 
                }
             }
        }
    }

As with the Freestyle project, the security tests are executed after the functional tests, in this case after the Unit tests stage, to ensure the application is working properly.

4\. Configure Jenkins to use the Jenkins file on your repository

<a href="https://downloads.intercomcdn.com/i/o/109269812/f4e5fe5ad221ff39edac7596/pipeline_1.png?expires=1769985000&amp;signature=72577d376ea95fa4a147efbcd7f3157b86e99b0ec64efff6e0fca6aa93bd7146&amp;req=dSAuFM93lYBdFb4f3HP0gPrR1TiCrTaFcdDqDelqT0R4gQQcjqfAmB4Wm3jU%0AsA4g6yP6eiZQl1%2BYdQ%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/6f72e1fbbaae7d3f545851c94ae4709207ee590d.png" width="1926" height="816" /></a>

If your Jenkinsfile is stored in a repository that was already configured here (Definition **Pipeline script from SCM)** you only need to commit the updated file to the repository

Hit **Save.** Your pipeline has now a stage that scans your target for vulnerabilities.\
​

You have some options on how Snyk API & Web scans your target, at the target settings page: you can choose different scan profiles, configure authentication in order to scan behind login pages, add custom headers and even enable automatic synchronisation with your project at Jira, just to name a few.

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F2771631-integrate-snyk-api-web-with-jenkins&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

