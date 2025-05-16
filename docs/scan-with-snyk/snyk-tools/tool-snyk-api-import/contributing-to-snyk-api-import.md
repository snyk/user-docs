# Contributing to snyk-api-import

## Contributor Agreement

A pull-request will only be considered for merging into the upstream codebase after you have signed Snyk's contributor agreement, assigning to Snyk the rights to the contributed code and granting you a license to use it in return. If you submit a pull request, you will be prompted to review and sign the agreement with one click (Snyk uses [CLA assistant](https://cla-assistant.io/)).

## Pull requests

### Commit messages

Commit messages must follow the [Angular-style](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format) commit format (but excluding the scope):

```
fix: minified scripts being removed

Also includes tests
```

This allows for the automatic change log to generate correctly.

Commit types must be one of the following:

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **test**: Adding missing tests
* **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons and other similar changes)
* **perf**: A code change that improves performance

To release a major update you must add `BREAKING CHANGE:` to the start of the body and the detail of the breaking change.

### Code standards

Ensure that your code adheres to the included `.eslintrc` config by running `npm run lint`.

### Sending pull requests

* New command line options are generally discouraged unless there's a _really_ good reason.
* Add tests for newly added code (and try to mirror directory and file structure if possible).
* Spell check.
* PRs will not be code reviewed unless all tests are passing.

{% hint style="info" %}
When fixing a bug, commit a failing test first so that CircleCI (or the approver) can show the code failing. Once that commit is in place, then commit the bug fix so that Snyk can test before and after.
{% endhint %}

Remember that you're developing for multiple platforms and versions of Node.js, so if the tests pass on your Mac or Linux or Windows machine, the tests may not pass elsewhere.
