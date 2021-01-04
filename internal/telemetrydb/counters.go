package telemetrydb

import (
  "fmt"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/log"
  "github.com/newrelic-experimental/newrelic-TSAK/internal/stdlib"
)

func TDBCounterAdd(key string) bool {
  if TDB != nil {
    tx, err := TDB.Begin()
    if err != nil {
      log.Error(fmt.Sprintf("Error building counter/add transaction:", err))
      return false
    }
    stmt, err := tx.Prepare("insert into Counters(timestamp,key,value) values(?,?,1) on conflict(key) do update set timestamp=?, value=value+1")
    if err != nil {
      log.Error(fmt.Sprintf("Error building counter/add query:", err))
    } else {
      _, err = stmt.Exec(stdlib.NowMilliseconds(), key, stdlib.NowMilliseconds())
      if err != nil {
        stmt.Close()
        log.Error(fmt.Sprintf("Error executing counter/add query:", err))
        return false
      } else {
        tx.Commit()
        stmt.Close()
      }
      return true
    }
  }
  return false
}

func TDBCounterGet(key string) (res int64, err error) {
  dbres, err := TDBValue(0, key)
  res = dbres.(int64)
  return
}
