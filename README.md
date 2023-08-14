# products-data-manager
## gRPC server and client which allows sync and recieve your products db data with another system

## Stack:

```
1. GO
2. gRPC
3. MongoDB
```

## Description:
This package contains server and client using gRPC connection
This is example of creation simple gRPC server which allows sync and recieve your products db data with another system.
## Server

TO update your protobuf use command below.

```protoc proto/products_manager.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.```

### Setting up:

Set on your `.env` file variables 
DB_URI
DB_USERNAME
DB_PASSWORD
DB_DATABASE_NAME
SERVER_PORT

To run MongoDB in Docker Container - use command below

```docker run -d -p {DB_PORT:DB_PORT} --name {DB_DATABASE_NAME} -e MONGO_INITDB_ROOT_USERNAME={DB_USERNAME} -e MONGO_INITDB_ROOT_PASSWORD={DB_PASSWORD} mongo:latest```

### Running
To run server use command

```go run cmd/main.go```

## Client
### Implemented functions:

```
1. Fetch(url) - use this to get external products data and update you data in MongoDB
2. List(paginationOptions, sortingOptions) - getting sorted data
``` 

### Example usage

```go
package main

import (
	"fmt"

	grpc_client "github.com/salesforceanton/products-data-manager/pkg/client"
)

func main() {
	client, err := grpc_client.NewClient(9000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := client.List(nil, &products_manager.SortOptions{Field: "price", Order: products_manager.DESC})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v", result)
}

```

```go
package main

import (
	"fmt"

	grpc_client "github.com/salesforceanton/products-data-manager/pkg/client"
)

func main() {
	client, err := grpc_client.NewClient(9000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = client.Fetch("http://164.92.251.245:8080/api/v1/products")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

```
