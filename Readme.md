# Simple GO lang webapp: using gin framework and mongodb

This is a simple golang webapp

### Important Instructions :

  - git clone : `git@github.com:rohit-raj/basic-go-webapp.git`
  - `go get package_path`
  - make changes to `config.dev.json` or `config.prod.json` inside config folder, by adding your mongodb url to url field.
  - `env=dev go run app.go` or `env=prod go run app.go`

Note: just make the changes in `config json files` file and make sure that the database exists.
You need not have to create the collection. When the server starts, it will automatically create the necessary collection and initialize them.

##### Libraries/packages used:
  - [Go][Go]
  - [gin-gonic][gin-gonic]
  - [bcrypt][bcrypt]
  - [uuid][uuid]
  - [mongo][mongo]

Feel free to post questions and star.

[gin-gonic]: <https://github.com/gin-gonic/gin>
[bcrypt]: <golang.org/x/crypto/bcrypt>
[Go]: <https://github.com/golang/go>
[uuid]: <https://github.com/satori/go.uuid>
[mongo]: <gopkg.in/mgo.v2>
