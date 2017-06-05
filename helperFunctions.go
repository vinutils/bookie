package main

import (
    "fmt"
    "os"
    "crypto/md5"
    "io"
    "encoding/hex"
    "strings"
    "bytes"
    "database/sql"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func markFileAsProcessed(fileDir, fileName string) {
    //move file to processed
    src := fmt.Sprint(fileDir, "/",fileName)
    dest := fmt.Sprint(fileDir, "/processed/",fileName)
    err := os.Rename(src, dest)
    check(err)
}

func getHashId(record []string) string {
    hash := md5.New()
    for _, field := range record {
        io.WriteString(hash, field)
    }
    return hex.EncodeToString(hash.Sum(nil))
}

func newNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
        String: s,
        Valid: true,
    }
}

func fixUSStyleDate(usDate string) sql.NullString {
    parts := strings.Split(strings.TrimSpace(usDate), "/") //MM/DD/YYYY
    if (len(parts) != 3) {
        return sql.NullString{}
    }
    var str bytes.Buffer
    str.WriteString(parts[2])
    str.WriteString("-")
    str.WriteString(parts[0])
    str.WriteString("-")
    str.WriteString(parts[1])

    return newNullString(str.String()) //YYYY-MM-DD
}