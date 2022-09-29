package core

import (
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 响应结构体, 参数:data,code,msg,status
func MakeResponse(args ...any) (ResponseData, int) {
	now := time.Now()
	status := 200
	resData := ResponseData{
		Data:    nil,
		Code:    0,
		Msg:     "ok",
		SysTime: now.UnixMilli(),
	}
	for i, v := range args {
		switch i {
		case 0:
			{
				resData.Data = v
				break
			}
		case 1:
			{
				resData.Code = v.(int)
				break
			}
		case 2:
			{
				resData.Msg = v.(string)
				break
			}
		case 3:
			{
				status = v.(int)
				break
			}
		}
	}
	return resData, status
}

// 获取客户端ip
func GetClientIp(ctx *gin.Context) string {
	aliCDNRealIp := ctx.Request.Header.Get("ali-cdn-real-ip")
	xForwardedFor := ctx.Request.Header.Get("x-forwarded-for")
	xRealIp := ctx.Request.Header.Get("x-real-ip")
	xForwardFor := ctx.Request.Header.Get("x-forward-for")

	if aliCDNRealIp != "" {
		return aliCDNRealIp
	} else if xForwardedFor != "" {
		return xForwardedFor
	} else if xRealIp != "" {
		return xRealIp
	} else if xForwardFor != "" {
		return xForwardFor
	} else {
		return ctx.ClientIP()
	}
}

// 获取服务端id
func GetServerIp() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return ""
}

// 检查参数
func CheckParams(ctx *gin.Context, queryBind any) {
	if queryBind != nil {
		err := ctx.ShouldBind(queryBind)
		if err != nil {
			log.Println(err)
			panic(ErrorData{Code: CODE_PARAMS_MISSING})
		}
	}
}

// 签名
func MakeSig(app int64, nonce string, ts int, ttl int, data string, appScrectKey string) string {
	strArr := strings.Split(fmt.Sprintf("%d%s%d%d%s", app, nonce, ts, ttl, data), "")
	sort.Strings(strArr)
	sig := EncryptHMACSHA1(strings.Join(strArr, ""), appScrectKey)
	return sig
}

// 签名
func MakeSignature(app int64, sig string, nonce string, ts int, ttl int, data string) string {
	expired := ts + ttl*1000
	str := fmt.Sprintf("%d|%d|%s|%s|%d|%d|%s", expired, app, sig, nonce, ts, ttl, data)
	key := fmt.Sprintf("%d", app)
	signature := EncryptAes(str, key)
	return signature
}

// 解析签名
func ParseSignature(signature string, app string) (SignatureData, error) {
	var data SignatureData
	parseStr := DecryptAes(signature, app)
	if len(parseStr) < 1 {
		return data, fmt.Errorf("parse fail")
	}
	parseArr := strings.Split(parseStr, "|")
	for k, v := range parseArr {
		switch k {
		case 0:
			data.Expired, _ = strconv.Atoi(v)
			continue
		case 1:
			data.App, _ = strconv.ParseInt(v, 10, 64)
			continue
		case 2:
			data.Sig = v
			continue
		case 3:
			data.Nonce = v
			continue
		case 4:
			data.Ts, _ = strconv.Atoi(v)
			continue
		case 5:
			data.Ttl, _ = strconv.Atoi(v)
			continue
		case 6:
			data.Data = v
			continue
		default:
			continue
		}

	}

	return data, nil
}
