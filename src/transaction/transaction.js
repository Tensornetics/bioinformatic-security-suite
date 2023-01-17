class Transaction {
  constructor(fromAddress, toAddress, amount) {
    this.fromAddress = fromAddress;
    this.toAddress = toAddress;
    this.amount = amount;
    this.timestamp = new Date();
  }

  calculateHash() {
    // code to calculate the transaction's hash
  }

  signTransaction(signingKey) {
    // code to sign the transaction using the signingKey
  }

  isValid() {
    // code to check if the transaction is valid
  }
}

module.exports = Transaction;
