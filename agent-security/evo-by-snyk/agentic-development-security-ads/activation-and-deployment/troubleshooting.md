---
description: >-
  Errors you can encounter when installing or removing ADS, and how to resolve
  them
---

# Troubleshooting

Use this page to resolve errors when you install or remove Agentic Development Security (ADS). To confirm that an install succeeded, visit [Verify an installation](verify-an-installation.md).

## Manifest not found, please make sure the configuration is published via the settings page

The installer could not find a published configuration for your Tenant. Open **Settings**, confirm your product selection, and click **Save & Publish**. Then run the install again.

## The install succeeded but no products are on the machine

On macOS, the package delivers the installer only. If onboarding failed, for example because of an incorrect push key, the installer is on the machine and the products are not. Check the output from that installer run for the error, then run the installer again with the correct credentials. If the error is not clear, collect the output and your Tenant ID, and contact your Snyk account team.

## Authentication fails outside the United States

A push key authenticates only against the region that issued it. If you use another region's host, authentication fails. The installer does not fall back to another region. Set the base URL to match the URL you log in with, as listed in [If you log in outside the United States](./#if-you-log-in-outside-the-united-states).

## `uninstall` returns an error and removes nothing

`uninstall` requires either `--all` or at least one of `--scan`, `--guard`, and `--studio`. `--scan` also requires `--tenant-id`. If the original install used a custom installation directory, pass the same directory with `--install-dir <DIRECTORY>`.

## `uninstall` does not remove a managed install

`uninstall` assumes user mode. To remove a managed install, run it with elevated privileges and add `--install-mode managed`, as described in [Remove a managed install](./#remove-a-managed-install).

## A machine stops sending data after you revoke a push key

A machine that still uses a revoked push key is rejected at its next check-in. Create a new push key if you do not have another active one, then redeploy the installer to the machine through your MDM tool with that key. To avoid this, revoke the old key only after the fleet has moved to the new key, as described in [Rotate a push key](./#rotate-a-push-key).

## Agent Behavior Governance does not govern Codex on a machine

In user mode, Codex prompts the developer after installation to allow hook configuration outside the managed path. If the developer does not accept the prompt, the hooks stay installed on the machine but Codex does not run them, so Agent Behavior Governance does not govern Codex on that machine. Codex warns the developer that hooks need review each time it starts, until the hook is trusted.

To fix this on a machine that is already installed, the developer reviews and trusts the Snyk hook manually in Codex. The hook does not need to be installed again.

Codex trusts a hook by its exact definition, so Codex asks the developer to review the hook again if the definition changes.

To avoid the prompt for a fleet, install in managed mode, as described in [Installation mode](README.md#installation-mode).
