package techniques

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kushao1267/Facade/facade/utils"
)

// WeChatTechnique ...
type WeChatTechnique BaseTechnique

func (t WeChatTechnique) setName(name string) {
	t.Name = name
}

// GetName wechat get name method
func (t WeChatTechnique) GetName() string {
	return t.Name
}

// Extract wechat extract method
func (t WeChatTechnique) Extract(html string) DirtyExtracted {
	extracted := GetEmptyDirtyExtracted()
	t.setName("WeChatTechnique")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return extracted
	}
	jsCode := ""
	jsCode = doc.Find("script").Text()
	/* 自定义提取信息 */
	// title
	titles := utils.MatchOneOf(jsCode, `msg_title = "(.+?)";`)
	if len(titles) > 1 {
		extracted[TitlesField] = append(extracted[TitlesField], titles[1:]...)
	}
	// image
	images := utils.MatchOneOf(jsCode, `msg_cdn_url = "(.+?)";`)
	if len(images) > 1 {
		extracted[ImagesField] = append(extracted[ImagesField], images[1:]...)
	}
	// description
	doc.Find("section").Each(func(i int, selection *goquery.Selection) {
		if i < 3 {
			extracted[DescriptionsField] = append(extracted[DescriptionsField], selection.Text())
		}
	})

	descriptions := utils.MatchOneOf(jsCode, `msg_desc = "(.+?)";`)
	if len(descriptions) > 1 {
		extracted[DescriptionsField] = append(extracted[DescriptionsField], descriptions[1:]...)
	}

	return extracted
}
