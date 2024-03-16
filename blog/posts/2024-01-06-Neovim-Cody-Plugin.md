---
templateKey: blog-post
title: "Neovim + Sourcegraph Cody Plugin Integration"
description: "Installing and Setting up Sourcegraph Plugin and Cody into Neovim. Enabling code generations, chat and sourcegraph search feature in Neovim."
date: 2024-01-06 20:15:00
status: published
slug: neovim-sourcegraph-cody
tags: ['vim']
image_url: https://meetgor-cdn.pages.dev/neovim-sourcegraph-cody.png
---

## Introduction

Have you ever used Sourcegraph's Cody? It is a great tool for developers, it is not just another LLM, it is tailored specifically for developers. Cody has some good features that allow parsing of context for the prompt in a smarter way.

### What is Sourcegraph's Cody

Cody is an AI assistant for developers that understands code context and can generate code. It goes beyond just answering questions - it can write code for you.

The major features of Cody from the rest of the LLMs or Chatbots are:

* Cody understands your code context - it reads your open files, repositories, etc. So it can answer questions specifically about your codebase, not just general programming questions.

* Cody can explain sections of code to you in plain English. This helps ramp up on unfamiliar code bases.

* Cody integrates into popular editors like VS Code, IntelliJ, Neovim, and others for frictionless use while coding.

