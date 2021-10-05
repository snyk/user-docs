# Configure env vars and commit changes

## Change the environment variables to support your Registry

In the SPC interface select github/workflow folder and main.yml file.

![](https://github.com/snyk/user-docs/tree/695c746d1b207ffdf923b84e4590d31b29e2cc73/docs/.gitbook/assets/env_var_change.png)

Update the environment variables at the top of the workflow to your Docker Hub registry/user Id.

![](https://github.com/snyk/user-docs/tree/695c746d1b207ffdf923b84e4590d31b29e2cc73/docs/.gitbook/assets/screen-shot-2020-08-25-at-3.35.08-pm.png)

Commit changes and validate GitHub action workflow has started.

![](https://github.com/snyk/user-docs/tree/695c746d1b207ffdf923b84e4590d31b29e2cc73/docs/.gitbook/assets/actions_running_purple_circle.png)

{% hint style="info" %}
This pipeline will take about 5 minutes to execute as it builds and tests SPC.
{% endhint %}

