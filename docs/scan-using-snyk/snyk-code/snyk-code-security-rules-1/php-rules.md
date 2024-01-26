# PHP rules

## Rule (1) Arbitrary File Write via Archive Extraction (Zip Slip)

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (2) Path Traversal

**CWE** (23) Relative Path Traversal

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (3) Command Injection

**CWE** (78) Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (4) Cross-site Scripting (XSS)

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (5) SQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (6) Code Injection

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (7) File Inclusion

**CWE** (98) Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (8) Information Exposure

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (9) Use of Hardcoded Credentials

**CWE** (259, 798) Use of Hard-coded Password, Use of Hard-coded Credentials

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (10) Inadequate Padding for Public Key Encryption

**CWE** (326) Inadequate Encryption Strength

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (11) Inadequate Encryption Strength

**CWE** (326) Inadequate Encryption Strength

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (12) Use of a Broken or Risky Cryptographic Algorithm

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (13) Use of Insufficiently Random Values

**CWE** (330) Use of Insufficiently Random Values

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (14) Origin Validation Error

**CWE** (346, 942) Origin Validation Error, Permissive Cross-domain Policy with Untrusted Domains

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (15) Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (16) Regular Expression Denial of Service (ReDoS)

**CWE** (400) Uncontrolled Resource Consumption

## Rule (17) Deserialization of Untrusted Data

**CWE** (502) Deserialization of Untrusted Data

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (18) Privacy Leak

**CWE** (532) Insertion of Sensitive Information into Log File

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A09:2021 - Security Logging and Monitoring Failures

## Rule (19) Hardcoded Secret

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (20) Open Redirect

**CWE** (601) URL Redirection to Untrusted Site ('Open Redirect')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (21) XML External Entity (XXE) Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (22) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**CWE** (614) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (23) Weak Password Recovery Mechanism for Forgotten Password

**CWE** (640) Weak Password Recovery Mechanism for Forgotten Password

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (24) XPath Injection

**CWE** (643) Improper Neutralization of Data within XPath Expressions ('XPath Injection')

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (25) Allocation of Resources Without Limits or Throttling

**CWE** (770) Allocation of Resources Without Limits or Throttling

## Rule (26) Use of Password Hash With Insufficient Computational Effort

**CWE** (916) Use of Password Hash With Insufficient Computational Effort

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (27) Server-Side Request Forgery (SSRF)

**CWE** (918) Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A10:2021 - Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (28) Sensitive Cookie Without 'HttpOnly' Flag

**CWE** (1004) Sensitive Cookie Without 'HttpOnly' Flag

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (29) Improper Restriction of Rendered UI Layers or Frames

**CWE** (1021) Improper Restriction of Rendered UI Layers or Frames

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design
