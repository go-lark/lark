package lark

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-lark/lark/card"
	"github.com/stretchr/testify/assert"
)

func assertCard(t *testing.T, c *card.Block, s string) {
	actMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(c.String()), &actMap)
	assert.Nil(t, err)
	expMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(s), &expMap)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(actMap["header"], expMap["header"]))
	assert.True(t, reflect.DeepEqual(actMap["elements"], expMap["elements"]))
}

// 以下的例子均为卡片构造工具（https://open.feishu.cn/tool/cardbuilder?from=howtoguide）提供的公开模板，使用CardBuilder转写。
func TestCardExample1(t *testing.T) {
	b := NewCardBuilder()
	c := b.Card(
		b.Markdown("🔥 **抖音文创“双十一”全年最优惠最低价活动今日开启** 🔥 \n🔥跨店每满300-30（可无限叠加）\n🔥店铺优惠可以和平台满减叠加：满199-20（叠加跨店满减，可以满300-50哦）").
			Href("urlVal",
				b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			),
		b.Img("img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg"),
		b.Hr(),
		b.Div().
			Extra(
				b.Img("img_v2_a4a1c992-7dba-42e8-8f07-859b315bcbeg").Alt("图片"),
			).
			Text(
				b.Text("**🎉 萌兔硅胶小夜灯**\n三档暖光轻松调节，适合大朋友更适合小朋友！\n原价199元，活动价119元，凑单后预估到手价**99元**。[查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)").
					LarkMd(),
			),
		b.Hr(),
		b.Div().
			Extra(
				b.Img("img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g").Alt("图片"),
			).
			Text(
				b.Text("**🎉 一包三用多功能双肩背包**\n手提、双肩、斜挎我都在行，男生女生都爱的吧!\n原价139元，活动价仅需**119元**。 [查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)").
					LarkMd(),
			),
		b.Hr(),
		b.Markdown("**🌟 特别福利：**\n下单加1元就送价值11元抖音热门梗文件夹，数量有限，每ID限1件").
			Href("urlVal", b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN")),
		b.Action(
			b.Button(b.Text("立即抢购")).Primary().URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			b.Button(b.Text("查看更多优惠券")).URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
		),
		b.Hr(),
		b.Note().
			AddText(b.Text("活动时间：2021年11月1日~2021年11月20日")),
	).Title("\U0001F973 抖音文创“双十一”年度大促").Purple()
	exp := "{\n  \"config\": {\n    \"wide_screen_mode\": true\n  },\n  \"elements\": [\n    {\n      \"content\": \"🔥 **抖音文创“双十一”全年最优惠最低价活动今日开启** 🔥 \\n🔥跨店每满300-30（可无限叠加）\\n🔥店铺优惠可以和平台满减叠加：满199-20（叠加跨店满减，可以满300-50哦）\",\n      \"href\": {\n        \"urlVal\": {\n          \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n        }\n      },\n      \"tag\": \"markdown\"\n    },\n    {\n      \"alt\": {\n        \"content\": \"\",\n        \"tag\": \"plain_text\"\n      },\n      \"img_key\": \"img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg\",\n      \"tag\": \"img\"\n    },\n    {\n      \"tag\": \"hr\"\n    },\n    {\n      \"extra\": {\n        \"alt\": {\n          \"content\": \"图片\",\n          \"tag\": \"plain_text\"\n        },\n        \"img_key\": \"img_v2_a4a1c992-7dba-42e8-8f07-859b315bcbeg\",\n        \"tag\": \"img\"\n      },\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"**🎉 萌兔硅胶小夜灯**\\n三档暖光轻松调节，适合大朋友更适合小朋友！\\n原价199元，活动价119元，凑单后预估到手价**99元**。[查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)\",\n        \"tag\": \"lark_md\"\n      }\n    },\n    {\n      \"tag\": \"hr\"\n    },\n    {\n      \"extra\": {\n        \"alt\": {\n          \"content\": \"图片\",\n          \"tag\": \"plain_text\"\n        },\n        \"img_key\": \"img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g\",\n        \"tag\": \"img\"\n      },\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"**🎉 一包三用多功能双肩背包**\\n手提、双肩、斜挎我都在行，男生女生都爱的吧!\\n原价139元，活动价仅需**119元**。 [查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)\",\n        \"tag\": \"lark_md\"\n      }\n    },\n    {\n      \"tag\": \"hr\"\n    },\n    {\n      \"content\": \"**🌟 特别福利：**\\n下单加1元就送价值11元抖音热门梗文件夹，数量有限，每ID限1件\",\n      \"href\": {\n        \"urlVal\": {\n          \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n        }\n      },\n      \"tag\": \"markdown\"\n    },\n    {\n      \"actions\": [\n        {\n          \"tag\": \"button\",\n          \"text\": {\n            \"content\": \"立即抢购\",\n            \"tag\": \"plain_text\"\n          },\n          \"type\": \"primary\",\n          \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n        },\n        {\n          \"tag\": \"button\",\n          \"text\": {\n            \"content\": \"查看更多优惠券\",\n            \"tag\": \"plain_text\"\n          },\n          \"type\": \"default\",\n          \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n        }\n      ],\n      \"tag\": \"action\"\n    },\n    {\n      \"tag\": \"hr\"\n    },\n    {\n      \"elements\": [\n        {\n          \"content\": \"活动时间：2021年11月1日~2021年11月20日\",\n          \"tag\": \"plain_text\"\n        }\n      ],\n      \"tag\": \"note\"\n    }\n  ],\n  \"header\": {\n    \"template\": \"purple\",\n    \"title\": {\n      \"content\": \"\U0001F973 抖音文创“双十一”年度大促\",\n      \"tag\": \"plain_text\"\n    }\n  }\n}"
	assertCard(t, c, exp)
}

func TestExample2(t *testing.T) {
	b := NewCardBuilder()
	c := b.Card(
		b.Img("img_v2_bfd72a81-1533-4699-995d-12a675708d0g"),
		b.Div().Text(b.Text("你是否曾因为一本书而产生心灵共振，开始感悟人生？\n你有哪些想极力推荐给他人的珍藏好书？\n\n加入 **4·23 飞书读书节**，分享你的**挚爱书单**及**读书笔记**，**赢取千元读书礼**！\n\n📬 填写问卷，晒出你的珍藏好书\n😍 想知道其他人都推荐了哪些好书？马上[入群围观](https://open.feishu.cn/)\n📝 用[读书笔记模板](https://open.feishu.cn/)（桌面端打开），记录你的心得体会\n🙌 更有惊喜特邀嘉宾 4月12日起带你共读").LarkMd()),
		b.Action(
			b.Button(b.Text("立即推荐好书")).Primary().URL("https://open.feishu.cn/"),
			b.Button(b.Text("查看活动指南")).URL("https://open.feishu.cn/"),
		),
	).Turquoise().Title("📚晒挚爱好书，赢读书礼金")
	//fmt.Println(c.String())
	exp := "{\n  \"config\": {\n    \"wide_screen_mode\": true\n  },\n  \"elements\": [\n    {\n      \"alt\": {\n        \"content\": \"\",\n        \"tag\": \"plain_text\"\n      },\n      \"img_key\": \"img_v2_bfd72a81-1533-4699-995d-12a675708d0g\",\n      \"tag\": \"img\"\n    },\n    {\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"你是否曾因为一本书而产生心灵共振，开始感悟人生？\\n你有哪些想极力推荐给他人的珍藏好书？\\n\\n加入 **4·23 飞书读书节**，分享你的**挚爱书单**及**读书笔记**，**赢取千元读书礼**！\\n\\n📬 填写问卷，晒出你的珍藏好书\\n😍 想知道其他人都推荐了哪些好书？马上[入群围观](https://open.feishu.cn/)\\n📝 用[读书笔记模板](https://open.feishu.cn/)（桌面端打开），记录你的心得体会\\n🙌 更有惊喜特邀嘉宾 4月12日起带你共读\",\n        \"tag\": \"lark_md\"\n      }\n    },\n    {\n      \"actions\": [\n        {\n          \"tag\": \"button\",\n          \"text\": {\n            \"content\": \"立即推荐好书\",\n            \"tag\": \"plain_text\"\n          },\n          \"type\": \"primary\",\n          \"url\": \"https://open.feishu.cn/\"\n        },\n        {\n          \"tag\": \"button\",\n          \"text\": {\n            \"content\": \"查看活动指南\",\n            \"tag\": \"plain_text\"\n          },\n          \"type\": \"default\",\n          \"url\": \"https://open.feishu.cn/\"\n        }\n      ],\n      \"tag\": \"action\"\n    }\n  ],\n  \"header\": {\n    \"template\": \"turquoise\",\n    \"title\": {\n      \"content\": \"📚晒挚爱好书，赢读书礼金\",\n      \"tag\": \"plain_text\"\n    }\n  }\n}"
	assertCard(t, c, exp)
}

func TestExample3(t *testing.T) {
	b := NewCardBuilder()
	c := b.Card(
		b.Markdown("感谢同学使用物品借用服务并及时归还物品～").
			Href(
				"urlVal",
				b.URL().
					MultiHref(
						"https://developer.android.com/",
						"lark://msgcard/unsupported_action",
						"https://www.feishu.com",
					),
			),
		b.Div().
			Text(b.Text("请对本次服务和系统使用体验进行满意度评价：").LarkMd()).
			Extra(
				b.SelectMenu(
					b.Option("1").Text("很好"),
					b.Option("2").Text("好"),
					b.Option("3").Text("一般"),
					b.Option("4").Text("差"),
					b.Option("5").Text("很差"),
				).
					Placeholder("请选择满意度评价").
					Value(map[string]interface{}{"key": "value"}),
			),
		b.Div().Text(b.Text("如有其它疑问及反馈，请点击下方按钮联系行政值班号。").LarkMd()),
		b.Action(b.Button(b.Text("联系值班号")).Primary().URL("https://open.feishu.cn/")),
	).Orange().Title("物品借用满意度调查")
	exp := "{\n  \"config\": {\n    \"wide_screen_mode\": true\n  },\n  \"elements\": [\n    {\n      \"content\": \"感谢同学使用物品借用服务并及时归还物品～\",\n      \"href\": {\n        \"urlVal\": {\n          \"android_url\": \"https://developer.android.com/\",\n          \"ios_url\": \"lark://msgcard/unsupported_action\",\n          \"pc_url\": \"https://www.feishu.com\",\n          \"url\": \"https://www.feishu.com\"\n        }\n      },\n      \"tag\": \"markdown\"\n    },\n    {\n      \"extra\": {\n        \"options\": [\n          {\n            \"text\": {\n              \"content\": \"很好\",\n              \"tag\": \"plain_text\"\n            },\n            \"value\": \"1\"\n          },\n          {\n            \"text\": {\n              \"content\": \"好\",\n              \"tag\": \"plain_text\"\n            },\n            \"value\": \"2\"\n          },\n          {\n            \"text\": {\n              \"content\": \"一般\",\n              \"tag\": \"plain_text\"\n            },\n            \"value\": \"3\"\n          },\n          {\n            \"text\": {\n              \"content\": \"差\",\n              \"tag\": \"plain_text\"\n            },\n            \"value\": \"4\"\n          },\n          {\n            \"text\": {\n              \"content\": \"很差\",\n              \"tag\": \"plain_text\"\n            },\n            \"value\": \"5\"\n          }\n        ],\n        \"placeholder\": {\n          \"content\": \"请选择满意度评价\",\n          \"tag\": \"plain_text\"\n        },\n        \"tag\": \"select_static\",\n        \"value\": {\n          \"key\": \"value\"\n        }\n      },\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"请对本次服务和系统使用体验进行满意度评价：\",\n        \"tag\": \"lark_md\"\n      }\n    },\n    {\n      \"tag\": \"div\",\n      \"text\": {\n        \"content\": \"如有其它疑问及反馈，请点击下方按钮联系行政值班号。\",\n        \"tag\": \"lark_md\"\n      }\n    },\n    {\n      \"actions\": [\n        {\n          \"tag\": \"button\",\n          \"text\": {\n            \"content\": \"联系值班号\",\n            \"tag\": \"plain_text\"\n          },\n          \"type\": \"primary\",\n          \"url\": \"https://open.feishu.cn/\"\n        }\n      ],\n      \"tag\": \"action\"\n    }\n  ],\n  \"header\": {\n    \"template\": \"orange\",\n    \"title\": {\n      \"content\": \"物品借用满意度调查\",\n      \"tag\": \"plain_text\"\n    }\n  }\n}"
	assertCard(t, c, exp)
}
