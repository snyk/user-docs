# Required packages missing when testing a Python project with pip

## Required packages missing when testing a Python project with pip

**Problem**:

When testing a python project with the Snyk CLI, you can sometimes get the following error:

```text
Required packages missing: package1, package2, package3, etc.
  Please run `pip install -r requirements.txt`. If the issue persists try again with --allow-missing.
```

**Resolution**:

First, you should make sure that you have installed your Python packages by running the command suggested in the error \(adjusting the filename as necessary\):

`pip install -r requirements.txt`

If the issue persists, then you can bypass missing python packages by passing the --allow-missing pip parameter through snyk by using the additional -- argument. For example:

```text
snyk test --org=org --file=requirements.txt -- --allow-missing
```

