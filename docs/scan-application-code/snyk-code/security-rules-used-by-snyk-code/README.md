# Security rules used by Snyk Code

{% hint style="info" %}
Snyk Code rules are updated continuously. The list expands continually, and the rules may change to provide the best protection and security solutions for your code.
{% endhint %}

This page lists the security rules used by Snyk Code when scanning your source code for vulnerabilities.

Each rule includes the following information.

* **Number and Rule Name**: Consecutive number for each rule and the Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **OWASP Top 10/SANS 25**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Supported Languages**: The programming languages to which this specific rule applies. Note that there might be two rules with the same name that apply to different languages.

## Rule (1) Use of Hardcoded Credentials

**CWE** (798) Use of Hard-coded Credentials

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** PHP

**CWE** (259) Use of Hard-coded Password

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Ruby, Go, Java, JavaScript, TypeScript, Python, C# & Asp.NET, Apex

## Rule (2) Use of Password Hash With Insufficient Computational Effort

**CWE** (916) Use of Password Hash With Insufficient Computational Effort

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Python, JavaScript, TypeScript, C# & ASP.NET. Java, Go, PHP, Apex

## Rule (3) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**CWE** (614) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** PHP, JavaScript, TypeScript, Ruby, C# & ASP,NET, Java, Python

## Rule (4) Hardcoded Secret

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** JavaScript, TypeScript, C# & ASP.NET, Java, Go, PHP, Python, Ruby, Apex

## Rule (5) Insecure Data Transmission

**CWE** (319) Cleartext Transmission of Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Ruby, C# & ASP.NET

## Rule (6) Command Injection

**CWE** (78) Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Python, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP, Apex

## Rule (7) Cross-site Scripting (XSS)

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP. Apex

## Rule (8) Server-Side Request Forgery (SSRF)

**CWE** (918) Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A10:2021 - Server-Side Request Forgery (SSRF)

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** JavaScript, TypeScript, C# & ASP.NET, Java, Go, PHP, Apex

## Rule (9) Open Redirect

**CWE** (601) URL Redirection to Untrusted Site ('Open Redirect')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Python, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP, Apex

## Rule (10) Regular expression injection

**CWE** (400) Uncontrolled Resource Consumption

**Supported languages:** Java

**CWE** (730)

**Supported languages:** C# & ASP.NET, Apex

## Rule (11) XML Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** C# & ASP.NET

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Apex

## Rule (12) SQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP, Apex

## Rule (13) Log Forging

**CWE** (117) Improper Output Neutralization for Logs

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A09:2021 - Security Logging and Monitoring Failures

**Supported languages:** C# & ASP.NET

## Rule (14) Use of Hardcoded Cryptographic Key

**CWE** (321) Use of Hard-coded Cryptographic Key

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Python, Ruby, Apex

## Rule (15) XML External Entity (XXE) Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** JavaScript, TypeScript

**OWASP Top 10/SANS 25:** SANS/CWE Top 25Ruby,&#x20;

**Supported languages:** Ruby, C# & ASP.NET, Java, PHP

## Rule (16) Inadequate Encryption Strength

**CWE** (326) Inadequate Encryption Strength

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** C# & ASP.NET, Java, Go, PHP

## Rule (17) Use of Insufficiently Random Values

**CWE** (330) Use of Insufficiently Random Values

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** PHP, Java, C# & ASP.NET, Go, JavaScript, TypeScript, Ruby

## Rule (18) Sensitive Cookie Without 'HttpOnly' Flag

**CWE** (1004) Sensitive Cookie Without 'HttpOnly' Flag

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** Python, Java, C# & ASP.NET, Go, JavaScript, TypeScript, PHP, Ruby

## Rule (19) Request Validation Disabled

**CWE** (554) ASP.NET Misconfiguration: Not Using Input Validation Framework

**Supported languages:** C# & ASP.NET

## Rule (20) IgnoreAntiforgeryToken in Use

