# Authenticate to use the CLI

Once you have [installed the Snyk CLI](install-the-snyk-cli/) using your chosen tool or operating system (OS), you need to authenticate with your Snyk account.

## Supported authentication methods

Snyk supports the following authentication methods:

* OAuth 2.0
* Personal Access Token (PAT)
* Snyk API token

## Prerequisites

If your Snyk account is hosted on a `SNYK-US-02`, `SNYK-EU-01,` or `SNYK-AU-01`  endpoint rather than the default `SNYK-US-01` endpoint, you need to configure the CLI environment before you authenticate with your Snyk account.

To do this, run the environment config command:

{% code title="Environment config command" %}
```bash
snyk config environment <ENVIRONMENT_NAME>
```
{% endcode %}

This example specifies the SNYK-EU-01 environment:

{% code title="Example command with specified environment" %}
```shellscript
snyk config environment SNYK-EU-01
```
{% endcode %}

Once the environment has been configured, authenticate using your preferred method.

## Authenticate locally

{% hint style="info" %}
This applies to: Homebrew, npm, Yarn, Scoop (Windows), and Direct binary downloads (standalone executables).&#x20;
{% endhint %}

<details>

<summary>OAuth 2.0 (browser based)</summary>

*   Run the following command:

    ```bash
    snyk auth
    ```
* A browser window will open on a Snyk address. Click **Authenticate**.
* Return to your terminal. You should see the message `Your account has been authenticated`.

For local service account testing, use the command:

```bash
 snyk auth --auth-type=oauth --client-id=<ID> --client-secret=<SECRET>
```

</details>

<details>

<summary>Personal Access Token (PAT)</summary>

To authenticate in the terminal and save the profile locally, run the command:

```bash
snyk auth <YOUR_PAT>
```

</details>

<details>

<summary>Snyk API token</summary>

Run the command:

```bash
snyk auth <YOUR_API_TOKEN>
```

</details>

<details>

<summary>Configuration command</summary>

Add your PAT or API token to the Snyk CLI configuration using the `snyk config` command:

```bash
snyk config set api=<YOUR_PAT_OR_API_TOKEN>
```

</details>

<details>

<summary>Environment variables</summary>

You can export local tokens to your local shell profile using the command:&#x20;

```bash
export SNYK_TOKEN=<YOUR_PAT_OR_API_TOKEN>
```

Or run them for each command using:

```bash
SNYK_TOKEN=<TOKEN> snyk <COMMAND>
```

</details>

## Authenticate in Docker

Pass your chosen token into the container using Docker's environment flag during execution:

```shellscript
docker run -e SNYK_TOKEN=<YOUR_PAT_OR_API_TOKEN> ...
```

```shellscript
docker run -e SNYK_OAUTH_TOKEN=<YOUR_OAUTH_TOKEN> ...
```

## Authenticate in a CI/CD or GitHub action

{% hint style="info" %}
This applies to: general CI/CD Pipelines (Jenkins, GitLab, CircleCI, and so on) and GitHub Actions.
{% endhint %}

Map your securely stored PAT or API token to the `SNYK_TOKEN` environment variable, or map your OAuth token to `SNYK_OAUTH_TOKEN`, directly within your pipeline configuration or GitHub Actions `env` block.
