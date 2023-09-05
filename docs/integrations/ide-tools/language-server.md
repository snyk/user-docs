---
description: Use this documentation to get started with the Language Server
---

# Snyk Language Server

Snyk offers IDE integrations that allow you to use the functionality of Snyk in your Integrated Development Environment or Editor. This page describes the Snyk Language Server that can provide diagnostics for any IDE or Editor that supports the [Language Server Protocol](https://microsoft.github.io/language-server-protocol/). For information about all of the IDE plugins and their use, see [Snyk for IDEs](https://docs.snyk.io/ide-tools) in the docs.

The Snyk Language Server scans for vulnerabilities, open source license issues, code quality, and infrastructure misconfigurations and returns results with security issues categorized by issue type and severity.

For open source, you receive automated, algorithm-based fix suggestions for both direct and transitive dependencies.

Snyk Language Server scans for the following types of issues:

* [**Open Source Security**](https://snyk.io/product/open-source-security-management/) - security vulnerabilities and license issues in both the direct and in-direct (transitive) open-source dependencies pulled into the Snyk Project. See also the [Open Source docs](https://docs.snyk.io/products/snyk-open-source).
* [**Code Security**](https://snyk.io/product/snyk-code/) and [**Code Quality**](https://snyk.io/product/snyk-code/) - security vulnerabilities and quality issues in your code. See also the [Snyk Code docs](https://docs.snyk.io/products/snyk-code).
* [**Infrastructure as Code (IaC) Security**](https://snyk.io/product/infrastructure-as-code-security/) - configuration issues in your IaC templates: Terraform, Kubernetes, CloudFormation, and Azure Resource Manager. See also the [Snyk Infrastructure as Code docs](https://docs.snyk.io/products/snyk-infrastructure-as-code).

After you have installed and configured the Language Server, every time you run it, open a file, or save, Snyk scans the manifest files, proprietary code, and configuration files in your project. Snyk delivers actionable vulnerability, license, code quality, or misconfiguration issue details and displays the results natively within the LSP supporting Editor or IDE.

This page explains supported environments, support, and giving feedback and provides installation instructions.

## Supported operating systems and architecture

You can use the Language Server in the following environments:

* Linux: AMD64 and ARM64
* Linux Alpine: 386 and AMD64
* Windows: 386, AMD64, ARM via 386 compatibility
* MacOS: AMD64 and ARM64

## Where you can download the Language Server

Currently, Snyk Language Server is automatically downloaded only when you use the Visual Studio Code (VS Code) and Eclipse plugins. Language Server can also be downloaded manually; the following shell script shows how to do that.

{% code title="getLanguageServer.sh" lineNumbers="true" %}
```bash
#!/bin/bash
# This file allows you to download the latest language server, which is helpful for integration into non-managed Editors and IDEs.
# Currently, these might be any editor that does not have a downloader built by Snyk and thus needs to download
# and update the language server regularly, and this script allows this for system administrators and users.
# Snyk recommends always using the latest version of the language server.

set -e
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')
if [[ $ARCH == "x86_64" ]]; then
  ARCH="amd64"
fi
if [[ $ARCH == "aarch64" ]]; then
  ARCH="arm64"
fi
PROTOCOL_VERSION=3
VERSION=$(curl https://static.snyk.io/snyk-ls/$PROTOCOL_VERSION/metadata.json | jq .version | sed -e s/\"//g)
wget -O /usr/local/bin/snyk-ls "https://static.snyk.io/snyk-ls/$PROTOCOL_VERSION/snyk-ls_${VERSION}_${OS}_${ARCH}"
```
{% endcode %}

The PROTOCOL\_VERSION currently is 3, but may increase with ongoing development.

## Configuration of Snyk Language Server

### Snyk LSP command line flags

`-c <FILE>` allows specifying a config file to load before all others

`-l <LOGLEVEL>` allows specifying the log level (`trace`, `debug`, `info`, `warn`, `error`, `fatal`). The default log level is `info`

`-o <FORMAT>` allows specifying the output format (`md` or `html`) for issues

`-f <FILE>` allows specifying a log file instead of logging to the console

`-licenses` displays the [licenses](https://github.com/snyk/snyk-ls/tree/main/licenses) used by Language Server

### **LSP initialization options**

As part of the [Initialize message](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize) within `initializationOptions?: LSPAny;` Snyk supports the following settings:

```json
{
  "activateSnykOpenSource": "true", // Enables Snyk Open Source - defaults to true
  "activateSnykCode": "false", // Enables Snyk Code, if enabled for your organization - defaults to false
  "activateSnykIac":  "true", // Enables Infrastructure as Code - defaults to true
  "insecure": "false", // Allows custom CAs (Certification Authorities)
  "endpoint":  "https://example.com", // Snyk API Endpoint required for non-default multi-tenant and single-tenant setups
  "additionalParams": "--all-projects", // Any extra params for the Snyk CLI, separated by spaces
  "additionalEnv":  "MAVEN_OPTS=-Djava.awt.headless=true;FOO=BAR", // Additional environment variables, separated by semicolons
  "path": "/usr/local/bin", // Adds to the system path used by the CLI
  "sendErrorReports":  "true", // Whether or not to report errors to Snyk - defaults to true
  "organization": "a string", // The name of your organization, e.g. the output of: curl -H "Authorization: token $(snyk config get api)"  https://snyk.io/api/cli-config/settings/sast | jq .org
  "enableTelemetry":  "true", // Whether or not user analytics can be tracked
  "manageBinariesAutomatically": "true", // Whether or not CLI/LS binaries will be downloaded & updated automatically
  "cliPath":  "/a/patch/snyk-cli", // The path where the CLI can be found, or where it should be downloaded to
  "token":  "secret-token", // The Snyk token, e.g.: snyk config get api
  "automaticAuthentication": "true", // Whether or not LS will automatically authenticate on scan start (default: true)
  "enableTrustedFoldersFeature": "true", // Whether or not LS will prompt to trust a folder (default: true)
  "trustedFolders": ["/a/trusted/path", "/another/trusted/path"], // An array of folder that should be trusted
}
```

## **Authentication for Snyk Language Server**

When Snyk Language Server starts, it checks for a token in the initializationOption `token`. If a token is not there, Snyk Language Server tries to retrieve and authenticate using the Snyk CLI. If the CLI is not authenticated either, Snyk Language Server opens a browser window to authenticate. After successful authentication in the web browser, Snyk Language Server automatically retrieves the Snyk authentication token from the CLI.

## **Environment variables for Snyk Language Server**

Snyk Language Server and Snyk CLI support and need certain environment variables to function:

1. `HTTP_PROXY`, `HTTPS_PROXY` and `NO_PROXY` to define the http proxy to be used
2. `JAVA_HOME` to analyze Java JVM-based projects via Snyk CLI
3. `PATH` to find `maven` when analyzing Maven projects, to find `python` and so on

## **Auto-configuration of environment variables for Snyk Language Server**

To automatically add these variables to the environment, Snyk Language Server searches for the following files, with the order determining precedence. If the executable is not called from an already configured environment (for example, via `zsh -i -c 'snyk-ls'`), you can also specify the config file with the `-c` command line flag for setting the required variables. Snyk Language Server reads the following files in the given precedence and order, not overwriting the already loaded variables.

```
given config file via -c flag
<working-dir>/.snyk.env
$HOME/.snyk.env
```

Any lines that contain an environment variable in the format `VARIABLENAME=VARIABLEVALUE` are added automatically to the environment if not already there. This adheres to the `dotenv` format. In the case of `.profile`, `.zshrc`and so on, if a variable is directly exported, for example, via `export VARIABLENAME=VARIABLEVALUE`, it is not loaded. The export would need to be split of and be in its own line, for example

```bash
VARIABLENAME=VARIABLEVALUE
export VARIABLENAME
```

The PATH variable is treated differently frrom all other variables, as it is an aggregate of all PATH variables found in the files and in the environment. Also, the current working directory `.` is automatically added to the path, so a download of the Snyk CLI into the current working directory by an LSP client would yield a found Snyk CLI for the Language Server.

In addition to configuring variables via config files, Snyk Language Server adds the following directories to the path on Linux and macOS:

* /bin
* $HOME/bin
* /usr/local/bin
* $JAVA\_HOME/bin

If no JAVA\_HOME is set, Snyk Language Server automatically searches for a java executable first in `path`, then in the following directories, and adds the parent directory of its parent as JAVA\_HOME. The following directories are recursively searched:

* /usr/lib
* /usr/java
* /opt
* /Library
* $HOME/.sdkman
* C:\Program Files
* C:\Program Files (x86)

The same directories are searched for a Maven executable and the parent directory is added to the path.

## **Snyk CLI**

To find the automatically managed Snyk CLI, the [XDG Data Home](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html#variables) and `PATH` path are automatically scanned for the OS-dependent file, for example, `snyk-macos` on macOS, `snyk-linux` on Linux and `snyk-win.exe` on Windows, and the first path where it is found is added to the environment. It is later used for all functionality that depends on the CLI. The path to the CLI can also be set manually using the `cliPath` initialization option.

## Folder trust

As part of examining the codebase for vulnerabilities, Snyk may automatically execute code on your computer to obtain additional data for analysis. This includes invoking the package manager (for example, pip, Gradle, Maven, Yarn, npm, and so on) to get dependency information for Snyk Open Source. Invoking these programs on untrusted code that has malicious configurations may expose your system to malicious code execution and exploits.

To safeguard against using the Language Server on untrusted folders, the Snyk Language Server asks for folder trust before running scans against these folders. When in doubt, do not grant trust.

The trust feature is enabled by default. When a folder is trusted, all sub-folders are also trusted. After a folder is trusted, Snyk Language Server notifies the Language Server Client with the custom `$/snyk.addTrustedFolders` notification, which contains a list of currently trusted folder paths. Based on this, a client can then implement logic to intercept this notification and persist the decision and trust in the IDE or Editor storage mechanism.

Trust dialogs can be disabled by setting `enableTrustedFoldersFeature` to `false` in the initialization options. This disables all trust prompts and checks.

An initial set of trusted folders can be provided by setting `trustedFolders` to an array of paths in the `initializationOptions`. These folders will be trusted on startup and will not prompt the user to trust them.

## Example configuration for Sublime Text

```
// Settings in here override those in "LSP/LSP.sublime-settings"
{
  // Show permanent language server status in the status bar.
  "show_view_status": true,

  // Language server configurations
  "clients": {
    "snyk": {
      // enable this configuration
      "enabled": true,
      "command": [
        "/usr/local/bin/snyk-ls", // path to the downloaded binary
        "-l", // define log level
        "info", // level info
        "-f", // file logging
        "/path/to/log/dir/snyk-ls-sublime.log" // log file
      ],
      // the selector that selects which type of buffers this language server attaches to
      "selector": "source", // all file types
      "initializationOptions": {
        "activateSnykCode":"true", // Enable Snyk Code
        "token": "xxx" // Set your Snyk Token to not authenticate on every start
      }
    },
  },
}
```

After opening a supported file, the Language Server should be started by Sublime Text and findings will be highlighted.

![Snyk Open Source findings displayed in Sublime Text](<../../.gitbook/assets/image (166) (1) (1) (1) (1) (1) (1) (2) (3).png>)

![Snyk Code findings displayed in Sublime Text](<../../.gitbook/assets/image (466) (1).png>)

## Example configuration for Neovim

The setup is as follows:

```
mkdir -p ~/.config/nvim
touch init.lua
```

The following is an example `init.lua` script:

```
local on_windows = vim.loop.os_uname().version:match 'Windows'

local function join_paths(...)
    local path_sep = on_windows and '\\' or '/'
    local result = table.concat({ ... }, path_sep)
    return result
end

vim.cmd [[set runtimepath=$VIMRUNTIME]]

local temp_dir = vim.loop.os_getenv 'TEMP' or '/tmp'

vim.cmd('set packpath=' .. join_paths(temp_dir, 'nvim', 'site'))

local package_root = join_paths(temp_dir, 'nvim', 'site', 'pack')
local install_path = join_paths(package_root, 'packer', 'start', 'packer.nvim')
local compile_path = join_paths(install_path, 'plugin', 'packer_compiled.lua')

local function load_plugins()
    require('packer').startup {
        {
            'wbthomason/packer.nvim',
            'neovim/nvim-lspconfig',
        },
        config = {
            package_root = package_root,
            compile_path = compile_path,
        },
    }
end

_G.load_config = function()
    vim.lsp.set_log_level 'trace'
    if vim.fn.has 'nvim-0.5.1' == 1 then
        require('vim.lsp.log').set_format_func(vim.inspect)
    end
    local nvim_lsp = require 'lspconfig'
    local on_attach = function(_, bufnr)
        local function buf_set_option(...)
            vim.api.nvim_buf_set_option(bufnr, ...)
        end

        buf_set_option('omnifunc', 'v:lua.vim.lsp.omnifunc')

        -- Mappings.
        local opts = { buffer = bufnr, noremap = true, silent = true }
        vim.keymap.set('n', 'gD', vim.lsp.buf.declaration, opts)
        vim.keymap.set('n', 'gd', vim.lsp.buf.definition, opts)
        vim.keymap.set('n', 'K', vim.lsp.buf.hover, opts)
        vim.keymap.set('n', 'gi', vim.lsp.buf.implementation, opts)
        vim.keymap.set('n', '<C-k>', vim.lsp.buf.signature_help, opts)
        vim.keymap.set('n', '<space>wa', vim.lsp.buf.add_workspace_folder, opts)
        vim.keymap.set('n', '<space>wr', vim.lsp.buf.remove_workspace_folder, opts)
        vim.keymap.set('n', '<space>wl', function()
            print(vim.inspect(vim.lsp.buf.list_workspace_folders()))
        end, opts)
        vim.keymap.set('n', '<space>D', vim.lsp.buf.type_definition, opts)
        vim.keymap.set('n', '<space>rn', vim.lsp.buf.rename, opts)
        vim.keymap.set('n', 'gr', vim.lsp.buf.references, opts)
        vim.keymap.set('n', '<space>e', vim.diagnostic.open_float, opts)
        vim.keymap.set('n', '[d', vim.diagnostic.goto_prev, opts)
        vim.keymap.set('n', ']d', vim.diagnostic.goto_next, opts)
        vim.keymap.set('n', '<space>q', vim.diagnostic.setloclist, opts)
    end

    local lspconfig = require('lspconfig')
    local configs = require 'lspconfig.configs'

    if not configs.snyk then
        configs.snyk = {
            default_config = {
                cmd = {'/usr/local/bin/snyk-ls','-f','/path/to/log/snyk-ls-vim.log'},
                root_dir = function(name)
                    return lspconfig.util.find_git_ancestor(name) or vim.loop.os_homedir()
                end,
                init_options = {
                    activateSnykCode = "true",
                }
            };
        }
    end
    lspconfig.snyk.setup {
      on_attach = on_attach
    }

    print [[You can find your log at $HOME/.cache/nvim/lsp.log. Please paste in a github issue under a details tag as described in the issue template.]]
end

if vim.fn.isdirectory(install_path) == 0 then
    vim.fn.system { 'git', 'clone', 'https://github.com/wbthomason/packer.nvim', install_path }
    load_plugins()
    require('packer').sync()
    vim.cmd [[autocmd User PackerComplete ++once lua load_config()]]
else
    load_plugins()
    require('packer').sync()
    _G.load_config()
end
```

![Snyk Code findings displayed in Neovim](<../../.gitbook/assets/image (219) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png>)

## Support for Snyk Language Server

If you need help, submit a [request](https://support.snyk.io/hc/en-us/requests/new) to Snyk Support.
