# Python

## Applicability and integration

{% hint style="info" %}
Python is supported for Snyk Code and Snyk Open Source.
{% endhint %}

Available integrations:

* SCM import
* CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM using `pkg:pypi`
* Test your app's packages using `pkg:pypi`

## Python for Snyk Code

For Python with Snyk Code, Python versions up to `3.11` are supported. Language features introduced in newer versions are not supported.

For Python with Snyk Code, the following file format is supported: `.py`

Available features:

* Reports
* Custom rules
* Interfile analysis

### Supported frameworks and libraries

For Python, the following frameworks and libraries are supported:

* AioHTTP
* iopg
* aiofiles
* argparse
* anthropic
* bottle
* CherryPy
* Django
* defusedxml
* fastapi
* fastMCP
* flask
* flask\_pymongo
* google.cloud.bigquery
* google\_generativeai
* grpcio
* huggingface\_hub
* httpx
* ldap3
* libxml
* lxml
* mistralai
* modelcontextprotocol/python-sdk
* mongoengine
* openai
* pandas
* paramiko
* peewee
* pickle
* pilyaml
* pyca/cryptography
* pymongo
* pymssql
* pyramid
* psycopg
* python-ldap
* Python Standard Library
* requests
* sqlite3 (or pysqlite2)
* sqlalchemy
* turboGears
* urllib
* werkzeug

### Directory layout

Snyk Code relies on Python projects to follow a standard directory layout for accurate analysis. Specifically, Snyk Code expects Projects to be compatible with [`setuptools` automatic discovery](https://setuptools.pypa.io/en/latest/userguide/package_discovery.html#auto-discovery), which identifies packages and modules automatically based on the directory structure. This includes support for `init.py` files to ensure that symbols defined in package initialization files are imported correctly, leading to a more accurate and deeper analysis.

Both `src-layout` and `flat-layout` are supported. Proper adherence to these conventions allows the scanner to trace code effectively and provide accurate results.

## Python for Snyk Open Source

{% hint style="info" %}
It is possible that some Python Projects contain dependencies that require specific versions of Python. Therefore, the version of Python used when scanning can affect the dependency tree that Snyk generates.
{% endhint %}

For Python with Snyk Open Source, the following Python versions are supported: `2.7`, `3.7`, `3.8`, `3.9`, `3.10`, `3.11`, `3.12`, `3.13`, `3.14`

### Supported package managers and registries

* Supported package managers: [Pip](https://pypi.org/project/pip/), [Poetry](https://python-poetry.org/), [pipenv](https://pipenv.pypa.io/en/latest/), and setup.py
* Supported package registry: [pypi.org](https://pypi.org/)

For pipenv Projects, Snyk uses Python version information specified in each `Pipfile` to choose the major and minor versions to use in scanning, for example:

```python
[requires]
python_version = "3.8"
```

Specific patch versions are ignored; Snyk uses a recent patch version from each series.

Snyk defaults to Python `3.10` if the `Pipfile` contains:

* No Python version information
* Only a major version
* An unsupported version

For Poetry projects, you do not need to specify the Python version. Poetry files contain sufficient information to build a full dependency tree without running native tooling.

### Supported file formats

* For poetry: `pyproject.toml` and `poetry.lock`
* For pip: `requirements.txt`
* For pipenv: `pipfile` and `pipfile.lock`
* For setup.py: `setup.py`

### Available features

* Fix PRs (supported only for pip)
* License scanning
* Reports

{% hint style="info" %}
Depending on your plan, some features may not be available. For more information, see [plans and pricing](https://snyk.io/plans/).
{% endhint %}

### IDE and CI/CD support

If you are using any of the supported IDEs to write Python, there are some configurations you must add to scan Python manifest files properly.

To scan your Projects, you must first install the relevant package manager and ensure that your Project contains the supported manifest files.

Python packages that are operating system-specific and not supported by Linux may not be compatible with Snyk SCM scans, leading to errors.

#### Virtual environment configuration

If you are using a virtual environment, you must add the following information:

* In your Snyk integration settings, locate the **Additional Options** field.
* Add the `PYTHON_PATH` to this field, for example `--command=.venv/bin/python`.

By default, the Snyk IDE integration looks for a `*req*.txt` file in the root of the Project as it is seen in the IDE.

#### Scan multiple directories

If you have manifest files in other directories within the root of the Project, Snyk cannot identify them unless directed to do so.

In your Snyk integration settings, locate the **Additional Options** field. Enable a recursive search and add the `--all-projects` option in the **Additional Options** field.

{% hint style="warning" %}
If each directory needs a different virtual environment, it is possible that the Snyk scan fails because it uses a single virtual environment for dependency detection. In such cases, Snyk recommends using the CLI or SCM integration instead of an IDE , in order to gather vulnerability details for all dependencies in each Project directory.
{% endhint %}
