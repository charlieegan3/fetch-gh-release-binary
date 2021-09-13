# fetch-gh-release-binary

GitHub Action generated to download binaries from GitHub releases and make them
available for use in later steps.

Example use:

```
- name: Download binary test
  uses: charlieegan3/fetch-gh-release-binary@main
  with:
    owner: charlieegan3
    repo: airtable-contacts
    asset-pattern: Linux_x86_64
    install-path: /usr/local/bin/airtable-contacts
    verbose: true
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```
