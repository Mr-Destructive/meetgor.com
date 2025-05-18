{"author":"meet","date":"2025-05-03","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #40","type":"newsletter"}

## Week #40

What a week, I knew the silence was just temporary, it would come at once, thankfully it was not too much releases to keep track of. There was, however, one release of a model (or a family of models) from Alibaba Cloud with the Qwen suite of models, that overthrew the Meta AI LLama Models from the SOTA standard for open source models.

This week, I explored a bit about LLMs and how to operate with them effectively, thanks to the prompt engineering masterclass from Anthropic. I created a bunch of things, leaving aside my fear and procrastination. This week, I was a bit self-critical of myself but was able to break through the chains and took some hard decisions and executed, the results? Not sure, maybe good, might not work. But that is not in my hands; whatever was, is done. Time to move ahead and explore this AI revolution firsthand.

### Quote of the week

> "Instead of a celebration of everything you know, an anti-library is an ode to everything you want to explore."
> 
> — Anne-Laure Le Cunff

I want to do this and this, nope, that might be output oriented

Let’s try to explore this and this, what might we need to do, what we need to go through to feel this?  
That is the question that is answered in a process-oriented way.

I want to change this, and I have been reading the book, Tiny Experiments, so far it has changed quite a few things about my procrastination behaviour and helped me take action over being scared or fretting over the perfectionism bug.

---

## Created

This week, I started a different approach in creating, just publishing at the end of the day.

### Built

