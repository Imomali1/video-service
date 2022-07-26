# api
This is CRUD API for video service project
## Usage

### Download libraries which are required to run this program

1. Download [gin](github.com/gin-gonic/gin) web framework. 
```sh
go get github.com/gin-gonic/gin 
```
2. Download [postgresql](https://www.postgresql.org/download/) open source object-relational database system and add to path.
3. Include it in your project
```sh
go get github.com/lib/pq 
```
4. Download [yaml](gopkg.in/yaml.v2) package that implements YAML support for the Go language.
```sh
go get gopkg.in/yaml.v2
```

### Run `setup.go` file to initialize database
```sh
 go run cmd/setup.go
```

### Run `main.go` file to use the CRUD API
```sh
go run cmd/setup.go
```

### In this project, there are functions like:
 1. `GET("/videos")            // gets all videos that are saved`
 2. `POST("/videos")           // creates a video`
 3. `GET("/videos/:id")        // gets video by id`
 4. `PUT("/videos/:id")        // updates video by id`
 5. `DELETE("/videos/:id")     // deletes video by id`
 
 ### Directory structure
 
 `
├───.idea  
├───cmd
|   ├───main.go
|   └───setup.go
├───configs
|   └───configs.yaml
├───entity
|   └───video.go
└───pkg
    ├───controller
    |   └───video-controller.go
    ├───db
    │   ├───SQL
    |   |    └───table.sql
    |   ├───db-connection.go
    |   └───db-functions.go
    ├───service
    |   └───video-service.go
    └───utils
        ├───convert
        |   └───strtoint.go
        └───reader
            └───readfile.go
 `
