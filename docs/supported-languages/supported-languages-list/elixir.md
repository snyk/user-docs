# Elixir

## Applicability and integration

{% hint style="info" %}
Elixir is supported only for Snyk Open Source.
{% endhint %}

Available integrations: CLI and IDE: test or monitor your app

Available functions:

* Test your app's SBOM using `pkg:hex`
* Test your app's packages using `pkg:hex`

## Technical specifications

* Supported package manager:  [Mix](https://hexdocs.pm/mix/Mix.html)/[Hex](https://hex.pm/)
* Supported package registry: [hex.pm](https://hex.pm/)

### Supported features

For Elixir, Snyk supports the **Reports** feature.&#x20;

{% hint style="info" %}
The **Snyk Fix PR** feature is not available for Elixir. This means that you will not be notified if the PR checks fail when the following conditions are met:&#x20;

* The **PR checks** feature is enabled and configured to **Only fail when the issues found have a fix available.**
* "**Fixed in" available** is set to **Yes.**
{% endhint %}

### Snyk CLI for Elixir

{% hint style="info" %}
To scan your dependencies, you must first install Elixir and Mix. For details, [see the Elixir installation instructions](https://elixir-lang.org/install.html).
{% endhint %}

Snyk offers security scanning to test your Elixir Projects for vulnerabilities using the [CLI](../../developer-tools/snyk-cli/).

Mix is a build tool that compiles, tests, and creates Elixir projects. Mix manages dependencies by integrating with the Hex package manager.

Snyk builds a dependency tree for your Project by analyzing your `mix.exs` and `mix.lock` files. The `mix.lock` file must be present and in sync with the `mix.exs` file. After Snyk builds the tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in the packages anywhere in the dependency tree.

#### **Project naming**

Projects in the Snyk UI are named according to the `app` keyword from the `project/0` function exported by `Mix.Project` in the main `mix.exs` file.

To override the name, use the `--project-name` CLI option.

#### **Mix umbrella projects**

If you test a Mix umbrella project, Snyk detects that it is an umbrella project and includes all the child apps automatically.

Along with the main `mix.exs`, each app `mix.exs` appears as a separate Project in the Snyk UI, named according to the path to the app.

Snyk fully supports all `:hex` packages listed in the Mix project, including all their transitive dependencies and any vulnerabilities.

Hex support includes both Elixir and Erlang packages.

Snyk also has limited support for `:path`, `:git` and `:github` dependencies, but not their transitive dependencies or vulnerabilities.

* `:path` dependencies appear in the dependency tree by name
* `:git` and `:github` dependencies appear in the dependency tree by repository URL and version (either `:branch`, `:tag` or `:ref`, as defined in the `mix.exs` file)

{% hint style="info" %}
When using `asdf`, ensure you set a version by running the `asdf global elixir <version of your choice>`.
{% endhint %}
