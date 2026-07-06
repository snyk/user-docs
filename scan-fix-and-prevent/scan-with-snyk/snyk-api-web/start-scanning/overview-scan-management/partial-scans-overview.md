# Run partial scans

To test only a subset of your Web target, run partial scans in Snyk API & Web.

Partial scans are helpful, especially on a continuous integration/continuous delivery (CI/CD) pipeline, because they provide faster feedback. They let you deliver code changes more frequently and reliably, without disregarding the security of your target.

You can run partial scans in different ways:

* By defining a reduced scope in your target.
* By setting up navigation sequences within your target.
* By enabling incremental scans.
* By using the API to specify scope on each scan request.

## Reduced scope

To set up the reduced scope, navigate to your target settings, open the **Scanner** tab, and locate the **PARTIAL SCANS: REDUCED SCOPE** section. Add the URLs you want Snyk to analyze during partial scans. Only the target hostname and defined extra hosts are allowed.

Assuming your target's URL is `https://example.com/`, some examples of possible reduced scopes are as follows:

* `https://example.com/admin*`
* `https://api.example.com/api/users*`
* `https://example.com/account*`
* `https://example.com/users/*/edit`
* `https://example.com/pages/users.php*`
* `https://auth-api.example.com/auth*`

{% hint style="info" %}
Add the wildcard character `*` so that all pages under that scope are analyzed and scanned as well. Otherwise, Snyk scans only the file or path itself.
{% endhint %}

You can add as many URLs as you need to define the intended reduced scope. If you have an API residing on a different hostname than your target (Extra Host), you can also add it here.

When you limit the scope of the scan, the crawler might not find some of the defined sub-scopes (URLs). The crawler always visits the root of your target to find possible valid endpoints. To test endpoints that are not accessible through your target root, navigate to your target settings, open the **Scanner** tab, locate the **SEEDS LIST** section, and add the URLs where those endpoints are present.

For example, to test the endpoint `https://example.com/users/*/edit`, which allows you to edit your users' information, add the list of users' URL, `https://example.com/users`, as a seed on your seeds list.

## Navigation sequences

Another way to run partial scans is to define navigation sequences. You can set your scans to run only navigation sequences.

To do this, navigate to your target settings, open the **Scanner** tab, and select the appropriate check boxes under the **NAVIGATION SEQUENCES** section: **On demand scans must only run navigation sequences** and **Scheduled scans must only run navigation sequences**.

During a navigation-sequences-only scan, the crawler navigates solely through the selected sequences, running their recorded actions and subsequent requests, instead of analyzing the whole target. In other words, during these scans, the crawler analyzes all requests intercepted during the sequence. To reduce the scope, complement this setting with the reduced scope setting described earlier.

## Incremental scans

Besides using the reduced scope and navigation sequences to narrow the scope of scans, you can also enable incremental scans. Incremental scans limit your scans to new URLs (pages that have not been scanned before) and updated URLs (pages that have changed since the previous scan).

Incremental scans help you understand the impact of new developments or changes made to your target, because they provide fast and meaningful feedback.

To enable incremental scans:

1. Navigate to your target settings and open the **Scanner** tab.
2. Locate the **PARTIAL SCANS: INCREMENTAL** section.
3. Activate this feature.

## Using the API for partial scans

The Snyk API & Web API lets you run partial scans on your targets without permanently configuring the reduced scope in target settings.

When you call the endpoint to start a vulnerability scan, you can send the reduced scope you want Snyk to scan in your start scan request. This approach is useful for running different partial scans dynamically without changing target configuration each time.

For more information about using the API, visit the [Snyk API & Web API documentation](https://developers.probely.com/api/reference/targets-scan-now-create/).
