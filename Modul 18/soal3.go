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

type Player struct {
	Name string
	Hand []Domino
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

func nilaiKartu(card Domino) int {
	return card.Left + card.Right
}

func printDomino(card Domino) string {
	return fmt.Sprintf("[%d|%d]", card.Left, card.Right)
}

func dealHands(deck *Dominoes, players []Player, handSize int) {
	for i := 0; i < handSize; i++ {
		for idx := range players {
			players[idx].Hand = append(players[idx].Hand, ambilKartu(deck))
		}
	}
}

func matches(card Domino, value int) bool {
	return card.Left == value || card.Right == value
}

func findPlayableCard(hand []Domino, leftValue, rightValue int) int {
	for i, card := range hand {
		if matches(card, leftValue) || matches(card, rightValue) {
			return i
		}
	}
	return -1
}

func placeCardOnTable(card Domino, leftValue, rightValue int) (Domino, int, int, bool) {
	if card.Right == leftValue {
		return card, card.Left, rightValue, true
	}
	if card.Left == leftValue {
		return Domino{Left: card.Right, Right: card.Left, IsDouble: card.IsDouble}, card.Right, rightValue, true
	}
	if card.Left == rightValue {
		return card, leftValue, card.Right, true
	}
	if card.Right == rightValue {
		return Domino{Left: card.Right, Right: card.Left, IsDouble: card.IsDouble}, leftValue, card.Left, true
	}
	return card, leftValue, rightValue, false
}

func handString(hand []Domino) string {
	str := ""
	for i, card := range hand {
		if i > 0 {
			str += " "
		}
		str += printDomino(card)
	}
	return str
}

func playGaple(players []Player, deck *Dominoes) {
	table := []Domino{}
	firstCard := ambilKartu(deck)
	table = append(table, firstCard)
	leftValue, rightValue := firstCard.Left, firstCard.Right

	fmt.Printf("Kartu pertama di meja: %s\n", printDomino(firstCard))

	passes := 0
	current := 0

	for passes < len(players) {
		player := &players[current]
		fmt.Printf("%s giliran. Tangan: %s\n", player.Name, handString(player.Hand))

		idx := findPlayableCard(player.Hand, leftValue, rightValue)
		if idx == -1 {
			if deck.Count > 0 {
				drawn := ambilKartu(deck)
				player.Hand = append(player.Hand, drawn)
				fmt.Printf("%s tidak bisa main, mengambil %s\n", player.Name, printDomino(drawn))
				idx = findPlayableCard(player.Hand, leftValue, rightValue)
			}
		}

		if idx >= 0 {
			card := player.Hand[idx]
			card, leftValue, rightValue, placed := placeCardOnTable(card, leftValue, rightValue)
			if placed {
				fmt.Printf("%s memasang %s. Meja sekarang: %d-%d ... %d-%d\n", player.Name, printDomino(card), leftValue, table[0].Left, table[len(table)-1].Right, rightValue)
				table = append(table, card)
				player.Hand = append(player.Hand[:idx], player.Hand[idx+1:]...)
				passes = 0

				if len(player.Hand) == 0 {
					fmt.Printf("%s menang!\n", player.Name)
					return
				}
			} else {
				passes++
				fmt.Printf("%s tidak bisa memasang kartu dari tangan.\n", player.Name)
			}
		} else {
			passes++
			fmt.Printf("%s melewatkan giliran.\n", player.Name)
		}

		current = (current + 1) % len(players)
	}

	fmt.Println("Permainan berakhir karena semua pemain melewatkan giliran.")
	for _, player := range players {
		total := 0
		for _, card := range player.Hand {
			total += nilaiKartu(card)
		}
		fmt.Printf("%s sisa nilai tangan: %d\n", player.Name, total)
	}
}

func main() {
	deck := newDominoes()
	kocokKartu(&deck)

	players := []Player{
		{Name: "Pemain 1"},
		{Name: "Pemain 2"},
	}

	dealHands(&deck, players, 7)
	fmt.Println("Kartu dibagikan kepada pemain.")

	playGaple(players, &deck)
}
