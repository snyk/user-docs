---
description: >-
  How to confirm ADS installed on a machine, and what to check when the results
  are not what you expect
---

# Verify an installation

Confirm that Agentic Development Security (ADS) installed correctly on a machine: check that the installer is present, then confirm that the machine appears in Evo.

## Check that the installer is present

These checks apply to package installs, which leave the installer at a fixed path.

{% tabs %}
{% tab title="macOS" %}
```bash
/usr/local/bin/snyk-ads/snyk-ads-installer --version
pkgutil --pkg-info io.snyk.ads-installer
```
{% endtab %}

{% tab title="Windows" %}
```powershell
& "C:\Program Files\Snyk\ADS Installer\snyk-ads-installer-windows-x86_64.exe" --version
```
{% endtab %}
{% endtabs %}

A command-line install on macOS or Linux deletes the installer when it finishes, so there is nothing to check on the machine. Confirm those installs in Evo instead.

## If a machine is missing from Evo

The first scan runs when the installer runs, so a machine appears in Evo once that run finishes. If it does not, check the causes in [Troubleshooting](troubleshooting.md): an unpublished configuration, a failed onboarding step, or the wrong region.

If none of these applies, collect the output from the installer run and your Tenant ID, and contact your Snyk account team.
