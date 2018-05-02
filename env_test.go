package gomicSocMedShared

import (
	"os"
	"testing"
)

func TestEnv_reads_environment_variables(t *testing.T) {
	expected := "testvalue"
	os.Setenv("testkey", "testvalue")

	actual := Env("testkey")
	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestIsProd_identifies_production_environment(t *testing.T) {
	os.Setenv(GOMIC_STAGE_PROD, "prod")
	os.Setenv(GOMIC_STAGE, "notprod")

	expected := false
	actual := IsProd()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}

	os.Setenv(GOMIC_STAGE_PROD, "prod")
	os.Setenv(GOMIC_STAGE, "prod")

	expected = true
	actual = IsProd()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
