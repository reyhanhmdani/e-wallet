package service

import (
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserService struct {
	UserRepository  domain.UserRepository
	CacheRepository domain.CacheRepository
	EmailService    domain.EmailService
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository, emailService domain.EmailService) domain.UserService {
	return &UserService{
		UserRepository:  userRepository,
		CacheRepository: cacheRepository,
		EmailService:    emailService,
	}
}

func (u UserService) AllUsers(c context.Context) ([]domain.Users, error) {
	users, err := u.UserRepository.AllUsers(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserService) Authenticate(c context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	//user, err := u.UserRepository.FindByUsernameOrEmail(c, req.Username, req.Email)
	criteria := map[string]interface{}{"username": req.Username, "email": req.Email}
	user, err := u.UserRepository.FindUser(c, criteria, false)
	if err != nil {
		logrus.Error(err)
		return dto.AuthRes{}, domain.OtpInvalid
	}
	if user == (domain.User{}) {
		logrus.Error(err)
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logrus.Error(err)
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	if !user.EmailVerifiedAtDB.Valid {
		return dto.AuthRes{}, domain.ErrEmailNotVerified
	}

	userData := dto.UserData{
		ID:       user.ID,
		Fullname: user.Fullname,
		Phone:    user.Phone,
		Username: user.Username,
		Email:    user.Email,
	}

	//jwt
	token, err := util.GenerateJWTToken(userData)
	if err != nil {
		logrus.Error(err)
		return dto.AuthRes{}, err
	}

	// Simpan token di cache atau penyimpanan yang sesuai
	u.CacheRepository.Set("user:"+token, []byte(token))

	return dto.AuthRes{
		UserID: user.ID,
		Token:  token,
	}, nil
}

func (u UserService) ValidateTokenWithCache(c context.Context, token string) (dto.UserData, error) {
	data, err := u.CacheRepository.Get("user:" + token)
	if err != nil {
		logrus.Error(err)
		return dto.UserData{}, domain.ErrAuthFailed
	}
	var user domain.User
	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		Fullname: user.Fullname,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}
func (u UserService) ValidateTokenJwt(c context.Context, token string) (dto.UserData, error) {
	// Parse token
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return util.JwtKey, nil
	})

	// Handle parsing errors
	if err != nil || !parsedToken.Valid {
		logrus.Error("Handle parsing error", err)
		return dto.UserData{}, domain.ErrAuthFailed
	}

	// Extract email from claims
	email, ok := claims["email"].(string)
	if !ok {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	// Retrieve user data based on email
	criteria := map[string]interface{}{"email": email}
	//user, err := u.UserRepository.FindByEmail(c, email)
	user, err := u.UserRepository.FindUser(c, criteria, false)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	// Convert domain.User to dto.UserData
	userData := dto.UserData{
		ID:       user.ID,
		Fullname: user.Fullname,
		Phone:    user.Phone,
		Username: user.Username,
		Email:    user.Email,
	}

	return userData, nil
}

func (u UserService) Register(c context.Context, req dto.UserCreateReq) (dto.UserRegisterRes, error) {
	// validasi
	if err := dto.ValidateUserRegisterReq(req); err != nil {
		logrus.Error(err)
		return dto.UserRegisterRes{}, err
	}
	//exist, err := u.UserRepository.FindByUsernameOrEmail(c, req.Username, req.Email)
	criteria := map[string]interface{}{"username": req.Username, "email": req.Email}
	exist, err := u.UserRepository.FindUser(c, criteria, false)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	if exist != (domain.User{}) {
		return dto.UserRegisterRes{}, domain.UsernameOrEmailTaken
	}
	if !strings.HasSuffix(req.Email, "@gmail.com") {
		return dto.UserRegisterRes{}, errors.New("invalid email format. Email must be a @gmail.com address")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	user := domain.User{
		Fullname: req.Fullname,
		Phone:    req.Phone,
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	err = u.UserRepository.Insert(c, &user)
	if err != nil {
		logrus.Error("error di Service user: ", err)
		return dto.UserRegisterRes{}, err
	}

	otpCode := util.GenerateRandomNumber(4)
	referenceID := util.GenerateRandomString(16)

	logrus.Info("your otp : ", otpCode)
	_ = u.CacheRepository.Set("otp:"+referenceID, []byte(otpCode))
	_ = u.CacheRepository.Set("user-ref:"+referenceID, []byte(user.Username))
	err = u.EmailService.Send(req.Email, "OTP CODE", "Otp Anda "+otpCode)
	if err != nil {
		logrus.Error("gagal send : ", err)
		return dto.UserRegisterRes{}, err
	}
	//kirim ke email
	//err = u.EmailService.SendOTPByEmail(req.Email, "Kode OTP Anda", "Code Otp anda adalah "+otpCode)
	//if err != nil {
	//	return dto.UserRegisterRes{}, err
	//}

	return dto.UserRegisterRes{
		ReferenceID: referenceID,
	}, nil
}

func (u UserService) ValidateOTP(c context.Context, req dto.ValidateOTPReq) error {
	val, err := u.CacheRepository.Get("otp:" + req.ReferenceID)
	if err != nil {
		logrus.Error(err)
		return domain.OtpInvalid
	}
	otp := string(val)
	if otp != req.OTP {
		return domain.OtpInvalid
	}

	val, err = u.CacheRepository.Get("user-ref:" + req.ReferenceID)
	if err != nil {
		return domain.OtpInvalid
	}

	criteria := map[string]interface{}{"username": string(val), "email": string(val)}
	user, err := u.UserRepository.FindUser(c, criteria, false)

	//user, err := u.UserRepository.FindByUsernameOrEmail(c, string(val), string(val))
	if err != nil {
		logrus.Error(err)
		return err
	}

	user.EmailVerifiedAt = time.Now()
	logrus.Info("EmailVerifiedAt before update:", user.EmailVerifiedAt)
	_ = u.UserRepository.Update(c, &user)

	return nil
}
