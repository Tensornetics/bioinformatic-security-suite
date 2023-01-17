//block.js
class Block {
  constructor(timestamp, data) {
    this.timestamp = timestamp;
    this.data = data;
    this.previousHash = null;
    this.hash = this.calculateHash();
  }

  calculateHash() {
    // code to calculate the block's hash
  }
}

//blockchain.js
class Blockchain {
  constructor() {
    this.chain = [this.createGenesisBlock()];
  }

  createGenesisBlock() {
    // code to create the genesis block
  }

  getLatestBlock() {
    // code to return the latest block in the chain
  }

  addBlock(newBlock) {
    // code to add a new block to the chain
  }
}
