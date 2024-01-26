# Snyk Code security rules

{% hint style="info" %}
Snyk Code rules are updated continuously. The list expands continually, and the rules may change to provide the best protection and security solutions for your code.
{% endhint %}

{% hint style="warning" %}
If you have followed a link for code quality from an IDE, see the language documentation for that information.
{% endhint %}

This page lists the security rules used by Snyk Code when scanning your source code for vulnerabilities.

Each rule includes the following information.

* **Number and Rule Name**: Consecutive number for each rule and the Snyk name of the rule.
* **CWE(s):** The [CWE numbers](https://cwe.mitre.org/) that are covered by this rule.
* **OWASP Top 10/SANS 25**: The [OWASP Top 10 ](https://owasp.org/Top10/)(2021 edition) category to which the rule belongs to, if any, and if it is included in [SANS 25](https://www.sans.org/top25-software-errors/).
* **Supported Languages**: The programming languages to which this specific rule applies. Note that there might be two rules with the same name that apply to different languages.

## Rule (1) External Control of System or Configuration Setting

**CWE** (15) External Control of System or Configuration Setting

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (2) Configuration Issue: Electron Disable Security Warnings

**CWE** (16) Configuration

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (3) Configuration Issues: Electron Insecure Web Preferences

**CWE** (16) Configuration

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (4) Configuration Issues: Electron Load Insecure Content

**CWE** (16) Configuration

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (5) Configuration Issue: Sinatra Protection Layers Disabled

**CWE** (16, 35, 79, 348, 352, 693, 1021) Configuration, Path Traversal: '.../...//', Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting'), Use of Less Trusted Source, Cross-Site Request Forgery (CSRF), Protection Mechanism Failure, Improper Restriction of Rendered UI Layers or Frames

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (6) Insufficient postMessage Validation

**CWE** (20) Improper Input Validation

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (7) Improper Input Validation

**CWE** (20) Improper Input Validation

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (8) Incomplete URL sanitization

**CWE** (20) Improper Input Validation

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (9) Arbitrary File Write via Archive Extraction (Tar Slip)

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (10) Arbitrary File Write via Archive Extraction (Zip Slip)

**CWE** (22) Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')

**Supported languages:** C# and ASP.NET, JavaScript and TypeScript, PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (11) Path Traversal

**CWE** (23) Relative Path Traversal

**Supported languages:** C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (12) Java Naming and Directory Interface (JNDI) Injection

**CWE** (74) Improper Neutralization of Special Elements in Output Used by a Downstream Component ('Injection')

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (13) Command Injection

**CWE** (78) Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')

**Supported languages:** Apex, C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (14) Indirect Command Injection via User Controlled Environment

**CWE** (78) Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (15) Disabling Strict Contextual escaping (SCE) could provide additional attack surface for Cross-site Scripting (XSS)

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (16) Cross-site Scripting (XSS)

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** Apex, C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (17) JavaScript Enabled

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (18) Jinja auto-escape is set to false.

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (19) Use dangerouslySetInnerHTML to be explicit that this function is dangerous and also trigger react updates

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (20) Unauthorized File Access

**CWE** (79) Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (21) GraphQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (22) SOQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (23) SOSL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (24) SQL Injection

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (25) Unsafe SOQL Concatenation

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (26) Unsafe SOSL Concatenation

**CWE** (89) Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (27) LDAP Injection

**CWE** (90) Improper Neutralization of Special Elements used in an LDAP Query ('LDAP Injection')

**Supported languages:** C++, C# and ASP.NET, Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (28) XML Injection

**CWE** (91) XML Injection (aka Blind XPath Injection)

**Supported languages:** Apex, C# and ASP.NET, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (29) Code Injection

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**Supported languages:** C# and ASP.NET, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (30) Remote Code Execution via Endpoint

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (31) Code Execution via Third Party Package Context

**CWE** (94) Improper Control of Generation of Code ('Code Injection')

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (32) Improper Neutralization of Directives in Statically Saved Code

**CWE** (96) Improper Neutralization of Directives in Statically Saved Code ('Static Code Injection')

**Supported languages:** JavaScript and TypeScript, Python, Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (33) File Inclusion

**CWE** (98) Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')

**Supported languages:** PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (34) Improper Neutralization of CRLF Sequences in HTTP Headers

**CWE** (113) Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Request/Response Splitting')

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (35) Disabled Neutralization of CRLF Sequences in HTTP Headers

**CWE** (113) Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Request/Response Splitting')

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (36) Process Control

**CWE** (114) Process Control

**Supported languages:** Java, Kotlin, Scala

## Rule (37) Log Forging

**CWE** (117) Improper Output Neutralization for Logs

**Supported languages:** C# and ASP.NET

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A09:2021 - Security Logging and Monitoring Failures

## Rule (38) Buffer Overflow

**CWE** (122) Heap-based Buffer Overflow

**Supported languages:** C++

## Rule (39) Potential buffer overflow from usage of unsafe function

**CWE** (122) Heap-based Buffer Overflow

**Supported languages:** C++

## Rule (40) Potential Negative Number Used as Index

**CWE** (125, 787) Out-of-bounds Read, Out-of-bounds Write

**Supported languages:** C++

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (41) Size Used as Index

**CWE** (125, 787) Out-of-bounds Read, Out-of-bounds Write

**Supported languages:** C++

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (42) Buffer Over-read

**CWE** (126) Buffer Over-read

**Supported languages:** JavaScript and TypeScript

## Rule (43) Use of Externally-Controlled Format String

**CWE** (134) Use of Externally-Controlled Format String

**Supported languages:** C++, Java, JavaScript and TypeScript, Kotlin, Scala

## Rule (44) Memory Allocation Of String Length

**CWE** (170) Improper Null Termination

**Supported languages:** C++

## Rule (45) Improper Null Termination

**CWE** (170) Improper Null Termination

**Supported languages:** C++

## Rule (46) Integer Overflow

**CWE** (190) Integer Overflow or Wraparound

**Supported languages:** C++

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (47) Clear Text Logging

**CWE** (200, 312) Exposure of Sensitive Information to an Unauthorized Actor, Cleartext Storage of Sensitive Information

**Supported languages:** Go, Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (48) Clear Text Sensitive Storage

**CWE** (200, 312) Exposure of Sensitive Information to an Unauthorized Actor, Cleartext Storage of Sensitive Information

**Supported languages:** Apex, JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (49) Information Exposure

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**Supported languages:** C# and ASP.NET, Java, JavaScript and TypeScript, Kotlin, PHP, Ruby, Scala, Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (50) File Access Enabled

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (51) Introspection Enabled

**CWE** (200) Exposure of Sensitive Information to an Unauthorized Actor

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (52) Observable Timing Discrepancy (Timing Attack)

**CWE** (208) Observable Timing Discrepancy

**Supported languages:** Java, JavaScript and TypeScript, Kotlin, Scala

## Rule (53) Generation of Error Message Containing Sensitive Information

**CWE** (209) Generation of Error Message Containing Sensitive Information

**Supported languages:** Go

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (54) Server Information Exposure

**CWE** (209) Generation of Error Message Containing Sensitive Information

**Supported languages:** Java, Kotlin, Python, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (55) Debug Features Enabled

**CWE** (215) Insertion of Sensitive Information Into Debugging Code

**Supported languages:** C# and ASP.NET, Visual Basic, XML

## Rule (56) Unprotected Storage of Credentials

**CWE** (256) Plaintext Storage of a Password

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (57) Use of Hardcoded Credentials

**CWE** (259, 798) Use of Hard-coded Password, Use of Hard-coded Credentials

**Supported languages:** Apex, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic, XML

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (58) Use of Sticky broadcasts

**CWE** (265) Privilege Issues

**Supported languages:** Java, Kotlin

## Rule (59) Android Uri Permission Manipulation

**CWE** (266) Incorrect Privilege Assignment

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (60) Improper Handling of Insufficient Permissions or Privileges

**CWE** (280) Improper Handling of Insufficient Permissions or Privileges

**Supported languages:** Java, Kotlin, Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (61) Access Violation

**CWE** (284, 285) Improper Access Control, Improper Authorization

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (62) Binding to all network interfaces may open service to unintended traffic

**CWE** (284) Improper Access Control

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (63) Improper Access Control: Email Content Injection

**CWE** (284) Improper Access Control

**Supported languages:** Apex, Go

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (64) Session Manipulation

**CWE** (285) Improper Authorization

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (65) Anonymous LDAP binding allows a client to connect without logging in

**CWE** (287) Improper Authentication

**Supported languages:** C++

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (66) Broken User Authentication

**CWE** (287) Improper Authentication

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (67) Device Authentication Bypass

**CWE** (287) Improper Authentication

**Supported languages:** Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (68) Improper Authentication

**CWE** (287) Improper Authentication

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (69) Authentication Bypass by Spoofing

**CWE** (290) Authentication Bypass by Spoofing

**Supported languages:** C++

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (70) Improper Certificate Validation

**CWE** (295) Improper Certificate Validation

**Supported languages:** Go, Java, Kotlin, Python, Ruby, Scala, Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (71) Improper Validation of Certificate with Host Mismatch

**CWE** (297) Improper Validation of Certificate with Host Mismatch

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (72) Cryptographic Issues

**CWE** (310) Cryptographic Issues

**Supported languages:** Java, JavaScript and TypeScript, Kotlin, Python, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (73) Selection of Less-Secure Algorithm During Negotiation (Force SSL)

**CWE** (311, 757) Missing Encryption of Sensitive Data, Selection of Less-Secure Algorithm During Negotiation ('Algorithm Downgrade')

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (74) The cipher text is equal to the provided input plain text

**CWE** (311) Missing Encryption of Sensitive Data

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (75) Cleartext Storage of Sensitive Information in a Cookie

**CWE** (315) Cleartext Storage of Sensitive Information in a Cookie

**Supported languages:** C# and ASP.NET, Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (76) ASP SSL Disabled

**CWE** (319) Cleartext Transmission of Sensitive Information

**Supported languages:** XML

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (77) Authentication over HTTP

**CWE** (319) Cleartext Transmission of Sensitive Information

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (78) Cleartext Transmission of Sensitive Information

**CWE** (319) Cleartext Transmission of Sensitive Information

**Supported languages:** Java, JavaScript and TypeScript, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (79) Insecure Data Transmission

**CWE** (319) Cleartext Transmission of Sensitive Information

**Supported languages:** Apex, C# and ASP.NET, Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (80) Use of Hardcoded Cryptographic Key

**CWE** (321) Use of Hard-coded Cryptographic Key

**Supported languages:** C++, Python, Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (81) Inadequate Padding for Public Key Encryption

**CWE** (326) Inadequate Encryption Strength

**Supported languages:** PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (82) Inadequate Padding for AES encryption

**CWE** (326) Inadequate Encryption Strength

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (83) Inadequate Encryption Strength

**CWE** (326) Inadequate Encryption Strength

**Supported languages:** C++, C# and ASP.NET, Go, Java, Kotlin, PHP, Python, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (84) Use of a Broken or Risky Cryptographic Algorithm

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**Supported languages:** C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (85) Insecure TLS Configuration

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**Supported languages:** Go, JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (86) Weak Cryptographic Primitive

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**Supported languages:** COBOL (Beta)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (87) Missing protocol in ssl.wrap\_socket

**CWE** (327) Use of a Broken or Risky Cryptographic Algorithm

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (88) Use of Hardcoded Cryptographic Initialization Value

**CWE** (329) Generation of Predictable IV with CBC Mode

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (89) Use of Insufficiently Random Values

**CWE** (330) Use of Insufficiently Random Values

**Supported languages:** C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (90) Origin Validation Error

**CWE** (346, 942) Origin Validation Error, Permissive Cross-domain Policy with Untrusted Domains

**Supported languages:** Java, JavaScript and TypeScript, Kotlin, PHP, Python, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (91) Insecure JWT Verification Method

**CWE** (347) Improper Verification of Cryptographic Signature

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (92) JWT Signature Verification Method Disabled

**CWE** (347) Improper Verification of Cryptographic Signature

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (93) JWT 'none' Algorithm Supported

**CWE** (347) Improper Verification of Cryptographic Signature

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (94) JWT Signature Verification Bypass

**CWE** (347) Improper Verification of Cryptographic Signature

**Supported languages:** Java

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (95) Anti-forgery token validation disabled

**CWE** (352) Cross-Site Request Forgery (CSRF)

**Supported languages:** C# and ASP.NET

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (96) Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**Supported languages:** Java, JavaScript and TypeScript, Kotlin, PHP, Python, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (97) Spring Cross-Site Request Forgery (CSRF)

**CWE** (352) Cross-Site Request Forgery (CSRF)

**Supported languages:** Java

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (98) Exposure of Private Personal Information to an Unauthorized Actor

**CWE** (359) Exposure of Private Personal Information to an Unauthorized Actor

**Supported languages:** C++, C# and ASP.NET

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (99) Division By Zero

**CWE** (369) Divide By Zero

**Supported languages:** C++

## Rule (100) Insecure Temporary File

**CWE** (377) Insecure Temporary File

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (101) Denial of Service (DoS) through Nested GraphQL Queries

**CWE** (400) Uncontrolled Resource Consumption

**Supported languages:** JavaScript and TypeScript

## Rule (102) Unchecked Input for Loop Condition

**CWE** (400, 606) Uncontrolled Resource Consumption, Unchecked Input for Loop Condition

**Supported languages:** JavaScript and TypeScript

## Rule (103) Regular expression injection

**CWE** (400, 730) Uncontrolled Resource Consumption, OWASP Top Ten 2004 Category A9 - Denial of Service

**Supported languages:** Apex, C# and ASP.NET, Java, Kotlin, Scala, Visual Basic

## Rule (104) Regular Expression Denial of Service (ReDoS)

**CWE** (400) Uncontrolled Resource Consumption

**Supported languages:** JavaScript and TypeScript, PHP, Python, Ruby

## Rule (105) Missing Release of Memory after Effective Lifetime

**CWE** (401) Missing Release of Memory after Effective Lifetime

**Supported languages:** C++

## Rule (106) Double Free

**CWE** (415) Double Free

**Supported languages:** C++

## Rule (107) Use After Free

**CWE** (416) Use After Free

**Supported languages:** C++

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (108) Insecure default value

**CWE** (453) Insecure Default Variable Initialization

**Supported languages:** Python

## Rule (109) Android Fragment Injection

**CWE** (470) Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (110) Unsafe Reflection

**CWE** (470) Use of Externally-Controlled Input to Select Classes or Code ('Unsafe Reflection')

**Supported languages:** Java, Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (111) Dereference of a NULL Pointer

**CWE** (476) NULL Pointer Dereference

**Supported languages:** C++

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (112) Android Debug Mode Enabled

**CWE** (489) Active Debug Code

**Supported languages:** XML

## Rule (113) Debug Mode Enabled

**CWE** (489) Active Debug Code

**Supported languages:** Python

## Rule (114) Struts Development Mode Enabled

**CWE** (489) Active Debug Code

**Supported languages:** XML

## Rule (115) Trust Boundary Violation

**CWE** (501) Trust Boundary Violation

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (116) Deserialization of Untrusted Data

**CWE** (502) Deserialization of Untrusted Data

**Supported languages:** C# and ASP.NET, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (117) Insecure Deserialization

**CWE** (502) Deserialization of Untrusted Data

**Supported languages:** Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (118) No Weak Password Requirements

**CWE** (521) Weak Password Requirements

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (119) Privacy Leak

**CWE** (532) Insertion of Sensitive Information into Log File

**Supported languages:** Java, PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A09:2021 - Security Logging and Monitoring Failures

## Rule (120) Hardcoded Secret

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**Supported languages:** Apex, COBOL (Beta), C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (121) Use of Hardcoded, Security-relevant Constants

**CWE** (547) Use of Hard-coded, Security-relevant Constants

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (122) Request Validation Disabled

**CWE** (554) ASP.NET Misconfiguration: Not Using Input Validation Framework

**Supported languages:** C# and ASP.NET, Visual Basic, XML

## Rule (123) Open Redirect

**CWE** (601) URL Redirection to Untrusted Site ('Open Redirect')

**Supported languages:** Apex, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (124) Insecure Xml Parser

**CWE** (611) Improper Restriction of XML External Entity Reference

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (125) XML External Entity (XXE) Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**Supported languages:** C++, C# and ASP.NET, Java, JavaScript and TypeScript, Kotlin, PHP, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (126) XAML Injection

**CWE** (611) Improper Restriction of XML External Entity Reference

**Supported languages:** C# and ASP.NET

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (127) Insufficient Session Expiration

**CWE** (613) Insufficient Session Expiration

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (128) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**CWE** (614) Sensitive Cookie in HTTPS Session Without 'Secure' Attribute

**Supported languages:** Apex, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (129) Unverified Password Change

**CWE** (620) Unverified Password Change

**Supported languages:** Apex

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (130) Weak Password Recovery Mechanism for Forgotten Password

**CWE** (640) Weak Password Recovery Mechanism for Forgotten Password

**Supported languages:** JavaScript and TypeScript, PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (131) XPath Injection

**CWE** (643) Improper Neutralization of Data within XPath Expressions ('XPath Injection')

**Supported languages:** C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A03:2021 - Injection

## Rule (132) SQL SELECT statement without WHERE clause

**CWE** (668) Exposure of Resource to Wrong Sphere

**Supported languages:** COBOL (Beta)

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (133) Use of Potentially Dangerous Function

**CWE** (676) Use of Potentially Dangerous Function

**Supported languages:** Java, Kotlin, Scala

## Rule (134) Android World Writeable/Readable File Permission Found

**CWE** (732) Incorrect Permission Assignment for Critical Resource

**Supported languages:** Java, Kotlin, Scala

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (135) Insecure File Permissions

**CWE** (732) Incorrect Permission Assignment for Critical Resource

**Supported languages:** Python

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (136) Incorrect Permission Assignment

**CWE** (732) Incorrect Permission Assignment for Critical Resource

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (137) Selection of Less-Secure Algorithm During Negotiation (SSL instead of TLS)

**CWE** (757) Selection of Less-Secure Algorithm During Negotiation ('Algorithm Downgrade')

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (138) Allocation of Resources Without Limits or Throttling

**CWE** (770) Allocation of Resources Without Limits or Throttling

**Supported languages:** JavaScript and TypeScript, PHP

## Rule (139) Missing Release of File Descriptor or Handle after Effective Lifetime

**CWE** (775) Missing Release of File Descriptor or Handle after Effective Lifetime

**Supported languages:** C++

## Rule (140) Memory Corruption

**CWE** (822) Untrusted Pointer Dereference

**Supported languages:** Swift

## Rule (141) Unrestricted Android Broadcast

**CWE** (862) Missing Authorization

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (142) Use of Expired File Descriptor

**CWE** (910) Use of Expired File Descriptor

**Supported languages:** C++

## Rule (143) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**CWE** (915) Improperly Controlled Modification of Dynamically-Determined Object Attributes

**Supported languages:** Ruby

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A08:2021 - Software and Data Integrity Failures

## Rule (144) Use of Password Hash With Insufficient Computational Effort

**CWE** (916) Use of Password Hash With Insufficient Computational Effort

**Supported languages:** Apex, C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A02:2021 - Cryptographic Failures

## Rule (145) Server-Side Request Forgery (SSRF)

**CWE** (918) Server-Side Request Forgery (SSRF)

**Supported languages:** Apex, C++, C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Scala, Swift, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A10:2021 - Server-Side Request Forgery (SSRF)

**OWASP Top 10/SANS 25:** SANS/CWE Top 25

## Rule (146) Insecure Data Storage

**CWE** (922) Insecure Storage of Sensitive Information

**Supported languages:** Swift

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A01:2021 - Broken Access Control

## Rule (147) Android Intent Forwarding

**CWE** (940) Improper Verification of Source of a Communication Channel

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (148) Code Execution via Third Party Package Installation

**CWE** (940) Improper Verification of Source of a Communication Channel

**Supported languages:** Java, Kotlin

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A07:2021 - Identification and Authentication Failures

## Rule (149) Permissive Cross-domain Policy

**CWE** (942) Permissive Cross-domain Policy with Untrusted Domains

**Supported languages:** JavaScript and TypeScript

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (150) NoSQL Injection

**CWE** (943) Improper Neutralization of Special Elements in Data Query Logic

**Supported languages:** Java, JavaScript and TypeScript, Python

## Rule (151) Sensitive Cookie Without 'HttpOnly' Flag

**CWE** (1004) Sensitive Cookie Without 'HttpOnly' Flag

**Supported languages:** C# and ASP.NET, Go, Java, JavaScript and TypeScript, Kotlin, PHP, Python, Ruby, Scala, Visual Basic

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A05:2021 - Security Misconfiguration

## Rule (152) Bad Coding Practices

**CWE** (1006) Bad Coding Practices

**Supported languages:** JavaScript and TypeScript

## Rule (153) Improper Restriction of Rendered UI Layers or Frames

**CWE** (1021) Improper Restriction of Rendered UI Layers or Frames

**Supported languages:** JavaScript and TypeScript, PHP

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A04:2021 - Insecure Design

## Rule (154) Python 2 source code

**CWE** (1104) Use of Unmaintained Third Party Components

**Supported languages:** Python

**OWASP Top 10/SANS 25:** OWASP Top Ten 2021 Category A06:2021 - Vulnerable and Outdated Components

## Rule (155) User Controlled Pointer

**CWE** (1285) Improper Validation of Specified Index, Position, or Offset in Input

**Supported languages:** C++

## Rule (156) Incorrect regular expression for validating values

**CWE** (1286) Improper Validation of Syntactic Correctness of Input

**Supported languages:** Ruby

## Rule (157) Improper Type Validation

**CWE** (1287) Improper Validation of Specified Type of Input

**Supported languages:** JavaScript and TypeScript

## Rule (158) Prototype Pollution

**CWE** (1321) Improperly Controlled Modification of Object Prototype Attributes ('Prototype Pollution')

**Supported languages:** JavaScript and TypeScript

## Rule (159) An optimizing compiler may remove memset non-zero leaving data in memory

**CWE** (1330) Remanent Data Readable after Memory Erase

**Supported languages:** C++
