# CAPTCHAs impact

A CAPTCHA, short for "Completely Automated Public Turing test to tell Computers and Humans Apart," is a challenge-response test used to confirm whether the user is human.

The primary goal of a CAPTCHA challenge is to ensure that a human initiates all actions taken on a website. This is typically accomplished through letter and image recognition.

Implementing a CAPTCHA on a website can help protect against bots and spam, and it is highly recommended for use in production.

## Helping Snyk API & Web bypass a CAPTCHA

When using Snyk API & Web, it is important to allow full access to the target you wish to scan.

The crawler, an automated module that crawls the target and gathers endpoints and injection points for the scanner, cannot solve a CAPTCHA independently. As a result, a CAPTCHA can block the crawling of a target and hinder the scanning process.

To help Snyk API & Web bypass a CAPTCHA, you can:

* Disable the challenge on a target in a staging environment
* Allow Snyk API & Web outbound IP addresses to bypass the CAPTCHA in its settings

If these options do not work, add a cookie that, when presented in a request, will disable the CAPTCHA and allow Snyk API & Web to bypass it successfully.

The cookie (implemented on your target) will check if the requests are being made using the configured header, disabling the CAPTCHAs and giving clear access to the target.

Example cookie snippet:

```
if (isset($_COOKIE['disable_captcha']) && $_COOKIE['disable_captcha'] === 'aaaaaffff5557777777cccccc99812333ab') {}
```

{% hint style="warning" %}
Do not use the value of the cookie in this example. Generate your own secure value.
{% endhint %}
