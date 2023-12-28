package resident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	residentRepository repository.Resident
}

type Service interface {
	GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	GetListResidentGroup(ctx context.Context) ([]dto.KecamatanInKabupaten, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		residentRepository: f.ResidentRepository,
	}
}

func (s *service) GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error) {
	resultResident, err := s.residentRepository.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultResident{}, err
	}

	return resultResident, nil
}

func (s *service) GetListResidentGroup(ctx context.Context) ([]dto.KecamatanInKabupaten, error) {
	var listResidentGroup []dto.KecamatanInKabupaten

	residents, err := s.residentRepository.GetKecamatanByKabupaten(ctx, "CIANJUR")
	if err != nil {
		return nil, err
	}

	for _, resident := range residents {
		residentDTO := dto.KecamatanInKabupaten{
			NamaKecamatan: resident.NamaKecamatan,
			NamaKabupaten: resident.NamaKabupaten,
			Count:         resident.Count,
		}
		listResidentGroup = append(listResidentGroup, residentDTO)
	}

	return listResidentGroup, nil
}
