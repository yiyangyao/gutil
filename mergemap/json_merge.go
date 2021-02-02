package mergemap

import "reflect"

const JsonMergeDepth int = 5

func JsonMerge(dst, src map[string]interface{}) map[string]interface{} {
	return mapMerge(dst, src, 0)
}

func mapMerge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > JsonMergeDepth {
		return dst
	}

	for key, srcVal := range src {

		if dstVal, ok := dst[key]; ok {
			srcMap, srcOk := assertMap(srcVal)
			dstMap, dstOk := assertMap(dstVal)

			if srcOk && dstOk {
				srcVal = mapMerge(dstMap, srcMap, depth+1)
			}
		}

		dst[key] = srcVal
	}

	return dst
}

func assertMap(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)

	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}

	return map[string]interface{}{}, false
}
