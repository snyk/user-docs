---
nav_context: agnostic
---

# Supported rules

Select your language to see the rules that support custom sanitizers.

{% tabs %}
{% tab title="Apex" %}
| Rule Key              | Rule name                                        |
| --------------------- | ------------------------------------------------ |
| CommandInjection      | Command Injection                                |
| EmailContentInjection | Improper Access Control: Email Content Injection |
| OR                    | Open Redirect                                    |
| PotentialXSS          | Cross-site Scripting (XSS)                       |
| RegexInjection        | Regular expression injection                     |
| Soqli                 | SOQL Injection                                   |
| Sosli                 | SOSL Injection                                   |
| Ssrf                  | Server-Side Request Forgery (SSRF)               |
| XmlInjection          | XML Injection                                    |
| XSS                   | Cross-site Scripting (XSS)                       |
{% endtab %}

{% tab title="C/C++" %}
| Rule Key                   | Rule name                                                         |
| -------------------------- | ----------------------------------------------------------------- |
| ClientSideRequestForgery   | Client-Side Request Forgery                                       |
| CodeInjection              | Code Injection                                                    |
| CommandInjection           | Command Injection                                                 |
| FormatString               | Use of Externally-Controlled Format String                        |
| ImproperNullTermination    | Improper Null Termination                                         |
| IntegerOverflow            | Integer Overflow                                                  |
| LdapInjection              | LDAP Injection                                                    |
| PrivateInformationExposure | Exposure of Private Personal Information to an Unauthorized Actor |
| PT                         | Path Traversal                                                    |
| Sqli                       | SQL Injection                                                     |
| Ssrf                       | Server-Side Request Forgery (SSRF)                                |
| UserControlledPointer      | User Controlled Pointer                                           |
| Xpath                      | XPath Injection                                                   |
| XSS                        | Cross-site Scripting (XSS)                                        |
| XXE                        | XML External Entity (XXE) Injection                               |
{% endtab %}

{% tab title="C#" %}
| Rule Key                   | Rule name                                                         |
| -------------------------- | ----------------------------------------------------------------- |
| CleartextCookieStorage     | Cleartext Storage of Sensitive Information in a Cookie            |
| CodeInjection              | Code Injection                                                    |
| CommandInjection           | Command Injection                                                 |
| HttpResponseSplitting      | Improper Neutralization of CRLF Sequences in HTTP Headers         |
| LdapInjection              | LDAP Injection                                                    |
| LogForging                 | Log Forging                                                       |
| OR                         | Open Redirect                                                     |
| PrivateInformationExposure | Exposure of Private Personal Information to an Unauthorized Actor |
| PT                         | Path Traversal                                                    |
| RegexInjection             | Regular expression injection                                      |
| ServerLeak                 | Information Exposure - Server Error Message                       |
| Sqli                       | SQL Injection                                                     |
| Ssrf                       | Server-Side Request Forgery (SSRF)                                |
| XamlInjection              | XAML Injection                                                    |
| XmlInjection               | XML Injection                                                     |
| Xpath                      | XPath Injection                                                   |
| XSS                        | Cross-site Scripting (XSS)                                        |
| XXE                        | XML External Entity (XXE) Injection                               |
| ZipSlip                    | Arbitrary File Write via Archive Extraction (Zip Slip)            |
{% endtab %}

{% tab title="Dart" %}
| Rule Key | Rule name                                              |
| -------- | ------------------------------------------------------ |
| OR       | Open Redirect                                          |
| PT       | Path Traversal                                         |
| ZipSlip  | Arbitrary File Write via Archive Extraction (Zip Slip) |
{% endtab %}

{% tab title="Go" %}
| Rule Key              | Rule name                                                      |
| --------------------- | -------------------------------------------------------------- |
| ClearTextLogging      | Clear Text Logging                                             |
| CommandInjection      | Command Injection                                              |
| EmailContentInjection | Improper Access Control: Email Content Injection               |
| ErrorMessage          | Generation of Error Message Containing Sensitive Information   |
| OR                    | Open Redirect                                                  |
| PT                    | Path Traversal                                                 |
| Sqli                  | SQL Injection                                                  |
| Ssrf                  | Server-Side Request Forgery (SSRF)                             |
| Ssti                  | Improper Neutralization of Directives in Statically Saved Code |
| Xpath                 | XPath Injection                                                |
| XSS                   | Cross-site Scripting (XSS)                                     |
{% endtab %}

