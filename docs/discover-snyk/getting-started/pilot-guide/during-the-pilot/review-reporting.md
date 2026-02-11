# Review reporting

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

## Reporting

The Snyk platform has a variety of [available reports](/broken/pages/Xet9BNiX2RviFlrTiatG) at the Organization and Group level, and [issue analytics](/broken/pages/1m1GoANo0CromJXOD7H1) to provide deeper insights at the tenant level. This guide covers a few of the key reports and capabilities.

## Issues Detail

The Issues Detail report displays all known issues in all of your Projects that Snyk is monitoring. It is a great place to review and prioritize issues across all of your repositories.

Applying filters such as fixability, exploit maturity, asset class, and severity are a great way to focus on the most important issues that affect your Organization.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXemxx2Nk4j5shNT04dsvthA0gPJKtsaCLEJZ3DE6csgid6FEOXV09BYykl-yP-Mpi9DOSivwidOwJdTAlWmJyTMfeKY9mmMS1xYY2QyJ-_548iUzQlmxpeSrT1P1VaGYqIGrRnUug?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

Try saving a view to come back to it later and keep all of the same filters applied:

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXf6HMW1Ld-T-lbX_rlij-vNILvjzLavoz6C8sosqgZ3uk8SIMp6bZX6_rXjMjg87Is11VTpP_Dzq8qtHD_xhC3rfuPBrQegKBpQqWN8xHdd0E4VsTIvJR5jsoUBs7-V20LSEboWqg?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

![](https://lh7-rt.googleusercontent.com/docsz/AD_4nXd1eLOowyZLyMAy73FjTE1LMw-Z0kE9M8uzTRhHWSkOEAKLstCkD6RgmmwEP6DBwSBD_EWUJGiF-QiF66V9xCwUaxkygcydGHOzHvR7WBgLWakvLbTyNyODLj4_Y8zkbs_JdVrfwg?key=i_CNrr-DvB8PGUAzq09BT3pc)

## SLA Management

The SLA Management report allows you to stay on top of any SLAs your company is required to adhere to. The SLA time can be adjusted for each severity and then saved as a new view to monitor in the future.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXeuN_Zmk6Ve1RThtILreE0_qPrFp15WsztUenS9bILnjgBkpa1t9pO6f8UXtszyYPFm25IJA0uOXJEcbqMicRtjsiBvl58olqEc5bBooaNPk8bssObP73uFe-oQi7lP4hBEuFKP?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

## Issues Analytics

Analytics can be accessed at the tenant level. This view breaks down the total issue count into baseline, preventable, and non-preventable issues. More details on this breakdown can be found [here](/broken/pages/1m1GoANo0CromJXOD7H1#delineation-of-how-risk-is-introduced), but the most important thing to know is that:

* Baseline issues are the issues identified during the first import of the repository.
* Preventable issues are known issues that could have been identified earlier on in the SDLC. In other words, the issue could have been caught by a Snyk PR check.
* Non-preventable issues are the result of an external factor, such as a new vulnerability being published or a new security rule being created, in contrast to developers not shifting left. E.g., a new zero-day vulnerability is identified.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXco8q4gfqWC3QdLqOY8N15kchfGb9_FKm28rXonSWmSbOnrTDIQpvQMluxOoiOBWQStylL_LKasaU7VhjbjkRzv0UIQ60UqKtX3yTwTO-XO1gz7tgiWQc2COU-frmYkUXl5FQAM?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

This is a small taste of what reporting has to offer. Check out more of the Snyk reports, available at both the org and group levels by clicking “Change Report”:\
![](https://lh7-rt.googleusercontent.com/docsz/AD_4nXfS5UzFsGMM5_N5fK6iFLN16rFFfSmj3W9BXkmDnZOvvOBoUjCIQD6j1afOaN9PySsB-MI4TNLtKdgFbVk1OMe5u1uCYDSKv1pjhkaUSqhGspmGqOggsPx5XCK7IZGVv7QQmN5NqQ?key=i_CNrr-DvB8PGUAzq09BT3pc)

{% hint style="info" %}
Additional Resources

* [Docs: Analytics](../../../../manage-risk/analytics/)
* [Docs: Reports](/broken/pages/ESO9SlllNm4KuC5roygB)
* [Product Training: Reports](https://learn.snyk.io/catalog/?type=product-training\&topics=Reporting)
{% endhint %}
