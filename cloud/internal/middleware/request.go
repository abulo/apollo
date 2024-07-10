package middleware

import (
	"bytes"
	"cloud/code"
	"cloud/initial"
	"cloud/internal/encrypt"
	"cloud/internal/response"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/cast"
)

// Params 加密之后的参数
type Params struct {
	Params string `json:"params"`
}

func Request() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		if !initial.Core.Config.Bool("encrypt.Disable", true) { // 先判断是不是需要加密
			// 先判断 method
			method := cast.ToString(newCtx.Request.Method())
			if method == "OPTIONS" {
				newCtx.Next(ctx)
			} else {
				HandleDecrypt(ctx, newCtx)
			}
		} else {
			newCtx.Next(ctx)
		}
	}
}

// HandleDecrypt 请求解密
func HandleDecrypt(ctx context.Context, newCtx *app.RequestContext) {
	// 获取请求类型
	contentType := newCtx.Request.Header.Get("Content-Type")
	// 判断是不是json请求
	isJsonRequest := strings.Contains(contentType, "application/json")
	// 判断是不是文件上传
	// isFileRequest := strings.Contains(contentType, "multipart/form-data")
	// 判断是不是表单提交
	isFormUrl := strings.Contains(contentType, "application/x-www-form-urlencoded")
	// 判断请求
	method := cast.ToString(newCtx.Request.Method())
	if method == "GET" {
		err := parseQuery(newCtx)
		if err != nil {
			response.JSON(newCtx, consts.StatusOK, utils.H{
				"code": code.DecryptError,
				"msg":  code.StatusText(code.DecryptError),
			})
			newCtx.Abort()
			return
		}
	} else if isJsonRequest {
		err := parseJson(newCtx)
		if err != nil {
			response.JSON(newCtx, consts.StatusOK, utils.H{
				"code": code.DecryptError,
				"msg":  code.StatusText(code.DecryptError),
			})
			newCtx.Abort()
			return
		}
	} else if isFormUrl {
		err := parseForm(newCtx)
		if err != nil {
			response.JSON(newCtx, consts.StatusOK, utils.H{
				"code": code.DecryptError,
				"msg":  code.StatusText(code.DecryptError),
			})
			newCtx.Abort()
			return
		}
	}
	newCtx.Next(ctx)
}

func parseFile(newCtx *app.RequestContext) error {

	contentType := newCtx.Request.Header.Get("Content-Type")
	_, params, _ := mime.ParseMediaType(contentType)
	boundary, ok := params["boundary"]
	if !ok {
		return errors.New("no multipart boundary param in Content-Type")
	}
	//准备重写数据
	bodyBuf := &bytes.Buffer{}
	wr := multipart.NewWriter(bodyBuf)
	mr := multipart.NewReader(newCtx.Request.BodyStream(), boundary)
	for {
		p, err := mr.NextPart() //p的类型为Part
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		fileByte, err := io.ReadAll(p)
		if err != nil {
			break
		}
		pName := p.FormName()
		fileName := p.FileName()
		if len(fileName) < 1 {
			if pName == "params" {
				decryptedData := encrypt.AesCBCDecrypt(string(fileByte), initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
				formData := make(map[string]interface{}, 0)
				if err := json.Unmarshal([]byte(decryptedData), &formData); err != nil {
					return err
				}
				for k, v := range formData {
					val := getStr(v)
					err = wr.WriteField(k, val)
					if err != nil {
						break
					}
				}
			} else {
				wr.WriteField(pName, string(fileByte))
			}
		} else {
			tmp, err := wr.CreateFormFile(pName, fileName)
			if err != nil {
				continue
			}
			tmp.Write(fileByte)
		}
	}
	//写结尾标志
	_ = wr.Close()
	newCtx.Request.Header.Set("Content-Type", wr.FormDataContentType())
	newCtx.Request.SetBody(bodyBuf.Bytes())
	return nil
}

func parseForm(newCtx *app.RequestContext) error {
	//读取数据 body处理
	payload := newCtx.GetRawData()
	if len(payload) < 1 {
		return nil
	}
	values, err := url.ParseQuery(string(payload))
	if err != nil {
		return err
	}
	payloadText := values.Get("params")
	if len(payloadText) > 0 {
		var jsonData Params
		err := json.Unmarshal([]byte(payloadText), &jsonData)
		if err != nil {
			return err
		}
		decryptedData := encrypt.AesCBCDecrypt(jsonData.Params, initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
		mapData := make(map[string]interface{}, 0)
		if err := json.Unmarshal([]byte(decryptedData), &mapData); err != nil {
			return err
		}
		for k, v := range mapData {
			values.Add(k, getStr(v))
		}
		formData := values.Encode()
		newCtx.Request.SetBody([]byte(formData))
	}
	return nil
}

func parseJson(newCtx *app.RequestContext) error {
	//读取数据 body处理
	payload := newCtx.GetRawData()
	if len(payload) < 1 {
		return nil
	}
	var jsonData Params
	err := json.Unmarshal(payload, &jsonData)
	if err != nil {
		return err
	}
	payloadText := jsonData.Params
	if len(payloadText) > 0 {
		// 解密字符串
		decryptedData := encrypt.AesCBCDecrypt(payloadText, initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
		newCtx.Request.SetBody([]byte(decryptedData))
	}
	return nil
}

func parseQuery(newCtx *app.RequestContext) error {
	params, exists := newCtx.GetQuery("params")
	if !exists {
		return nil
	}
	if len(params) < 1 {
		return nil
	}
	// 解密字符串
	queryData := encrypt.AesCBCDecrypt(params, initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
	if len(queryData) < 1 {
		return nil
	}
	formData := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(queryData), &formData); err != nil {
		return err
	}
	var args []string
	for k, v := range formData {
		val := getStr(v)
		args = append(args, fmt.Sprintf("%s=%s", k, url.QueryEscape(val)))
	}
	newCtx.Request.SetQueryString(strings.Join(args, "&"))
	return nil
}

func getStr(v interface{}) string {
	val := ""
	switch v.(type) {
	case float64:
		fl, _ := v.(float64)
		val = strconv.FormatFloat(fl, 'f', -1, 64)
	default:
		val = fmt.Sprintf("%v", v)
	}
	return val
}
