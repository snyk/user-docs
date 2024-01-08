# Templates and predicates

This section provides an introduction to the Snyk Code Query Language standard library of templates and predicates, to give some practical examples and documentation on each predicate and template.

## Methods, literals, and arguments

A basic capability of Snyk Code is to find method calls and reason about their arguments. The goal here is to discover certain patterns of method calls and their arguments and to check if certain properties hold for these objects.&#x20;

Consider the following Python program to be analyzed and searched. If a similar program is provided, the same examples will work for any other language Snyk Code supports. The code does not need to be compiled to be queried.

```python
def safesend(x, y):
 lock(x)
 x.send(y)
 unlock(x)


def finalsend(x):
 x.send('final')
 x.disconnect()


o = connect()
o1 = connect()
safesend(o, 'connect')
safesend(o, 'message')
finalsend(o)
o.send('unsafe')
```

The first query finds the method `connect`. The query `"connect"`returns both the string with this value and the method call with this name. These can be separated by putting the value in a template. We can use autocompletion to find how the name `connect` can be wrapped. `Literal` or `StringLiteral` will restrict the search results to the string _`'connect'`_, whereas `CallExpression` will restrict the results to the function call `connect()`.

<figure><img src="../../../.gitbook/assets/SnykCodeQueryStringvsCallExpr.png" alt="Environment suggesting possible use for &#x22;connect&#x22;"><figcaption><p>Environment suggesting possible use for "connect"</p></figcaption></figure>

Note that you can find function calls to functions outside of the file being scanned. Trying to find `CallExpression<"safesend">` will not yield results. The reason is that the analysis may inline local functions in order to reason about their behavior.

Look at the most called method in the file, `send`. This method is called on an object returned by connect and takes various strings as input. To see its arguments, you can use some of the templates for its arguments. These are `HasArg0`, `HasArg1`, and so on.

<figure><img src="../../../.gitbook/assets/SnykCodeQuery2ndParameterSugg.png" alt="Environment suggests HasArg"><figcaption><p>Environment suggests HasArg</p></figcaption></figure>

For example, you can find all calls to send on an object returned by `connect` with the following query:

`CallExpression<"send"> HasArg0<CallExpression<"connect">>`\
These are all locations, but we can find the places where we call send with the first argument taking the value `connect`.

<figure><img src="../../../.gitbook/assets/SnykCodeQueryCallExprsSend.png" alt="Call send with the first argument taking the value connect"><figcaption><p>Call send with the first argument taking the value connect</p></figcaption></figure>

This gives a different picture. The interprocedural analysis figured out that the message connect was sent in a call to a local function.

Different tasks around the state of the object `o` can be explored. Assume we want to find all the calls to _send_ after _disconnect_. These should be pretty bad cases of the programs where the connection may be in a bad state. To do this query, you can perform the following query:\
`CallExpression<"send"> HasArg0<DataFlowAfter<Arg0In<CallExpression<"disconnect">>>>`

This query searches for calls to _send_ with its argument `0` satisfying the following property: in the dataflow, it is after a location that is an argument `0` in a call to `disconnect`. This matches only the final unsafe `send` call.

For negation, you can search for objects that are returned by `connect`, but not calling `disconnect` for the returned object.

`CallExpression<"connect"> Not<ForSameObject<Arg0In<CallExpression<"disconnect">>>>`

Similarly, you can call `send` with `connect` but not call `disconnect`. The following example has no matches in the preceding code snippet:

`CallExpression<"send"> HasArg0<CallExpression<"connect">> Not<HasArg0<ForSameObject<Arg0In<CallExpression<"disconnect">>>>>`

In all cases, the auto-completion for the rules should guide the search through the examples to make it easier to write such queries and those that are even more complex.

## Taint flows and data sources

In many cases, you want to ensure that certain types of data have no way to flow to certain sensitive locations in the program. This is often done for security reasons, both to ensure compliance and correctness.

The first important element to query is sensitive data sources. Snyk has built in the following set of  hierarchical data sources that you can query:

