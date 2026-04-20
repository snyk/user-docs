# Connect your development tools

To roll out Snyk efficiently at scale, your first major milestone is configuring a Template Organization. Rather than setting up every new team from scratch, this template acts as your master blueprint. By configuring your core tools, source control integrations, and default security behaviors here first, you establish a standardized baseline that can be easily cloned, either manually or using the API, across your entire business.

As you work through this page, you will:

1. [Configure SCM integrations for Groups.](connect-your-development-tools.md#configure-scm-integrations-for-groups)
2. [Configure SCM integrations for Organizations.](connect-your-development-tools.md#configure-scm-integrations-for-organizations)
3. [Set up Container registries.](connect-your-development-tools.md#set-up-container-registry-integrations)
4. [Set up additional integrations.](connect-your-development-tools.md#set-up-additional-integrations)
5. [Configure Universal Broker.](connect-your-development-tools.md#configure-the-universal-broker-connection)
6. [Enable Snyk Code.](connect-your-development-tools.md#enable-snyk-code)

## Set up your SCM integrations

When implementing your SCM integrations, you must configure connections at two Snyk levels: **Group** and **Organization**.

The following table outlines the difference between these two implementation levels.

| Category                 |                         Group level SCM                        |                            Organization level SCM                            |
| ------------------------ | :------------------------------------------------------------: | :--------------------------------------------------------------------------: |
| **Description**          |               Monitors your entire SCM landscape               |     Created so specific teams can actively find and fix vulnerabilities.     |
| **Primary purpose**      |                     Global asset discovery                     |   Running Snyk security tests and automatically raising Fix Pull Requests.   |
| **Required permissions** |                            Read only                           |                                Read and write                                |
| **Scope**                | Access to all repositories in your SCM, both new and existing. | Restrict access to repositories owned by that specific Organization or team. |

### Configure SCM integrations for Groups

Configure your Source Control Manager (SCM) at the Group level to centralize authentication. This allows Snyk to access your repositories across multiple Organizations without requiring individual configuration for each one.

{% embed url="https://res.cloudinary.com/snyk/video/upload/v1775661962/3._Group-Level_Repository_Discovery_jvicn7.mp4" %}
Group level repository discovery video guide
{% endembed %}

{% stepper %}
{% step %}
#### Select your SCM platform

{% hint style="success" %}
**Key decision:** Choose the SCM platform that hosts your primary development work and determine if you require **Snyk Broker** for an on-premise connection.
{% endhint %}

1. In the Snyk web UI, navigate to **Group Settings** > **Integrations**.
2. Select your SCM platform (for example, GitHub, GitLab, or Azure Repos).

If your SCM is behind a firewall, you must install and configure Snyk Broker to establish a secure, outbound-only connection.
{% endstep %}

{% step %}
#### Authenticate the integration

{% hint style="success" %}
**Key decision:** Determine which service account or administrative user will provide the initial OAuth or Personal Access Token (PAT) to ensure the connection remains stable.
{% endhint %}

1. Follow the prompts to authorize Snyk to access your SCM organization.
2. Use a service account: Snyk recommends using a dedicated service account rather than a personal user account to prevent integration failure if an individual leaves the company.
3. Grant the required permissions for Snyk to read manifest files and, if desired, create pull requests (PRs) for security fixes later in your rollout.
{% endstep %}

{% step %}
#### Align with your Organization structure

{% hint style="success" %}
**Key decision:** Decide if you will use a single Group-level integration or if specific Organizations require separate credentials based on your established hierarchy.
{% endhint %}

Your integration approach should match the structure you selected during the planning phase:

* **SCM organization-based structure**: If you have 10+ SCM organizations and are using the `snyk-api-import` tool, ensure your Group-level integration has the scope to access all relevant repositories.
* **Team-based structure**: Use Group-level authentication to ensure consistency across different developer teams while maintaining isolated Organizations.
* **Product-based structure**: Centralizing at the Group level allows developers to seamlessly access Projects across multiple product Organizations.

| Structure type    | Integration strategy                                                        |
| ----------------- | --------------------------------------------------------------------------- |
| **SCM-based**     | Use one Group integration to discover and import all SCM organizations.     |
| **Team-based**    | Use Group-level authentication for centralized credential management.       |
| **Product-based** | Centralize at the Group level to simplify access for full-stack developers. |
{% endstep %}
{% endstepper %}

### Configure SCM integrations for Organizations

Configure SCM integrations at the Organization level to establish granular connections for specific teams, products, or business units. While Group-level integrations provide a global baseline, Organization-level settings allow for isolated credentials and team-specific automation.

{% hint style="info" %}
If you are using multiple SCMs, Snyk recommends using separate Organizations for separate SCM integrations.
{% endhint %}

{% stepper %}
{% step %}
#### Establish granular authentication

{% hint style="success" %}
**Key decision:** Determine if this specific Organization requires a unique access token or a different service account than the one used at the Group level.
{% endhint %}

Unlike Group-level setup, Organization-level integrations allow you to:

* **Isolate access:** Use a unique Personal Access Token (PAT) or OAuth connection that only has access to a specific team's repositories.
* **Override Group defaults:** If a specific business unit uses a different SCM instance (for example, a separate GitHub Org or GitLab Group), you can configure it here without affecting the rest of the company.

Set up your Org-level integrations by navigating in your Organization to the **Integrations** page and selecting the relevant SCM tile.
{% endstep %}

{% step %}
#### Consider specific Snyk Broker tokens

{% hint style="success" %}
**Key decision:** Identify if this Organization requires a dedicated Snyk Broker token to segment network traffic or satisfy distinct security requirements.
{% endhint %}

When configuring at the Organization level, consider:

* **Token segmentation:** You can assign a unique Broker token to an Organization. This is useful if different teams operate in different VPCs or data centers.
* **Granular troubleshooting:** Dedicated tokens allow you to monitor and troubleshoot connectivity issues for a specific team without impacting the entire Group.

{% hint style="info" %}
If you are using Azure Repos, Snyk recommends using Universal Broker to avoid Azure limitations on organization mapping.
{% endhint %}
{% endstep %}

{% step %}
#### Define team-specific automation

{% hint style="success" %}
**Key decision:** Decide which PR check behaviors and fix strategies apply to this team’s specific development workflow.
{% endhint %}

While the Group integration provides the connection, the Organization level is where you define how Snyk interacts with your code, through:

* **Targeted PR checks:** Enable or disable PR checks for specific Organizations based on the team's maturity level.
* **Customized fix PRs:** Configure whether Snyk should automatically open fix PRs for this Organization's Projects, allowing for a phased rollout of "prevention" features.

| Feature             | Organization-level specific benefit                                    |
| ------------------- | ---------------------------------------------------------------------- |
| Access control      | Limits repository visibility to only the members of that Organization. |
| Snyk Broker mapping | Maps specific SCM instances to specific internal network segments.     |
| Customized rollout  | Allows for different "prevention" settings (PR checks) per team.       |
{% endstep %}
{% endstepper %}

## Set up Container registry integrations

Integrate Snyk with your container registries to import and monitor images for known vulnerabilities. This ensures that the base images and layers used in your deployments are continuously scanned against the Snyk vulnerability database.

{% stepper %}
{% step %}
### Select your registry provider

{% hint style="success" %}
**Key decision:** Identify which container registries (for example, Docker Hub, Amazon ECR, Google Artifact Registry) host your production-ready images and determine if they reside behind a firewall.
{% endhint %}

1. In the Snyk web UI, navigate to **Integrations** > **Container Registries**.
2. Select your specific registry provider.

If your registry is on-premise or behind a firewall, you must use Snyk Broker to establish a secure connection.
{% endstep %}

{% step %}
### Authenticate and authorize

{% hint style="success" %}
**Key decision:** Use a dedicated service account with read-only permissions to the registry to maintain a stable connection and follow the principle of least privilege.
{% endhint %}

1. Provide the required credentials (for example, Access Keys, Role ARNs, or JSON keys) as specified by your registry provider.
2. Ensure the account has sufficient permissions to list and pull images for scanning.

{% hint style="info" %}
For Amazon ECR, Snyk recommends using Cross-Account Role authentication for enhanced security.
{% endhint %}
{% endstep %}

{% step %}
### Configure scan frequency and visibility

{% hint style="success" %}
**Key decision:** Decide on a monitoring frequency that balances security visibility with your team's remediation capacity.
{% endhint %}

Once integrated, Snyk allows you to manage how often images are re-tested:

* **Continuous monitoring**: Snyk automatically rescans imported images daily to detect new vulnerabilities.
* **Avoid duplication**: If you already scan images in your CI/CD pipeline using the Snyk CLI, decide if you also need registry-level monitoring. Registry integration provides a last line of defense for images currently in storage.

As with SCM integrations, ensure email notifications are disabled at the Organization level during the initial bulk import of images to prevent alert fatigue.

| Integration level  | Recommended for    | Advantage                                                          |
| ------------------ | ------------------ | ------------------------------------------------------------------ |
| Container registry | Security teams     | Visibility into all stored images and production snapshots.        |
| CI/CD pipeline     | DevOps/Developers  | Catching vulnerabilities before the image is pushed to a registry. |
| Kubernetes         | Platform engineers | Monitoring vulnerabilities in active, running workloads.           |
{% endstep %}
{% endstepper %}

### Use cases for direct registry integration

Snyk strongly recommends integrating your container registry if your organization fits any of the following profiles:

* **Empowering AppSec autonomy**: This solution enables Application Security (AppSec) teams to manage and scan container images independently. It allows security teams to maintain oversight without bottlenecking developers or requiring them to change existing workflows.
* **Scanning legacy and offline images:** If you have a large number of older images or multiple legacy tags that no longer have an active build pipeline, registry integration is the most efficient way to ensure they are still scanned for vulnerabilities.
* **Small-scale repositories:** Direct integration and UI management work best for smaller volumes of images (under 1,000). A smaller inventory makes the import process seamless and keeps the user interface fast and responsive.
* **Limited developer bandwidth:** If your development team lacks the availability to update build pipelines or integrate new security tools into their current processes, direct registry scanning provides a frictionless alternative.

### Use cases for CI/CD or CLI

Direct registry integration is not a one-size-fits-all solution. Consider pivoting to pipeline or command-line testing under the following conditions:

* **High-volume registries:** If your repository exceeds 1,000 images, manually selecting and managing images using the UI can become slow over time.
* **Active, high-velocity workflows:** In practice, many teams skip direct registry integration for their active builds. Instead, they integrate security scanning directly into their CI/CD pipelines or use CLI testing. This "shift-left" approach is often preferred for actively developed applications because it catches vulnerabilities during the build process rather than after the image has been pushed to the registry.

## Set up additional integrations

Configure additional integrations to build a complete inventory of your code-based assets. By connecting Snyk to your broader ecosystem, you can identify coverage gaps and ensure all repositories have the necessary security controls in place.<br>

{% stepper %}
{% step %}
### Access the inventory

{% hint style="success" %}
**Key decision**: Determine if you have the necessary Group Administrator or **Edit Essentials** permissions to manage the global asset inventory.
{% endhint %}

To start building your inventory:

1. In the Snyk web UI, navigate to your Group.
2. Select **Inventory** from the side menu.
3. Navigate to **Integrations** to view your existing connections or add new ones.
{% endstep %}

{% step %}
### Configure SCM integrations for asset discovery

{% hint style="success" %}
**Key decision**: Decide whether to use a broad-access service account token to ensure Snyk can discover all repositories across your development teams.
{% endhint %}

This configuration is specific to asset management and is separate from the Organization-level integrations used for security scanning.

1. Click **Add integration** and choose your SCM provider (GitHub, GitLab, Azure DevOps, or Bitbucket).
2. Apply Group-level tokens: Snyk recommends setting tokens at the Group level. This provides a comprehensive view of all repositories and ensures alignment between security and development teams.
3. Handle large environments: If you have more than 1,000 repositories or encounter API rate limits, deploy a dedicated Snyk Broker instance for asset management.
{% endstep %}

{% step %}
### Define application context and tags

{% hint style="success" %}
**Key decision**: Choose which metadata (tags) and application structures are most critical for your risk assessment and reporting.
{% endhint %}

Once assets are imported, use the following features to organize your inventory:

* **Automatic tags**: Snyk automatically adds tags for detected technologies (for example, Python, Terraform).

{% hint style="info" %}
Bitbucket users must add language tags manually.
{% endhint %}

* **Coverage cap filters**: Use the **Coverage** and **Coverage Gap** filters to identify which repositories are missing security products like Snyk Code or Snyk Open Source.
* **Asset hierarchy**: View your assets in a nested structure to see how packages relate to their parent repositories.

| Feature             | Purpose                                                                             |
| ------------------- | ----------------------------------------------------------------------------------- |
| Asset dashboard     | Provides a high-level view of issue trends and control coverage.                    |
| Technology grouping | Organizes assets by the programming languages or tools detected.                    |
| Team mapping        | Groups repositories based on your SCM team structure for easier ownership tracking. |
{% endstep %}
{% endstepper %}

## Configure brokered connections

Snyk Universal Broker improves the management of Broker deployments by supporting multiple connections of any type with a single running client (or replica group). Credentials remain in your network and are securely referenced without being transmitted to Snyk. Use Universal Broker to enable Snyk to scan repositories, container registries, or other assets hosted behind your firewall without exposing your internal network.

{% hint style="info" %}
If you prefer to use helm charts for Broker configuration or are configuring a CR agent, Snyk recommends you use [Classic Broker](../../enterprise-setup/snyk-broker/classic-broker/).
{% endhint %}

{% stepper %}
{% step %}
### Determine deployment requirements

{% hint style="success" %}
**Key decision**: Determine your redundancy strategy. While a single Universal Broker instance can manage multiple integrations, Snyk recommends configuring at least two replicas of the client for high availability.
{% endhint %}

Before installation, verify your environment:

* You can run Universal Broker as a Docker container or a Kubernetes pod using Helm. This requires at least 1 CPU and 256 MB RAM.
* Ensure the Broker has outbound HTTPS access (port 443) to [https://broker.snyk.io](https://broker.snyk.io/). No inbound ports need to be opened on your firewall.
* You must have Node.js (version 20 or higher) installed to run the setup CLI, a personal Snyk API token, and **Snyk Tenant Admin** permissions.
{% endstep %}

{% step %}
### Configure the Universal Broker connection

{% hint style="success" %}
**Key decision**: Choose whether to link the Broker token to the Group level for broad asset discovery or to a specific Organization for isolated team access. Universal Broker uses a CLI tool to dynamically configure connections rather than generating a static Broker token in the UI.
{% endhint %}

1. Install the setup CLI tool: `npm install -g snyk-broker-config`
2. Start the interactive creation workflow: `snyk-broker-config workflows connections create`
3. Follow the on-screen prompts to input your Snyk Token, Tenant or Org ID, and select the specific integration type you want to connect (for example, GitHub Enterprise, GitLab, or Artifactory).
4. The CLI will generate your `DEPLOYMENT_ID`, `CLIENT_ID`, and `CLIENT_SECRET`. It will also prompt you to create a credential reference name (for example, `MY_GITHUB_TOKEN`) that maps to your actual secret. Keep these values secure for the next step.

{% hint style="info" %}
If you are setting up Snyk Essentials for asset management with over 1,000 repositories, Snyk recommends a dedicated Broker and Organization.
{% endhint %}
{% endstep %}

{% step %}
### Deploy the Broker instance

{% hint style="success" %}
**Key decision**: Decide on the deployment method (Docker or Kubernetes) that best fits your internal DevOps standards.
{% endhint %}

#### Docker deployment

Run the Docker command using the unified `snyk/broker:universal` image and your specific environment variables:

```
docker run --restart always \
  -p 8000:8000 \
  -e DEPLOYMENT_ID=<YOUR_DEPLOYMENT_ID> \
  -e CLIENT_ID=<YOUR_CLIENT_ID> \
  -e CLIENT_SECRET=<YOUR_CLIENT_SECRET> \
  -e PORT=8000 \
  -e <YOUR_CREDENTIALS_REFERENCE>=<secret_value> \
  snyk/broker:universal
```

#### Kubernetes deployment

Deploy using the official Snyk Universal Broker Helm chart. Ensure your secrets are stored securely as Kubernetes secrets.

```
helm pull oci://registry-1.docker.io/snyk/snyk-universal-broker helm install my-snyk-broker oci://registry-1.docker.io/snyk/snyk-universal-broker
--set deploymentId='YOUR_DEPLOYMENT_ID'
--set clientId='YOUR_CLIENT_ID'
--set clientSecret='YOUR_CLIENT_SECRET'
--set credentialReferences.<YOUR_CREDENTIALS_REFERENCE>='<secret_value>'

```
{% endstep %}

{% step %}
### Verify the connection

{% hint style="success" %}
**Key decision**: Determine if you need to configure additional environment variables for a proxy server (`HTTPS_PROXY`) or a custom certificate authority (`NODE_EXTRA_CA_CERTS`) to establish the connection out to Snyk.
{% endhint %}

1. If your connection is not fully mapped to an Organization, run `snyk-broker-config workflows connections integrate` in the CLI and select your deployment.
2. Check the Broker client container or pod logs to confirm the connection is established without any missing credentials reference errors.
3. Return to the Snyk web UI and navigate to your Organization **Settings** > **Integrations** page. Confirm the integration tile is marked as **Configured**.
4. Test the integration by importing a Project from your on-premise repository.

| Environment                       | Variable requirement                                                          |
| --------------------------------- | ----------------------------------------------------------------------------- |
| Standard                          | `DEPLOYMENT_ID`, `CLIENT_ID`, `CLIENT_SECRET`, `<YOUR_CREDENTIALS_REFERENCE>` |
| Proxy server                      | `HTTPS_PROXY`                                                                 |
| Custom CA (Certificate Authority) | `NODE_EXTRA_CA_CERTS`                                                         |
{% endstep %}
{% endstepper %}

## Enable Snyk Code

{% hint style="info" %}
Snyk Code is disabled for Organizations by default.
{% endhint %}

Enable Snyk Code to activate static application security testing (SAST) for your Organizations. Snyk Code analyzes your source code in real time to identify vulnerabilities and provides developer-friendly remediation advice.

{% stepper %}
{% step %}
### Verify Snyk Code availability

{% hint style="success" %}
**Key decision**: Determine if Snyk Code should be enabled globally for all Organizations or phased in for specific high-priority development teams.
{% endhint %}

Before enabling Snyk Code, ensure your Snyk license includes SAST capabilities. Snyk Code supports most major programming languages and integrates directly with your existing SCM and IDE setups.
{% endstep %}

{% step %}
### Enable Snyk Code in Settings

{% hint style="success" %}
**Key decision**: You must enable Snyk Code before importing your first Projects to ensure Snyk performs a code analysis scan during the initial onboarding.
{% endhint %}

Snyk Code is disabled by default in new Organizations. If you enable it after you have already imported a Project, Snyk does not automatically detect the code files. You must re-import the Project to trigger the scan.

1. In the Snyk web UI, navigate to **Settings** > **Snyk Code**.
2. Toggle the switch to **Enabled**.
3. Click **Save changes**.
{% endstep %}

{% step %}
### (Optional) Enable Snyk Code at scale using the API

{% hint style="success" %}
**Key decision**: If you are managing dozens or hundreds of Organizations, use the Snyk API to enable Snyk Code programmatically rather than using the web UI.
{% endhint %}

To enable Snyk Code for multiple Organizations:

* Use the `Enable/Disable the Snyk Code` settings API endpoint.
* Incorporate this call into your automated Organization provisioning script (refer to your Organization Template logic).
{% endstep %}

{% step %}
### Align with your import strategy

{% hint style="success" %}
**Key decision**: Choose the import method that provides the best resolution for your specific programming languages.
{% endhint %}

While SCM integration is the fastest way to gain visibility, some environments benefit from additional CLI scanning.

| Import strategy   | Benefit of Snyk Code (SAST)                                                      |
| ----------------- | -------------------------------------------------------------------------------- |
| SCM repository    | Provides continuous monitoring and centralized visibility without building code. |
| Snyk CLI          | Best for complex builds or local testing before a developer commits code.        |
| PR or MR scanning | Delivers immediate feedback on new code changes during the review process.       |
{% endstep %}
{% endstepper %}
