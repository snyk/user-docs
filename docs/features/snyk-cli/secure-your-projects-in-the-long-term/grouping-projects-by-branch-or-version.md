# Separating projects by branch or version

{% hint style="warning" %}
This feature is currently in Open Beta. There will be areas where it is not fully supported. Currently [Snyk Open Source](../../../products/snyk-open-source/) is supported.
{% endhint %}

Sometimes your project might have multiple states which you want to monitor separately. These could be branches, releases, deployments and so on. You can use `--target-reference` to separate projects into these specific groupings.

`--target-reference` takes any text so you can combine it with another command to automatically set it to a value. For example, below is a script that when executed will set `--target-reference` to the current Git branch.

```
snyk monitor --target-reference="$(git branch --show-current)"
```

Similarly if you'd like to use the latest Git tag you can use:

```
snyk monitor --target-reference="$(git describe --tags --abbrev=0)"
```

Adjust the command to your liking using whichever developer tools your project uses.

`--target-reference` will be used for sub-groupings in the Projects page.

![A project grouping with sub-groups.](<../../../.gitbook/assets/Screenshot 2021-10-21 at 13-08-58 Projects.png>)
