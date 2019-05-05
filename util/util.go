package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/microcosm-cc/bluemonday"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/speps/go-hashids"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
	"gopkg.in/russross/blackfriday.v2"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// encode uid 为字符串
func EncodeUid(uid int64) string {
	hd := hashids.NewData()
	hd.Salt = "1024casts_uid"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	e, err := h.EncodeInt64([]int64{uid})
	if err != nil {
		log.Warn("encode uid err")
	}

	return e

}

// decode uid 为int64
func DecodeUid(encodedUid string) (uid int64) {
	hd := hashids.NewData()
	hd.Salt = "1024casts_uid"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	d, err := h.DecodeInt64WithError(encodedUid)

	if err != nil {
		log.Warn("decode uid err")
	}

	if len(d) > 0 {
		return d[0]
	}

	return 0

}

func GetUserId(ctx *gin.Context) uint64 {
	cookie, err := ctx.Cookie(viper.GetString("cookie.name"))
	if err != nil {
		return 0
	}

	userId := DecodeUid(cookie)
	return uint64(userId)
}

// 获取七牛资源的私有链接
func GetQiniuPrivateAccessUrl(path string) string {
	accessKey := viper.GetString("qiniu.access_key")
	secretKey := viper.GetString("qiniu.secret_key")
	mac := qbox.NewMac(accessKey, secretKey)

	domain := viper.GetString("qiniu.cdn_url")
	key := strings.TrimPrefix(path, "/")

	imageStyle := "imageView2/2/w/200/h/200/q/75|imageslim"
	key = key + "?" + imageStyle
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期

	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	return privateAccessURL
}

// 获取七牛资源的公有链接
// 无需配置bucket, 域名会自动到域名所绑定的bucket去查找
func GetQiNiuPublicAccessUrl(path string) string {
	domain := viper.GetString("qiniu.public_cdn_url")
	key := strings.TrimPrefix(path, "/")

	publicAccessURL := storage.MakePublicURL(domain, key)

	return publicAccessURL
}

func TimeLayout() string {
	return "2006-01-02 15:04:05"
}

func TimeToString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format(TimeLayout())
}

func TimeToDateString(ts time.Time) string {
	return time.Unix(ts.Unix(), 00).Format("2006年01月02日")
}

func TimeToInt64(ts time.Time) int64 {
	return ts.Unix()
}

func GetDate() string {
	return time.Now().Format("2006/01/02")
}

func GetImageFullUrl(uri string) string {
	return viper.GetString("image_domain") + uri
}

func GenerateOrderNo() (uint64, error) {
	dateStr := time.Now().Format("20060102150405")
	log.Infof("data str: %s", dateStr)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randStr := fmt.Sprintf("%05v", rnd.Intn(10000))
	log.Infof("rand str: %s", randStr)

	orderNoStr := dateStr + randStr
	orderNo, err := strconv.ParseUint(orderNoStr, 10, 64)
	if err != nil {
		log.Warnf("[util] convert: %s err: %+v", randStr, err)
		return 0, err
	}

	log.Infof("orderNo: %d", orderNo)

	return orderNo, nil
}

func MarkdownToHtml(con string) string {
	mdByte := []byte(con)
	unsafe := blackfriday.Run(mdByte)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return string(html)
}

// 随机生成字符串
// more: https://colobu.com/2018/09/02/generate-random-string-in-Go/
func RandStr(strLen int) string {
	var (
		codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/~!@#$%^&*()_="
		codeLen = len(codes)
	)
	data := make([]byte, strLen)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < strLen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}
	return string(data)
}

func GenPasswordToken() string {
	hashKey := []byte("aHG^OKHJ^)R$sp1q_key")
	h := hmac.New(sha256.New, hashKey)
	randStr := RandStr(40)
	log.Infof("rand str:%s", randStr)
	h.Write([]byte(randStr))

	return hex.EncodeToString(h.Sum(nil))
}

// see: https://blog.csdn.net/haibo0668/article/details/77648875
func FormatTime(needTime time.Time) string {
	var timeText string
	var curTime = time.Now().Unix()
	var needTimeTs = needTime.Unix()
	if needTimeTs < 0 {
		return timeText
	}

	// 是否跨年
	year := time.Now().Year() - needTime.Year()
	isCrossYear := false
	if year > 0 {
		isCrossYear = true
	}

	var tempStr string

	log.Infof("test: %+v", needTime.Unix())

	// 时间差，单位：秒
	switch t := curTime - needTimeTs; {
	case t == 0:
		timeText = "刚刚"
	case t < 60:
		timeText = fmt.Sprintf("%d秒前", t) // 一分钟内
	case t < 60*60:
		var temp = math.Floor(float64(t / 60))
		tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
		timeText = fmt.Sprintf("%s分钟前", tempStr) // 一小时内
	case t < 60*60*24:
		var temp = math.Floor(float64(t / (60 * 60)))
		tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
		timeText = fmt.Sprintf("%s小时前", tempStr) // 一天内
	case t < 60*60*24*2:
		timeText = fmt.Sprintf("昨天%s", needTime.Format("15:04")) // 昨天
	case t < 60*60*24*30:
		var temp = math.Floor(float64(t / (60 * 60 * 24)))
		tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
		// timeText = needTime.Format("01月02日 15:04") // 一个月内
		timeText = fmt.Sprintf("%s天前", tempStr)
	case t < 60*60*24*365 && isCrossYear == false:
		var temp = math.Floor(float64(t / (60 * 60 * 24 * 30)))
		tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
		//timeText = needTime.Format("01月02日") // 一年内
		timeText = fmt.Sprintf("%s个月前", tempStr)
	default:
		timeText = needTime.Format("2006年01月02日") // 一年以前
	}

	return timeText
}

/**
* @des 时间转换函数
* @param atime string 要转换的时间戳（秒）
* @return string
 */
func StrTime(atime time.Time) string {
	var byTime = []int64{
		365 * 24 * 60 * 60,
		24 * 60 * 60 * 30,
		24 * 60 * 60 * 2,
		24 * 60 * 60,
		60 * 60,
		60,
		1,
	}
	var unit = []string{"年前", "个月前", "天前", "昨天", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime.Unix()
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i])
		}
		break
	}
	return res
}

/**
* @desc 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
