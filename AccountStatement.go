package main

type AccountStatement interface {
    dookie()
    importStatement(filePath string)
    getFilePath() string
}
