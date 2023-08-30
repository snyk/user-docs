# List of predefined predicates and templates

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

### DataFlowsFrom

Matches on places which a taint data can flow from.

Template parameter: Source

### DataFlowsInto

Matches on places that tainted data can flow into.

Template parameter: Sink

### ForSameObject

Matches on entities that happen on the same object as its parameter.

Template parameter: ObjectAction The action that happens on the object.

### HasArg0

Matches on entities that take an argument in the 0th index, that is, receiver object for method calls, with the provided value.

Template parameter: Value

### HasArg1

Matches on entities that take an argument in the first index with the provided value.

Template parameter: Value

### HasArg2

Matches on entities that take an argument in the second index with the provided value.

Template parameter: Value

### HasArg3

Matches on entities that take an argument in the third index with the provided value.

Template parameter: Value

### HasArg4

Matches on entities that take an argument in the fourth index with the provided value.

Template parameter: Value

### HasArg5

Matches on entities that take an argument in the fifth index with the provided value.

Template parameter: Value

### HasArg6

Matches on entities that take an argument in the sixth index with the provided value.

Template parameter: Value

### HasArg7

Matches on entities that take an argument in the seventh index with the provided value.

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
