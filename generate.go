package sdk

//go:generate openapi-generator generate -i api-docs/reference/SpaceTraders.json -o . -g go -p packageName=sdk --git-user-id realfatcat --git-repo-id go-space-traders -p enumClassPrefix=true
