# Go Base Structure

This is the base source code that provide the basic structure of a Golang Project. Developer can rely on this to build a simple API application.

```
/api: contains backend code
	/cmd
		/<app_name>
			/main.go: entry point of project
	/internal: private application and library code.
		/api
			/router: handle routing
			/rest: contains all restful apis
				/erros.go
			/graphql: contains all graphql apis (not use in this project)
			/middleware: contains middleware code
		/config: contains code to init server + init db
		/controller
		/repository
		/models
		/docs: api documentation
		/pkg: utility codes: respondJSON, etc….
	/data/migrations: contains db migrations file
	/pkg: library code that's ok to use by external applications
/web: contains client code 
/go.mod
/go.sum
/dockerfile
```

## 1. Summary

- Programming Language: Golang
- Database: Postgresql
- Deployment: Docker, Linux
- Tools: Goland, Git

## 3. Deployment

- This project can be deployed by Docker to Linux server at: http://localhost:3000/

### #Deployment process

```
make setup
make boilerplate
make run
```

### #Testing

```
make test
```
