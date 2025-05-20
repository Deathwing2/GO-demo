package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ping(address string, timeout time.Duration) (time.Duration, error) {
	start := time.Now()
	client := http.Client{
		Timeout: timeout * time.Second,
	}
	req, err := http.NewRequest(
		"GET", address, nil,
	)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	defer resp.Body.Close()

	fmt.Println("Пинг к", address, "успешен!", time.Since(start))
	return time.Since(start), err
}

func main() {
	var deniedError, deniedTimeout int
	var timeout, pingMin, pingMax time.Duration
	var amountSuccesPing int = 0
	var sumSuccesPing time.Duration = 0
	var address string
	var count, s int

	flag.String("address", "google.com", "Укажите адрес сайта")
	flag.Int("s", 5, "Укажите размер тайм-аута в секундах")
	flag.Int("count", 10, "Введите количество запросов")
	flag.Parse()
	address = os.Args[1]
	numberString2 := os.Args[2]
	s, err := strconv.Atoi(numberString2)
	if err != nil {
		fmt.Println("Ошибка преобразования аргумента в число:", err)
		return
	}
	numberString3 := os.Args[3]
	count, err = strconv.Atoi(numberString3)
	if err != nil {
		fmt.Println("Ошибка преобразования аргумента в число:", err)
		return
	}
	pingMin = 10 * time.Second
	for i := 0; i < count; i++ {
		go ping(address, time.Duration(s))
		result, err := ping(address, timeout)
		if err != nil {
			deniedError = deniedError + 1
		}
		if err, ok := err.(net.Error); ok && err.Timeout() {
			deniedTimeout = deniedTimeout + 1
		}
		if result > 0 {
			if pingMin > result {
				pingMin = result
			}
			if pingMax < result {
				pingMax = result
			}
			amountSuccesPing = amountSuccesPing + 1
			sumSuccesPing = sumSuccesPing + result
		}
	}
	fmt.Println("Всего отправлено запросов на сервер", count)
	fmt.Println("Самый быстрый запрос вернулся через", pingMin)
	fmt.Println("Самый медленный запрос -", pingMax)
	fmt.Println("Среднее значение времени отклика", sumSuccesPing/time.Duration(amountSuccesPing))
	fmt.Println("Неудачных запросов", deniedError)
	fmt.Println("Неудачных запросов по тайм-ауту", deniedTimeout)
}
