{
    "author":"meet",
    "date":"2025-09-27",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #61",
    "type":"newsletter",
    "slug": "techstructive-weekly-61"
}


## Week #61

A bit of slow and disappointing week on a personal note. I tried my best to be a backend developer but was not worthy of being one, I was stranded as a product developer. I am not saying I hate being a product developer, it just gets too menial and boring once you know the limitations and the quirks of a product.

Apart from the grill, I learnt a lot about tokenization. I also continued to write occasionally about SQLite in the SQLog, three more entries. Wasn’t able to livestream due to guests and was unwell for the weekend with cold.

Read a lot as usual, still keeping out from the doom-scrolling spiral. It is helping me on the mental level, but need to find something in life that will push me without looking for other things constantly. Maybe my work is not helping me here, its barely keeping me in peace, need to find a switch as fast as possible, the job market is not what you want to be in.

Let’s see being conscious and specific about life is getting important and need to take some actions to lift my spirits up.

### Quote of the week

> *“A home isn’t always the house we live in. It’s also the people we choose to surround ourselves with.”*
> 
> — From the book “The House in the Cerulean Sea” by T.J. Klune

I completed reading the book and felt deeply satisfied and cozy. What a book, hits hard, hits home. The way it presents dull and boring life into a beautiful narrative for a transformation. The lessons of empathy and kindness, being a human is so important in today’s age. Sorry, this is not a book review, but getting into the specific of this quote is grounded in the story.

The main character finds his new home which once he thought would just be a part of his job to do the needful for a month. A home is not something you just live in, its a relation with the people you make, the care each other makes, the choices and memories we make is what make a home, family.

## Created

Continued to write more about SQLite, not quite continuously. I had a break on Sunday, then got back after one day gaps. So only three articles for the week. But keeping it strong and sturdy.