[AnySource](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#anysource)

* [SourceServer](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourceserver)
  * [SourceCookie](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcecookie)
  * [SourceHttpParam](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcehttpparam)
  * [SourceRequestUrl](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcerequesturl)
  * [SourceHttpHeader](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcehttpheader)
  * [SourceWebForm](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcewebform)
  * [SourceHttpBody](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcehttpbody)
  * [SourceHttpFileUpload](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcehttpfileupload)
  * [SourceRpcApiParam](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcerpcapiparam)
* [SourceNonServer](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcenonserver)
  * [SourceResourceAccess](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourceresourceaccess)
    * [SourceDatabase](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcedatabase)
    * [SourceFile](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcefile)
    * [SourceArchive](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcearchive)
    * [SourceClientFramework](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourceclientframework)
  * [SourceLocalEnv](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcelocalenv)
    * [SourceCLI](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcecli)
    * [SourceStdin](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourcestdin)
    * [SourceEnvironmentVariable](../../../scan-application-code/snyk-code/custom-rules-beta/list-of-predefined-predicates-and-templates.md#sourceenvironmentvariable)

The first category of sources (SourceServer) is defined for programs that implement servers. These sources are typically fully user-controllable. This means that a malicious user can use them to launch an attack against the application or that one needs to handle such data with additional care. For example, you may want to check that authentication is always performed or that some other property is enforced.

The non-server predicates also apply to programs that do not implement server functionality.\


Each of the predicates in the SourceServer category is returned by querying _`PRED:SourceServer`_ or _`PRED:AnySource`_. Consider the following TypeScript code example:

```javascript
import { Request, Response, NextFunction } from 'express';


module.exports = function productReviews () {
 return (req: Request, res: Response, next: NextFunction) => {
   let user = req.signedCookies;
   doSomething(user);
   console.log('Some message ' + user);
 }
}

```

This implements a request handler for an express server. In this case, the code reads the user cookie and logs it on the console. This might be a security vulnerability and a compliance problem for many applications. The first capability of Snyk Code is that it can discover these cookie locations, and you can connect them to check a lot of properties about them. In this case, running a _`PRED:SourceCookie`_ query will find the first line of the request handler.

You can now verify that cookies are handled correctly by the code. For example, you can check that cookies do not end up logged anywhere. You can try to use data flow or _`ForSameObject.`_ In this case, report if the cookie is logged as part of some other object, string concatenation, or other simple transformation.&#x20;

To achieve this, there is a taint analysis done with the `taint` predicate. This takes the following shape: `Taint< source, sanitizer, sink >` .&#x20;

_Source_ is the source of sensitive data, _sanitizer_ gives code patterns that would transform the data to be non-sensitive, and _sink_ is the location where the sensitive data should reach a report to be made. The report is then made at the sink location.

Now, consider finding places where the user is logged in. You can then use the following query:

`Taint<PRED:SourceCookie, PRED:None, CallExpression<"log">>`

Of course, one may want to say that if a cookie is hashed using the function `hash123`, then it is safe to be logged. Then, the query would look like:\
\
`Taint<PRED:SourceCookie, CallExpression<"hash123">, CallExpression<"log">>`

## Predefined sinks and sanitizers

Using the preceding taint template, you can start writing vulnerability detectors. However, Snyk Code provides predicates for various types of vulnerabilities. For example, if you want to detect SQL injection, this can be performed fully with the following query:

`Taint<PRED:AnySource, PRED:SqliSanitizer, PRED:SqliSink>`

Of course, this assumes that any of the sources in _`AnySource`_ (see the hierarchy above) is one that a malicious actor may control. For example, not every application is set up in a way that users control environment variables or command line arguments. If you want to find only such SQL injections, you can run a query like:\
\
`Taint<Or<PRED:AnySource, PRED:SourceResourceAccess>, PRED:SqliSanitizer, PRED:SqliSink>`

In addition to SQL injection, Snyk Code can detect tens of other vulnerabilities and has corresponding predicates accessible from search and custom rules. The number of predicates is growing over time, and more rules are getting open to modifications.



\
