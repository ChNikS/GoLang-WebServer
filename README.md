# GoLang Web-server
Web-app project based on Gin Web Framework.<br/>
Starts on 9000 port.<br/>
Request example:<br/>
http://localhost:9000/exchangeRate?from=rub&to=USD&to=EUR&to=JPY
Response example: 
```json
{
    "Rates": [
        {
            "From": "Рубль",
            "To": "Доллар США",
            "Value": 65.5906
        },
        {
            "From": "Рубль",
            "To": "Евро",
            "Value": 76.2294
        },
        {
            "From": "Рубль",
            "To": "Японских иен",
            "Value": 57.8273
        }
    ]
}
```
