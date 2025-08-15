package mapper

import (
	"time"

	"git.inet.co.th/ekyc-platform-backend/model"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	one_id "git.inet.co.th/ekyc-platform-backend/pkg/one-id"
)

func MapToAccount(data map[string]interface{}) (*model.Account, error) {
	account := &model.Account{}

	if v, ok := data["account_id"].(string); ok {
		account.AccountOneId = v
	}
	if v, ok := data["account_title_th"].(string); ok {
		account.TitleTh = v
	}
	if v, ok := data["special_title_name_th"].(string); ok {
		account.SpecialTitleTh = v
	}
	if v, ok := data["first_name_th"].(string); ok {
		account.FirstNameTh = v
	}
	if v, ok := data["middle_name_th"].(string); ok {
		account.MiddleNameTh = v
	}
	if v, ok := data["last_name_th"].(string); ok {
		account.LastNameTh = v
	}
	if v, ok := data["account_title_eng"].(string); ok {
		account.TitleEng = v
	}
	if v, ok := data["special_title_name_eng"].(string); ok {
		account.SpecialTitleEng = v
	}
	if v, ok := data["first_name_eng"].(string); ok {
		account.FirstNameEng = v
	}
	if v, ok := data["middle_name_eng"].(string); ok {
		account.MiddleNameEng = v
	}
	if v, ok := data["last_name_eng"].(string); ok {
		account.LastNameEng = v
	}
	// if v, ok := data["account_category"].(string); ok {
	// 	// Add to model if field exists
	// }
	// if v, ok := data["account_sub_category"].(string); ok {
	// 	// Add to model if field exists
	// }
	if v, ok := data["birth_date"].(string); ok {
		account.BirthDate = v
	}
	if v, ok := data["hash_id_card_num"].(string); ok {
		account.CidHash = v
	}
	// if v, ok := data["id_card_num"].(string); ok {
	// 	account.CidEncrypt = v
	// }
	// if v, ok := data["id_card_type"].(string); ok {
	// 	// Add to model if field exists
	// }
	// if v, ok := data["thai_email"].(string); ok {
	// 	account.Email = v
	// }
	// if v, ok := data["thai_email2"].(string); ok {
	// 	// Add to model if exists
	// }
	// if v, ok := data["thai_email3"].(string); ok {
	// 	// Add to model if exists
	// }
	// if v, ok := data["status_cd"].(string); ok {
	// 	// Add to model if field exists
	// }

	account.LastLogin = time.Now()

	return account, nil
}

func MapResponseApiAccountOneIdToResponseUserProfile(src one_id.ResponseApiAccountOneId) *dto.ResponseUserProfile {
	return &dto.ResponseUserProfile{
		AccountId:           src.ID,
		FirstNameTH:         src.FirstNameTH,
		MiddleNameTH:        src.MiddleNameTH,
		LastNameTH:          src.LastNameTH,
		FirstNameEng:        src.FirstNameEng,
		MiddleNameEng:       src.MiddleNameEng,
		LastNameEng:         src.LastNameEng,
		SpecialTitleNameTH:  src.SpecialTitleNameTH,
		AccountTitleTH:      src.AccountTitleTH,
		SpecialTitleNameEng: src.SpecialTitleNameEng,
		AccountTitleEng:     src.AccountTitleEng,
		IdCardType:          src.IDCardType,
		IdCardNum:           src.IDCardNum,
		HashIdCardNum:       src.HashIDCardNum,
		AccountCategory:     src.AccountCategory,
		AccountSubCategory:  src.AccountSubCategory,
		ThaiEmail:           src.ThaiEmail,
		ThaiEmail2:          src.ThaiEmail2,
		ThaiEmail3:          src.ThaiEmail3,
		StatusCD:            src.StatusCD,
		BirthDate:           src.BirthDate,
	}
}

func MapModelAccountToResponseUserProfile(src *model.Account) *dto.ResponseUserProfile {

	res := dto.ResponseUserProfile{
		AccountId:           src.AccountOneId,
		FirstNameTH:         src.FirstNameTh,
		MiddleNameTH:        src.MiddleNameTh,
		LastNameTH:          src.LastNameTh,
		FirstNameEng:        src.FirstNameEng,
		MiddleNameEng:       src.MiddleNameEng,
		LastNameEng:         src.LastNameEng,
		SpecialTitleNameTH:  src.SpecialTitleTh,
		AccountTitleTH:      src.TitleTh,
		SpecialTitleNameEng: src.SpecialTitleEng,
		AccountTitleEng:     src.TitleEng,
		IdCardType:          src.CidType,
		IdCardNum:           src.CidEncrypt,
		HashIdCardNum:       src.CidHash,
		Email:               src.Email,
		Mobile:              src.Phone,
		BirthDate:           src.BirthDate,
	}

	return &res
}

func MapAccountOneIdToModelAccount(src *one_id.ResponseApiAccountOneId) *model.Account {
	return &model.Account{
		AccountOneId:    src.ID,
		TitleTh:         src.AccountTitleTH,
		SpecialTitleTh:  src.SpecialTitleNameTH,
		MiddleNameTh:    src.MiddleNameTH,
		LastNameTh:      src.LastNameTH,
		FirstNameEng:    src.FirstNameEng,
		MiddleNameEng:   src.MiddleNameEng,
		LastNameEng:     src.LastNameEng,
		TitleEng:        src.AccountTitleEng,
		SpecialTitleEng: src.SpecialTitleNameEng,
		CidType:         src.IDCardType,
		CidHash:         src.HashIDCardNum,
		Email:           src.ThaiEmail,
		BirthDate:       src.BirthDate,
	}
}
