# Predefined predicates and templates

## Predicates

### Any

A "catchall" rule. Matches on anything.

### AnySource

Matches on various types of potentially user-controlled data sources, both servers, for example, HTTP parameters/header/body, URLs, cookies, and so on, or indirect ones such as database fields, local files, I/O, or environment variables.

### DeserializationSanitizer

Matches on deserialization sanitizers.

### DeserializationSink

Matches on deserialization sinks.

### None

An "anti-catchall" rule. Matches on nothing.

### PTSanitizer

Matches on path-traversal sanitizers.

### PtSink

Matches on path-traversal sinks.

### SourceArchive

Matches on reading values that are coming from zip, tar, or other archives.

### SourceCLI

Matches on reading command line arguments.

### SourceClientFramework

Matches on reading values that are coming from a client-side framework such as Android, SwiftUI, UIKit, or the DOM of an HTML page.

### SourceCookie

Matches on reading values of cookies on an http server. These values are of security interest because they can be fully controlled by malicious users.

### SourceDatabase

Matches on reading values that are coming from a database.

### SourceEnvironmentVariable

Matches on reading environment variables of a process.

### SourceFile

Matches on reading values that are coming from files.

### SourceHttpBody

Matches on reading the http request body on an http server. These values are of security interest because they may be fully controlled by malicious actors.

### SourceHttpFileUpload

Matches on the name and content of the file uploaded to an http server. These values are of security interest because they may be fully controlled by malicious actors.

### SourceHttpHeader

Matches on reading values of http headers on a server. These values are of security interest because they may be fully controlled by malicious actors.

### SourceHttpParam

Matches on reading values of http parameters on an http server. These values are of security interest because they may be fully controlled by malicious actors.

### SourceLocalEnv

Matches on reading values from the local environment of the running process. This includes command line arguments, standard input, or environment variables.

### SourceNonServer

Matches on reading values that may be controlled by an adversary but not directly by sending requests to a server. For example, if an application fetches a value from a URL, an adversary in control of that URL may use it to control its content.

### SourceRequestUrl

Matches on reading request URLs on a server. The URLs are of security interest because they may be fully controlled by malicious actors.

### SourceResourceAccess

Matches on reading values that may be controlled by an adversary if they gain access to a resource. The resources this matches are remote URLs, files, database fields, or other framework-specific cases such as Android intents.

### SourceRpcApiParam

Matches on parameters of RPCs implemented in an RPC server. These values are of security interest because they may be fully controlled by malicious actors.

### SourceServer

Matches on reading values that an attacker can send to a server. Examples are http parameters/header/body, URLs, or cookies. Since these may be directly controllable by attackers, these sources are of significant security interest.

### SourceStdin

Matches on reading input from the standard input of a process.

### SourceWebForm

Matches on reading values of web forms in a web server. These values are of security interest because they may be fully controlled by malicious actors.

### SqliSanitizer

Matches on SQL injection sanitizers.

### SqliSink

Matches on SQL injection sinks.

### SsrfSanitizer

Matches on SSRF sanitizers.

### SsrfSink

Matches on SSRF sinks.

### XssSanitizer

Matches on XSS sanitizers.

### XssSink

Matches on XSS sinks.

## Templates

### And

A binary conjunction. Matches only if both arguments match.

Template parameters:

* First conjunct
* Second conjunct

### Arg0In

Matches on the 0th index argument, that is, on the receiver object for method calls for the provided method or function.

Template parameter: Function

### Arg1In

Matches on the first index argument for the provided method or function.

Template parameter: Function

### Arg2In

Matches on the second index argument for the provided method or function.

Template parameter: Function

### Arg3In

Matches on the third index argument for the provided method or function.

Template parameter: Function

### Arg4In

Matches on the fourth index argument for the provided method or function.

Template parameter: Function

### Arg5In

Matches on the fifth index argument for the provided method or function.

Template parameter: Function

### Arg6In

Matches on the sixth index argument for the provided method or function.

Template parameters: Function

### Arg7In

Matches on the seventh index argument for the provided method or function.

