package secret_reader

import (
	"crypto/rsa"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrivateKey(t *testing.T) {
	nKey := new(big.Int)
	nKey.SetString("test", 10)

	tests := []struct {
		name        string
		pemData     []byte
		privatePath string
		ExpectedKey *rsa.PrivateKey
		wantErr     bool
	}{
		{
			name:        "private key not found",
			privatePath: "/test",
			ExpectedKey: nil,
			wantErr:     true,
		},
		{
			name:        "private key found",
			privatePath: "mock_private.pem",
			ExpectedKey: &rsa.PrivateKey{
				PublicKey: rsa.PublicKey{
					N: nKey,
					E: 65537,
				},
			},
			wantErr: false,
		}}

	tempFileInit()
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			p := pemReader{
				secretPath: tt.privatePath,
			}

			privateKey, err := p.GetPrivateKey()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NotNil(t, privateKey)
			}
		})
	}

	tempFileCleanup()
}

func TestGetPublicKey(t *testing.T) {
	tests := []struct {
		name        string
		pemData     []byte
		publicPath  string
		expectedKey *rsa.PublicKey
		wantErr     bool
	}{
		{
			name:       "success ReadPublicKeyPKCS8",
			publicPath: "mock_public.pem",
			wantErr:    false,
		},
		{
			name:        "error InvalidPublicKeyFormat",
			pemData:     []byte("invalid_pem_data_here"),
			publicPath:  "mock",
			expectedKey: nil,
			wantErr:     true,
		},
		{
			name:        "error invalid public path",
			publicPath:  "invalid",
			expectedKey: nil,
			wantErr:     true,
		},
	}

	for _, test := range tests {
		tempFileInit()

		t.Run(test.name, func(t *testing.T) {
			p := pemReader{
				publicPath: test.publicPath,
			}

			pubKey, err := p.GetPublicKey()

			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NotNil(t, pubKey)
			}
		})

		tempFileCleanup()
	}

}

func tempFileInit() {
	mockPublic := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyP7eGiN7ExU0xy8aXXdR
xaiqDFMnhxWvToqv/zhpTDg7EeSR7BfiLkX2q/1PB13uo4NI+bC6f+Q/GouTIqr8
lycXdY/BLiirP6nkK7UtLrKqiWCQzM3t2F9XiB6w+SztggYiYT6eoCIVDkn4SJsS
RDGjA4qLp2ZqCvT+uiaxCo3nUW7lYA5+D7cUi691c2GEWT2uuQ+R843SJjrqLbmz
Dd5m1kVoYnE9qDNvX03t3CZc0QHGjponD3UbTU39P1pPowg3oicaC7ismfyFElbs
yJeIJzBijPyoXkjzpU0cFZvnh1o69vsQeCmX9NBTaCU7vk11BHtKXWFX8IqHmr1B
3QIDAQAB
-----END PUBLIC KEY-----`

	mockPrivate := `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDI/t4aI3sTFTTH
Lxpdd1HFqKoMUyeHFa9Oiq//OGlMODsR5JHsF+IuRfar/U8HXe6jg0j5sLp/5D8a
i5MiqvyXJxd1j8EuKKs/qeQrtS0usqqJYJDMze3YX1eIHrD5LO2CBiJhPp6gIhUO
SfhImxJEMaMDiounZmoK9P66JrEKjedRbuVgDn4PtxSLr3VzYYRZPa65D5HzjdIm
OuotubMN3mbWRWhicT2oM29fTe3cJlzRAcaOmicPdRtNTf0/Wk+jCDeiJxoLuKyZ
/IUSVuzIl4gnMGKM/KheSPOlTRwVm+eHWjr2+xB4KZf00FNoJTu+TXUEe0pdYVfw
ioeavUHdAgMBAAECggEBAJNI8EgHJ/Db4Uj0Y0WKYgmNhs5xQM3kPgo35rAHDmIj
8mUyMRvohH2UFyYBASBM3MpFMfyGXKPLBdLV5IPK+D1rD+294bmJY7PLMsA0i19k
3UK92F27qUac1u+QTe7J1WEqTZck4+hEEVnfKmlJ+SCvntzBcYTBr4NH9EFEiQdJ
l1fMfMxS6HgpIcY2wxnypqjpeC/LdG4543iGpP8zbnunxW43yFgkNC8AJ/KuY9eL
dX/bES25F0mmhQRruzpFczre+4S2tiVhMq73VuyXJiYIvaCxA1JHQjQBCjmkJ7r7
lNKZYlf7Qfcg5FZ9njGdFWpDeM40vyiGYZbKVV6X4aECgYEA9n84ldyPD4U8Kmnz
56pJ5lK+2ztC90pXCNS+vpCqg3VdLf7eE1nUAp2f6JBA/qm3kMIfju49bpTFM6ZD
Kj4jI8WI7JdBy1gfNls9HF5nVtoHzTHG9E3wT0YAOPHmyCfohpfFIFaWX63VwZ15
/jKP9OqW7aeMLoOWv9Mg+WQbVrMCgYEA0L6TKuttEZmyLDFhb+4T52Yl8ofNLbxe
6BhqDy44xwjBsi2a0BDc9l1/kv2ACsHEfMNFO/jW1IZ4C+2S3WDSoH+ff7uFVxS0
1wGrNNM/ch/zX3oLzX+ZevAYSys95IshMQMfoePp533ZYk504nRsxM4JVGSUoSHK
MvmnlRMJzS8CgYEAiJOE7sP+IENaSsXZ9opL1+oRBbeYKxxtjN8TsNLHJ39n2YxV
z7L93VUovNrwqCmxI+vrQG6QayzS9wMwQ7+aCL/yVeSY9+ojoSJ8gbNs3pp/qBnk
eoiUldfbV7HwhQZXt/tvpbNULj9LKLPwW//382PnrFYhPcR7Sl3Y71WgMDECgYBv
DzXVe/RHjPJSuOMSXiSQ1LQT2VS8pKAJ9BNZiEoE+w+y8LiRQqeNHCmn1t+s2XLk
vi+zvKzv3as5DWk6By2I3t3JY8eJkSa1zdl8/XegDIe7oH9vEhhiZCNIuvTvB2bd
YMAPrebglwB1YTCm2zKTcttb3zeEkym0/Ua/9aUdWQKBgQC43+cW1iWcw3U8LhO7
3Wkys/d07xkZEoEGJRFymqhyUNIBRY3fe4ukQOp3Qq0bff0Oj5YBDD6n+TES9nSh
wrKPt/kduqwqr9Ob4SwaUwX18rQlQwRxo2O3EQLcIIMiMVSuDKdc68AT8uc1Dbqn
gmyBqtD/AVn+rKit0f7HDuOrcw==
-----END PRIVATE KEY-----`

	err := os.WriteFile("mock_public.pem", []byte(mockPublic), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("mock_private.pem", []byte(mockPrivate), 0644)
	if err != nil {
		panic(err)
	}
}

func tempFileCleanup() {
	os.Remove("mock_private.pem")
	os.Remove("mock_public.pem")
}