- [SQLite SQL: Many to Many Table without RowID](https://www.meetgor.com/sqlog/sqlite-many-to-many-table-without-rowid)
- [SQLite SQL: Collate Column Modifier](https://www.meetgor.com/sqlog/sqlite-collate-column-modifier)
- [SQLite SQL: Create Temporary Table](https://www.meetgor.com/sqlog/sqlite-create-temp-table)

## Read

1. [I don’t want to code with LLMs](https://blaines-blog.com/I-dont-want-to-code-with-LLMs#footnote-ref-1)
    - This is the best of the lot. Nailed every point.
          - Coding was never a bottleneck, communication was (is).
          - Yes, it can do trivial task, but that is not 20% of developers work.
          - Its still bad at complex stuff. Vibe coders hit a ceiling after a while.
          - Reviewing is worse than writing it yourself.
          - AI is a tool just like IDEs, you are not losing or missing out on it.
          - When did learning deeper and low-level knowledge get uncool or not important? It hasn’t, its just hype, hiding form us the actual pillar behind the rise in information accessibility.
2. [A New kind of Code](https://registerspill.thorstenball.com/p/a-new-kind-of-code)
    - I like the analogy of glue code, but I don’t like reviewing code. Its skill issue on my side, yes, but I can’t stand reading code, I want to see it, I want to feel it. Generating with a button doesn’t give me any excitement to push forward to get things done.
    - Maybe I am naive, but I don’t like that button. Sometimes yes I do use it, sometimes forced to, but I like to keep nuts and bolt and prepare it myself. I am learning still maybe that is the reason. Can’t disagree to Thorsten, who can write 2 books about crafting programming languages and interpreters and compilers.
3. [How I a non-developer describe a developer](https://anniemueller.com/posts/how-i-a-non-developer-read-the-tutorial-you-a-developer-wrote-for-me-a-beginner)
    - This is so funny. I can feel this. Documentation, who writes that.
    - Tests, what are those? We push -f to prod
    - By the way, here’s the rollback script if the production is doomed.
4. [Processes and Threads: Planetscale Blog](https://planetscale.com/blog/processes-and-threads)
    - This was cool actually. The interactive elements really make things so clear. They remove the textual-ness in the blog and add a depth to it.
    - There is a clear explanation of program, process, ram, fork and low level details that are quintessential to a developer, any f-ing developer.
5. [Be careful with Go struct embedding](https://mattjhall.co.uk/posts/be-careful-with-go-struct-embedding.html)
    - Yikes, this hurts sometimes. Its just like SQLite, but without guardrails. Golang tries to be too lenient with the inferring the values for a struct and create a bag of unexpected behaviours.
6. [Artists are losing work, wages, and hope as bosses and clients embrace AI](https://www.bloodinthemachine.com/p/artists-are-losing-work-wages-and)
    - This is a serious topic, I have stopped using AI to generate thumbnails for my posts and whatever artistic form I used. Firstly it looks and feels dull and un-energetic.
    - I do respect artists, this might be a disrespect for them to use their knowledge and wisdom without giving them a credit.
    - I feel bad about the writers whose style and wisdom is now in-grained in models like GPT, Llama and Claude and others too.
7. [Brace yourself](https://thedailywtf.com/articles/brace-yourself)
    - It is always that damn comment. Developers write code, run it, if it works, fine, but if it doesn’t work, then debug it (printf debugging) and fix it and leave, we don’t test.
8. [One last id](https://thedailywtf.com/articles/one-last-id)
    - Man! SQL is tricky sometimes. The number of abstractions people create in the dialects is so jarring. It really breaks stuff.

## Watched

- [Tokenization from scratch from Andrej Karapathy](https://youtu.be/zduSFxRajkE)
    - What a beautiful piece of content. Archive and store it in a museum. The depth with which he explained it, the low-level details, the pythonic bits, is so fun and contagious to watch, and feel.
    - I learnt a few tricks about interaction with LLMs and understood certain quirks. This could give a intuition for why certain LLMs won’t be able to give good completions for certain tasks.
    - I also don’t quite liked the Sentence piece tokenization logic. But I can see where it could be probably come handy, in PDFs for example, the scope of sentence is well defined. In arbitrary piece of text on the internet, it might not be.

Double click to interact with video
- [Tokenization in C from Tsoding](https://youtu.be/6dCqR9p0yWY)
    - This was another great livestream like tutorial. The depth with which he communicated and came up with the solution is what helped me get better understanding of tokenization.
    - This is an example, why learning from first fundamentals is still cool. The ability to learn and explain with confidence and comfort is remarkable quality for a developer. This is also I am still thinking about doing livestream and being able to develop something from scratch.

Double click to interact with video
- [What is a Tensor? A beautiful intuition and question and answer based explanation](https://youtu.be/k2FP-T6S1x0)
    - This is a great piece of explanation. The question from Richard Feymann are so deep, provoking and sensible. The way the author explained and questioned his own thinking is really great. I liked that way of teaching.
    - The direction bit and the animation also helped a lot. Nice editing skills.

Double click to interact with video
- [I hate myself more for seeing this, than I hate javascript](https://youtu.be/7bvBVBy_CrM)
    - This is gross. I really hate seeing this now. I can’t bare this.
    - I mean, javascript is a good language, but why people just used it and didn’t improve it. The author developed it in a week, weren’t the industry leaders a bit mature to make it better? The Java developers have fixed things from it, but the thing that was copied to Javascript is still ain’t? This goes back to the meme of pillars holding the bigger stones.

Double click to interact with video

## Learnt

- How BPE or Byte Pair Encoding Transformer works
    - It basically finds the most frequent occurring two characters and groups them as a new token (apart from the individual token / characters)
    - These pair of word is added to the vocabulary or a lookup reference for mapping the character set to the id of the token.
    - Then it does the above iteratively until the most frequent pair has count of one, we can limit that. Then at the end the vocabulary developed will be used to encode and decode the content (text)
    - There is a sweet spot of minimum token size instead of converging to the entire dataset for maximum tokens, it can compress a lot of information about a single word that it might not be able to understand it.
    - This is a classical example and limitation of why LLMs can’t do math or count how many r’s in strawberry as strawberry might be a single token and it can’t get enough data about the characters in just one token. It has a lot of knowledge to get that number but just sheer looking at the tokens, it can’t.
    - The other example of adding less tokens is that we might have a lot of tokens for a given content and we might run of out context window to process the request. This is true for languages other than English, LLMs aren’t good at that the earlier ones especially as the tokenisation for those was not optimized or had fewer tokens due to selection of dataset.
- SQLite
    - Create temporary tables, the temporary table is actually stored in the temp location file and in a separate database.
    - If you had two tables with the same name one as the primary permanent table and other as the temporary table, if you just refer the table name, the temporary table will be preferred over the permanent. You will have to use `main.<table name> `to actually refer to the permanent table.

## Tech News

- [Postgres 18 released](https://www.postgresql.org/about/news/postgresql-18-released-3142/):
    - This has some cool features like generated virtual column as default generated columns.
- [Planetscale makes Postgres Generally Available](https://planetscale.com/blog/planetscale-for-postgres-is-generally-available)
    - This is a big one, they are going all in postgres.
- [Github launches Copilot CLI: Another entry for the Agentic CLIs, oh god!](https://github.blog/changelog/2025-09-25-github-copilot-cli-is-now-in-public-preview/)
    - We have Claude Code, Gemini CLI, Codex, Warp, Amp, Cursor agent CLI, and oh my gosh, this list is getting too big. Developers are using it, this is going to create a IDE-like war in the near future, if developers exist.
- [Improved Gemini 2.5 Flash and Flash lite models](https://developers.googleblog.com/en/continuing-to-bring-you-our-latest-models-with-an-improved-gemini-2-5-flash-and-flash-lite-release/)
    - This looks like a good improvement, reducing the number of output token is a good sign from Google. They are not looking to increase their revenue, they are taking a hit for it.
- [ChatGPT Pulse](https://openai.com/index/introducing-chatgpt-pulse/)
    - Not sure this is a right step for OpenAI, they are going to far with this one.
- [Meta releases open weight coding model](https://ai.meta.com/research/publications/cwm-an-open-weights-llm-for-research-on-code-generation-with-world-models/)
    - This is a bit cool, it has almost the intelligence of Gemini 2.5 Thinking with a lot less parameters (32 B)
    - It also has a different and quite novel approach for a coding model

That’s it from this week. It has been a harsher week. We are finally into the last quarter of 2025. Ah! That ended pretty quickly. LLMs have just consumed me.

---

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-764) (#764th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
