package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func main() {
	// Definir o argumento de linha de comando para a chave privada em hexadecimal
	var hexKey string
	flag.StringVar(&hexKey, "key", "", "Chave privada em hexadecimal (sem prefixo 0x)")
	flag.Parse()

	// Verificar se a chave privada foi fornecida
	if hexKey == "" {
		fmt.Println("Por favor, forneça a chave privada usando o parâmetro -key.")
		return
	}

	// Converter a chave privada de hexadecimal para bytes
	privKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		fmt.Printf("Erro ao decodificar a chave privada: %v\n", err)
		return
	}

	// Adicionar o prefixo 0x80
	versionedPayload := append([]byte{0x80}, privKeyBytes...)

	// Realizar o primeiro hash SHA-256
	firstSHA := sha256.Sum256(versionedPayload)

	// Realizar o segundo hash SHA-256
	secondSHA := sha256.Sum256(firstSHA[:])

	// Adicionar os primeiros 4 bytes do segundo hash como checksum
	checksum := secondSHA[:4]
	fullPayload := append(versionedPayload, checksum...)

	// Codificar em Base58
	wif := base58.Encode(fullPayload)

	fmt.Println("Chave privada WIF:", wif)
}
