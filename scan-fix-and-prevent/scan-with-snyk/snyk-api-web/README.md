# Snyk API & Web

Snyk API & Web is a cloud-based dynamic application security testing (DAST) solution that identifies security vulnerabilities in your running web applications and APIs. Snyk simulates real-world attacks against your deployed applications to discover security vulnerabilities before attackers can exploit them.

Modern applications expose complex attack surfaces through web interfaces and API endpoints. Security vulnerabilities in production applications can expose sensitive data, compromise user accounts, and damage the reputation of your company. Traditional static analysis tools examine code at rest, but many vulnerabilities only appear when code executes in its runtime environment.

Snyk tests your applications in their running state, finding runtime vulnerabilities such as SQL injection, cross-site scripting (XSS), authentication bypasses, and configuration weaknesses. By integrating DAST throughout the software development lifecycle (SDLC), you can catch and fix vulnerabilities before they reach production.

## Scan web applications and APIs

Snyk supports two scanning approaches:

* Web applications: the scanner crawls your application, discovers pages and functionality, interacts with forms and buttons, and performs comprehensive security tests across your entire web application.
* APIs: the scanner tests all endpoints defined in your API schema to ensure complete coverage of your API surface.

Configure authentication to scan protected areas accessible only to logged-in users. Snyk supports multiple authentication methods, including login forms, login sequences, Basic Auth, API keys, and two-factor authentication (2FA).

## Test behind authentication

Snyk provides flexible authentication options to scan protected application areas:

* Login forms and sequences: automate login flows using form selectors or recorded browser interactions.
* Two-factor authentication: TOTP-based 2FA (for example, Google Authenticator, Authy) and email or SMS-based OTP.
* API authentication: API keys, login endpoints, and token-based authentication.
* Logout detection: re-authenticate automatically if your application logs out mid-scan.

Comprehensive authentication support ensures security coverage across your entire application, including administrative interfaces and user-specific functionality.

## Scan internal and private applications

For applications not accessible from the public internet, deploy the scanning agent in your private network. The agent creates a secure tunnel between Snyk cloud scanning infrastructure and your internal applications to support DAST for:

* Development and staging environments on private networks
* Internal tools and administrative interfaces
* Applications behind VPNs or firewalls

The agent deploys as a Docker container or Kubernetes workload in your infrastructure.

## Integrate into your development workflow

Snyk provides multiple interfaces for integrating security testing into your workflow:

* Web UI: configure targets, review findings, and generate reports.
* REST API: programmatic access for custom integrations and automation.
* CLI: command-line interface for scripting and CI/CD pipelines.
* CI/CD integrations: native support for Jenkins, GitHub Actions, GitLab CI, and Azure DevOps.

Trigger scans automatically after deployments, fail builds based on vulnerability thresholds, and sync findings with Jira or other ticketing systems to streamline remediation across your development and security teams.

## Prioritize and fix vulnerabilities

Snyk detects and reports security vulnerabilities with detailed information, including:

* Severity ratings: Snyk assigns these automatically based on vulnerability type, exploitability, impact, and scope, with CVSS scores to prioritize fixes.
* Affected endpoints: specific URLs and API endpoints where Snyk finds vulnerabilities.
* Remediation guidance: detailed explanations and recommended fixes for each vulnerability type.
* Vulnerability evidence: evidence demonstrating how an attacker can exploit the vulnerability.

Use customizable scan profiles to balance speed and coverage based on your needs, from fast SSL/TLS checks for CI/CD gates to comprehensive full scans for thorough security assessments.
