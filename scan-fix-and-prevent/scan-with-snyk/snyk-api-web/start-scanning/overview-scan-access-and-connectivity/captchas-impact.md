---
description: How CAPTCHAs affect Snyk API and Web scans
nav_context: agnostic
---

# CAPTCHAs impact

A CAPTCHA, short for "Completely Automated Public Turing test to tell Computers and Humans Apart," is a challenge-response test used to confirm whether the user is human.

The primary goal of a CAPTCHA challenge is to ensure that a human initiates all actions taken on a website. CAPTCHAs typically accomplish this through letter and image recognition.

Implementing a CAPTCHA on a website helps protect against bots and spam. Snyk recommends a CAPTCHA for use in production.

## Helping Snyk API & Web bypass a CAPTCHA

When using Snyk, allow full access to the target you want to scan.

The crawler, an automated module that crawls the target and gathers endpoints and injection points for the scanner, cannot solve a CAPTCHA independently. As a result, a CAPTCHA can block the crawling of a target and hinder the scanning process.

To help Snyk bypass a CAPTCHA, you can:

* Disable the challenge on a target in a staging environment
* Allow Snyk outbound IP addresses to bypass the CAPTCHA in its settings

If these options do not work, add a cookie that disables the CAPTCHA when presented in a request, allowing Snyk to bypass it.

Implement the cookie on your target. The cookie checks whether the requests use the configured header, then disables the CAPTCHAs and gives clear access to the target.

Example cookie snippet:

```
if (isset($_COOKIE['disable_captcha']) && $_COOKIE['disable_captcha'] === 'aaaaaffff5557777777cccccc99812333ab') {}
```

{% hint style="warning" %}
Do not use the value of the cookie in this example. Generate your own secure value.
{% endhint %}
