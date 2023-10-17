# commercionetworkethermint

## Deploy di smart contract

### Hardhat

- creazione di un nuovo progetto con `npx hardhat init`
  - selezionare `javascript` dal menu stampato
- sostituito il contratto (`Lock.sol`) creato dal commando precedente con `HelloWorld.sol` 

>HelloWorld:
>
>```solidity
>pragma solidity ^0.8.9;
>
>contract HelloWorld {
>    string hello = "hello";
>
>    function helloworld() public returns (string memory){
>        return hello;
>    }
>}
>```

- compilazione del contratto con il commando `npx hardhat compile`
- configurazione di `scripts/deploy.js`
  - modificato lo script di deploy per addattarlo al contratto `HelloWorld.sol`

>scripts/deploy.js:
>
>```js
>const hre = require("hardhat");
>
>async function main() {
>  const contract = await hre.ethers.deployContract("HelloWorld");
>
>  await contract.waitForDeployment();
>
>  const a = await contract.getAddress();
>  console.log(`Contract deployed to ` + a);
>
>}
>
>// We recommend this pattern to be able to use async/await everywhere
>// and properly handle errors.
>main().catch((error) => {
>  console.error(error);
>  process.exitCode = 1;
>});
>```

- configurazione di `hardhat.config.js` per poter fare deploy su `commercionetworkethermint` in localhost
    - aggiunta della json-rpc in localhost e la chiave privata dell'account che fa la deploy

>hardhat.config.js:
>
>```js
>require("@nomicfoundation/hardhat-toolbox");
>
>
>/** @type import('hardhat/config').HardhatUserConfig */
>module.exports = {
>  solidity: "0.8.19",
>  networks: {
>    commercionetwork: {
>      url: `http://127.0.0.1:8545`,
>      accounts: ["782c8fe0081d1a1ec17fc612400e31b6650113378e4e3812f16ae9fe20b5f84f"]
>    }
>  }
>};
>
>```

- deploy del contratto sulla nostra chain con `npx hardhat run scripts/deploy.js --network commercionetwork`

Arrivato a questo punto la chain restituisce il segjuente errore:

```bash
ProviderError: failed to check sender balance: sender balance < tx cost (0 < 1659632): insufficient funds: insufficient funds
    at HttpProvider.request (/home/cheikh/Lavoro/Hello_world/node_modules/hardhat/src/internal/core/providers/http.ts:88:21)
    at processTicksAndRejections (node:internal/process/task_queues:95:5)
    at HardhatEthersSigner.sendTransaction (/home/cheikh/Lavoro/Hello_world/node_modules/@nomicfoundation/hardhat-ethers/src/signers.ts:125:18)
    at ContractFactory.deploy (/home/cheikh/Lavoro/Hello_world/node_modules/ethers/src.ts/contract/factory.ts:111:24)
    at main (/home/cheikh/Lavoro/Hello_world/scripts/deploy.js:18:20)
```

### TRUFFLE

- creazione nuovo progetto e inizializzazione 

```bash
mkdir cosmoprj
cd cosmoprj
truffle init
```

- creazione nuovo contratto con `truffle create contract Hello`

- modifica del contratto `contracts/Hello.sol` con il seguente codice

```solidity
// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Hello {
  function sayHello() public pure returns (string memory)  {
    return "Hello, world!";
  }
}
```

- configurazione di `truffle-config.js` per poter fare la deploy su commercionetwork in localhost

```js
module.exports = {
  networks: {
    // Development network is just left as truffle's default settings
    commercionetwork: {
      host: "0.0.0.0",     // Localhost (default: none)
      port: 8545,            // Standard Ethereum port (default: none)
      network_id: "*",       // Any network (default: none)
      gas: 5000000,          // Gas sent with each transaction
      gasPrice: 1000000000,  // 1 gwei (in wei)
    },
  },
  compilers: {
    solc: {
      version: "0.5.17",
    },
  },
}
```

- compilazione del contratto con `truffle compile`

output:

```bash
Compiling your contracts...
===========================
> Compiling ./contracts/Hello.sol
> Artifacts written to /home/cheikh/Lavoro/cosmoprj/build/contracts
> Compiled successfully using:
   - solc: 0.5.17+commit.d19bba13.Emscripten.clang
```

- deploy del contratto con `truffle migrate --network commercionetwork`

output:

```bash
Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.
Network up to date.
```
> Il contratto NON è stato inviato sulla rete! Truffle *non ha fatto niente*.

### CURL

- controllo accounts e balance dall'interfaccia

```bash
curl --data '{"jsonrpc":"2.0","method":"eth_coinbase", "id":1}' -H "Content-Type: application/json" localhost:8545
{"jsonrpc":"2.0","id":1,"result":"0xF0AF26d6b7D4Ed90CF70A591428127d2A48349Dc"}

