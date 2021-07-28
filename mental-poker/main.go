package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

// Do not use anywhere near serious or production-grade systems.
// Commutative scheme idea from https://asecuritysite.com/encryption/go_comm

func main() {
	// Bob and Alice select their passphrace, and get their
	// encoding/decoding functions pairs; Eb/Db and Ea/Da.
	// These will remain secrets until the end of the game, when they will be revealed
	// to verify that no cheating has occured.
	passBob := "pass1"
	passAlice := "pass2"
	Eb, Db := getED(passBob)
	Ea, Da := getED(passAlice)

	// Bob creates a new Deck, and encrypts each card using his key. He then shuffles it.
	deck := buildNewDeck()

	encryptedDeck := make([][]byte, 0)
	for _, card := range deck {
		encryptedDeck = append(encryptedDeck, Eb(card.name()))
	}
	bobRng := rand.New(rand.NewSource(time.Now().Unix()))

	bobRng.Shuffle(
		len(encryptedDeck),
		func(i, j int) {
			encryptedDeck[i], encryptedDeck[j] = encryptedDeck[j], encryptedDeck[i]
		})

	// Alice receives the deck, and selects five cards at random, and sends them back to Bob.
	// Τhis will be Bob's hand, and he can laterdecrypt the values to see what he has been dealt.
	aliceRng := rand.New(rand.NewSource(time.Now().Unix()))

	aliceRng.Shuffle(
		len(encryptedDeck),
		func(i, j int) {
			encryptedDeck[i], encryptedDeck[j] = encryptedDeck[j], encryptedDeck[i]
		})

	bobHand := encryptedDeck[0:5]
	encryptedDeck = encryptedDeck[5:]
	for _, c := range bobHand {
		fmt.Println(string(Db(c)))
	}

	// Now Alice selects five other cards at random, encrypts them with her key.
	// Each card in Alice's hand is doubly-encrypted now, and she sends them over to Bob.
	aliceRng.Shuffle(
		len(encryptedDeck),
		func(i, j int) {
			encryptedDeck[i], encryptedDeck[j] = encryptedDeck[j], encryptedDeck[i]
		})
	aliceHand := encryptedDeck[0:5]
	encryptedDeck = encryptedDeck[5:]
	for i := range aliceHand {
		aliceHand[i] = Ea(aliceHand[i])
	}

	// Bob decrypts Alice's hand with his own decryptor.
	// The hand is now only encrypted with her key, and Bob has no knowledge of it.
	for i := range aliceHand {
		aliceHand[i] = Db(aliceHand[i])
	}

	// Each player now has a hand, only encrypted with his own key,
	// and thus can decode it to see what was dealt to him, and who has the better hand.
	realBobHand := make([][]byte, 0)
	realAliceHand := make([][]byte, 0)

	for _, cb := range bobHand {
		realBobHand = append(realBobHand, Db(cb))
	}
	for _, ca := range aliceHand {
		realAliceHand = append(realAliceHand, Da(ca))
	}

	// The two players reveal their secret keys; now each player
	// can check that the other was actually dealt the cards he claimed to have played.
	fmt.Println(realBobHand)
	fmt.Println(realAliceHand)

	// Bob reveals his passphrase, and Alice verifies that she can replicate the above claim by Bob
	revealedBobPassphrase := passBob
	_, revealedDb := getED(revealedBobPassphrase)
	for i := range bobHand {
		if !reflect.DeepEqual(revealedDb(bobHand[i]), realBobHand[i]) {
			panic("Bob may have cheated!")
		}
	}

	// Alice reveals her passphrase, and Bob can do the same
	revealedAlicePassphrase := passAlice
	_, revealedDa := getED(revealedAlicePassphrase)

	for i := range aliceHand {
		if !reflect.DeepEqual(revealedDa(aliceHand[i]), realAliceHand[i]) {
			panic("Alice may have cheated!")
		}
	}

}

func getED(passphrase string) (func([]byte) []byte, func([]byte) []byte) {

	e := func(src []byte) []byte {
		key := sha256.Sum256([]byte(passphrase))
		nonce := make([]byte, chacha20poly1305.NonceSizeX)

		ecipher, _ := chacha20.NewUnauthenticatedCipher(key[:32], nonce)

		res := make([]byte, len(src))
		ecipher.XORKeyStream(res, src)
		return res
	}

	d := func(src []byte) []byte {
		key := sha256.Sum256([]byte(passphrase))
		nonce := make([]byte, chacha20poly1305.NonceSizeX)

		dcipher, _ := chacha20.NewUnauthenticatedCipher(key[:32], nonce)

		res := make([]byte, len(src))
		dcipher.XORKeyStream(res, src)
		return res
	}

	return d, e
}

type card struct {
	Rank string
	Suit string
}

func buildNewDeck() []card {
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	var deck []card
	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, card{Rank: r, Suit: s})
		}
	}

	return deck
}

func (c card) name() []byte {
	return []byte(c.Rank + " of " + c.Suit)
}
