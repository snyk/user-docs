---
nav_context: agnostic
---

# Custom sanitizers

Taint flow analysis can miss your application's unique security controls. Defining custom sanitizers helps you get more accurate results by teaching the scanner about your specific data cleaning and validation methods, reducing false positives.

Use custom sanitizers in the following cases:

* In-house libraries: For your organization's own `Sanitize.clean()` functions.
* Third-party libraries: When a library you use safely handles data.
* Mapper functions: For simple, custom utility functions that clean or validate data.

### How custom sanitizers work

When you register a function as a custom sanitizer, Snyk Code treats the data it cleans or validates as safe. A taint flow that reaches a sink only after passing through your sanitizer is no longer reported, which removes the corresponding false positive. You tell Snyk Code *which* function to trust by its [fully qualified name (FQN)](identify-your-sanitizers-fqn.md), and *how* it cleans data by choosing a sanitizer type.

## Types of sanitizers

### Flow Through

_API `extension_type`: `flows_through_sanitizer`_

Sanitize all the data that flows through the function. The return value of the function is always sanitized, even if the input parameters are not sanitized. These sanitizers turn unsanitized data into sanitized data.

{% tabs %}
{% tab title="Java" %}
In this example, a sanitizer exists with the following FQN: `com.company.utils.SecurityUtils.escapeHtml`

{% code fullWidth="false" %}
```java
package com.company.utils;

public class SecurityUtils {
    public static String escapeHtml(String data) {
        if (data == null) return "";
        return data.replace("<", "&lt;").replace(">", "&gt;");
    }
}
```
{% endcode %}

It is used in this HttpServlet:

```java
import com.company.utils.SecurityUtils;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class CommentServlet extends HttpServlet {
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws IOException {
        var userComment = request.getParameter("comment");
        var safeComment = SecurityUtils.escapeHtml(userComment);
        
        response.getWriter().println("<html><body>" + safeComment + "</body></html>");
    }
}
```
{% endtab %}
{% endtabs %}

### If True

_API `extension_type`: `if_true_sanitizer`_

Arguments to these sanitizers are considered sanitized on ‌branches where the return value of the sanitizer is checked for trueness. These sanitizers are expected to be used in a condition.

{% tabs %}
{% tab title="Java" %}
In this example, a sanitizer exists with the following FQN `com.company.utils.SecurityUtils.isAlphaNumeric`

{% code fullWidth="false" %}
```java
package com.company.utils;

public class SecurityUtils {
    public static boolean isAlphaNumeric(String data) {
        return data != null && data.matches("^[a-zA-Z0-9]*$");
    }
}
```
{% endcode %}

It is used in this HttpServlet:

```java
import com.company.utils.SecurityUtils;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.IOException;

public class ProfileServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
        var username = request.getParameter("username");

        if (SecurityUtils.isAlphaNumeric(username)) {
            var filePath = "/var/user_profiles/" + username + "/avatar.png";
            var userAvatar = new File(filePath);
        } else {
            response.sendError(400, "Invalid username.");
        }
    }
}
```
{% endtab %}
{% endtabs %}

### If False

_API `extension_type`: `if_false_sanitizer`_

Arguments to these sanitizers are considered sanitized on ‌branches where the return value of the sanitizer is checked for falseness. These sanitizers are expected to be used in a condition.

{% tabs %}
{% tab title="Java" %}
In this example, a sanitizer exists with the following FQN: `com.company.utils.SecurityUtils.containsJavascriptProtocol`

{% code fullWidth="false" %}
```java
package com.company.utils;

public class SecurityUtils {
    public static boolean containsJavascriptProtocol(String data) {
        if (data == null) return false;
        return data.toLowerCase().trim().startsWith("javascript:");
    }
}
```
{% endcode %}

It is used in this HttpServlet:

```java
import com.company.utils.SecurityUtils;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class RedirectServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
        var redirectUrl = request.getParameter("url");

        if (!SecurityUtils.containsJavascriptProtocol(redirectUrl)) {
            response.sendRedirect(redirectUrl);
        } else {
            response.sendError(400, "Invalid redirect URL detected.");
        }
    }
}
```
{% endtab %}
{% endtabs %}

### Any Usage

_API `extension_type`: `any_usage_sanitizer`_

Arguments to these sanitizers are considered sanitized after the execution of the sanitizer. This form of sanitizer is either expected to mutate data that a reference is passing, or throw exceptions if non-sanitized data is passed.

Using the wrong sanitizer type can cause false negatives and hide incorrect usage of sanitizer functions in your code.

{% tabs %}
{% tab title="Java" %}
In this example, a sanitizer exists with the following FQN: `com.company.utils.SecurityUtils.validateFileId`

{% code fullWidth="false" %}
```java
package com.company.utils;

public class SecurityUtils {
    public static void validateFileId(String data) {
        if (data == null || !data.matches("^[a-zA-Z0-9_-]+$")) {
            throw new IllegalArgumentException("Invalid characters in File ID.");
        }
    }
}
```
{% endcode %}

It is used in this HttpServlet:

```java
import com.company.utils.SecurityUtils;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

public class FileAccessServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
        var fileId = request.getParameter("fileId");
        
        try {
            SecurityUtils.validateFileId(fileId);
            var filePath = "/var/data/uploads/" + fileId;
            var requestedFile = new File(filePath);
            
            response.setContentType("application/octet-stream");
            try (var in = new FileInputStream(requestedFile)) {
                in.transferTo(response.getOutputStream());
            }

        } catch (IllegalArgumentException e) {
            response.sendError(400, "Invalid File ID.");
        }
    }
}
```
{% endtab %}
{% endtabs %}
