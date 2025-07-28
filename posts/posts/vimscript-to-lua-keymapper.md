{
  "title": "Map Vimscript Keymaps to Lua with a single function",
  "status": "published",
  "slug": "vimscript-to-lua-keymapper",
  "date": "2025-04-05T12:33:39Z"
}

<h2>Introduction</h2>
<p>Are you bored of writing all the keymaps from vimscript to lua? Try the below function to create all your keymaps to lua equivalent maps in Neovim.</p>
<p>Take your vimscript keymaps and put them in lua don't write any lua for it ;)</p>
<h2>The Lua Function</h2>
<p>The below-provided snippet is a lua function that takes in a table of strings(list of strings), the strings will be your keymaps. The function then maps these keymaps using lua functions. You don't have to type out all the keymaps by yourself. It can also print out the lua equivalent function calls required to map the existing keymaps from vimscript to lua runtime in Neovim. Though it won't handle all the options, we have passed in a default value to the keymap.</p>
<pre><code class="language-lua">function key_mapper(keymaps)
  for _, keymap in ipairs(keymaps) do
    local mode = keymap:sub(1,1)
    local delimiter = &quot; &quot;
    local lhs = ''
    local rhs_parts = {}
    local m = 0
    local options = {noremap = true}
    for matches in (keymap..delimiter):gmatch(&quot;(.-)&quot;..delimiter) do
      if m == 1 then
        lhs = matches
      end
      if m &gt;= 2 then
        table.insert(rhs_parts, matches)
      end
      m = m + 1
    end
    rhs = ''
    for _, p in ipairs(rhs_parts) do
      rhs = rhs .. &quot; &quot; .. p
    end
    --print(&quot;vim.keymap.set(&quot;..&quot;\'&quot;..mode..&quot;\'&quot;..&quot;, &quot;..&quot;\'&quot;..lhs..&quot;\'&quot;..&quot;, &quot;..&quot;\'&quot;..rhs..&quot;\'&quot;..&quot;, &quot;..vim.inspect(options)..&quot;)&quot;)
    vim.keymap.set(mode, lhs, rhs, options)
  end
end
</code></pre>
<p>You can uncomment the print statement <strong>once</strong> to grab the keymaps and paste them into the config file. If you leave it uncommented, it might print every time you open up a new neovim instance. The function can be called like below:</p>
<pre><code class="language-lua">key_mapper({
  'nnoremap cpp :!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;i',
  'nnoremap c, :!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;',
  'nnoremap py :!python %&lt;cr&gt;',
  'nnoremap go :!go run %&lt;cr&gt;',
  'nnoremap sh :!bash %&lt;CR&gt;'
})
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657559501/blog-media/neovim-lua-keymapper.gif" alt="Keymapper demonstration"></p>
<p>We pass in a table of strings, these strings are just the vimscript keymaps. This function call will then map the keymaps into equivalent lua maps. You can customize it as per your needs.</p>
<p>For further references, you can check out my <a href="https://github.com/Mr-Destructive/dotfiles">dotfiles</a> on GitHub.</p>
<h2>How the function works</h2>
<p>The function is simply a text scrapping from lua strings. We extract the first character in the string for the mode, grab the strings which are space-separated and finally sort out which are lhs and rhs sides of the maps.</p>
<p>We iterate over the table in lua with the help of ipairs function which allows us to iterate over an ordered list of items in a table. Using the gmatch function, we find a pattern to split the string with the space as the delimiter. Thereby, we can have separate sets of strings identified as rhs and lhs. We can store them in variables as strings as the lua functions require them as strings.</p>
<p>We simply add those variables into the <a href="https://neovim.io/doc/user/lua.html#:~:text=set(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7Bopts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20*vim.keymap.set()*">vim.keymap.set</a> or <a href="https://neovim.io/doc/user/api.html#nvim_set_keymap():~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*">vim.api.nvim_set_keymap</a> functions. We by default set the value of <code>{noremap: True}</code> to avoid teh recursive mapping of the keys. These option parameter is the one which needs to be a bit more dynamic in terms of wide variety of keymaps.</p>
<p>So, this is how we can convert the vimscript keymaps to lua in Neovim. Hope you found this useful. Thanks for reading. Happy Viming :)</p>
