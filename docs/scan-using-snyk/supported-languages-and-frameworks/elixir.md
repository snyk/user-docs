# Elixir

{% hint style="info" %}
Snyk for Elixir is supported only for Snyk Open Source.
{% endhint %}

## Supported frameworks and package managers

{% hint style="warning" %}
You might encounter false positives or false negatives for partially supported frameworks and package managers.
{% endhint %}

Snyk offers security scanning to test your Elixir Projects for vulnerabilities using the [CLI](../../snyk-cli/).

{% hint style="info" %}
Features may not be available, depending on your plan. See [pricing plans](https://snyk.io/plans/) for more details.
{% endhint %}

| Package managers / Features                                  | CLI support | Git support | License scanning | Fix PRs |
| ------------------------------------------------------------ | ----------- | ----------- | ---------------- | ------- |
| [Mix](https://hexdocs.pm/mix/Mix.html)/[Hex](https://hex.pm) | ✔︎          |             |                  |         |

Snyk builds a dependency tree for your Project by analyzing your manifest and lock files.

After Snyk builds the tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in the packages anywhere in the dependency tree.

## Getting started with Snyk for Elixir across environments

### Snyk CLI&#x20;

#### Mix/Hex

{% hint style="info" %}
To scan your dependencies, first install Elixir and Mix. For details, [see the Elixir installation instructions](https://elixir-lang.org/install.html).
{% endhint %}

Mix is a build tool that manages dependencies, compiles, tests, and creates Elixir projects.

Mix manages dependencies by integrating with the Hex package manager.

To build the dependency tree, Snyk analyzes your `mix.exs` and `mix.lock` files. The `mix.lock` file must be present and in sync with the `mix.exs` file.

#### **Project naming**

Projects in the Snyk UI are named according to the `app` keyword from the `project/0` function exported by `Mix.Project` in the main `mix.exs` file.

To override the name, use the `--project-name` CLI option.

#### **Umbrella projects**

If you test a Mix Umbrella project, Snyk detects that this is an umbrella project and includes all the child apps automatically.

Along with the main `mix.exs`, each app `mix.exs` appears as a separate Project in the Snyk UI, named according to the path to the app.

Snyk fully supports all `:hex` packages listed in the Mix project, including all their transitive dependencies and any vulnerabilities.

Hex support includes both Elixir and Erlang packages.

Snyk also has limited support for `:path`, `:git` and `:github` dependencies, but not their transitive dependencies or vulnerabilities.

* `:path` dependencies appear in the dependency tree by name
* `:git` and `:github` dependencies appear in the dependency tree by repository URL and version (either `:branch`, `:tag` or `:ref`, as defined in the `mix.exs` file)

### Snyk integrations&#x20;

:link: For integrated development environments, see [Use Snyk in your IDE](../../integrate-with-snyk/ide-tools/).

:link: If you prefer continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software (see [Snyk CI/CD](../../integrate-with-snyk/snyk-ci-cd-integrations/) and [Snyk API](../../snyk-api/)).

## Troubleshooting

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).&#x20;
