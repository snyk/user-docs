# Python for open source

## Python for open source support

**Package manager**: Pip, Poetry, pipenv, setup.py

**Package manager versions**: Suitable with `Python 2 -> 2.7.16`, and `Python 3 -> 3.7.4`.

**Package registry**: [pypi.org](https://pypi.org)

**Import your app through SCM**: Available for Pip, pipenv, and Poetry

**Test or monitor your app through CLI and IDE**: Available

**Test your app's SBOM**: Available, `pkg:pypi`

**Test your app's packages**: Available, `pkg:pypi`

**Features**:&#x20;

* Fix PRs&#x20;
* License scanning
* Reports

## Open source and licensing

{% hint style="info" %}
Some features may not be available, depending on your plan. For more information, see  [Plans and pricing.](https://snyk.io/plans/)
{% endhint %}

The following summarizes the package managers supported by Snyk for Python for open source.

| Package managers Features                                                                                                        | CLI support                                                                                                                                 | Git support | License scanning | Fix PRs |
| -------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | ---------------- | ------- |
| [Pip](https://pypi.org/project/pip/)                                                                                             | ✔︎                                                                                                                                          | ✔︎          | ✔︎               | ✔︎      |
| [Poetry](https://python-poetry.org)                                                                                              | ✔︎                                                                                                                                          | ✔︎          | ✔︎               |         |
| [Pipenv](https://pipenv.pypa.io/en/latest/)                                                                                      | ✔︎                                                                                                                                          | ✔︎          | ✔︎               |         |
| [setup.py](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/snyk-cli-for-python#setup.py-and-cli) | ✔︎[ see instructions](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/snyk-cli-for-python#setup.py-and-cli) |             | ✔︎               |         |

To scan your Projects, you must first install the relevant package manager and ensure that your Project contains the supported manifest files.
