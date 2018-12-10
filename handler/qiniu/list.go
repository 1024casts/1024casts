package qiniu

import (
	"fmt"

	. "github.com/1024casts/1024casts/handler"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

func List(c *gin.Context) {

	accessKey := viper.GetString("qiniu.AccessKey")
	secretKey := viper.GetString("qiniu.SecretKey")
	bucket := viper.GetString("qiniu.Bucket")

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	infos := make([]storage.ListItem, 0)

	limit := 10
	prefix := c.Query("prefix")
	delimiter := ""
	//初始列举marker为空
	marker := c.Query("next_marker")

	entries, _, nextMarker, hashNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
	if err != nil {
		fmt.Println("list error,", err)
	}
	//print entries
	for _, entry := range entries {
		//fmt.Println(entry.Key)
		infos = append(infos, entry)

	}
	if hashNext {
		fmt.Println("nextMarker", nextMarker)
		marker = nextMarker
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: uint64(len(infos)),
		List:       infos,
		NextMaker:  marker,
	})

}
