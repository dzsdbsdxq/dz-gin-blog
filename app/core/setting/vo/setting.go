package vo

import "github.com/dzsdbsdxq/dz-gin-blog/app/core/setting/model"

type SettingVO struct {
	Required int    `json:"required"`
	Key      string `json:"key"`
	Val      string `json:"val"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Type     string `json:"type"`
	Comp     string `json:"comp"`
}

func InitializeData() []*model.SysSetting {
	return []*model.SysSetting{
		{
			Required: 2,
			Key:      "BlogTitle",
			Val:      "我的博客",
			Name:     "博客标题",
			Desc:     "",
			Type:     "normal",
			Comp:     "text",
		},
		{
			Required: 2,
			Key:      "BlogLink",
			Val:      "http://127.0.0.1:10000",
			Name:     "博客地址",
			Desc:     "",
			Type:     "normal",
			Comp:     "text",
		},
		{
			Required: 1,
			Key:      "BlogLogo",
			Val:      "",
			Name:     "Logo",
			Desc:     "",
			Type:     "normal",
			Comp:     "attachment",
		},
		{
			Required: 1,
			Key:      "BlogFavicon",
			Val:      "",
			Name:     "Favicon",
			Desc:     "",
			Type:     "normal",
			Comp:     "attachment",
		},
		{
			Required: 1,
			Key:      "BlogFooterInfo",
			Val:      "",
			Name:     "页脚信息",
			Desc:     "",
			Type:     "normal",
			Comp:     "textarea",
		},

		{
			Required: 1,
			Key:      "BlogSeoKeywords",
			Val:      "",
			Name:     "关键词",
			Desc:     "",
			Type:     "seo",
			Comp:     "text",
		},
		{
			Required: 1,
			Key:      "BlogSeoDescription",
			Val:      "",
			Name:     "博客描述",
			Desc:     "",
			Type:     "seo",
			Comp:     "textarea",
		},

		{
			Required: 1,
			Key:      "BlogSeoDescription",
			Val:      "",
			Name:     "博客描述",
			Desc:     "",
			Type:     "post",
			Comp:     "textarea",
		},

		{
			Required: 1,
			Key:      "EnableEmail",
			Val:      "0",
			Name:     "是否启用",
			Desc:     "",
			Type:     "email",
			Comp:     "radio",
		},
		{
			Required: 2,
			Key:      "SmtpAddress",
			Val:      "",
			Name:     "SMTP地址",
			Desc:     "",
			Type:     "email",
			Comp:     "text",
		},
		{
			Required: 2,
			Key:      "SmtpSendProtocol",
			Val:      "smtp",
			Name:     "发送协议",
			Desc:     "",
			Type:     "email",
			Comp:     "text",
		},
		{
			Required: 2,
			Key:      "SmtpSslPort",
			Val:      "465",
			Name:     "SSL端口",
			Desc:     "",
			Type:     "email",
			Comp:     "text",
		},
		{
			Required: 2,
			Key:      "SmtpEmailAccount",
			Val:      "",
			Name:     "邮箱账号",
			Desc:     "",
			Type:     "email",
			Comp:     "text",
		},
		{
			Required: 2,
			Key:      "SmtpEmailPassword",
			Val:      "",
			Name:     "邮箱密码",
			Desc:     "",
			Type:     "email",
			Comp:     "password",
		},
		{
			Required: 2,
			Key:      "SmtpEmailName",
			Val:      "",
			Name:     "发件人",
			Desc:     "",
			Type:     "email",
			Comp:     "text",
		},

		{
			Required: 1,
			Key:      "BlogCustomHead",
			Val:      "",
			Name:     "自定义全局head",
			Desc:     "",
			Type:     "other",
			Comp:     "textarea",
		},
		{
			Required: 1,
			Key:      "BlogCustomPostHead",
			Val:      "",
			Name:     "自定义内容页head",
			Desc:     "",
			Type:     "other",
			Comp:     "textarea",
		},
		{
			Required: 1,
			Key:      "BlogStatisticsCode",
			Val:      "",
			Name:     "统计代码",
			Desc:     "",
			Type:     "other",
			Comp:     "textarea",
		},
		{
			Required: 1,
			Key:      "BlogEnableTheme",
			Val:      "",
			Name:     "启用的主题",
			Desc:     "",
			Type:     "readonly",
			Comp:     "text",
		},
	}
}
