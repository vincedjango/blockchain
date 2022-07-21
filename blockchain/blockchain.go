package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/****************
*   Interface   *
*****************/
type Tblockchain interface {
	AddBlock(from, to string, amount float32)
}

/*************
*   Struct   *
**************/
type BlockBTC struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type BlockADA struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type BlockchainBTC struct {
	genesisBlock BlockBTC
	chain        []BlockBTC
	difficulty   int
}

type BlockchainADA struct {
	genesisBlock BlockADA
	chain        []BlockADA
	difficulty   int
}

func (b BlockBTC) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b BlockADA) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *BlockBTC) mine(difficulty int) {
	var i int
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
		i++
	}
	fmt.Println(i)
}

func (b *BlockADA) mine(difficulty int) {
	var i int
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
		i++
	}
	fmt.Println(i)
}

func CreateBlockchainBTC(difficulty int) *BlockchainBTC {
	genesisBlock := BlockBTC{
		hash:      "0",
		timestamp: time.Now(),
	}
	return &BlockchainBTC{
		genesisBlock,
		[]BlockBTC{genesisBlock},
		difficulty,
	}
}

func CreateBlockchainADA(difficulty int) *BlockchainADA {
	genesisBlock := BlockADA{
		hash:      "0",
		timestamp: time.Now(),
	}
	return &BlockchainADA{
		genesisBlock,
		[]BlockADA{genesisBlock},
		difficulty,
	}
}

func IAddBlock(bc Tblockchain, from, to string, amount float32) {
	bc.AddBlock(from, to, amount)
}

func (b *BlockchainBTC) AddBlock(from, to string, amount float32) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := BlockBTC{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b *BlockchainADA) AddBlock(from, to string, amount float32) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := BlockADA{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b BlockchainBTC) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (b BlockchainADA) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
