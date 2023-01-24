# golang-testing-task

Task:
Build API to check if a list of given numbers are primes. API should only have a single
'POST' endpoint to accept the request.
The request body must be a slice of integers, otherwise return an error.
If the request is valid, return a slice of booleans either the number was prime or not.


<img width="613" alt="image" src="https://user-images.githubusercontent.com/61024343/214285372-dc87604b-7ee0-46ee-b2e3-36d6cb43808c.png">


[Vadim Kravets]
Проект сделан с учетом дальнейшего расширения, архиктуру выбрал больше похожую на DDD
Бизнес логика в сервисах, хендлеры как прокладки, хелперы вынесены в отдельные папки

Инструкция:
(Makefile)

1. git clone https://github.com/Kravets21/golang-testing-task.git
2. cd golang-testing-task

3. (Если ГО установлен тогда make run)
Если нет, можно использовать докер контейнер

поднять проект на докере:
make build

закрыть проект:
make down

endpoint: http://localhost:8080

