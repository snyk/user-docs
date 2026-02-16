# Troubleshooting: Login failed in target authentication with a login form

Learn how to troubleshoot issues in target authentication with a login form.

<img src="../../../.gitbook/assets/snyk-api-web/6c9ff9a1029bc64335fc2a2d9574cf4b0b0ecb44.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Ana Pascoal avatar" />

<span class="text-body-secondary-color"></span>

In targets with authentication, Snyk API & Web scans must log in to reach areas reserved for authenticated users to scan them for vulnerabilities.

# The problem

When running scans on a target with a login form, Snyk API & Web fails to log in.\
​

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1396124277/783a0c5cc9e86b95b89e45153001/Screenshot+2025-02-25+at+15_28_21.png?expires=1769985000&amp;signature=16be438639ecc69d9fdabd8fd3adffe4a7ed6c49dd767a02b39241fc246e7bc0&amp;req=dSMuEMh8mYNYXvMW1HO4zbXajBt9Do%2BMh6tDfjmcU1LPYtFgG9wRZ2izf5xZ%0AG9jhwDoSYcljAYmqNec%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/9af6d2942b27081f6f8de4c8b3d993a1535894da.png" width="2590" height="730" /></a>

# Troubleshoot the problem

To troubleshoot this problem, go through the following steps to identify the possible causes and respective solutions to fix it.

## Step 1: Test the current credentials

Test if the current credentials configured in the target settings are still valid, as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the Snyk API & Web, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list and click on the cogwheel to open its settings.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the **AUTHENTICATION** tab, check the **LOGIN FORM** configuration, get the URL of the login form, and the current login credentials.

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser and type the target’s login URL.

    </div>

5.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Log in with the current credentials.

    </div>

If the login fails, check the following possible causes and apply the respective solution:

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
<p>The credentials are invalid.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Obtain valid login credentials and update them in the target settings.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form">How to set up Target Authentication with a Login Form</a>.</p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The credentials expired.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Obtain new login credentials and update them in the target settings.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form">How to set up Target Authentication with a Login Form</a>.</p>
</div></td>
</tr>
</tbody>
</table>

## Step 2: Test the login flow

Test if the login flow is still a login form, as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser and type the target’s login URL.

    </div>

If the login is not a login form, the target authentication fails.

Check the following possible causes and apply the respective solution:

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
<p>The login flow is not a login form but a complex login (e.g., multi-step login).</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Configure the target authentication to use a login sequence, which supports complex logins.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/5402712-how-to-set-up-target-authentication-with-a-login-sequence">How to set up Target Authentication with a Login Sequence</a>.</p>
</div></td>
</tr>
</tbody>
</table>

## Step 3: Check the field names

Check the values configured in field names in the target authentication with a login form as follows

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the Snyk API & Web, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Identify the target in the list and click on the cogwheel to open its settings.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the **AUTHENTICATION** tab, go to the **LOGIN FORM** configuration to see the configured **field names** (typically, one for the username and another for the password).

    </div>

4.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser and type the target’s login URL.

    </div>

5.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Right-click and select **Inspect** to see the attributes of the input fields on the login form.

    </div>

6.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the **LOGIN FORM** configuration, check if the values set in the **field names** for the input fields contain a valid “id”, “name”, or CSS selector from the login form.

    </div>

If the values configured in the **field names** are not valid, Snyk API & Web scans cannot authenticate.

Check the following possible causes, and apply the respective solution:

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
<p>The value configured in a <strong>field name</strong> does not contain a valid “id”, “name”, or CSS selector that identifies that input field in the login form.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Go to the target authentication settings and set the value of the <strong>field name</strong> with the “id,” “name,” or the CSS selector that uniquely identifies that input field in the login form.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form">How to set up Target Authentication with a Login Form</a>.</p>
</div></td>
</tr>
</tbody>
</table>

## Step 4: Test for a blocking WAF

Test if there is a WAF blocking access to the authentication page with the login form as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Open a browser in incognito mode and type the target’s URL.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Go to the authentication page with the login form.

    </div>

If a WAF blocks access to the authentication page with the login form, Snyk API & Web scans cannot authenticate.

Check the following possible causes and apply the respective solution:

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
<p>A WAF is blocking access to the authentication page with the login form.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Add Snyk API &amp; Web IPs to the WAF’s whitelist.</p>
<p>Learn more about <a href="https://help.probely.com/en/articles/8931142-how-to-configure-probely-s-ips-in-wafs">How to configure Snyk API &amp; Web IPs in WAFs</a>.</p>
</div></td>
</tr>
</tbody>
</table>

After following these steps, identifying the causes, and applying the respective solutions, scans should be able to log in to your target.

Learn more about this subject in the following articles:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [Scanning Targets with Authentication](https://help.probely.com/en/articles/9177848-scanning-targets-with-authentication)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to set up Target Authentication with a Login Form](https://help.probely.com/en/articles/3292779-how-to-set-up-target-authentication-with-a-login-form)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to set up Target Authentication with a Login Sequence](https://help.probely.com/en/articles/5402712-how-to-set-up-target-authentication-with-a-login-sequence)

  </div>

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F9639573-troubleshooting-login-failed-in-target-authentication-with-a-login-form&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

