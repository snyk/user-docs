# Review reporting

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

## Reporting

The Snyk platform has a variety of [available reports](../../../../manage-risk/analytics/reports-tab/) at the Organization and Group level to provide deeper insights. This guide covers a few of the key reports and capabilities.

## Issues Detail

The Issues Detail report displays all known issues in all of your Projects that Snyk is monitoring. It is a great place to review and prioritize issues across all of your repositories.

Applying filters such as fixability, exploit maturity, asset class, and severity are a great way to focus on the most important issues that affect your Organization.

## SLA Management

The SLA Management report allows you to stay on top of any SLAs your company is required to adhere to. The SLA time can be adjusted for each severity and then saved as a new view to monitor in the future.

## Analytics

[Analytics](../../../../manage-risk/analytics/) can be accessed at the Group or Organization level. This view breaks down the total issue count into baseline, preventable, and non-preventable issues. The most important thing to know is that:

* Baseline issues are the issues identified during the first import of the repository.
* Preventable issues are known issues that could have been identified earlier on in the SDLC. In other words, the issue could have been caught by a Snyk PR check.
* Non-preventable issues are the result of an external factor, such as a new vulnerability being published or a new security rule being created, in contrast to developers not shifting left. for example, a new zero-day vulnerability is identified.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXco8q4gfqWC3QdLqOY8N15kchfGb9_FKm28rXonSWmSbOnrTDIQpvQMluxOoiOBWQStylL_LKasaU7VhjbjkRzv0UIQ60UqKtX3yTwTO-XO1gz7tgiWQc2COU-frmYkUXl5FQAM?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

This is a small taste of what reporting has to offer. Check out more of the Snyk reports, available at both the Organization and Group levels by clicking “Change Report”:\
![](https://lh7-rt.googleusercontent.com/docsz/AD_4nXfS5UzFsGMM5_N5fK6iFLN16rFFfSmj3W9BXkmDnZOvvOBoUjCIQD6j1afOaN9PySsB-MI4TNLtKdgFbVk1OMe5u1uCYDSKv1pjhkaUSqhGspmGqOggsPx5XCK7IZGVv7QQmN5NqQ?key=i_CNrr-DvB8PGUAzq09BT3pc)

{% hint style="info" %}
Additional Resources

* [Docs: Analytics](../../../../manage-risk/analytics/)
* [Docs: Reports](../../../../manage-risk/analytics/reports-tab/)
* [Product Training: Reports](https://learn.snyk.io/catalog/?type=product-training\&topics=Reporting)
{% endhint %}