* Published the [meta-ai-api tool-call](https://github.com/Mr-Destructive/meta_ai_api_tool_call) python package
    
* [Agentic Calculator](https://agentic-calculator.vercel.app/): LLM with math tools using Pydnatic AI and Vercel
    
    * LLM: Meta AI Llama 4 (on API Wrapper with tool calls)
        
    * Agent Framework: Pydantic AI
        
    * Deployment: Vercel Lambda Functions
        
    

### Wrote

* Outlined a couple of articles (will publish and reveal soon)
    
* Thought: [ChatGPT on WhatsApp](https://meetgor.bearblog.dev/chatgpt-is-online-on-whatsapp/) → might redirect in some time to → [meetgor.com](https://meetgor.com/thoughts/chatgpt-is-online-on-whatsapp/)
    

### Recorded

* 2 YouTube Shorts on [ChatGPT](https://www.youtube.com/shorts/R0FA-GFNgcM) and [Perplexity](https://www.youtube.com/shorts/V4WM1SmxrEI) released as WhatsApp Chatbots
    
* Livestreamed
    
    * [Creating the Meta AI API wrapper with a tool call and a Pydantic AI model wrapper for the Meta AI API](https://www.youtube.com/live/tBRQSlpgEUg)
        
    * [Building an Agentic Calculator with Pydantic, Appwrite (failed, then moved to Vercel), Meta AI API wrapper with tool calls](https://www.youtube.com/live/yA5d-R6O8h0)
        

## Read

* [Horseless Carriages](https://koomen.dev/essays/horseless-carriages/) :
    
    * This was a great read. People making AI-based applications are too naive to understand the actual technology behind LLM and how the approach needs to be changed.
        
    * Right now, it is the assumption that the developer is attached to the system prompt; however, LLMs at this stage need customisation from the user perspective and ndo ot have the dependency on the developer.
        
    * The shift in developer-user responsibility is quite unnoticed, and this article highlights that effectively.
        
* [Be Kind](https://boz.com/articles/be-kind) :
    
    * As a developer, we need to be kind, not just yes, thank you but really understand the person from the other end and be humble about his/her situation and feelings.
        
    * That might be too realistic for a developer to do (right? We are nerds, I think), but having that kindness and insight about the person we are interacting with sets us apart as an effective engineer.
        
* [Habits I recommend to a software developer](https://www.saiyangrowthletter.com/p/habits-i-recommend-to-succeed-as)
    
    * Reading a book
        
    * Build projects (keep building)
        
    * Write to reflect the learnings
        
    * Learn and build in public
        
* [Getting started with Pulumi](https://awsfundamentals.com/blog/pulumi-getting-started)
    
    * I always wanted to understand Infrastructure and code, but no article helps me understand what it actually solves, no one has yet shown the problem before the solution. I might find and write it myself one day (day one?).
        
* [Why can’t I be technical](https://www.fightforthehuman.com/why-i-cannot-be-technical)?
    
    * I didn’t completely read this, but this makes an interesting point, that we need to understand the opposite end of the thinking on the thing we are working on, kind of wired but that makes sense now.
        
    

## Watched

* [Lessons on AI Agents from Claude playing Pokémon](https://youtu.be/CXhYDOvgpuU)
    
    * Wow, agents are already playing games, not exactly, but quite fair, I would say. Anthropic is really a lab, like they are researching LLM behaviors through and through; they are technical scientists.
        
    * Claude plays Pokémon Red:
        
        * Send a screenshot of the current state
            
        * Describe the game mechanics
            
        * Ask for the action
            
        * Iterate
            
    * It’s quite a fascinating experiment. Maybe we can try with different types of games with LLMs. They tried a Pokémon-like game, because that is a very user-paced game, not a very rapid pace, or live-like games. Very smooth transitions and turn based game.
        

* [AI Prompt Engineering: Deep dive](https://youtu.be/T9aRN5JkmL8)
    
    * This is a masterclass in prompt engineering. Must watch
        
    * Anthropic really cares about the craft and art of LLMs, they really understand what to think when interfacing with an LLM.
        
    * First principle thinking
        

* [Advice for juniors on manager-engineer relation and Theo’s experience (must watch)](https://youtu.be/3VuM1GCadt4)
    
    * TLDR: Don’t try to do a job that you are not asked, if the environment is pulling each other down.
        
    * It doesn’t mean you are not following your curiosity or doing the things you are excited to do, rather than knowing when to and when not to.
        

* [AI Coding is not enough we need Agnetic Coding:](https://youtu.be/2TIXl2rlA6Q)
    
    * This is wild. I have been sleeping on Claude Code and Warp
        
    * Those are the ones that are truly agentic editors.
        
    * AI Coding (which people are pissed at) versus, Agentic Coding, that is really the difference. Agentic coding opens a wild number of possibilities. This is the second time I am overwhelmed in life in programming. One was with Vim and Linux, there were so many things to learn and experiment with. And this time, it’s LLMs, models, tools, and so many details to learn, so many behaviours to understand. This is fascinating. Just watch this video, I can’t be thankful to this person enough.
        

## Learnt

* [GGUF is the file format for storing LLM model weights](https://vickiboykis.com/2024/02/28/gguf-the-long-way-around/)
    
    * I wanted to evaluate an idea for a project. Running models from a file, and this file format is what I needed. Using this format and a binding with llama.cpp and other libraries, this can be used for inference later to actually run the model
        
* Using llama.cpp python bindings to run a model with a gguf file
    
    * We can use the llama.cpp or other library binding to load the file in memory, and the binding library will use the inference to get the tokens out from the given prompt
        
* Creating Python Lambda functions in Vercel
    
    * The snippet is what you need to get up and running with Python serverless functions in Vercel
        
        ```bash
        import json
        from http.server import BaseHTTPRequestHandler
        from urllib.parse import parse_qs, urlparse
        
        
        class handler(BaseHTTPRequestHandler):
            
            def do_GET(self):
                parsed_path = urlparse(self.path)
                query_params = parse_qs(parsed_path.query)
                
                query = query_params.get('q', [''])[0]
                result = {"message": "hello, world!"}
                response_data = {"result": result}
        
                self.send_response(200)
                self.send_header('Content-type', 'application/json')
                self.end_headers()
                self.wfile.write(json.dumps(response_data).encode('utf-8'))
                print(response_data)
                return
        
        
            def do_POST(self):
                try:
                    content_length = int(self.headers.get('Content-Length', 0))
        
                    post_data_bytes = self.rfile.read(content_length)
        
                    try:
                        request_body = json.loads(post_data_bytes.decode('utf-8')
                    except json.JSONDecodeError:
                        self.send_response(400) # Bad Request
                        self.send_header('Content-type', 'application/json')
                        self.end_headers()
                        self.wfile.write(json.dumps({"error": "Invalid JSON"}).encode('utf-8'))
                        return
        
                   
                    response_message = {"status": "success", "received": request_body}
        
                    self.send_response(200)
                    self.send_header('Content-type', 'application/json')
                    self.end_headers()
                    self.wfile.write(json.dumps(response_message).encode('utf-8'))
        
                except Exception as e:
                    print(f"Error handling POST request: {e}")
                  
                    if not self.headers_sent:
                         try:
                             self.send_error(500, f"Server error during POST: {e}")
                         except Exception as send_err:
                             print(send_err)
                return
        ```
        
* Using the Warp terminal
    
    * Warp is a terminal and it also has agentic flow, as well as nice auto-completion, really awesome.
        

## Tech News

* [Qwen 3](https://qwenlm.github.io/blog/qwen3/) punches Meta on the face, becomes the sota model in open source models
    
    * Alibaba released the Qwen 3 family of models with almost Gemini 2.5-like results on an open weight model
        
    * The release was almost perfect, the integration so nice, the’ve put effort in making sure the model is available to all vendors on day 1, a state-of-the-art model release, set a standard for release as observed by [Simon Wilson](https://simonwillison.net/2025/Apr/29/qwen-3/) and a [few others](https://www.interconnects.ai/p/qwen-3-the-new-open-standard).
        
* [ChatGPT](https://help.openai.com/en/articles/10193193-1-800-chatgpt-calling-and-messaging-chatgpt-with-your-phone) and [Perplexity](https://x.com/AravSrinivas/status/1918138605203333264) make the chatbot available to chat on [WhatsApp](https://x.com/OpenAI/status/1916947244852646202)
    
    * This is a bit of a hot take, it can be revolutionary, it is exposing the LLMs to a wider audience, adoption rate might go high. I am not negative or positive at the moment, but it’s kind of revolutionary to see this thing. Open AI is again heading a revolution(not able to lead, however).
        
* [Redis becomes open source again](https://redis.io/blog/agplv3/)
    
    * Wow! A comeback from the tiny database that holds the internet.
        
    * More than that, the [v8 release](https://redis.io/blog/redis-8-ga/) sounds amazing! deep dive next week
        
* [Cloudflare launches a Python runtime for serverless functions](https://blog.cloudflare.com/streamable-http-mcp-servers-python/)
    
    * This is interesting, Cloudflare has Python runtime for serverless functions for quite some time now. I found that a few months back. It’s kind of there, and not. As it’s a Pyodide runtime and not a native Python runtime, so there are limitations there. Only a few selected packages are compatible or made available to the Pydodide runtime.
        
    * Kind of limiting, since other providers like Appwrite, Vercel, support full-fledged support for Python
        

---

Phew — that was a long, wild week. Feels like the whole ecosystem is shifting under our feet. Everything’s speeding up. If I had to bet, next week’s bombshell will come from Google or Anthropic — probably on May the 4th. Just a feeling.

That’s it for this week, will see you in the next one!

Thanks for reading

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public, so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
