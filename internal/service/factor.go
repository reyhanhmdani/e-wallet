package service

import (
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type factorService struct {
	factorRepository domain.FactorRepository
}

func NewFactor(factorRepository domain.FactorRepository) domain.FactorService {
	return &factorService{
		factorRepository: factorRepository,
	}
}

func (f factorService) ValidatePin(ctx context.Context, req dto.ValidatePinReq) error {
	factor, err := f.factorRepository.FindByUser(ctx, req.UserID)
	if err != nil {
		logrus.Error("Find By user di service factor gagal", err)
		return domain.UserNotFound
	}

	if factor == (domain.Factor{}) {
		logrus.Error("Find By user di service kosong", err)
		return domain.PinInvalid
	}

	err = bcrypt.CompareHashAndPassword([]byte(factor.Pin), []byte(req.PIN))
	if err != nil {
		logrus.Error("Compare password gagal", err)
		return err
	}

	return nil
}
