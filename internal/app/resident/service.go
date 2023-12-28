package resident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	residentRepository repository.Resident
}

type Service interface {
	GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	Store(ctx context.Context, payload dto.PayloadStoreResident) error
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

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreResident) error {
	dataStore := model.Resident{
		Nama:          payload.Nama,
		Alamat:        payload.Alamat,
		JenisKelamin:  payload.JenisKelamin,
		NamaKabupaten: payload.NamaKabupaten,
		NamaKecamatan: payload.NamaKecamatan,
		NamaKelurahan: payload.NamaKelurahan,
		Nik:           payload.Nik,
		Rt:            payload.Rt,
		Rw:            payload.Rw,
		Usia:          payload.Usia,
		Telp:          payload.Telp,
		Tps:           payload.Tps,
	}
	if err := s.residentRepository.Store(ctx, dataStore); err != nil {
		return err
	}
	return nil
}
