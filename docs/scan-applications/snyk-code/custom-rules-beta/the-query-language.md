# The query language

Snyk Code provides custom queries using a domain-specific language for code search. In our case, it is a logic declarative programming language that is not Turing complete. This brings the advantage that every query written in the query language is guaranteed to terminate and return zero, one, or more matches.

The query language is independent of the programming language used for the code, and rules work across all Snyk-supported languages. If a code snippet is provided, you have to pick the language of the provided code snippet.

{% hint style="info" %}
The query language is case-sensitive.
{% endhint %}

The query language is a language for finding matches in code. Every query discovers some elements in the queried code for which the specified properties match.

The first capability of the query language is to match program elements by their value. This is done by quoting the value in double quotes `"`. Program elements are identified by their fully qualified names. Consider the following Java code example:

```java
import java.time.LocalDate;
class Test {
 static void test() {
   System.out.println("test" + 123);
   System.out.println(LocalDate.now());
 }
}
```

You can match the method call for taking the current time by quoting its fully qualified name with the query `"java.time.LocalDate.now"`. &#x20;

You can match both the function declaration _test_ and the string '`test`' by using the query `"test"`.&#x20;

The number value 123 can be matched by using the query `"123"`. Quotes are used to match elements regardless of their type, identifier, string, number, or other value. Elements can also be matched by using regular expressions. Regular expressions are identified by putting the symbol `~` in front of the quotes. For example, the program element `123` can be matched by the expression `~"12.*".` The print statements can be matched by queries such as: `"java.lang.System.out.println"` or `~".*\.println"`.

To make sure that the correct, fully-qualified names of elements are used, the search interface provides autocompletion of the values for program elements that exist in the given code snippet or the provided repository.

## **Predicate (**_**PRED**_**)**&#x20;

A predicate matches program elements based on some predefined condition. The main advantage of predicates is that you can use them to leverage the existing Snyk Code knowledge base. For example, if you need to find all program locations where an HTTP server handles cookies, you can use the predefined predicate `PRED:SourceCookie`.&#x20;

Similarly, there is a predicate `PRED:SqliSink` for matching all program locations where SQL queries are handled. To support discovering all available predicates, custom rules provide autocompletion capabilities. There are two special predicates `PRED:Any` and `PRED:None` that match all program elements or no program elements, respectively.

{% hint style="info" %}
When multiple matches are provided in a sequence, the result is a combination of all of them.&#x20;

For example, using a query like: `PRED:SourceCookie ~"get.*"` will only match on methods that are both returning cookies and have a name that starts with `get` (logical AND of the two conditions that match elements).
{% endhint %}

## **Template**

Templates are used for combining one or more conditions provided as their parameters. Templates themselves, just like predicates and value matches, also describe rules for matching elements in the given code. The parameters of templates are provided in chevrons or angle brackets, `<` and `>,` and are comma separated. The predefined templates are designed for multiple use cases that are defined here.

Templates can be used to restrict the matches of their parameters. For example `StringLiteral<"test">` takes all the program elements with value test and only returns the elements that are string literals.

Templates can be used to relate different elements of the program. For example, the following query will find all program entities that have the string literal test as the first argument:\
`HasArg1<StringLiteral<"test">>`.

Note that the template `HasArg1` encodes semantic relations between program elements. For example, the preceding query will match only for the first print statement in the following Java code:

```java
class Test {
  String x;
  void test() {
    this.x = "test";
    System.out.println(x);
    this.x = "test2";
    System.out.println(x);
  }
}
```

Some templates encode logical relations, like the conjunction `And` and disjunctions `Or`. The following query finds all calls to `println` that receives the string literal test as the first argument:  `And<"java.lang.System.out.println", HasArg1<StringLiteral<"test">>>`

## Formal syntax

The following two paragraphs provide the formal definition and relation of the query language to Datalog. This information is not needed to use the query language successfully but is provided for completeness.

The syntax of a query is defined as:

```ebnf
<query> ::= <term> | <term> “ ” <query>
<term> ::= <literal> | <regexp> | <predicate> | <template>
<literal> ::= “"” <value> “"”
<regexp> ::= “~"” <value> “"”
<predicate> ::= “PRED:” <predicate-name>
<template> ::= <template-name> “<” <template-params>  “>”
<template-params> ::= <term> | <term> “,” <template-params>
```

A query can contain one or more terms. Semantically, each term must be satisfied for the query to match, such as the conjunction of the terms. A term is a literal, regular expression, predicate, or template. Literals and regular expressions are in quotes; regular expressions are preceded by the symbol `~`. Predicates define complex conditions program elements should satisfy. Templates are used to combine multiple predicates.

## Semantics and relationship to Datalog

Datalog is not a Turing complete language, and every program in Datalog is guaranteed to terminate. The query language is a subset of Datalog, focusing mostly on unary predicates. That is, a query either matches a program element or does not. The query language is not designed to define new relationships between pairs of program elements. It allows only querying existing relations as computed by the program analysis in Snyk Code, such as dataflow and taint.

Existing relations are queried via templates. Templates are essentially predefined snippets of Datalog code with some parts left as holes that are replaced by the customizing Snyk rules (v2) values in the templates used at their instantiation. Some templates include recursion in their implementation, allowing the custom query rule to include recursion, but only in the shape defined in the templates.

As a result, the Snyk Code Query Language becomes a subset of Datalog, effectively disallowing computations that could have large time or memory complexity. As a result, Snyk Code Query Language is guaranteed to terminate and will produce its matches fast. In the experience of Snyk, the restrictions that Snyk Code Query Language puts on Datalog rarely affect its expressibility in security-scanning settings.

\


\
\
\
\
\
