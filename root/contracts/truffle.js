/*
 * NB: since truffle-hdwallet-provider 0.0.5 you must wrap HDWallet providers in a 
 * function when declaring them. Failure to do so will cause commands to hang. ex:
 * ```
 * mainnet: {
 *     provider: function() { 
 *       return new HDWalletProvider(mnemonic, 'https://mainnet.infura.io/<infura-key>') 
 *     },
 *     network_id: '1',
 *     gas: 4500000,
 *     gasPrice: 10000000000,
 *   },
 */

module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!

    compiler: {
        solc: "0.4.23"
    },

    networks: {
        development: {
            host: 'localhost',
            port: 8545,
            network_id: '10',
            from: '0x70043fa2d7b7125b8f6351b4b6b96c748f8a8336',
            gas: 4000000,
            gasPrice: 10000000000
        }
    }
};
