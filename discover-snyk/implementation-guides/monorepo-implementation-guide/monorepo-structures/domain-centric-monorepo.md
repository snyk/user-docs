---
description: How to scan a Domain-centric monorepo with Snyk, where code is grouped by business domain and team
nav_context: agnostic
---

# Domain-centric monorepo

```
Root/
├── inventory-domain/          # Team A
│   ├── ci.inventory.yml       # Pipeline: owns test, build, and deploy for this folder
│   ├── inventory-ui/          # Micro frontend
│   └── inventory-svc/         # Microservice
├── checkout-domain/           # Team B
│   ├── ci.checkout.yml        # Pipeline: owns test, build, and deploy for this folder
│   ├── checkout-ui/           # Micro frontend
│   └── checkout-svc/          # Microservice
├── shared-platform/           # Global design system
└── renovate.json              # Global dependency automation
```

## Layout

* **Organized by business capability**: Focused on which team owns the code, for example Payments versus Inventory.
* **High cohesion and low coupling**: Code for a specific feature, frontend and backend, stays together in one domain folder to minimize cross-team friction.
* **Shared core**: Usually includes a small `shared/` or `platform/` folder for cross-cutting concerns such as authentication or logging.
* **Distributed pipelines**: Each domain is an island. A change in `inventory-domain/` triggers only the `ci.inventory.yml` pipeline.
* **Isolated source code**: Source code in `inventory-domain` rarely references code in `checkout-domain`. When this holds, you can scan code within the domain itself, because the source and sink are in the same domain.
* **Typical consumption model**: Published to a registry.

## Snyk best practices

* **Targeted execution**: Make sure the pipeline runs only when domain files change. To maintain isolation and speed, `cd` into the domain directory before running Snyk commands.
* **Domain pull request gate**: In the domain-specific pipeline, run blocking `snyk test` and `snyk code test`. This prevents a team from introducing risk within its own domain boundary.
* **Shared platform registry gate**: If any file in `shared-platform/` is modified, run `snyk code test` and `snyk test` on that folder during the publish to the registry.

{% hint style="info" %}
This scan identifies insecure logic (sinks) in the library in isolation. It does not trace a full source-to-sink path from a consuming domain. Domains consume these libraries as published artifacts, for example through Artifactory, rather than as raw source code. Root-level Snyk Code scans are ineffective in this case, because Snyk cannot trace data flow into an opaque artifact. This gate prevents potential sinks from being pushed to the registry, where they would be invisible to a domain's local Snyk Code analysis.
{% endhint %}

* **Snyk Code reporting**: If domain-specific code changes, run the following command within that domain's pipeline. This scopes static analysis results to the domain's assets and pushes them to the Snyk web UI with the correct product context.

  ```bash
  snyk code test --report --project-name="<DOMAIN_NAME>/<SERVICE>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Snyk Open Source monitoring**: Run the following command within the domain pipeline to create a persistent heartbeat for that domain's production dependencies.

  ```bash
  snyk monitor --project-name="<DOMAIN_NAME>/<SERVICE>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Unified Organization strategy**: Keep all Project data for the monorepo in a single, dedicated Snyk Organization (`ORG_ID`).
* **Naming precision**: Set `--project-name` to the domain name, for example `Inventory` or `Checkout`, to distinguish assets in the Snyk web UI.
