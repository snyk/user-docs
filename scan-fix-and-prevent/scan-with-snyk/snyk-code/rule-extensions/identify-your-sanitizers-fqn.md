---
description: How to identify the fully qualified name (FQN) of a sanitizer function for a Snyk Code Rule Extension
nav_context: agnostic
---

# Identify your sanitizer's FQN

When creating a custom sanitizer, you must provide its Fully Qualified Name (FQN). The FQN is how Snyk Code's security engine identifies your function across your codebase. This guide walks you through identifying the correct FQN format for your sanitizer function.

## Why FQNs matter

Snyk Code traces data flow through your application to identify security issues. To recognize your custom sanitizer, the engine must resolve the function's name to its exact definition. The name you see in your IDE may differ from how the engine resolves it, especially when imports, namespaces, or module systems are involved.

**Critical rule**: Snyk Code requires Fully Qualified Names (FQNs). Partially Qualified Names (PQNs) are not supported.

## Step-by-step identification process

Follow these steps to identify your sanitizer's FQN:

### Step 1: Identify your programming language

Select your language to see the specific FQN format:

{% tabs %}
{% tab title="C#" %}
#### C# FQN format

**Format**: `Namespace.Class.Method`

**Important**: C# FQNs must include the `global::` prefix.

**Pattern**: `global::Namespace.Class.Method`

**How to identify your FQN**

1. Find the namespace declaration in your sanitizer's file.
2. Find the class name containing your sanitizer method.
3. Find the method name.
4. Combine them with dots: `global::Namespace.Class.Method`

**Example**

```csharp
namespace MyCorp.Security {
    public static class InputSanitizer {
        public static string Clean(string input) { 
            return input.Trim(); 
        }
    }
}
```

**FQN**: `global::MyCorp.Security.InputSanitizer.Clean`

**Common pitfalls**

* **Using `using` directives**: If your sanitizer is imported via a `using` directive (especially wildcard imports), the engine may not resolve the FQN correctly. Use explicit namespace references at the call site.
* **Missing `global::` prefix**: Always include `global::` at the start of your FQN.

**Instance methods**

For instance methods, use the class name followed by the method name:

```csharp
namespace MyCorp.Web {
    public class RequestHelper {
        public string FilterParams(string raw) { /*...*/ }
    }
}
```

**FQN**: `global::MyCorp.Web.RequestHelper.FilterParams`
{% endtab %}

{% tab title="Java" %}
#### Java FQN format

**Format**: `package.name.ClassName.methodName`

**How to identify your FQN**

1. Find the `package` declaration at the top of your file.
2. Find the class name containing your sanitizer method.
3. Find the method name.
4. Combine them with dots: `package.ClassName.methodName`

**Example**

```java
package com.bank.security;

public class XSSFilter {
    public static String stripTags(String unsafe) {
        return unsafe.replaceAll("<[^>]*>", "");
    }
}
```

**FQN**: `com.bank.security.XSSFilter.stripTags`

**Common pitfalls**

* **Wildcard imports**: Avoid `import com.example.security.*;`. The engine struggles to trace functions imported this way. Use explicit imports: `import com.example.security.Sanitizer;`
* **Static vs instance**: Both static and instance methods are supported, but ensure the class is explicitly typed at the call site.

**Instance methods**

```java
package com.bank.utils;

public class StringUtils {
    public String makeSafe(String input) { /*...*/ }
}
```

**FQN**: `com.bank.utils.StringUtils.makeSafe`
{% endtab %}

{% tab title="Kotlin" %}
#### Kotlin FQN format

**Format**: `package.name.ClassName.methodName` or `package.name.FileNameKt.functionName` (for top-level functions)

**How to identify your FQN**

**For class methods:**

1. Find the `package` declaration.
2. Find the class name.
3. Find the method name.
4. Combine: `package.ClassName.methodName`

**For top-level functions:**

1. Find the `package` declaration.
2. Find the file name (without extension).
3. Find the function name.
4. Combine: `package.FileNameKt.functionName`

**Example: Top-level function**

