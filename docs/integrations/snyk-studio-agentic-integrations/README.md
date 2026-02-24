# Agentic security with Snyk Studio

Agentic workflows accelerate software development and innovation but introduce security risks, for example, AI-generated vulnerabilities or providing insecure practice suggestions. To address these risks, Snyk Studio embeds directives directly into AI-assisted development workflows to guide the agent towards secure software development.

Snyk Studio connects the Snyk platform, your development environment, and your AI tools, in four interconnected layers:

* [Snyk MCP Server](../../discover-snyk/getting-started/glossary.md#snyk-mcp-server): An MCP server that enables easy integration with coding assistants, providing security context to AI agents. This runs locally using the Snyk CLI.
* AI agents or [ADE's](../../discover-snyk/getting-started/glossary.md#ade): Agentic development environments that enable developers to interact with coding agents, which can integrate with Snyk Studio to drive security.
* [Directives](../../discover-snyk/getting-started/glossary.md#directive): Policies, rules and commands provided to coding assistants to guide them on coding and security standards.
* The Snyk platform: Security intelligence from Snyk products (Snyk Code, Snyk Open Source, and more) which can be used to identify security issues and drive resolutions.

## MCP Server supported tools

Snyk Studio supports integrating the following Snyk security tools into an AI system:

* `snyk_sca_scan` (Open Source scan)
* `snyk_code_scan` (Code scan)
* `snyk_iac_scan` (IaC scan)
* `snyk_container_scan` (Container scan)
* `snyk_sbom_scan` (SBOM file scan)
* `snyk_aibom` (create AI-BOM)
* `snyk_trust` (trust a given folder before running a scan)
* `snyk_auth` (authentication)
* `snyk_logout` (logout)
* `snyk_version` (version information)
* `snyk_send_feedback` (summarizing issues fixed)
* `snyk_package_health_check` \[Experimental]\(evaluate package health during dependency selection)

{% hint style="info" %}
Running `snyk_sca_scan` can execute third-party ecosystem tools (for example, Gradle or Maven) on your machine to fetch the project's dependency tree.
{% endhint %}

## Coding assistant support

Snyk supports any AI agent and ADE that integrates with a local MCP server and offers [quickstart guides](quickstart-guides-for-snyk-studio/) for some of these to provide additional guidance.

## Resources

Snyk provides a general setup guide with instructions applicable to all coding assistants. To learn more, visit [Getting started with Snyk Studio](getting-started-with-snyk-studio.md).

To learn about Snyk directives which govern agent behaviour and are automatically injected into agent interactions, visit [Guardrail Directives](directives.md#guardrail-directives). To learn about Snyk directives which can be manually invoked by human developers or AI agents, visit [Command Directives](directives.md#command-directives).

Snyk Studio supports large-scale distribution and centralized governance. For guidance on how to handle scaling across your organization, visit [Distribution at scale](distribution-at-scale.md).
