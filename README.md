# URL Shortener

Simple URL shortener, built as a challenge in 4 hours! Saw the challenge on YouTube, and wanted to try it myself.

You can access it <https://apk.rip>!

## 4 Hours

Within the 4 hours, I got a working product:

- Docker Setup
- API: Create shortened URLs
- API: List shortened URLs
- API: Redirect when visiting shortened URLs
- UI: Basic UI listing latest URLs and creating new URLs
- UI styling to much of what you see today (table, header, search bar)
- Heroku Deployment

After the initial 4 hours, I spent some time working on it to make it actually usable

- UI: Firebase Auth so only I can create and edit URLs
- Go 1.16 embedded FS to include UI assets in the single executable file
- Filter/search for URLs when filling out the search bar

TODO:

- API: Consume firebase auth
- Infinite scrolling to show more pages of URLs when scrolling
- Cleanup login flow

## Deployments

Deployments are done locally with:

```bash
make deploy
```