{% tab title="Groovy" %}
| Rule Key                 | Rule name                          |
| ------------------------ | ---------------------------------- |
| CodeInjection            | Code Injection                     |
| CommandInjection         | Command Injection                  |
| IndirectCommandInjection | Command Injection                  |
| OR                       | Open Redirect                      |
| PT                       | Path Traversal                     |
| Sqli                     | SQL Injection                      |
| Ssrf                     | Server-Side Request Forgery (SSRF) |
| XSS                      | Cross-site Scripting (XSS)         |
{% endtab %}

{% tab title="Java" %}
| Rule Key                   | Rule name                                                  |
| -------------------------- | ---------------------------------------------------------- |
| CodeInjection              | Code Injection                                             |
| CommandInjection           | Command Injection                                          |
| EnvCommandInjection        | Indirect Command Injection via User Controlled Environment |
| ExternalConfigControl      | External Control of System or Configuration Setting        |
| ExternalProcessControl     | Process Control                                            |
| HttpResponseSplitting      | Improper Neutralization of CRLF Sequences in HTTP Headers  |
| IndirectCommandInjection   | Command Injection                                          |
| IntentForwarding           | Android Intent Forwarding                                  |
| JndiInjection              | Java Naming and Directory Interface (JNDI) Injection       |
| LdapInjection              | LDAP Injection                                             |
| NoSqli                     | NoSQL Injection                                            |
| OR                         | Open Redirect                                              |
| PT                         | Path Traversal                                             |
| Reflection                 | Unsafe Reflection                                          |
| RegexInjection             | Regular expression injection                               |
| Sqli                       | SQL Injection                                              |
| Ssrf                       | Server-Side Request Forgery (SSRF)                         |
| ThirdPartyCodeExec         | Code Execution via Third Party Package Context             |
| ThirdPartyInstall          | Code Execution via Third Party Package Installation        |
| TooPermissiveCorsHeader    | Origin Validation Error                                    |
| TrustBoundaryViolation     | Trust Boundary Violation                                   |
| UnauthorizedFileAccess     | Unauthorized File Access                                   |
| UriPermissionManipulation  | Android Uri Permission Manipulation                        |
| UserControlledFormatString | Use of Externally-Controlled Format String                 |
| Xpath                      | XPath Injection                                            |
| XSS                        | Cross-site Scripting (XSS)                                 |
{% endtab %}

{% tab title="JavaScript" %}
| Rule Key                 | Rule name                                                      |
| ------------------------ | -------------------------------------------------------------- |
| CodeInjection            | Code Injection                                                 |
| CommandInjection         | Command Injection                                              |
| Deserialization          | Deserialization of Untrusted Data                              |
| DOMXSS                   | DOM-based Cross-site Scripting (XSS)                           |
| FormatString             | Use of Externally-Controlled Format String                     |
| Graphqli                 | GraphQL Injection                                              |
| ImproperCodeSanitization | Improper Code Sanitization                                     |
| IndirectCommandInjection | Command Injection                                              |
| NoSqli                   | NoSQL Injection                                                |
| OR                       | Open Redirect                                                  |
| PrototypePollution       | Prototype Pollution                                            |
| PT                       | Path Traversal                                                 |
| reDOS                    | Regular Expression Denial of Service (ReDoS)                   |
| reDOSPolynomial          | Regular Expression Denial of Service (ReDoS)                   |
| Spoof                    | Weak Password Recovery Mechanism for Forgotten Password        |
| Sqli                     | SQL Injection                                                  |
| Ssrf                     | Server-Side Request Forgery (SSRF)                             |
| Ssti                     | Improper Neutralization of Directives in Statically Saved Code |
| UnsafeJqueryPlugin       | Unsafe JQuery Plugin                                           |
| Xpath                    | XPath Injection                                                |
| XSS                      | Cross-site Scripting (XSS)                                     |
| XXE                      | XML External Entity (XXE) Injection                            |
| ZipSlip                  | Arbitrary File Write via Archive Extraction (Zip Slip)         |
{% endtab %}

