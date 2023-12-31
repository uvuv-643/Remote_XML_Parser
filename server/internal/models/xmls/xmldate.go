package xmls

import (
	"encoding/xml"
	"time"
)

type XMLDate struct {
	time.Time
}

func (c *XMLDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	forms := make([]string, 0)
	forms = append(forms, "02 Jan 2006")
	forms = append(forms, "Jan 2006")
	forms = append(forms, "2006")
	var v string
	d.DecodeElement(&v, &start)
	for _, form := range forms {
		parse, err := time.Parse(form, v)
		if err == nil {
			*c = XMLDate{parse}
		}
	}
	return nil
}
