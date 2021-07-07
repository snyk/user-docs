# Why is my setup.py file failing to scan or finding 0 dependencies?

##  Why is my setup.py file failing to scan or finding 0 dependencies?

Commonly, when executing the 

```text
snyk test --file=setup.py
```

command, there are some python 'setup.py' projects that fail or output 0 dependencies. This is because Snyk at the current moment only reads through the

```text
install_requires 
```

key that has all the packages listed. 

Any dependencies in other 'requires' key will not be scanned.

Further support is subject to change in the future.

