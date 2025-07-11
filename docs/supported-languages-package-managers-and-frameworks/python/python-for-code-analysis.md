# Python for code analysis

Snyk Code relies on Python Projects' adhering to a **s**tandard directory layout for accurate analysis. Specifically, Snyk Code expects Projects to be compatible with [`setuptools` automatic discovery](https://setuptools.pypa.io/en/latest/userguide/package_discovery.html#auto-discovery), which identifies packages and modules automatically based on the directory structure. This includes support for `init.py` files to ensure that symbols defined in package initialization files are imported correctly, leading to a more accurate and deeper analysis. Both `src-layout` and `flat-layout` are supported.

For further details, see the [`setuptools` automatic package discovery documentation](https://setuptools.pypa.io/en/latest/userguide/package_discovery.html#auto-discovery). Proper adherence to these conventions allows the scanner to trace code effectively and provide accurate results.

Refer to the [Python detail](./) for supported frameworks, libraries, and features.

