# Snyk Code AI Engine

Snyk Code is powered by a semantic, AI-based analysis engine, and can analyze the following in your code:

## **Hardcoded secrets**

Snyk Code invokes hardcoded secrets detection rules during SAST scans but does not act as a standalone secrets scanning tool. For an enhanced secrets solution, see our partnership with [GitGuardian](https://snyk.io/blog/supercharge-app-security-code-to-cloud/).

<figure><img src="../../../.gitbook/assets/Introduction - AI Engine - Hardcoded secrets.png" alt="Hardcoded secret found"><figcaption><p>Hardcoded secret found</p></figcaption></figure>

## **Coding issues**

Finds problems such as dead code, branches that are predefined, and branches having the same code on each side.

## **Type inference**

Determines the initial type and its changes. This is of special interest for dynamically typed languages.

## **Value ranges**

Infers possible values for variables used to call functions to track off-by-one errors in arrays, division-by-zero errors, and null dereferences.

## **Data flow**

Follows the flow of data within the application from the source to the sink. Combined with AI-based learning of external insecure data sources, data sinks, and sanitation functions, this enables a strong taint analysis.

## **API usage**

Identifies multiple potential issues including API misuses, null dereferences, and type mismatches by modeling the use of memory in variables and references. This mechanism can also identify the use of insecure functions.

## **Control flow**

Identifies null dereference or race conditions by modeling each possible control flow in the application.

## **Point-to analysis**

Identifies multiple potential issues including buffer overruns, null dereferences, and type mismatches by modeling the use of memory in variables and references.
