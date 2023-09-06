# Snyk for Python

{% hint style="success" %}
Snyk provides security scanning on Projects for vulnerabilities both through the [Snyk CLI ](../../../snyk-cli/)and from the Snyk Web UI through different [Snyk Integrations](../../../integrations/).
{% endhint %}

This page describes how to use Snyk to scan Python Projects.

## Features of Snyk for Python

| Package managers / Features                  | CLI support | Git support    | License scanning | Fix PRs |
| -------------------------------------------- | ----------- | -------------- | ---------------- | ------- |
| [Pip](https://pypi.org/project/pip/)         | ✔︎          | ✔︎             | ✔︎               | ✔︎      |
| [Poetry](https://python-poetry.org)          | ✔︎          | ✔︎             | ✔︎               |         |
| [Pipenv](https://pipenv.pypa.io/en/latest/)  | ✔︎          | ✔︎ (Open beta) | ✔︎               |         |
| setup.py                                     | ✔︎          |                | ✔︎               |         |

{% hint style="info" %}
**Feature availability**\
Features may not be available, depending on your plan. See the [Plans and pricing ](https://snyk.io/plans/) page for more details.
{% endhint %}

## How Snyk for Python works

Snyk builds a dependency tree, and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the dependencies in the tree.

How Snyk analyzes and builds the tree varies depending on the language and package manager for the Project and the location of your Project.

## Snyk CLI for Python Projects

When you use the CLI, how Snyk analyzes and builds the tree varies depending on the language and package manager for the Project.

{% hint style="info" %}
To scan your Projects, you must first install the relevant package manager and ensure that your Project contains the supported manifest files.
{% endhint %}

### Snyk CLI options for Python

For information about the `snyk test` options available for use with Python, see [Options for Python Projects](https://docs.snyk.io/snyk-cli/commands/test#options-for-python-projects) in the `test` command help.&#x20;

For the available `snyk monitor` options, see [Options for Python Projects](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-python-projects) in the `monitor` command help.

### Pip and Snyk CLI

{% hint style="info" %}
Run `pip install` before scanning with the CLI,`for` example:

```
pip install -r requirements.txt
```
{% endhint %}

Pip `requirements.txt` files specify only top-level dependencies, not nested or transitive ones. Therefore, the full Pip Project must be installed to ensure the CLI can build a complete dependency tree.

### Poetry and Snyk CLI

To build the dependency tree for a Poetry application, Snyk uses `pyproject.toml` and `poetry.lock` files. Both files must be present for Snyk to scan Poetry dependencies and identify issues.

If no `poetry.lock` file is present; you should run `poetry lock` to generate one before scanning.

{% hint style="info" %}
[PEP 621](https://peps.python.org/pep-0621/) is a standard for defining direct dependencies in `pyproject.toml` files, which is different from how Poetry does this.&#x20;

Snyk does not currently support PEP 621.
{% endhint %}

### Pipenv and Snyk CLI

To build the dependency tree for a Pipenv application, Snyk uses `Pipfile` and `Pipfile.lock` files. Both files must be present for Snyk to scan Poetry dependencies and identify issues.

{% hint style="info" %}
Run `pip install` before scanning with the CLI
{% endhint %}

Run `pipenv install` to ensure the CLI can build an up-to-date, accurate dependency tree using `pipenv graph`.

### setup.py and Snyk CLI

To build the dependency tree, Snyk analyzes the `setup.py` file, and detects packages listed in the `install_requires` key.

This file will not be discovered automatically by the CLI. It must be specified manually using the  `--file` option, for example:

```
snyk test --file=setup.py
```

You can also convert `setup.py` to `requirements.txt` by installing the packages into a virtual environment and then running `pip freeze`.

## Git services for Python Projects

To scan your Projects, you must ensure your repository contains the supported manifest files.

### Pip and Git repositories

{% hint style="warning" %}
URLs in `requirements.txt` files are not supported, as this poses a security risk.\
They are removed before resolving the dependencies in the files.
{% endhint %}

To scan Pip Projects, Snyk analyzes your `requirements.txt` files using native `pip` tooling in an isolated Linux environment.

Snyk imports any file that follows the `**/*req*.txt` pattern. This can help if you have renamed the `requirements.txt` files, for example, to `requirements-dev.txt`.

Snyk also looks for files using the `**/requirements/*.txt` pattern. This can help if you have placed your files in a `requirements` folder, for example, `requirements/requirements.txt`.

If you are using a package manager that creates different manifest file formats from `requirements.txt`, then you may be able to convert or export the manifest file to the `requirements.txt` format.

An example follows of how `dephell` is used to convert from  Conda `environments.yml` to a `requirements.txt`.

```python
dephell deps convert --from=conda --to=requirements.txt
```

### Poetry and Git repositories

To scan Poetry Projects, Snyk inspects your `pyproject.toml` and `poetry.lock` files.

You can choose whether Snyk should include [dev dependencies](https://python-poetry.org/docs/managing-dependencies/) when scanning your Poetry Projects.

Snyk regards non-dev dependencies to be those declared in `tool.poetry.dependencies`, the implicit `main` group. All others are classed as dev dependencies.

Poetry dev dependencies are not included in scans by default. To change this, modify your settings as follows:

1. Log in to your Snyk account and navigate to the relevant Group and Organization.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for **Python**.
4. Enable or disable the **Scan Poetry dev dependencies** option under the **Poetry dev dependencies** settings.&#x20;

<figure><img src="../../../.gitbook/assets/Screenshot 2023-04-03 at 16.00.22.png" alt="Poetry dev dependency settings"><figcaption><p>Poetry dev dependency settings</p></figcaption></figure>

### Pipenv and Git repositories

{% hint style="warning" %}
Pipenv is momentarily available as an open beta feature, meaning that some of its functionalities might be subject to change.
{% endhint %}

To scan Pipenv Projects, Snyk analyzes your `Pipfile` and `Pipfile.lock` files using native `pipenv` tooling in an isolated Linux environment.

{% hint style="info" %}
When you scan in the isolated environment, private packages and those with non-Linux OS requirements may be unresolvable and omitted from the dependency tree.

If a `Pipfile.lock` is present, any unresolved packages it contains are added to the top level of the dependency tree using versions from the lock file.&#x20;
{% endhint %}

You can choose whether Snyk should include dependencies specified in `[dev-packages]` when scanning your Pipenv Projects.

Pipenv dev dependencies are not included in scans by default. To change this, modify your settings as follows:

1. Log in to your Snyk account and navigate to the relevant Group and Organization.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for **Python**.
4. Enable or disable the **Scan Pipenv dev dependencies** option under the **Pipenv** settings.&#x20;

<figure><img src="../../../.gitbook/assets/Screenshot 2023-08-25 at 15.40.53.png" alt="Pipenv dev dependency settings"><figcaption><p>Pipenv dev dependency settings</p></figcaption></figure>

## Python version support

Some Python Projects may have dependencies that are valid with certain versions of Python only.&#x20;

Therefore, the version of Python used when scanning can affect the dependency tree Snyk generates.

You can specify the version of Python that Snyk uses to scan dependencies in both the CLI and Git integration.

### Setting the Python version in the CLI

Add the following option to `snyk test` or `snyk monitor` with the value of the Python binary:

```
--command=python3
```

For details, see the [Test command help](https://docs.snyk.io/snyk-cli/commands/test#options-for-python-projects) and the [Monitor command help](https://docs.snyk.io/snyk-cli/commands/monitor#options-for-python-projects).

### Setting the Python version in Git Projects

#### Pip and Python versions

{% hint style="info" %}
When you are scanning Pip Projects imported from Git, Snyk uses Python 2 or 3.&#x20;

Currently supported versions are `2.7.16` and `3.7.4`.
{% endhint %}

By default, Snyk tests Pip Projects using Python 3.

To define which Python Major version Snyk uses to test your Git-imported Pip Projects, use either Organization settings or a [`.snyk` policy file](../../../manage-issues/policies/the-.snyk-file.md).

To define the Python version for all Projects in an Organization:

1. Log in to your Snyk account and navigate to the relevant Group and Organization.
2. Select **Settings**, then **Languages**.
3. Select **Edit settings** for **Python**.
4. From the **Python version** dropdown, select the **Python 2** or **Python 3** version to use when testing Projects for this Organization.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-08-25 at 16.07.26.png" alt="Pip Python version settings"><figcaption><p>Pip Python version settings</p></figcaption></figure>

Snyk recommends you create different Organizations to work with different Python versions.

If you prefer one Organization but require Projects to use different Python versions, you may add a `.snyk` file to a Project repository and specify the desired version.

The `.snyk` file must be in the same directory as the Project manifest file.

Snyk detects the `Major version` specified and uses this to decide whether to test with Python 2 or Python 3, for example:

```python
language-settings:
  python: '3.7.2'
```

In this example, Snyk runs the scan using its currently supported version of Python 3. Snyk does not use the exact minor and patch version specified.

#### Pipenv and Python versions

{% hint style="warning" %}
Pipenv is momentarily available as an open   feature, meaning that some of its functionalities might be subject to change.
{% endhint %}

{% hint style="info" %}
Currently supported Python versions are `3.8`, `3.9`, `3.10`, `3.11`.
{% endhint %}

Snyk uses Python version information specified in each `Pipfile` to choose the major and minor versions to use in scanning, for example:

```python
[requires]
python_version = "3.6"
```

Specific Patch versions are ignored; Snyk uses a recent Patch version from each series.

Snyk defaults to Python `3.10` if the `Pipfile` contains:

* No Python version information
* Only a Major version
* An unsupported version

#### Poetry and Python versions

There is no need to inform Snyk about Python versions for Poetry Projects.

Poetry files contain sufficient information to build a full dependency tree without running native tooling.

## Scanning Python Projects in IDEs

If you are using any of the supported IDEs to write Python, there are some configurations you must add to scan Python manifest files properly.

If you are using a virtual environment, you must add the `PYTHON_PATH` to the **Additional Options** text input in the Snyk integration settings, for example, `--command=.venv/bin/python`. Snyk tries to look for a `*req*.txt` file in the root of the Project as it is seen in the IDE.

However, if you have manifest files in other directories within the root of the Project, Snyk cannot identify them. For Snyk to find them, you must use the `--all-projects` option. Snyk then recursively searches each Project directory to find all the manifest files.

If those directories each require a different virtual environment to run, the Snyk scan will not be successful because it will use one virtual environment to search for installed dependencies. In this case, it is best to use the CLI or the Git integration to get vulnerability information on all the dependencies listed in each Project directory.
