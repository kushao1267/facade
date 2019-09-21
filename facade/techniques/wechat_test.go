package techniques

import (
	"github.com/kushao1267/Facade/facade/utils"
	"testing"
)

func TestWeChatTechnique_Extract(t *testing.T) {
	var technique WeChatTechnique

	_, html := utils.GetHtml("https://mp.weixin.qq.com/s/VRzeIxFO_sHTOHAyZRX7xw")

	extracted := technique.Extract(html)
	allEmpty := true
	for _, value := range extracted {
		if len(value) > 0 {
			allEmpty = false
		}
	}
	if allEmpty {
		t.Fail()
	}
	t.Log(extracted)
}
