---
description: >-
  What an AI-BOM scan detects in configuration and manifest files, without analyzing source code
nav_context: agnostic
---

# Non-code scanning

In addition to analyzing source code, an AI-BOM scan reads configuration and manifest files that declare AI components. Each finding appears on the **Inventory** page with the file and line where Evo found it.

## What Evo detects

| Asset | Detected from | Details captured |
| --- | --- | --- |
| MCP server | MCP configuration files for AI coding tools | Name, transport, endpoint URL, command, and arguments |
| Skill | `SKILL.md` files | Name, description, license, version, author, and allowed tools |
| Plugin | `.claude-plugin/plugin.json` | Name, description, author, and keywords |
| Hook | `hooks/hooks.json` | Event, matcher, hook types, and commands |
| Model | Configuration and manifest files | Model name, file, and line |

Evo records only these details. Other content, such as environment variable values and headers in MCP configuration files, is not stored.

A plugin includes the skills, MCP servers, and hooks in its directory.

## MCP configuration files

Evo detects MCP servers in project-committed configuration files at any depth in the repository:

| Client | File |
| --- | --- |
| Claude Code | `.mcp.json` |
| Cursor | `.cursor/mcp.json` |
| VS Code | `.vscode/mcp.json` |
| Windsurf | `.windsurf/mcp_config.json` |
| Kiro | `.kiro/settings/mcp.json` |
| Antigravity | `.agents/mcp_config.json` |
| Amazon Q | `.amazonq/mcp.json` |
| Gemini CLI | `.gemini/settings.json` |
| Claude Desktop | `claude_desktop_config.json` |
| Amp | `.amp/settings.json`, `.amp/settings.jsonc` |
| OpenCode | `opencode.json`, `opencode.jsonc` |
| Codex | `.codex/config.toml` |
| Other clients | Any `mcp.json` or `.mcp.json` |

Evo skips servers without a name.

## Skills, plugins, and hooks

* **Skills:** Evo detects `SKILL.md` files that start with YAML frontmatter containing a `description`. If there is no `name`, Evo uses the parent directory name. Evo also tracks the other files in the skill directory, such as scripts, so a change to any of them is treated as a change to the skill.
* **Plugins:** Evo detects Claude Code plugin manifests that include a name.
* **Hooks:** Evo reports one hook for each event entry, such as each `PreToolUse` entry, in `hooks/hooks.json`.

## Models in configuration files

Evo detects references to known AI models, such as `gpt-4o` or `meta-llama/Llama-3`, in these files:

* **Extensions:** `.yaml`, `.yml`, `.json`, `.toml`, `.ini`, `.conf`, `.xml`, `.properties`, `.tf`, `.tfvars`, `.env`, `.dockerfile`, and `.containerfile`.
* **Names:** `Dockerfile`, `Containerfile`, `Makefile`, `Justfile`, `Pipfile`, `Gemfile`, and dependency manifests such as `requirements.txt`, `package.json`, `pyproject.toml`, `go.mod`, `pom.xml`, `build.gradle`, `Cargo.toml`, and `composer.json`.

Evo does not scan Markdown, `.txt`, `.rst`, CSV, or other free-text files.

Model detection is heuristic. Evo uses the surrounding context, such as the key a name appears under, to filter out false matches, and reports these findings with low confidence.

{% hint style="info" %}
Evo includes models found in configuration files only if the scan also finds a model or model-loading function in your source code. A repository with model names in configuration but none in code produces no model assets.
{% endhint %}

## What Evo skips

* **Directories:** `.git`, `node_modules`, `vendor`, `.venv`, `venv`, `dist`, `build`, `target`, `out`, `.idea`, `.cache`, `tmp`, `temp`, `test`, `tests`, `.test`, `docs`, `doc`, `examples`, `example`, `fixtures`, `__fixtures__`, `mocks`, `__mocks__`, and `gems`. Evo does not scan anything inside them, including real configuration files.
* **Large files:** Evo skips most files over 50 KB. Configuration and dependency manifest files can be up to 1 MB. MCP server, plugin, and hook files can be up to 64 KB.
* **Malformed files:** Evo skips unreadable or malformed files without failing the scan.

## Limitations

* Evo cannot see per-user configuration, such as `~/.cursor/mcp.json`, because it scans repositories.
* Evo does not detect hooks defined in `.claude/settings.json`. It detects only `hooks/hooks.json`.
* Evo does not detect models in notebooks, Markdown, or free text.
* Evo does not detect models it does not recognize.
* Plugin detection supports the Claude Code format only.
