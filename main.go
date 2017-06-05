package main

import (
    "io/ioutil"
    "strings"
    "fmt"
)

func main() {
	//readDir properties file that contains file path
	/*file, err := os.Open("./filepath.txt")
	check(err)
	defer file.Close()

	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		readDir(fscanner.Text())
	}

	if err := fscanner.Err(); err != nil {
		log.Fatal(err)
	}
	*/

    readDirAndProcessStatements(ChaseBankRecord{filePath:"/Users/lolmac/Documents/fins/chs/svg"})
    readDirAndProcessStatements(ChaseBankRecord{filePath:"/Users/lolmac/Documents/fins/chs/chk"})
    readDirAndProcessStatements(ChaseCCRecord{filePath:"/Users/lolmac/Documents/fins/chs/reserve"})
    readDirAndProcessStatements(ChaseCCRecord{filePath:"/Users/lolmac/Documents/fins/chs/sapphire"})
    readDirAndProcessStatements(ChaseCCRecord{filePath:"/Users/lolmac/Documents/fins/chs/united"})
    readDirAndProcessStatements(AllyRecord{filePath:"/Users/lolmac/Documents/fins/aly"})
    readDirAndProcessStatements(FidelityRecord{filePath:"/Users/lolmac/Documents/fins/fid"})
    readDirAndProcessStatements(ETradeRecord{filePath:"/Users/lolmac/Documents/fins/etr"})

}

func readDirAndProcessStatements(stmt AccountStatement) {
    files, _ := ioutil.ReadDir(stmt.getFilePath())
    for _, f := range files {
        if f.IsDir() || !strings.HasSuffix(strings.ToLower(f.Name()), ".csv") {
            continue
        }

        fileName := fmt.Sprint(stmt.getFilePath(), "/",f.Name())
        //processFile(fileName)
        stmt.importStatement(fileName)
        markFileAsProcessed(stmt.getFilePath(), f.Name())
    }
}