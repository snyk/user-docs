# Python

## Applicability

Snyk supports [Python for code analysis](python-for-code-analysis.md) and [Python for open source](python-for-open-source.md).

For specific information about the use of versions and package managers, See [Snyk CLI for Python](snyk-cli-for-python.md) and [Git repositories and Python](scm-integrations-and-python.md).

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.

Available functions:

* SCM import, available for Snyk Open Source and Snyk Code. For Python used with Snyk Open Source, the SCM import is available for Pip, pipenv, and Poetry.
* Test or monitor your app through CLI and IDE, available for Snyk Open Source and Snyk Code.
* Test your app's SBOM using pkg:pypi
* Test your app's packages using pkg:pypi

## Package managers and supported file extensions

Snyk for Python supports Pip, Poetry, pipenv, and setup.py as package managers. For the list of supported Python versions, check the [Git repositories and Python](scm-integrations-and-python.md) page.

As a package registry, [pypi.org](https://pypi.org/) is supported.

Snyk for Python supports the following file formats:

* Snyk Open Source:
  * For poetry: `pyproject.toml`, `poetry.lock`
  * For pip: `requirements.txt`
  * For pipenv: `pipfile`, `pipfile.lock`
  * For setup.py: `setup.py`
* Snyk Code: `.py`

## Frameworks and libraries

The following frameworks and libraries are supported in Snyk for Python:

* AioHTTP - Comprehensive
* iopg - Comprehensive
* aiofiles - Comprehensive
* argparse - Comprehensive
* anthropic - Comprehensive
* bottle - Comprehensive
* CherryPy - Comprehensive
* Django - Comprehensive
* defusedxml - Comprehensive
* fastapi - Partial
* fastMCP - Comprehensive
* flask - Comprehensive
* flask\_pymongo - Comprehensive
* google.cloud.bigquery - Comprehensive
* google\_generativeai - Comprehensive
* grpcio - Comprehensive
* huggingface\_hub - Comprehensive
* httpx - Comprehensive
* ldap3 - Comprehensive
* libxml - Comprehensive
* lxml - Comprehensive
* mistralai - Comprehensive
* modelcontextprotocol/python-sdk - Comprehensive
* mongoengine - Comprehensive
* openai - Comprehensive
* pandas - Partial
* paramiko - Comprehensive
* peewee - Comprehensive
* pickle - Comprehensive
* pilyaml - Comprehensive
* pyca/cryptography - Comprehensive
* pymongo - Comprehensive
* pymssql - Comprehensive
* pyramid - Comprehensive
* psycopg - Comprehensive
* python-ldap - Comprehensive
* Python Standard Library - Comprehensive
* requests - Comprehensive
* sqlite3 (or pysqlite2) - Comprehensive
* sqlalchemy - Comprehensive
* turboGears - Comprehensive
* urllib - Comprehensive
* werkzeug - Comprehensive

## Features

The following features are supported in Snyk for Python:

| Snyk Open Source                                                   | Snyk Code                                                                 |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------- |
| <ul><li>Fix PRs</li><li>License scanning</li><li>Reports</li></ul> | <ul><li>Reports</li><li>Custom rules</li><li>Interfile analysis</li></ul> |

{% hint style="info" %}
For details on support for specific package managers in Snyk Open Source, see [Python for Open Source](python-for-open-source.md).&#x20;
{% endhint %}

## Python version support

Some Python Projects may contain dependencies that require specific versions of Python. Therefore, the version of Python used when scanning can affect the dependency tree that Snyk generates.

You can specify the version of Python that Snyk uses to scan dependencies in both the CLI and Git integration.

See [Snyk CLI for Python](snyk-cli-for-python.md) for information about the Python version and installation and use information for Pip, Poetry, Pipevn, and setup.py.

See [Git repositories and Python](scm-integrations-and-python.md) for information about the Python version and installation and use for Python and pip and use of Poetry and pipenv.

### Pipenv and Python versions supported

The supported versions of Python are `2.7`, `3.7`, `3.8`, `3.9`, `3.10`, `3.11`, `3.12`, `3.13`.

Snyk uses Python version information specified in each `Pipfile` to choose the major and minor versions to use in scanning, for example:

```python
[requires]
python_version = "3.8"
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
