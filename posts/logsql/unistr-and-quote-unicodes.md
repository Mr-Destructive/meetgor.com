---
type: "logsql"
title: "SQLite functions: unistr and unistr_quote"
date: 2025-08-26
---

SQLite 3.50 introduced unistr and unistr_quote functions

These are for taking in escape sequences and returning the actual unicode string like emoji or other non-english characters outside the ascii convention forming the utf-8 character set

The unistr_quote is used to escape those escape sequences in a sql query so your data can be stored with emojis.

I wonder these functions were added because of LLMs emoji-heavy outputs?

Here’s how you use unistr in a simple query

SELECT unistr('\u2665, \U0001F600');

This will print the :heart and :smiley emojis

I won’t print the emojis, it might look like a AI-generated post (grunts with pride)

This is handy to render the values but first you need to get the values stored in the db, for that you need to have it like a string i.e. escaped string

That is where the unistr_quote comes in

SELECT unistr_quote('\U0001F600');

This will simply print '\U0001F600' notice the quotes here, it will be a escaped character sequence making it like a special string in middle of a string,

so if we have something like this

SELECT printf('what the heck is this%s symbol?', unistr_quote('\U0001F449'));

output: what the heck is this'\U0001F449' symbol?

This will make the distinction clear, we get a quoted string for that special set of characters or escape sequence which can be useful if the storage medium would make storing emojis obsolete.
