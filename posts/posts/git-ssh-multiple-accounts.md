{
  "title": "Adding SSH Keys for Multiple Accounts in Git",
  "status": "published",
  "slug": "git-ssh-multiple-accounts",
  "date": "2025-04-05T12:33:27Z"
}

<p>Let's  say you have multiple github accounts. One for your personal projects, one for your company that you work at, and one other remote repository account (let's say gitlab).</p>
<p>You are juggling with multiple accounts, you should not waste much time and pick a SSH from those remote repository and pull it in your local machine, that makes the process just smooth and saves a ton of time.</p>
<h3>Create a SSH Key</h3>
<p>To create a SSH key, in linux you can use <code>ssh-keygen</code> command.</p>
<pre><code class="language-bash">ssh-keygen -t ed25519 -C &quot;alice@example.com&quot;
</code></pre>
<p>The above command will prompt you for two things</p>
<ol>
<li>The location where you want to store the key</li>
<li>The passphrase for accessing the key</li>
</ol>
<h3>Add SSH Key to Github</h3>
<p>Locate to the <code>ssh</code> folder and copy the generated <code>.pub</code> file to your <code>github</code> account.</p>
<p>For example, if you have created the key at <code>~/.ssh/your_name</code> then copy the contents of the file <code>~/.ssh/your_name.pub</code> to your clipbaord.</p>
<p>Navigate to your <code>github</code> account and in the settings, <code>SSH and GPG keys</code> tab, click on <code>Add SSH key</code> and copy the contents of your clipboard to the <code>Key</code> field.</p>
<h3>Configuring the SSH keys for multiple accounts</h3>
<pre><code class="language-config">Host your_company
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_company

Host your_name
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_name

Host some_name
  HostName gitlab.com
  User git
  IdentityFile ~/.ssh/some_name
</code></pre>
<p>You can change the <code>Host</code> config tag values in the <code>~/.ssh/conFig</code></p>
<p>The next time you clone/create a repository on those remote git providers, you need to specify the ssh key for that account.</p>
<p>For example, if you have a repository <code>github.com/StartUp_company/some_wired_project</code> then you can specify the remote as <code>git@your_company.com:StartUp_company/some_wired_project</code>. Here, the <code>git@your_company</code> is the <code>Host</code> value tag from the <code>~/.ssh/config</code>. If that repository is from your <code>your_company</code> organisation/user scope, you need to add the <code>git@your_company</code> tag, if that's your project, simply add <code>git@your_name</code> before the repository url i.e. <code>your_name/repo_name</code> which would set the origin as <code>git@your_name:your_name/repo_name</code>, here the 1st <code>your_name</code> is the tag from the <code>Host</code> config and the 2nd <code>your_name</code> is the github username.</p>
<p>So, in summary if you wanted to use multiple accounts in the same machine, you can understand in the following example:</p>
<pre><code class="language-bash">ssh -T git@your_name

git clone https://github.com/your_name/repo_name
</code></pre>
<p>However, you will need to authenticate with the ssh keys in this way everytime you push/pull a repository. So for that, you can set the origin with the <code>git@your_name</code> tag as the host for automatically authenticating the ssh keys on every push/pull or other activities.</p>
<p>Thanks for reading, Happy Coding :)</p>
