const statusEl = document.getElementById('status');

function log(msg, cls = '') {
  statusEl.className = `status ${cls}`.trim();
  statusEl.textContent = `${new Date().toLocaleString()}\n${msg}`;
}

function getCreds() {
  const username = document.getElementById('username').value.trim();
  const password = document.getElementById('password').value;
  if (!username || !password) {
    throw new Error('username and password are required');
  }
  return { username, password };
}

function parseTags(value) {
  return value
    .split(',')
    .map((s) => s.trim())
    .filter(Boolean);
}

async function postJSON(url, payload) {
  const res = await fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });

  const raw = await res.text();
  let body;
  try {
    body = JSON.parse(raw);
  } catch {
    body = raw;
  }

  if (!res.ok) {
    throw new Error(typeof body === 'string' ? body : JSON.stringify(body));
  }
  return body;
}

async function createPost() {
  try {
    const { username, password } = getCreds();
    const title = document.getElementById('title').value.trim();
    const content = document.getElementById('content').value;
    const type = document.getElementById('type').value;
    const slug = document.getElementById('slug').value.trim();
    const tags = parseTags(document.getElementById('tags').value);

    if (!title || !content) {
      throw new Error('title and content are required');
    }

    const metadata = {
      title,
      type,
      post_dir: type,
      published: 'published',
      tags,
    };
    if (slug) {
      metadata.slug = slug;
    }

    const payload = { username, password, title, content, metadata };
    const body = await postJSON('/cms/api/create', payload);
    log(`post created successfully\n${JSON.stringify(body, null, 2)}`, 'ok');
  } catch (err) {
    log(`create failed: ${err.message}`, 'err');
  }
}

async function syncDb() {
  try {
    const payload = getCreds();
    const body = await postJSON('/cms/api/sync-db', payload);
    log(`sync successful\n${JSON.stringify(body, null, 2)}`, 'ok');
  } catch (err) {
    log(`sync failed: ${err.message}`, 'err');
  }
}

async function triggerBuild() {
  try {
    const payload = getCreds();
    const body = await postJSON('/cms/api/trigger-build', payload);
    log(`build trigger successful\n${JSON.stringify(body, null, 2)}`, 'ok');
  } catch (err) {
    log(`build trigger failed: ${err.message}`, 'err');
  }
}

document.getElementById('createBtn').addEventListener('click', createPost);
document.getElementById('syncBtn').addEventListener('click', syncDb);
document.getElementById('buildBtn').addEventListener('click', triggerBuild);
