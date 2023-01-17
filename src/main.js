const Blockchain = require('./blockchain');
const myChain = new Blockchain();

myChain.addBlock(new Block(1, "Block 1 data"));
myChain.addBlock(new Block(2, "Block 2 data"));
myChain.addBlock(new Block(3, "Block 3 data"));

console.log(JSON.stringify(myChain, null, 4));
console.log("Is blockchain valid? " + myChain.isChainValid());
