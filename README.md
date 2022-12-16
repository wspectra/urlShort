# :heavy_check_mark: urlShort
**urlShort** - сервис, который предоставляет API по созданию сокращённых ссылок следующего формата:
- Сокращенная ссылка состоит из 10 символов
- Cимволы состоят из латинского алфавита в нижнем и верхнем регистре, цифр и символа '_'
# Сервис принимает следующие запросы по http:
1. Метод `POST`, который сохраняет оригинальный URL в базе и возврает сокращённый
----
* **URL**: /post/
*  **URL Params**: None
* **Data Params**
   **Required:**
   ```json
  {"Url":"ссылка"}
  ```
* **Success Response:**
  * **Code:** 200 <br />
    **Content:** `{"Status":"success",
    "Message":"rtr"}`
    **Context:** `Success`

* **Error Response:**
  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Status":"fail",
    "Message":"invalid character 'U' looking for beginning of value"}`
    **Context:** `invalid JSON`
  OR

   * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Status":"fail",
    "Message":"Key: 'RequestStruct.Url' Error:Field validation for 'Url' failed on the 'required' tag"}n`
    **Context:** `wrong JSON structure`

  OR

   * **Code:** 400 BAD REQUEST <br />
    **Content:** `{"Status":"fail",
    "Message":"parse \"qewerfwet\": invalid URI for request"}`
    **Context:** `invalid URL`

  OR
     * **Code:** 500 INTERNAL SERVER ERROR <br />
      **Content:** `{"Status":"fail",
      "Message":"parse \"qewerfwet\": invalid URI for request"}`
      **Context:** `error during post in database`

2. Метод Get, который принимает сокращённый URL и выполняет редирект на оригинальный URL
----
* **URL**: /укороченная_ссылка
*  **URL Params**: None
* **Data Params**
   **Required:** None

* **Success Response:**
  * **Code:** 303 SEE OTHER <br />
    **Content:** `{"Status":"fail",
    "Message":<a href="исходная ссылка">See Other</a>."}`

* **Error Response:**
  * **Code:** 404 NOT FOUND <br />
    **Content:** `{"Status":"fail",
    "Message":"long Url not found"}`
# Хранилище
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;В качестве хранилища возможно использование in-memory решения и базы данных - postgresql. Какое хранилище использовать указывается параметром при запуске сервиса.
![image](https://user-images.githubusercontent.com/80648065/155390687-8f427f70-a635-4e98-98f9-ee1aca628551.png)
# Usage
По умолчанию поднимается контейнер в котором работает сервис

    make

Выполняются тесты

    make test

Завершить работу контейнеров

    make clean
    
## Other
**Author:**
*[wspectra](https://github.com/wspectra)*
