# How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI

* [ Unable to test Go Modules repository using CLI](/hc/en-us/articles/360015917317-Unable-to-test-Go-Modules-repository-using-CLI)
* [ Can I use Python 3 instead of Python 2?](/hc/en-us/articles/360014386678-Can-I-use-Python-3-instead-of-Python-2-)
* [ Private Gem support for Ruby projects](/hc/en-us/articles/360006435718-Private-Gem-support-for-Ruby-projects)
* [ How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI](/hc/en-us/articles/360004167317-How-to-install-the-SBT-dependency-graph-plugin-to-test-Scala-projects-with-Snyk-CLI)
* [ Does Snyk support Mongo Express Node \(MERN\) stack?](/hc/en-us/articles/360004048137-Does-Snyk-support-Mongo-Express-Node-MERN-stack-)
* [ How do I scan .NET project in CLI?](/hc/en-us/articles/360004026097-How-do-I-scan-NET-project-in-CLI-)
* [ I can't import or scan a Python project](/hc/en-us/articles/360002553098-I-can-t-import-or-scan-a-Python-project)
* [ Gradle Java project: java.lang.OutOfMemoryError: Java heap space](/hc/en-us/articles/360001857918-Gradle-Java-project-java-lang-OutOfMemoryError-Java-heap-space)
* [ What does it mean if Snyk offers remediation advice only regarding my out of date dependencies?](/hc/en-us/articles/360001667158-What-does-it-mean-if-Snyk-offers-remediation-advice-only-regarding-my-out-of-date-dependencies-)
* [ Java - Gradle : Variables appearing in dependency names in SCM](/hc/en-us/articles/360001451358-Java-Gradle-Variables-appearing-in-dependency-names-in-SCM)

 [See more](/hc/en-us/sections/360000935397-Language-support)

##  How to install the SBT dependency graph plugin to test Scala projects with Snyk CLI

The need for the sbt plugin is due to the fact that there is no way for us to understand the dependencies, other languages and package managers have a way to give us this list \(such as a lockfile\), but sbt does not. Therefore, the use of the plugin is to understand and get a view of the relationship between the different dependencies. 

**The following steps are required:**

1. Install Homebrew  Homebrew will be used to install both sbt and scala
2. Install sbt then run sbt by typing sbt in the CLI
3. Install scala
4. Install the sbt dependency graph plugin
5. Configure sbt dependency graph plugin

**1. Install Homebrew \(Homebrew will allow you to install sbt and scala\)**

1. 1. In your browser navigate to: brew.sh
   2. Copy the command displayed under the “Install Homebrew” banner
   3. Paste that command into the command line in terminal
   4. Type “Brew” to run Brew

**2. Now install Sbt**

1. 1. Type “brew install sbt”
      1. NOTE: You may have to run the command below first 
      2. “brew cask install adoptopenjdk” \(Brew will let you know if you do need to do this\)
      3. If you did need to run this extra command, then you will have to run the “brew install sbt” again afterward
      4. Run sbt by typing _sbt_ in the command line

**3. Now install Scala**

1. Type “brew install scala” NOTE: Although other package managers provide a way to allow Snyk to determine the list/graph/tree of dependencies, Sbt does not. So an Sbt dependency graph plugin is required to address this issue.

**4. Now follow these instructions for installing the Sbt dependency graph plugin**  
  
In order to use the Snyk CLI to test Scala projects, you will need to install the Sbt dependency graph plugin. 

Installing the Sbt dependency graph plugin for sbt 0.13  
Prerequisites

* Ensure you have installed Scala. \(See steps outlined above\)
* Ensure you have installed Sbt and ran sbt. \(See steps outlined above\)

NOTE: The steps below will install the Sbt dependency plugin as a global plugin.

1. First navigate to the correct directory by typing the following command: _cd ~/.sbt_
2. This will take you to the Sbt directory. From there you will need to navigate to the 0.13 directory. Typing the _ls_ command will show if 0.13 and/or 1.0 exists in the directory
3. Navigate to 0.13 by typing: _cd_ _0.13_ and then make a directory called plugins by typing: _mkdir plugins_
4. Navigate to the new directory by typing: _cd plugins_ and then proceed to create a file called “plugins.sbt” by typing: _touch plugins.sbt_
5. Edit the plugins.sbt file via the CLI using the nano command: _nano plugins.sbt_ 
6. \(you can also use a text editor is you prefer\)
7. Add the following line to the file using this command: _addSbtPlugin\("net.virtual-void" % "sbt-dependency-graph" % "0.10.0-RC1"\)_
8. Press: “_Control X”_ to exit, Press: “_Y_” when prompted to save the file, then press: _return_ to close the nano editor 

5. Take the following steps for the 1.0 directory. Check if 1.0 exists by typing _ls_ in the sbt directory:

1. If the 1.0 does NOT exist in the sbt directory, type mkdir 1.0 in the sbt directory
2. If 1.0 exists in the directory, run the following command: _cd ~/.sbt/1.0_
3. Make a directory called “plugins” in that folder by typing: _mkdir plugins_
4. Copy the existing “plugins.sbt” file from the 0.13 directory to the current 1.0 directory by typing the following: _cp ../0.13/plugins/plugins.sbt ./plugins_
5. Validate that the plugin has been installed correctly by running the following command:  _sbt "-Dsbt.log.noformat=true" dependencyTree_  \(This should be tested in the directory of the project and running the command will generate the dependency graph. You can also run it each time you want to generate the dependency graph\) 

You should now be able to successfully run “snyk test” via the CLI to uncover vulnerabilities in the scala project you imported. 

