// Package diffiehellman provides functions necessary for exchanging
// keys using the Diffie-Hellman key exchange protocol.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey selects a private key, an integer, between 1 and the
// supplied prime number, p.
func PrivateKey(p *big.Int) *big.Int {

	// calculate the max value to ensure the random number generated
	// lies in the range 1 < n < p.
	max := big.NewInt(0)
	max = max.Sub(p, big.NewInt(2))

	// generate the random number and adjust for the offest applied above
	privKey, _ := rand.Int(rand.Reader, max)
	privKey.Add(privKey, big.NewInt(2))

	return privKey
}

// PublicKey generates the public key corresponding to the supplied
// private key, prime number and primitive root, g.
func PublicKey(private, p *big.Int, g int64) *big.Int {

	// calculate the public key based on the following formula
	// pubKey = g**privKey mod p
	G := big.NewInt(g)
	pubKey := G.Exp(G, private, p)

	return pubKey
}

// NewPair returns a corresponding pair of public and private keys.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	privKey := PrivateKey(p)
	pubKey := PublicKey(privKey, p, g)
	return privKey, pubKey
}

// SecretKey returns the shared secret for a given key pair.
func SecretKey(private1, public2, p *big.Int) *big.Int {

	// calculate the secret key based on the following formula
	// secKey = pubKey**privKey mod p
	pk := *public2
	secKey := pk.Exp(public2, private1, p)

	return secKey
}
