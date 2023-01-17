const WebSocket = require('ws');
const Blockchain = require('../blockchain/blockchain');

class Node {
  constructor(port) {
    this.blockchain = new Blockchain();
    this.sockets = [];
    this.server = new WebSocket.Server({ port });
    this.server.on('connection', socket => this.connectSocket(socket));
  }
  connectSocket(socket) {
    this.sockets.push(socket);
    console.log('Socket connected');
  }
  broadcast(data) {
    this.sockets.forEach(socket => socket.send(JSON.stringify(data)));
  }
  handleBlockchainResponse(socket) {
    socket.on('message', message => {
      const data = JSON.parse(message);
      this.blockchain.replaceChain(data);
    });
  }
}

module.exports = Node;
