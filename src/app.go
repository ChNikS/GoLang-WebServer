package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
	"encoding/json"
	"errors"
)

type ResponseItem struct {
	From string
	To string
	Value float64
}

//structure for pasrsing data from https://www.cbr-xml-daily.ru/
type BankData struct {
	Msg interface{} `json:"Valute"`
}


func main() {
	//init gin web server
	r := gin.Default()
    //routing
	r.GET("/exchangeRate", func(c *gin.Context) {
		
		//check if correct response
		if data, err := getExchangeRate(c); err != nil {
			c.JSON(200, gin.H{
				"Error": err.Error(),
			})	
		} else {
			c.JSON(200, gin.H{
				"Rates": data,
			})
		}
	})

	//run server with proper port setting
	r.Run(":9000")
}

func getExchangeRate(c *gin.Context) ([]ResponseItem, error) {
	var responseArray []ResponseItem
	
	//parse data
	from := string(c.Request.URL.Query().Get("from")) // -> rub (check if rub)
	to := c.Request.URL.Query()["to"]
	
	if from == "" || to == nil {
		//exception required params are missing
		return responseArray, errors.New("required params are missing")
	}

	if from != "rub" {
		return responseArray, errors.New("unspported operation, try to use from=rub")
	}
	
	//get data from bank
	resp, err := resty.R().Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return responseArray, errors.New("Cannot connect to bank server")
	}
	
	//parse data from bank
	var env BankData
	if err := json.Unmarshal([]byte(resp.String()), &env); err != nil {
		return responseArray, errors.New("Cannot parse data from bank server. Contact administrator")
	}

	dataMap := env.Msg.(map[string]interface{})

	for index := range to {
    	//check if currency code is correct
        if dataMap[to[index]] != nil {
	        currency := (dataMap[to[index]]).(map[string]interface{})
	        var responseTo string
	        var responseValue float64

	        if str, ok := currency["Name"].(string); ok {
			    responseTo = str
			} else {
				//exception
			    return responseArray, errors.New("Cannot parse data from bank server. Contact administrator")
			}

			if str, ok := currency["Value"].(float64); ok {
			    responseValue = str
			} else {
			    //exception
			    return responseArray, errors.New("Cannot parse data from bank server. Contact administrator")
			}

	        response := ResponseItem{
		    	From: "Рубль",
		    	To: responseTo,
		    	Value: responseValue,
		    }
		    responseArray = append(responseArray, response)
		}
    } 

    return responseArray, nil
}







