package domain

import "errors"

// 401
var ErrAuthFailed = errors.New("error authentication failed") //
var ErrEmailNotVerified = errors.New("email not verified")

// 400
var UsernameOrEmailTaken = errors.New("username or Email al ready taken") //
var OtpInvalid = errors.New("otp invalid")                                //
var InsufficientBalance = errors.New("insufficient balance")
var EmailAlreadyVerified = errors.New("email sudah terifikasi")
var SelfTransfer = errors.New("transfer ke diri sendiri tidak diperbolehkan")
var ValidationError = errors.New("")

// 404
var AccountNotFound = errors.New("account not found")
var InquiryNotFound = errors.New("inquiry Not found")
var UserNotFound = errors.New("user Not found")
var TemplNotFound = errors.New("template not found")
var TopUpReqNotFound = errors.New("topup request not found")
var PinInvalid = errors.New("pin Invalid")
