package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Book struct {
    ID       string
    Title    string
    Author   string
    Publisher string
    Copies   int
    Year     int
    Rating   int
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    if err != nil {
        return
    }
    line = strings.TrimSpace(line)
    if line == "" {
        return
    }

    n, err := strconv.Atoi(line)
    if err != nil || n <= 0 {
        return
    }

    books := make([]Book, 0, n)
    for i := 0; i < n; i++ {
        raw, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        raw = strings.TrimSpace(raw)
        if raw == "" {
            i--
            continue
        }

        fields := strings.Fields(raw)
        if len(fields) < 7 {
            continue
        }

        copies, err1 := strconv.Atoi(fields[len(fields)-3])
        year, err2 := strconv.Atoi(fields[len(fields)-2])
        rating, err3 := strconv.Atoi(fields[len(fields)-1])
        if err1 != nil || err2 != nil || err3 != nil {
            continue
        }

        books = append(books, Book{
            ID:        fields[0],
            Title:     fields[1],
            Author:    fields[2],
            Publisher: fields[3],
            Copies:    copies,
            Year:      year,
            Rating:    rating,
        })
    }

    searchLine, err := reader.ReadString('\n')
    if err != nil {
        return
    }
    searchLine = strings.TrimSpace(searchLine)
    searchRating, err := strconv.Atoi(searchLine)
    if err != nil {
        return
    }

    if len(books) == 0 {
        return
    }

    favorite := findFavorite(books)
    fmt.Printf("%s %s %s %s %d %d %d\n", favorite.ID, favorite.Title, favorite.Author, favorite.Publisher, favorite.Copies, favorite.Year, favorite.Rating)

    topTitles := topTitlesByRating(books, 5)
    fmt.Println(strings.Join(topTitles, " "))

    for _, b := range books {
        if b.Rating == searchRating {
            fmt.Printf("%s %s %s %s %d %d %d\n", b.ID, b.Title, b.Author, b.Publisher, b.Copies, b.Year, b.Rating)
        }
    }
}

func findFavorite(books []Book) Book {
    favorite := books[0]
    for _, b := range books {
        if b.Rating > favorite.Rating {
            favorite = b
        }
    }
    return favorite
}

func topTitlesByRating(books []Book, count int) []string {
    sorted := make([]Book, len(books))
    copy(sorted, books)
    sort.SliceStable(sorted, func(i, j int) bool {
        return sorted[i].Rating > sorted[j].Rating
    })

    if count > len(sorted) {
        count = len(sorted)
    }

    titles := make([]string, 0, count)
    for i := 0; i < count; i++ {
        titles = append(titles, sorted[i].Title)
    }
    return titles
}
