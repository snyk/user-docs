# List of built-in Predicates and Templates

## List of built-in predicates

### Any

A "catchall" rule. Matches on anything.

### AnySource

Matches on various types of potentially user controlled data sources, both servers (e.g., HTTP parameters/header/body, URLs, cookies, etc.) or indirect ones such as database fields, local files, I/O or environment variables.

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

Matches on reading values that are coming from zip, tar or other archives.

### SourceCLI

Matches on reading command line arguments.

### SourceClientFramework

Matches on reading values that are coming from a client-side framework such as Android, SwiftUI, UIKit, the DOM of an HTML page.

### SourceCookie

Matches on reading values of cookies in an http server. These values are of security interest, because they can be fully controlled by malicious users.

### SourceDatabase

Matches on reading values that are coming from a database.

### SourceEnvironmentVariable

Matches on reading environment variables of a process.

### SourceFile

Matches on reading values that are coming from files.

### SourceHttpBody

Matches on reading http request body in an http server. These values are of security interest, because they may be fully controlled by malicious actors.

### SourceHttpFileUpload

Matches on the name and content of file uploaded to an http server. These values are of security interest, because they may be fully controlled by malicious actors.

### SourceHttpHeader

Matches on reading values of http headers in a server. These values are of security interest, because they may be fully controlled by malicious actors.

### SourceHttpParam

Matches on reading values of http parameters in an http server. These values are of security interest, because they may be fully controlled by malicious actors.

### SourceLocalEnv

Matches on reading values from the local environment of the running process. This includes command line arguments, standard input or environment variables.

### SourceNonServer

Matches on reading values that may be controlled by an adversary, but not directly by sending requests to a server. E.g. if an application fetches a value from a URL, an adversary in control of that URL may use it to control its content.

### SourceRequestUrl

Matches on reading request URLs in a server. The URLs are of security interest, because they may be fully controlled by malicious actors.

### SourceResourceAccess

Matches on reading values that may be controlled by an adversary if they gain access to a resource. The resources this matches are remote URLs, files, database fields or other framework-specific cases such as Android intents.

### SourceRpcApiParam

Matches on parameters of RPCs implemented in an RPC server. These values are of security interest, because they may be fully controlled by malicious actors.

### SourceServer

Matches on reading values that an attacker can send to a server. Examples are HTTP parameters/header/body, URLs or cookies. Since these may be directly controllable by attacker, these sources are of significant security interest.

### SourceStdin

Matches on reading input from the standard input of a process.

### SourceWebForm

Matches on reading values of web forms in a web server. These values are of security interest, because they may be fully controlled by malicious actors.

### SqliSanitizer

Matches on SQL-injection sanitizers.

### SqliSink

Matches on SQL-injection sinks.

### SsrfSanitizer

Matches on SSRF sanitizers.

### SsrfSink

Matches on SSRF sinks.

### XssSanitizer

Matches on XSS sanitizers.

### XssSink

Matches on XSS sinks.

## List of built-in templates

### And

A binary conjunction. Matches only if both arguments match.

List of template parameters:

* First conjunct
* Second conjunct

### Arg0In

Matches on the 0th index argument (i.e. the receiver object for method calls) for the provided method or function.

List of template parameters:

* Function

### Arg1In

Matches on the 1st index argument for the provided method or function.

List of template parameters:

* Function

### Arg2In

Matches on the 2nd index argument for the provided method or function.

List of template parameters:

* Function

### Arg3In

Matches on the 3rd index argument for the provided method or function.

List of template parameters:

* Function

### Arg4In

Matches on the 4th index argument for the provided method or function.

List of template parameters:

* Function

### Arg5In

Matches on the 5th index argument for the provided method or function.

List of template parameters:

* Function

### Arg6In

Matches on the 6th index argument for the provided method or function.

List of template parameters:

* Function

### Arg7In

Matches on the 7th index argument for the provided method or function.

List of template parameters:

* Function

### BooleanLiteral

Matches on boolean type literals.

List of template parameters:

* Value

### CallExpression

Matches when a given name is called.

List of template parameters:

* Callee Function or method to call.

### DataFlowAfter

Matches on entities that happen after in the dataflow of its parameter.

List of template parameters:

* PrevAction The previous action executed.

### DataFlowsFrom

Matches on places which a taint data can flow from.

List of template parameters:

* Source

### DataFlowsInto

Matches on places which a taint data can flow into.

List of template parameters:

* Sink

### ForSameObject

Matches on entities that happen on the same object as its parameter.

List of template parameters:

* ObjectAction The action that happens on the object.

### HasArg0

Matches on entities that take an argument in the 0th index (i.e. receiver object for method calls) with the provided value.

List of template parameters:

* Value

### HasArg1

Matches on entities that take an argument in the 1st index with the provided value.

List of template parameters:

* Value

### HasArg2

Matches on entities that take an argument in the 2nd index with the provided value.

List of template parameters:

* Value

### HasArg3

Matches on entities that take an argument in the 3rd index with the provided value.

List of template parameters:

* Value

### HasArg4

Matches on entities that take an argument in the 4th index with the provided value.

List of template parameters:

* Value

### HasArg5

Matches on entities that take an argument in the 5th index with the provided value.

List of template parameters:

* Value

### HasArg6

Matches on entities that take an argument in the 6th index with the provided value.

List of template parameters:

* Value

### HasArg7

Matches on entities that take an argument in the 7th index with the provided value.

List of template parameters:

* Value

### HasNamedArg

Matches on entities that take a named argument with the provided value.

List of template parameters:

* Name The name of the argument.
* Value The value the named argument should have.

### Identifier

Matches on an identifier.

List of template parameters:

* Name The entity that should be an identifier.

### Literal

Matches on string/boolean or number type literals.

List of template parameters:

* Value

### NamedArgIn

Matches on the named argument for the provided method or function.

List of template parameters:

* Name The name of the argument.
* Function The provided method or function.

### Not

A negation. Matches only if the argument does not match.

List of template parameters:

* Negated property

### NumberLiteral

Matches on numeric type literals.

List of template parameters:

* Value

### Or

A binary disjunction. Matches if either (or both) arguments match.

List of template parameters:

* First disjunct
* Second disjunct

### ReturnedBy

Matches on the returned entity.

List of template parameters:

* Function The entity that returns.

### Returns

Matches on the entity (e.g. a function or a method) that returns the value provided as argument.

List of template parameters:

* Value What is returned.

### StringLiteral

Matches on string type literals.

List of template parameters:

* Value

### Taint

Template for taint vulnerability, for Custom Rules app, can be used with anything (not just sources/sinks).

List of template parameters:

* Source
* Sanitizer
* Sink
