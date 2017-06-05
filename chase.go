package main

import (
    "fmt"
    "database/sql"
    "encoding/csv"
    "bufio"
    "io"
    "os"
    "strings"
    "strconv"
)

func chaseAccountType (s string) string {
    if(strings.HasSuffix(s, "/chk")) {
        return "CHECKING";
    }
    if(strings.HasSuffix(s, "/svg")) {
        return "SAVINGS";
    }
    if(strings.HasSuffix(s, "/reserve")) {
        return "RESERVE";
    }
    if(strings.HasSuffix(s, "/sapphire")) {
        return "SAPPHIRE";
    }
    if(strings.HasSuffix(s, "/united")) {
        return "UNITED";
    }

    return "ERROR"

}

type ChaseBankRecord struct {
    filePath string
}


func (record ChaseBankRecord) dookie() {
    fmt.Println(record.filePath)
}

func (rec ChaseBankRecord) importStatement(stmtPath string) {
    file, err := os.Open(stmtPath)
    check(err)
    defer file.Close()

    db, err := sql.Open("mysql", "bookie:bookie@/bookie")
    check(err)
    defer db.Close()

    dbWriter, err := db.Prepare("INSERT INTO chase_banking VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
    check(err)
    defer dbWriter.Close()

    csvReader := csv.NewReader(bufio.NewReader(file))
    csvReader.Read()

    for {
        record, err := csvReader.Read()
        if (err == io.EOF) {
            break;
        }
        if (len(record) != 8) {
            continue;
        }

        amt, err := strconv.ParseFloat(record[3], 64)
        balance, err := strconv.ParseFloat(record[5], 64)

        _, err = dbWriter.Exec(
            getHashId(record),
            record[0], //details
            fixUSStyleDate(record[1]), //posting date
            record[2], //desc
            amt,
            record[4], //type
            balance,
            record[6], //check#
            chaseAccountType(rec.getFilePath()))

        check(err)
    }
}


func (rec ChaseBankRecord) getFilePath() string {
    return rec.filePath
}


/***************************************************************************************/
/*********************              Credit Card          *******************************/
/***************************************************************************************/


type ChaseCCRecord struct {
    filePath string
}

func (rec ChaseCCRecord) dookie() {
    fmt.Println(rec.filePath)
}

func (rec ChaseCCRecord) importStatement(stmtPath string) {
    file, err := os.Open(stmtPath)
    check(err)
    defer file.Close()

    db, err := sql.Open("mysql", "bookie:bookie@/bookie")
    check(err)
    defer db.Close()

    dbWriter, err := db.Prepare("INSERT INTO chase_cc VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
    check(err)
    defer dbWriter.Close()

    csvReader := csv.NewReader(bufio.NewReader(file))
    csvReader.Read()

    for {
        record, err := csvReader.Read()
        if (err == io.EOF) {
            break;
        }
        if (len(record) != 7) {
            continue;
        }

        amt, err := strconv.ParseFloat(record[4], 64)

        _, err = dbWriter.Exec(
            getHashId(record),
            record[0], //type
            fixUSStyleDate(record[1]), //tx date
            fixUSStyleDate(record[2]), //post date
            record[3], //desc
            amt,
            record[5], //category
            record[6], //memo
            chaseAccountType(rec.getFilePath()))

        check(err)
    }
}


func (rec ChaseCCRecord) getFilePath() string {
    return rec.filePath
}