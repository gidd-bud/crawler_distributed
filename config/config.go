package config

const (
	// Elastic search
	ElasticIndex = "dating_profile"

	// RPC endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Rate limiting
	Qps = 20
)
