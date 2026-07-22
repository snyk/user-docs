---
description: How to upgrade the Snyk CLI to stay secure and current
nav_context: agnostic
---

# Upgrade the Snyk CLI

Upgrade the Snyk CLI regularly to maintain a secure environment. Snyk releases frequent updates to provide the latest vulnerability definitions, introduce new features, and ensure compatibility with package managers. Outdated versions can limit scanning capabilities, trigger false positives, or cause inaccurate results.

Select the upgrade command that matches your original installation method.

## Package managers

<details>

<summary>Homebrew</summary>

Upgrade the Snyk CLI to the latest version using the standard Homebrew upgrade command. Snyk updates the Homebrew tap daily:

{% code overflow="wrap" expandable="true" %}
```bash
brew upgrade snyk
```
{% endcode %}

</details>

<details>

<summary>npm</summary>

Check your current version using the command `npm -v`.

To upgrade to the latest version of the Snyk CLI with npm, use one of these commands:

{% code overflow="wrap" expandable="true" %}
```bash
npm install -g snyk@latest
```
{% endcode %}

{% code overflow="wrap" expandable="true" %}
```bash
npm install -g snyk
```
{% endcode %}

</details>

<details>

<summary>Yarn</summary>

To upgrade the Snyk CLI globally using Yarn, run the upgrade command or add the package again to fetch the latest version:

{% code overflow="wrap" expandable="true" %}
```bash
yarn global upgrade snyk
```
{% endcode %}

{% code overflow="wrap" expandable="true" %}
```bash
yarn global add snyk
```
{% endcode %}

</details>

<details>

<summary>Scoop (Windows)</summary>

Check your current version using the command `scoop status`.

To upgrade to the latest version of Scoop, use the command:

{% code overflow="wrap" expandable="true" %}
```bash
scoop update snyk
```
{% endcode %}

</details>

## Docker

Upgrade the Snyk CLI in a Docker environment by pulling the latest image from Docker Hub. This action overwrites the local cached image:

{% code overflow="wrap" expandable="true" %}
```bash
## Pull the latest universal image
docker pull snyk/snyk
```
{% endcode %}

{% code overflow="wrap" expandable="true" %}
```bash
## Pull the latest language-specific image (for example, Python)
docker pull snyk/snyk:python
```
{% endcode %}

## GitHub actions

The Snyk GitHub Action type in your workflow determines how the Snyk CLI updates.

Using the `@master` branch tag (for example, `uses: snyk/actions/node@master`) ensures your pipeline pulls the latest version of the GitHub Action. This tag does not dictate the CLI version.

Snyk updates the CLI version in one of two ways:

### Language-specific actions (Docker based)

Language-specific actions (such as Node and Python) run using Snyk Docker images. These images use the latest version of the Snyk CLI by default.

{% code overflow="wrap" expandable="true" %}
```yaml
- name: Run Snyk to check for vulnerabilities
  uses: snyk/actions/node@master
```
{% endcode %}

### Snyk setup action (binary installation)

The Snyk Setup action installs the CLI directly from a binary. By default, this action downloads and uses the latest version of the Snyk CLI.

{% code overflow="wrap" expandable="true" %}
```yaml
- name: Set up Snyk CLI
  uses: snyk/actions/setup@master
```
{% endcode %}

Unlike Docker-based actions, the Snyk Setup action allows you to pin a specific CLI version to prevent automatic updates. Use the `version` input to specify the version:

{% code overflow="wrap" expandable="true" %}
```yaml
- name: Set up Snyk CLI
  uses: snyk/actions/setup@master
  with:
    snyk-version: latest # Alternatively, pin a specific version (for example, 'v1.1305.0')
```
{% endcode %}

## Direct binary download

Update the Snyk CLI manually when using standalone executables. To update, download the latest binary and replace the existing executable. For example, on Linux or MacOS, run the curl download command again.

For more information including the download commands, visit [Install the Snyk CLI](../install-the-snyk-cli/#direct-binary-download).
