# httpServer

This repository implements http server through socket in go language.

Although the Go language has imbedded http server in net/http package, 
but I want to implement it by myself. It is so funny to do so.

In the development process, I hope that I can think deeply 
and be more familiar with the network programming.

### Usage

1.  Firstly, install the repository and build the main.go.


    ```bash
    $ go get github.com/lwlwilliam/httpServer
    $ go build main.go
    ```


2.  Run the command below.


    ```bash
    $ ./main.go
    ```
    

    After that, try to access [http://localhost:8000](http://localhost:8000) on the
    browser directly. Also, you can specify the host and port with flags. Just like this.


    ```bash
    $ ./main -h host -p port
    ```
