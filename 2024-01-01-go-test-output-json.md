---
article_html: "<p>I just discovered that we can generate a JSON output of test results
  in Golang. I found this <a href=\"https://youtu.be/cf72gMBrsI0?t=80\">here</a>.</p>\n<p>Let's
  take a fresh simple example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">jsontest</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">hello</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">jsontest</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;testing&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">TestHello</span><span class=\"p\">(</span><span class=\"nx\">t</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">testing</span><span
  class=\"p\">.</span><span class=\"nx\">T</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">want</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Logf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q, want: %q&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have a function <code>hello</code> that simply returns a string, the <code>TestHello</code>
  function in the <code>jsontest</code> package will check if the returned string
  is correctly returned or not.</p>\n<p>So, we can test this with <code>go test ./...</code>
  command, this will give out the output in a standard output/error in text format.
  However, if we add the <code>-json</code> flag, we can get the output in JSON format.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> ./... -json\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.861974085+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863133332+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863142397+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestHello\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863148346+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:10: got: \\&quot;Hello, World\\&quot;\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863151351+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:12: got: \\&quot;Hello, World\\&quot;, want:
  \\&quot;Hello, World!\\&quot;\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863157014+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;---
  FAIL: TestHello (0.00s)\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863160418+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;fail&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mi\">0</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863411555+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;FAIL\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863438344+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;FAIL\\tjson-test\\t0.001s\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863443461+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;fail&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mf\">0.001</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Pretty
  cool right?</p>\n<p>This is really useful for programmatically taking the output
  and parsing it to get the metrics.</p>\n<p>We can even combine with the coverage
  flag to get the coverage metrics as well.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> ./... -json -cover\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.771976961+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.775118482+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775172535+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestHello\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775201647+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:10: got: \\&quot;Hello, World!\\&quot;\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775231759+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;--- PASS: TestHello (0.00s)\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775253928+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">0</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.775269402+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;PASS\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.776153185+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;coverage:
  100.0% of statements\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.776808599+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;ok  \\tjsontest\\t0.004s\\tcoverage:
  100.0% of statements\\n&quot;</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.777814589+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mf\">0.006</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>I am
  planning to use this in my workflow for integrating the output of the tests suite
  with specific tests.</p>\n<p>For running the specific tests, you can use <code>go
  test -run TestName</code> command, this will only run the provided test function.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// main.go</span><span class=\"w\"></span>\n\n<span class=\"kn\">package</span><span
  class=\"w\"> </span><span class=\"nx\">jsontest</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello,
  World!&quot;</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">add</span><span
  class=\"p\">(</span><span class=\"nx\">x</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">x</span><span
  class=\"w\"> </span><span class=\"o\">+</span><span class=\"w\"> </span><span class=\"nx\">y</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// main_test.go</span><span class=\"w\"></span>\n<span class=\"kn\">package</span><span
  class=\"w\"> </span><span class=\"nx\">jsontest</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">import</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"s\">&quot;testing&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">TestHello</span><span
  class=\"p\">(</span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">testing</span><span class=\"p\">.</span><span class=\"nx\">T</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">want</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Logf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q, want: %q&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">TestAdd</span><span
  class=\"p\">(</span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">testing</span><span class=\"p\">.</span><span class=\"nx\">T</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">want</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">add</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">t</span><span class=\"p\">.</span><span
  class=\"nx\">Logf</span><span class=\"p\">(</span><span class=\"s\">&quot;got: %d&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %d, want: %d&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we have two test in this go module, we can run a specific test using <code>go test
  -run TestName</code> command as so:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> -run TestAdd -json\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.19397581+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198067398+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198150156+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestAdd\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198197444+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:19: got: 2\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198217057+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;--- PASS: TestAdd (0.00s)\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198230965+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">0</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198241628+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;PASS\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.19869148+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;ok
  \ \\tjsontest\\t0.004s\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198822637+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;pass&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span
  class=\"p\">:</span><span class=\"mf\">0.005</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As we can see, there is only one
  test being executed and the output of the test is in JSON format.</p>\n<p>These
  are really good flags and options to have as they make the output more portable.
  I will be planning to use this to improve my workflow in testing and developing
  open source projects and personal projects as well. I am really inspired by the
  Teej's video of executing anything in NeoVim.</p>\n<div class='prevnext'>\n    <style
  type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/django-bulk-update-queryset'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Django
  Bulk Update QuerySet objects&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next'
  href='/vimscript-to-lua-keymapper'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Map
  Vimscript Keymaps to Lua with a single function&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg
  width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot;
  xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n        &lt;path d=&quot;M10.5
  15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot; stroke-width=&quot;1.5&quot;
  stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2024-01-01
