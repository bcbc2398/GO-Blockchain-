package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp time.Time
	transaction []string
	prevHash []byte
	Hash []byte
}

func main() {
	genesisTransactions := []string{"Tony sent Montana 50 MATIC", "Montana sent Tony 30 MATIC"}
	genesisBlock := NewBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInformation(genesisBlock)

	block2Transactions := []string{"Ricky sent Tony 30 MATIC"}
	block2 := NewBlock(block2Transactions, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInformation(block2)

	block3Transactions := []string{"Montana sent Tony 45 MATIC", "Tony sent Montana 10 MATIC"}
	block3 := NewBlock(block3Transactions, block2.Hash)
	fmt.Println("--- Third Block ---")
	printBlockInformation(block3)
}

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block {
		timestamp: currentTime,
		transactions: transactions,
		prevHash: prevHash,
		Hash: NewHash(currentTime, transactions, prevHash),
	}
}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}

