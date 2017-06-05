package main

import "fmt"

type ETradeRecord struct {
    filePath string
}

func (stmt ETradeRecord) dookie() {
    fmt.Println(stmt.filePath)
}
func (stmt ETradeRecord) importStatement(stmtPath string) {
    fmt.Println(stmtPath)
}

func (rec ETradeRecord) getFilePath() string {
    return rec.filePath
}
