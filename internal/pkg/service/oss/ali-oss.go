package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-liam/util/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var aliUtil *AliUtil
var onceAliUtil sync.Once

//GetAliUtilInstance :单实例
func GetAliUtilInstance(accessKeyID string, accessKeySecret string, roleAcs string) *AliUtil {
	onceAliUtil.Do(func() {
		config := &AliConfig{
			AccessKeyID:     accessKeyID,     //conf.EnvConfig.Ali.AccessKeyID,
			AccessKeySecret: accessKeySecret, //conf.EnvConfig.Ali.AccessKeySecret,
			RoleAcs:         roleAcs,         //conf.EnvConfig.Ali.RoleAcs,
		}

		aliUtil = &AliUtil{
			config: config,
		}
	})
	return aliUtil
}

type AliConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	RoleAcs         string
}

type AliUtil struct {
	config *AliConfig
}

//GetSTSToken GetSTSToken
func (p *AliUtil) GetSTSToken(sessionName string, resArn string) (map[string]interface{}, error) {
	policy := `{"Statement":[{"Action":["oss:*"],"Effect":"Allow","Resource":["%s"]}],"Version":"1"}`
	policy = fmt.Sprintf(policy, resArn)
	url, err := p.GenerateStsSignatureURL(sessionName, policy, "3600")
	if err != nil {
		return nil, err
	}
	data, err := GetStsResponse(url)
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

//GenerateStsSignatureURL GenerateStsSignatureURL
func (p *AliUtil) GenerateStsSignatureURL(sessionName, policy, durationSeconds string) (string, error) {
	assumeURL := "AccessKeyId=" + p.config.AccessKeyID
	assumeURL += "&Action=AssumeRole"
	assumeURL += "&DurationSeconds=" + durationSeconds
	assumeURL += "&Format=JSON"
	assumeURL += "&Policy=" + url.QueryEscape(policy)
	assumeURL += "&RoleArn=" + url.QueryEscape(p.config.RoleAcs)
	assumeURL += "&RoleSessionName=" + sessionName
	assumeURL += "&SignatureMethod=HMAC-SHA1"
	assumeURL += "&SignatureNonce=" + uuid.UUID() //NewUUID()
	assumeURL += "&SignatureVersion=1.0"
	assumeURL += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	assumeURL += "&Version=2015-04-01"

	// 解析成V type
	signToString, err := url.ParseQuery(assumeURL)
	if err != nil {
		return "", err
	}

	// URL顺序化
	result := signToString.Encode()

	// 拼接
	StringToSign := "GET" + "&" + "%2F" + "&" + url.QueryEscape(result)

	// HMAC
	hashSign := hmac.New(sha1.New, []byte(p.config.AccessKeySecret+"&"))
	hashSign.Write([]byte(StringToSign))

	// 生成signature
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	// Url 添加signature
	assumeURL = "https://sts.aliyuncs.com/?" + assumeURL + "&Signature=" + url.QueryEscape(signature)

	return assumeURL, nil
}

// GetStsResponse 请求构造好的URL,获得授权信息
// TODO: 安全认证 HTTPS
func GetStsResponse(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return body, err
}
