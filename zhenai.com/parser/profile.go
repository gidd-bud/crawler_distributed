package parser

import (
	"IMOOC/crawler_distributed/engine"
	"IMOOC/crawler_distributed/model"
	"regexp"
	"strconv"
)

//	"age":24,
//	"basicInfo":["天蝎座(10.23-11.21)",,"50kg",,"医疗管理",],
//	"detailInfo":["汉族","籍贯:江苏苏州","体型:富线条美","不吸烟","社交场合会喝酒","租房","未买车","没有小孩","是否想要孩子:视情况而定","何时结婚:两年内"],
//	"educationString":"大学本科",
//	"emotionStatus":0,
//	"gender":1,
//	"genderString":"女士",
//	"hasIntroduce":true,
//	"heightString":"165cm",
//	"hideVerifyModule":false,
//	"marriageString":"未婚",
//	"memberID":1172388090,
//	"nickname":"思忆",
//	"praisedIntroduce":false,
//	"salaryString":"20001-50000元",
//	"showValidateIDCardFlag":false,
//	"totalPhotoCount":7,
//	"validateEducation":false,
//	"validateFace":false,
//	"validateIDCard":true,
//	"videoCount":0,
//	"videoID":0,
//	"workCity":10103001,
//	"workCityString":"上海",
//	"workProvinceCityString":"上海浦东新区",

var (
	idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	memberIDRe = regexp.MustCompile(`"memberID":([0-9]+)`)
	//nicknameRe = regexp.MustCompile(`"nickname":"([^"]+)"`)
	//trueNameRe = regexp.MustCompile(`"trueName":([0-9]+)`)
	ageRe = regexp.MustCompile(`"age":([\d]+)`)
	genderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)
	heightRe = regexp.MustCompile(`"heightString":"([\d]+)cm"`)
	educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
	marriageRe = regexp.MustCompile(`"marriageString":"([^"]+)"`)
	salaryRe = regexp.MustCompile(`"salaryString":"([^"]+)"`)
	workCityRe = regexp.MustCompile(`"workCityString":"([^"]+)"`)
)

func ParseProfile(url string, contents []byte, nickname string) engine.ParseResult {
	profile := model.Profile{
		extractString(contents, memberIDRe),
	nickname,
	extractInt(contents, ageRe),
	extractString(contents, genderRe),
	extractInt(contents, heightRe),
	extractString(contents, educationRe),
	extractString(contents, workCityRe),
	extractString(contents, marriageRe),
	extractString(contents, salaryRe),
	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id:extractString([]byte(url), idUrlRe),
				Payload:profile,
			}},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	submatch := re.FindSubmatch(contents)
	if len(submatch) >= 2 {
		return string(submatch[1])
	}else{
		return ""
	}
}
func extractInt(contents []byte, re *regexp.Regexp) int {
	if retVal, err := strconv.Atoi(extractString(contents, re)); err == nil{
		return retVal
	} else {
		return 0
	}
}
