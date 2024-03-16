import datetime
from pathlib import Path, PosixPath
from markata.hookspec import hook_impl
from jinja2 import Template, FileSystemLoader, Environment

@hook_impl 
def save(markata):
    projects_dir = Path('markout/projects')
    projects_page = projects_dir / 'index.html'
    projects_dir.mkdir(exist_ok=True)
    template_file = (
        Path(__file__).resolve().parents[1] / "layouts" / "projects_list_template.html"
    )
    project_template_file = (
        Path(__file__).resolve().parents[1] / "layouts" / "project_template.html"
    )
    with open(template_file, 'r') as f:
        template = Template(f.read())

    with open(template_file) as f:
        env = Environment()
        env.loader = FileSystemLoader('layouts/')
        if type(template_file) is PosixPath:
            template = template_file.name
        template = env.get_template(template)

    
    with open(project_template_file, 'r') as f:
        project_template = Template(f.read())

    
    projects = [post for post in markata.articles if post['templateKey'] == 'project']
    tags = []
    for project in projects:
        tags = tags + project['tags']
    tags = list(set(tags))
    tags.sort()
    body = ""
    body += "<ul class='tag-list'>"
    for tag in tags:
        body += f"""
        <a href="#{tag}">{tag}</a> |
        """
    body += "</ul><ul class='project-list'>"
    project_cards = []
    for post in markata.articles:
        if post['templateKey'] == 'project':
            project_slug = post.get('slug')
            project_page = Path(projects_dir) / project_slug 
            project_page.mkdir(exist_ok=True, parents=True)
            post['cover_image'] = post.get('cover_image', 'https://github.com/Mr-Destructive/meetgor.com/blob/gh-pages/static/tbicon.png?raw=true')
            
            # Create project page
            with open(project_page / 'index.html', 'w') as f:
                f.write(
                    project_template.render(
                        title=post['title'], 
                        cover_image=post['cover_image'],
                        body=post.content,
                        today=datetime.datetime.today(),
                        github_link=post['github_link'],
                    )
                )
                
            # Add card to projects page
            project_cards.append(f"""
            <section class="project-card">
                <h2>{post['title']}</h2>
                <p>{post['description']}</p>
                <button><a href="{post['slug']}">View Details</a></button>
            </section>
            """)
    body = body + "".join(project_cards) + "</ul>"
            
    # Generate projects page
    with open(projects_page, "w+") as f:
        f.write(
            template.render(
                title="Projects",
                body=body,
                today=datetime.datetime.today(),
            )
        )
