# Activation and deployment

You configure and deploy Agentic Development Security (ADS) from **Settings** in Evo. Choose which products to roll out, then install on a single machine or across your company through your mobile device management (MDM) tool.

## Authenticate

ADS uses a push key to bind installed agents to your Snyk Tenant, so each machine's data lands in your Tenant. You can rotate the key at any time.

## Choose products

{% hint style="info" %}
**Feature availability**

**Observe** (Agent Behavior Governance) is in open preview. As such it is not connected to the platform surfaces (Inventory, Policies & issues, Reports, and Evo chat) and comes with limitations described here.
{% endhint %}

Under **Products**, select the products to roll out:

* **Machines** (Agent Supply Chain Security)
* **Observe** (Agent Behavior Governance)
* **Snyk Studio** (Trusted Output Assurance)

By default, each product stays on the latest version and updates automatically.

## Install

Follow the on-screen instructions to install on a local machine or through your MDM tool.

### What the installer does

When the installer runs, it sets up each selected product on the machine:

* **Agent Supply Chain Security** discovers the skills and MCP servers in the known directories, performs a risk assessment, and sends the results to the Evo Tenant associated with the push key.
* **Agent Behavior Governance** checks whether the supported agents have the required hooks configured. If they do not, it writes them with the push key, so subsequent agent activity is pushed to Evo and evaluated against your Tenant's policies.
* **Trusted Output Assurance** checks whether the required configuration is in place — the Snyk CLI and its MCP server, the Secure at inception hooks, and the package health check, fix commands, and skills — and adds anything missing, so subsequent agent activity runs through the Secure at inception loop. Each developer still authenticates the Snyk CLI and MCP server individually.
