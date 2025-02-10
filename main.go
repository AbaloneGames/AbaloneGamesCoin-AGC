package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Transaction structure
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

// Block structure
type Block struct {
	Index        int
	Timestamp    string
	PrevHash     string
	Hash         string
	Nonce        int
	Transactions []Transaction
}

// Blockchain Slice
var Blockchain []Block
var transactionPool []Transaction

// Calculate hash function
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%d", block.Index, block.Timestamp, block.PrevHash, block.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// Mine a new block
func mineBlock(prevBlock Block, transactions []Transaction) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		PrevHash:     prevBlock.Hash,
		Transactions: transactions,
	}

	nonce := 0
	for {
		newBlock.Nonce = nonce
		newBlock.Hash = calculateHash(newBlock)
		if newBlock.Hash[:4] == "0000" { // Adjust mining difficulty
			break
		}
		nonce++
	}
	return newBlock
}

// Create the Genesis Block
func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		PrevHash:  "0",
		Nonce:     0,
		Hash:      calculateHash(Block{Index: 0}),
	}
}

// Add a new transaction
func addTransaction(sender string, receiver string, amount float64) {
	transaction := Transaction{Sender: sender, Receiver: receiver, Amount: amount}
	transactionPool = append(transactionPool, transaction)
}

func main() {
	// Initialize Blockchain
	genesisBlock := createGenesisBlock()
	Blockchain = append(Blockchain, genesisBlock)

	// Simulating Transactions
	addTransaction("Alice", "Bob", 10.5)
	addTransaction("Charlie", "Dave", 5.0)

	// Mine a new block
	newBlock := mineBlock(Blockchain[len(Blockchain)-1], transactionPool)
	Blockchain = append(Blockchain, newBlock)
	transactionPool = []Transaction{}

	// Print Blockchain
	for _, block := range Blockchain {
		fmt.Printf("\nIndex: %d\nTimestamp: %s\nPrevHash: %s\nHash: %s\nNonce: %d\n",
			block.Index, block.Timestamp, block.PrevHash, block.Hash, block.Nonce)
		for _, tx := range block.Transactions {
			fmt.Printf("Transaction: %s -> %s (%f AGC)\n", tx.Sender, tx.Receiver, tx.Amount)
		}
	}

	fmt.Println("\nBlockchain is running...")
}
