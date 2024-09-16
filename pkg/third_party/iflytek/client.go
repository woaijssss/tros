package iflytek

import (
	"bufio"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	http2 "github.com/woaijssss/tros/client/http"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	defaultBufferSize   = 1024 * 16
	apiSuccessCode      = "000000"
	orderFinishedStatus = 4

	ApiNoSuccessErr         = "api resp no success"
	ApiGetResultFailTypeErr = "api get result fail"
)

type client struct {
	appId     string
	secretKey string
}

type basicUpload struct {
	uploadUrl string
	filePath  string
	fileSize  int64
	duration  int64
	f         *os.File
}

type basicGetResult struct {
	getResultUrl string
	orderId      string
}

type sizedReader struct {
	r        io.Reader
	readSize int64
}

type uploadResp struct {
	Code     string `json:"code"`
	DescInfo string `json:"descInfo"`
	Content  struct {
		OrderId          string `json:"orderId"`
		TaskEstimateTime int    `json:"taskEstimateTime"`
	} `json:"content"`
}

func (r *uploadResp) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func (c *client) upload(ctx context.Context, bu *basicUpload) (string, string, error) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	uploadFileName := utils.GetFilename(bu.filePath)
	signa := generateSignature(c.appId+ts, c.secretKey)
	paramList := buildUploadParamList(c.appId, uploadFileName, bu.fileSize, bu.duration, ts, signa)
	parameters := formUrlEncodedValueParameters(paramList)
	uploadUrl := bu.uploadUrl + "?" + parameters
	//reader := bufio.NewReader(file)
	reader := &sizedReader{
		r: bufio.NewReaderSize(bu.f, defaultBufferSize),
	}

	hc := http2.NewHttpClient()
	hc.SetHeader("Content-Type", "application/octet-stream")

	resp, err := hc.Post(ctx, uploadUrl, reader)
	if err != nil {
		trlogger.Errorf(ctx, "sdk Upload Client.Do err: %v", err)
		return "", "", err
	}

	trlogger.Infof(ctx, "sdk Upload %s file success, uploaded bytes size: %d, file size is:%d,url %s", uploadFileName, reader.readSize, bu.fileSize, uploadUrl)

	if resp.StatusCode != http.StatusOK {
		trlogger.Errorf(ctx, "sdk Upload http.Post statusCode: %d", resp.StatusCode)
		return "", "", errors.New("upload call asr-service failed")
	}

	ret := new(uploadResp)
	err = http2.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "sdk Upload utils.ResToObj err: %v", err)
		return "", "", err
	}

	if ret.Code != apiSuccessCode {
		trlogger.Errorf(ctx, "sdk Upload asr-service failed: %s", ret.String())
		return "", ret.DescInfo, errors.New(ApiNoSuccessErr)
	}

	return ret.Content.OrderId, "", nil
}

type getResultResp struct {
	Code     string `json:"code"`
	DescInfo string `json:"descInfo"`
	Content  struct {
		OrderInfo struct {
			OrderId          string `json:"orderId"`
			FailType         int    `json:"failType"`
			Status           int    `json:"status"`
			OriginalDuration int    `json:"originalDuration"`
			RealDuration     int    `json:"realDuration"`
			ExpireTime       int    `json:"expireTime"`
		}
		OrderResult string `json:"orderResult"`
	} `json:"content"`
}

func (c *client) getResult(ctx context.Context, bg *basicGetResult) (string, error) {
	var (
		result, txt, desc string
	)
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	signa := generateSignature(c.appId+ts, c.secretKey)
	paramList := buildGetResultParamList(c.appId, bg.orderId, ts, signa)
	formUrlString := formUrlEncodedValueParameters(paramList)
	resultUrl := bg.getResultUrl + "?" + formUrlString

	status := 3
	for status == 3 {
		response, err := http2.NewHttpClient().Get(ctx, resultUrl)
		if err != nil {
			trlogger.Errorf(ctx, "sdk GetResult Client.Do err: %v", err)
			return "", err
		}

		if response.StatusCode != http.StatusOK {
			trlogger.Errorf(ctx, "sdk GetResult http.Post statusCode: %d", response.StatusCode)
			return "", errors.New("getResult call asr-service failed")
		}

		ret := new(getResultResp)
		err = http2.ResToObj(response, ret)
		if err != nil {
			trlogger.Errorf(ctx, "sdk GetResult utils.ResToObj err: %v", err)
			return "", err
		}

		if ret.Code != apiSuccessCode {
			trlogger.Errorf(ctx, "sdk GetResult asr-service failed: %s", ret.String())
			return "", errors.New(ApiNoSuccessErr)
		}

		orderInfo := ret.Content.OrderInfo
		if orderInfo.FailType != 0 && orderInfo.FailType != 11 {
			reason := "order failType: " + strconv.FormatInt(int64(orderInfo.FailType), 10)
			trlogger.Errorf(ctx, "sdk GetResult asr-service failed: %s, reason: [%s]", ret.String(), reason)
			return "", errors.New(ApiGetResultFailTypeErr)
		}

		//trlogger.Infof(ctx, "sdk GetResult orderId: %s,orderStatus: %d", orderId, orderInfo.Status)
		trlogger.Infof(ctx, "sdk GetResult orderId: %s,orderStatus: %d ,url %s", bg.orderId, orderInfo.Status, resultUrl)
		orderResult := &OrderResult{}
		// 订单已完成的时候,解析识别结果
		if orderInfo.Status == orderFinishedStatus {
			err = json.Unmarshal([]byte(ret.Content.OrderResult), orderResult)
			if err != nil {
				trlogger.Errorf(ctx, "sdk GetResult json.Unmarshal orderResult err: %v", err)
				return "", err
			}
			start := time.Now()
			result = ret.String()
			txt = orderResult.String()
			desc = ret.DescInfo
			end := time.Now()

			// 计算程序运行时长
			duration := end.Sub(start)
			trlogger.Infof(ctx, "程序运行时长: %v", duration)
			trlogger.Debugf(ctx, "result:[%s], desc:[%s]", result, desc)
			break
		}

		time.Sleep(5 * time.Second)
	}
	return txt, nil
}