{% tab title="Kotlin" %}
| Rule Key                   | Rule name                                                  |
| -------------------------- | ---------------------------------------------------------- |
| CodeInjection              | Code Injection                                             |
| CommandInjection           | Command Injection                                          |
| Deserialization            | Deserialization of Untrusted Data                          |
| EnvCommandInjection        | Indirect Command Injection via User Controlled Environment |
| ExternalConfigControl      | External Control of System or Configuration Setting        |
| ExternalProcessControl     | Process Control                                            |
| HttpResponseSplitting      | Improper Neutralization of CRLF Sequences in HTTP Headers  |
| IndirectCommandInjection   | Command Injection                                          |
| IntentForwarding           | Android Intent Forwarding                                  |
| JndiInjection              | Java Naming and Directory Interface (JNDI) Injection       |
| LdapInjection              | LDAP Injection                                             |
| OR                         | Open Redirect                                              |
| PT                         | Path Traversal                                             |
| RegexInjection             | Regular expression injection                               |
| Sqli                       | SQL Injection                                              |
| Ssrf                       | Server-Side Request Forgery (SSRF)                         |
| ThirdPartyCodeExec         | Code Execution via Third Party Package Context             |
| ThirdPartyInstall          | Code Execution via Third Party Package Installation        |
| TooPermissiveCorsHeader    | Origin Validation Error                                    |
| TrustBoundaryViolation     | Trust Boundary Violation                                   |
| UnauthorizedFileAccess     | Unauthorized File Access                                   |
| UriPermissionManipulation  | Android Uri Permission Manipulation                        |
| UserControlledFormatString | Use of Externally-Controlled Format String                 |
| Xpath                      | XPath Injection                                            |
{% endtab %}

{% tab title="PHP" %}
| Rule Key              | Rule name                                              |
| --------------------- | ------------------------------------------------------ |
| CodeInjection         | Code Injection                                         |
| CommandInjection      | Command Injection                                      |
| Deserialization       | Deserialization of Untrusted Data                      |
| EmailContentInjection | Improper Access Control: Email Content Injection       |
| FileInclusion         | File Inclusion                                         |
| OR                    | Open Redirect                                          |
| PT                    | Path Traversal                                         |
| reDOS                 | Regular Expression Denial of Service (ReDoS)           |
| ServerLeak            | Information Exposure - Server Error Message            |
| Sqli                  | SQL Injection                                          |
| Ssrf                  | Server-Side Request Forgery (SSRF)                     |
| Xpath                 | XPath Injection                                        |
| XSS                   | Cross-site Scripting (XSS)                             |
| XxePostPHP8           | XML External Entity (XXE) Injection - Post-PHP 8       |
| XxePrePHP8            | XML External Entity (XXE) Injection - Pre-PHP 8        |
| ZipSlip               | Arbitrary File Write via Archive Extraction (Zip Slip) |
{% endtab %}

{% tab title="Python" %}
| Rule Key                  | Rule name                                                      |
| ------------------------- | -------------------------------------------------------------- |
| BrokenUserAuthentication  | Broken User Authentication                                     |
| CodeInjection             | Code Injection                                                 |
| CommandInjection          | Command Injection                                              |
| Deserialization           | Deserialization of Untrusted Data                              |
| DjangoUnvalidatedPassword | Password Requirements Not Enforced in Django Application       |
| LdapInjection             | LDAP Injection                                                 |
| NoSqli                    | NoSQL Injection                                                |
| OR                        | Open Redirect                                                  |
| PT                        | Path Traversal                                                 |
| reDOS                     | Regular Expression Denial of Service (ReDoS)                   |
| ServerInformationExposure | Server Information Exposure                                    |
| Sqli                      | SQL Injection                                                  |
| Ssrf                      | Server-Side Request Forgery (SSRF)                             |
| Ssti                      | Improper Neutralization of Directives in Statically Saved Code |
| Xpath                     | XPath Injection                                                |
| XSS                       | Cross-site Scripting (XSS)                                     |
| XXE                       | XML External Entity (XXE) Injection                            |
{% endtab %}

