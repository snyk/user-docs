# Troubleshooting: Low coverage in a scan

Learn how to troubleshoot issues in a target scan with low coverage.

<img src="../../../.gitbook/assets/snyk-api-web/6c9ff9a1029bc64335fc2a2d9574cf4b0b0ecb44.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Ana Pascoal avatar" />

<span class="text-body-secondary-color"></span>

Scans should cover as much of the target scope as possible to identify the maximum number of vulnerabilities. Learn more about [how to generate a coverage report](https://help.probely.com/en/articles/3292591-how-to-generate-a-csv-coverage-report) and [what is the meaning of the coverage report?](https://help.probely.com/en/articles/5522983-what-is-the-meaning-of-the-csv-coverage-report)

# The problem

When running a scan on a target, the coverage is low.\
​

<a href="https://probely.intercom-attachments-7.com/i/o/1125491678/abf57de7dc95fdbffade8819/AD_4nXdljbhciLaF3jCcrUUsLnz9cCZtye8uScXAQh9UDnvlREzCWdbJPPps9kctFRDoIDYvoh8bKBfLjn3-HTkTNIWoJRdEeuI1xYotVULVgEmVtGVkeCaz0QEDBLHZQCbKIUtr3Ckz5FqrGFNUEENSf2jT2z2a?expires=1769985000&amp;signature=7b68a28865d4e9ab335caaa63da87dcccb0d002f3b016e374908e3df174d1b3d&amp;req=dSElE813nIdYUfMW1HO4zcCwDeiOG%2FmZ6sdDVZFoJfR3tcHLyYP%2BWzWr8C3L%0A7nlUJQrYSKveMAIXqHs%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/5bec8f84f07fe6f814f798ef89cb80c2e7fcffad.png" /></a>

# Troubleshoot the problem

To troubleshoot this problem, go through the following steps to identify the possible causes and respective solutions to fix it.

## Step 1: Check for target authentication

If the target has authentication, check if the scanner was able to log in:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the Snyk API & Web, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list, and click on its name to see its details.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **SCAN ACTIVITY** to see the list of scans.

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the scan in the list, and click on **VIEW**.

    </div>

5.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **CRAWLING REPORT** to get the spreadsheet with the scanned URLs.

    </div>

6.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify URLs that are only available for authenticated users.

    </div>

If no URLs for authenticated users are listed, the scan must have failed to log in.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #e8e8e8"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Cause</strong></p>
</div></td>
<td style="background-color: #e8e8e8"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Solution</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The scan failed to log in to the target.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Check the target authentication configuration.</p>
<p>Learn more about <a href="https://help.probely.com/en/collections/9914791-troubleshooting-target-autentication">Troubleshooting Target Authentication</a>.</p>
&#10;</div></td>
</tr>
</tbody>
</table>

## Step 2: Check for missing SPA API

If the target is a SPA (Single-Page Application) with a backing API, check if the API is in a different URL. For example:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  SPA URL: <a href="https://example.com" rel="nofollow noopener noreferrer" target="_blank">https://example.com</a>

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  SPA API URL: <a href="https://api.example.com" rel="nofollow noopener noreferrer" target="_blank">https://api.example.com</a>

  </div>

  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  </div>

If the backing API has a URL different from the SPA, Snyk API & Web scans need to know the API URL to scan the SPA properly.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Cause</strong></p>
</div></td>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Solution</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The target is a SPA with its backing API in a different URL.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Go to the target settings, and add an extra host with the URL of the backing API.</p>
&#10;</div></td>
</tr>
</tbody>
</table>

## Step 3: Check for a blocking WAF

Check if scan requests started being blocked by a WAF after the scan has started:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the Snyk API & Web, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list, and click on its name to see its details.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **SCAN ACTIVITY** to see the list of scans.

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the scan in the list, and click on **VIEW**.

    </div>

5.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **CRAWLING REPORT** to get the spreadsheet with the scanned URLs.

    </div>

6.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Check if, at some point, the URLs started having HTTP error status 403.

    </div>

7.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser in incognito mode, type those URLs to test them, and see if a WAF is blocking the access.

    </div>

If a WAF starts blocking access to URLs, Snyk API & Web cannot scan them.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Cause</strong></p>
</div></td>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Solution</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>A WAF started blocking access to the URLs during the scan.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Add Snyk API &amp; Web IPs to the WAF’s whitelist.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/8931142-how-to-configure-probely-s-ips-in-wafs">How to configure Snyk API &amp; Web IPs in WAFs</a>.</p>
</div></td>
</tr>
</tbody>
</table>

## Step 4: Check for blocking WordPress plugin

If the target is WordPress, check if scan requests are being blocked by a WordPress plugin (e.g., WordFence):

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the Snyk API & Web, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list, and click on its name to see its details.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **SCAN ACTIVITY** to see the list of scans.

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the scan in the list, and click on **VIEW**.

    </div>

5.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Click on **CRAWLING REPORT** to get the spreadsheet with the scanned URLs.

    </div>

6.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Check if the URLs have HTTP error status 403.

    </div>

7.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser in incognito mode, type those URLs to test them, and see if a WordPress plugin is blocking the access.

    </div>

If a WordPress plugin is blocking access to URLs, Snyk API & Web cannot scan them.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Cause</strong></p>
</div></td>
<td style="background-color: #e8e8e880"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Solution</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>A WordPress plugin (e.g., WordFence) is blocking access to the URLs.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Configure the WordPress plugin to allow requests from Snyk API &amp; Web IPs.</p>
<p>Refer to this article <a href="https://help.probely.com/en/articles/1975548-what-is-the-scanner-s-outgoing-ip-address">What is the scanner's outgoing IP address?</a></p>
</div></td>
</tr>
</tbody>
</table>

After following these steps, identifying the causes, and applying the respective solutions, scans should have the expected coverage for your targets.

Learn more about this subject in the following articles:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [What is the meaning of the coverage report?](https://help.probely.com/en/articles/5522983-what-is-the-meaning-of-the-csv-coverage-report)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to generate a coverage report?](https://help.probely.com/en/articles/3292591-how-to-generate-a-csv-coverage-report)

  </div>

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F9662415-troubleshooting-low-coverage-in-a-scan&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

