package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	AuthUtil "github.com/salimkun/Efishery-Test/Auth/common/util"
	"github.com/salimkun/Efishery-Test/Fetch/common/util"
	"github.com/salimkun/Efishery-Test/Fetch/model"
)

func GetResource(c *gin.Context) {

	response, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var responseObject []*model.Resource
	json.Unmarshal(responseData, &responseObject)

	// check redis
	var currency float64
	client := AuthUtil.SetUpRedis()
	countryVal, err := client.Get("currency").Result()
	if err == redis.Nil {
		currency, err = getCurrencyNow()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// set redis
		err = client.Set("currency", fmt.Sprintf("%.2f", currency), 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		currency, _ = strconv.ParseFloat(countryVal, 64)
	}

	for idx, i := range responseObject {
		val, _ := strconv.ParseFloat(i.Price, 64)
		lastCurrency := float64(val) / currency
		responseObject[idx].Price = string(fmt.Sprintf("%.2f", lastCurrency))
	}

	c.JSON(http.StatusOK, gin.H{"data": responseObject})
}

func AgregateResource(c *gin.Context) {

	response, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var responseObject []*model.Resource
	json.Unmarshal(responseData, &responseObject)

	var result []*model.AgregateResource
	// var obj *model.AgregateObj
	keys := make(map[string]int)
	for _, i := range responseObject {
		price, _ := strconv.Atoi(i.Price)
		size, _ := strconv.Atoi(i.Size)

		tm := time.Now()
		if i.Timestamp != "" {
			k, err := strconv.ParseInt(i.Timestamp, 10, 64)
			if err != nil {
				panic(err)
			}
			tm = time.Unix(k, 0)
		}

		resource := &model.AgregateResource{
			AreaProv: i.AreaProv,
			Price: model.AgregateObj{
				Maximum:   int64(price),
				Mininum:   int64(price),
				Median:    util.GetMedian([]float64{float64(price)}),
				Avg:       util.GetAvg([]float64{float64(price)}),
				ArrayData: []float64{float64(price)},
			},
			Size: model.AgregateObj{
				Maximum:   int64(size),
				Mininum:   int64(size),
				Median:    util.GetMedian([]float64{float64(size)}),
				Avg:       util.GetAvg([]float64{float64(size)}),
				ArrayData: []float64{float64(size)},
			},
			DateResource: tm.String(),
		}

		if j, ok := keys[i.AreaProv]; ok {
			// set max price
			if int64(price) > result[j].Price.Maximum {
				result[j].Price.Maximum = int64(price)
			}

			// set min price
			if int64(price) < result[j].Price.Mininum {
				result[j].Price.Mininum = int64(price)
			}

			// set Median price
			result[j].Price.ArrayData = append(result[j].Price.ArrayData, float64(price))
			result[j].Price.Median = util.GetMedian(result[j].Price.ArrayData)
			result[j].Price.Avg = util.GetAvg(result[j].Price.ArrayData)

			// set max size
			if int64(size) > result[j].Size.Maximum {
				result[j].Size.Maximum = int64(size)
			}

			// set min price
			if int64(price) < result[j].Size.Mininum {
				result[j].Size.Mininum = int64(size)
			}

			// set median size
			result[j].Size.ArrayData = append(result[j].Size.ArrayData, float64(size))
			result[j].Size.Median = util.GetMedian(result[j].Size.ArrayData)
			result[j].Size.Avg = util.GetAvg(result[j].Size.ArrayData)

			result[j].AreaProv = i.AreaProv
		} else {
			// Unique key found. Record position and collect
			// in result.
			keys[i.AreaProv] = len(result)
			result = append(result, resource)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func getCurrencyNow() (float64, error) {
	url := "https://api.apilayer.com/exchangerates_data/convert?to=IDR&from=USD&amount=1"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("apikey", "Xz0bS7GbwzTmVidjtteaDekQ9SYUmgdX")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 {
		return 0, err
	}

	var responseObject *model.ConvertCurrency
	json.Unmarshal(body, &responseObject)

	return responseObject.Result, nil
}
