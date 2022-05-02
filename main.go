package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 请求参数结构体
type ParamData struct {
	AppId           string `json:"appId"`
	PaasHeader      string `json:"paasHeader"`
	TimestampHeader int64  `json:"timestampHeader"`
	NonceHeader     string `json:"nonceHeader"`
	SignatureHeader string `json:"signatureHeader"`
	Key             string `json:"key"`
}

func main() {
	timeStamp := time.Now().Unix()
	timeStampStr := strconv.FormatInt(timeStamp, 10)
	paramData := getParamData(timeStamp, timeStampStr)
	centerSign := "fTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA"
	signature := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(timeStampStr+centerSign+timeStampStr))))
	result := postApi(paramData, signature, timeStampStr)
	fmt.Println(result)
}

// 获取请求参数
func getParamData(timeStamp int64, timeStampStr string) ParamData {
	i := "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA"
	a := "123456789abcdefg"
	s := "zdww"
	str := timeStampStr + i + a + timeStampStr
	signature := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(str))))

	return ParamData{
		AppId:           "NcApplication",
		PaasHeader:      s,
		TimestampHeader: timeStamp,
		NonceHeader:     a,
		SignatureHeader: signature,
		Key:             "3C502C97ABDA40D0A60FBEE50FAAD1DA",
	}
}

func postApi(urlValues ParamData, signature string, timeStampStr string) string {
	client := &http.Client{}

	jsonValue, _ := json.Marshal(urlValues)
	url := "https://bmfw.www.gov.cn/bjww/interface/interfaceJson"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("x-wif-nonce", "QkjjtiLM2dCratiA")
	req.Header.Set("x-wif-paasid", "smt-application")
	req.Header.Set("x-wif-signature", signature)
	req.Header.Set("x-wif-timestamp", timeStampStr)
	req.Header.Set("content-Type", "application/json; charset=utf-8")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	return strBody
}
