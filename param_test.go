package broccli

import (
	"testing"
)

func h(c *CLI) int {
	return 0
}

func TestParamValidationBasic(t *testing.T) {
	p := &param{}
	if p.validateValue("") != nil {
		t.Errorf("Empty param should validate")
	}

	for _, typ := range []int64{TypeInt, TypeFloat, TypeAlphanumeric, TypePathFile} {
		p.valueType = typ
		if p.validateValue("") != nil {
			t.Errorf("Empty param should validate")
		}
	}

	p.valueType = TypeInt
	if p.validateValue("48") != nil {
		t.Errorf("Int param should validate")
	}
	if p.validateValue("aa") == nil {
		t.Errorf("Int param should not validate string")
	}

	p.valueType = TypeFloat
	if p.validateValue("48.998") != nil {
		t.Errorf("Float param should validate")
	}
	if p.validateValue("48") == nil {
		t.Errorf("Float param should not validate int")
	}
	if p.validateValue("aa") == nil {
		t.Errorf("Float param should not validate string")
	}

	p.valueType = TypeAlphanumeric
	if p.validateValue("a123aaAEz") != nil {
		t.Errorf("Alphanumeric param should validate")
	}
	if p.validateValue("a.z") == nil {
		t.Errorf("Alphanumeric param should not validate")
	}

	p.valueType = TypePathFile
	if p.validateValue("anything/here") != nil {
		t.Errorf("TypePathFile param should validate")
	}
}

func TestParamValidationRequired(t *testing.T) {
	p := &param{
		flags: IsRequired,
	}
	if p.validateValue("") == nil {
		t.Errorf("Empty param with IsRequired should not validate")
	}
	if p.validateValue("aa") != nil {
		t.Errorf("Param with IsRequired should validate")
	}
}