For more insights, check out the blog [all you need is Cody](https://about.sourcegraph.com/blog/all-you-need-is-cody). This is a great article about what and how Cody is tailored specifically to assist developers.

## Prerequisites

To setup sourcegraph on neovim, you will require the following setup:

* Neovim 0.9 or above

* Node.js &gt;= 18.17 (LTS)

* Cargo (Rust) (optional)

To install neovim latest/nightly release, you can follow the [INSTALL](https://github.com/neovim/neovim/blob/master/INSTALL.md) or [BUILD](https://github.com/neovim/neovim/blob/master/BUILD.md) documentation of the neovim project.

Cargo is optional, as the plugin will install the binaries itself, however, if you prefer to have cargo, just install it in case something goes wrong.

## Installing sg.nvim

There is a specific plugin for neovim for interacting with the Sourcegraph products, and Cody is one of them. The [sg.nvim](https://github.com/sourcegraph/sg.nvim) is a plugin for integrating sourcegraph search, Cody, and other features provided by soucegraph.

#### Using packer.nvim

```lua
use { 'sourcegraph/sg.nvim', run = 'nvim -l build/init.lua' }
use { 'nvim-lua/plenary.nvim' }
```

Source your lua file where you have configured all the plugins and then enter the command `:PackerInstall` or `:PackerSync` to install the plugin.

#### Using vim-plug

If you are using vim-plug as the plugin manager, you can add the plugin to the configuration as below:

```plaintext
Plug 'sourcegraph/sg.nvim', { 'do': 'nvim -l build/init.lua' }

" Required for various utilities
Plug 'nvim-lua/plenary.nvim'

" Required if you want to use some of the search functionality
Plug 'nvim-telescope/telescope.nvim'
```

You can source the file and run the command `:PlugInstall` to install the plugin.

#### Using Lazy.nvim

If you are using Lazy.nvim as the plugin manager, you can add the plugin to the Configuration as below:

```lua
return {
  {
    "sourcegraph/sg.nvim",
    dependencies = { "nvim-lua/plenary.nvim", "nvim-telescope/telescope.nvim" },

    -- If you have a recent version of lazy.nvim, you don't need to add this!
    build = "nvim -l build/init.lua",
  },
}
```

You can source the file and run the command `:Lazy install` to install the plugin.

For other installation instructions, you can refer to the [README](https://github.com/sourcegraph/sg.nvim?tab=readme-ov-file#install) of sg.nvim.

### Installing Binaries and Building the Plugin

After the plugin is installed, you can move into the building and setup process of the sourcegraph Cody plugin.

To install the binaries which are the dependencies of the plugin, you can run the command `:SourcegraphDownloadBinaries` which will force downloading the binaries, making sure that the plugin is properly installed.

![SourcegraphDownloadBinaries Command Output](https://meetgor-cdn.pages.dev/sg-nvim-build.png)

To build the plugin, you can simply run the command from within neovim as `:SourcegraphBuild`, this will rebuild the Sourcegraph rust crates and its dependencies (which might have failed during installation).

### Sourcegraph Authentication

You need to now authenticate to your Sourcegraph account to use the sourcegraph features such as search and Cody.

You can do that by running the command `:SourcegraphLogin` in neovim. This will require two inputs, the sourcegraph endpoint and the access token. If you are using sourcegraph cloud and not a self-hosted sourcegraph, you do not need to change the endpoint, just press enter and move ahead. This will redirect you to the browser for authentication and creating an access token. Log in with your credentials to sourcegraph and copy the access token.

This will prompt you back to the neovim interface to provide the access token. Paste the access token there and you will be good to go.

### Health Check

Once the plugins are installed then you can check the plugin is correctly downloaded by checking the health with the `:checkhealth sg` command.

Below is the health check report on the sourcegraph plugin in neovim.

```plaintext
sg: require("sg.health").check()

sg.nvim report ~
- Machine: x86_64, sysname: Linux
- OK Valid nvim version: table: 0x7ffa0b7bce38
- OK Found `cargo` (cargo 1.70.0) is executable
-     Use `:SourcegraphDownloadBinaries` to avoid building locally.
- OK Found `sg-nvim-agent`: "/home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/sg-nvim-agent"
- OK Found `node` (config.node_executable) is executable.
  Version: '20.10.0'
- OK Found `cody-agent`: /home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/cody-agent.js
- OK   Authentication setup correctly
- OK     endpoint set to: https://sourcegraph.com
- OK Found correct binary versions: "1.0.5" = "1.0.5"
- OK   Sourcegraph Connection info: {
  access_token_set = true,
  endpoint = "https://sourcegraph.com",
  sg_nvim_version = "1.0.5",
  sourcegraph_version = {
		  build = "256174_2023-12-30_5.2-dbb20677711c",
		  product = "256174_2023-12-30_5.2-dbb20677711c"
  }
  }
- To manage your Cody Account, navigate to: https://sourcegraph.com/cody/manage
- OK Cody Account Information: {
	  chat_limit = 20,
	  chat_usage = 53,
	  code_limit = 500,
	  code_usage = 0,
	  cody_pro_enabled = false,
	  username = "Mr-Destructive"
  }
- OK sg.nvim is ready to run
```

At this point, the sourcegraph plugin is ready to be used. However, we need to set up the plugin in neovim with the default configurations.

### Configuration

In your lua setup files, you can set the plugin like this:

```lua
require("sg").setup()
```

Source the lua file and restart neovim, this should properly make sourcegraph commands available in the editor. After these steps, Cody is right into neovim.

## Usage

To use the plugin, there are multiple commands available within the editor, the complete list of them is given below:

```plaintext
SourcegraphBuild
SourcegraphClear
SourcegraphDownloadBinaries
SourcegraphInfo
SourcegraphLink
SourcegraphLogin
SourcegraphSearch

CodyAsk
CodyChat
CodyDo
CodyRestart
CodyTask
CodyTaskAccept
CodyTaskNext
CodyTaskPrev
CodyTaskView
CodyToggle
```

You can get more info about these commands with the `help :commandname` command. The command are however self explanatory and simple to use.

### Cody Ask

You can quickly parse a prompt as a string to the `:CodyAsk` command as `:CodyAsk "what is neovim by the way?"` and you will get Cody's response in the side vertical split.

### Cody Chat

You can start Cody chats from neovim with the command `:CodyChat`. This will open a vertical split with the Cody chat interface, the bottom split has the user input prompt and the upper window will have the generated Cody response. You can enter the prompt in the bottom buffer get into normal mode and hit enter to send the prompt to generate a response from Cody.

![Sourcegraph Cody Chat Interface](https://meetgor-cdn.pages.dev/sg-nvim-cody-chat.png)

## Conclusion

Sourcegraph Cody is a great tool for getting quick solutions to trivial as well as specific problems to the current file/package or project. The context parsing of Cody makes it valuable for developers to ask for specific questions and it answers them really in a straightforward way, without the developer needing to parse the context for the prompt. Thank you for reading. Happy Coding :)
