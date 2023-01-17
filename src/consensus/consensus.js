const axios = require('axios');
const Blockchain = require('../blockchain/blockchain');

class Consensus {
  constructor(blockchain) {
    this.blockchain = blockchain;
    this.networkNodes = [];
  }

  registerNode(address) {
    this.networkNodes.push(address);
  }

  async resolveConflicts() {
    const promises = this.networkNodes.map(node => axios.get(`http://${node}/blockchain`));
    const blocks = await Promise.all(promises);
    const maxLength = blocks.reduce((max, block) => (block.data.length > max ? block.data.length : max), 0);
    const longestChain = blocks.find(block => block.data.length === maxLength).data;
    if (longestChain && this.blockchain.chainIsValid(longestChain)) {
      this.blockchain.chain = longestChain;
      return true;
    }
    return false;
  }
}

module.exports = Consensus;
