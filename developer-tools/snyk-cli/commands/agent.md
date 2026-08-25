---
description: The snyk agent command space, a scanning surface built for AI coding agents
---

# Snyk agent

{% hint style="warning" %}
`snyk agent` is experimental. Commands, flags, and output format may change in future releases.
{% endhint %}

## Prerequisites

- Snyk CLI v1.1307.0 (or later).

## Usage

`$ snyk agent <COMMAND> [<OPTION>]`

**See also:** [`snyk agent test`](agent-test.md) — run Snyk Open Source, Code, and Secrets scans together.

## Description

The `snyk agent` command space is a scanning surface designed for AI coding agents rather than human readers. It provides token-optimized output and ergonomics so agentic tools can request a scan and parse the result efficiently.

## Debug

Use the `-d` or `--debug` option to output the debug logs.
