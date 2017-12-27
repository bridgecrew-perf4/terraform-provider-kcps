package kcps

import (
	gk "github.com/uesyn/gokcps"
)

func flattenTags(tags []gk.Tag) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(tags))
	for _, t := range tags {
		tagMap := make(map[string]interface{})
		tagMap["key"] = t.Key
		tagMap["value"] = t.Value
		result = append(result, tagMap)
	}
	return result
}

func convertStringArrToInterface(strs []string) []interface{} {
	arr := make([]interface{}, len(strs))
	for i, str := range strs {
		arr[i] = str
	}
	return arr
}

func remove(strings []string, search string) []string {
	var result []string
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
