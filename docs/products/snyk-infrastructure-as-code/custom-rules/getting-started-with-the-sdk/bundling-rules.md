# Bundling rules

Once you are ready, you can build a custom rules bundle by running the following command:

```text
snyk-iac-rules build
```

If you have more than your generated rules in the current folder consider using the  `--ignore` flag to exclude the folders and files irrelevant for a production-ready bundle. This can both speed up the process and ensures the size of the generated bundle stays small.

#### Overriding the default entry point

If you have chosen to name the rule that evaluates something different than **`deny`** \(e.g. `allow`,`violation`, etc. \), you can override it by running:

```text
snyk-iac-rules build --entrypoint "<package name>/<function name>"
```

Finally, you can check the contents of the bundle without extracting it by running:

```text
tar -tf bundle.tar.gz
```

That will output all the files included in the bundle:

```text
/data.json
/lib/main.rego
/rules/my_rule/main.rego
/policy.wasm
/.manifest
```

 You can now [run snyk iac test with your newly built custom bundle. ](../how-to-run-custom-rules-with-the-snyk-cli.md)

