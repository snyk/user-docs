# (Recommended) Test your Custom Rule

Use the following CLI command:

```
snyk iac rules test
```

This command tests your Custom Rules by locally running your Custom Rule specs against those rules.

## **Example output**

```
Running specs...
7/7 specs passed.
Running rego tests...
PASS: 1/1
```

## **Update expected output**

Running `snyk iac rules test --update-expected` tests your Custom rRules and creates or overwrites the JSON file in your expected directory.
