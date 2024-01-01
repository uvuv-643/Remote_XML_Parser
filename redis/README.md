## Redis: применение

Redis применяется для оптимизации выполнения операции обновления базы данных.
Данные, которые были получены с удаленного сервера и были записаны в базу данных,
сохраняются также и в Redis с определенным TTL ([.env.example](../server/.env.example))

При повторном получении данных с удаленного ресурса проверяется соответствие
данных в Redis (оперативная память) и в случае их совпадения запись пропускается
и не выполняется обращение к хранилищу. В противном случае cache-miss и перезапись в базу данных.

Используется для создания новых записей, редактирования и удаления текущих записей.

## Redis: контракт

Информация в Redis хранится в следующе  м виде:
В качестве ключа выступает составная строка вида ```<package_name>.<type>.<uid>```
Например, ```dbs.SDNAddress.1234```.
В качестве значения - объект, сериализованный в формат JSON.
Список первичных ключей хранится в кеше с ключём вида ```<package_name>.<type>```.

## Redis: производительность

Результаты замеров производительности ```GET /update```:

* База данных пуста, данные заполняются заново:
    ```
    server  | Request latency": "418.404142s"
  ```

* База данных заполнена, данные обновляются полностью без кеша:
    ```
    server  | Request latency": "91.460347s"
  ```
  
* База данных заполнена, данные обновляются полностью из кеша:
    ```
    server  | Request latency": "25.714319s"
  ```

В зависимости от объёма обновляемых данных, время будет линейно увеличиватся, и 
кеш промахи не существенно влияют на показатель скорости выполнения запроса.
Обновление половины данных в базе теоретически должно занимать приблизительно 50s.