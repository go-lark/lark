package lark

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertI18NCard(t *testing.T, c *I18NCardBlock, s string) {
	actMap := map[string]interface{}{}
	s2 := c.String()
	err := json.Unmarshal([]byte(s2), &actMap)
	assert.Nil(t, err)
	expMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(s), &expMap)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(actMap["header"], expMap["header"]))
	assert.True(t, reflect.DeepEqual(actMap["elements"], expMap["elements"]))
}

func TestI18NCardExample1(t *testing.T) {
	b := NewI18NCardBuilder()

	c := b.NewI18NCard().Purple().
		CnTitle("\U0001F973 抖音文创“双十一”年度大促").
		EnTitle("\U0001F973 Douyin Creativity \"Double Eleven\" Anniversary Sale").
		JpTitle("\U0001F973 Douyin Creativity \"Double Eleven\" Anniversary Sale")

	c.AddCnContent(
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
	)

	// 英语不好 机翻意思一下哈
	c.AddEnContent(
		b.Markdown("🔥 **Tiktok cultural and creative \"double 11\" annual best and lowest price activity opens today** 🔥 \n🔥Cross store 300-30 (unlimited stacking)\n🔥Store discount can be superimposed with platform full reduction: full 199-20 (superimposed cross store full reduction can be full 300-50)").
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
				b.Text("**🎉 Cute rabbit silicone night light**\nThree gear warm light easy adjustment, suitable for big friends and children!\nThe original price is 199 yuan, the activity price is 119 yuan, and the estimated hand price is **99 yuan**.[Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)").
					LarkMd(),
			),
		b.Hr(),
		b.Div().
			Extra(
				b.Img("img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g").Alt("图片"),
			).
			Text(
				b.Text("**🎉 One bag three use multifunctional Backpack**\nI'm good at hand, shoulder and cross car. Boys and girls love it!\nThe original price is 139 yuan, and the activity price is only **119 yuan**. [Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)").
					LarkMd(),
			),
		b.Hr(),
		b.Markdown("**🌟 Special benefits**:\nAdd 1 yuan to an order, and you will get a folder of Tiktok popular stems worth 11 yuan. The number is limited, and each ID is limited to 1 piece").
			Href("urlVal", b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN")),
		b.Action(
			b.Button(b.Text("Buy Now")).Primary().URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			b.Button(b.Text("View more coupons")).URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
		),
		b.Hr(),
		b.Note().
			AddText(b.Text("Time: November 1, 2021 ~ November 20, 2021")),
	)

	c.AddJpContent(
		b.Markdown("🔥 **Tiktok cultural and creative \"double 11\" annual best and lowest price activity opens today** 🔥 \n🔥Cross store 300-30 (unlimited stacking)\n🔥Store discount can be superimposed with platform full reduction: full 199-20 (superimposed cross store full reduction can be full 300-50)").
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
				b.Text("**🎉 Cute rabbit silicone night light**\nThree gear warm light easy adjustment, suitable for big friends and children!\nThe original price is 199 yuan, the activity price is 119 yuan, and the estimated hand price is **99 yuan**.[Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)").
					LarkMd(),
			),
		b.Hr(),
		b.Div().
			Extra(
				b.Img("img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g").Alt("图片"),
			).
			Text(
				b.Text("**🎉 One bag three use multifunctional Backpack**\nI'm good at hand, shoulder and cross car. Boys and girls love it!\nThe original price is 139 yuan, and the activity price is only **119 yuan**. [Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)").
					LarkMd(),
			),
		b.Hr(),
		b.Markdown("**🌟 Special benefits**:\nAdd 1 yuan to an order, and you will get a folder of Tiktok popular stems worth 11 yuan. The number is limited, and each ID is limited to 1 piece").
			Href("urlVal", b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN")),
		b.Action(
			b.Button(b.Text("Buy Now")).Primary().URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			b.Button(b.Text("View more coupons")).URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
		),
		b.Hr(),
		b.Note().
			AddText(b.Text("Time: November 1, 2021 ~ November 20, 2021")),
	)

	exp := "{\n  \"config\": {\n    \"wide_screen_mode\": true,\n    \"enable_forward\": true,\n    \"update_multi\": false\n  },\n  \"header\": {\n    \"title\": {\n      \"tag\": \"plain_text\",\n      \"i18n\": {\n        \"zh_cn\": \"🥳 抖音文创“双十一”年度大促\",\n        \"en_us\": \"🥳 Douyin Creativity \\\"Double Eleven\\\" Anniversary Sale\",\n        \"ja_jp\": \"🥳 Douyin Creativity \\\"Double Eleven\\\" Anniversary Sale\"\n      }\n    },\n    \"template\": \"purple\"\n  },\n  \"i18n_elements\": {\n    \"en_us\": [\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"🔥 **Tiktok cultural and creative \\\"double 11\\\" annual best and lowest price activity opens today** 🔥 \\n🔥Cross store 300-30 (unlimited stacking)\\n🔥Store discount can be superimposed with platform full reduction: full 199-20 (superimposed cross store full reduction can be full 300-50)\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"img\",\n        \"img_key\": \"img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg\",\n        \"alt\": {\n          \"tag\": \"plain_text\",\n          \"content\": \"\"\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 Cute rabbit silicone night light**\\nThree gear warm light easy adjustment, suitable for big friends and children!\\nThe original price is 199 yuan, the activity price is 119 yuan, and the estimated hand price is **99 yuan**.[Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_a4a1c992-7dba-42e8-8f07-859b315bcbeg\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 One bag three use multifunctional Backpack**\\nI'm good at hand, shoulder and cross car. Boys and girls love it!\\nThe original price is 139 yuan, and the activity price is only **119 yuan**. [Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"**🌟 Special benefits**:\\nAdd 1 yuan to an order, and you will get a folder of Tiktok popular stems worth 11 yuan. The number is limited, and each ID is limited to 1 piece\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"action\",\n        \"actions\": [\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"Buy Now\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"primary\"\n          },\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"View more coupons\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"default\"\n          }\n        ]\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"note\",\n        \"elements\": [\n          {\n            \"tag\": \"plain_text\",\n            \"content\": \"Time: November 1, 2021 ~ November 20, 2021\"\n          }\n        ]\n      }\n    ],\n    \"ja_jp\": [\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"🔥 **Tiktok cultural and creative \\\"double 11\\\" annual best and lowest price activity opens today** 🔥 \\n🔥Cross store 300-30 (unlimited stacking)\\n🔥Store discount can be superimposed with platform full reduction: full 199-20 (superimposed cross store full reduction can be full 300-50)\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"img\",\n        \"img_key\": \"img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg\",\n        \"alt\": {\n          \"tag\": \"plain_text\",\n          \"content\": \"\"\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 Cute rabbit silicone night light**\\nThree gear warm light easy adjustment, suitable for big friends and children!\\nThe original price is 199 yuan, the activity price is 119 yuan, and the estimated hand price is **99 yuan**.[Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_a4a1c992-7dba-42e8-8f07-859b315bcbeg\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 One bag three use multifunctional Backpack**\\nI'm good at hand, shoulder and cross car. Boys and girls love it!\\nThe original price is 139 yuan, and the activity price is only **119 yuan**. [Detail>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"**🌟 Special benefits**:\\nAdd 1 yuan to an order, and you will get a folder of Tiktok popular stems worth 11 yuan. The number is limited, and each ID is limited to 1 piece\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"action\",\n        \"actions\": [\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"Buy Now\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"primary\"\n          },\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"View more coupons\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"default\"\n          }\n        ]\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"note\",\n        \"elements\": [\n          {\n            \"tag\": \"plain_text\",\n            \"content\": \"Time: November 1, 2021 ~ November 20, 2021\"\n          }\n        ]\n      }\n    ],\n    \"zh_cn\": [\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"🔥 **抖音文创“双十一”全年最优惠最低价活动今日开启** 🔥 \\n🔥跨店每满300-30（可无限叠加）\\n🔥店铺优惠可以和平台满减叠加：满199-20（叠加跨店满减，可以满300-50哦）\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"img\",\n        \"img_key\": \"img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg\",\n        \"alt\": {\n          \"tag\": \"plain_text\",\n          \"content\": \"\"\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 萌兔硅胶小夜灯**\\n三档暖光轻松调节，适合大朋友更适合小朋友！\\n原价199元，活动价119元，凑单后预估到手价**99元**。[查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_a4a1c992-7dba-42e8-8f07-859b315bcbeg\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"div\",\n        \"text\": {\n          \"tag\": \"lark_md\",\n          \"content\": \"**🎉 一包三用多功能双肩背包**\\n手提、双肩、斜挎我都在行，男生女生都爱的吧!\\n原价139元，活动价仅需**119元**。 [查看详情>>](https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN)\"\n        },\n        \"extra\": {\n          \"tag\": \"img\",\n          \"img_key\": \"img_v2_678c0e33-bf1e-45ca-9ef6-98386302206g\",\n          \"alt\": {\n            \"tag\": \"plain_text\",\n            \"content\": \"图片\"\n          }\n        }\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"markdown\",\n        \"content\": \"**🌟 特别福利：**\\n下单加1元就送价值11元抖音热门梗文件夹，数量有限，每ID限1件\",\n        \"href\": {\n          \"urlVal\": {\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\"\n          }\n        }\n      },\n      {\n        \"tag\": \"action\",\n        \"actions\": [\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"立即抢购\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"primary\"\n          },\n          {\n            \"tag\": \"button\",\n            \"text\": {\n              \"tag\": \"plain_text\",\n              \"content\": \"查看更多优惠券\"\n            },\n            \"url\": \"https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN\",\n            \"type\": \"default\"\n          }\n        ]\n      },\n      {\n        \"tag\": \"hr\"\n      },\n      {\n        \"tag\": \"note\",\n        \"elements\": [\n          {\n            \"tag\": \"plain_text\",\n            \"content\": \"活动时间：2021年11月1日~2021年11月20日\"\n          }\n        ]\n      }\n    ]\n  }\n}"
	assertI18NCard(t, c, exp)
}