```kotlin
package com.app.helpers

// Defined directly in the file, not in a class
fun sanitizeInput(input: String): String {
    return input.trim()
}
```

**FQN**: `com.app.helpers.UtilsKt.sanitizeInput`

_(Note: Kotlin compiles top-level functions into a class named `FileName + Kt`)_

**Example: Class method**

```kotlin
package com.app.security
class Validator {
    fun clean(text: String): String { /*...*/ }
}
```

**FQN**: `com.app.security.Validator.clean`

**Common pitfalls**

* **Implicit typing**: Avoid `val x = Validator()`. Use explicit types: `val v: Validator = Validator()`
* **Wildcard imports**: Same as Java—avoid wildcard imports for sanitizer classes.
{% endtab %}

{% tab title="Python" %}
#### Python FQN format

**Format**: `module.submodule.function_name` or `module.ClassName.method_name`

**How to identify your FQN**

1. Determine the module path from your project root.
2. Replace directory separators (`/`) with dots (`.`).
3. Add the function or class name.
4. For class methods, add the class name before the method: `module.ClassName.method_name`

**Example: Module function**

```python
# File: project/security/text.py
def remove_bad_chars(user_input):
    return user_input.replace(";", "")
```

**FQN**: `project.security.text.remove_bad_chars`

**Example: Class method**

```python
# File: project/utils/cleaner.py
class InputScrubber:
    def scrub(self, data):
        return data.strip()
```

**FQN**: `project.utils.cleaner.InputScrubber.scrub`

**Common pitfalls**

* **Decorators**: You cannot define a sanitizer that is applied as a decorator (e.g., `@my_sanitizer`).
{% endtab %}

{% tab title="JavaScript / TypeScript" %}
#### JavaScript/TypeScript FQN format

**Format**: `module_name.export_name`

**How to identify your FQN**

**For named exports:**

1. Identify the module path from your project root or npm package name.
2. Add the exported function name.

**For npm packages:**

1. Use the package name as it appears in `package.json`.
2. Add the exported function name.

**Example: Named export**

```typescript
// File: security/index.ts
export function escapeSql(query: string): string {
    return query.replace(/'/g, "''");
}
```

**FQN**: `security.index.escapeSql`

**Example: NPM package**

```typescript
import { sanitizeHtml } from 'awesome-sanitizer';
```

**FQN**: `awesome-sanitizer.sanitizeHtml`

**Common pitfalls**

* **Relative imports**: The engine does not reliably resolve FQNs for functions imported using relative paths (e.g., `import { clean } from '../../utils'`). Use absolute module paths or npm package names instead.
* **Default exports**: Default exports are supported. Append `.default` to the module path (e.g., `security.index.default`).
{% endtab %}

{% tab title="Go" %}
#### Go FQN format

**Format**: `host.com/user/project/package.Function`

**How to identify your FQN**

1. Find the import path (e.g., `github.com/myuser/utils`).
2. Find the function name (must be exported—starts with capital letter).
3. Combine: Use the full import path followed by `.FunctionName`.

**Example: Public function**

```go
// File: pkg/input/clean.go
package input

func SanitizeString(s string) string {
    return strings.TrimSpace(s)
}
```

**Import Path**: `github.com/org/repo/pkg/input`\
**FQN**: `github.com/org/repo/pkg/input.SanitizeString`

_(Note the dots replacing slashes)_

**Example: Struct method**

```go
package security

type XSSBlocker struct {}

func (x *XSSBlocker) Block(s string) string { /*...*/ }
```

**FQN**: `github.com/org/repo/pkg/security.XSSBlocker.Block`

**Common pitfalls**

* **Factory returns**: Methods called on structs returned by interface factories often fail resolution. Use direct struct instantiation when possible.
* **Package path format**: If the standard format fails, try replacing `/` with `.` in the package path.
{% endtab %}

{% tab title="Apex" %}
#### Apex FQN format

**Format**: `ClassName.methodName` or `Namespace.ClassName.methodName`

**How to identify your FQN**

