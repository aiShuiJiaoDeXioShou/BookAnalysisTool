package main

import (
	"testing"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// 爬虫练习bilibili
func TestPanchon(t *testing.T) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.bilibili.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"),
	)
	// 添加相关网站对应的Cookie值
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", "buvid3=F5851E4F-C6AF-C2A5-14A6-09D83A5C95ED30346infoc; _uuid=43A8DEC9-5E710-5E85-63109-13510104436EC231110infoc; buvid4=E13CA173-A467-9DAF-80B9-EFCD46C125B833062-022082422-Pk1O31qDhl6zNmF/YjG6D6h192oJs6ei0DogXwF+zb/1nd7ZT5nPCw%3D%3D; fingerprint=687b2ef47da61a0c35769dd6bb569ba4; buvid_fp_plain=undefined; DedeUserID=284563289; DedeUserID__ckMd5=5009e33c09d521a6; LIVE_BUVID=AUTO6716613522662624; b_ut=5; buvid_fp=277bf8d9143df6eca65e664b80fa6462; rpdid=0zbfAHUzwv|fU5EQ8Pc|2gE|3w1Or9di; is-2022-channel=1; hit-dyn-v2=1; CURRENT_BLACKGAP=0; CURRENT_QUALITY=0; nostalgia_conf=-1; b_nut=100; SESSDATA=43929334%2C1678964732%2Ceb5e2%2A91; bili_jct=cf0ae3e794e5b39e5f636c1f5b6a134f; bp_video_offset_284563289=710519025523753000; go_old_video=1; i-wanna-go-feeds=-1; i-wanna-go-back=2; CURRENT_FNVAL=4048; innersign=0; b_lsid=CE944CFD_183841D9D52; PVID=2")
	})
	// 将自己伪装成普通浏览器
	extensions.RandomUserAgent(c)
	c.Wait()
	// 来开始爬bilibili

}
