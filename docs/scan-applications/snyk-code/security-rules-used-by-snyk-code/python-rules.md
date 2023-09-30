# Python rules

## Rule (1) Incomplete URL sanitization

**CWE** (20) Improper Input Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (2) Arbitrary File Write via Archive Extraction (Tar Slip)

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (3) Path Traversal

**CWE** (23) Relative Path Traversal

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (4) Command Injection

**CWE** (78) Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (5) Cross-site Scripting (XSS)

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (6) Jinja auto-escape is set to false.

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (7) SQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (8) Code Injection

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (9) Improper Neutralization of Directives in Statically Saved Code

**CWE** (96) Improper Neutralization of Directives in Statically Saved Code ('Static Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (10) Server Information Exposure

**CWE** (209) Generation of Error Message Containing Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (11) Use of Hardcoded Credentials

**CWE** (259, 798) Use of Hard-coded Password, Use of Hard-coded Credentials

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (12) Improper Handling of Insufficient Permissions or Privileges

**CWE** (280) Improper Handling of Insufficient Permissions or Privileges

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (13) Binding to all network interfaces may open service to unintended traffic

**CWE** (284) Improper Access Control

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (14) Broken User Authentication

**CWE** (287) Improper Authentication

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (15) Improper Certificate Validation

**CWE** (295) Improper Certificate Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (16) Cryptographic Issues

**CWE** (310) Cryptographic Issues

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (17) Authentication over HTTP

**CWE** (319) Cleartext Transmission of Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (18) Use of Hardcoded Cryptographic Key

**CWE** (321) Use of Hard-coded Cryptographic Key

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (19) Inadequate Encryption Strength

**CWE** (326) Inadequate Encryption Strength

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (20) Use of a Broken or Risky Cryptographic Algorithm

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (21) Missing protocol in ssl.wrap\_socket

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (22) Use of Hardcoded Cryptographic Initialization Value

**CWE** (329) Generation of Predictable IV with CBC Mode

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (23) Origin Validation Error

**CWE** (346, 942) Origin Validation Error, Permissive Cross-domain Policy with Untrusted Domains

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (24) Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (25) Insecure Temporary File

**CWE** (377) Insecure Temporary File

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (26) Regular Expression Denial of Service (ReDoS)

**CWE** (400) Uncontrolled Resource Consumption

## Rule (27) Insecure default value

**CWE** (453) Insecure Default Variable Initialization

## Rule (28) Debug Mode Enabled

**CWE** (489) Active Debug Code

## Rule (29) Deserialization of Untrusted Data

**CWE** (502) Deserialization of Untrusted Data

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (30) Hardcoded Secret

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (31) Open Redirect

**CWE** (601) URL Redirection to Untrusted Site ('Open Redirect')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (32) Insecure Xml Parser

**CWE** (611) Improper Restriction of XML External Entity Reference

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (33) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**CWE** (614) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (34) XPath Injection

**CWE** (643) Improper Neutralization of Data within XPath Expressions ('XPath Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (35) Insecure File Permissions

**CWE** (732) Incorrect Permission Assignment for Critical Resource

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (36) Selection of Less-Secure Algorithm During Negotiation (SSL instead of TLS)

**CWE** (757) Selection of Less-Secure Algorithm During Negotiation ('Algorithm Downgrade')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (37) Use of Password Hash With Insufficient Computational Effort

**CWE** (916) Use of Password Hash With Insufficient Computational Effort

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (38) Server-Side Request Forgery (SSRF)

**CWE** (918) Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A10:2021 - Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (39) NoSQL Injection

**CWE** (943) Improper Neutralization of Special Elements in Data Query Logic

## Rule (40) Sensitive Cookie Without 'HttpOnly' Flag

**CWE** (1004) Sensitive Cookie Without 'HttpOnly' Flag

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (41) Python 2 source code

**CWE** (1104) Use of Unmaintained Third Party Components

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A06:2021 - Vulnerable and Outdated Components
