# Learning GRPC with golang + gin
- golang
- grpc
- gin 

In this project I created a project with grpc, go and gin, just to learn how microservices work. It is a fun project and it has three parts
One is the API part and two are the service part.
The service parts are - auth-svc, product-svc

The client will communicate with the api part and the api part will communicate with the service. The client will not connect directly to the service.

For now, the data can be saved in json format using postman by running this project.


## Installation

this apps requires go, gin, grpc

clone this project and Install the dependencies and devDependencies and start the server.
go to ```auth-svc``` (auth-svc)[https://github.com/asadlive84/school/learn-mircroserive-grpc-golang-gin/tree/main/auth-svc]   folder and run this commaand in your terminal
```go
make build
```
Verify the deployment by navigating to your server address in
your postman.
postman url ( METHOD is ```POST``` )
```sh
localhost:3000/auth/register
```
>> insert a new user
```json
{
    "email":"asad1@me.com",
    "password":"123"

}
```

>> insert a new product

postman url (METHOD is ```POST```)

```sh
localhost:3000/product/create
```

```sh
{
    "name":"Nokia",
    "description":"nokia description",
    "is_active":true,
    "user_id":"a8cf63c6-edd9-4024-84fc-6ecf43722c3d"

}
```




