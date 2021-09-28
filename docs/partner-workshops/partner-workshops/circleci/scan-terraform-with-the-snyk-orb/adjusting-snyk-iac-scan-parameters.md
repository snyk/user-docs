---
description: >-
  After adding Snyk IAC to your workflow, you might want to change the
  sensitivity of the scan or the severity assigned to the rules. This section
  shows ways to accomplish that.
---

# Adjusting Snyk IAC Scan parameters

## Adjusting CLI Severity Threshold

One way to change the pass/fail criteria for the Snyk Orb is to pass the `--severity-threshold` argument. This allows you to choose what severity misconfiguration needs to be found to break the build.

In this example, we'll tell Snyk to only identify High severity misconfigurations:

```text
workflows:
  build_test:
    jobs:
      - run_tests
      - build_docker_image
      - snyk/scan-iac:
          args: part03/iac_gke_cluster --severity-threshold=high
      - gke_create_cluster:
```

Since the defects found are all Medium and Low severity, the Snyk stage in the workflow will pass.

![Snyk workflow passes with 0 issues found.](../../../.gitbook/assets/image%20%284%29.png)

In this example, we only break the pipeline when High Severity issues are found. 

## Adjusting rule thresholds in the Snyk UI

If you want to override the [default severity thresholds](https://snyk.io/security-rules) provided by Snyk, you can do so in the Snyk UI. This allows you to change the severity of each of the IAC security rules at an organization level.

To do this, navigate to your Organization -&gt; Settings -&gt; Infrastructure as Code. You'll see the list below:

![Security Rule severity adjustments in the Snyk UI.](../../../.gitbook/assets/image%20%283%29.png)

Adjust the rules as you see fit. Next time the workflow runs, the Snyk Orb will evaluate the Terraform files according to the new severity levels. 

### Configuring the Snyk Organization used by the Snyk Orb

The Snyk Orb output identifies which Snyk Org is used for evaluating security rules.

![Snyk Org used by the Orb](../../../.gitbook/assets/image%20%281%29.png)

If you want to use the rules in a different org, pass the `--org` parameter to the `args`. 

{% hint style="info" %}
The Snyk Orb evaluates severity levels for the Org for which the SNYK\_TOKEN was created. If you specify a different Org, the token must have access to that Org.
{% endhint %}

```text
workflows:
  build_test:
    jobs:
      - run_tests
      - build_docker_image
      - snyk/scan-iac:
          args: part03/iac_gke_cluster --severity-threshold=high --org=neworg
      - gke_create_cluster:
```

This allows you to have different rulesets for different environments where your applications run. 

