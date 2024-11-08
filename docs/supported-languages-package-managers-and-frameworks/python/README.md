# Python

## Applicability

Snyk supports [Python for code analysis](python-for-code-analysis.md) and [Python for open source](python-for-open-source.md).&#x20;

For specific information about the use of versions and package managers, See [Snyk CLI for Python](snyk-cli-for-python.md) and [Git repositories and Python](git-repositories-and-python.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

Available functions:

* SCM import. For Python used with Snyk Open Source, the SCM import is available for Pip, pipenv, and Poetry.
* Test or monitor your app through CLI and IDE.
* Test your app's SBOM using pkg:pypi
* Test your app's packages using pkg:pypi

## Package managers

This language supports Pip, Poetry, pipenv, setup.py as package managers, suitable with the following Python versions: `Python 2 -> 2.7.16`, and `Python 3 -> 3.7.4`.

As a package registry, [pypi.org](https://pypi.org/) is supported.

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Python:&#x20;

* AioHTTP - Comprehensive a
* iopg - Comprehensive&#x20;
* argparse - Comprehensive&#x20;
* anthropic - Comprehensive&#x20;
* bottle - Comprehensive&#x20;
* CherryPy - Comprehensive&#x20;
* Django - Comprehensive&#x20;
* defusedxml - Comprehensive&#x20;
* fastapi - Partial&#x20;
* flask - Comprehensive&#x20;
* flask\_pymongo - Comprehensive&#x20;
* google.cloud.bigquery - Comprehensive&#x20;
* google\_generativeai - Comprehensive&#x20;
* huggingface\_hub - Comprehensive&#x20;
* httpx - Comprehensive&#x20;
* ldap3 - Comprehensive&#x20;
* libxml - Comprehensive&#x20;
* lxml - Comprehensive&#x20;
* mistralai - Comprehensive&#x20;
* mongoengine - Comprehensive&#x20;
* openai - Comprehensive&#x20;
* pandas - Partial&#x20;
* paramiko - Comprehensive&#x20;
* peewee - Comprehensive&#x20;
* pickle - Comprehensive&#x20;
* pilyaml - Comprehensive&#x20;
* pyca/cryptography - Comprehensive&#x20;
* pymongo - Comprehensive&#x20;
* pymssql - Comprehensive&#x20;
* pyramid - Comprehensive&#x20;
* psycopg - Comprehensive&#x20;
* python-ldap - Comprehensive&#x20;
* Python Standard Library - Comprehensive&#x20;
* requests - Comprehensive&#x20;
* sqlite3 (or pysqlite2) - Comprehensive&#x20;
* sqlalchemy - Comprehensive&#x20;
* turboGears - Comprehensive&#x20;
* urllib - Comprehensive&#x20;
* werkzeug - Comprehensive

## Features

The following features are supported in Snyk for Python:

* Fix PRs&#x20;
* License scanning&#x20;
* Reports
* Custom rules
* Interfile analysis

## Python version support

Some Python Projects may contain dependencies that require specific versions of Python. Therefore, the version of Python used when scanning can affect the dependency tree that Snyk generates.

You can specify the version of Python that Snyk uses to scan dependencies in both the CLI and Git integration.

See [Snyk CLI for Python](snyk-cli-for-python.md) for information about the Python version and installation and use information for Pip, Poetry, Pipevn, and setup.py.

See [Git repositories and Python](git-repositories-and-python.md) for information about the Python version and installation and use for Python and pip and use of Poetry and pipenv.

### Pipenv and Python versions supported

Supported Python versions are `3.8`, `3.9`, `3.10`, `3.11`, `3.12`.

Snyk uses Python version information specified in each `Pipfile` to choose the major and minor versions to use in scanning, for example:

```python
[requires]
python_version = "3.6"
```

Specific patch versions are ignored; Snyk uses a recent patch version from each series.

Snyk defaults to Python `3.10` if the `Pipfile` contains:

* No Python version information
* Only a major version
* An unsupported version

### Poetry and Python versions supported

There is no need to tell Snyk the Python version for Poetry Projects.

Poetry files contain sufficient information to build a full dependency tree without running native tooling.

## IDE and CI/CD with Snyk for Python

If you are using any of the supported IDEs to write Python, there are some configurations you must add to scan Python manifest files properly.

If you are using a virtual environment, you must add the `PYTHON_PATH` to the **Additional Options** text input in the Snyk integration settings, for example, `--command=.venv/bin/python`. Snyk tries to look for a `*req*.txt` file in the root of the Project as it is seen in the IDE.

However, if you have manifest files in other directories within the root of the Project, Snyk cannot identify them. For Snyk to find them, you must use the `--all-projects` option. Snyk then recursively searches each Project directory to find all the manifest files.

If those directories each require a different virtual environment to run, the Snyk scan will not be successful because it will use one virtual environment to search for installed dependencies. In this case, it is best to use the CLI or the Git integration rather than an IDE to get vulnerability information on all the dependencies listed in each Project directory.

## Troubleshooting Snyk for Python

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).
