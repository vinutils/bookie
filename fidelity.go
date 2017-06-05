package main

import (
    "fmt"
    "os"
    "database/sql"
    "encoding/csv"
    "bufio"
    "io"
    "strconv"
)

type FidelityRecord struct {
    filePath string
}

func (stmt FidelityRecord) dookie() {
    fmt.Println(stmt.filePath)
}

func (stmt FidelityRecord) importStatement(stmtPath string) {
    file, err := os.Open(stmtPath)
    check(err)
    defer file.Close()

    db, err := sql.Open("mysql", "bookie:bookie@/bookie")
    check(err)
    defer db.Close()

    dbWriter, err := db.Prepare("INSERT INTO fidelity VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    check(err)
    defer dbWriter.Close()

    csvReader := csv.NewReader(bufio.NewReader(file))
    csvReader.Read()

    for {
        record, err := csvReader.Read()
        if (err == io.EOF) {
            break;
        }
        if (len(record) != 17) {
            continue;
        }

        qty, err := strconv.ParseFloat(record[8], 64)
        price, err := strconv.ParseFloat(record[10], 64)
        xRate, err := strconv.ParseFloat(record[11], 64)
        comm, err := strconv.ParseFloat(record[12], 64)
        fees, err := strconv.ParseFloat(record[13], 64)
        intr, err := strconv.ParseFloat(record[14], 64)
        amt, err := strconv.ParseFloat(record[15], 64)


        _, err = dbWriter.Exec(
            getHashId(record),
            fixUSStyleDate(record[0]), //rundate
            record[1], //account
            record[2], //action
            record[3], //symbol
            record[4], //desc
            record[5], //type
            record[6], //xqty
            record[7], //xcurr
            qty,
            record[9], //curr
            price,
            xRate,
            comm,
            fees,
            intr,
            amt,
            fixUSStyleDate(record[16])) //settleDate


        check(err)
    }
}

func (rec FidelityRecord) getFilePath() string {
    return rec.filePath
}