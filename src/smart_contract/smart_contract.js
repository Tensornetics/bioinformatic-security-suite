const Web3 = require('web3');

class SmartContract {
  constructor(address) {
    this.web3 = new Web3(new Web3.providers.HttpProvider(address));
  }

  deploy(abi, bytecode, args) {
    return this.web3.eth.contract(abi).deploy({ data: bytecode, arguments: args });
  }

  call(contract, functionName, args) {
    return contract.methods[functionName](...args).call();
  }

  sendTransaction(contract, functionName, args) {
    return contract.methods[functionName](...args).send();
  }
}

module.exports = SmartContract;
