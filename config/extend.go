package config

var ExtConfig *Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！

type Extend struct {
	Ldap       *Ldap // 这里配置对应配置文件的结构即可
	Notify     *Notify
	AesSecrets *AesSecrets // 对称加密 key
}

type Ldap struct {
	Host         string `yaml:"host"`
	Tls          bool   `yaml:"tls"`
	Port         int    `yaml:"port"`
	BindDN       string `yaml:"bindDN"`
	BindPassword string `yaml:"bindPassword"`
	SearchDomain string `yaml:"searchDomain"`
}

type Notify struct {
	BotCredit *BotCredit `yaml:"bot_credit"`
}

type BotCredit struct {
	AppID     string `yaml:"appid" `
	AppSecret string `yaml:"app_secret" `
}

type AesSecrets struct {
	Key string `yaml:"key"`
}
