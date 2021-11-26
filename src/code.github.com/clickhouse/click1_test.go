package clickhouse

import (
	"testing"
)

func TestSchema(t *testing.T)  {

	sql := "desc test_ccm_dws_rpt_scada_farm_1d"

	rows, _ := DB.Exec(sql).Rows()
	for rows.Next() {

	}
}
