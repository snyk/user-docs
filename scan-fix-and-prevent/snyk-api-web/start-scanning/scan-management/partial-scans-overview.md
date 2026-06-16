# Run partial scans in Web targets

If you want to test only a subset of your Web target, Snyk API & Web allows you to run partial scans.

Partial scans can be very helpful, especially on a continuous integration and continuous delivery (CI/CD) pipeline, since they provide faster feedback, allowing you to deliver code changes more frequently and reliably, without disregarding the security of your target.

You can run partial scans in different ways:

* By defining a reduced scope in your target.
* By setting up navigation sequences within your target.
* By enabling incremental scans.
* By using the API to specify scope on each scan request.

## Reduced scope

To set up the reduced scope, go to your target's settings, open the **Scanner** tab, and locate the **PARTIAL SCANS: REDUCED SCOPE** section. Add the URLs you want Snyk API & Web to analyze during partial scans. Note that only the target's hostname and defined extra hosts are allowed.

Assuming your target's URL is `https://example.com/`, some examples of possible reduced scopes are as follows:

* `https://example.com/admin*`
* `https://api.example.com/api/users*`
* `https://example.com/account*`
* `https://example.com/users/*/edit`
* `https://example.com/pages/users.php*`
* `https://auth-api.example.com/auth*`

{% hint style="info" %}
Add the wildcard character `*`, so that all pages under that scope are analyzed and scanned as well. Otherwise, only the file or path itself will be scanned.
{% endhint %}

You can add as many URLs as you need in order to define the intended reduced scope. If you have an API residing on a different hostname than your target (Extra Host), you can also add it here.

By limiting the scope of the scan, it is possible that some of the sub-scopes (URLs) defined are not found. The root of your target will always be visited by the crawler in order to find possible valid endpoints. If you want to test some endpoints that are not accessible through your target's root, go to your target settings, open the **Scanner** tab, locate the **SEEDS LIST** section, and add the URLs where those endpoints are present.

For instance, assuming you want to test the endpoint `https://example.com/users/*/edit`, which allows you to edit your users' information, you may need to add the list of users' URL, `https://example.com/users`, as a seed on your seeds list.

## Navigation sequences

Another way to run partial scans is to define navigation sequences. You can decide that your scans should only run navigation sequences.

This can be done by going to your target's settings, opening the **Scanner** tab, and checking the appropriate checkboxes under the **NAVIGATION SEQUENCES** section: **On demand scans must only run navigation sequences** and **Scheduled scans must only run navigation sequences**.

<figure><img src="../../../../.gitbook/assets/partial-scans-navigation-sequences.png" alt="Navigation Sequences section showing checkboxes to enable navigation sequences only scans"></figure>

During the navigation sequences only scan, the crawler will navigate solely through the selected sequences, running their recorded actions and subsequent requests, instead of analyzing the whole target. In other words, during these scans, all intercepted requests during the sequence will be analyzed. If you need to reduce the scope, you can complement this setting with the reduced scope setting detailed above.

## Incremental scans

Besides using the reduced scope and navigation sequences to narrow down the scope of scans, you can also enable incremental scans. With incremental scans, you are limiting your scans to new URLs (that is, pages that have not been scanned before) and to updated URLs (which are pages that have changed since the previous scan).

Incremental scans are a great way to understand the impact of new developments or changes made to your target, since they provide you with fast and meaningful feedback.

To enable incremental scans:

1. Go to your target's settings and open the **Scanner** tab.
1. Locate the **PARTIAL SCANS: INCREMENTAL** section.
1. Activate this feature.

## Using the API for partial scans

The Snyk API & Web API allows you to run partial scans on your targets without permanently configuring the reduced scope in target settings.

When you call the endpoint to start a vulnerability scan, you can send the reduced scope you want Snyk API & Web to scan in your start scan request. This approach is useful for running different partial scans dynamically without changing target configuration each time.

For more information about using the API, visit the [Snyk API & Web API documentation](https://developers.probely.com/api/reference/targets-scan-now-create/).