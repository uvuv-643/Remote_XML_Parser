package models

import (
	"Remote_XML_Parser/internal/models/dbs"
	"Remote_XML_Parser/internal/models/xmls"
	"reflect"
)

const XMLS_PREFIX = "xmls"

func ConvertXmlToDb(value interface{}) interface{} {
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNProgram" {
		program := value.(xmls.SDNProgram)
		return &dbs.SDNProgram{Program: string(program)}
	}
	if reflect.TypeOf(value).String() == XMLS_PREFIX+".SDNAka" {
		aka := value.(xmls.SDNAka)
		return &dbs.SDNAka{
			UID:      aka.UID,
			Type:     aka.Type,
			Category: aka.Category,
			LastName: aka.LastName,
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
	return nil
}
