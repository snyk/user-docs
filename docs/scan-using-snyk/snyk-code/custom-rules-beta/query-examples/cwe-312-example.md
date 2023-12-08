---
description: Cleartext Storage of Sensitive Information
---

# CWE 312 Example

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

```
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
  And<
    CallExpression<"global::System.IO.File.WriteAllText">, 
    HasArg1<"testFile.txt">
  >
>
```

{% hint style="info" %}
**`CallExpression`** and **`HasArg1`** may also be used on it's own. Connecting them using the **`And`** statement establishes the relationship & Snyk Code will attempt to match them in combination only.
{% endhint %}

### Catching all File writers

In .NET, files can be written using [`WriteAllText`](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.writealltext?view=net-7.0). There'a also `WriteAllLines` and `WriteAllBytes`.

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
  And<
    CallExpression<
      ~"global::System\.IO\.File\.WriteAll(Text|Lines|Bytes)"
    >, 
    HasArg1<Or<"testFile.txt", "testFile.bin">>
  >
>
```

{% hint style="info" %}
Notice how CallExpression now contains a **regular expression,** whereas **`HasArg1`** utilises the **`Or`** template. It could be written either way.
{% endhint %}

Finally, let's add support for .NETs `Async` variants and also the `Append` methods:

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  And<
    CallExpression<
      ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
    >, 
    HasArg1<Or<"testFile.txt", "testFile.bin">>
  >
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

The first call to `WriteAllText` should be prevented while the second call is allowed. In order to accomplish this, you can simply defer to specifying the field name in the query:

```ada
Taint<
  "EmailAddress",
  PRED:None,
  And<
    CallExpression<
      ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
    >, 
    HasArg1<Or<"testFile.txt", "testFile.bin">>
  >
>
```

