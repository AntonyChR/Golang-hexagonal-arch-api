## Install dependencies

```
go get -d ./..
```

To use the MySQL repository, uncomment dbConfig in cmd/main.go and change the value of the Host field with your local host IP
```golang
	dbConfig := repositorie.MysqlConfig{
		Database: "test",
		Table:    "todos",
		Host:     "192.168.18.46",
		Port:     "3306",
		User:     "root",
	}

	todoService := application.TodoService{
		Repo: repositorie.InitializeMysqlRepository(dbConfig),
		//Repo: repositorie.InitializeMemoryRepository("todos", "test"),
	}

```
## Run
```
go run cmd/main.go
```
