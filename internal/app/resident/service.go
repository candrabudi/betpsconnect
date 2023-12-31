package resident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	residentRepository    repository.Resident
	subdistrictRepository repository.SubDistrict
	districtRepository    repository.District
}

type Service interface {
	GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultTpsResidents, error)
	DetailResident(ctx context.Context, ResidentID int) (dto.DetailResident, error)
	CheckResidentByNik(ctx context.Context, Nik string) (dto.DetailResident, error)
	GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error)
	Store(ctx context.Context, payload dto.PayloadStoreResident) error
	GetListResidentGroup(ctx context.Context) error
	residentValidate(ctx context.Context, updatePayload dto.PayloadUpdateValidInvalid) ([]int, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		residentRepository:    f.ResidentRepository,
		subdistrictRepository: f.SubDistrictRepository,
		districtRepository:    f.DistrictRepository,
	}
}

func (s *service) GetListResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultTpsResidents, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultTpsResidents{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.residentRepository.GetResidentTps(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultTpsResidents{}, err
	}

	return resultTpsResidents, nil
}

func (s *service) DetailResident(ctx context.Context, ResidentID int) (dto.DetailResident, error) {
	resultResident, err := s.residentRepository.DetailResident(ctx, ResidentID)
	if err != nil {
		return dto.DetailResident{}, err
	}

	return resultResident, nil
}

func (s *service) CheckResidentByNik(ctx context.Context, Nik string) (dto.DetailResident, error) {
	resultResident, err := s.residentRepository.CheckResidentByNik(ctx, Nik)
	if err != nil {
		return dto.DetailResident{}, err
	}

	return resultResident, nil
}

func (s *service) residentValidate(ctx context.Context, updatePayload dto.PayloadUpdateValidInvalid) ([]int, error) {
	if updatePayload.IsTrue != true && updatePayload.IsTrue != false {
		return nil, errors.New("Please insert a valid value (true or false) for IsTrue")
	}

	payload := dto.PayloadUpdateValidInvalid{
		ResidentID: updatePayload.ResidentID,
		IsTrue:     updatePayload.IsTrue,
	}

	results, err := s.residentRepository.ResidentValidate(ctx, payload)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *service) GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error) {
	resultTps, err := s.residentRepository.GetTpsBySubDistrict(ctx, filter)
	if err != nil {
		return nil, err
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
