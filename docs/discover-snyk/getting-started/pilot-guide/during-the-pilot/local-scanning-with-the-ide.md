# Local scanning with the IDE

{% include "../../../../.gitbook/includes/pilot-guide-toc.md" %}

## Configure the Snyk IDE

To start scanning code in the IDE, navigate to your IDE plugin or extension marketplace and search for Snyk. This guide focuses on VS Code, but the [other supported IDEs](../../../../developer-tools/snyk-ide-plugins-and-extensions/) follow a very similar setup flow.

* Open VS Code, Extensions and search for Snyk.
* Install the Snyk extension.

<figure><img src="../../../../.gitbook/assets/image (313).png" alt=""><figcaption></figcaption></figure>

* Select **Connect & Trust Workspace** to let Snyk scan your code.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXesxK_N2_rft7boTn2XoMiTvyyi9VClXBZ64kU4Jprojyu3C3mYYaQlpamq5PWDULrUFsG2MyOFXezWWUn3_4oMlF__xu_4PjVJEF7lQQ5jFc2sqT3NhZt5bheUJOqpajKXGqpJmA?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

* After being redirected to Snyk, you can click **Grant app access**. This authenticates your Snyk IDE extension.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXd9acXhko99qCN0uD5gaOu7N_sEDxHeQKn8GVdmfcMFLJ4IUs5Y0BVaYtHdPQcwqaJW63iEVEZ-37wl7DBrkUT_zsjUBv3Ar-loSvnhjm0tuV5ay9qF1_83iPDNV50pj3caXg22pg?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

* Return to your IDE, where you can start scanning immediately.&#x20;

## Use the Snyk IDE

### Read the results

You can see the results are broken out by scan type. All the Open Source issues are grouped together, followed by Code, then Infrastructure as Code.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfEwAZTkQDih4_j2v7UqXQeLV2OlGXmgEAzPi3TOGojBAsUwhUuRPU09SKOECs4ho7vamwk-0P-2eeR3o_y3_R2G8tRrLxxn-0SsYHWqVfFYA35QvqqDlkaBVKZolcoCJ3TTWfj?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

### Filter issues

Issues can be filtered in a number of ways to tune what gets presented to you while coding. Two of the most common settings to adjust are ‌**Severity** and **Total vs. new**.

To adjust which severity is shown, navigate to the **Severity** section in the extension settings.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfrWs66nt4A0swZ1i7tgWaXzdevqIG3sUXhUnp5Ac2AyQP6dF678QBLoo3_o64Tcp7KuhT6wFKJbhnEh5ijdIgb7DPCIP04p9-4Dx0uzu6feCM5F7Wa8G6PfL2v3kqymqP9tzhr?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

You can also change from looking at the total set of issues in the codebase to just new issues. This method scans the main branch and the version that the developer is working on, then reports only the new vulnerabilities. This mode can be easily toggled at the top of the extension menu.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXcyxcHyZ-CvP-Ri9F4pHCRKGCRpzgABhB06WozPHT3w3J_BTdUFBeX5rzJajhKNIQe0u3zYay2MDPT_LlAXco-nNq3akF0ASA7drSGhM7Yzcc_6KNEGkWvMD5lbj02T3hOlaTGH?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

This option helps developers avoid introducing new vulnerabilities without getting overwhelmed by their backlog.

### Check the vulnerabilities

The results for each scan type and vulnerability type vary, but expect to see inline feedback as well as a more detailed window that can be accessed by clicking on the vulnerability.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXc-IFnmluEMchIOAgwW7-eDTzIV_k4M8XCd3ss4athOzyoLi4MBTqRgXEJhDh6y-0mH6B7hrCb6kaeUnfFH0ynFtFTjcq3Pm-Pf2I1okriX3jBg5spJ8IzDqMhuLwJPGccn2AZhOA?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

For Snyk Code issues, you can see the Fix Analysis, Data Flow, and an Issue Overview. You can also click “Learn about this issue type” which brings you to a more detailed lesson in Snyk’s developer education platform, [Snyk Learn](https://learn.snyk.io/).

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXfp_x_fKA8lVV_wmvtw4eXnn1p47MmbQsXrO92GgLeVyw5PiPa4fqAp1fjIfIdBrTP5jeYGkY3P26vMMt5gCkVbepgKp27b6hBBu0RFSorkVs-E_niw8qKsx9M9ASJBtqXRBesb7g?key=i_CNrr-DvB8PGUAzq09BT3pc" alt=""><figcaption></figcaption></figure>

The next step in the guide covers fixing issues.

{% hint style="info" %}


Additional Resources

* [Product Training: Snyk in the IDE](https://learn.snyk.io/lesson/snyk-in-an-ide)
* [Snyk Docs: Snyk IDE plugins and extensions](../../../../developer-tools/snyk-ide-plugins-and-extensions/)
* [Troubleshooting](../../../../developer-tools/snyk-ide-plugins-and-extensions/troubleshooting-ides/)
{% endhint %}
