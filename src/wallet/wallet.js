const crypto = require('crypto');

class Wallet {
  constructor() {
    this.address = null;
    this.keyPair = null;
    this.publicKey = null;
    this.balance = 0;
  }

  generateKeyPair() {
    this.keyPair = crypto.generateKeyPairSync('rsa', {
      modulusLength: 2048,
      publicKeyEncoding: {
        type: 'spki',
        format: 'pem'
      },
      privateKeyEncoding: {
        type: 'pkcs8',
        format: 'pem'
      }
    });
    this.publicKey = this.keyPair.publicKey;
    this.address = crypto
      .createHash('sha256')
      .update(this.publicKey)
      .digest('hex');
  }

  sign(data) {
    return crypto.sign('sha256', Buffer.from(data), this.keyPair.privateKey);
  }

  verify(data, signature) {
    return crypto.verify('sha256', Buffer.from(data), this.publicKey, signature);
  }
}

module.exports = Wallet;
