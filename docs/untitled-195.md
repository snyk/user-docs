# Snyk test - Could not find or load main class org.gradle.wrapper.GradleWrapperMain

##  Snyk test - Could not find or load main class org.gradle.wrapper.GradleWrapperMain

### **Problem**:

When testing a gradle project, you may receive the following Gradle error:

```text
>>> stderr:
Error: Could not find or load main class org.gradle.wrapper.GradleWrapperMain
Caused by: java.lang.ClassNotFoundException: org.gradle.wrapper.GradleWrapperMain
```

### **Discussion**:

This solution was found and confirmed to fix an issue for a Snyk client from this link: [https://stackoverflow.com/questions/29805622/could-not-find-or-load-main-class-org-gradle-wrapper-gradlewrappermain](https://stackoverflow.com/questions/29805622/could-not-find-or-load-main-class-org-gradle-wrapper-gradlewrappermain)

Your gradle wrapper is missing, broken or corrupted.

**What is gradle wrapper?**  
`gradlew` is the gradle wrapper executable - batch script on windows and shell script elsewhere. The wrapper script when invoked, downloads the defined gradle version and executes it. By distributing the wrapper with your project, anyone can work with it without needing to install Gradle beforehand. Even better, users of the build are guaranteed to use the version of Gradle that the build was designed to work with.

### **Restoring Gradle wrapper:**

Most newer versions of gradle now have a built-in task to resolve this. Just run:

```text
gradle wrapper
```

It used to be that you needed to add a `wrapper` task to your build.gradle to restore gradle wrapper and all its dependencies. For instance:

```text
task wrapper(type: Wrapper) {
    gradleVersion = '4.1'
}
```

You can also supply additional flags to specify versions etc

```text
gradle wrapper --gradle-version 6.2 --distribution-type all
```

When you run this task, a gradle wrapper script, and the required jar files are added to your source folders. Properties are stored in gradle/wrapper/gradle-wrapper.properties

\(You may need to install gradle locally to run this. `brew install gradle` on mac for instance. See more detailed instructions [here](https://docs.gradle.org/current/userguide/installation.html#installation)\)

**Why was it missing in the first place?**  
OP seems to have deleted something that gradle wrapper depends on.

But a common reason is that a .gitignore entry prevents wrapper jars from being checked into git. Note that the .gitignore in effect may be in the source folder, or a global one in your user home folder or git global configuration. It is common to have a `*.jar` entry in .gitignore.

You can add an exception for gradlew's jar files in .gitignore

```text
*.jar
!gradle/wrapper/gradle-wrapper.jar
```

or force add the wrapper jar into git

```text
git add -f gradle/wrapper/gradle-wrapper.jar
```

ref: [Gradle Wrapper](http://gradle.org/docs/current/userguide/gradle_wrapper.html)

