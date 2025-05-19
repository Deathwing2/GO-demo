package main

import (
	"fmt"
	"net"
	"time"
)

func ping(address string, timeout time.Duration) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout("ip:icmp", address, timeout)
	if err != nil {
		fmt.Println("Ошибка", err)
		return 0, err
	}
	defer conn.Close()
	fmt.Println("Пинг к", address, "успешен!", time.Since(start))
	return time.Since(start), err
}

func main() {
	var deniedError int
	var deniedTimeout int
	var pingMin time.Duration
	var pingMax time.Duration = 0
	var amountSuccesPing int = 0
	var sumSuccesPing time.Duration = 0
	var address string
	var count int
	var s time.Duration
	// fmt.Println("Введите название сайта, который будем пинговать")
	// fmt.Scan(&address)
	address = "google.com"
	fmt.Println("Введите желаемое количество запросов")
	fmt.Scan(&count)
	fmt.Println("Введите желаемое время тайм-аута")
	fmt.Scan(&s)
	timeout := s * time.Millisecond

	for i := 0; i < count; i++ {
		result, err := ping(address, timeout)
		if err != nil {
			deniedError = deniedError + 1
		}
		// if err == dial ip:icmp: lookup google.com: i/o timeout {
		// 	deniedTimeout = deniedTimeout + 1
		// }
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

		time.Sleep(3 * time.Second)
	}
	fmt.Println("Всего отправлено запросов на сервер", count)
	fmt.Println("Самый быстрый запрос вернулся через", pingMin)
	fmt.Println("Самый медленный запрос -", pingMax)
	fmt.Println("Среднее значение времени отклика", sumSuccesPing/time.Duration(amountSuccesPing))
	fmt.Println("Неудачных запросов", deniedError)
	fmt.Println("Неудачных запросов по тайм-ауту", deniedTimeout)
}
