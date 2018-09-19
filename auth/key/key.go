package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	pKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4w5xhil8YFSLptRxzQsiJgQm7DxfVx7nEFAndQDw/7a1VfIf
hhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV4wrQAOSe6pDK5UP/xZQx8ygy70lG
fJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+KdFIsD7pt+FxV6mMmPbnAvDN+hF5
NwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfuiUIEX3GxhskzV/ROiS+R/NbQZlsfm
QqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTTtu3ys/hs5m6AqNPPkLKqp6R8iXF1
Lg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+GuwIDAQABAoIBADlemeKLMujoE80Y
WpSzXnJ6lBcWfgR2Q23EwuN2VG5YDONlZP+u5G8qKEyzO6hvNkYgn2DPuyS8VNR9
VT6OcMmIHtxK57he01UwZDzY3/IPUydQvWWZbd4lBy7y5Q1MUbAK29avF7cgxD6+
qwncBtusDJCzpLwYU1oR9ftkTyRXl8WzHUQ+/QILNnSCDsTrP8JsVaVxbd6FhKKn
5sSyqM+dX7mtvVAOcj0OJSHZiit7fk5QG9Pi/5iP4pCdZf42sImsr++2GFOezfJd
H5UU+ujTf+b4oGirnqgEDRrSr5IyykagWc07D2KJgyPzrkfFDxoB5C/ZC3C6C9AA
Xwzd+GECgYEA5SPDfCMVBRFkYBoxKgbWEElquGiPMDSe+p6QSlX24UXFv8gzdtbT
f33d27v2cpIOWYym3Er5JiSFq6oCr1cg9+mLP/tNc50sHrdHb8vRfn190nawFJHa
eOe0b3ZePUtAxdd1HaZgq4bNnLYSbi//spdHuu6E1jZrzcmbvIm7PJECgYEA/awp
rILMDvqHuGNlVr+kdcGfmFxA8y9Z1tZHLgqNjPQQlaOuyJn1cfYbIqghMLjk//Au
VQ5gfKLc2abHQaVQ2dLqV846eNQvr+cnLQUrUqk41IZuN0HTMbvLHgOLkQNdsUMs
1TmmPeMxh9X9cLqp7mZoY5CeWeWFOe3EJA1dZIsCgYEAklbf3yUMpJrx7wprQbrx
9Z7dwH5OjGve6JJh9oemT0LfQ1dZvtj+ZBr/mPkXMR6keX6Bhol/S2Ph1ruSUWck
0A/gdfFKCr9jUQ6eWgDif5UnyUUxuUFZNQRN0S3Yi+7GpFOxIUmDzagfIqmJZcPT
2rwQ/IqeXayN9vR+ONABu3ECgYAECn4PdXXytyL6WPsASsU/6vmz36RZO2Pe/ELe
BOUEXc7100mxgGJckmMURkFhGVDsktLqH/SBh8ak4PdDoHKNRcLd6zcbPaYU00XY
fcCW7IMvP4T59F586FTwAXZztO4FKODJ9MUlLz1WwJ3s8cxLM+5tx5v+Kp3YsmTx
fhUCyQKBgDCEkFexrqC2a1rHLh+pwTyvnE4JCVNt72FF8L51aEsG5tGGFvTvgUN6
IlRCYASNhUK/3+hu337uOSolKXu0W+dFnp1/OLo6sUkuhxWGx3YLwGJygjSrOl5f
3wIikQ0U/RjRr+/pI0/yw/w3Xcr7iUjei6SBxkiIeZL/749EcLNB
-----END RSA PRIVATE KEY-----`

	data, _ := pem.Decode([]byte(pKey))

	parsedKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		err = errors.Wrap(err, "Error parsing PrivateKey")
		return nil, err
	}

	return parsedKey, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {

	pKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4w5xhil8YFSLptRxzQsi
JgQm7DxfVx7nEFAndQDw/7a1VfIfhhzZlUYx6u+57kP4+JPhqLMl9hEPnJh2DMPV
4wrQAOSe6pDK5UP/xZQx8ygy70lGfJ6MVo7mkXKaofKobOhkFIOhqtLU/6CrzFl+
KdFIsD7pt+FxV6mMmPbnAvDN+hF5NwU6N61WGAZER8z7SSTgayGpuHdUKCdPwfui
UIEX3GxhskzV/ROiS+R/NbQZlsfmQqcBJ5FxhOtAVevi9s7x6LLTSQKopuuunSTT
tu3ys/hs5m6AqNPPkLKqp6R8iXF1Lg0DMeQlFHYwEo3oRweMNhfYRzC3ukioSf+G
uwIDAQAB
-----END PUBLIC KEY-----`

	data, _ := pem.Decode([]byte(pKey))

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		err = errors.Wrap(err, "Error parsing PublicKey")
		return nil, err
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Error converting PublicKey-String to RSA-PublicKey")
	}

	return rsaPub, nil
}
