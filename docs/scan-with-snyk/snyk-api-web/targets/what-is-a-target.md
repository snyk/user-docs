# What is a target?

Definition of target and scope

<img src="../../../.gitbook/assets/snyk-api-web/fba47833324c7dd805705ec6734b332a5ec9fed0.png" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Nuno Loureiro avatar" />

<span class="text-body-secondary-color"></span>

A Target is the URL of a Web Application, Website, or API. All the following are examples of Targets:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  <a href="https://app1.example.com" rel="nofollow noopener noreferrer" target="_blank">https://app1.example.com</a>

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  <a href="https://example.com/app1" rel="nofollow noopener noreferrer" target="_blank">https://example.com/app1</a>

  </div>

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  <a href="https://example.com" rel="nofollow noopener noreferrer" target="_blank">https://example.com</a>

  </div>

The Target defines the scope of the scan. The scanner will never leave its scope, i.e., it will never scan any page that is not prefixed with the Target's base URL.

If the Target is <a href="https://example.com" rel="nofollow noopener noreferrer" target="_blank">https://example.com</a>, the scanner will not scan, for instance, <a href="https://www.example.com" rel="nofollow noopener noreferrer" target="_blank">https://www.example.com</a> or any other hosts. In other words, the scanner will only scan URLs prefixed by example.com.

You can also think of a Target as to how you want to organize your security testing. Imagine you have a big app at <a href="https://example.com" rel="nofollow noopener noreferrer" target="_blank">https://example.com</a>. This app includes different sections or modules that can even be built by different teams. You can split it into different Targets to facilitate your workflows, like <a href="https://example.com/sectionA" rel="nofollow noopener noreferrer" target="_blank">https://example.com/sectionA</a> and <a href="https://example.com/sectionB" rel="nofollow noopener noreferrer" target="_blank">https://example.com/sectionB</a>.

Watch this video for an overall view of Targets.

# A apărut o eroare.

Nu se poate executa JavaScript.

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F1974191-what-is-a-target&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

