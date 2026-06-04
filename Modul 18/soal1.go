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
            deck.Cards[idx] = Domino{
                Left:     left,
                Right:    right,
                IsDouble: left == right,
            }
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
        return Domino{Left: -1, Right: -1, IsDouble: false}
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

func main() {
    deck := newDominoes()
    fmt.Printf("Awal set domino: %d kartu\n", deck.Count)

    kocokKartu(&deck)
    fmt.Println("Deck sudah dikocok")

    fmt.Println("Mengambil 5 kartu pertama:")
    for i := 0; i < 5; i++ {
        card := ambilKartu(&deck)
        fmt.Printf("Kartu %d: %s, nilai=%d, sisi1=%d, sisi2=%d, balak=%t\n",
            i+1, printDomino(card), nilaiKartu(card), gambarKartu(card, 1), gambarKartu(card, 2), card.IsDouble)
    }

    fmt.Printf("Kartu tersisa dalam deck: %d\n", deck.Count)
}
