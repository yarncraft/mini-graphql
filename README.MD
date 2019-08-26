# mini-graphql 

_type safe, lightning fast & minimal GraphQL server template_

<img src="https://i.imgur.com/XrC77wp.png" border="0">


### dependencies
* gqlgen  : the GraphQL toolkit
* gin     : the web-server
* gorm    : the ORM


### get started

Make sure the Go binaries and dep manager are installed first!

```
dep install
go run src/cmd/server/main.go

or

./compile.sh
docker build -t gql .
docker run -it -p 8080:8080 gql
```

### final docker size

The main binary is copied into a builder container, then stripped and compressed and transferred to a scratch image! 

The final executable binary results in a stunning 4MB image 🖤

<img src="https://i.imgur.com/CojyIGu.png" border="0">