curl --data '{"jsonrpc":"2.0","method":"eth_getBalance", "params": ["0xF0AF26d6b7D4Ed90CF70A591428127d2A48349Dc", "latest"], "id":2}' -H "Content-Type: application/json" localhost:8545
{"jsonrpc":"2.0","id":2,"result":"0x5f5e100"}
```

- compilazione del contratto con `solc --input-file contracts/HelloWorld.sol --bin`

risultato:

```bash
======= contracts/HelloWorld.sol:HelloWorld =======
Binary:
60806040526040518060400160405280600581526020017f68656c6c6f0000000000000000000000000000000000000000000000000000008152505f90816100479190610293565b50348015610053575f80fd5b50610362565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d457607f821691505b6020821081036100e7576100e6610090565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010e565b610153868361010e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019761019261018d8461016b565b610174565b61016b565b9050919050565b5f819050919050565b6101b08361017d565b6101c46101bc8261019e565b84845461011a565b825550505050565b5f90565b6101d86101cc565b6101e38184846101a7565b505050565b5b81811015610206576101fb5f826101d0565b6001810190506101e9565b5050565b601f82111561024b5761021c816100ed565b610225846100ff565b81016020851015610234578190505b610248610240856100ff565b8301826101e8565b50505b505050565b5f82821c905092915050565b5f61026b5f1984600802610250565b1980831691505092915050565b5f610283838361025c565b9150826002028217905092915050565b61029c82610059565b67ffffffffffffffff8111156102b5576102b4610063565b5b6102bf82546100bd565b6102ca82828561020a565b5f60209050601f8311600181146102fb575f84156102e9578287015190505b6102f38582610278565b86555061035a565b601f198416610309866100ed565b5f5b828110156103305784890151825560018201915060208501945060208101905061030b565b8683101561034d5784890151610349601f89168261025c565b8355505b6001600288020188555050505b505050505050565b6102178061036f5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80632f2f48591461002d575b5f80fd5b61003561004b565b6040516100429190610164565b60405180910390f35b60605f8054610059906101b1565b80601f0160208091040260200160405190810160405280929190818152602001828054610085906101b1565b80156100d05780601f106100a7576101008083540402835291602001916100d0565b820191905f5260205f20905b8154815290600101906020018083116100b357829003601f168201915b5050505050905090565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156101115780820151818401526020810190506100f6565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610136826100da565b61014081856100e4565b93506101508185602086016100f4565b6101598161011c565b840191505092915050565b5f6020820190508181035f83015261017c818461012c565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806101c857607f821691505b6020821081036101db576101da610184565b5b5091905056fea2646970667358221220d221fb0e3658dcf383726583fcfb95d6dada037b64f247d739d8df3b7e7e359364736f6c63430008150033
```

- controllo della stima di gas che serve per la deploy del contratto

```bash
BINARY=0x60806040526040518060400160405280600581526020017f68656c6c6f0000000000000000000000000000000000000000000000000000008152505f90816100479190610293565b50348015610053575f80fd5b50610362565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d457607f821691505b6020821081036100e7576100e6610090565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010e565b610153868361010e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019761019261018d8461016b565b610174565b61016b565b9050919050565b5f819050919050565b6101b08361017d565b6101c46101bc8261019e565b84845461011a565b825550505050565b5f90565b6101d86101cc565b6101e38184846101a7565b505050565b5b81811015610206576101fb5f826101d0565b6001810190506101e9565b5050565b601f82111561024b5761021c816100ed565b610225846100ff565b81016020851015610234578190505b610248610240856100ff565b8301826101e8565b50505b505050565b5f82821c905092915050565b5f61026b5f1984600802610250565b1980831691505092915050565b5f610283838361025c565b9150826002028217905092915050565b61029c82610059565b67ffffffffffffffff8111156102b5576102b4610063565b5b6102bf82546100bd565b6102ca82828561020a565b5f60209050601f8311600181146102fb575f84156102e9578287015190505b6102f38582610278565b86555061035a565b601f198416610309866100ed565b5f5b828110156103305784890151825560018201915060208501945060208101905061030b565b8683101561034d5784890151610349601f89168261025c565b8355505b6001600288020188555050505b505050505050565b6102178061036f5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c80632f2f48591461002d575b5f80fd5b61003561004b565b6040516100429190610164565b60405180910390f35b60605f8054610059906101b1565b80601f0160208091040260200160405190810160405280929190818152602001828054610085906101b1565b80156100d05780601f106100a7576101008083540402835291602001916100d0565b820191905f5260205f20905b8154815290600101906020018083116100b357829003601f168201915b5050505050905090565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156101115780820151818401526020810190506100f6565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610136826100da565b61014081856100e4565b93506101508185602086016100f4565b6101598161011c565b840191505092915050565b5f6020820190508181035f83015261017c818461012c565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806101c857607f821691505b6020821081036101db576101da610184565b5b5091905056fea2646970667358221220d221fb0e3658dcf383726583fcfb95d6dada037b64f247d739d8df3b7e7e359364736f6c63430008150033

curl --data '{"jsonrpc":"2.0","method": "eth_estimateGas", "params": [{"from": "0xF0AF26d6b7D4Ed90CF70A591428127d2A48349Dc", "data": "'$BINARY'"}], "id": 5}' -H "Content-Type: application/json" localhost:8545
```

> Il commando precedente ritorna il seguente errore: `{"jsonrpc":"2.0","id":5,"error":{"code":-32000,"message":"rpc error: code = Unknown desc = invalid opcode: PUSH0"}}`

- deploy del contratto: `curl --data '{"jsonrpc":"2.0","method": "eth_sendTransaction", "params": [{"from": "0xF0AF26d6b7D4Ed90CF70A591428127d2A48349Dc", "gas": "0x1c31e", "data": "'$BINARY'"}], "id": 6}' -H "Content-Type: application/json" localhost:8545`

> Il commando precedente ritorna il seguente errore: `{"jsonrpc":"2.0","id":6,"error":{"code":-32000,"message":"method handler crashed"}}`

## Riferimenti e tutorials

- https://docs.evmos.org/protocol/evmos-cli/configuration#pruning
- https://ethereum.org/en/developers/docs/development-networks/
- https://ethereum.org/en/developers/local-environment/
- https://hardhat.org/hardhat-runner/docs/getting-started
- https://ethereum.org/en/developers/docs/apis/json-rpc/#deploying-contract
- https://ethereum.org/en/developers/docs/apis/json-rpc/
- https://docs.evmos.org/protocol/concepts/transactions