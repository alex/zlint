// lint_sub_cert_key_usage_crl_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestCrlSignBitSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageInvalid.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_key_usage_crl_sign_bit_set"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCrlSignBitNotSet(t *testing.T) {
	inputPath := "../testlint/testCerts/subKeyUsageValid.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_key_usage_crl_sign_bit_set"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
