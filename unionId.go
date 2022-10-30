package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var nextNum int = 0
var serviceId string = ""

/**
 * 唯一id 16位 = 2 + 2 + 3 + 7 + 2
 * year(2)randId(2)serviceId(3)time(7)nextId(2)
 * */
func CreateUnionId() int64 {
	nextNum += 1
	if nextNum >= 100 {
		nextNum = 0
	}
	now := time.Now()
	yearId := fmt.Sprintf("%d", (now.Year()))[2:]
	timeId := fmt.Sprintf("%d", now.UnixMicro()/1e3)
	timeId = timeId[len(timeId)-7:]
	randId := ""
	randNum := rand.Intn(99)
	if randNum < 10 {
		randId = fmt.Sprintf("0%d", randNum)
	} else {
		randId = fmt.Sprintf("%d", randNum)
	}
	nextId := ""
	if nextNum < 10 {
		nextId = fmt.Sprintf("0%d", nextNum)
	} else {
		nextId = fmt.Sprintf("%d", nextNum)
	}
	finalId, _ := strconv.ParseInt(yearId+randId+serviceId+timeId+nextId, 10, 64)
	return finalId
}

func init() {
	serverIp := strings.Join(strings.Split(GetServerIp(), "."), "")
	serviceId = serverIp[len(serverIp)-3:]
}
