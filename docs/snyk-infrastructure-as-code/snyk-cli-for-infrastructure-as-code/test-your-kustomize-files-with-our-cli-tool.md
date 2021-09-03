# Test your Kustomize files with our CLI tool

You scan a Kustomize template by building the Kuberenetes manifest file and then scanning this using the Snyk IaC CLI. 

```text
kustomize build > kubernetes.yaml
snyk iac test kubernetes.yaml
```

Depending on your kustomize templates, you may need to provide a name after the build argument. 

 
<br><br><hr>

{% hint style="success" %}
Ready to get started with Snyk? [Sign up for free!](https://snyk.io/login?cta=sign-up&loc=footer&page=support_docs_page)
{% endhint %}