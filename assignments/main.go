package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world")

	arraySign([]int{2, 1})                    // 1
	arraySign([]int{-2, 1})                   // -1
	arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	isAnagram("anak", "kana") // true
	isAnagram("anak", "mana") // false
	isAnagram("anagram", "managra") // true

	findTheDifference("abcd", "abcde") // 'e'
	findTheDifference("abcd", "abced") // 'e'
	findTheDifference("", "y")         // 'y'

	canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	sign := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			sign *= -1
		}
	}

	return sign // if positive
	// return -1 // if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	for _, char := range t {
		if count[char] < 0 {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	var result byte

    for i := range s {
        result ^= s[i]
    }

    for i := range t {
        result ^= t[i]
    }

    return result
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	
		diff := arr[1] - arr[0]
	
		for i := 2; i < len(arr); i++ {
			if arr[i] - arr[i-1] != diff {
				return false
			}
		}
	
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
		symbol := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
		number := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

		d.cards = []Card{}

		for _, symbols := range symbol {
			for _, numbers := range number {
				d.cards = append(d.cards, Card{symbol: symbol, number: number})
			}
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}

	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}

	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	if idx < 0 || idx >= len(d.cards) {
		return Card{} 
		
	}
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	// Implementasi Fisher-Yates Shuffle
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	if n < 0 {
		n = 0
	}

	// Pisahkan kartu berdasarkan jumlah n
	cut := d.cards[:n]
	rest := d.cards[n:]

	// Gabungkan kartu, bagian bawah diikuti bagian atas
	d.cards = append(rest, cut...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}

func main(){
	tesDeck()
}
