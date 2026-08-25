---
description: The snyk agent test command that runs Open Source, Code, and Secrets scans together
---

# Snyk agent test

{% hint style="warning" %}
`snyk agent test` is experimental. Commands, flags, and output format may change in future releases.
{% endhint %}

## Prerequisites

- Snyk CLI v1.1307.0 (or later).

## Usage

`$ snyk agent test [<OPTION>]`

## Description

The `snyk agent test` command runs Snyk Open Source, Snyk Code, and Snyk Secrets scans together in a single pass, returning token-optimized output intended for consumption by AI coding agents rather than direct human review.

## Exit codes

Possible exit codes and their meaning:

**0**: success (scan completed), no issues found.\
**1**: action_needed (scan completed), one or more issues found.\
**2**: failure, try to re-run the command. Use `-d` to output the debug logs.

## Debug

Use the `-d` or `--debug` option to output the debug logs.

## Options

_To be confirmed with engineering — for example org scoping (`--org=<ORG_ID>`), severity threshold, and output format flags._
