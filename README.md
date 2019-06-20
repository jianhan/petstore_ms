# Petstore Micro Service
Implementation of [petstore](https://petstore.swagger.io/) using micro service architecture. It only has two end points, one for creating pet, the other is updating pet.

### Over all architecture
It contains 2 micro services registered in [consul](https://www.consul.io/) and communicated via RPC.
The flow is simple:
1. request comes in via api service
2. api service call store service via RPC call
3. store service will perform the businss logic and save/update data into mysql.
4. store service return the response and api service return back to client.

Services are dockrized via docker compose, the micro web is just a tool to test RPC end points for individual service, the rest defined in docker compose yml file are all essential. 

### Design principles
To ensure code is testable and well written, following design princple was implemented in mind.
1. SOLID
2. KISS
3. unix design philosophy
4. programming to interface not implementation 

### Tools & package used
1. GO 1.12.5
2. go module for dep management
3. go micro for RPC framework
4. go mock for mocking interface to test
5. golangci-lint for linting
6. gorilla mux for api end point development
7. makefile for build , test, local development etc...
8. micro web for building, testing RPC calls.
9. docker
10. mysql 5.7

### Start all services
Clone the repo
`git clone git@github.com:jianhan/petstore_ms.git && cd petstore_ms` 
run via docker compose
`docker-compose up -d && docker-compose ps`

### Testing end points
Currently there are two end points:

`http://localhost:9090/api/v2/pet` to create new pet, it has a very simple validation built into it,
pet can not be empty, pet name can not be empty, pet status can not be empty and must be valid.

To test create pet successfuly:
```
curl -d '{"pet":  {"name": "pet name 1", "photo_urls": ["url1", "url2"], "status": "pending"}}' -H "Content-Type: application/json" -X POST http://localhost:9090/api/v2/pet
```

To test pet must be require:
```
curl -d '{}' -H "Content-Type: application/json" -X POST http://localhost:9090/api/v2/pet
```

To test pet name must be required:
```
curl -d '{"pet": {}}' -H "Content-Type: application/json" -X POST http://localhost:9090/api/v2/pet
```

To test pet status must be valid
```
curl -d '{"pet":  {"name": "pet name 1", "photo_urls": ["url1", "url2"], "status": "invalid status"}}' -H "Content-Type: application/json" -X POST http://localhost:9090/api/v2/pet
```

`http://localhost:9090/api/v2/pet/{id}` to update pet. Id must be a valid id in mysql database, same validation for name and status, to test successfuly update pet 
```
curl -d '{"id": 1,"name": "name updated","status": "available"}' -H "Content-Type: application/json" -X POST http://localhost:9090/api/v2/pet/1
```

### Unit test
from root of project, `cd srv/store` then run `make test` to test store service.
same way to test api service. Only two tests was written.

### Note
I only tested on mac an linux, both works fine.