# Claude Code guide

You can add the Snyk MCP server to Claude Code to secure code generated with agentic workflows through an LLM. This integration enables real-time security analysis and vulnerability scanning during development. When you use it for the first time, the MCP server will ask for trust and trigger authentication if necessary.

## Prerequisites

* [Install Claude Code CLI](https://docs.anthropic.com/en/docs/claude-code/setup)
* [Install the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/)
* [Install the Snyk MCP](claude-code-guide.md#install-the-snyk-mcp-server-in-claude-code)

### Install Claude Code CLI

Install Claude Code CLI using npm:

```bash
npm install -g @anthropic-ai/claude-code
```

Ensure you have a working Claude Code environment with proper authentication.

### Install the Snyk MCP Server in Claude Code

Install the Snyk MCP Server using the method that best suits your operating system and local development environment.

#### Install with Node.js and `npx` (Recommended)

Use the Claude Code CLI to add the Snyk MCP server:

```bash
# Add Snyk MCP server using npx
claude mcp add snyk -- npx -y snyk@latest mcp -t stdio

# Verify MCP server is registered
claude mcp list
```

#### Install with pre-installed Snyk CLI

If you have the Snyk CLI installed and accessible on your system path, you can use the local installation:

```bash
# First, find your Snyk CLI path
which snyk

# Add Snyk MCP server using local CLI
claude mcp add snyk -- /path/to/snyk mcp -t stdio

# Verify MCP server is registered
claude mcp list
```

#### Alternative: `.mcp.json` Configuration

Create or edit the file `.mcp.json` in the root directory of your project:

```json
{
  "mcpServers": {
    "snyk": {
      "command": "npx",
      "args": [
        "-y",
        "snyk@latest",
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}
```

For pre-installed Snyk CLI:

```json
{
  "mcpServers": {
    "snyk": {
      "command": "/absolute/path/to/snyk",
      "args": [
        "mcp",
        "-t",
        "stdio"
      ]
    }
  }
}
```

If the `snyk` command is not available, add it by following the instructions on the [Installing or updating the Snyk CLI](../../../developer-tools/snyk-cli/install-or-update-the-snyk-cli/) page.

## Setting up the Snyk MCP Server

As a one-time setup, you may need to authenticate and trust the current project directory. Claude Code can handle this automatically through the MCP integration.

### First-Time Authentication

```bash
# Check authentication status through Claude Code
claude "Check Snyk authentication status"

# If authentication is required, Claude can authenticate via MCP
claude "Authenticate with Snyk"
```

### Trust Your Project Directory

```bash
# Trust your project directory via MCP
claude "Trust this directory for Snyk scanning"
```

{% hint style="info" %}
These MCP commands are executed directly in Claude Code - no need for external CLI setup. Claude handles authentication and trust automatically through the MCP integration.
{% endhint %}

## Examples

### Scanning code and dependencies for security vulnerabilities

You can ask Claude Code to perform comprehensive security scans of your codebase:

{% code title="prompt" overflow="wrap" %}
```
Scan this directory for code security & dependency vulnerabilities and security issues
```
{% endcode %}

Claude Code will use the Snyk MCP Server to perform various security scans and provide detailed vulnerability reports.

### Security Code Reviews

Request security-focused code reviews with vulnerability context:

{% code title="prompt" overflow="wrap" %}
```
Review this codebase for security issues using Snyk data
```
{% endcode %}

### Dependency-Specific Analysis

Check specific dependencies for known vulnerabilities:

{% code title="prompt" overflow="wrap" %}
```
Check if the axios dependency has any known vulnerabilities
```
{% endcode %}

### AWS Infrastructure Security

Scan cloud infrastructure code for security best practices:

{% code title="prompt" overflow="wrap" %}
```
Analyze this AWS CloudFormation template for security best practices and known vulnerabilities
```
{% endcode %}

### Pre-Commit Security Checks

Integrate security scanning into your development workflow:

{% code title="prompt" overflow="wrap" %}
```
Scan my staged changes for security vulnerabilities before I commit
```
{% endcode %}

## Claude Code Security Best Practices

To practice [Secure at Inception](https://snyk.io/solutions/secure-ai-generated-code/), Snyk recommends adopting workflows that align Claude Code for secure code generation.

### Recommended Security Workflow

When working with Claude Code, follow these security guidelines:

{% code title="Security workflow prompts" overflow="wrap" %}
```
Always run Snyk Code scanning for new first-party code generated.
Always run Snyk SCA scanning for new dependencies or dependency updates.
If any security issues are found in newly introduced or modified code or dependencies, attempt to fix the issues using the results context from Snyk.
Rescan the code after fixing issues to ensure they were resolved and no new issues were introduced.
Repeat this process until no issues are found.
```
{% endcode %}

### Integration with Git Workflows

You can integrate security scanning into your Git workflow:

```bash
# Security review before committing
git add . && claude "Security review of staged changes" && git commit

# Pre-push security validation
claude "Perform final security scan before push" && git push
```

### Advanced Configuration

For projects requiring specific security policies or custom scanning configurations, you can enhance your `.mcp.json` with environment variables:

```json
{
  "mcpServers": {
    "snyk": {
      "command": "npx",
      "args": [
        "-y",
        "snyk@latest",
        "mcp",
        "-t",
        "stdio"
      ],
      "env": {
        "SNYK_TOKEN": "your-snyk-token-here",
        "SNYK_CFG_ORG": "your-org-id"
      }
    }
  }
}
```

{% hint style="warning" %}
Never commit sensitive tokens or credentials to version control. Use environment variables or secure secret management solutions.
{% endhint %}