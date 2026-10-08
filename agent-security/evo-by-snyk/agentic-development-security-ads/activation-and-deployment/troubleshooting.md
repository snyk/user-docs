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
