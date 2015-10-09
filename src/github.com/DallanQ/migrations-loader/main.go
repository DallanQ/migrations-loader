package main
import (
    "flag"
    "os"
    "fmt"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "bufio"
    "regexp"
    "strings"
    "strconv"
)

// protocol is typically tcp
var conn = flag.String("c", "", "Database connection string: [user[:pass]@][protocol[(address:port)]]/dbname (required)")
var immFilename = flag.String("i", "", "immigrations filename")
var emFilename = flag.String("e", "", "emigrations filename ")

func main() {
    flag.Parse()
   	if *conn == "" || (*immFilename == "" && *emFilename == "") {
   		flag.Usage()
   		os.Exit(1)
   	}

    db, err := sql.Open("mysql", *conn)
    if err != nil {
        log.Fatal("Open", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("Ping", err)
    }

    if *immFilename != "" {
        parseFile(*immFilename, "To: \\{(.*) (\\d+)\\} From: (.*)", db, "immigrations")
    }

    if *emFilename != "" {
        parseFile(*emFilename, "From: \\{(.*) (\\d+)\\} To: (.*)", db, "emigrations")
    }

}

func parseFile(filename string, pattern string, db *sql.DB, table string) {
    re := regexp.MustCompile(pattern)
    prepare := "insert into "+table+" values(?,?,?,?)"

    file, err := os.Open(*immFilename)
    if err != nil {
        log.Fatal("Open",err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        data := re.FindStringSubmatch(line)
        if data == nil || len(data) != 4 {
            log.Println("FindStringSubmatch", line)
            continue
        }
        place := strings.TrimSpace(data[1])
        if place == "" {
            continue
        }
        year, err := strconv.Atoi(data[2])
        if err != nil {
            log.Println("Atoi year", data[2])
            continue
        }

        fmt.Printf("loadPlaces p=%s y=%d counts=%s\n", place, year, data[3])
        loadPlaces(place, year, data[3], prepare, db)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func loadPlaces(place string, year int, placeCounts string, prepare string, db *sql.DB) {
    re := regexp.MustCompile("(.*) \\((\\d+)\\)")

    // start a tx
    tx, err := db.Begin()
    if err != nil {
        log.Println("Begin", err)
        return
    }
    defer tx.Rollback()

    // prepare statement
    stmt, err := tx.Prepare(prepare)
    if err != nil {
        log.Println("Prepare", err)
        return
    }
    defer stmt.Close()

    for _, placeCount := range strings.Split(strings.TrimSpace(placeCounts),";") {
        if placeCount == "" {
            continue
        }
        data := re.FindStringSubmatch(placeCount)
        if data == nil || len(data) != 3 {
            log.Println("FindStringSubmatch", placeCount)
            continue
        }
        place2 := strings.TrimSpace(data[1])
        if place2 == "" {
            continue
        }
        count, err := strconv.Atoi(data[2])
        if err != nil {
            log.Println("Atoi count", data[2])
            continue
        }

        _, err = stmt.Exec(place, year, place2, count)
        if err != nil {
            log.Println("Exec", err)
        }
    }

    err = tx.Commit()
    if err != nil {
        log.Println("Commit", err)
    }
}
