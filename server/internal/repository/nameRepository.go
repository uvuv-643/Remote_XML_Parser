package repository

import (
	"Remote_XML_Parser/internal/dto"
	"Remote_XML_Parser/internal/services"
	"gorm.io/gorm"
	"strings"
)

type NameRepository struct {
}

func NewNameRepository() *NameRepository {
	return &NameRepository{}
}

func (r *NameRepository) Weak(db *gorm.DB, nameComponents ...string) ([]dto.UserName, error) {
	query := `
		SELECT *
FROM (SELECT CONCAT('aka', sdn_akas.uid) as "key", sdn_items.uid as "uid", sdn_akas.first_name as "first_name", sdn_akas.last_name as "last_name"
      FROM sdn_items
               INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
               INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
      UNION
      SELECT CONCAT('item', sdn_items.uid) as "key", sdn_items.uid as "uid", sdn_items.first_name as "first_name", sdn_items.last_name as "last_name"
      FROM sdn_items) as users
WHERE users."uid" = (SELECT best_match.item_id
                     FROM (SELECT users."uid" as item_id, COUNT(*) as cnt
                           FROM (SELECT sdn_items.uid as "uid", null as "first_name", sdn_akas.last_name as "last_name"
                                 FROM sdn_items
                                          INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
                                          INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
                                 WHERE LOWER(sdn_akas.last_name) LIKE ?
                                    OR LOWER(sdn_akas.last_name) LIKE ?
                                    OR LOWER(sdn_akas.first_name) LIKE ?
                                    OR LOWER(sdn_akas.first_name) LIKE ?
                                 UNION
                                 SELECT sdn_items.uid        as "uid",
                                        sdn_items.first_name as "first_name",
                                        sdn_items.last_name  as "last_name"
                                 FROM sdn_items
                                 WHERE LOWER(sdn_items.last_name) LIKE ?
                                    OR LOWER(sdn_items.first_name) LIKE ?
                                    OR LOWER(sdn_items.first_name) LIKE ?
                                    OR LOWER(sdn_items.last_name) LIKE ?) as users
                           GROUP BY "uid"
                           ORDER BY cnt DESC
                           LIMIT 1) AS best_match)
	`
	params := make([]interface{}, 8)
	if len(nameComponents) == 1 {
		for i := range params {
			params[i] = "%" + strings.ToLower(nameComponents[0]) + "%"
		}
	} else if len(nameComponents) >= 2 {
		for i := range params {
			params[i] = "%" + strings.ToLower(nameComponents[i%2]) + "%"
		}
	} else {
		return nil, services.UserBadRequest
	}

	queryResult := db.Raw(query, params...)
	rows, err := queryResult.Rows()
	if err != nil {
		return nil, services.ServerUnavailable
	}
	var rowSlice []dto.UserName
	for rows.Next() {
		var s dto.UserName
		err := rows.Scan(&s.RecordKey, &s.UID, &s.FirstName, &s.LastName)
		if err != nil {
			return nil, services.ServerUnavailable
		}
		rowSlice = append(rowSlice, s)
	}
	return rowSlice, nil

}

func (r *NameRepository) Strong(db *gorm.DB, nameComponents ...string) ([]dto.UserName, error) {
	query := `
SELECT *
FROM (SELECT CONCAT('aka', sdn_akas.uid) as "key", sdn_items.uid as "uid", sdn_akas.first_name as "first_name", sdn_akas.last_name as "last_name"
      FROM sdn_items
               INNER JOIN items_akas ON sdn_items.uid = items_akas.sdn_item_uid
               INNER JOIN sdn_akas ON sdn_akas.uid = items_akas.sdn_aka_uid
      UNION
      SELECT CONCAT('item', sdn_items.uid) as "key", sdn_items.uid as "uid", sdn_items.first_name as "first_name", sdn_items.last_name as "last_name"
      FROM sdn_items) as users
WHERE LOWER(first_name) LIKE ? AND LOWER(last_name) LIKE ?
   OR LOWER(last_name) LIKE ? AND LOWER(first_name) LIKE ?
	`
	params := make([]interface{}, 4)
	if len(nameComponents) == 1 {
		for i := range params {
			if i%2 == 0 {
				params[i] = strings.ToLower(nameComponents[0])
			} else {
				params[i] = "%"
			}
		}
	} else if len(nameComponents) >= 2 {
		for i := range params {
			params[i] = strings.ToLower(nameComponents[i%2])
		}
	} else {
		return nil, services.UserBadRequest
	}

	queryResult := db.Raw(query, params...)
	rows, err := queryResult.Rows()
	if err != nil {
		return nil, services.ServerUnavailable
	}
	var rowSlice []dto.UserName
	for rows.Next() {
		var s dto.UserName
		err := rows.Scan(&s.RecordKey, &s.UID, &s.FirstName, &s.LastName)
		if err != nil {
			return nil, services.ServerUnavailable
		}
		rowSlice = append(rowSlice, s)
	}
	return rowSlice, nil

}
