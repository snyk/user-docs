# Missing or differing results in Snyk Code

## .gitignore and .dcignore files <a href="#gitignore-and-.dcignore-files" id="gitignore-and-.dcignore-files"></a>

* Rename all `.gitignore` and `.dcignore` files to, for example, `.dcignore.bak.`
* Rerun Snyk Code.
* If the result is now better, check the rules in `.gitignore` and `.dcignore` files to see if they cause the ignoring of files.
* Ignores can be analyzed in Jetbrains by enabling debug logging for the Snyk Code logger.
* Ignores can be analyzed in the Language Server by enabling trace logging, for example, by setting the env-variable `SNYK_DEBUG_LEVEL=trace.`
* Enable ignore patterns step-by-step until results are correct, and the only files that are ignored should be ignored.
* If the `.dcignore` is filled with a lot of ignores and was created automatically, it should be emptied. IntelliJ before 2.4.63 added this file automatically when no `.gitignore` was found at the top level. It contained many rules that can interfere with repositories and source files.

## Snyk Code Quality <a href="#snyk-code-quality" id="snyk-code-quality"></a>

* The IDE plugins and Language Server report Snyk Code Quality results.
* The CLI does not report Snyk Code Quality results.
* This can lead to a difference in reported issues.
