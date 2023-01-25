# golang-testing-task

Task:
Build API to check if a list of given numbers are primes. API should only have a single
'POST' endpoint to accept the request.
The request body must be a slice of integers, otherwise return an error.
If the request is valid, return a slice of booleans either the number was prime or not.


<img width="613" alt="image" src="https://user-images.githubusercontent.com/61024343/214285372-dc87604b-7ee0-46ee-b2e3-36d6cb43808c.png">


[Vadim Kravets]
The Project architecture (DDD) was made for future extending of its code base. 
Business logic is at services, handlers just as "proxy", helpers are located outside in additional folders

Instruction:
(Makefile)

1. git clone https://github.com/Kravets21/golang-testing-task.git
2. cd golang-testing-task

3. (if Golang is downloaded on your PC then type "make run")
If Golang isnt on your PC then you should work with docker containers:

up the project:
make build

down the project:
make down

to run tests:
make test (without docker)
make test-docker (from docker)

endpoint: http://localhost:8080

You can test app with Postman

