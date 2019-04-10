package controllers

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WeixinPayController struct {
	BaseController
}

type WXPayResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Nonce_str   string `xml:"nonce_str"`
	Prepay_id   string `xml:"prepay_id"`
	AppId       string `xml:"appid"`
	MchId       string `xml:"mch_id"`
	CodeUrl     string `xml:"code_url"`
	ResultCode  string `xml:"result_code"`
}

//微信支付
func (c *WeixinPayController) WxPay() {
	//info := make(map[string]interface{}, 0)
	ip := c.Ctx.Request.RemoteAddr[0:strings.Index(c.Ctx.Request.RemoteAddr, ":")]
	ip = "127.0.0.1"

	total_fee, _ := strconv.ParseFloat(c.GetString("total_fee"), 64) //单位分
	total_fee = 1.01
	//openId
	//openId := c.GetString("openId") //"oKYr_0GkE-Izt9N9Wn43sapI9Pqw"
	openId := "o8wdO5RIivHTkMLZEo5wWVzLBXhE"
	body := "费用说明"
	// //订单号
	//orderNo := c.GetString("orderNo") //"wx"+utils.ToStr(time.Now().Unix()) + string(utils.Krand(4, 0))
	orderNo := "wsl39182189"
	// //随机数
	rand.Seed(time.Now().UnixNano())
	nonceStr := time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(10000))
	var reqMap = make(map[string]interface{}, 0)
	reqMap["appid"] = "wxa8073b7865bc35e3"                               //微信小程序appid
	reqMap["body"] = body                                                //商品描述
	reqMap["mch_id"] = "1530840201"                                      //商户号
	reqMap["nonce_str"] = nonceStr                                       //随机数
	reqMap["notify_url"] = "http://test.com.cn/weixinNotice.jspx"        //通知地址
	reqMap["openid"] = openId                                            //商户唯一标识 openid
	reqMap["out_trade_no"] = orderNo                                     //订单号
	reqMap["spbill_create_ip"] = ip                                      //用户端ip   //订单生成的机器 IP
	reqMap["total_fee"] = strconv.FormatFloat(total_fee*100, 'f', 0, 64) //订单总金额，单位为分
	reqMap["trade_type"] = "JSAPI"                                       //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, "VDMwoWGWGLeTJqmB9XkR7aQ4n52n8EYJ")

	reqStr := Map2Xml(reqMap)
	fmt.Println(reqStr)

	fmt.Println("请求xml", reqStr)

	client := &http.Client{}

	// 调用支付统一下单API
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(reqStr))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("解析响应内容失败", err)
		return
	}

	var resp1 WXPayResp
	err = xml.Unmarshal(body2, &resp1)
	if err != nil {
		panic(err)
	}
	fmt.Println("响应数据", resp1.Nonce_str)
	fmt.Println("响应数据", resp1.Return_msg)
	fmt.Println("响应数据", resp1.Return_code)
	fmt.Println("响应数据", resp1.Prepay_id)
	fmt.Println("响应数据", resp1.AppId)
	fmt.Println("响应数据", resp1.MchId)
	fmt.Println("响应数据", resp1.CodeUrl)
	fmt.Println("响应数据", resp1.ResultCode)
	fmt.Println(string(body2))

	// // 返回预付单信息
	if strings.ToUpper(resp1.Return_code) == "SUCCESS" {
		fmt.Println("预支付申请成功")
		// 再次签名
		// 	var resMap = make(map[string]interface{}, 0)
		// 	resMap["appId"] = "wxa8073b7865bc35e3"
		// 	resMap["nonceStr"] = resp1.Nonce_str                 //商品描述
		// 	resMap["package"] = "prepay_id=" + resp1.Prepay_id   //商户号
		// 	resMap["signType"] = "MD5"                           //签名类型
		// 	resMap["timeStamp"] = utils.ToStr(time.Now().Unix()) //当前时间戳

		// 	resMap["paySign"] = WxPayCalcSign(resMap, utils.WX_KEY)
		// 	// 返回5个支付参数及sign 用户进行确认支付

		// 	fmt.Println("支付参数", resMap)
		// 	c.Console(resMap)
		// } else {
		// 	info["msg"] = "微信请求支付失败"
		// 	c.Ctx.Console(info)
	}
	return
}

func Map2Xml(mReq map[string]interface{}) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k, v := range mReq {
		sb.WriteString("<" + k + ">" + v.(string) + "</" + k + ">")
	}
	sb.WriteString("</xml>")
	return sb.String()
}

func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	for _, k := range sorted_keys {
		log.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
