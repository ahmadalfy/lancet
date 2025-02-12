package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestRsaEncrypt(t *testing.T) {
	GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	data := []byte("hello world")
	encrypted := RsaEncrypt(data, "rsa_public.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private.pem")

	if string(data) != string(decrypted) {
		internal.LogFailedTestInfo(t, "RsaEncrypt/RsaDecrypt", string(data), string(data), string(decrypted))
		t.FailNow()
	}
}
