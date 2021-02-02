package mergemap

import (
	"fmt"
	"testing"
)

func TestJsonMerge(t *testing.T) {
	dst := map[string]interface{}{
		"k3": "v3",
		"k4": map[string]interface{}{
			"k41": "v41",
			"k42": "v42",
		},
	}

	src := map[string]interface{}{
		"k5": "v5",
		"k4": map[string]interface{}{
			"kk41": "vv41",
			"kk42": "vv42",
		},
		"k3": "vv3",
	}
	m := JsonMerge(dst, src)
	fmt.Printf("%+v", m)
}
