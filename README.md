# changelog-nightly-parser
A lambda function that parses the [Changelog Nightly](http://nightly.changelog.com/) page for the last night (`http://nightly.changelog.com/YYYY/MM/DD`) and uploads to GitHub the list of trending repositories found (as a JSON file).

# How to use
Define the following environment variables in configuration of AWS Lambda:
- `GITHUB_REPOSITORY` - name of repository to which to upload the JSON file (eg. "trending-daily").
- `GITHUB_OWNER` - Github username (eg. "myusername")
- `GITHUB_TOKEN` - Github personal token (eg. "myusername")
