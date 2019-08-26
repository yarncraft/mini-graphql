# mini-graphql 🧩

_type safe, lightning fast & minimal GraphQL server template featuring_

- [gqlgen](https://gqlgen.com/feature-comparison/)  the GraphQL toolkit 
- [gin](https://gin-gonic.com/)                     the web-server      
- [gorm](https://gorm.io/)                          the ORM             

<img src="https://i.imgur.com/XrC77wp.png" border="0">



### get started

Make sure the [Go binaries](https://golang.org/doc/install) and [dep manager](https://golang.github.io/dep/docs/installation.html) are installed first! 

```
dep install
go run src/cmd/migration/up.go
```

Now we can start running the server locally, or with Docker:

```
go run src/cmd/server/main.go

./compile.sh
docker build -t gql .
docker run -it -p 8080:8080 gql
```

### final docker size

The 'main' binary is copied into a builder container, then stripped and compressed and transferred to a scratch image! 

The final executable binary results in a stunning 4MB image!!!

<img src="https://i.imgur.com/CojyIGu.png" border="0">