{% tab title="Ruby" %}
| Rule Key            | Rule name                                                      |
| ------------------- | -------------------------------------------------------------- |
| CommandInjection    | Command Injection                                              |
| Deserialization     | Deserialization of Untrusted Data                              |
| OR                  | Open Redirect                                                  |
| PT                  | Path Traversal                                                 |
| reDOS               | Regular Expression Denial of Service (ReDoS)                   |
| Reflection          | Unsafe Reflection                                              |
| SessionManipulation | Session Manipulation                                           |
| Sqli                | SQL Injection                                                  |
| Ssti                | Improper Neutralization of Directives in Statically Saved Code |
| Xpath               | XPath Injection                                                |
| XXE                 | XML External Entity (XXE) Injection                            |
{% endtab %}

{% tab title="Rust" %}
| Rule Key         | Rule name                          |
| ---------------- | ---------------------------------- |
| CommandInjection | Command Injection                  |
| OR               | Open Redirect                      |
| PT               | Path Traversal                     |
| Sqli             | SQL Injection                      |
| Ssrf             | Server-Side Request Forgery (SSRF) |
| XSS              | Cross-site Scripting (XSS)         |
{% endtab %}

{% tab title="Scala" %}
| Rule Key                   | Rule name                                                  |
| -------------------------- | ---------------------------------------------------------- |
| CodeInjection              | Code Injection                                             |
| CommandInjection           | Command Injection                                          |
| Deserialization            | Deserialization of Untrusted Data                          |
| EnvCommandInjection        | Indirect Command Injection via User Controlled Environment |
| ExternalConfigControl      | External Control of System or Configuration Setting        |
| ExternalProcessControl     | Process Control                                            |
| HttpResponseSplitting      | Improper Neutralization of CRLF Sequences in HTTP Headers  |
| IndirectCommandInjection   | Command Injection                                          |
| JndiInjection              | Java Naming and Directory Interface (JNDI) Injection       |
| LdapInjection              | LDAP Injection                                             |
| OR                         | Open Redirect                                              |
| PT                         | Path Traversal                                             |
| RegexInjection             | Regular expression injection                               |
| Sqli                       | SQL Injection                                              |
| Ssrf                       | Server-Side Request Forgery (SSRF)                         |
| TooPermissiveCorsHeader    | Origin Validation Error                                    |
| TrustBoundaryViolation     | Trust Boundary Violation                                   |
| UserControlledFormatString | Use of Externally-Controlled Format String                 |
| Xpath                      | XPath Injection                                            |
| XSS                        | Cross-site Scripting (XSS)                                 |
{% endtab %}

{% tab title="Swift" %}
| Rule Key         | Rule name                           |
| ---------------- | ----------------------------------- |
| CodeInjection    | Code Injection                      |
| CommandInjection | Command Injection                   |
| MemoryCorruption | Memory Corruption                   |
| PT               | Path Traversal                      |
| Sqli             | SQL Injection                       |
| Ssrf             | Server-Side Request Forgery (SSRF)  |
| XSS              | Cross-site Scripting (XSS)          |
| XXE              | XML External Entity (XXE) Injection |
{% endtab %}

{% tab title="VB.NET" %}
| Rule Key              | Rule name                                                 |
| --------------------- | --------------------------------------------------------- |
| CodeInjection         | Code Injection                                            |
| CommandInjection      | Command Injection                                         |
| HttpResponseSplitting | Improper Neutralization of CRLF Sequences in HTTP Headers |
| OR                    | Open Redirect                                             |
| PT                    | Path Traversal                                            |
| RegexInjection        | Regular expression injection                              |
| Sqli                  | SQL Injection                                             |
| Ssrf                  | Server-Side Request Forgery (SSRF)                        |
| XmlInjection          | XML Injection                                             |
| Xpath                 | XPath Injection                                           |
| XSS                   | Cross-site Scripting (XSS)                                |
| XXE                   | XML External Entity (XXE) Injection                       |
{% endtab %}
{% endtabs %}
