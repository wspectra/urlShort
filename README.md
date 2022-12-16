# :heavy_check_mark: urlShort
**urlShort** - сервис, который предоставляет API по созданию сокращённых ссылок следующего формата:
- Сокращенная ссылка состоит из 10 символов
- Cимволы состоят из латинского алфавита в нижнем и верхнем регистре, цифр и символа '_'
# Сервис принимает следующие запросы по http:
1. Метод `POST`, который сохраняет оригинальный URL в базе и возврает сокращённый
----
* **URL**: /post
*  **URL Params**: None
* **Data Params**
   **Required:**
   ```json
  {  "Url":"ссылка"  }
  ```
# Success Response:
  * **Code:** 200 <br />
    **Context:** `Success`<br />
    **Content:** 
    ```json
    {
      "Status":"success",
      "Message":"rtr"
    }
    ```


# Error Response:
  * **Code:** 400 BAD REQUEST <br />
    **Context:** `invalid JSON` <br />
    **Content:** 
    ```json
    {
      "Status":"fail",
      "Message":"invalid character 'U' looking for beginning of value"
    }
    ```
    
  OR

* **Code:** 400 BAD REQUEST <br />
  **Context:** `wrong JSON structure`<br />
**Content:**
   ```json
     {
       "Status":"fail",
       "Message":"Key: 'RequestStruct.Url' Error:Field validation for 'Url' failed on the 'required' tag"
     }
    ```

  OR

   * **Code:** 400 BAD REQUEST <br />
     **Context:** `invalid URL`<br />
    **Content:**
  ```json
  {
    "Status":"fail",
    "Message":"parse \"qewerfwet\": invalid URI for request"
  }
   ```

  OR
     * **Code:** 500 INTERNAL SERVER ERROR <br />
  **Context:** `error during post in database`<br />
      **Content:**
  ```json
  {
    "Status":"fail",
    "Message":"parse \"qewerfwet\": invalid URI for request"
  }
  ```

2. Метод Get, который принимает сокращённый URL и выполняет редирект на оригинальный URL
----
* **URL**: /укороченная_ссылка
*  **URL Params**: None
* **Data Params**
   **Required:** None

* **Success Response:**
  * **Code:** 303 SEE OTHER <br />
    **Content:**
  ```json
  {
    "Status":"fail",
    "Message":<a href="исходная ссылка">See Other</a>."
  }
  ```

* **Error Response:**
  * **Code:** 404 NOT FOUND <br />
    **Content:**
  ```json
  {
    "Status":"fail",
    "Message":"long Url not found"
  }
  ```
  

# Хранилище
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;В качестве хранилища возможно использование in-memory решения и базы данных - postgresql. Какое хранилище использовать указывается параметром при запуске сервиса.<br />
  ![image](https://user-images.githubusercontent.com/75119633/208161156-aa2dbb36-be9e-42c5-b165-48b081c415cd.jpg)
# Usage
По умолчанию поднимается контейнер в котором работает сервис

    make

Выполняются тесты

    make test

Завершить работу контейнеров

    make clean
    
## Other
**Author:**
:pig:**[wspectra](https://github.com/wspectra)**
