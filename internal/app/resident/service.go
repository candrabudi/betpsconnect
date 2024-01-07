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
	GetValidateResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultValidateResidents, error)
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
		return dto.ResultTpsResidents{
			Items:    []dto.FindTpsResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultTpsResidents.Items) == 0 {
		return dto.ResultTpsResidents{
			Items:    []dto.FindTpsResidents{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultTpsResidents, nil
}

func (s *service) GetValidateResident(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultValidateResidents, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultValidateResidents{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultValidateResidents, err := s.residentRepository.GetListValidate(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultValidateResidents{
			Items:    []dto.FindValidateResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultValidateResidents.Items) == 0 {
		return dto.ResultValidateResidents{
			Items:    []dto.FindValidateResidents{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultValidateResidents, nil
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
		Items:  updatePayload.Items,
		IsTrue: updatePayload.IsTrue,
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
		return []string{}, err
	}
	if len(resultTps) == 0 {
		return []string{}, nil
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
	dataStore := model.TrueResident{
		FullName:    payload.FullName,
		Address:     payload.Address,
		Gender:      payload.Gender,
		City:        payload.City,
		District:    payload.District,
		SubDistrict: payload.Subdistrict,
		Nik:         payload.Nik,
		Age:         payload.Age,
		NoHandphone: payload.NoHandphone,
		Tps:         payload.TPS,
	}
	if err := s.residentRepository.Store(ctx, dataStore); err != nil {
		return err
	}
	return nil
}
