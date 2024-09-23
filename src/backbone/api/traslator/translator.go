package traslator

import "github.com/ali-mahdavi-bn/service-site/src/backbone/helper/utils"

var MapTranslate = map[string]map[string]map[string]string{
	"error": {
		"fa": utils.ReadJSONFile("/Users/ali/Desktop/go-p/cosmic-go/src/backbone/api/traslator/phrases/error/fa.json"),
	},
}
