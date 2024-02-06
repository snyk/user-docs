# Ruby rules

## Rule (1) Configuration Issue: Sinatra Protection Layers Disabled

**CWE** (16, 35, 79, 348, 352, 693, 1021) Configuration, Path Traversal: '.../...//', Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting'), Use of Less Trusted Source, Cross-Site Request Forgery (CSRF), Protection Mechanism Failure, Improper Restriction of Rendered UI Layers or Frames

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (2) Improper Input Validation

**CWE** (20) Improper Input Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

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

## Rule (6) SQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (7) Code Injection

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (8) Remote Code Execution via Endpoint

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (9) Improper Neutralization of Directives in Statically Saved Code

**CWE** (96) Improper Neutralization of Directives in Statically Saved Code ('Static Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (10) Information Exposure

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (11) Use of Hardcoded Credentials

**CWE** (259, 798) Use of Hard-coded Password, Use of Hard-coded Credentials

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (12) Session Manipulation

**CWE** (285) Improper Authorization

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (13) Improper Certificate Validation

**CWE** (295) Improper Certificate Validation

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (14) Selection of Less-Secure Algorithm During Negotiation (Force SSL)

**CWE** (311, 757) Missing Encryption of Sensitive Data, Selection of Less-Secure Algorithm During Negotiation ('Algorithm Downgrade')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (15) Insecure Data Transmission

**CWE** (319) Cleartext Transmission of Sensitive Information

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (16) Use of Hardcoded Cryptographic Key

**CWE** (321) Use of Hard-coded Cryptographic Key

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (17) Use of a Broken or Risky Cryptographic Algorithm

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (18) Use of Insufficiently Random Values

**CWE** (330) Use of Insufficiently Random Values

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (19) Regular Expression Denial of Service (ReDoS)

**CWE** (400) Uncontrolled Resource Consumption

## Rule (20) Unsafe Reflection

**CWE** (470) Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (21) Deserialization of Untrusted Data

**CWE** (502) Deserialization of Untrusted Data

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (22) No Weak Password Requirements

**CWE** (521) Weak Password Requirements

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (23) Hardcoded Secret

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (24) Open Redirect

**CWE** (601) URL Redirection to Untrusted Site ('Open Redirect')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (25) XML External Entity (XXE) Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (26) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**CWE** (614) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (27) XPath Injection

**CWE** (643) Improper Neutralization of Data within XPath Expressions ('XPath Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (28) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**CWE** (915) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

## Rule (29) Use of Password Hash With Insufficient Computational Effort

**CWE** (916) Use of Password Hash With Insufficient Computational Effort

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (30) Sensitive Cookie Without 'HttpOnly' Flag

**CWE** (1004) Sensitive Cookie Without 'HttpOnly' Flag

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (31) Incorrect regular expression for validating values

**CWE** (1286) Improper Validation of Syntactic Correctness of Input
