# Query examples

language-specificUse Snyk Code custom rules to create queries with [suggestive AI support](how-custom-rules-work.md#suggestive-ai-support). You can choose from provided [templates](how-custom-rules-work.md#query-templates) and [predicates](how-custom-rules-work.md#query-predicates). Alternatively, you can create your own predicates and [save them as a custom rule](create-custom-rules.md).&#x20;

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

Run the following query over your own code `~"([a-zA-Z0-9+/]{40})"` If you find something, check it out first, as you might leak your AWS secrets.

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

language-specific

If the data flow is also going through a sanitizer, you can use a specialized template. Change the query to ​​`Taint<PRED:AnySource, PRED:SqliSanitizer, PRED:SqliSink>`

{% hint style="info" %}
There is nothing language-specific in the query. It would work on similar code in other languages.
{% endhint %}

## **Net new data flow rule**

Create a new rule because Snyk is not aware of the proprietary source built in-house, resulting in missed findings.

Use a data flow [template](how-custom-rules-work.md#query-templates) known as `Taint` when [creating a data flow query](run-query.md#run-query-on-a-repository).&#x20;

<pre class="language-javascript"><code class="lang-javascript"><a data-footnote-ref href="#user-content-fn-1">Taint</a>&#x3C;PRED:"SourceFoo",PRED:XssSanitizer,PRED:XssSink>
</code></pre>

You can configure the following parameters:

* **Source:** The first parameter indicates where the data flow starts.
* **Sanitizer:** The second parameter indicates a known sanitizer that would sanitize the data resulting in it not being tainted
* **Sink**_**:**_ The third parameter indicating where the data flow ends

Custom [predicates](how-custom-rules-work.md#query-predicates) are indicated by writing their names within brackets. In this scenario, the custom method is called `SourceFoo`.

With this query, you can look for the data flow that originates in `SourceFoo`. A source unknown to Snyk ends up in a known vulnerable cross-site scripting (XSS) Sink and does not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

## **Extend a data flow rule**

Recreate a Snyk rule and add a source to the current Snyk known vulnerable source list because they are not being taken into account in the scans, resulting in missed vulnerabilities.&#x20;

Like the [Net new data flow rule](query-examples.md#net-new-data-flow-rule), the `Taint` data flow template is used with an `Or` operator. Operators are available to create logical statements for your queries, such as `Or` or `And`.

Run the data flow rule using both the Snyk known sources but also a custom source called [`SourceFoo`](#user-content-fn-2)[^2]_._

```javascript
Taint<Or<PRED:AnySource,"SourceFoo">,PRED:XssSanitizer,PRED:XssSink>
```

With this query, you look for the data flow that originates in a Snyk known source OR “`SourceFoo`” . A source unknown to Snyk ends up in a known vulnerable cross-site scripting (XSS) Sink and does not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

Any statement that uses an operator will be written within angle brackets  _`< statement >`_.&#x20;

## **Context added to data flow rule**

Recreate a Snyk rule and remove a source from the current Snyk known vulnerable sources because this source is not vulnerable within the context of an application.&#x20;

Like the [Net new data flow](query-examples.md#net-new-data-flow-rule) and [Extend a data flow](query-examples.md#extend-a-data-flow-rule) rules, the `Taint` data flow template is used with an `And` operator. A declarative negative statement (`Not`) is used to indicate the false case of the statement and not the true case.

Run the data flow rule using the Snyk known sources, removing `SnykSource` from the results. In this example, `SnykSource` is a Snyk known source that is used within the regular general `AnySource` [predicate](how-custom-rules-work.md#query-predicates).

```javascript
Taint<And<PRED:AnySource,Not<PRED:”SnykSource”>>,PRED:XssSanitizer,PRED:XssSink>
```

With this query, you look for the data flow that originates in a known Snyk source but remove results that come from `SnykSource` that end up in a known vulnerable cross-site scripting (XSS) Sink and do not pass through a known cross-site scripting (XSS) Sanitizer. Therefore, the assumption is that the data is tainted.

## **High recall mode**

See all the sources in the code no matter what to understand all the places that data can flow from.

Find all the sources available in their code, no matter where they end up or the data flow. This is often used as an investigatory query to understand the code stack better and where data can come from in the application.

To run this query, call the `AnySource` [predicate](how-custom-rules-work.md#query-predicates) that hosts all the Snyk known sources.&#x20;

```javascript
PRED:AnySource
```

This query highlights all Snyk known sources within the code that it is being tested against.

[^1]: C

[^2]: 
