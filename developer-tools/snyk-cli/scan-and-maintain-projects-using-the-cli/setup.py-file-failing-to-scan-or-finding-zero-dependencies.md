# Setup.py file failing to scan or finding zero dependencies

When you run the command `snyk test --file=setup.py`, typically there are some Python `setup.py` projects that fail or output 0 dependencies. This is because Snyk reads through only the `install_requires` key that has all the packages listed.

Any dependencies in other `requires` keys are not scanned.

Future support is subject to change.
