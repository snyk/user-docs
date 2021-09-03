# Test your Helm Charts with our CLI tool

You scan a Helm chart by converting the template to a rendered Kuberenetes manifest file and then scanning this using the Snyk IaC CLI. 

```text
helm template ./iac-helm > helm.yaml
snyk iac test helm.yaml
```

change \`iac-helm\` for your Helm chart name. 

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}