Template parameter: Function

### BooleanLiteral

Matches on boolean type literals.

Template parameter: Value

### CallExpression

Matches when a given name is called.

Template parameter: Callee Function or method to call.

### DataFlowAfter

Matches on entities that happen after in the dataflow of its parameter.

Template parameter: PrevAction The previous action executed.

<details>

<summary>Example</summary>

In the following code snippet, `replaceAll` is used as a sanitizer for data read from the source `scanner.nextLine()`:

```java
import java.util.Scanner;

public class RegexSanitize {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String input = scanner.nextLine();
        input = input.replaceAll("[^a-zA-Z0-9 ]", "");
        System.out.println(input);
    }
}
```

To verify **replaceAll** with specific parameters is part of this **call chain**, the following rule can be used:

```starlang
And<CallExpression<"replaceAll">, HasArg1<"[^a-zA-Z0-9 ]">>
HasArg0<DataFlowAfter<
  And<
    CallExpression<"nextLine">, 
    And<
      HasArg0<"java.util.Scanner">, 
      HasArg0<HasArg1<"java.lang.System.in">>
    >
  >
>>
```

Using `HasArg0`, a relationship to `DataFlowAfter` is established in order to express that a `replaceAll` call (matched by the preceding `And` template) appears on the data flow path after `scanner.nextLine()`.

***

Note this will only match on the function call being actually executed, regardless of whether `input` is re-assigned with the sanitised value or not.&#x20;

It does not match on the data being sent to `System.out.println`. Combine this function with `Taint` to achieve such a check.

```starlang
Taint<
  And<
    And<CallExpression<"replaceAll">, HasArg1<"[^a-zA-Z0-9 ]">>,
    HasArg0<DataFlowAfter<
      And<
        CallExpression<"nextLine">, 
        And<
          HasArg0<"java.util.Scanner">, 
          HasArg0<HasArg1<"java.lang.System.in">>
        >
      >
    >>
  >,
  PRED:None,
  "java.lang.System.out.println"
>
```

</details>

### DataFlowsFrom

Matches on places which a taint data can flow from.

Template parameter: Source

### DataFlowsInto

Matches on places that tainted data can flow into.

Template parameter: Sink

### ForSameObject

Matches on entities that happen on the same object as its parameter.

Template parameter: ObjectAction The action that happens on the object.

### HasAnnotation

Matches on entities which are decorated with a specific annotation.

Template parameter: Value

<details>

<summary>Example</summary>

```java
package snippets.java.docs;

public class HasAnnotationDemo {
    public void oldMethod(@Deprecated String parameter1, String parameter2) {

    }
}
```

To capture `parameter1`, use the following query:

```
HasAnnotation<"java.lang.Deprecated">
```

</details>

<details>

<summary>Example (named parameters)</summary>

```java
package snippets.java.docs;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target({ElementType.FIELD, ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@interface Sensitive {
    String reason();
}


public class HasAnnotationWithArgsDemo {
    private String email = "support@snyk.io";

    @Sensitive(reason = "PII")
    public String getEmail() {
        return email;
    }
}
```

To capture the `getEmail` method, use the following query:

```
HasAnnotation<And<"Sensitive", HasNamedArg<"reason", "PII">>>
```

</details>

{% hint style="warning" %}
`HasAnnotation` can not be used within `Taint` workflows yet.
{% endhint %}

### HasAnyArg

Matches on entities that take an argument at _any_ index with the provided value.

Template parameter: Value

<details>

<summary>Example</summary>

```java
package snippets.java.docs;

public class VarArgsDemo {
    public void method(String... args) {}

    public static void main(String[] args) {
        VarArgsDemo check = new VarArgsDemo();
        check.method("tainted", "sample2", "sample3");
        check.method("sample0", "tainted", "sample2", "sample3");
        check.method("sample2", "sample3", "tainted");
    }
}
```

To match all 3 method calls which receive `tainted` values, use the following query:

```
HasAnyArg<"tainted">
```

</details>

### HasArg0