**CWE** (352) Cross-Site Request Forgery (CSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** C# & ASP.NET

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (21) Debug Features Enabled

**CWE** (215) Insertion of Sensitive Information Into Debugging Code

**Supported languages:** C# & ASP.NET

## Rule (22) Deserialization of Untrusted Data

**CWE** (502) Deserialization of Untrusted Data

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, PHP

## Rule (23) ASP SSL Disabled

**CWE** (319) Cleartext Transmission of Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** C# & ASP.NET

## Rule (24) Code Injection

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Python, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, PHP

## Rule (25) Information Exposure

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** PHP

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** PHP, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java

## Rule (26) Exposure of Private Personal Information to an Unauthorized Actor

**CWE** (359) Exposure of Private Personal Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** C# & ASP.NET

## Rule (27) Cleartext Storage of Sensitive Information in a Cookie

**CWE** (315) Cleartext Storage of Sensitive Information in a Cookie

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** Java, C# & ASP.NET

## Rule (28) LDAP Injection

**CWE** (90) Improper Neutralization of Special Elements used in an LDAP Query ('LDAP Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java, C# & ASP.NET

## Rule (29) Path Traversal

**CWE** (23) Relative Path Traversal

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Python, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP

## Rule (30) XPath Injection

**CWE** (643) Improper Neutralization of Data within XPath Expressions ('XPath Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Python, JavaScript, TypeScript, Ruby, C# & ASP.NET, Java, Go, PHP

## Rule **(31) Arbitrary File Write via Archive Extraction (Zip Slip)**

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** PHP

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** JavaScript, TypeScript, C# & ASP.NET

## Rule (32) Improper Certificate Validation

**CWE** (295) Improper Certificate Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Go, Java, Ruby

## Rule (33) Insecure TLS Configuration

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Go

## Rule (34) Clear Text Logging

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Go

**CWE** (312) Cleartext Storage of Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule **(35) Generation of Error Message Containing Sensitive Information**

**CWE** (209) Generation of Error Message Containing Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** Go

## Rule (36) Improper Authentication

**CWE** (287) Improper Authentication

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (37) Use of a Broken or Risky Cryptographic Algorithm

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Java, JavaScript, TypeScript

## Rule **(38) Use of Potentially Dangerous Function**

**CWE** (676) Use of Potentially Dangerous Function

**Supported languages:** Java

## Rule (39) Use of Hardcoded, Security-relevant Constants

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** Java

## Rule (40) JWT Signature Verification Bypass

**CWE** (347) Improper Verification of Cryptographic Signature

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Java

## Rule (41) Cleartext Transmission of Sensitive Information

**CWE** (319) **Cleartext Transmission of Sensitive Information**

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** Java, JavaScript, TypeScript

## Rule (42) Improper Validation of Certificate with Host Mismatch

**CWE** (297) Improper Validation of Certificate with Host Mismatch

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java

## Rule (43) Insufficient Session Expiration

**CWE** (613) Insufficient Session Expiration

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java

## Rule (44) Origin Validation Error

**CWE** (942) Permissive Cross-domain Policy with Untrusted Domains

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** Python

**CWE** (346) Origin Validation Error

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java, JavaScript, TypeScript, PHP

## Rule (45) Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Python

**OWASP Top 10/SANS 25:**  SANS/CWE Top 25

**Supported languages:** Ruby, Java, JavaScript, TypeScript, PHP

## Rule (46) Disabled Neutralization of CRLF Sequences in HTTP Headers

**CWE** (113) Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Response Splitting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

## Rule (47) JavaScript Enabled

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (48) File Access Enabled

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Java

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (49) Android Fragment Injection

**CWE** (470) Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

## Rule (50) Spring Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Java

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (51) Struts Development Mode Enabled

**CWE** (489) Active Debug Code

**Supported languages:** Java

## Rule (52) Android Debug Mode Enabled

**CWE** (489) Active Debug Code

**Supported languages:** Java

## Rule **(53) Process Control**

**CWE** (114) Process Control

**Supported languages:** Java

## Rule (54) Use of Externally-Controlled Format String

**CWE** (134) Use of Externally-Controlled Format String

**Supported languages:** Java, JavaScript, TypeScript

## Rule (55) External Control of System or Configuration Setting

**CWE** (15) External Control of System or Configuration Setting

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** Java

## Rule **(56) Server Information Exposure**

**CWE** (209) Generation of Error Message Containing Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** Python, Java

## Rule **(57) Improper Neutralization of CRLF Sequences in HTTP Headers**

**CWE** (113) Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Response Splitting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

## Rule (58) Trust Boundary Violation

**CWE** (501) Trust Boundary Violation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** Java

## Rule (59) Android Intent Forwarding

**CWE** (940) Improper Verification of Source of a Communication Channel

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java

## Rule (60) Unauthorized File Access

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

**OWASP Top 10/SANS 25:**  SANS/CWE Top 25

## Rule (61) Code Execution via Third Party Package Context

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

## Rule (62) Android Uri Permission Manipulation

**CWE** (266) Incorrect Privilege Assignment

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** Java

## Rule (63) Java Naming and Directory Interface (JNDI) Injection

**CWE** (74) Improper Neutralization of Special Elements in Output Used by a Downstream Component ('Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Java

## Rule (64) Code Execution via Third Party Package Installation

**CWE** (940) Improper Verification of Source of a Communication Channel

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Java

## Rule (65) Observable Timing Discrepancy (Timing Attack)

**CWE** (208) Observable Timing Discrepancy

**Supported languages:** JavaScript, TypeScript

## Rule (66) Buffer Over-read

**CWE** (126) Buffer Over-read

**Supported languages:** JavaScript, TypeScript

## Rule (67) Improper Restriction of Rendered UI Layers or Frames

**CWE** (1021) Improper Restriction of Rendered UI Layers or Frames

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** PHP, JavaScript, TypeScript

## Rule (68) Unchecked Input for Loop Condition

**CWE** (400) Uncontrolled Resource Consumption

**OWASP Top 10/SANS 25:**&#x20;

**Supported languages:** JavaScript, TypeScript

**CWE** (606) Unchecked Input for Loop Condition

## Rule (69) Improper Input Validation

**CWE** (20) Improper Input Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** JavaScript, TypeScript

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Ruby

## Rule (70) Allocation of Resources Without Limits or Throttling

**CWE** (770) Allocation of Resources Without Limits or Throttling

**OWASP Top 10/SANS 25:**&#x20;

**Supported languages:** PHP, JavaScript, TypeScript

## Rule (71) Permissive Cross-domain Policy

**CWE** (942) Permissive Cross-domain Policy with Untrusted Domains

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:**  JavaScript, TypeScript

## Rule (72) Denial of Service (DoS) through Nested GraphQL Queries

**CWE** (400) Uncontrolled Resource Consumption

**Supported languages:** JavaScript, TypeScript

## Rule (73) Introspection Enabled

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** JavaScript, TypeScript

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (74) Weak Password Recovery Mechanism for Forgotten Password

**CWE** (640) Weak Password Recovery Mechanism for Forgotten Password

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** PHP, JavaScript, TypeScript

## Rule (75) Prototype Pollution

**CWE** (1321) Improperly Controlled Modification of Object Prototype Attributes ('Prototype Pollution')

**Supported languages:** JavaScript, TypeScript

## Rule (76) Regular Expression Denial of Service (ReDoS)

**CWE** (400) Uncontrolled Resource Consumption

**OWASP Top 10/SANS 25:**&#x20;

**Supported languages:** PHP. JavaScript, TypeScript, Python, Ruby

## Rule (77) Improper Neutralization of Directives in Statically Saved Code

**CWE** (96) Improper Neutralization of Directives in Statically Saved Code ('Static Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** JavaScript, TypeScript, Python, Ruby

## Rule (78) GraphQL Injection

**CWE**  (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** JavaScript, TypeScript

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Apex

## Rule (79) NoSQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** JavaScript, TypeScript

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Apex

## Rule (80) XML internal entity expansion

**CWE**  (776) Improper Restriction of Recursive Entity References in DTDs ('XML Entity Expansion')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**Supported languages:** JavaScript, TypeScript

## Rule (81) Inadequate Padding for Public Key Encryption

**CWE** (326) Inadequate Encryption Strength

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

**Supported languages:** PHP

## Rule (82) File Inclusion

**CWE**  (98) Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** PHP

## Rule (83) Broken User Authentication

**CWE** (287) Improper Authentication

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25&#x20;

## Rule (84) Insecure File Permissions

**CWE** (732) Incorrect Permission Assignment for Critical Resource

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

**Supported languages:** Python

## Rule (85) Improper Handling of Insufficient Permissions or Privileges

**CWE** (280) Improper Handling of Insufficient Permissions or Privileges

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**Supported languages:** Python

## Rule (86) Arbitrary File Write via Archive Extraction (Tar Slip)

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (87) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**CWE** (915) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**Supported languages:** Ruby

## Rule (88) Unsafe Reflection

**CWE** (470) Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**Supported languages:** Ruby
