package service

import (
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ipCheckerService struct {
}

func NewIpChecker() domain.IpCheckerService {
	return ipCheckerService{}
}

func (i ipCheckerService) Query(ctx context.Context, ip string) (checker dto.IpChecker, err error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,lat,lon,timezone,query", ip)
	// hide apinya
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("error di bagian ipchecker service query", err)
		return dto.IpChecker{}, err
	}
	// nanti ketika sudah beroperasi selesai, func nya kita akan jalankan closenya
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("io.readall error ", err)
		return dto.IpChecker{}, err
	}

	err = json.Unmarshal(body, &checker)
	if err != nil {
		logrus.Error("convert body nya ke checker error ", err)
		return dto.IpChecker{}, err
	}
	return
}
