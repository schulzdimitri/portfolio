# Dimitri Schulz Amado — Portfolio

Personal portfolio website built with pure HTML, CSS, and JavaScript.  
No frameworks, no build step. Deployable to GitHub Pages in one command.

## Project Structure

```
portfolio/
├── index.html              # Entry point — only structural markup
├── README.md
│
├── css/
│   └── style.css           # All styles, organized by section/component
│
├── js/
│   └── main.js             # DOM interaction & dynamic rendering
│
├── img/
│   ├── profile.jpg         # Profile photo (hero section)
│   └── about.jpg           # Photo (about section)
│
├── assets/
│   ├── fonts/              # Self-hosted fonts (optional)
│   └── icons/              # Custom SVG icons or favicon assets
│
└── data/
    └── portfolio.json      # Content data — experiences, links, about text
```

## Design Decisions

| Concern | Solution |
|---|---|
| **Separation of concerns** | Content lives in `data/portfolio.json`. Structure in `index.html`. Style in `css/style.css`. Behavior in `js/main.js`. |
| **No unnecessary dependencies** | Zero npm, zero bundlers. Opens in any browser with no setup. |
| **Scalability** | Adding a new job experience requires only a JSON change, not touching HTML. |
| **Deployability** | Push to any static host (GitHub Pages, Netlify, Vercel) with no build step. |

## Running locally

Just open `index.html` in your browser, or serve it with any static server:

```bash
python3 -m http.server 3000
# Open http://localhost:3000
```

## Deploying to GitHub Pages

```bash
git add .
git commit -m "feat: initial portfolio release"
git push origin main
# Enable Pages in repo Settings → Pages → Source: main branch
```
