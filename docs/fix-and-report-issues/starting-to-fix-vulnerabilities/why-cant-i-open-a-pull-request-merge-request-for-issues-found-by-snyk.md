# Why can't I open a Pull Request/Merge Request for issues found by Snyk?

You would not necessarily get a Fix PR button for your project whenever you import it from your Source Control tool or run a scan. Snyk only opens PR for direct dependencies.

If your project has transitive dependencies that contains vulnerabilities, Snyk would not open PRs against them.

