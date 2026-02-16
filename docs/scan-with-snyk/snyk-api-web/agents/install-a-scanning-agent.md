# How to install a Scanning Agent

Learn how to install Snyk API & Web Scanning Agent

<img src="../../../.gitbook/assets/snyk-api-web/5530a2b28285ea60de10f16a64f57b44160bec88.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Martim Valente avatar" />

<span class="text-body-secondary-color"></span>

The Scanning Agent allows scanning internal applications with a secure, clean, and easy-to-set-up solution to scan your non-public applications. You can learn more about the Scanning Agent in this article on [how to scan internal applications with a Scanning Agent](https://help.probely.com/en/articles/4615595-how-to-scan-internal-applications-with-a-scanning-agent).

Installing a Scanning Agent involves two steps:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Create the Scanning Agent in Snyk API & Web.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Install Snyk API & Web's Farcaster Agent on your network.

    </div>

This article describes these steps in detail.

# Step 1: Create the Scanning Agent

This first step creates the Scanning Agent on the Snyk API & Web side:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open the **Settings** dropdown menu on the bottom-left corner of the navigation bar and click on **Scanning Agents**. If you do not see this option, contact your account owner.\
    ​

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-center">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1344737459/75f62bc4ae5162f03868c3a0f695/settings_agents.png?expires=1769985000&amp;signature=c0d03dfa17393da5f7425d7ea7544ebbfa961b1ab4158604e533475f639f1ba3&amp;req=dSMjEs59moVaUPMW1HO4zZJ9vnegueZIic4Mo1TIgS9%2FRNgyygp5%2FbJ4sbnT%0AgkPs%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/94027d31c236b0c7818aeabdbac8bfe567f43b84.png" width="225" height="517" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **ADD AGENT**.

    </div>

    <div class="intercom-interblocks-ordered-nested-list">

    1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

        Type the name of the Scanning Agent.

        </div>

    2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

        If the Scanning Agent is restricted to targets of some teams, tick the checkbox and select those teams from the dropdown. Learn [how to get started with Teams](https://help.probely.com/en/articles/6345396-how-to-get-started-with-teams).\
        ​

        </div>

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **GENERATE**.\
    ​

    </div>

    <div class="intercom-interblocks-image intercom-interblocks-align-left">

    <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1704832194/59b8f95bdb719ff000ce105efb7c/Screenshot+2025-09-02+at+11_53_04.png?expires=1769985000&amp;signature=5a4aaedbaa29d569fe57e0ba727433859dab4e86ac28641d3710c605549774fe&amp;req=dScnEsF9n4BWXfMW1HO4zTt%2BqFhd3WTHyddQHfVlN2%2FIVbjv73tUZB8R2EZ8%0Apvnv%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/d8efd6c9b328014459946ac5958fbb75aaff3149.png" width="1168" height="1426" /></a>

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    A pop-up window is displayed with important information that, for security reasons, will not be visible again. So, you should do the following:

    </div>

    <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    </div>

    <div class="intercom-interblocks-ordered-nested-list">

    1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

        Under **Agent Token**, copy and save the token securely.\
        ​

        </div>

    2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

        Under **Installation**, go to the tabs for the way you want to install the agent:\
        ​

        </div>

        <div class="intercom-interblocks-ordered-nested-list">

        1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

            **DOCKER** - To use Docker, copy and save securely the following:

            </div>

            <div class="intercom-interblocks-ordered-nested-list">

            1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Docker command to install the agent.

                </div>

            2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Docker command to check the agent logs.\
                ​

                </div>

            </div>

        2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

            **DOCKER-COMPOSE** - To use Docker-compose, copy and save securely the following:

            </div>

            <div class="intercom-interblocks-ordered-nested-list">

            1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The `docker-compose.yml` manifest for the agent.

                </div>

            2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Docker-compose command to start the agent.\
                ​

                </div>

            </div>

        3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

            **KUBERNETES** - To use Kubernetes, copy and save securely the following:

            </div>

            <div class="intercom-interblocks-ordered-nested-list">

            1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Kubernetes command to create the Snyk API & Web namespace.

                </div>

            2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Kubernetes command to create the agent token secret.

                </div>

            3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

                The Kubernetes command to deploy the agent pod.

                </div>

            </div>

            <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

            </div>

        </div>

    </div>

# Step 2: Install Snyk API & Web's Farcaster Agent

In this step, you will install Snyk API & Web’s Farcaster Agent on your network. It is the piece that will allow connecting Snyk API & Web Cloud servers and your internal applications to perform scans. For that, read and follow <a href="https://github.com/Probely/farcaster-onprem-agent/" rel="nofollow noopener noreferrer" target="_blank">the instructions on GitHub</a> using the information you saved in Step 1. of this article.\
​

<a href="https://probely.intercom-attachments-7.com/i/o/920910900/c4498935c2aee36da8f93ba4/Z9dBejPHOPIjpSKPmlcYN3WdniH7_PzjKWfTrKcQpTalAuxFhvKxdvkUS_akYOhGVBqpli08BwKeHhZZbljq3PiPZPW6opFhe8ArYLOdZZ0K5Nvasn8STo3Ug9uVqrIU2Uk9Rydts_YE9lHMQS7VFGk?expires=1769985000&amp;signature=7cbbb8cb26aeabc8a72b147ab7301c6986fd79014dbf66865f8294134eeee31b&amp;req=fSInH8h%2BlIFfFb4f3HP0gEqzTh4hWQm8yPqK1cbYD39j5JZQGtAkL%2BhquZdB%0A8%2Btv8ojaPjVUj7FgQQ%3D%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/f69f651edd1489db1e7dbe4da5aebd9d43695fbd.png" /></a>

In any case, you can quickly install the Farcaster Agent by executing the following command in your terminal:

    docker run -d \ 
    --cap-add NET_ADMIN \ 
    --device /dev/net/tun \ 
    -e FARCASTER_AGENT_TOKEN=<YOUR_FARCASTER_AGENT_TOKEN> \ 
    probely/farcaster-onprem-agent:v3

In the command, replace **\<YOUR_FARCASTER_AGENT_TOKEN\>** with the token you generated and saved in Step 1.

After running this command, the container should be up and running. You just have to set the Scanning Agent in your targets (see the section **How to scan with a Scanning Agent?** in [this article](https://help.probely.com/en/articles/4615595-how-to-scan-internal-applications-with-a-scanning-agent)) and run scans on those targets to scan your internal applications.

Please don't hesitate to contact us with any questions about the Scanning Agent.

P.S.: Why the name Farcaster? Read about it in [Why is the Scanning Agent named Farcaster?](https://help.probely.com/en/articles/4615879-why-is-the-scanning-agent-named-farcaster)

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F6503388-how-to-install-a-scanning-agent&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

