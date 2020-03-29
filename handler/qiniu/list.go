package qiniu

import (
	"fmt"
	"strconv"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
)

func List(c *gin.Context) {

	accessKey := viper.GetString("qiniu.access_key")
	secretKey := viper.GetString("qiniu.secret_key")
	log.Infof("accessKey", accessKey)
	log.Infof("secretKey", secretKey)

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{
		Zone: &storage.ZoneHuabei,
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
		// 上传是否使用CDN上传加速
		UseCdnDomains: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	infos := make([]storage.ListItem, 0)

	limitStr := c.DefaultQuery("pageSize", "20")
	limit, _ := strconv.Atoi(limitStr)

	prefix := c.DefaultQuery("prefix", "/")
	delimiter := ""
	//初始列举marker为空
	marker := c.DefaultQuery("next_marker", "")

	bucket := c.DefaultQuery("bucket", viper.GetString("qiniu.bucket"))

	log.Debugf("req params:", bucket, prefix, delimiter, marker, limit)
	entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
	if err != nil {
		fmt.Println("list error,", err)
	}
	//print entries
	for _, entry := range entries {
		//fmt.Println(entry.Key)
		infos = append(infos, entry)

	}

	if hasNext {
		log.Infof("nextMarker: %s", nextMarker)
	}

	app.Response(c, nil, ListResponse{
		TotalCount: uint64(len(infos)),
		List:       infos,
		NextMaker:  nextMarker,
	})

}
