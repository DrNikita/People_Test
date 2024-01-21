         ---Task---
Реализовать сервис, который будет получать по апи ФИО, из открытых апи обогащать
ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в
БД. По запросу выдавать инфу о найденных людях. Необходимо реализовать следующее
1. Выставить rest методы для получения данных с различными фильтрами и пагинацией
2. Для удаления по идентификатору
3. Для изменения сущности
4. Для добавления новых людей в формате

   {
         "name": "Dmitriy",
         "surname": "Ushakov",
         "patronymic": "Vasilevich" // необязательно
   }

2. Корректное сообщение обогатить
   1. Возрастом - https://api.agify.io/?name=Dmitriy
   2. Полом - https://api.genderize.io/?name=Dmitriy
   3. Национальностью - https://api.nationalize.io/?name=Dmitriy
3. Обогащенное сообщение положить в БД postgres (структура БД должна быть создана
   путем миграций)
4. Покрыть код debug- и info-логами
5. Вынести конфигурационные данные в .env

         ---Instructions---
1) create db "people"
2) migrate -path internal/db/migration/ -database "postgresql://postgres:postgres@localhost:5432/people?sslmode=disable" -verbose up
3) migrate -path internal/db/migration/ -database "postgresql://postgres:postgres@localhost:5432/people?sslmode=disable" -verbose down