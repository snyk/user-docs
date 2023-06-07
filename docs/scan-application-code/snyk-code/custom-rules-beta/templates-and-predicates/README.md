# Templates and Predicates

This section provides an introduction to the Snyk Code’s Query Language standard library of templates and predicates. The goal of the section is both to give some practical examples and to give documentation on each predicate and template.

### Methods, literals and arguments

A basic capability of Snyk Code is to find method calls and reason about its arguments. The goal here is to discover certain patterns of method calls and their arguments and to check if certain properties hold for these objects. Consider the following imaginary Python program we are going to analyze and search in. The same examples will work for any other language Snyk Code supports if a similar program is provided. The code does not need to compile to be able to be queried:

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

The first query we will make is to find the method connect. The query `"connect"` will return both the string with this value and the method call with this name. We can separate these by putting the value in a template. We can use autocompletion to find how the name _connect_ can be wrapped. `Literal` or `StringLiteral` will restrict the search results to the string _'connect'_, whereas `CallExpression` will restrict the results to the function call `connect()`.

<figure><img src="../../../../.gitbook/assets/SnykCodeQueryStringvsCallExpr.png" alt=""><figcaption><p>Environment suggesting possible use for "connect"</p></figcaption></figure>

Note that you can find function calls to functions outside of the file being scanned. Trying to find `CallExpression<"safesend">` will not yield results. The reason is that the analysis may inline local functions in order to reason about their behavior.

Now, let’s look at the most called method in the file, `send`. This method is called on an object returned by connect and takes various strings as input. To see its arguments, one can use some of the templates about its arguments. These are `HasArg0`, `HasArg1`, etc.

<figure><img src="../../../../.gitbook/assets/SnykCodeQuery2ndParameterSugg.png" alt=""><figcaption><p>Environment suggests HasArg</p></figcaption></figure>

So, for example, we can find all calls to send on an object returned by connect. This is done by the following query:

`CallExpression<"send"> HasArg0<CallExpression<"connect">>`\
These are all locations, but we can find the places where we call send with the first argument taking the value _connect_.

<figure><img src="../../../../.gitbook/assets/SnykCodeQueryCallExprsSend.png" alt=""><figcaption></figcaption></figure>

This gives a more interesting picture. The interprocedural analysis figured out that the message connect was sent in a call to a local function.

Now, coming to more interesting tasks, around the state of the object o. Assume we want to find all the calls to _send_ after _disconnect_. These should be pretty bad cases of our imaginary programs where the connection may be in a bad state. To do this query, we can perform the following query:\
`CallExpression<"send"> HasArg0<DataFlowAfter<Arg0In<CallExpression<"disconnect">>>>`

This query searches for calls to _send_ with its argument 0 satisfying the following property: in the dataflow it is after a location that is argument 0 in a call to `disconnect`. This matches only the final unsafe `send` call.

Finally, we can do other things with negation, we can search for objects that are returned by `connect`, but not calling `disconnect` for the returned object.

`CallExpression<"connect"> Not<ForSameObject<Arg0In<CallExpression<"disconnect">>>>`

or similarly, calling `send` with `connect` but not calling `disconnect` (this one has no matches in the above code snippet):

`CallExpression<"send"> HasArg0<CallExpression<"connect">> Not<HasArg0<ForSameObject<Arg0In<CallExpression<"disconnect">>>>>`

In all cases, the auto-completion for the rules should guide the search through the examples to make it easier to write such and even more complex queries.

### Taint flows and data sources

In many cases, one wants to ensure that certain types of data have no way to flow to certain sensitive locations in the program. These are often done for security reasons, compliance reasons as well as ensuring correctness.

The first important element to query is sensitive data sources. Snyk has built-in the following set of interesting hierarchy of data sources that one can query:

* AnySource
  * SourceServer
    * SourceCookie
    * SourceHttpParam
    * SourceRequestUrl
    * SourceHttpHeader
    * SourceWebForm
    * SourceHttpBody
    * SourceHttpFileUpload
    * SourceRpcApiParam
  * SourceNonServer
    * SourceResourceAccess
      * SourceDatabase
      * SourceFile
      * SourceArchive
      * SourceClientFramework
    * SourceLocalEnv
      * SourceCLI
      * SourceStdin
      * SourceEnvironmentVariable

The first category of sources (SourceServer) is defined for programs that implement servers. These sources are typically fully user controllable. This means that a malicious user can use them to launch an attack against the application or that simply one needs to handle such data with additional care. For example, one may want to check that authentication is always performed or some other property that should be enforced.

The non-server predicates apply also for programs that do not implement server functionality.\


Each of the predicates in the SourceServer category is returned by querying _PRED:SourceServer_ or _PRED:AnySource_. Consider the following example TypeScript code:

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

This implements a request handler for an express server. In this case, the code reads the user cookie and then logs it on the console. This might be a security vulnerability but is also probably a compliance problem for many applications. The first capability of Snyk Code is that it can simply discover these cookie locations and one can connect them to check a lot of properties about them. In this case, simply running a _PRED:SourceCookie_ query will find the first line of the request handler.

One can now verify that cookies are handled correctly by the code, e.g. to check that cookies do not end up logged anywhere. One can try to use data flow or _ForSameObject_, but in this case we also want to report in case the cookie is logged as part of some other object, string concatenation, or other simple transformation. To do this, there is a taint analysis done with the `taint` predicate. This takes the following shape: `Taint< source, sanitizer, sink >` . _Source_ is the source of sensitive data, _sanitizer_ gives code patterns that would transform the data to be non-sensitive and _sink_ is the location where the sensitive data should reach a report to be made. The report is then made at the sink location.

Now, consider we want to find places where the user is logged. We can then use the following query:

`Taint<PRED:SourceCookie, PRED:None, CallExpression<"log">>`

Of course, one may want to say that if a cookie is hashed using the function `hash123`, then it is safe to be logged. Then, the query would look like:\
\
`Taint<PRED:SourceCookie, CallExpression<"hash123">, CallExpression<"log">>`

### Predefined sinks and sanitizers

Using the above taint template, one can start writing vulnerability detectors. However, Snyk Code provides predicates for various types of vulnerabilities. For example, if one wants to detect sql injection, this can be performed fully with the following query:

`Taint<PRED:AnySource, PRED:SqliSanitizer, PRED:SqliSink>`

Of course, this assumes that any of the sources in _AnySource_ (see hierarchy above) is one that a malicious actor may control. For example, not every application is set up in a way that environment variables or command line arguments are controlled by users. If one would want to find only such SQL injections, one can run a query like:\
\
`Taint<Or<PRED:AnySource, PRED:SourceResourceAccess>, PRED:SqliSanitizer, PRED:SqliSink>`

In addition to SQL injection, Snyk Code can detect tens of other vulnerabilities and has corresponding predicates accessible from search and custom rules. The number of predicates is growing over time and more of the rules are getting open to modifications.



\
