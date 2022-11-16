# internship_backend_2022

<a href=https://github.com/avito-tech/internship_backend_2022>Текст тестового задания Авито 2022</a>

Для запуска проекта:
1. создать файл docker/.env (образец docker/.env.sample, можно указать ТОЛЬКО MYSQL_DATA_PATH свой, а остальное оставить как есть)
2. в корне проекта выполнить команду make start запустятся 2 docker контейнера
3. импортировать дамп базы данных, который находится docker/mysql/user_balance_2022-10-24.sql

swagger файл будет доступен по адресу http://localhost:9000/swagger/index.html если не менялись значения ENV

Примеры запросов:
1. пополнить баланс пользователя:
curl -H 'Content-Type: application/json' --data '{"user_id":2, "type":1, "money":25000, "service_id":1, "service_name":"Пополнение банковской картой", "order_id":12451252}' http://127.0.0.1:9000/balance/add

2. добавить транзакцию оплаты услуги с заморозкой средств:
curl -H 'Content-Type: application/json' --data '{"user_id":5, "type":0, "money":20000, "service_id":2, "service_name":"турбопродажа", "order_id":12451261}' http://127.0.0.1:9000/transaction/add

3. подтверждение транзакции:
curl -H 'Content-Type: application/json' --data '{"user_id":5, "money":20000, "service_id":2, "service_name":"турбопродажа", "order_id":12451261}' http://127.0.0.1:9000/transaction/confirm

4. отмена транзакции и возврат средств на баланс пользователя:
curl -H 'Content-Type: application/json' --data '{"user_id":2, "money":20000, "service_id":2, "service_name":"турбопродажа", "order_id":12451262}' http://127.0.0.1:9000/transaction/cancel

5. получить баланс пользователя:
curl -H 'Content-Type: application/json' --data '{"user_id":2}' http://127.0.0.1:9000/balance/get

6. получить список транзакций пользователя:
    6.1. сортировка по умолчанию, первая страница:
    curl -H 'Content-Type: application/json' --data '{"user_id":2}' http://127.0.0.1:9000/transaction/list

    6.2. добавляем пагинацию (нумерация страниц начинается с нуля):
    curl -H 'Content-Type: application/json' --data '{"page_num":1, "id":2,"user_id":2}' http://127.0.0.1:9000/transaction/list

    6.3. добавляем сортировку по одному из полей (money - деньги и confirmed - дата подтверждения), варианты сортировки ASC или DESC
    curl -H 'Content-Type: application/json' --data '{"page_num":1, "id":2,"user_id":2,"filter":{"order_by":"money", "order_direction":"DESC"}}' http://127.0.0.1:9000/transaction/list

7. получить ссылку на файл с отчетом для бухгалтерии:
curl -H 'Content-Type: application/json' --data '{"month":10, "year":2022}' http://127.0.0.1:9000/accounting/list


Возникавшие вопропросы и способы из решения:
1. не совсем понятно было резервирование средств с основного баланса на отдельном счете, я реализовал это добавив к транзакции метод подтверждения или отмены, которые либо подтверждают транзакцию и средства больше нельзя вернуть на баланс пользователя или отменяют транзакцию и возвращают срадства на баланс пользователя
2. возник вопрос какой тип использовать для денег, решил использовать целочисленный, так же принимать и отдавать в таком же виде. Пример: 12050 = 120 рублей 50 копеек

Проблемы:
1. Unit тесты не работают, не доделаны, не хватило времени честно говоря не набита рука
2. Уже в послендний момент заметил слабое место в логике, которое может, при передачи ошибочных данных привести к ошибкам в рассчетах. А именно: поле type в структурах, которое отвечает за тип транзакции 0 - оплата, 1 - пополнение. Нужно не передавать ее, а жестко привязать к методам в коде. Если это метод с оплатой, то type == 0, если пополнение баланса пользователя, то type == 1. В моей реализации, есть валидация, которая пропустит только 0 или 1, а вот передать несоответствующий тип получится, что приведет к ошибке в расчетах.