Matches on entities that take an argument at the 0th index, that is, receiver object for method calls, with the provided value.

Template parameter: Value

<details>

<summary>Example (Java)</summary>

The following code snippet demonstrates two types of logging into a network service using a method named `login`.

```java
import sun.net.ftp.FtpProtocolException;
import sun.net.ftp.impl.FtpClient;

import javax.security.auth.login.LoginContext;
import javax.security.auth.login.LoginException;
import java.io.IOException;

public class X {
    private static void ftpExample() {
        FtpClient client = new FtpClient();
        try {
            client.login("user", "pass".toCharArray());
        } catch (FtpProtocolException | IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static void loginContextExample() {
        try {
            LoginContext lc = new LoginContext("MyLoginConfig");
            lc.login();
            System.out.println("Authentication succeeded!");

        } catch (LoginException e) {
            System.err.println("Authentication failed: " + e.getMessage());
        }
    }

    public static void main(String[] args) {
        ftpExample();
        loginContextExample();
    }
}
```

For this example, the target may be to **capture only the login call on the FtpClient class**. To select it, the following query will match:

```
And<HasArg0<"sun.net.ftp.impl.FtpClient">, CallExpression<"login">>
```

</details>

<details>

<summary>Example (Python)</summary>

The following code snippet demonstrates two types of logging into a network service using a method named `login`.

```python
from ftplib import FTP, error_perm
import smtplib

def ftp_example():
    try:
        ftp = FTP('ftp.example.com')
        ftp.login('user', 'pass')
        print("FTP login successful")
    except error_perm as e:
        print(f"FTP login failed: {e}")

def smtp_example():
    try:
        smtp = smtplib.SMTP('smtp.example.com', 587)
        smtp.ehlo()
        smtp.starttls()
        smtp.login('user@example.com', 'password')
        print("SMTP login successful")
    except smtplib.SMTPException as e:
        print(f"SMTP login failed: {e}")

if __name__ == '__main__':
    ftp_example()
    smtp_example()
```

For this example, the target may be to **capture only the login call on the FTP class**. To select it, the following query will match:

```starlang
And<HasArg0<"ftplib.FTP">, CallExpression<"login">>
```

</details>

### HasArg1

Matches on entities that take an argument at the first index with the provided value.

Template parameter: Value

### HasArg2

Matches on entities that take an argument at the second index with the provided value.

Template parameter: Value

### HasArg3

Matches on entities that take an argument at the third index with the provided value.

Template parameter: Value

### HasArg4

Matches on entities that take an argument at the fourth index with the provided value.

Template parameter: Value

### HasArg5

Matches on entities that take an argument at the fifth index with the provided value.

Template parameter: Value

### HasArg6

Matches on entities that take an argument at the sixth index with the provided value.

Template parameter: Value

### HasArg7

Matches on entities that take an argument at the seventh index with the provided value.

Template parameter: Value

### HasNamedArg

Matches on entities that take a named argument with the provided value.

Template parameters:

* Name The name of the argument.
* Value The value the named argument should have.

### Identifier

Matches on an identifier.

Template parameter: Name The entity that should be an identifier.

### Literal

Matches on string/boolean or number type literals.

Template parameter: Value

### NamedArgIn

Matches on the named argument for the provided method or function.

Template parameters:

* Name The name of the argument.
* Function The provided method or function.

### Not

A negation. Matches only if the argument does not match.

Template parameter: Negated property

### NumberLiteral

Matches on numeric type literals.

Template parameter: Value

### Or

A binary disjunction. Matches if either or both arguments match.

Template parameters:

* First disjunct
* Second disjunct

### ReturnedBy

Matches on the returned entity.

Template parameter: Function The entity that returns.

### Returns

Matches on the entity, for example, a function or a method, that returns the value provided as an argument.

Template parameter: Value What is returned.

### StringLiteral

Matches on string-type literals.

Template parameter: Value

### Taint

Identifies all data flows originating from specified sources and flowing to the vulnerable method (sink) without passing through the specified sanitizers.

Template parameters:

* Source
* Sanitizer
* Sink
