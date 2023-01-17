class Mod {
  constructor(options) {
    this.options = options;
  }

  performModOperation(data) {
    // code to perform the mod operation
    return data % this.options.divisor;
  }
}

module.exports = Mod;
