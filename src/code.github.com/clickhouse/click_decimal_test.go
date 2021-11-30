package clickhouse

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Decimal(t *testing.T) {
	const (
		ddl = `
			CREATE TABLE clickhouse_test_nullable (
				decimal  Decimal(18,5),
				decimalNullable  Nullable(Decimal(15,3))
			) Engine=Memory;
		`
		dml = `
			INSERT INTO clickhouse_test_nullable (
				decimal,
				decimalNullable
			) VALUES (
				?,
				?
			)
		`
		query = `
			SELECT
				decimal
				
			FROM clickhouse_test_nullable
		`
	)
	if connect, err := sql.Open("clickhouse", "tcp://10.0.50.100:9000?debug=true"); assert.NoError(t, err) {
		if _, err := connect.Begin(); assert.NoError(t, err) {
			if rows, err := connect.Query(query); assert.NoError(t, err) {
				columnTypes, _ := rows.ColumnTypes()
				for i, column := range columnTypes {
					switch i {
					case 0:
						precision, scale, ok := column.DecimalSize()
						println("%v-%v-%v", precision, scale, ok)
					}
				}
				for rows.Next() {
					var decimal int
					rows.Scan(&decimal)
					println("%v-%v-%v", decimal)
				}
			}
		}
	}
}
