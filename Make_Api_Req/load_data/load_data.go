package load_data

import (
	"Make_Api_Req/get_data"
	"encoding/json"
	"os"
)

type Records struct {
	Index   string          `json:"index"`
	Records []get_data.User `json:"records"`
	Total   int             `json:"total_records"`
}

type Load interface {
	LoadRecord(records []Records) error
	WriteFile(name string, data []byte, perm os.FileMode) error
}
type WriteFile struct{}

func (l WriteFile) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func AppendRecords(c map[string][]get_data.User, l Load) error {
	var records []Records
	for i, j := range c {
		tempR := Records{Index: i, Records: j, Total: len(j)}
		records = append(records, tempR)
	}
	err := l.LoadRecord(records)
	if err != nil {
		return err
	}
	return nil
}
func (w *WriteFile) LoadRecord(records []Records) error {
	for _, rec := range records {
		file, err := json.MarshalIndent(rec, "", "\t")
		if err != nil {
			return err
		}
		name := rec.Index + ".json"
		err = w.WriteFile(name, file, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
