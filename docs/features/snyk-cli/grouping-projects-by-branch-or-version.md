# Grouping projects by branch or version

You can use `--target-reference` to separate projects into specific groupings like branch or version.

`--target-reference` takes any text so you can combine it with another command to automatically set it to a value. For example, below is a script that when executed will set `--target-reference` to the current Git branch.

```text
#!/usr/bin/env bash
snyk monitor --target-reference="$(git branch --show-current)"
```

Similarly if you'd like to use the latest Git tag you can use:

```text
#!/usr/bin/env bash
snyk monitor --target-reference="$(git describe --tags --abbrev=0)"
```

Adjust the command to your liking using whichever developer tools your project uses.

