package mapper

import (
	"git.inet.co.th/ekyc-platform-backend/model"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	one_id "git.inet.co.th/ekyc-platform-backend/pkg/one-id"
)

func MapRequestResgisterToRequestRegisterOneId(src *dto.RequestRegisterUser) *one_id.RequestApiRegisterOneId {
	return &one_id.RequestApiRegisterOneId{
		TitleTh:         src.TitleTh,
		SpecialTitleTh:  src.SpecialTitleTh,
		MiddleNameTh:    src.MiddleNameTh,
		FirstNameTh:     src.FirstNameTh,
		LastNameTh:      src.LastNameTh,
		TitleEng:        src.TitleEng,
		SpecialTitleEng: src.SpecialTitleEng,
		MiddleNameEng:   src.MiddleNameEng,
		FirstNameEng:    src.FirstNameEng,
		LastNameEng:     src.LastNameEng,
		Email:           src.Email,
		MobileNo:        src.MobileNo,
		BirthDate:       src.BirthDate,
		Username:        src.Username,
		Password:        src.Password,
		IdCardNum:       src.IdCardNum,
	}
}

func MapRequestResgisterToModelAccount(src *dto.RequestRegisterUser, accountIdOne string) *model.Account {
	return &model.Account{
		AccountOneId:    accountIdOne,
		TitleTh:         src.TitleTh,
		SpecialTitleTh:  src.SpecialTitleTh,
		MiddleNameTh:    src.MiddleNameTh,
		LastNameTh:      src.LastNameTh,
		FirstNameEng:    src.FirstNameEng,
		MiddleNameEng:   src.MiddleNameEng,
		LastNameEng:     src.LastNameEng,
		TitleEng:        src.TitleEng,
		SpecialTitleEng: src.SpecialTitleEng,
		CidType:         "ID_CARD",
		CidHash:         src.IdCardNum, //ใส่ไปก่อนเดี๋ยวค่อยhash
		CidEncrypt:      src.IdCardNum,
		Email:           src.Email,
		BirthDate:       src.BirthDate,
	}
}
