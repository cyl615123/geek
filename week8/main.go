package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

// dial wraps DialDefaultServer() with a more suitable function name for examples.
func dial() (redis.Conn, error) {
	return redis.Dial("tcp", "127.0.0.1:6379")
}

func main() {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	singleBytes := []int{10, 20, 50, 100, 200, 1000, 5000}
	totalBytes := []int{10000, 20000, 50000, 100000, 500000}
	for _, t := range totalBytes {
		for _, s := range singleBytes {
			SetBytes(c, s, t)
		}
	}
}

func SetBytes(c redis.Conn, size, total int) {
	c.Do("flushdb")
	memBeforeSet, _ := c.Do("info", "memory")
	value := make([]byte, size)
	for i := range value {
		value[i] = 'a'
	}
	for i := 0; i < total/size; i++ {
		s := strconv.Itoa(i)
		c.Do("SET", s, value)
	}
	memAfterSet, _ := c.Do("info", "memory")

	memBeforeSetTotal := getUsedMemory(string(memBeforeSet.([]uint8)))
	memAfterSetTotal := getUsedMemory(string(memAfterSet.([]uint8)))
	fmt.Println("total=", total, ",bytes=", size,
		",avg =", (memAfterSetTotal-memBeforeSetTotal)*size/total,
		",memBeforeSet:", memBeforeSetTotal,
		",memAfterSet:", memAfterSetTotal)
}

func getUsedMemory(content string) int {
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		kv := strings.Split(line, ":")
		if kv[0] == "used_memory" {
			r, _ := strconv.Atoi(kv[1])
			return r
		}
	}
	return 0
}
