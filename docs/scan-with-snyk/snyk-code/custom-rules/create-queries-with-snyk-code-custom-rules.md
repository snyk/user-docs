# Create query

To use Snyk Code custom rules to create queries with [suggestive AI support](./#suggestive-ai-support), you can choose from provided [templates](./#template) and [predicates.](./#predicate-pred) Alternatively, you can create your own predicates and [save them as a custom rule](create-snyk-code-custom-rules.md).&#x20;

Consider the following query examples and rules to use with Snyk Code custom rules.

## Simple syntactical query

Copy the following source code snippet in the snippet window and select **C#** as the language

{% hint style="info" %}
It is only a snippet and not a full program. It will not compile.
{% endhint %}

<pre class="language-csharp"><code class="lang-csharp"><strong>// Read request body
</strong>string body;
using (var reader = new StreamReader(context.Request.Body))
{
   body = await reader.ReadToEndAsync();
}
// Parse JSON data
var form = JsonConvert.DeserializeObject&#x3C;SignupForm>(body);
var sql = String.Format("INSERT INTO submissions(email, name) VALUES('%s', '%s')", form.Email, form.Name);
form.Email = "nobody@notrealdomain.co.uk";
using var cmd = new NpgsqlCommand(sql, conn);
</code></pre>

### Running the query

Enter the following queries in the query window and press **Run Query** to see the results.

1. Select `body` by using the query: `“body”`&#x20;

{% hint style="info" %}
This query does not select the Body with a capital B. The query language is case-sensitive.
{% endhint %}

2. Add `Body` to the findings so the query becomes `Or<”body”,”Body”>`.
3. You can achieve the same outcome using a regex `~"body|Body"` or `~"[Bb]ody"`
4. Do something more complex regex and query: \
   ``~"[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"``\
   It matches the hardcoded email address.

<figure><img src="../../../.gitbook/assets/simple syntactical query.png" alt="Syntactical query example"><figcaption><p>Syntactical query example</p></figcaption></figure>

### Try it yourself

Run the following query over your code `~"([a-zA-Z0-9+/]{40})"` If you find something, check it out first, as you might leak your AWS secrets.

If you are interested in a certain type of object, you can use [templates](templates-and-predicates.md). For example, the query `CallExpression<"Format">` matches a function call or `Literal<"nobody@notrealdomain.co.uk">` matches the string with the email address.

## A data flow or taint analysis

For this example, a JavaScript code snippet is used. You can copy it in the snippet window and select **JavaScript**.

```javascript
const express = require('express');
const bodyParser = require('body-parser');
const { Client } = require('pg');
const fs = require('fs');


const app = express();
app.use(bodyParser.json());


const client = new Client({
   host: 'localhost',
   user: 'youruser',
   password: 'yourpassword',
   database: 'yourdbname'
});


async function connectDb(client) {
   await client.connect();
}


async function insertSubmission(client, email, name) {
   await client.query(`INSERT INTO submissions(email, name) VALUES(${email}, ${name})`);
}


function logSubmission(email, name) {
   const logMessage = `New submission: Email=${email}, Name=${name}\n`;
   fs.appendFileSync('myapp.log', logMessage);
}


app.post('/signup', async (req, res) => {
   try {
       const { email, name } = req.body;
       await insertSubmission(client, email, name);
       logSubmission(email, name);
       res.send({ message: 'Signup successful!' });
   } catch (err) {
       console.error(err);
       res.status(500).send({ message: 'An error occurred.' });
   }
});


connectDb(client).then(() => {
   app.listen(3000, () => console.log('Server is running on port 3000'));
});

```

Snyk Code knows a list of possible sources of external data in the predicate `PRED:AnySource`. The following query shows you that `app.post()` is identified.&#x20;

Query `PRED:SqliSinks` shows you that `query()` is part of that list of SQL injection sinks. The query engine comes with many different predicates for various source, sink, and sanitizer types. Check the list of predicates to see them all.

To check whether the data flows into a SQL injection sink, use the following: `DataFlowsInto<PRED:SqliSink>`. It shows you that in the program, data from the `req` parameter flow into `query()` taking several turns.

If the data flow is also going through a sanitizer, you can use a specialized template. Change the query to ​​`Taint<PRED:AnySource, PRED:SqliSanitizer, PRED:SqliSink>`

{% hint style="info" %}
There is nothing language-specific in the query. It would work on similar code in other languages.
{% endhint %}

## **Net new data flow rule**

Create a new rule because Snyk is not aware of the proprietary source built in-house, resulting in missed findings.

Use a data flow [template](templates-and-predicates.md) known as `Taint` when [creating a data flow query](run-query.md#run-query-on-a-repository).&#x20;

```javascript
Taint<PRED:"SourceFoo",PRED:XssSanitizer,PRED:XssSink>
```

You can configure the following parameters:

* **Source:** The first parameter indicates where the data flow starts.
* **Sanitizer:** The second parameter indicates a known sanitizer that would sanitize the data, resulting in the data's not being tainted
* **Sink**_**:**_ The third parameter indicating where the data flow ends

Custom [predicates](templates-and-predicates.md) are indicated by writing their names within brackets. In this scenario, the custom method is called `SourceFoo`.

With this query, you can look for the data flow that originates in `SourceFoo`. A source unknown to Snyk ends up in a known vulnerable cross-site scripting (XSS) Sink and does not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

## **Extend a data flow rule**

Recreate a Snyk rule and add a source to the current Snyk known vulnerable source list because they are not being taken into account in the scans, resulting in missed vulnerabilities.&#x20;

Like the [Net new data flow rule](create-queries-with-snyk-code-custom-rules.md#net-new-data-flow-rule), the `Taint` data flow template is used with an `Or` operator. Operators are available to create logical statements for your queries, such as `Or` or `And`.

Run the data flow rule using both the Snyk known sources but also a custom source called [`SourceFoo`](#user-content-fn-1)[^1]_._

```javascript
Taint<Or<PRED:AnySource,"SourceFoo">,PRED:XssSanitizer,PRED:XssSink>
```

With this query, you look for the data flow that originates in a Snyk known source OR “`SourceFoo`” . A source unknown to Snyk ends up in a known vulnerable cross-site scripting (XSS) Sink and does not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

Any statement that uses an operator will be written within angle brackets  _`< statement >`_.&#x20;

## **Context added to data flow rule**

Recreate a Snyk rule and remove a source from the current Snyk known vulnerable sources because this source is not vulnerable within the context of an application.&#x20;

Like the [Net new data flow](create-queries-with-snyk-code-custom-rules.md#net-new-data-flow-rule) and [Extend a data flow](create-queries-with-snyk-code-custom-rules.md#extend-a-data-flow-rule) rules, the `Taint` data flow template is used with an `And` operator. A declarative negative statement (`Not`) is used to indicate the false case of the statement and not the true case.

Run the data flow rule using the Snyk known sources, removing `SnykSource` from the results. In this example, `SnykSource` is a Snyk-known source that is used within the regular general `AnySource` [predicate](templates-and-predicates.md).

```javascript
Taint<And<PRED:AnySource,Not<PRED:”SnykSource”>>,PRED:XssSanitizer,PRED:XssSink>
```

With this query, you look for the data flow that originates in a known Snyk source but remove results that come from `SnykSource` that end up in a known vulnerable cross-site scripting (XSS) Sink and do not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

## **High recall mode**

See all the sources and sinks in the source code to understand every location where data can flow from or to. This analysis is conducted regardless of the presence of data flows, allowing users to assess coverage comprehensively.

This mode is often used as an investigatory tool to gain deeper insights into the code stack and to comprehend where data originates and terminates within the application.

To initiate this analysis, use the `AnySource` or `AnySink` predicates, which encompass all Snyk-known sources and sinks, respectively.

```javascript
PRED:AnySource
```

This query identifies and highlights every known source within the code being analyzed.

```
PRED:AnySink
```

Similarly, this query identifies and highlights every known sink within the code, providing a complete view of potential data endpoints.

## Assessing coverage and identifying missing data flow links

Identifying gaps in data flow paths is crucial for understanding the destinations of user data, especially if it does not end up in expected locations such as databases or file systems. These gaps may reveal the use of unsupported libraries or frameworks—or components thereof—potentially leading to false negatives. This insight is essential for comprehensive security assessments and ensuring robust coverage.

<details>

<summary><strong>Java example: Interaction between custom WebServer and WebServlet</strong></summary>

This Java example demonstrates two components: `WebServer` and `WebServlet`.

* **WebServer**: A custom HTTP server that represents the use of an unsupported or proprietary component.
* **WebServlet**: Uses Java's standard Servlet API for web interactions but connects to a custom database for user record queries.

```java
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.sql.Connection;
import java.sql.Statement;
import java.util.Map;

public class CoverageDemo {
    static class WebServer extends MyWebEndpoint {
        private Connection connection;

        @Override
        void handlePostRequest(Map<String, String> parameters) throws Throwable {
            final String username = parameters.get("username");
            final String query = String.format("SELECT * FROM users WHERE user = '%s'", username);
            final Statement statement = connection.createStatement();
            statement.execute(query);
            statement.close();
        }
    }

    static class WebServlet extends HttpServlet {
        private MySpecialDatabase database;

        @Override
        protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
            final String username = req.getParameter("username");
            final String query = String.format("SELECT * FROM users WHERE user = '%s'", username);
            database.performSQL(query);
        }
    }
}

```

</details>

Out of the box, Snyk will not show any vulnerabilities for these two classes, so the following query will not yield any results in the code snippet above.

```
Taint<PRED:AnySource,  PRED:None, PRED:AnySink>
```

### Find unmatched sources

The following query enhances security coverage by revealing unmatched sources, pinpointing situations where the `HttpServletRequest` parameter in the`WebServlet`'s `doPost` method is not linked to known sinks, thus identifying gaps in data handling.

```starlang
PRED:AnySource and not DataFlowsInto<Taint<PRED:AnySource, PRED:None, PRED:AnySink>>
```

### Find unmatched sinks

Similarly, to improve coverage, another query locates unmatched sinks by finding elements like the `java.sql.Connection` object in the `WebServer` class that are poised to receive data but lack incoming data flows, highlighting areas for potential gaps in library and framework coverage.

```starlang
PRED:AnySink and not DataFlowsFrom<PRED:AnySource>
```

## CWE 312 query example

In the context of Static Application Security Testing (SAST), identifying vulnerabilities associated with CWE-312 poses complex challenges in terms of analysis and accuracy. Specifically, the following concerns must be addressed:

1. Data Sensitivity Classification: Accurately categorizing **which data** elements are sensitive and which are not is a non-trivial task. An incorrect classification can result in false positives.
2. Sink Protection Validation: Another challenge is the identification of **data sinks** (endpoints where data is stored or transmitted) that are adequately secure.
3. Geographical Data Protection Requirements: Moreover, the data protection regulations can vary significantly depending on the jurisdiction.

Given these complexities, custom rules offer a more flexible and tailored approach for detecting CWE-312 vulnerabilities. In this example, we will employ the C# programming language and leverage the `Taint` template within the Custom Rules framework to address these challenges.

Let's start with a simple program:

```csharp
using System;
using System.IO;
using System.Text;
using System.Threading;

namespace CWE_312_Example
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("Please enter your username: ");
            string username = Console.ReadLine();

            string userData = $"Username: {username}";

            File.WriteAllText("testFile.txt", userData);
        }
    }
}
```

Building a rule that matches on sensitive data (`username`) being sent to a text file is simple enough. A naive first approach is:

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  "global::System.IO.File.WriteAllText"
>
```

This matches a sensitive data flow from `ReadLine` to `WriteAllText`. Given the broad nature of the rule, it may create quite a lot of noise.

### Restricting to specific log files only

The first caveat is that perhaps only `testFile.txt` is considered unsafe. Files like `cache.txt` should be considered safe.&#x20;

```csharp
// Create a warning on this one
File.WriteAllText("testFile.txt", userData);

// We can ignore this file
File.WriteAllText("cache.txt", userData);
```

To achieve this, we use the `CallExpression` and `HasArg1` templates.&#x20;

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  CallExpression<"global::System.IO.File.WriteAllText">
    and
      HasArg1<"testFile.txt">
>
```

{% hint style="info" %}
**`CallExpression`** and **`HasArg1`** may also be used on their own. Connecting them using the **`and`** operator establishes the relationship, and Snyk Code will attempt to match them in combination only.
{% endhint %}

### Catching all File writers

In .NET, files can be written using [`WriteAllText`](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.writealltext?view=net-7.0). There are also `WriteAllLines` and `WriteAllBytes`.

Our code snippet may be extended this way:

```csharp
string userData = $"Username: {username}";
byte[] userBytes = Encoding.UTF8.GetBytes(userData);
string[] userLines = new string[] {userData};

File.WriteAllText("testFile.txt", userData);
File.WriteAllLines("testFile.txt", userLines);
File.WriteAllBytes("testFile.bin", userBytes);
```

Let's capture the variants first using regular expressions. We will look for the functions and also both filename variants (`.bin` and `.txt`):

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  > 
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

{% hint style="info" %}
Notice how `CallExpression now` contains a **regular expression,** whereas **`HasArg1`** uses the **`or`** operator. It could be written either way.
{% endhint %}

Finally, let's add support for the .NET`Async` variants and also the `Append` methods:

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  >
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

### Defining "interesting data"

In our previous example, we treated `ReadLine` as a source of sensitive data. Let's consider a simple object with only specific sensitive fields:

```csharp
public class MyUser
{
    public string EmailAddress;

    public string MembershipType;
}
```

For this class, only `EmailAddress` should be considered.

Therefore, given a code snippet like this:

```csharp
MyUser user = new MyUser();
user.EmailAddress = "support@snyk.io";
user.MembershipType = "SampleRole";

string sensitiveData = $"Username: {user.EmailAddress}";
string notSensitiveData = $"MembershipType: {user.MembershipType}";

File.WriteAllText("testFile.txt", sensitiveData);
File.WriteAllText("testFile.txt", notSensitiveData);
```

The first call to `WriteAllText` should be prevented while the second call is allowed. To accomplish this, you can defer to specifying the field name in the query:

```ada
Taint<
  "EmailAddress",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  >
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

[^1]: 
