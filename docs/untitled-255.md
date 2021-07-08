# Which projects must be built before testing with CLI?

* [ Clojure Projects](/hc/en-us/articles/360016431777-Clojure-Projects)
* [ Snyk reports an invalid Gradle dependency or dependency version in CLI](/hc/en-us/articles/360016003758-Snyk-reports-an-invalid-Gradle-dependency-or-dependency-version-in-CLI)
* [ Which projects must be built before testing with CLI?](/hc/en-us/articles/360015552617-Which-projects-must-be-built-before-testing-with-CLI-)
* [ Snyk test - Could not find or load main class org.gradle.wrapper.GradleWrapperMain](/hc/en-us/articles/360007745957-Snyk-test-Could-not-find-or-load-main-class-org-gradle-wrapper-GradleWrapperMain)
* [ Required packages missing when testing a Python project with pip](/hc/en-us/articles/360007419257-Required-packages-missing-when-testing-a-Python-project-with-pip)
* [ Failing with exit code : Generic Error](/hc/en-us/articles/360006786558-Failing-with-exit-code-Generic-Error)
* [ How do I test a different branch of a project via the CLI?](/hc/en-us/articles/360006052837-How-do-I-test-a-different-branch-of-a-project-via-the-CLI-)
* [ What version of Node is required for Snyk?](/hc/en-us/articles/360004183317-What-version-of-Node-is-required-for-Snyk-)
* [ Out of Memory Error when testing Scala \(SBT\) project](/hc/en-us/articles/360003143417-Out-of-Memory-Error-when-testing-Scala-SBT-project)
* [ My local build is working but when using the CLI tool it fails](/hc/en-us/articles/360002879278-My-local-build-is-working-but-when-using-the-CLI-tool-it-fails)

 [See more](/hc/en-us/sections/360000923118-CLI)

##  Which projects must be built before testing with CLI?

### Introduction

For some types of projects, you must build the project before analyzing it with the [Snyk CLI](https://support.snyk.io/hc/en-us/categories/360000456217-Snyk-CLI).

This is because manifests provide some dependency information, other dependencies are only resolved after the project is built. Similarly, lock files giving dependency information may not be available.

### Which projects must be built?

The following types of projects must be built before analysis with the CLI.

<table>
  <thead>
    <tr>
      <th style="text-align:left"><b>Language</b>
      </th>
      <th style="text-align:left"><b>Project type</b>
      </th>
      <th style="text-align:left"><b>Build required</b>
      </th>
      <th style="text-align:left"><b>Notes</b>
      </th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">JavaScript</td>
      <td style="text-align:left">npm</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>package-lock.json</code> file present, run <code>npm install</code> to
        generate</td>
    </tr>
    <tr>
      <td style="text-align:left">yarn</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>yarn.lock</code> file present, run <code>yarn install</code> to
        generate</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Java</td>
      <td style="text-align:left">Maven</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">Run <code>mvn install</code> before testing</td>
    </tr>
    <tr>
      <td style="text-align:left">Gradle</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">.NET</td>
      <td style="text-align:left">nuget</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>packages.config</code> file present</td>
    </tr>
    <tr>
      <td style="text-align:left">paket</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Python</td>
      <td style="text-align:left">pip</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">
        <p>Run <code>pip install -r requirements.txt</code> before testing. Install
          is required so that the full dependency tree (nested dependencies included)
          can be tested.</p>
        <p>Alternatively, pass the CLI <code>--skip-unresolved=true</code> param when
          testing to only test the dependencies available without building</p>
      </td>
    </tr>
    <tr>
      <td style="text-align:left">pipenv</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">Run <code>pipenv update</code> before testing</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">setup.py</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">Run <code>pip install -e .</code> before testing</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Poetry</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>poetry.lock</code> file present, run <code>poetry lock</code> to
        generate</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Golang</td>
      <td style="text-align:left">dep</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">Run <code>dep ensure</code> before testing</td>
    </tr>
    <tr>
      <td style="text-align:left">govendor</td>
      <td style="text-align:left">Yes</td>
      <td style="text-align:left">Run <code>govendor sync</code> before testing</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">go modules</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left"></td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Swift / Objective-C</td>
      <td style="text-align:left">Cocoapods</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>Podfile.lock</code> file present, run <code>pod install</code> to
        generate</td>
    </tr>
    <tr>
      <td style="text-align:left">Scala</td>
      <td style="text-align:left">sbt</td>
      <td style="text-align:left">No</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Ruby</td>
      <td style="text-align:left">bundler</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>Gemfile.lock</code> file present, run <code>bundle install</code> to
        generate</td>
    </tr>
    <tr>
      <td style="text-align:left">PHP</td>
      <td style="text-align:left">composer</td>
      <td style="text-align:left">No*</td>
      <td style="text-align:left">Build only required if no <code>composer.lock</code> file present, run <code>composer install</code> to
        generate</td>
    </tr>
  </tbody>
</table>

