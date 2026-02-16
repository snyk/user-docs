# How to scan internal applications with a Scanning Agent

Scan your internal applications with Snyk API & Web's Scanning Agent. A secure, clean, and easy-to-set-up solution to scan non-public applications.

<img src="../../../.gitbook/assets/snyk-api-web/0581f948f3d016b0996055b218f51b7485035d04.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Tiago Mendo avatar" />

<span class="text-body-secondary-color"></span>

This article overviews what a Scanning Agent is, how it works, and how to install and scan with it.

# What is a Scanning Agent for?

Snyk API & Web's Scanning Agent allows you to scan internal applications for vulnerabilities without exposing them to the Internet or even to our IP addresses. It is the ideal approach to scan any application that is only reachable from within your network, including development/staging/pre-release and/or internal production applications that support your business.

You can use a single Scanning Agent to scan multiple internal targets, but you can also have different Scanning Agents, each one reaching a part of your network. There is no need for a single Scanning Agent to connect to the whole network.

# How does a Scanning Agent work?

A Scanning Agent creates an encrypted and authenticated tunnel where traffic flows securely between Snyk API & Web and your network.

To make sure we meet your security expectations, we follow a set of principles:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  All code is open source and <a href="https://github.com/Probely/farcaster-onprem-agent/" rel="nofollow noopener noreferrer" target="_blank">publicly available</a>.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  You have complete control over the Scanning Agent, including the right to change it.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  Snyk API & Web cannot access the Scanning Agent.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  The Scanning Agent runs in containers with the least required privileges.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  All traffic is encrypted end-to-end.

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  The Scanning Agent does not open any network port.

  </div>

# How to install a Scanning Agent?

To install a Scanning Agent, refer to this article on [how to install a Scanning Agent](https://help.probely.com/en/articles/6503388-how-to-install-a-scanning-agent) and the installation reference and source code for the installer available at <a href="https://github.com/Probely/farcaster-onprem-agent/" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web's GitHub repositories</a>.

# How to scan a target with a Scanning Agent?

When a Scanning Agent is configured and running, you must choose which targets will use it:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the <a href="https://plus.probely.app/" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web app</a>, go to the **Targets** menu.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list for which you want to set the Scanning Agent and click on the cogwheel to open its settings.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Under the **Scanner** tab, go to the **SCANNING AGENT** section and select the Scanning Agent you want to use.

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click **Save**.

    </div>

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1421050675/bb482151483ffc3c74c0dfc3771a/Screenshot+2025-03-13+at+15_46_33.png?expires=1769985000&amp;signature=53bfb100d4c1d171a2e5c1c130acad8db9425d58f6f8ef2c443f3a62239d4b99&amp;req=dSQlF8l7nYdYXPMW1HO4zVJRwp1SHvvPfwD%2FE9kpNUkUUepze2NKFCU%2FrSLW%0AGU4lMuTP0OPnrIC7dRA%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/55011671f9a9c6651d8392f08610a47d6767a175.png" width="1272" height="838" /></a>

Clicking **Unlink** removes the Scanning Agent for the target.

You can also assign/remove a Scanning Agent to/from multiple targets in the targets list. Select the targets you want to configure, and the options will appear:

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1421048788/7e5cf24db9dcb9564960e5ceedba/Screenshot+2025-03-13+at+15_44_26.png?expires=1769985000&amp;signature=85def0a236625a583fc890aeb5e2c8732ef4b3ddc076ec23681d282541879724&amp;req=dSQlF8l6lYZXUfMW1HO4zUnkzviGM2sbIu4npukmiJWGijiSxzVTMvd3ca4P%0AAT%2BJnE8pbyYioC31HPk%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/7c6bbaeffeeea9a91f5b5b5894211d0f7cd653b1.png" width="2734" height="820" /></a>

As the image above shows, targets configured to use a Scanning Agent will show a cloud icon.

# What are the statuses of a Scanning Agent?

A Scanning Agent can have one of the following statuses:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Status</strong></p>
</div></td>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Description</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><span class="image placeholder" data-original-image-src="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1344787811/5449e4722edd85696dc4bb4d89cc/State%3Don.png?expires=1770066000&amp;signature=9e77509c4580b7ce9e0c536845eb651aa6b02f5ab617fb76d18a376196e64342&amp;req=dSMjEs52moleWPMW3Hu4gSt9vB4uaJZH8iJIfxU3Q0u2PIoVw5TkFZcU2zQ6%0AkQ%3D%3D%0A" data-original-image-title="" width="20"></span> Connected</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The scanning agent is connected. It was working in the last 180 seconds.</p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><span class="image placeholder" data-original-image-src="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1344788748/949754dcb9e6bc5e3ba21630e9b6/State%3Dwarning.png?expires=1770066000&amp;signature=2df77c08af06b655844848edfa131a525ce3865f9bac27d96ef8def81361eb90&amp;req=dSMjEs52lYZbUfMW3Hu4gbbqkQ%2F0iQmLE4kIzhbG2T01QYjxcBmucUkMbm6w%0AlA%3D%3D%0A" data-original-image-title="" width="20"></span> Connected with issues</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The scanning agent is connected, but it may have poor network performance if it uses, for example, an HTTP proxy or a direct TCP connection to Snyk API &amp; Web.</p>
<p>For more information, see this article about the <a href="https://web.archive.org/web/20220103191127/http://sites.inka.de/bigred/devel/tcp-tcp.html" rel="nofollow noopener noreferrer" target="_blank">TCP Meltdown</a> problem and check the documentation on <a href="https://github.com/Probely/farcaster-onprem-agent?tab=readme-ov-file#launch-the-agent" rel="nofollow noopener noreferrer" target="_blank">launching the agent</a>.</p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><span class="image placeholder" data-original-image-src="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1344789068/ffb6fe7836037b0a7b2d8c8b0198/State%3Doff.png?expires=1770066000&amp;signature=165ed7590e7b54e3158280f5f2fbf4e2b6fcd02d48f197d544896ce1cf0a5b3d&amp;req=dSMjEs52lIFZUfMW3Hu4gTJ6JuIiaMTEUR0qSQ%2FAbaxlilQG3mJLVrv1oZ28%0Axg%3D%3D%0A" data-original-image-title="" width="20"></span> Disconnected</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The scanning agent is disconnected, maybe due to misconfiguration.</p>
<p>Check the scanning agent configuration or the firewall rules, for example.</p>
<p>For more information, check the <a href="https://github.com/Probely/farcaster-onprem-agent?tab=readme-ov-file#installation" rel="nofollow noopener noreferrer" target="_blank">Installation</a> and <a href="https://github.com/Probely/farcaster-onprem-agent?tab=readme-ov-file#network-requirements" rel="nofollow noopener noreferrer" target="_blank">Network Requirements</a> documentation.</p>
</div></td>
</tr>
</tbody>
</table>

If you still need help, don't hesitate to message us or send us an email to <a href="/cdn-cgi/l/email-protection#0f7c7a7f7f607d7b4f7f7d606d6a6376216c6062" rel="nofollow noopener noreferrer" target="_blank"><span class="__cf_email__" data-cfemail="abd8dedbdbc4d9dfebdbd9c4c9cec7d285c8c4c6">[email protected]</span></a>.

P.S.: Why is the Scanning Agent named Farcaster? Learn more about it [here](https://help.probely.com/en/articles/4615879-why-is-the-scanning-agent-named-farcaster).

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F4615595-how-to-scan-internal-applications-with-a-scanning-agent&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

