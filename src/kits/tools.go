package kits

import (
	"cmdb-go/config"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"
)

func Log(Msg, MsgType string) {
	Files := map[string]string{"info": config.InfoFile, "error": config.ErrorFile, "debug": config.DebugFile}
	Prefix := map[string]string{"info": "[Info]", "error": "[Error]", "debug": "[Debug]"}
	_, err := os.Stat(Files[MsgType])
	if err == nil {
		logFile, err := os.OpenFile(Files[MsgType], syscall.O_RDWR|syscall.O_APPEND, 0666)
		if err == nil {
			defer logFile.Close()
			debugLog := log.New(logFile, Prefix[MsgType], log.LstdFlags)
			debugLog.Println(Msg)
		}
	} else {
		logFile, err := os.Create(Files[MsgType])
		if err == nil {
			defer logFile.Close()
			debugLog := log.New(logFile, Prefix[MsgType], log.LstdFlags)
			debugLog.Println(Msg)
		}
	}
}
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Max(value []uint64) uint64 {
	var max uint64
	for _, val := range value {
		if val > max {
			max = val
		}
	}
	return max
}

func CheckFile(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
func StringToMap(content string) map[string]interface{} {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		fmt.Println("string转map失败", err)
	}
	return resMap
}
