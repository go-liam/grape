package oss

/*
{"RequestId":"BAD371B3-C4B2-4E60-868B-66F2E1BF3B87","AssumedRoleUser":{"Arn":"acs:ram::1462761290059388:role/oss/cbox",
"AssumedRoleId":"377319537320419904:cbox"},"Credentials":{"SecurityToken":"CAIShgJ1q6Ft5B2yfSjIr5f4P/CNrrcZ4Y6RZWyIoEI
kerlEqvafsDz2IH9Lf3BhBOsZtf43mGlZ5/4elq14QZhVHfpU4xmlqMY5yxioRqackc7XhOV2hf/IMGyXDAGBq622Su7lTdTbV+6wYlTf7EFayqf7cjPQMD
7INoaS29wdPbZxZASjaidcD9p7PxZrrNRgVUHcLvGwKBXn8B2yeUNjoVh7kngtq/b9kMGR9F/EgUa/welSrI72Mt/mTbE1YMklDInshr0uLveQi34A0XUQ
qvcq1p4j0Czco9SQD2NW5xi7KOfO+rVtVlQmOfNjSvMf86OmzaEh672Ly56czA1WLXWh7PRN5yTKGoABf5D+evWkPYMDfVnjTIW/8lfGeaPrnOaStqvpDH
cx6mkrsLt4tiNCLFLncLNexEBAw7HQh8oJRdC00Vwplh/4kCfBg2uPVtX/BXTEawArbEuFpK611AY9hbxPb+2kk2trGIxE8S8fTvZg6zCNpzRl5lzC9NK5YcZOBIb24AZ96aA=",
"AccessKeyId":"STS.NTMtJ9Ch8VLzgJ9DBqv6hE34R","AccessKeySecret":"EMLoSSj9R8D568kqPps3DY78xXYemjsGTvA1w9gf6Jnz","Expiration":"2020-08-10T07:50:49Z"}}

2020/08/10 14:50:49 v1={"AssumedRoleUser":{"Arn":"acs:ram::1462761290059388:role/oss/cbox",
"AssumedRoleId":"377319537320419904:cbox"},"Credentials":{"AccessKeyId":"STS.NTMtJ9Ch8VLzgJ9DBqv6hE34R",
"AccessKeySecret":"EMLoSSj9R8D568kqPps3DY78xXYemjsGTvA1w9gf6Jnz","Expiration":"2020-08-10T07:50:49Z",
"SecurityToken":"CAIShgJ1q6Ft5B2yfSjIr5f4P/CNrrcZ4Y6RZWyIoEIkerlEqvafsDz2IH9Lf3BhBOsZtf43mGlZ5/4elq14QZhVHfpU4xmlqMY5y
xioRqackc7XhOV2hf/IMGyXDAGBq622Su7lTdTbV+6wYlTf7EFayqf7cjPQMD7INoaS29wdPbZxZASjaidcD9p7PxZrrNRgVUHcLvGwKBXn8B2yeUNjoVh
7kngtq/b9kMGR9F/EgUa/welSrI72Mt/mTbE1YMklDInshr0uLveQi34A0XUQqvcq1p4j0Czco9SQD2NW5xi7KOfO+rVtVlQmOfNjSvMf86OmzaEh672L
y56czA1WLXWh7PRN5yTKGoABf5D+evWkPYMDfVnjTIW/8lfGeaPrnOaStqvpDHcx6mkrsLt4tiNCLFLncLNexEBAw7HQh8oJRdC00Vwplh/4kCfBg2uPVt
X/BXTEawArbEuFpK611AY9hbxPb+2kk2trGIxE8S8fTvZg6zCNpzRl5lzC9NK5YcZOBIb24AZ96aA="},
"RequestId":"BAD371B3-C4B2-4E60-868B-66F2E1BF3B87"}

2020/08/10 14:50:49 v2=<nil>
--- PASS: TestGetAliUtilInstance (1.21s)
https://shimo.im/docs/VXxQqDrykDhgGvcj
*/

/*
func TestServer_GetSTSToken(t *testing.T) {
	accessKeyID := "LTAIP5***"
	accessKeySecret := "yiGZyVXeK7**"
	roleAcs := "acs:ram::1462761290059388:role/oss"
	GetAliUtilInstance(accessKeyID, accessKeySecret, roleAcs)
	resArn := "acs:oss:*:*:%s/%s*"
	resArn = fmt.Sprintf(resArn, "dev-res-cn", "cbox/") //conf.Bucket, conf.PrefixDir)
	name := "cbox"
	//token, err := helpers.GetAliUtilInstance().GetSTSToken(conf.Name, resArn)
	token, err := aliUtil.GetSTSToken(name, resArn)
	log.Printf("v1=%+v\n", conv.StructToJsonString(token))
	log.Printf("v2=%+v\n", err)
}
*/

/*
if token != nil && token["Credentials"] != nil {
			cred := token["Credentials"].(map[string]interface{})
			tokenRes := &proto.GetStsTokenRsp{}
			tokenRes.OssConfig = transformOssConfigData(conf)
			tokenRes.AccessKeyId = cred["AccessKeyId"].(string)
			tokenRes.AccessKeySecret = cred["AccessKeySecret"].(string)
			tokenRes.StsToken = cred["SecurityToken"].(string)
			tokenRes.Host = conf.OssHost
			if conf.DomainHost != "" {
				tokenRes.Host = conf.DomainHost
			}
			return tokenRes, nil
		}
*/

// client test:

// 测试时要换上新申请的数据

/* *
func TestClient_STS_UploadFile(t *testing.T) {
	//"<yourBucketName>"
	bucketName := "**-res-cn"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := "content/org/test.jpeg"
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := "/Users/Downloads/study-rs/hezi.jpeg"
	yourEndpoint := "http://oss-cn-shenzhen.aliyuncs.com"
	// STS 临时数据
	yourAccessKeyId := "STS.NT1hp5tZ5qpbE*****"
	yourAccessKeySecret := "FmUnY7fCHTjpK7mL9wFeG1e77w31x*****"
	yourSecurityToken := "CAISl2yfSjIr5eEI8q***W1lQiNOFovYDksztbZkG5c"
	// 用户拿到STS临时凭证后，通过其中的安全令牌（SecurityToken）和临时访问密钥（AccessKeyId和AccessKeySecret）生成OSSClient。
	// 创建OSSClient实例。
	client, err := oss.New(yourEndpoint, yourAccessKeyId, yourAccessKeySecret, oss.SecurityToken(yourSecurityToken))
	if err != nil {
		fmt.Println("1-Error:", err)
		os.Exit(-1)
	}
	// 获取存储空间。
	bucket, err2 := client.Bucket(bucketName)
	if err2 != nil {
		fmt.Println("2-Error:", err2)
		os.Exit(-1)
	}
	// 上传文件。
	err3 := bucket.PutObjectFromFile(objectName, localFileName)
	if err3 != nil {
		fmt.Println("3-Error:", err3)
		os.Exit(-1)
	}
}
* */
