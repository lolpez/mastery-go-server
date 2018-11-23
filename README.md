# mastery-go-server
Small go HTTP server for practicing purposes.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for Development environment.

### Prerequisites

What things you need to install the software and how to install them.

* Install [Golang](https://golang.org/) for Windows.

That's all! :D

### Installing

A step by step series of examples that tell you have to get a Development env. running.

* Clone the project to your src Golang project folder. (Ex Windows. C:\\Users\YOUR USER\go\src)
* Inside the project folder, open a command prompt and type:
```
go get -u github.com/golang/dep/cmd/dep
```
```
dep init -v
```
```
dep ensure -v 
```

* Execute the http server:
```
go run main.go
```

* Open an Internet Browser and go to [Localhost port 9000, route documents](http://localhost:9000/documents)

### Output

* The program will return all MD5 checksum, names and sizes of the files located in the "files" folder in JSON format.

Example:
Making a GET request to [http://localhost:9000/documents](http://localhost:9000/documents),
The server will respond with a dictionary of all files located in "files" folder in JSON format:

```
[{"ID":"af526914b1724469467f85ae09e90f3e","Name":"javascript file.js","Size":27},{"ID":"9bd8007769ac2fa6077b9fee7561881c","Name":"json file.json","Size":34},{"ID":"661668af31d00980cac52b7509ef5d14","Name":"pug file.pug","Size":14},{"ID":"b10a8db164e0754105b7a99be72e3fe5","Name":"txt file.txt","Size":11}]
```

## Authors

* **Luis Daniel** - *Developer*

## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE) file for details

## Acknowledgments

* Golang
* Videogame gods
