package auth

import (
    "github.com/pquerna/otp/totp"
    "github.com/skip2/go-qrcode"
    // "time"
)

func Generate2FASecret(email string) (string, []byte, error) {
    key, err := totp.Generate(totp.GenerateOpts{
        Issuer:      "go-webapp-boilerplate",
        AccountName: email,
    })
    if err != nil {
        return "", nil, err
    }

    // Generate QR code image
    qrCode, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
    if err != nil {
        return "", nil, err
    }

    return key.Secret(), qrCode, nil
}

func Validate2FACode(secret, code string) bool {
    return totp.Validate(code, secret)
}
