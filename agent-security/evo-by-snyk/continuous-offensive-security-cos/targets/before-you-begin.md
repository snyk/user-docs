# Before you begin

Adding a target takes three steps. This page walks through all three in order, with full detail for each step inline.

{% hint style="warning" %}
**Test only what you are authorized to test**

Configuring a target arms Snyk Continuous Offensive Security to send real attack payloads at it. Confirm you have authorization to test the application, use test accounts rather than real user data, and prefer a staging environment where one exists.
{% endhint %}

## Prerequisites

Have the following ready:

* The main URL of the application you want to test, reachable from the internet.
* The list of other hosts your application depends on: your API, your identity provider, any CDN or external JavaScript source. You will need these to define the scope.
* Test credentials for at least one user. Two or more users with different roles produce a better assessment, because agents can then test for privilege escalation and broken authorization.
* If your network controls block automated traffic, the scanner IP addresses to allow through your WAF or firewall.

## Add a Target

{% stepper %}
{% step %}
## Name, URL, and scope

### Name

A display name for the target. This is what appears in the targets list, in the scans and jobs view, and in reports. Use something a colleague would recognize, for example `Acme Staging` rather than `target-1`.

### Main URL

The main entry point of your application. This is where agents begin their reconnaissance.

Your main URL is in scope automatically. You do not need to add it to the allowlist.

### Scope

The scope is the set of hosts and paths agents are allowed to reach. Every request an agent makes passes through a proxy, and anything outside the scope is rejected before it leaves.

Add every host that is part of your application, including:

* Your API, if it is served from a different hostname.
* Your third-party login provider, such as Okta or GitHub.
* Any CDN you serve assets from.
* Any external endpoint you fetch JavaScript or other resources from.

You can also add a reject list of paths that must never be touched. If something appears in both the allowlist and the reject list, the reject list wins.

Scope is the setting most worth getting right on the first pass, because anything you leave out is silently never tested. Visit Define the target scope for the full treatment, including worked examples.

{% hint style="info" %}
Your main URL is in scope automatically, but anything else—including your own API on a different hostname—is rejected by the proxy and never tested if you leave it out.
{% endhint %}
{% endstep %}

{% step %}
## Authentication

Specify the credentials of the users you want agents to test as. You can add any number of users.

{% hint style="info" %}
Add more than one user, with different roles. Agents compare what each user can reach against what that user should be able to reach, which is how they detect privilege escalation and broken authorization such as BOLA. With a single user there is nothing to compare against, and this entire class of vulnerability goes untested.
{% endhint %}

Label each user with the role it represents, then choose one of two authentication types:

| Type                  | When to use it                                                                                                                                                                                                                                           |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Credentials login** | The login is a form. Provide any number of key-value pairs: username or email, password, and whatever other fields the form needs. Supports a TOTP secret for two-factor authentication, and free-text hints for small obstacles such as a cookie banner |
| **Custom login**      | The login flow cannot be expressed as form fields, for example single sign-on or a multi-page login. Describe the flow in natural language and the agents follow your description                                                                        |

You can also give the URL of your login page. If you omit it, it is detected automatically.

Visit Configure authentication for both types in detail.
{% endstep %}

{% step %}
## Extra instructions

The third step is optional but improves assessment quality. It has three parts.

### Custom headers

Provide headers that agents send with every request.

#### Get scan traffic past a firewall or WAF

The most common use. A web application firewall that sees a burst of attack payloads will usually start blocking, which means agents spend the scan being rate limited rather than testing your application. Configure a header your firewall recognizes and allows through, and scan traffic reaches the application intact.

Coordinate this with whoever operates the firewall. The usual arrangement is a shared secret header that the firewall is configured to allow.

The alternative is to allow the scanner IP addresses at the firewall instead. Either approach works: a header is easier when you cannot change network policy, an IP allowlist is easier when you cannot inject headers. Visit Scanner IP addresses.

{% hint style="info" %}
Bypassing your firewall for a scan is deliberate. You are testing the application, not the firewall in front of it. A vulnerability that a WAF happens to mask today is still a vulnerability, and WAF rules change.
{% endhint %}

#### Other header uses

| Use                                | Example                                                                 |
| ---------------------------------- | ----------------------------------------------------------------------- |
| Route to a specific environment    | A header your load balancer or service mesh uses to select a deployment |
| Satisfy a required API header      | An API version header, or a tenant identifier your gateway expects      |
| Identify scan traffic in your logs | A header you can filter on to separate scan requests from real users    |

### Custom cookies

Provide cookies that agents send with every request.

| Use                            | Notes                                                                                                   |
| ------------------------------ | ------------------------------------------------------------------------------------------------------- |
| Carry a session token directly | When your session cannot be created by submitting the login form, supply a valid session cookie instead |
| Set required application state | Feature flags, locale, tenant selection, or consent cookies that the application expects to be present  |
| Skip a blocking interstitial   | A cookie that suppresses an onboarding tour or cookie banner which would otherwise obstruct crawling    |

Session tokens expire. If you supply one directly rather than configuring authentication, it may expire partway through a long scan and the authenticated surface stops being tested from that point on. Configuring credentials is more durable, because agents can re-establish the session themselves. Use a supplied cookie when credentials are not an option, not as the default.

### Application context

Free-text context about the application. This is the part most often skipped and the part that most improves the relevance of results.

Agents use this context to decide where to concentrate effort and how seriously to treat what they find. The same technical vulnerability can be a minor issue in one application and a critical one in another, and the difference is usually business context that cannot be inferred from the outside.

#### What to include

| Include                                    | Why it helps                                                                             |
| ------------------------------------------ | ---------------------------------------------------------------------------------------- |
| **What the application does**              | Orients the agents. A payments platform and an internal wiki warrant different attention |
| **Who the users are and what roles exist** | Lets agents reason about which cross-role actions should be forbidden                    |
| **Where sensitive data lives**             | Focuses effort on the paths where a breach would actually matter                         |
| **The parts you care most about**          | Concentrates effort on the surface you are worried about                                 |
| **Known quirks**                           | Behavior that looks broken but is intentional, so agents do not spend the scan on it     |
| **Tenancy model**                          | If the application is multi-tenant, say so, and describe how tenants are separated       |

#### What not to include

* Credentials. Those belong in authentication.
* Instructions to avoid a host or path. Those belong in the reject list, which is enforced at the proxy. See Define the target scope.

{% hint style="warning" %}
Context guides the agents; it does not constrain them. Only the reject list is an enforced boundary. If a path must never be touched, put it in the reject list rather than asking for it in the context field.
{% endhint %}
{% endstep %}
{% endstepper %}

## Link a source repository

Linking a source repository upgrades the assessment from blackbox to greybox. Agents gain source code context, which lets them reach logic flaws that are invisible from the outside, and remediation guidance changes from generic advice to a specific file, line, and suggested code change.

There is no separate setting to enable greybox. It is selected automatically whenever a repository is linked.

## After you save

The target appears in your targets list, ready to scan. Nothing is tested until you start a scan.&#x20;

Some configuration problems are caught immediately: a scan will not start at all if the target is unreachable or if authentication fails.

Others only show up once the scan is underway. Because a scan takes a few hours, check the recon numbers early rather than waiting for the result. If recon discovers only your login page, or far less of the application than you expect, agents are most likely not getting past authentication or the rest of your application is out of scope. Cancel, fix the target, and start again. See [Monitor and manage scans](../scans/monitor-and-manage-scans.md).

