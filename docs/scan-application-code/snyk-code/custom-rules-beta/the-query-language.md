# The Query Language

Snyk Code provides Custom Queries using a domain-specific language for code search. It is a logic declarative programming language that is not Turing complete - in our case. This brings the advantage that every query written in the querying language is guaranteed to terminate and return zero, one or more matches.

The querying language is independent of the programming language used for the code and rules work across all Snyk-supported languages. If a code snippet is provided, you have to pick the language of the provided code snippet.

| Note: The querying language is case-sensitive. |
| ---------------------------------------------- |

The querying language is a language for finding matches in code. Every query discovers some elements in the queried code for which the specified properties match.

The first capability of the querying language is to match program elements by their value. This is done by quoting the value in double quotes (`"`). Program elements are identified by their fully qualified names. Consider the following example Java code:

```java
import java.time.LocalDate;
class Test {
 static void test() {
   System.out.println("test" + 123);
   System.out.println(LocalDate.now());
 }
}
```

One can match the method call for taking the current time by quoting its fully qualified name, i.e. with the query `"java.time.LocalDate.now"`.  One can match both the function declaration _test_ and the string '`test`' by using the query `"test"`. The number value 123 can be matched by using the query `"123"`. Quotes are used to match elements regardless of their type - identifier, string, number, or other value. Elements can also be matched by using regular expressions. Regular expressions are identified by putting the symbol `~` in front of the quotes. For example, the program element 123 can be matched by the expression `~"12.*".` The print statements can be matched by queries such as: `"java.lang.System.out.println"` or `~".*\.println"`.

To make sure that the correct fully qualified names of elements are used, the search interface provides autocompletion of the values for program elements that exist in the given code snippet or the provided repository.

**Predicate (**_**PRED**_**):** A predicate matches program elements based on some predefined condition. The main advantage of predicates is that one can use them to leverage the existing Snyk Code knowledge base. For example, if one needs to find all program locations where an HTTP server handles cookies, one can use the predefined predicate `PRED:SourceCookie`. Similarly, there is a predicate `PRED:SqliSink` for matching all program locations where SQL queries are handled. To discover all available predicates, custom rules provide autocompletion capabilities. There are two special predicates `PRED:Any` and `PRED:None` that match all program elements or no program elements, respectively.

| Note: When multiple matches are provided in a sequence, the result is a combination for all of them. E.g. using a query like: `PRED:SourceCookie ~"get.*"` will only match on methods that are both returning cookies and their name starts with get (i.e. logical AND of the two conditions that match elements) |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

**Template:** Templates are used for combining one or more conditions provided as their parameters. Templates themselves, just like predicates and value matches, also describe rules for matching elements in the given code. The parameters of templates are provided in chevrons (or angle brackets, i.e. `<` and `>`) and are comma separated. The predefined templates are designed for multiple use cases that we define below.

Templates can be used to restrict the matches of their parameters. E.g. `StringLiteral<"test">` takes all the program elements with value test and only returns the elements that are string literals.

Templates can be used to relate different elements of the program. For example, the following query will find all program entities that have the string literal test as first argument:\
`HasArg1<StringLiteral<"test">>`.

Note that the template `HasArg1` encodes semantic relations between program elements. For example, the above query will match only for the first print statement in the following Java code:

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



Some templates encode logical relations, like conjunction (And) and disjunctions (Or). The following find all calls to `println` that receives the string literal test as the first argument:  `And<"java.lang.System.out.println", HasArg1<StringLiteral<"test">>>`

### Formal Syntax

The following two paragraphs provide the formal definition and relation of the querying language to Datalog. It is not necessarily needed to successfully use the querying language but it is provided for completeness.

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

A query can contain one or more terms. Semantically, each term must be satisfied for the query to match (i.e. conjunction of the terms). A term is either a literal, regular expression, predicate or a template. Literals and regular expressions are quotes, regular expressions are preceded by the symbol `~`. Predicates define complex conditions program elements should satisfy. Templates are used to combine multiple predicates.

#### Semantics and Relationship to Datalog

Datalog is not a Turing complete language, and every program in Datalog is guaranteed to terminate. The query language itself is a subset of Datalog with mostly focus on unary predicates - i.e. a query either matches for a program element or does not. The query language is not designed to define new relationships between pairs of program elements. Still, it only allows querying existing relations as computed by the program analysis in Snyk Code, such as dataflow, taint, and others.

The way existing relations are queried is via templates. Templates are essentially pre-built snippets of Datalog code with some parts left as holes that are then replaced by the Customizing Snyk Rules (V 2) values in the templates used at their instantiation. Some templates include recursion in their implementation, which allows the custom query rule to itself include recursion, but only in the shape defined in the templates.

As a result, Snyk Code Query Language becomes a subset of Datalog, which effectively disallows computations that could have large time or memory complexity. As a result, Snyk Code Query Language is not only guaranteed to terminate, but will produce its matches fast. In the experience of Snyk, the restrictions that Snyk Code Query Language puts on Datalog rarely affect its expressibility in security-scanning settings.

\


\
\
\
\
\
