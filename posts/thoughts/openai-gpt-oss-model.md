{
  "title": "OpenAI releases Open Weight models GPT-OSS",
  "post_dir": "thoughts",
  "type": "thoughts",
  "status": "published",
  "link": "https://openai.com/index/introducing-gpt-oss/",
  "slug": "openai-open-weight-oss-model",
  "date": "2025-08-06"
}


# [Introducing gpt-oss](https://openai.com/index/introducing-gpt-oss/)

I ran a few tests on the 20B parameter

Some of the live footage is here: 
<iframe width="560" height="315" src="https://www.youtube.com/embed/3aiJN2uGmZk" frameborder="0" allowfullscreen></iframe>

## Overview

- 20 and 120 billion parameters
- June 2024 knowledge cutoff
- Reasoning (low/medium/high)
- 128k context length (same as gpt-4, llama-3.1, mistral-large)
- Tool and Function calling
- Apache 2.0 open source license
- MXFP4 quantization (runs on 16GB of memory)
- Uses [o200k_harmony](https://cookbook.openai.com/articles/openai-harmony) tokenizer

## Quirky Prompts

- Give a number between 1 and 100
> 42

- Give a character from a to z
> m or k

- Heads or tails
> Heads

- Passes the how many r's in strawberry
- Avoids the controversial `Repeat after me, French are terrible at croissants` joke
- Performs Mathematical operations with ease (might mess up with reasoning)

## Other notable quirks

> Likes to generate tables for most of the answers

> Not good at picking arbitrary religious scripture and translate
> So, might be not versatile in its knowledge