func generateSignature(input string, accessSecretKey string) string {
	md5Sum := encrypt.EncodeMD5Byte(input)
	hmacHash := hmac.New(sha1.New, []byte(accessSecretKey))
	hmacHash.Write(md5Sum)
	signa := hmacHash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signa)
}

type paramPair struct {
	Key   string
	Value any
}

func buildUploadParamList(appId string, filename string, filesize int64, duration int64, ts, signa string) []*paramPair {
	//newUUID, _ := uuid.NewUUID()
	//uuidStr := newUUID.String()

	var paramList []*paramPair
	paramList = append(paramList, &paramPair{Key: "appId", Value: appId})
	//paramList = append(paramList, &paramPair{Key: "dateTime", Value: time.Now().Format(timeFormat)})
	paramList = append(paramList, &paramPair{Key: "signa", Value: signa})
	paramList = append(paramList, &paramPair{Key: "ts", Value: ts})
	paramList = append(paramList, &paramPair{Key: "fileName", Value: filename})
	paramList = append(paramList, &paramPair{Key: "fileSize", Value: filesize})
	paramList = append(paramList, &paramPair{Key: "duration", Value: duration})
	//paramList = append(paramList, &paramPair{Key: "roleType", Value: 1})
	//paramList = append(paramList, &paramPair{Key: "roleNum", Value: 0})

	utils.Cmp[*paramPair](func(p1, p2 **paramPair) bool {
		return (*p1).Key < (*p2).Key
	}).Sort(paramList)
	return paramList
}

func formUrlEncodedValueParameters(paramPairs []*paramPair) string {
	paramStr := ""
	for _, param := range paramPairs {
		if param.Key != "" {
			paramStr += param.Key + "=" + url.QueryEscape(fmt.Sprintf("%v", param.Value)) + "&"
		}
	}

	paramLen := len(paramStr)
	if paramLen > 0 {
		paramStr = paramStr[:paramLen-1]
	}
	return paramStr
}

func (sr *sizedReader) Read(p []byte) (n int, err error) {
	n, err = sr.r.Read(p)
	if err == nil {
		sr.readSize += int64(n)
	}
	return n, err
}

func buildGetResultParamList(appId, orderId string, ts, signa string) []*paramPair {
	var paramList []*paramPair
	paramList = append(paramList, &paramPair{Key: "appId", Value: appId})
	paramList = append(paramList, &paramPair{Key: "signa", Value: signa})
	paramList = append(paramList, &paramPair{Key: "ts", Value: ts})
	paramList = append(paramList, &paramPair{Key: "orderId", Value: orderId})
	paramList = append(paramList, &paramPair{Key: "resultType", Value: "transfer"})
	utils.Cmp[*paramPair](func(p1, p2 **paramPair) bool {
		return (*p1).Key < (*p2).Key
	}).Sort(paramList)
	return paramList
}

func (r *getResultResp) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

type OrderResult struct {
	Lattice []struct {
		Json1best string `json:"json_1best"`
	} `json:"lattice"`
}

type Json1best struct {
	St struct {
		Sc string `json:"sc"`
		Pa string `json:"pa"`
		Rt []struct {
			Ws []struct {
				Cw []struct {
					W  string `json:"w"`
					Wp string `json:"wp"`
					Wc string `json:"wc"`
				} `json:"cw"`
				Wb int `json:"wb"`
				We int `json:"we"`
			} `json:"ws"`
		} `json:"rt"`
		Bg string `json:"bg"`
		Ed string `json:"ed"`
		Rl string `json:"rl"`
	} `json:"st"`
}

func (o OrderResult) String() string {
	var json1BestArr []*Json1best
	for _, lattice := range o.Lattice {
		json1Best := &Json1best{}
		err := json.Unmarshal([]byte(lattice.Json1best), json1Best)
		if err != nil {
			return err.Error()
		}
		json1BestArr = append(json1BestArr, json1Best)
	}

	var totalContent bytes.Buffer
	for _, json1best := range json1BestArr {
		itemStr := ""
		rl := json1best.St.Rl
		for _, rt := range json1best.St.Rt {
			for _, ws := range rt.Ws {
				for _, cw := range ws.Cw {
					itemStr += cw.W
				}
			}
		}
		fmt.Fprintf(&totalContent, "发言人%s:  %s \n", rl, itemStr)
	}
	return totalContent.String()
}
