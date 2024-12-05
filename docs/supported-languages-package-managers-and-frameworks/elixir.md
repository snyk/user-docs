# Elixir

## Applicability

Snyk for Elixir is supported **only for Snyk Open Source**.

Check the language availability to be imported as an application, tested, or monitored using the Snyk products.&#x20;

Available functions:

* Import your app through SCM: N/A
* Test or monitor your app through CLI and IDE, available for Snyk Open Source.&#x20;
* Test your app's SBOM using `pkg:hex`&#x20;
* Test your app's packages using `pkg:hex`

For information on using the Snyk CLI for code analysis, see [Snyk CLI for Snyk Code](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/).

## Package managers and supported file extensions

Snyk for Elixir supports [Mix](https://hexdocs.pm/mix/Mix.html)/[Hex](https://hex.pm/) as a package manager and [hex.pm](https://hex.pm/) as a package registry and does not support any file formats.

## Frameworks and libraries

Snyk for Elixir does not have any available frameworks and libraries.&#x20;

## Features

The Reports feature is supported by Snyk for Elixir.&#x20;

## Snyk CLI for Elixir

{% hint style="info" %}
To scan your dependencies, you must first install Elixir and Mix. For details, [see the Elixir installation instructions](https://elixir-lang.org/install.html).
{% endhint %}

Snyk offers security scanning to test your Elixir Projects for vulnerabilities using the [CLI](../snyk-cli/).

Mix is a build tool that compiles, tests, and creates Elixir projects. Mix manages dependencies by integrating with the Hex package manager.

Snyk builds a dependency tree for your Project by analyzing your `mix.exs` and `mix.lock` files. The `mix.lock` file must be present and in sync with the `mix.exs` file. After Snyk builds the tree, Snyk uses the [vulnerability database](https://snyk.io/vuln) to find vulnerabilities in the packages anywhere in the dependency tree.

### **Project naming**

Projects in the Snyk UI are named according to the `app` keyword from the `project/0` function exported by `Mix.Project` in the main `mix.exs` file.

To override the name, use the `--project-name` CLI option.

### **Mix umbrella projects**

If you test a Mix umbrella project, Snyk detects that it is an umbrella project and includes all the child apps automatically.

Along with the main `mix.exs`, each app `mix.exs` appears as a separate Project in the Snyk UI, named according to the path to the app.

Snyk fully supports all `:hex` packages listed in the Mix project, including all their transitive dependencies and any vulnerabilities.

Hex support includes both Elixir and Erlang packages.

Snyk also has limited support for `:path`, `:git` and `:github` dependencies, but not their transitive dependencies or vulnerabilities.

* `:path` dependencies appear in the dependency tree by name
* `:git` and `:github` dependencies appear in the dependency tree by repository URL and version (either `:branch`, `:tag` or `:ref`, as defined in the `mix.exs` file)

## IDE and CI/CD

For integrated development environments, see [Snyk IDEs](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/).

If you use continuous integration/continuous delivery workflows, you can scan with Snyk based on the integration with your automation software.

## Troubleshooting Snyk for Elixir

When using `asdf`, be sure to set a version by running the `asdf global elixir <version of your choice>`.

If you need help, [contact Snyk Support](https://support.snyk.io/hc/en-us).