datetime: 2024-01-01 00:00:00+00:00
description: Obtain JSON output of test results in Golang
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2024-01-01-go-test-output-json.md
html: "<p>I just discovered that we can generate a JSON output of test results in
  Golang. I found this <a href=\"https://youtu.be/cf72gMBrsI0?t=80\">here</a>.</p>\n<p>Let's
  take a fresh simple example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">jsontest</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">hello</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">jsontest</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;testing&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">TestHello</span><span class=\"p\">(</span><span class=\"nx\">t</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">testing</span><span
  class=\"p\">.</span><span class=\"nx\">T</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">want</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Logf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q, want: %q&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have a function <code>hello</code> that simply returns a string, the <code>TestHello</code>
  function in the <code>jsontest</code> package will check if the returned string
  is correctly returned or not.</p>\n<p>So, we can test this with <code>go test ./...</code>
  command, this will give out the output in a standard output/error in text format.
  However, if we add the <code>-json</code> flag, we can get the output in JSON format.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> ./... -json\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.861974085+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863133332+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863142397+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestHello\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863148346+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:10: got: \\&quot;Hello, World\\&quot;\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863151351+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:12: got: \\&quot;Hello, World\\&quot;, want:
  \\&quot;Hello, World!\\&quot;\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863157014+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;---
  FAIL: TestHello (0.00s)\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T20:52:45.863160418+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;fail&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mi\">0</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863411555+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;FAIL\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863438344+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;FAIL\\tjson-test\\t0.001s\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T20:52:45.863443461+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;fail&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;json-test&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mf\">0.001</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Pretty
  cool right?</p>\n<p>This is really useful for programmatically taking the output
  and parsing it to get the metrics.</p>\n<p>We can even combine with the coverage
  flag to get the coverage metrics as well.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> ./... -json -cover\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.771976961+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.775118482+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775172535+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestHello\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775201647+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:10: got: \\&quot;Hello, World!\\&quot;\\n&quot;</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775231759+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;--- PASS: TestHello (0.00s)\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.775253928+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestHello&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">0</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.775269402+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;PASS\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.776153185+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;coverage:
  100.0% of statements\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:13:30.776808599+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;ok  \\tjsontest\\t0.004s\\tcoverage:
  100.0% of statements\\n&quot;</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:13:30.777814589+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span class=\"mf\">0.006</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>I am
  planning to use this in my workflow for integrating the output of the tests suite
  with specific tests.</p>\n<p>For running the specific tests, you can use <code>go
  test -run TestName</code> command, this will only run the provided test function.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// main.go</span><span class=\"w\"></span>\n\n<span class=\"kn\">package</span><span
  class=\"w\"> </span><span class=\"nx\">jsontest</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello,
  World!&quot;</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">add</span><span
  class=\"p\">(</span><span class=\"nx\">x</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">x</span><span
  class=\"w\"> </span><span class=\"o\">+</span><span class=\"w\"> </span><span class=\"nx\">y</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// main_test.go</span><span class=\"w\"></span>\n<span class=\"kn\">package</span><span
  class=\"w\"> </span><span class=\"nx\">jsontest</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">import</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"s\">&quot;testing&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">TestHello</span><span
  class=\"p\">(</span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">testing</span><span class=\"p\">.</span><span class=\"nx\">T</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">want</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello, World!&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">hello</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Logf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">got</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %q, want: %q&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">TestAdd</span><span
  class=\"p\">(</span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">testing</span><span class=\"p\">.</span><span class=\"nx\">T</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">want</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">add</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">t</span><span class=\"p\">.</span><span
  class=\"nx\">Logf</span><span class=\"p\">(</span><span class=\"s\">&quot;got: %d&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">t</span><span class=\"p\">.</span><span class=\"nx\">Errorf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;got: %d, want: %d&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">got</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">want</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we have two test in this go module, we can run a specific test using <code>go test
  -run TestName</code> command as so:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  <span class=\"nb\">test</span> -run TestAdd -json\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.19397581+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;start&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198067398+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;run&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Test&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198150156+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;=== RUN   TestAdd\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198197444+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;    main_test.go:19: got: 2\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198217057+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;--- PASS: TestAdd (0.00s)\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.198230965+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;pass&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Test&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;TestAdd&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">0</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198241628+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;output&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Output&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;PASS\\n&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;2024-01-01T21:33:44.19869148+05:30&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;output&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span
  class=\"p\">:</span><span class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Output&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;ok
  \ \\tjsontest\\t0.004s\\n&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"nt\">&quot;Time&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;2024-01-01T21:33:44.198822637+05:30&quot;</span><span class=\"p\">,</span><span
  class=\"nt\">&quot;Action&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;pass&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;Package&quot;</span><span class=\"p\">:</span><span
  class=\"s2\">&quot;jsontest&quot;</span><span class=\"p\">,</span><span class=\"nt\">&quot;Elapsed&quot;</span><span
  class=\"p\">:</span><span class=\"mf\">0.005</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As we can see, there is only one
  test being executed and the output of the test is in JSON format.</p>\n<p>These
  are really good flags and options to have as they make the output more portable.
  I will be planning to use this to improve my workflow in testing and developing
  open source projects and personal projects as well. I am really inspired by the
  Teej's video of executing anything in NeoVim.</p>\n<div class='prevnext'>\n    <style
  type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/django-bulk-update-queryset'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Django
  Bulk Update QuerySet objects&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next'
  href='/vimscript-to-lua-keymapper'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Map
  Vimscript Keymaps to Lua with a single function&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg
  width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot;
  xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n        &lt;path d=&quot;M10.5
  15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot; stroke-width=&quot;1.5&quot;
  stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: I just discovered that we can generate a JSON output of test results
  in Golang. I found this  Let Here, we have a function  So, we can test this with  Pretty
  cool right? This is really useful for programmatically taking the output and parsing
  it to g
