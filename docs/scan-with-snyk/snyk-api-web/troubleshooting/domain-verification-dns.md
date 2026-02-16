# Troubleshooting: Can’t verify a domain using DNS

Learn how to troubleshoot issues when verifying a domain using CNAME or TXT records in the DNS.

<img src="../../../.gitbook/assets/snyk-api-web/6c9ff9a1029bc64335fc2a2d9574cf4b0b0ecb44.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Ana Pascoal avatar" />

<span class="text-body-secondary-color"></span>

In order for Snyk API & Web to run full-fledged scans on your target, you need to verify its domain. Learn more about [why do we require you to verify the ownership of a domain](https://help.probely.com/en/articles/3285602-why-do-we-require-you-to-verify-the-ownership-of-your-target).

# The problem

Domain verification using CNAME or TXT records in your DNS fails with the following errors:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  DNS (CNAME) error: `The DNS query name does not exist: <your CNAME record>`\
  ​

  </div>

  <div class="intercom-interblocks-image intercom-interblocks-align-left">

  <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1480098541/370f93d0d51b623b66cca69de968/dns1.png?expires=1769985000&amp;signature=9328a2e1062908b79009b8896df42e4f3216f89b0175718dd059278528608193&amp;req=dSQvFsl3lYRbWPMW1HO4zYis%2BN9mn4zEEDaRP270zwNsMV%2BIO6RoEepTH3lK%0AzBBV%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/edb05a8f2281569da7829f26fc994b7ab457c7fc.png" width="2764" height="1576" /></a>

  </div>

  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  </div>

  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  DNS (TXT) error: `Token not found`\
  ​

  </div>

  <div class="intercom-interblocks-image intercom-interblocks-align-left">

  <a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1480098742/31a213a714be9f5cc36dc6c099ea/dns2.png?expires=1769985000&amp;signature=2d2c06e81efc8dfa4dafebbb6d882bd8bb55cc979b5fb3dfeb5e35c4a4032c31&amp;req=dSQvFsl3lYZbW%2FMW1HO4zSxnmKhPtJyqv16P867SCmnj1WZi18ikhE4lHEQP%0A8aZT%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/5a82f7751613de952f7803d3511f557bf8e46b84.png" width="2764" height="1574" /></a>

  </div>

  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  \
  ​

  </div>

# Troubleshoot the problem

To troubleshoot this problem, go through the following steps to identify the possible causes and respective solutions to fix it.\
​

## Step: Check the TTL value

Check the Time to Live (TTL) as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Go to <a href="https://toolbox.googleapps.com/apps/dig/" rel="nofollow noopener noreferrer" target="_blank">https://toolbox.googleapps.com/apps/dig/</a> and type the target's URL.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Depending on your domain verification method, click on **CNAME** or **TXT** to see the TTL value.

    </div>

If the TTL value is high, the DNS configuration with the CNAME / TXT record has not been propagated yet, and the domain verification fails.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #cccccc"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Cause</strong></p>
</div></td>
<td style="background-color: #cccccc"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p><strong>Solution</strong></p>
</div></td>
</tr>
<tr>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>The DNS configuration with the CNAME / TXT record has not been propagated yet.</p>
</div></td>
<td><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Wait more time for the DNS configuration to propagate, or go to your authoritative DNS server (e.g., Cloudflare) and reduce the TTL value.</p>
</div></td>
</tr>
</tbody>
</table>

After following these steps, identifying the causes, and applying the respective solutions, you should be able to verify your domain using a meta tag.

Learn more about this subject in the following articles:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [Why do we require you to verify the ownership of the domain](https://help.probely.com/en/articles/3285602-why-do-we-require-you-to-verify-the-ownership-of-your-target)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [Verifying your domain](https://help.probely.com/en/collections/3869012-verifying-your-target)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to verify the ownership of a domain using DNS (CNAME Records)](https://help.probely.com/en/articles/5642359-how-to-verify-the-ownership-of-a-target-using-dns-cname-records)

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to verify the ownership of a domain using DNS (TXT Records)](https://help.probely.com/en/articles/3285635-how-to-verify-the-ownership-of-a-target-using-dns-txt-records)

  </div>

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F9657976-troubleshooting-can-t-verify-a-domain-using-dns&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

