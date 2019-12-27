package config

//RPC Config
const (
	RpcServerHost = ":1234"
	ElasticHost = "http://www.vowcloud.cn:9200"
	ElasticIndex = "dating_profile"
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
	CrawlServiceHost1 = ":9000"

	//parse name
	//NilParseResult = "NilParseResult"
	FuncParser     = "FuncParser"
	CityUserParse = "CityUserParse"
	ProfileParse    = "ProfileParse"
	CityListParer  = "CityListParser"

	// RateLimitTime 毫秒
	Qps = 1
)