package core_test

import (
	"testing"

	"github.com/keenoho/go-core"
)

func TestUtil(t *testing.T) {
	str := `Key: 'SignatureParams.Sign' Error:Field validation for 'Sign' failed on the 'required' tag
Key: 'SignatureParams.RandId' Error:Field validation for 'RandId' failed on the 'required' tag
Key: 'SignatureParams.Ts' Error:Field validation for 'Ts' failed on the 'required' tag
Key: 'SignatureParams.Ttl' Error:Field validation for 'Ttl' failed on the 'required' tag`

	tags := core.FilterOutMissingTags(str)
	t.Log(tags)
}
