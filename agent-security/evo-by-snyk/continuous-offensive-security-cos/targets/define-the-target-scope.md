# Define the Target Scope

The scope defines how far agents may go. It is the boundary between the application you authorized for testing and everything else on the internet.

## How scope is enforced

Every request an agent makes passes through a proxy. The proxy checks the request against your scope and rejects anything that falls outside it, before the request leaves.

```
Agent  →  Proxy  →  in scope?  ──yes──→  your application
                        │
                        └──no───→  rejected
```

This matters for two reasons:

* The boundary is not advisory. Scope is enforced by the proxy, not by instructing the agents to stay inside it. An agent that follows a link to a system you did not authorize is stopped at the proxy.
* What you leave out is never tested. There is no warning and no error. Requests to an out-of-scope host are simply rejected, so that part of your application silently produces no findings.

## What agents do inside the scope

Scope is a permission boundary, not a to-do list. Within the scope you define, agents examine the application, work out which parts genuinely belong to it, and concentrate their attacks there. Putting a CDN in scope does not mean the CDN gets pentested; it means agents can load the assets they need in order to render and understand your application properly.

## The allowlist

Your main URL is in scope automatically. Everything else must be listed.

Work through this checklist for your application:

| Add to scope                                       | Why                                                                                                                                                                           |
| -------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Your API**, if served from a different hostname  | The front end calls it, and it is usually where the highest severity findings are. Leaving it out is the single most common scoping mistake                                   |
| **Your identity provider**, such as Okta or GitHub | If the login flow redirects off your main domain and the provider is out of scope, the redirect is rejected and authentication fails, so agents never get past the login page |
| **Your CDN**                                       | Agents cannot render pages correctly without the assets they reference                                                                                                        |
| **External JavaScript sources**                    | Application behavior defined in externally hosted scripts is invisible if those scripts cannot load                                                                           |
| **Any other host in the same application**         | Separate hostnames for admin interfaces, file uploads, websockets, or authentication services                                                                                 |

### Worked example

For an application at `https://app.example.com` that serves its API from `api.example.com`, authenticates through Okta, and loads assets from a CDN:

| Setting  | Value                                                    |
| -------- | -------------------------------------------------------- |
| Main URL | `https://app.example.com`                                |
| Scope    | `api.example.com`, `example.okta.com`, `cdn.example.com` |

With `api.example.com` omitted, agents would crawl the front end, watch its API calls get rejected at the proxy, and report almost nothing. With `example.okta.com` omitted, authentication would fail outright and the entire authenticated surface would go untested.

## The reject list

The reject list names hosts and paths agents must never touch. Like the allowlist, it is enforced at the proxy.

The reject list always wins. If something matches both the allowlist and the reject list, it is rejected. You can therefore put a broad host in the allowlist and carve specific paths out of it, without having to enumerate every permitted path.

Use the reject list for anything where the act of testing is itself the problem:

| Reject                                                         | Why                                                                |
| -------------------------------------------------------------- | ------------------------------------------------------------------ |
| Payment and checkout endpoints                                 | Attacks may trigger real charges                                   |
| Outbound email and notification endpoints                      | Attacks may send email to real recipients                          |
| Destructive operations, such as account deletion or data purge | The action cannot be undone                                        |
| Third-party integrations you do not own                        | You are not authorized to test them                                |
| Administrative operations with real-world side effects         | Deployments, provisioning, or anything that changes infrastructure |

{% hint style="warning" %}
The reject list is a scope control, not a safety guarantee for production. Snyk Continuous Offensive Security sends real attack payloads to everything inside your scope. Prefer a staging environment, and where you must test production, be deliberate about the reject list.
{% endhint %}

## Getting scope right

Start broader than feels comfortable, then narrow. An overly tight scope produces a clean-looking report that reflects untested surface rather than a secure application. Missing findings are more dangerous than noisy ones.

Diagnose scope problems from recon. If the recon stage of your first scan discovers far fewer pages and endpoints than you expect, scope is the first thing to check. See Monitor and manage scans.

Revisit scope when the application changes. A new service on a new hostname is out of scope until you add it.