1. Find the Apex class name containing your sanitizer method.
2. Find the method name.
3. If the class is in a managed package namespace, prepend the namespace.
4. Combine with a dot: `ClassName.methodName` or `Namespace.ClassName.methodName`

**Example: Global class method**

```java
public class InputSanitizer {
    public static String clean(String input) {
        return input.replaceAll('<[^>]*>', '');
    }
}
```

**FQN**: `InputSanitizer.clean`

**Example: Managed package class method**

```java
// Deployed in managed package namespace MyCorp
public class SecurityUtils {
    public static String escapeHtml(String input) {
        return input;
    }
}
```

**FQN**: `MyCorp.SecurityUtils.escapeHtml`

**Common pitfalls**

* **Inner classes**: Include the outer class name:
  `OuterClass.InnerClass.methodName`
* **Validation**: Use the [Impact Testing API](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#post-groups-group_id-sast-rule_extensions-tests)
  to confirm the FQN resolves correctly in your org.
{% endtab %}

{% tab title="C/C++" %}
#### C/C++ FQN format

**Format**: `namespace::functionName` or `namespace::ClassName::methodName`

**How to identify your FQN**

1. Find the namespace declaration (if any).
2. Find the function or class method name.
3. Combine with `::` separators.

**Example: Namespace function**

```cpp
namespace mycorp {
namespace security {

std::string sanitizeInput(const std::string& input) {
    return input;
}

}  // namespace security
}  // namespace mycorp
```

**FQN**: `mycorp::security::sanitizeInput`

**Example: Class method**

```cpp
namespace mycorp {
namespace security {

class InputSanitizer {
public:
    static std::string clean(const std::string& input) { /*...*/ }
};

}  // namespace security
}  // namespace mycorp
```

**FQN**: `mycorp::security::InputSanitizer::clean`

**Common pitfalls**

* **Missing namespace**: If your function is in the global namespace, use
  `functionName` or `ClassName::methodName` only.
* **Header vs implementation**: Use the namespace and name from the
  declaration the engine resolves at the call site.
{% endtab %}

{% tab title="Dart" %}
#### Dart FQN format

**Format**: `package_name/path/file.dart.functionName` or
`package_name/path/file.dart.ClassName.methodName`

**How to identify your FQN**

1. Find the package name from `pubspec.yaml`.
2. Find the library file path relative to the `lib/` directory.
3. Add the function or class method name.

**Example: Top-level function**

```dart
// File: lib/security/sanitizer.dart
// Package: my_app

String removeBadChars(String input) {
  return input.replaceAll(';', '');
}
```

**FQN**: `my_app/security/sanitizer.dart.removeBadChars`

**Example: Class method**

```dart
// File: lib/security/sanitizer.dart

class InputScrubber {
  String scrub(String data) {
    return data.trim();
  }
}
```

**FQN**: `my_app/security/sanitizer.dart.InputScrubber.scrub`

**Common pitfalls**

* **Barrel exports**: Prefer the file that defines the sanitizer, not a
  re-export file.
* **Private methods**: Sanitizers must be public (no leading `_`).
{% endtab %}

{% tab title="Groovy" %}
#### Groovy FQN format

**Format**: `package.ClassName.methodName`

Groovy follows the same FQN pattern as Java.

**How to identify your FQN**

1. Find the `package` declaration at the top of your file.
2. Find the class name containing your sanitizer method.
3. Find the method name.
4. Combine them with dots: `package.ClassName.methodName`

**Example**

```groovy
package com.example.security

class XSSFilter {
    static String stripTags(String unsafe) {
        return unsafe.replaceAll('<[^>]*>', '')
    }
}
```

**FQN**: `com.example.security.XSSFilter.stripTags`

**Common pitfalls**

* **Dynamic dispatch**: Prefer explicit class references at the call site
  (for example, `XSSFilter.stripTags(input)`).
* **Scripts without packages**: Use `ScriptClassName.methodName` where
  `ScriptClassName` is the compiled script class name.
{% endtab %}

{% tab title="PHP" %}
#### PHP FQN format

**Format**: `Namespace\ClassName\methodName` or `methodName`

**How to identify your FQN**

1. Find the `namespace` declaration in your sanitizer's file.
2. Find the class name containing your sanitizer method.
3. Find the method name.
4. Combine with backslashes: `Namespace\ClassName\methodName`

For some call patterns—especially chained instance method calls—the engine
may resolve the sanitizer to the method name alone. Validate the FQN with
the [Impact Testing API](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/sast/rule_extensions/tests).

**Example 1: Static call on a namespaced class**

```php
<?php
namespace App\Security;

abstract class Sanitizer
{
    public static function sanitizeInput(string $input): string
    {
        return htmlspecialchars($input, ENT_QUOTES, 'UTF-8');
    }
}

// Call site: Sanitizer::sanitizeInput($userInput);
```

**FQN**: `App\Security\Sanitizer\sanitizeInput`

**Example 2: Chained instance method call**

```php
<?php
namespace App\Security;

class HtmlEscaper
{
    public function escape(string $input): string
    {
        return htmlspecialchars($input, ENT_QUOTES, 'UTF-8');
    }
}

// Call site: $escaper->escape($userInput);
```

**FQN**: `escape`

**Common pitfalls**

* **Abstract or parent classes**: Static calls on abstract classes resolve to
  the declaring class namespace path, not the concrete subclass.
* **Chained calls**: When a sanitizer is invoked through a chain
  (for example, `$object->escape($input)`), the FQN may be the method name
  only. Always confirm with an impact test.
* **Global functions**: Functions defined outside a namespace use the
  function name only (for example, `htmlspecialchars`).
{% endtab %}

{% tab title="Ruby" %}
#### Ruby FQN format

**Format**: `Module::Class.method_name` or `Module.method_name`

**How to identify your FQN**

1. Find the module or class path to your sanitizer.
2. Find the method name.
3. Combine modules with `::` and separate the method with `.`:
   `Module::Class.method_name`

**Example: Class method**

```ruby
module MyApp
  module Security
    class InputSanitizer
      def self.clean(input)
        input.gsub(/[<>]/, '')
      end
    end
  end
end

# Call site: MyApp::Security::InputSanitizer.clean(user_input)
```

**FQN**: `MyApp::Security::InputSanitizer.clean`

**Example: Module function**

```ruby
module MyApp
  module Security
    module Sanitizer
      def self.escape_html(input)
        CGI.escapeHTML(input)
      end
    end
  end
end
```

**FQN**: `MyApp::Security::Sanitizer.escape_html`

**Common pitfalls**

* **Mixins and `include`**: Prefer the module or class where the method is
  defined, not where it is mixed in.
* **Unique method names**: If the method name is unique in your project,
  the engine may resolve to the method name alone. Confirm with an impact
  test.
{% endtab %}

{% tab title="Rust" %}
#### Rust FQN format

**Format**: `crate::module::function_name` or
`crate::module::StructName::method_name`

**How to identify your FQN**

1. Find the crate name from `Cargo.toml`.
2. Find the module path to your sanitizer.
3. Add the function name, or the struct and method name for `impl` methods.

**Example: Public function**

```rust
// File: src/security/sanitize.rs
// Crate: my_app

pub fn sanitize_string(input: &str) -> String {
    input.trim().to_string()
}
```

**FQN**: `my_app::security::sanitize::sanitize_string`

**Example: Struct method**

```rust
pub struct InputSanitizer;

impl InputSanitizer {
    pub fn clean(input: &str) -> String {
        input.to_string()
    }
}
```

**FQN**: `my_app::security::sanitize::InputSanitizer::clean`

**Common pitfalls**

* **Re-exports**: Prefer the module where the function is defined.
* **Private functions**: Sanitizers must be `pub`.
{% endtab %}

{% tab title="Scala" %}
#### Scala FQN format

**Format**: `package.ClassName.methodName` or `package.ObjectName.methodName`

Scala follows the same FQN pattern as Java.

**How to identify your FQN**

1. Find the `package` declaration at the top of your file.
2. Find the class or object name containing your sanitizer method.
3. Find the method name.
4. Combine them with dots: `package.ClassName.methodName`

**Example: Class method**

```scala
package com.bank.security

class XSSFilter {
  def stripTags(unsafe: String): String = {
    unsafe.replaceAll("<[^>]*>", "")
  }
}
```

**FQN**: `com.bank.security.XSSFilter.stripTags`

**Example: Object method**

```scala
package com.bank.security

object SecurityUtils {
  def escapeHtml(input: String): String = input
}
```

**FQN**: `com.bank.security.SecurityUtils.escapeHtml`

**Common pitfalls**

* **Companion objects**: Use the object name directly:
  `package.MyClass$.methodName` only if the standard object name fails.
* **Wildcard imports**: Avoid wildcard imports for sanitizer classes.
{% endtab %}

{% tab title="Swift" %}
#### Swift FQN format

**Format**: `ModuleName.ClassName.methodName` or `ModuleName.functionName`

**How to identify your FQN**

1. Find the module name (typically your target or framework name).
2. Find the type or function name.
3. Combine with dots.

**Example: Struct method**

```swift
// Module: MyApp

struct InputSanitizer {
    static func clean(_ input: String) -> String {
        return input.trimmingCharacters(in: .whitespaces)
    }
}

// Call site: InputSanitizer.clean(userInput)
```

**FQN**: `MyApp.InputSanitizer.clean`

**Example: Top-level function**

```swift
// Module: MyApp

func sanitizeInput(_ input: String) -> String {
    return input
}
```

**FQN**: `MyApp.sanitizeInput`

**Common pitfalls**

* **Extensions**: Prefer the type that owns the method at the call site.
* **Framework modules**: For sanitizers in a framework target, use that
  target's module name as the prefix.
{% endtab %}

{% tab title="VB.NET" %}
#### VB.NET FQN format

**Format**: `global::Namespace.Class.Method`

VB.NET follows the same FQN pattern as C#.

**Important**: VB.NET FQNs must include the `global::` prefix.

**How to identify your FQN**

1. Find the namespace declaration in your sanitizer's file.
2. Find the class name containing your sanitizer method.
3. Find the method name.
4. Combine them with dots: `global::Namespace.Class.Method`

**Example**

```vb
Namespace MyCorp.Security
    Public Class InputSanitizer
        Public Shared Function Clean(input As String) As String
            Return input.Trim()
        End Function
    End Class
End Namespace
```

**FQN**: `global::MyCorp.Security.InputSanitizer.Clean`

**Common pitfalls**

* **Missing `global::` prefix**: Always include `global::` at the start
  of your FQN.
* **Modules**: For methods in a VB module, use
  `global::Namespace.ModuleName.Method`.
{% endtab %}
{% endtabs %}

### Step 2: Verify your function meets requirements

Before proceeding, ensure your sanitizer function:

* ✅ Is **publicly accessible** (public/exported)
* ✅ Has a **unique name** within your codebase
* ✅ Is **not imported via wildcard** (where applicable)
* ✅ Uses **explicit types** at call sites (for languages that support type inference)

### Step 3: Check your call site

The way your sanitizer is called affects FQN resolution:

**Supported patterns:**

* Static methods called with explicit class names
* Instance methods on explicitly typed objects
* Functions imported with absolute paths (not relative)

**Problematic patterns:**

* Functions imported via wildcard imports (`import pkg.*`)
* Functions called via relative imports (`import { func } from './utils'`)
* Implicitly typed variables (for languages with type inference)

### Step 4: Test your FQN

After identifying your FQN:

1. Create a draft Rule Extension with your identified FQN.
2. Run an [Impact Testing API](https://docs.snyk.io/developer-tools/snyk-api/reference/sastruleextensions#post-groups-group_id-sast-rule_extensions-tests) test on a
   project that uses your sanitizer. This is the recommended way to confirm
   the FQN resolves correctly and to check the expected reduction in findings before you
   publish the Rule Extension.
3. Verify that false positives are reduced in the impact test results.
4. If issues persist, review the troubleshooting section below.

For more information on permissions and usage, see
[FAQ & troubleshooting](faq-and-troubleshooting.md#rule-extensions).

## Quick reference: FQN formats by language

| Language              | Format                                                              | Example                                              |
| --------------------- | ------------------------------------------------------------------- | ---------------------------------------------------- |
| Apex                  | `ClassName.methodName` or `Namespace.ClassName.methodName`          | `InputSanitizer.clean`                               |
| C/C++                 | `namespace::functionName` or `namespace::ClassName::methodName`     | `mycorp::security::sanitizeInput`                    |
| C#                    | `global::Namespace.Class.Method`                                    | `global::MyCorp.Security.InputSanitizer.Clean`       |
| Dart                  | `package/path/file.dart.functionName`                               | `my_app/security/sanitizer.dart.removeBadChars`      |
| Go                    | `host.com/user/project/package.Function`                            | `github.com/org/repo/pkg/input.SanitizeString`       |
| Groovy                | `package.ClassName.methodName`                                      | `com.example.security.XSSFilter.stripTags`           |
| Java                  | `package.ClassName.methodName`                                      | `com.bank.security.XSSFilter.stripTags`              |
| JavaScript/TypeScript | `module_name.export_name`                                           | `security.index.escapeSql`                           |
| Kotlin                | `package.ClassName.methodName` or `package.FileNameKt.functionName` | `com.app.security.Validator.clean`                   |
| PHP                   | `Namespace\ClassName\methodName` or `methodName`                    | `App\Security\Sanitizer\sanitizeInput`               |
| Python                | `module.function_name` or `module.ClassName.method_name`            | `project.security.text.remove_bad_chars`             |
| Ruby                  | `Module::Class.method_name`                                         | `MyApp::Security::InputSanitizer.clean`              |
| Rust                  | `crate::module::function_name` or `crate::module::StructName::method_name` | `my_app::security::sanitize::sanitize_string`        |
| Scala                 | `package.ClassName.methodName`                                      | `com.bank.security.XSSFilter.stripTags`              |
| Swift                 | `ModuleName.ClassName.methodName`                                   | `MyApp.InputSanitizer.clean`                          |
| VB.NET                | `global::Namespace.Class.Method`                                    | `global::MyCorp.Security.InputSanitizer.Clean`       |

## Troubleshooting FQN issues

### My sanitizer isn't being recognized

**Check these common issues:**

1. **Missing `global::` prefix (C# and VB.NET)**: Ensure C# and VB.NET FQNs
   start with `global::`
2. **Wildcard imports**: Replace wildcard imports with explicit imports
3. **Relative imports (JS/TS)**: Use absolute module paths or npm package names
4. **Case sensitivity**: FQNs are case-sensitive—verify exact capitalization
5. **PHP namespace separators**: Use backslashes (`\`), not dots, in PHP FQNs
6. **Chained or dynamic calls (PHP, Ruby)**: The engine may resolve to the
   method name only—validate with the
   [Impact Testing API](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/sast/rule_extensions/tests)

### Still having issues?

If your sanitizer meets all requirements but still isn't recognized:

1. Verify the function is exported/public
2. Check that the function name matches exactly (case-sensitive)
3. Ensure the function is called with explicit types (where applicable)
4. Review the language-specific limitations in the tabs above
5. Run an [Impact Testing API](https://apidocs.snyk.io/?version=2026-03-25#post-/groups/-group_id-/sast/rule_extensions/tests) test to
   confirm the FQN and sanitization type
6. Contact your Snyk account team or [Snyk support](https://support.snyk.io/s/) for assistance

## Next steps

Once you've identified your FQN:

1. Read about [sanitization types](custom-sanitizers.md#types-of-sanitizers) to determine which type matches your function's behavior
2. Review the [supported rules](supported-rules.md) to select which rules your sanitizer applies to
3. Follow the [configuration guide](configure-rule-extensions.md) to create your Rule Extension
