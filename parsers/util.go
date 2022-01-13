package parsers

import "github.com/mitchellh/mapstructure"

func ExplainToStandard(body map[string]interface{}) *StandardEntity {
	_, ok := body["msgtype"]
	if ok {
		// 钉钉
		var dingtalkEntity DingtalkEntity
		mapstructure.Decode(body, &dingtalkEntity)
		return dingtalkEntity.ToStandard()
	}

	_, ok = body["user"]
	if ok {
		// pushover
		var pushoverEntity PushoverEntity
		mapstructure.Decode(body, &pushoverEntity)
		return pushoverEntity.ToStandard()
	}
	return nil
}
