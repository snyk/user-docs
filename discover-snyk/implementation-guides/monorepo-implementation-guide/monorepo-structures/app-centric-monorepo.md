---
description: How to scan an App-centric monorepo with Snyk, where code is grouped by deployment target
nav_context: agnostic
---

# App-centric monorepo

```
Root/
├── toolchain/                 # Shared scripts
│   ├── web-build-logic.sh     # Global micro frontend build script
│   └── svc-build-logic.sh     # Global container build script
├── web-apps/                  # Target: CDN or Amazon S3
│   ├── pipeline.web.yml       # Pipeline: runs web-build-logic.sh for all micro frontends
│   ├── store-front/           # Micro frontend
│   └── admin-panel/           # Micro frontend
├── services/                  # Target: Kubernetes or Fargate
│   ├── pipeline.svc.yml       # Pipeline: runs svc-build-logic.sh for all services
│   ├── order-api/             # Microservice
│   └── email-worker/          # Worker
└── shared-packages/           # Internal NuGet, npm, or source packages
```

## Layout

* **Organized by deployment target**: Grouped by where the code runs, for example browser or Amazon S3 versus server or Docker.
* **Shared infrastructure**: Common build logic and CI scripts are usually in a `toolchain/` or `eng/` folder.
* **Environment-specific logic**: All projects in a folder, such as `services/`, share the same base images, runtime versions, and deployment scripts.
* **Batch pipelines**: Often one pipeline per platform. One CI run handles all micro frontends and another handles all microservices.
* **Isolated source code**: Source code is likely isolated between platforms, because the technical stacks and runtimes, for example browser logic versus server logic, are fundamentally different and do not share execution paths.
* **Typical consumption model**: Published to a registry.

## Snyk best practices

* **Platform-level triggers**: Configure pipelines to trigger only when files in the specific platform directory, for example `web-apps/` or `services/`, are modified.
* **Targeted directory scanning**: For each changed application or service, `cd` into the project directory so Snyk accurately identifies the local manifest files and source code.
* **Shared package registry gate**: If any file in `shared-packages/` is modified, run a blocking `snyk code test` and `snyk test` on that folder during its publish pipeline.

{% hint style="info" %}
This scan identifies insecure logic (sinks) in the package in isolation. It does not trace a full source-to-sink path from a consuming application. Applications consume these packages as published artifacts rather than as raw source code. Root-level Snyk Code scans are ineffective in this case, because Snyk cannot trace data flow into an opaque artifact. This gate prevents potential sinks from being published to the internal registry, where they would be invisible to an application's local Snyk Code analysis.
{% endhint %}

* **Snyk Code reporting**: If source files do not reference each other, run the following command within the individual platform pipelines only. This keeps reporting scoped to the specific application logic and avoids large, unnecessary root-level scans.

  ```bash
  snyk code test --report --project-name="<APP_CATEGORY>/<APP_NAME>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Snyk Open Source monitoring for deployables**: Run the following command within individual service or micro frontend pipelines. This provides a persistent heartbeat for the version of the code currently deployed.

  ```bash
  snyk monitor --project-name="<APP_CATEGORY>/<APP_NAME>" --target-reference="$(git branch --show-current)" --org=<ORG_ID>
  ```
* **Unified Organization strategy**: Keep all Project data for the monorepo in a single, dedicated Snyk Organization (`ORG_ID`).
* **Naming and tracking**: Use unique `--project-name` values, for example `web-apps/store-front`, to differentiate assets in the Snyk web UI.
