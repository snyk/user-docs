# CLI support for Python

## Set the Python version in the CLI

To set the Python version in the CLI, add the following option to `snyk test` or `snyk monitor` with the name of the Python binary:

```sh
--command=python3
```

For details, see the options for Python Projects in the [`snyk test`](../../../developer-tools/snyk-cli/commands/test.md) and [`snyk monitor`](../../../developer-tools/snyk-cli/commands/monitor.md) help.

## Pip and  CLI

{% hint style="info" %}
Run `pip install` before scanning with the CLI, for example:

```
pip install -r requirements.txt
```
{% endhint %}

Pip `requirements.txt` files specify only top-level dependencies, not nested or transitive ones. Therefore, the full Pip Project must be installed to ensure the CLI can build a complete dependency tree.

## Poetry and CLI

{% hint style="info" %}
[PEP 621](https://peps.python.org/pep-0621/) is a standard for defining direct dependencies in `pyproject.toml` files. Snyk supports PEP 621 only for Poetry v2. Poetry v1 uses an alternative approach.

Snyk does not support PEP 621 for any package manager other than Poetry.
{% endhint %}

Poetry v1 and v2 are supported.

To build the dependency tree for a Poetry application, Snyk uses `pyproject.toml` and `poetry.lock` files. Both files must be present for Snyk to scan Poetry dependencies and identify issues.

If no `poetry.lock` file is present; you should run `poetry lock` to generate one before scanning.

## Pipenv and CLI

To build the dependency tree for a Pipenv application, Snyk uses `Pipfile` and `Pipfile.lock` files. Both files must be present for Snyk to scan Pipenv dependencies and identify issues.

Run `pip install` before scanning with the CLI.

Run `pipenv install` to ensure the CLI can build an up-to-date, accurate dependency tree using `pipenv graph`.

## setup.py and CLI

To build the dependency tree, Snyk analyzes the `setup.py` file, and detects packages listed in the `install_requires` key.

This file will not be discovered automatically by the CLI. It must be specified manually using the  `--file` option, for example:

```python
snyk test --file=setup.py
```

You can also convert `setup.py` to `requirements.txt` by installing the packages into a virtual environment and then running `pip freeze`.