now: 2025-05-01 18:11:33.315138
path: blog/tils/2024-01-01-go-test-output-json.md
slug: golang-test-output-json
status: published-til
tags:
- go
templateKey: til
title: 'Golang: Test Output JSON'
today: 2025-05-01
---

I just discovered that we can generate a JSON output of test results in Golang. I found this [here](https://youtu.be/cf72gMBrsI0?t=80).

Let's take a fresh simple example.

```go
package jsontest

func hello() string {
    return "Hello, World!"
}
```
```go
package jsontest

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello, World!"
    got := hello()
    t.Logf("got: %q", got)
    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}
```

Here, we have a function `hello` that simply returns a string, the `TestHello` function in the `jsontest` package will check if the returned string is correctly returned or not.

So, we can test this with `go test ./...` command, this will give out the output in a standard output/error in text format. However, if we add the `-json` flag, we can get the output in JSON format.


```bash
go test ./... -json
```

```json
{"Time":"2024-01-01T20:52:45.861974085+05:30","Action":"start","Package":"json-test"}
{"Time":"2024-01-01T20:52:45.863133332+05:30","Action":"run","Package":"json-test","Test":"TestHello"}
{"Time":"2024-01-01T20:52:45.863142397+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"=== RUN   TestHello\n"}
{"Time":"2024-01-01T20:52:45.863148346+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"    main_test.go:10: got: \"Hello, World\"\n"}
{"Time":"2024-01-01T20:52:45.863151351+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"    main_test.go:12: got: \"Hello, World\", want: \"Hello, World!\"\n"}
{"Time":"2024-01-01T20:52:45.863157014+05:30","Action":"output","Package":"json-test","Test":"TestHello","Output":"--- FAIL: TestHello (0.00s)\n"}
{"Time":"2024-01-01T20:52:45.863160418+05:30","Action":"fail","Package":"json-test","Test":"TestHello","Elapsed":0}
{"Time":"2024-01-01T20:52:45.863411555+05:30","Action":"output","Package":"json-test","Output":"FAIL\n"}
{"Time":"2024-01-01T20:52:45.863438344+05:30","Action":"output","Package":"json-test","Output":"FAIL\tjson-test\t0.001s\n"}
{"Time":"2024-01-01T20:52:45.863443461+05:30","Action":"fail","Package":"json-test","Elapsed":0.001}
```

Pretty cool right?

This is really useful for programmatically taking the output and parsing it to get the metrics.


We can even combine with the coverage flag to get the coverage metrics as well.

```bash
go test ./... -json -cover
```

```json
{"Time":"2024-01-01T21:13:30.771976961+05:30","Action":"start","Package":"jsontest"}
{"Time":"2024-01-01T21:13:30.775118482+05:30","Action":"run","Package":"jsontest","Test":"TestHello"}
{"Time":"2024-01-01T21:13:30.775172535+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"=== RUN   TestHello\n"}
{"Time":"2024-01-01T21:13:30.775201647+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"    main_test.go:10: got: \"Hello, World!\"\n"}
{"Time":"2024-01-01T21:13:30.775231759+05:30","Action":"output","Package":"jsontest","Test":"TestHello","Output":"--- PASS: TestHello (0.00s)\n"}
{"Time":"2024-01-01T21:13:30.775253928+05:30","Action":"pass","Package":"jsontest","Test":"TestHello","Elapsed":0}
{"Time":"2024-01-01T21:13:30.775269402+05:30","Action":"output","Package":"jsontest","Output":"PASS\n"}
{"Time":"2024-01-01T21:13:30.776153185+05:30","Action":"output","Package":"jsontest","Output":"coverage: 100.0% of statements\n"}
{"Time":"2024-01-01T21:13:30.776808599+05:30","Action":"output","Package":"jsontest","Output":"ok  \tjsontest\t0.004s\tcoverage: 100.0% of statements\n"
{"Time":"2024-01-01T21:13:30.777814589+05:30","Action":"pass","Package":"jsontest","Elapsed":0.006}
```

I am planning to use this in my workflow for integrating the output of the tests suite with specific tests.

For running the specific tests, you can use `go test -run TestName` command, this will only run the provided test function.

```go
// main.go

package jsontest

func hello() string {
    return "Hello, World!"
}

func add(x, y int) int {
    return x + y
}
```

```go
// main_test.go
package jsontest

import (
	"testing"
)

func TestHello(t *testing.T) {
    want := "Hello, World!"
    got := hello()
    t.Logf("got: %q", got)
    if got != want {
        t.Errorf("got: %q, want: %q", got, want)
    }
}

func TestAdd(t *testing.T) {
    want := 2
    got := add(1, 1)
    t.Logf("got: %d", got)
    if got != want {
        t.Errorf("got: %d, want: %d", got, want)
    }
}
```

So, we have two test in this go module, we can run a specific test using `go test -run TestName` command as so:

```bash
go test -run TestAdd -json
```

```json
{"Time":"2024-01-01T21:33:44.19397581+05:30","Action":"start","Package":"jsontest"}
{"Time":"2024-01-01T21:33:44.198067398+05:30","Action":"run","Package":"jsontest","Test":"TestAdd"}
{"Time":"2024-01-01T21:33:44.198150156+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"=== RUN   TestAdd\n"}
{"Time":"2024-01-01T21:33:44.198197444+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"    main_test.go:19: got: 2\n"}
{"Time":"2024-01-01T21:33:44.198217057+05:30","Action":"output","Package":"jsontest","Test":"TestAdd","Output":"--- PASS: TestAdd (0.00s)\n"}
{"Time":"2024-01-01T21:33:44.198230965+05:30","Action":"pass","Package":"jsontest","Test":"TestAdd","Elapsed":0}
{"Time":"2024-01-01T21:33:44.198241628+05:30","Action":"output","Package":"jsontest","Output":"PASS\n"}
{"Time":"2024-01-01T21:33:44.19869148+05:30","Action":"output","Package":"jsontest","Output":"ok  \tjsontest\t0.004s\n"}
{"Time":"2024-01-01T21:33:44.198822637+05:30","Action":"pass","Package":"jsontest","Elapsed":0.005}
```

As we can see, there is only one test being executed and the output of the test is in JSON format.

These are really good flags and options to have as they make the output more portable. I will be planning to use this to improve my workflow in testing and developing open source projects and personal projects as well. I am really inspired by the Teej's video of executing anything in NeoVim.
<div class='prevnext'>
    <style type='text/css'>
    :root {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="light"] {
      --prevnext-color-text: #1f2022;
      --prevnext-color-angle: #ffeb00;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="dark"] {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    .prevnext {
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: flex-start;
    }
    .prevnext a {
      display: flex;
      align-items: center;
      width: 100%;
      text-decoration: none;
    }
    a.next {
      justify-content: flex-end;
    }
    .prevnext a:hover {
      background: #00000006;
    }
    .prevnext-subtitle {
      color: var(--prevnext-color-text);
      filter: brightness(var(--prevnext-subtitle-brightness));
      font-size: .8rem;
    }
    .prevnext-title {
      color: var(--prevnext-color-text);
      font-size: 1rem;
    }
    .prevnext-text {
      max-width: 30vw;
    }
    </style>
    
    <a class='prev' href='/django-bulk-update-queryset'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Django Bulk Update QuerySet objects</p>
        </div>
    </a>
    
    <a class='next' href='/vimscript-to-lua-keymapper'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Map Vimscript Keymaps to Lua with a single function</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>