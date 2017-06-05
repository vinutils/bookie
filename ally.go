package main

import (
    "os"
    "bufio"
    "encoding/csv"
    "io"
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "fmt"
)

type AllyRecord struct {
    filePath string
}

func (rec AllyRecord) dookie() {
    fmt.Println(rec.filePath)
}

func (rec AllyRecord) getFilePath() string {
    return rec.filePath
}

func (ally AllyRecord) importStatement(stmtPath string) {
    file, err := os.Open(stmtPath)
    check(err)
    defer file.Close()

    db, err := sql.Open("mysql", "bookie:bookie@/bookie")
    check(err)
    defer db.Close()

    dbWriter, err := db.Prepare("INSERT INTO ally VALUES (?, ?, ?, ?, ?, ?)")
    check(err)
    defer dbWriter.Close()

    csvReader := csv.NewReader(bufio.NewReader(file))
    csvReader.Read()

    for {
        record, err := csvReader.Read()
        if (err == io.EOF) {
            break;
        }

        amt, err := strconv.ParseFloat(record[2], 64)
        _, err = dbWriter.Exec(getHashId(record), record[0],record[1], amt,record[3],record[4])
        check(err)
    }
}