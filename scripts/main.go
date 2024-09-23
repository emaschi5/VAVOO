name: Generate and Publish M3U Files

on:
  schedule:
    - cron: '0 */6 * * *' # Runs every 6 hours
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout repository
      uses: actions/checkout@v2
      with:
        fetch-depth: 0

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.16'

    - name: Run main.go
      run: go run scripts/main.go

    - name: Remove untracked files
      run: git clean -f

    - name: Commit and push changes to main-m3u
      run: |
        git config --global user.name 'github-actions[bot]'
        git config --global user.email 'github-actions[bot]@users.noreply.github.com'
        git fetch origin main-m3u || git branch main-m3u
        git checkout main-m3u
        git add *.m3u index.html  # Only add M3U files and index.html
        git commit -m 'Automated update of M3U files' || echo "No changes to commit"
        git push origin main-m3u

    - name: Deploy to gh-pages
      uses: peaceiris/actions-gh-pages@v3
      with:
        github_token: ${{ secrets.GITHUB_TOKEN }}
        publish_dir: ./ # Publish the entire repository
        publish_branch: gh-pages
        force_orphan: true
