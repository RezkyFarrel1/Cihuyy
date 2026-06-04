package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Domino struct {
    Left     int
    Right    int
    IsDouble bool
}

type Dominoes struct {
    Cards [28]Domino
    Count int
}

func newDominoes() Dominoes {
    var deck Dominoes
    idx := 0
    for left := 0; left <= 6; left++ {
        for right := left; right <= 6; right++ {
            deck.Cards[idx] = Domino{Left: left, Right: right, IsDouble: left == right}
            idx++
        }
    }
    deck.Count = idx
    return deck
}

func kocokKartu(deck *Dominoes) {
    rand.Seed(time.Now().UnixNano())
    for i := deck.Count - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
    }
}

func ambilKartu(deck *Dominoes) Domino {
    if deck.Count == 0 {
        return Domino{Left: -1, Right: -1}
    }
    deck.Count--
    return deck.Cards[deck.Count]
}

func gambarKartu(card Domino, suit int) int {
    if suit == 1 {
        return card.Left
    }
    return card.Right
}

func nilaiKartu(card Domino) int {
    return card.Left + card.Right
}

func printDomino(card Domino) string {
    return fmt.Sprintf("[%d|%d]", card.Left, card.Right)
}

func samaGambar(card Domino, value int) bool {
    return card.Left == value || card.Right == value
}

func galiKartu(deck *Dominoes, target Domino) []Domino {
    var drawn []Domino
    if deck.Count == 0 {
        return drawn
    }

    targetValues := []int{target.Left, target.Right}
    for deck.Count > 0 {
        card := ambilKartu(deck)
        drawn = append(drawn, card)

        for _, val := range targetValues {
            if samaGambar(card, val) {
                return drawn
            }
        }
    }
    return drawn
}

func sepasangKartu(a, b Domino) bool {
    return nilaiKartu(a)+nilaiKartu(b) == 12
}

func main() {
    deck := newDominoes()
    kocokKartu(&deck)

    target := Domino{Left: 4, Right: 2}
    fmt.Printf("Target kartu: %s\n", printDomino(target))

    drawn := galiKartu(&deck, target)
    fmt.Printf("Mengambil %d kartu sampai cocok:\n", len(drawn))
    for i, card := range drawn {
        fmt.Printf("  %d. %s (nilai=%d)\n", i+1, printDomino(card), nilaiKartu(card))
    }

    a := Domino{Left: 6, Right: 0}
    b := Domino{Left: 3, Right: 3}
    fmt.Printf("Sepasang %s dan %s -> %t\n", printDomino(a), printDomino(b), sepasangKartu(a, b))
}
