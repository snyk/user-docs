# Testing Helm Charts with our CLI tool

You scan a Helm chart by rendering the Helm templates into Kuberenetes manifest files and then scanning these using the `snyk iac` command in the Snyk CLI.

For example if you have a Helm project located in a `./helm` directory you would run the following command to output the templated files into a directory called `./output`:

{% tabs %}
{% tab title="macOS/Linux/Unix" %}
```text
helm template ./helm --output-dir ./output
snyk iac test ./output
```
{% endtab %}

{% tab title="Windows PowerShell" %}
```
helm template .\helm\ --output-dir .\output\
snyk iac test .\output\
```
{% endtab %}
{% endtabs %}

For unix based terminals you can also pipe the output of `helm template` directly into a single file:

```text
helm template ./helm > output.yaml
snyk iac test output.yaml
```

Currently the Snyk CLI is not able to read from standard input but this is a feature we are currently working on.

