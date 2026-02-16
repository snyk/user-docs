# How to use Seeds and Reject lists

Learn how to use Seeds and Reject lists to optimize target scans.

<img src="../../../.gitbook/assets/snyk-api-web/5530a2b28285ea60de10f16a64f57b44160bec88.jpg" class="inline-flex items-center justify-center rounded-full bg-primary text-lg font-bold leading-6 text-primary-text shadow-solid-2 shadow-body-bg [&amp;:nth-child(n+2)]:hidden lg:[&amp;:nth-child(n+2)]:inline-flex h-8 w-8 sm:h-9 sm:w-9" width="24" height="24" alt="Martim Valente avatar" />

<span class="text-body-secondary-color"></span>

The Seeds and Reject lists allow you to control the behavior of your target scans by including or excluding areas of the target from scans.

For example, if an area in the target can’t be reached in scans (e.g., an admin area that is only reached through a direct link), add it to the Seeds list to be included in scans. On the other hand, if there’s a critical area being scanned in the target (e.g., a form that sends e-mails to the users) and you don’t want it to be tested, you can add it to the Reject list to exclude it from scanning.

The Seeds list allows you to add areas of the target that otherwise would be “hidden” and not caught by scans. This helps to ensure that scans cover everything you need in the target’s scope.

The Reject List allows you to limit what we visit on the target. It's a collection of URLs that the crawler is instructed to avoid. This helps to exclude unnecessary or sensitive content, such as login pages and logout options, private user data, or repetitive structures like paginated lists.

<table>
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td style="background-color: #fff2cc"><div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">
<p>Keep in mind that the Reject list takes <strong>precedence</strong> over the Seeds list: areas in the Seeds list will be ignored if they match areas in the Reject list.</p>
</div></td>
</tr>
</tbody>
</table>

# Using the Seeds List

Add the area you want to include in the target scans to the Seeds list as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the <a href="https://plus.probely.app/" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web app</a>, go to the **Targets** entry.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Find the target in the list and click on the cogwheel to edit its settings.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the **SCANNER** tab, go to the **SEEDS LIST** section, type the path to the area to include in scans in **ADD PATH**, and click on **Add**.\
    You can add more than one path, and they must be relative to the target URL.

    </div>

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1431911539/b675ae28bc3ce7740644c1c8df6d/Screenshot+2025-03-20+at+11_48_16.png?expires=1769985000&amp;signature=4b86435470018697ca66b6cb026a58883e7eed7d2f6933417cd58a65ea01f51a&amp;req=dSQkF8B%2FnIRcUPMW1HO4zbTrzmr5TkkYH%2F2C19MJLhmGfIxZNSlFYjcL%2Fu56%0ArTcztGdgheYek0lDsaI%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/7b44e4ad9c4b4271ab8462212f966b4fd08edce7.png" width="1256" height="902" /></a>

Scans will follow these paths and explore them, thus expanding the reach across the target.\
​

# Using the Reject List

Add the area you want to exclude from the target scans to the Reject list as follows:

1.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the <a href="https://plus.probely.app/" rel="nofollow noopener noreferrer" target="_blank">Snyk API &amp; Web app</a>, go to the **TARGETS** tab.

    </div>

2.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    Find the target in the list and click on the cogwheel to edit its settings.

    </div>

3.  <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

    In the **SCANNER** tab, go to the **REJECT LIST** section, type the area to exclude from scans in **ADD URL**, and click on **Add**.\
    The URLs you type must be absolute, and you can use **\*** as a wildcard.

    </div>

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1431911718/dcafa7b3fcaa2eb9b0b7af4ff70a/Screenshot+2025-03-20+at+11_48_08.png?expires=1769985000&amp;signature=6f23ef06296226e0ef5c0e74ef19b89e9a8d3f5d2d4242604de4aeb1e83f2a85&amp;req=dSQkF8B%2FnIZeUfMW1HO4zcVbYsunnAUqnRprb4mtrFpUdI6gwnrlhgx9KnaY%0AIqB6McauOezyD6wq9zA%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/89b86fad7cf69679a1bedc4e07137a57131400bc.png" width="1254" height="976" /></a>

Scans will not follow these paths, thus avoiding these areas in the target scans.

Remember, you can always delete entries in these lists as your scanning needs evolve. To remove an entry, simply click the trash icon.

<a href="https://downloads.intercomcdn.com/i/o/r3ylwg3q/1431912620/9a5a74fc51ef48d10c8ca8ab4932/Screenshot+2025-03-20+at+11_49_16.png?expires=1769985000&amp;signature=fbc192e9e617e3cd318500fc5d80b932e49fc5290004a8afc17d86eb42b8db53&amp;req=dSQkF8B%2Fn4ddWfMW1HO4zeFONifoqIntpxLiHO%2BgbmhpcpCZhG38BiuQpwSd%0ASDOkLY1kVT2Kx57AdRY%3D%0A" target="_blank" rel="noreferrer nofollow noopener"><img src="../../../.gitbook/assets/snyk-api-web/f6d8e99ee505e9ae0da128157eea579233e3393d.png" width="1112" height="146" /></a>

This tailored approach ensures that “hidden” areas are thoroughly tested while sensitive or unnecessary areas are excluded, which helps optimize your scanning process.

For further details or to explore related topics, please check out the links below:

- <div class="intercom-interblocks-paragraph no-margin intercom-interblocks-align-left">

  [How to add a target](https://help.probely.com/en/articles/5733114-how-to-add-a-target)

  </div>

Did this answer your question?

<span title="Disappointed">😞</span>

<span title="Neutral">😐</span>

<span title="Smiley">😃</span>

<a href="/en/" class="no-underline"><img src="../../../.gitbook/assets/snyk-api-web/a15a09243c0347906a8894e098b4c169a184c5b1.png" class="max-h-8 contrast-80 inline" data-testid="logo-img" alt="Snyk API &amp; Web Help Center" /></a>

- <a href="https://www.facebook.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-0"><img src="../../../.gitbook/assets/snyk-api-web/c8736e4703de6bdf91ff6e530d9eb32d874a05cb.svg" loading="lazy" data-testid="social-icon-facebook" width="16" height="16" /></a>
- <a href="https://www.twitter.com/snyksec" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-1"><img src="../../../.gitbook/assets/snyk-api-web/2ef8a64440ad5e13da982461b8f8ed37a7c21515.svg" loading="lazy" data-testid="social-icon-twitter" width="16" height="16" /></a>
- <a href="https://www.linkedin.com/company/snyk" class="no-underline" target="_blank" rel="nofollow noreferrer noopener" data-testid="footer-social-link-2"><img src="../../../.gitbook/assets/snyk-api-web/b327bf9a16233fc3f3268ea536d10d4e884df405.svg" loading="lazy" data-testid="social-icon-linkedin" width="16" height="16" /></a>

![](../../../.gitbook/assets/snyk-api-web/43570513a16202d124d06a905897b6aef7146ecc.svg)<a href="https://www.intercom.com/intercom-link?company=Snyk+API+%26+Web&amp;solution=customer-support&amp;utm_campaign=intercom-link&amp;utm_content=We+run+on+Intercom&amp;utm_medium=help-center&amp;utm_referrer=https%3A%2F%2Fhelp.probely.com%2Fen%2Farticles%2F9777912-how-to-use-seeds-and-reject-lists&amp;utm_source=desktop-web" class="pl-2 align-middle no-underline">We run on Intercom</a>

