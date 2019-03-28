package main

import (
	"fmt"
	"time"
)

//          appid    左右眼     块号   byte数组
var fmap map[int64]map[int32]map[int32][]*[512]byte

func AllocMap(appId int64, eyeFlag int32, block int32,size int)  {
	arrayBlock := make([]*[512]byte,size)
	fmap[appId][eyeFlag][block] = arrayBlock;
}

/**
  添加虹膜数据
 */
func AddFeature(appId int64, eyeFlag int32, block int32, feature *[512]byte) bool {

	fmap[appId][eyeFlag][block][0] = feature;
	return true
}

/**
  获取一整块的虹膜数据
 */
func GetFeature(appId int64, eyeFlag int32, block int32) (featureArray []*[512]byte) {

	if _, ok := fmap[appId]; ok {
		if _, ok := fmap[appId][eyeFlag]; ok {
			if _, ok := fmap[appId][eyeFlag][block]; ok {
				featureArray = fmap[appId][eyeFlag][block]
			}
		}
	}
	return
}

func gos(u int)  {

	go func(u int) {
		fmt.Println(u)
	}(u)
}

func main() {

	for n:=1;n<10;n++ {

		go func(){
			fmt.Println(n)
		}()
	}

	time.Sleep(time.Hour)
}
