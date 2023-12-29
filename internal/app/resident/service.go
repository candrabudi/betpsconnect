package resident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	residentRepository    repository.Resident
	subdistrictRepository repository.SubDistrict
	districtRepository    repository.District
}

type Service interface {
	GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) (dto.FindTpsByDistrict, error)
	Store(ctx context.Context, payload dto.PayloadStoreResident) error
	GetListResidentGroup(ctx context.Context) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		residentRepository:    f.ResidentRepository,
		subdistrictRepository: f.SubDistrictRepository,
		districtRepository:    f.DistrictRepository,
	}
}

func (s *service) GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error) {
	resultResident, err := s.residentRepository.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultResident{}, err
	}

	return resultResident, nil
}

func (s *service) GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) (dto.FindTpsByDistrict, error) {
	resultTps, err := s.residentRepository.GetTpsBySubDistrict(ctx, filter)
	if err != nil {
		return dto.FindTpsByDistrict{}, err
	}

	return resultTps, nil
}

func (s *service) GetListResidentGroup(ctx context.Context) error {

	results, err := s.residentRepository.GetKecamatanByKabupaten(ctx, "BOGOR")
	if err != nil {
		return err
	}

	for _, data := range results {
		filter := dto.GetByCity{
			NamaKabupaten: data.NamaKabupaten,
			NamaKecamatan: data.NamaKecamatan,
		}
		_, err := s.districtRepository.FindOne(ctx, filter)
		if err == mongo.ErrNoDocuments {
			result, _ := s.districtRepository.FindLastOne(ctx)

			SubDistrictId := result.ID + 1
			dataStore := model.District{
				ID:            SubDistrictId,
				NamaKabupaten: data.NamaKabupaten,
				NamaKecamatan: data.NamaKecamatan,
			}
			if err := s.districtRepository.Store(ctx, dataStore); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
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
