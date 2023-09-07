import datetime
from jinja2 import Template
from markata.hookspec import hook_impl
from pathlib import Path

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
    
    with open(project_template_file, 'r') as f:
        project_template = Template(f.read())
    
    body = "<ul class='project-list'>"
    project_cards = []
    for post in markata.articles:
        if post['templateKey'] == 'project':
            project_slug = post['slug']
            project_page = Path(projects_dir) / project_slug 
            project_page.mkdir(exist_ok=True, parents=True)
            
            # Create project page
            with open(project_page / 'index.html', 'w') as f:
                f.write(
                    project_template.render(
                        title=post['title'], 
                        cover_image=post['cover_image'],
                        body=post.content,
                        today=datetime.datetime.today(),
                    )
                )
                
            # Add card to projects page
            project_cards.append(f"""
            <section class="project-card">
                <img src="{post['cover_image']}" alt="{post['title']}">
                <h2>{post['title']}</h2>
                <p>{post['description']}</p>
                <a href="#">View Details</a>
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
