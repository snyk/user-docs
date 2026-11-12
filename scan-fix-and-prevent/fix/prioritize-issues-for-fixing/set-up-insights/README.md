---
description: Snyk retired Insights. Its risk factors and settings are no longer available.
nav_context: classic
last_reviewed: 2026-10-09
last_reviewed_source: bulk-sync
---

{% include "../../../.gitbook/includes/new-navigation-banner.md" %}

# Insights (end of life)

{% hint style="warning" %}
Snyk has retired Insights. The Insights risk factors and the Insights settings are no longer available.
{% endhint %}

Insights computed the **Deployed**, **Loaded package**, **OS condition**, and **Public facing** risk factors from your container images, packages, and Kubernetes resources. Snyk no longer computes these risk factors.

If you installed the Kubernetes connector for Insights, uninstall it from your clusters. For instructions, visit [Insights Kubernetes connector (end of life)](set-up-insights-kubernetes-connector.md).

To prioritize your issues, use the [Risk Score](../risk-score.md) and [Prioritization for Snyk Essentials](../prioritization-for-snyk-essentials.md).
