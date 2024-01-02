package models

import (
	"Remote_XML_Parser/internal/models/dbs"
	"Remote_XML_Parser/internal/models/xmls"
	"reflect"
)

const XMLS_PREFIX = "xmls"

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	if s.IsNil() {
		return nil
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func convertToSDNProgramSlice(converted []interface{}) []*dbs.SDNProgram {
	result := make([]*dbs.SDNProgram, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNProgram)
		result[i] = resetStructFields(result[i], "Program").(*dbs.SDNProgram)
	}
	return result
}

func convertToSDNAkaSlice(converted []interface{}) []*dbs.SDNAka {
	result := make([]*dbs.SDNAka, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNAka)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNAka)
	}
	return result
}

func convertToSDNIdSlice(converted []interface{}) []*dbs.SDNId {
	result := make([]*dbs.SDNId, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNId)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNId)
	}
	return result
}

func convertToSDNAddressSlice(converted []interface{}) []*dbs.SDNAddress {
	result := make([]*dbs.SDNAddress, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNAddress)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNAddress)
	}
	return result
}

func convertToSDNNationalitySlice(converted []interface{}) []*dbs.SDNNationality {
	result := make([]*dbs.SDNNationality, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNNationality)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNNationality)
	}
	return result
}

func convertToSDNDateOfBirthSlice(converted []interface{}) []*dbs.SDNDateOfBirth {
	result := make([]*dbs.SDNDateOfBirth, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNDateOfBirth)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNDateOfBirth)
	}
	return result
}

func convertToSDNPlaceOfBirthSlice(converted []interface{}) []*dbs.SDNPlaceOfBirth {
	result := make([]*dbs.SDNPlaceOfBirth, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNPlaceOfBirth)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNPlaceOfBirth)
	}
	return result
}

func convertToSDNCitizenshipSlice(converted []interface{}) []*dbs.SDNCitizenship {
	result := make([]*dbs.SDNCitizenship, len(converted))
	for i, v := range converted {
		result[i] = v.(*dbs.SDNCitizenship)
		result[i] = resetStructFields(result[i], "UID").(*dbs.SDNCitizenship)
	}
	return result
}

func ConvertXmlToDbArray(data []interface{}) []interface{} {
	converted := make([]interface{}, 0)
	for _, entity := range data {
		converted = append(converted, ConvertXmlToDb(entity))
	}
	return converted
}

func resetStructFields(s interface{}, keepField string) interface{} {
	val := reflect.ValueOf(s)
	val = val.Elem()
	typeOfS := val.Type()
	newStruct := reflect.New(typeOfS).Elem()
	var keepFieldIndex int
	for i := 0; i < val.NumField(); i++ {
		fieldName := typeOfS.Field(i).Name
		if fieldName == keepField {
			keepFieldIndex = i
			break
		}
	}
	newStruct.Field(keepFieldIndex).Set(val.Field(keepFieldIndex))
	return newStruct.Addr().Interface()
}

func ConvertXmlToDb(value interface{}) interface{} {
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNProgram" {
		program := value.(xmls.SDNProgram)
		return &dbs.SDNProgram{Program: string(program)}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNAka" {
		aka := value.(xmls.SDNAka)
		return &dbs.SDNAka{
			UID:       aka.UID,
			Type:      aka.Type,
			Category:  aka.Category,
			FirstName: aka.FirstName,
			LastName:  aka.LastName,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNId" {
		id := value.(xmls.SDNId)
		return &dbs.SDNId{
			UID:       id.UID,
			Type:      id.Type,
			Number:    id.Number,
			Country:   id.Country,
			IssueDate: id.IssueDate.Time,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNAddress" {
		address := value.(xmls.SDNAddress)
		return &dbs.SDNAddress{
			UID:             address.UID,
			City:            address.City,
			Address1:        address.Address1,
			Address2:        address.Address1,
			Address3:        address.Address1,
			StateOrProvince: address.Address1,
			PostalCode:      address.Address1,
			Country:         address.Address1,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNNationality" {
		nationality := value.(xmls.SDNNationality)
		return &dbs.SDNNationality{
			UID:       nationality.UID,
			Country:   nationality.Country,
			MainEntry: nationality.MainEntry,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNDateOfBirth" {
		date := value.(xmls.SDNDateOfBirth)
		return &dbs.SDNDateOfBirth{
			UID:         date.UID,
			DateOfBirth: date.DateOfBirth.Time,
			MainEntry:   date.MainEntry,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNPlaceOfBirth" {
		place := value.(xmls.SDNPlaceOfBirth)
		return &dbs.SDNPlaceOfBirth{
			UID:       place.UID,
			Place:     place.Place,
			MainEntry: place.MainEntry,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNCitizenship" {
		citizenship := value.(xmls.SDNCitizenship)
		return &dbs.SDNCitizenship{
			UID:       citizenship.UID,
			Country:   citizenship.Country,
			MainEntry: citizenship.MainEntry,
		}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNItem" {
		item := value.(xmls.SDNItem)
		return &dbs.SDNItem{
			UID:              item.UID,
			FirstName:        item.FirstName,
			LastName:         item.LastName,
			Title:            item.Title,
			SDNType:          item.SDNType,
			Remarks:          item.Remarks,
			Program:          convertToSDNProgramSlice(ConvertXmlToDbArray(InterfaceSlice(item.ProgramList.Program))),
			Aka:              convertToSDNAkaSlice(ConvertXmlToDbArray(InterfaceSlice(item.AkaList.Aka))),
			RID:              convertToSDNIdSlice(ConvertXmlToDbArray(InterfaceSlice(item.IdList.ID))),
			Address:          convertToSDNAddressSlice(ConvertXmlToDbArray(InterfaceSlice(item.AddressList.Address))),
			Nationality:      convertToSDNNationalitySlice(ConvertXmlToDbArray(InterfaceSlice(item.NationalityList.Nationality))),
			DateOfBirthItem:  convertToSDNDateOfBirthSlice(ConvertXmlToDbArray(InterfaceSlice(item.DateOfBirthList.DateOfBirthItem))),
			PlaceOfBirthItem: convertToSDNPlaceOfBirthSlice(ConvertXmlToDbArray(InterfaceSlice(item.PlaceOfBirthList.PlaceOfBirthItem))),
			Citizenship:      convertToSDNCitizenshipSlice(ConvertXmlToDbArray(InterfaceSlice(item.CitizenshipList.Citizenship))),
		}
	}

	return nil
}
