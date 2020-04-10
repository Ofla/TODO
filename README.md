# TODO

Simple TODO List application with GRPC and REST API


## Getting started


## Setup 

1. Install go

    See https://golang.org/doc/install

2. Install gRPC tools:

    ```
    $ go get google.golang.org/grpc
    $ go get -a github.com/golang/protobuf/protoc-gen-go
    ```

3. Install project with dependencies
   
   ```
   $ go get -u github.com/Ofla/TODO
   ```

4. Generate needed files 
   
   ```
   make generate
   ```
5. Run GRPC Server

    ```
    make runGRPCServer
    ```

6. Run GRPC Client 
    
    ```
    make runGRPCClient
    ```

7. Run Http Server 
   
    ```
    make HTTPServer
    ```
    
   * then you can test the end-points using swagger-ui 
  
    ```
    localhost:9977/swagger-ui/
    ```
8. Testing

    ```
    make Test
    ```
  