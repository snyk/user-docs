# Python for open source

**Note:** The language pages are being restructured and updated on June 22 and June 23.

## Python support

**Package manager**: Pip, Poetry, pipenv, setup.py

**Package resgirty**: [pypi.org](https://pypi.org)

**Import your app through SCM**: Available for Pip, pipenv and Poetry

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:pypi`

**Test your app's packages**: Available, `pkg:pypi`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

**Package manager versions**: Suitable with `Python 2 -> 2.7.16`, and `Python 3 -> 3.7.4`.

## Open source and licensing

{% hint style="info" %}
Features may not be available, depending on your plan. See the [Plans and pricing ](https://snyk.io/plans/) page for more details.
{% endhint %}

{% hint style="info" %}
Before testing your Open Source Project for vulnerabilities, with limited exceptions, you must **build your Project**. For details, see [Open Source Projects that must be built before testing with the Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md).
{% endhint %}

| Package managers / Features                  | CLI support | Git support | License scanning | Fix PRs |
| -------------------------------------------- | ----------- | ----------- | ---------------- | ------- |
| [Pip](https://pypi.org/project/pip/)         | ✔︎          | ✔︎          | ✔︎               | ✔︎      |
| [Poetry](https://python-poetry.org)          | ✔︎          | ✔︎          | ✔︎               |         |
| [Pipenv](https://pipenv.pypa.io/en/latest/)  | ✔︎          | ✔︎          | ✔︎               |         |
| setup.py                                     | ✔︎          |             | ✔︎               |         |

Snyk builds a dependency tree and then uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in any of the dependencies in the tree.

How Snyk analyzes and builds the tree varies depending on the language and package manager for the Project and the location of your Project.

{% hint style="info" %}
To scan your Projects, you must first install the relevant package manager and ensure that your Project contains the supported manifest files.
{% endhint %}

## G
