---
description: The report types available in Snyk API and Web
nav_context: agnostic
---

# Report types

Snyk API & Web offers several types of target scan reports, available in PDF or DOCX format:

* Standard.
* Executive Summary.
* PCI-DSS v.4.0.1 and v.3.2.1.
* OWASP Top 10 2021 and 2025.
* ISO 27001.
* HIPAA.

Use the type of report that fits your situation.

The following sections describe the content of each report.

## Standard

This is the default report type for target scans and contains the following:

* A scan summary.
* A settings summary.
* A technical summary.
* An exhaustive test list.
* All findings.
* Information about vulnerabilities (impact, causes, and prevention methods).

## Executive Summary

This is a high-level view report of the target scan and only contains:

* A scan summary.
* A settings summary.
* A technical summary.

## PCI-DSS v.4.0.1 and v.3.2.1

This is a target scan report specific to PCI-DSS compliance. You can use any of these reports to verify which controls a target is passing or failing in the respective PCI-DSS version. The PCI-DSS report is similar to the Standard report but adds a section to the scan summary with the PCI-DSS requirements checklist.

Snyk indicates whether a target was tested for the requirements checklist and whether it passed each item on the list.

## OWASP Top 10 2021 and 2025

This is similar to the PCI-DSS report but considers OWASP Top 10 2021 or 2025. OWASP Top 10 is a popular framework provided by OWASP that lists the top 10 security risks of web applications. Auditors often use this framework when performing a company's security audit.

Snyk indicates whether a target was tested for the requirements checklist and whether it passed each item on the list.

## ISO 27001

This is a specific target scan report on compliance with ISO/IEC 27001 (2022 revision). You can use this report to verify which controls a target is passing or failing. ISO 27001 is similar to the Standard report but adds a section to the scan summary with the ISO 27001 requirements checklist.

Snyk indicates whether a target was tested for the requirements checklist and whether it passed each item on the list.

## HIPAA

This is a specific target scan report on HIPAA compliance. You can use this report to verify which controls a target is passing or failing. HIPAA is similar to the Standard report but adds a section to the scan summary with the HIPAA requirements checklist.

Snyk indicates whether a target was tested for the requirements checklist and whether it passed each item on the list.

You do not need to start a target scan again to issue a new type of report. Visit [Change report type](change-report-type.md) to generate and download a different report type.
