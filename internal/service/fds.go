package service

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"log"
	"time"
)

type fdsService struct {
	ipCheckerService   domain.IpCheckerService
	loginLogRepository domain.LoginLogRepository
}

func NewFds(ipCheckerService domain.IpCheckerService,
	loginLogRepository domain.LoginLogRepository) domain.FdsService {
	return &fdsService{
		ipCheckerService:   ipCheckerService,
		loginLogRepository: loginLogRepository,
	}

}

func (f fdsService) IsAuthorized(ctx context.Context, ip string, userId uuid.UUID) bool {
	locationCheck, err := f.ipCheckerService.Query(ctx, ip)
	if err != nil || locationCheck == (dto.IpChecker{}) {
		logrus.Error("Query gagal", err)
		return false
	}

	// buat domain
	newAccess := domain.LoginLog{
		UserId:       userId,
		IsAuthorized: false,
		IpAddress:    ip,
		Timezone:     locationCheck.Timezone,
		Lat:          locationCheck.Lat,
		Lon:          locationCheck.Lon,
		AccessTime:   time.Now(),
	}

	lastLogin, err := f.loginLogRepository.FindLastAuthorized(ctx, userId)
	if err != nil {
		logrus.Error("Find Last authorized gagal", err)
		_ = f.loginLogRepository.Save(ctx, &newAccess)
		return false
	}
	// kalau belum pernah login
	if lastLogin == (domain.LoginLog{}) {
		newAccess.IsAuthorized = true
		_ = f.loginLogRepository.Save(ctx, &newAccess)
		return true
	}
	// letak pengecekan jarak
	distanceHour := newAccess.AccessTime.Sub(lastLogin.AccessTime)
	distanceChange := util.GetDistance(lastLogin.Lat, lastLogin.Lon, newAccess.Lat, newAccess.Lon)

	log.Printf("hour: %f, distance: %f\n", distanceHour.Hours(), distanceChange)

	if (distanceChange / distanceHour.Hours()) > 400 {
		_ = f.loginLogRepository.Save(ctx, &newAccess)
		return false
	}

	newAccess.IsAuthorized = true
	_ = f.loginLogRepository.Save(ctx, &newAccess)
	return true
}
