---
title: This directive content guid...
---

This directive content guides the agent to:

* Execute one or more security tests.
* Filter the results if any parameters are provided.
* Rescan to validate the security fix resolved the issues and did not introduce new ones.
* Generate a pull request that you can review, approve, and merge into the codebase.

This directive is invoked by calling the directive in an agent chat interaction, for example, `/snyk-fix`. This is invoked without any arguments, which results in any detected findings being remediated, or it is invoked with specific Snyk product filters or other indications of which resulting findings should be fixed:
