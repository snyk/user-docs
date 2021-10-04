# Snyk Test using CLI

## Testing your SPC applications

From the root directory of the SPC application,  execute the Synk command below. The Snyk test command will test your dependencies for vulnerabilities and tell you how many vulnerabilities are found.

```text
snyk test 
```

{% hint style="info" %}
Synk CLI offers a number of switches for specific use cases and formatting of output.
{% endhint %}

{% hint style="info" %}
Since we did not build the application the scan will take longer. This will take ~2 mins. If we packaged the application the scan will take 30 seconds. 
{% endhint %}

### Results of running Snyk test

Snyk test displays a list of vulnerabilities grouped by fixable and non-fixable issues followed by license compliance issues. Take a minute to review the output, similar to the results below. Snyk provides remediation advice for fixable vulnerabilities.

```text
Tested 83 dependencies for known issues, found 22 issues, 22 vulnerable paths.


Issues to fix by upgrading:

  Upgrade org.webjars:bootstrap@3.3.6 to org.webjars:bootstrap@3.4.1 to fix
  ✗ Cross-site Scripting (XSS) [Medium Severity][https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-451160] in org.webjars:bootstrap@3.3.6
  ✗ Cross-site Scripting (XSS) [Medium Severity][https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-451162] in org.webjars:bootstrap@3.3.6
  ✗ Cross-site Scripting (XSS) [Medium Severity][https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-451164] in org.webjars:bootstrap@3.3.6
  ✗ Cross-site Scripting (XSS) [Medium Severity][https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-451168] in org.webjars:bootstrap@3.3.6
  ✗ Cross-site Scripting (XSS) [Medium Severity][https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-479505] in org.webjars:bootstrap@3.3.6
```

## Additional Examples of using Synk Test

### Testing a Container image

Tests a local container image. If the image is not available locally, Snyk will try to pull it from Docker Hub.

```text
snyk test --docker openjdk:8
```

{% hint style="info" %}
This scan takes about 30 seconds as the image is downloaded.
{% endhint %}

### Test public repositories before use

To test a public Github, BitBucket or GitLab repository, run `snyk test` and include the Github URL to the repo.

```text
snyk test https://github.com/snyk/goof
```

The following git URL formats are supported:

* `git://github.com/user/project.git#commit-ish`
* `https://github.com/user/project#commit-ish`
* `user/project#commit-ish`

