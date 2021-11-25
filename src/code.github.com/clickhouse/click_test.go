package clickhouse

import (
	"testing"
	"time"
)

type jdbcExample struct {
	Day  time.Time `db:"day"`
	Name string    `db:"name"`
	Age  uint8     `db:age`
}

func TestA(t *testing.T) {

}
