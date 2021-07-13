package main

import (
	"github.com/spf13/viper"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/cbirc"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/mee"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/miit"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/api/v1/sarm"
	"github.com/xiaogogonuo/cct-spider/internal/minserver/store"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/callback"
	"github.com/xiaogogonuo/cct-spider/internal/pkg/filter"
	"github.com/xiaogogonuo/cct-spider/pkg/config"
	"github.com/xiaogogonuo/cct-spider/pkg/encrypt/md5"
	"github.com/xiaogogonuo/cct-spider/pkg/logger"
	"sync"
)

var (
	minV      *viper.Viper
	filt      *filter.Filter
	urlKeyMap map[string]byte
)

func minConfig() *viper.Viper {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/min",
	}
	v, err := c.NewConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	minV = minConfig()
	filt = &filter.Filter{
		Filepath: "urlKey.txt",
		ThisUrlKey: make(map[string]byte),

	}
	urlKeyMap = filt.ReadUrlKey()
}

func ministries() {
	wg := &sync.WaitGroup{}
	urlChannel := make(chan *callback.UrlChan)   // url请求池
	infoChannel := make(chan *callback.InfoChan) // info请求池
	errChannel := make(chan *callback.InfoChan)  // 异常池
	message := make(chan *callback.Message)      // 数据池
	save := store.InsertIntoSQL                  // 保存数据的函数

	wg.Add(5)
	go miit.GetPageUrlList(minV.GetString("工业和信息化部"), urlChannel, wg)
	go sarm.GetPageUrlList(minV.GetString("国家市场监督管理总局"), urlChannel, wg)
	go mee.GetFirstUrl(minV.GetString("生态环境部"), urlChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会928"), infoChannel, wg)
	go cbirc.GetPageUrlList(minV.GetString("银保监会927"), infoChannel, wg)

	go func() {
		for v := range urlChannel {
			if _, ok := urlKeyMap[md5.MD5(v.Url)]; ok {
				logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
				continue
			}
			wg.Add(1)
			go v.GetUrlFunc(urlChannel, infoChannel, wg)
		}
	}()
	go func() {
		for v := range infoChannel {
			if _, ok := urlKeyMap[md5.MD5(v.Url)]; ok {
				logger.Info("Obtained, no need to update", logger.Field("url", v.Url))
				continue
			}
			wg.Add(1)
			go v.GetInfoFunc(errChannel, message, wg)
		}
	}()
	go func() {
		wg.Wait()
		close(urlChannel)
		close(infoChannel)
		close(message)
	}()
	save(filt, message)
	filt.SaveUrlKey()
}

func main() {
	ministries()
}
