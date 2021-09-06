# Monitor your projects at regular intervals

With test and protect, you’re well set up to address currently known vulnerabilities. However, new vulnerabilities are constantly disclosed - which is where monitor comes in.

cd ~/projects/myproject/snyk monitor

Just before you deploy, run `snyk monitor` in your project directory. This takes a snapshot of your current dependencies, so we can notify you about newly disclosed vulnerabilities in them, or when a previously unavailable patch or upgrade path is created. If you take multiple snapshots of the same project, we will only alert you to new information about the latest one.

Log in and go to [snyk.io/monitor](https://app.snyk.io/monitor/) to see the latest snapshot and history of your project.

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}

