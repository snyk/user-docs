# Example configurations for Snyk Language Server

## Example configuration for Sublime Text

```json5
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
        "/usr/local/bin/snyk", // path to the downloaded CLI binary
        "language-server", // start CLI in language server mode
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

![Snyk Open Source findings displayed in Sublime Text](<../../../.gitbook/assets/image (109) (1).png>)

![Snyk Code findings displayed in Sublime Text](<../../../.gitbook/assets/image (56).png>)

## Example configuration for Neovim

{% hint style="info" %}
If the project root directory cannot be obtained from Git information, this script tries to set the root directory to your home directory by `vim.loop.os_homedir().` Depending on the contents of your home directory, this may consume a large amount of memory. If this happens, set an alternative root directory for your environment.
{% endhint %}

The set up is as follows:

```
mkdir -p ~/.config/nvim
touch init.lua
```

The following is an example `init.lua` script:

```lua
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
                cmd = {'/usr/local/bin/snyk', 'language-server', '-f','/path/to/log/snyk-ls-vim.log'},
                root_dir = function(name)
                    return lspconfig.util.find_git_ancestor(name) or vim.loop.os_homedir()
                end,
                init_options = {
                    activateSnykCode = "true",
                    token = "xxx"
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

![Snyk Code findings displayed in Neovim](<../../../.gitbook/assets/image (12).png>)
