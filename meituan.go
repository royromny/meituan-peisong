package meituan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// true会打印出logs
const DebugStatus = true

type Client struct {
	apiDomain string
	Client    *http.Client
	location  *time.Location
	appKey    string
	appSecret string
}

// New 初始化客户端
// appKey
// appSecret
// 测试环境与正式环境使用测试账号及线上账号区分
func New(appKey, appSecret string) (client *Client, err error) {
	location, _ := time.LoadLocation("Local")
	client = &Client{}
	client.appKey = appKey
	client.appSecret = appSecret
	client.apiDomain = kProductionURL
	client.Client = http.DefaultClient
	client.location = location

	return client, nil
}

// 补充处理基本请求参数
func (this *Client) URLValues(param Param) (map[string]string, error) {
	value := make(map[string]string)
	value["appkey"] = this.appKey
	value["timestamp"] = strconv.Itoa(int(time.Now().In(this.location).Unix()))
	value["version"] = kVersion
	p := StructToMapString(param)
	if p != nil {
		for k, v := range p {
			value[k] = v
		}
	}
	value["sign"] = sign(value, this.appSecret)
	return value, nil
}

func (this *Client) doRequest(method string, param Param, result interface{}) (err error) {
	DataUrlVal := url.Values{}
	if param != nil {
		p, err := this.URLValues(param)
		if err != nil {
			return err
		}
		for k, v := range p {
			if DebugStatus {
				println(k)
				println(v)
			}
			DataUrlVal.Add(k, v)
		}
	}
	req, err := http.NewRequest(method, this.apiDomain+param.APIName(), strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("********dataStr********")
	fmt.Println(string(data))

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return this.doRequest(method, param, result)
}

// 生成签名
func sign(p map[string]string, appSecret string) string {
	// 按key顺序整理出来
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strParam string
	for _, k := range keys {
		if k != "sign" && p[k] != "" {
			strParam += k + p[k]
		}
	}
	// 首加上app_secret秘钥
	strParam = appSecret + strParam

	return SHA1(strParam)
}

func (this *Client) AckNotify(w http.ResponseWriter) {
	AckNotify(w)
}

func AckNotify(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"code":0}`))
}
