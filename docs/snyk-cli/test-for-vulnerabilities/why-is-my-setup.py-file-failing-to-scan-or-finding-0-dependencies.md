# Why is my setup.py file failing to scan or finding 0 dependencies?

Commonly, when executing the command, there are some python 'setup.py' projects that fail or output 0 dependencies. This is because Snyk at the current moment only reads through the key that has all the packages listed. 

```text
snyk test --file=setup.py
```

```text
install_requires 
```

Any dependencies in other 'requires' key will not be scanned.

Further support is subject to change in the future.

