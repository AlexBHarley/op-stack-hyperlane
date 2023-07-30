# op-stack-hyperlane

An OP Stack Mod that bundles a [Hyperlane](https://hyperlane.xyz) ⏩ deployment into your L2. With Hyperlane now a core part your L2 you benefit from:

- opt in fast token withdrawals 🔥
- predeployed Mailbox's and associated ISMs
  - L1 => L2 leverages an Optimism ISM that maintains the security of Ethereum
  - L2 => L1 configured with an M of N multisig
- default relayer and validator infrastructure

See the diff with vanilla Optimism [here](https://github.com/AlexBHarley/op-stack-hyperlane/compare/op-stack-checkpoint...develop).

## What's included?

### Predeploys

- [Mailbox](https://docs.hyperlane.xyz/docs/protocol/messaging), available at `0x4200000000000000000000000000000000000068`

  Allows for bidirectional communication between the L1 and your L2 with Hyperlane

- [OptimismISM](https://docs.hyperlane.xyz/docs/protocol/sovereign-consensus/hook-ism), available at `0x420000000000000000000000000000000000006b`

  Allows Hyperlane messages from the L1 to your L2 to inherit the security guarantees of the L1.

- [ValidatorAnnounce](https://docs.hyperlane.xyz/docs/operators/validators/setup#announcing-your-validator), available at `0x420000000000000000000000000000000000006c`

  Enables Hyperlane validators to register themselves as validating withdrawals.

### Infrastructure

- Default ISM validators

  A Multisig ISM configured with the M of N genesis validator set, leveraged for messages from L2 to L1.

- Gas paymaster and relayer

  A relayer configured to work with the default interchain gas paymaster.

## Fast withdrawals

The core functionality offered by `op-stack-hyperlane` is the ability to modularise token withdrawal security. The 7 day finalization time offered by OP Mainnet works well for the largest of transfers, but the rise of bridges that leverage other security models speaks volumes to the user experience of this.

With `op-stack-hyperlane`, withdrawal security is configured on a per token basis. Meaning withdrawals finalised according to any logic you can imagine. An example that could be implemented is the following,

- USDC withdrawals < $10,000 USD, 5/10 validator attestations
- USDC withdrawals > $10,000 USD, 9/10 validator attestations

For the largest of transfers, a fraud proof window could also be added to somewhat mimic the native bridge,

- USDC withdrawals > $1,000,000 USD, an [AggregationISM](https://docs.hyperlane.xyz/docs/protocol/sovereign-consensus/aggregation-ism) comprising 9/10 validator attestations + an M of N 3 day fraud proof window

### Demo

The quicket way to familiarise yourself with this mod is to spin up the devnet and transfer some ERC20 tokens back and forth between the L1 and L2. Following along with the below commands will get you setup and running in under five minutes.

We'll first setup our devnet,

```
make devnet-up-deploy
```

which will run both the L1 and L2 pieces of the rollup, super handy. After that's completed we'll set some variables to make deploying and interacting with our contracts a bit easier.

```
export L2=http://localhost:9545
export L1=http://localhost:8545
export PRIVATE_KEY=0xc481977dbc3300c306daad7cb02ac13930f559e799b2e4e3c0d3ad80cb0f88dc
export FUNDED_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
ADDRESS=$(cast wallet address --private-key=$PRIVATE_KEY)

cast send $ADDRESS --value 100ether --rpc-url $L2 --private-key $FUNDED_PRIVATE_KEY
cast send $ADDRESS --value 100ether --rpc-url $L1 --private-key $FUNDED_PRIVATE_KEY
```

The `FUNDED_PRIVATE_KEY` comes from the [devnet initialisation file](bedrock-devnet/devnet/__init__.py) and gets set as the admin for a bunch of the proxy contracts. So we want to just transfer some ether to a different account which we can use for interacting with said contracts.

Next we need to deploy the contracts for the token we want to transfer between chains. We'll do the rest of this walkthrough from the `contracts-bedrock` package,

```
cd packages/contracts-bedrock
```

Then run the deployment scripts,

```
forge script scripts/DeployERC20.s.sol:Deploy --private-key $PRIVATE_KEY --broadcast --rpc-url $L1
forge script scripts/DeployMintableERC20.s.sol:Deploy --private-key $PRIVATE_KEY --broadcast --rpc-url $L2
```

These scripts deploy two sides of the same coin. On the L1 we have a normal `ERC20` contract and on the L2 we use a special `ERC20` that can only be minted by the `L2StandardBridge`.

Now we can approve and deposit the tokens from our L1,

```
cast send 0x6523Fd7e1B28113aB7350B7C312b541fd4866492 "approve(address,uint256)" 0x0165878A594ca255338adfa4d48449f69242Eb8F 100000 --rpc-url $L1 --private-key $PRIVATE_KEY
cast send 0x0165878A594ca255338adfa4d48449f69242Eb8F "depositERC20(address,address,uint256,uint32,bytes)" 0x6523Fd7e1B28113aB7350B7C312b541fd4866492 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2 1 300000 "01" --private-key $PRIVATE_KEY --rpc-url $L1
```

Before checking that they were received on the L2,

````
cast call 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2 "balanceOf(address)" "$ADDRESS" --rpc-url $L2```
0x0000000000000000000000000000000000000000000000000000000000000001
````

At this stage we've simply transferred some tokens from L1 to L2, which is great but this functionality comes bundled with Optimism natively. We haven't done anything Hyperlane specific. Now we'll be showing how we can use Hyperlane to modularise the architecture of our deployment and enable fast withdrawals. Keep in mind that the usual withdraw, prove, finalise flow takes 7 days on OP mainnet. We want to achieve near instantaneous withdrawals.

The first step is to call a function, `registerTokenForMailboxWithdrawals(address)`, on the `L2StandardBridge`. This marks the token as one that will go via the Hyperlane `Mailbox` for withdrawals as opposed to the native Optimism `CrossDomainMessenger`.

```
cast send 0x4200000000000000000000000000000000000010 "registerTokenForMailboxWithdrawals(address)" 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2 --private-key $PRIVATE_KEY --rpc-url $L2
```

Now we'll can `approve` and `withdraw` tokens from the L2,

```
cast send 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2 "approve(address,uint256)" 0x4200000000000000000000000000000000000010  --rpc-url $L2 --private-key $PRIVATE_KEY
cast send 0x4200000000000000000000000000000000000010 "withdraw(address,uint256,uint32,bytes)" 0x7c6b91D9Be155A6Db01f749217d76fF02A7227F2 1 300000 "01" --private-key $PRIVATE_KEY --rpc-url $L2
```

Almost instantly these tokens are now available on the L1,

```
cast call 0x6523Fd7e1B28113aB7350B7C312b541fd4866492 "balanceOf(address)" $ADDRESS --rpc-url $L1 --private-key $PRIVATE_KEY
0x00000000000000000000000000000000000000000000d3c21bcecceda1000000
```

<div align="center">
  <br />
  <br />
  <a href="https://optimism.io"><img alt="Optimism" src="https://raw.githubusercontent.com/ethereum-optimism/brand-kit/main/assets/svg/OPTIMISM-R.svg" width=600></a>
  <br />
  <h3><a href="https://optimism.io">Optimism</a> is Ethereum, scaled.</h3>
  <br />
</div>

## What is Optimism?

[Optimism](https://www.optimism.io/) is a project dedicated to scaling Ethereum's technology and expanding its ability to coordinate people from across the world to build effective decentralized economies and governance systems. The [Optimism Collective](https://app.optimism.io/announcement) builds open-source software for running L2 blockchains and aims to address key governance and economic challenges in the wider cryptocurrency ecosystem. Optimism operates on the principle of **impact=profit**, the idea that individuals who positively impact the Collective should be proportionally rewarded with profit. **Change the incentives and you change the world.**

In this repository, you'll find numerous core components of the OP Stack, the decentralized software stack maintained by the Optimism Collective that powers Optimism and forms the backbone of blockchains like [OP Mainnet](https://explorer.optimism.io/) and [Base](https://base.org). Designed to be "aggressively open source," the OP Stack encourages you to explore, modify, extend, and test the code as needed. Although not all elements of the OP Stack are contained here, many of its essential components can be found within this repository. By collaborating on free, open software and shared standards, the Optimism Collective aims to prevent siloed software development and rapidly accelerate the development of the Ethereum ecosystem. Come contribute, build the future, and redefine power, together.

## Documentation

- If you want to build on top of OP Mainnet, refer to the [Optimism Community Hub](https://community.optimism.io)
- If you want to build your own OP Stack based blockchain, refer to the [OP Stack docs](https://stack.optimism.io)
- If you want to contribute to the OP Stack, check out the [Protocol Specs](./specs)

## Community

General discussion happens most frequently on the [Optimism discord](https://discord.gg/optimism).
Governance discussion can also be found on the [Optimism Governance Forum](https://gov.optimism.io/).

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of the contributing process for this repository.
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Optimism Monorepo.
Then check out the list of [Good First Issues](https://github.com/ethereum-optimism/optimism/contribute) to find something fun to work on!

## Security Policy and Vulnerability Reporting

Please refer to the canonical [Security Policy](https://github.com/ethereum-optimism/.github/blob/master/SECURITY.md) document for detailed information about how to report vulnerabilities in this codebase.
Bounty hunters are encouraged to check out [the Optimism Immunefi bug bounty program](https://immunefi.com/bounty/optimism/).
The Optimism Immunefi program offers up to $2,000,042 for in-scope critical vulnerabilities.

## The Bedrock Upgrade

OP Mainnet is currently preparing for [its next major upgrade, Bedrock](https://dev.optimism.io/introducing-optimism-bedrock/).
You can find detailed specifications for the Bedrock upgrade within the [specs folder](./specs) in this repository.

Please note that a significant number of packages and folders within this repository are part of the Bedrock upgrade and are NOT currently running in production.
Refer to the Directory Structure section below to understand which packages are currently running in production and which are intended for use as part of the Bedrock upgrade.

## Directory Structure

<pre>
~~ Production ~~
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts.
│   ├── <a href="./packages/contracts-periphery">contracts-periphery</a>: Peripheral contracts for Optimism
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Optimism easier
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/fault-detector">fault-detector</a>: Service for detecting Sequencer faults
│   ├── <a href="./packages/replica-healthcheck">replica-healthcheck</a>: Service for monitoring the health of a replica node
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Optimism
├── <a href="./op-bindings">op-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./op-batcher">op-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./op-bootnode">op-bootnode</a>: Standalone op-node discovery bootnode
├── <a href="./op-chain-ops">op-chain-ops</a>: State surgery utilities
├── <a href="./op-challenger">op-challenger</a>: Dispute game challenge agent
├── <a href="./op-e2e">op-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./op-exporter">op-exporter</a>: Prometheus exporter client
├── <a href="./op-heartbeat">op-heartbeat</a>: Heartbeat monitor service
├── <a href="./op-node">op-node</a>: rollup consensus-layer client
├── <a href="./op-program">op-program</a>: Fault proof program
├── <a href="./op-proposer">op-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./op-service">op-service</a>: Common codebase utilities
├── <a href="./op-signer">op-signer</a>: Client signer
├── <a href="./op-wheel">op-wheel</a>: Database utilities
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
└── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade

~~ Pre-BEDROCK ~~
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-periphery">contracts-periphery</a>: Peripheral contracts for Optimism
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Optimism easier
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/fault-detector">fault-detector</a>: Service for detecting Sequencer faults
│   ├── <a href="./packages/replica-healthcheck">replica-healthcheck</a>: Service for monitoring the health of a replica node
│   └── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Optimism
├── <a href="./indexer">indexer</a>: indexes and syncs transactions
├── <a href="./op-exporter">op-exporter</a>: A prometheus exporter to collect/serve metrics from an Optimism node
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
└── <a href="./technical-documents">technical-documents</a>: audits and post-mortem documents
</pre>

## Branching Model

### Active Branches

| Branch                                                                 | Status                                                                                                |
| ---------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| [master](https://github.com/ethereum-optimism/optimism/tree/master/)   | Accepts PRs from `develop` when intending to deploy to production.                                    |
| [develop](https://github.com/ethereum-optimism/optimism/tree/develop/) | Accepts PRs that are compatible with `master` OR from `release/X.X.X` branches.                       |
| release/X.X.X                                                          | Accepts PRs for all changes, particularly those not backwards compatible with `develop` and `master`. |

### Overview

This repository generally follows [this Git branching model](https://nvie.com/posts/a-successful-git-branching-model/).
Please read the linked post if you're planning to make frequent PRs into this repository.

### Production branch

The production branch is `master`.
The `master` branch contains the code for latest "stable" releases.
Updates from `master` **always** come from the `develop` branch.

### Development branch

The primary development branch is [`develop`](https://github.com/ethereum-optimism/optimism/tree/develop/).
`develop` contains the most up-to-date software that remains backwards compatible with the latest experimental [network deployments](https://community.optimism.io/docs/useful-tools/networks/).
If you're making a backwards compatible change, please direct your pull request towards `develop`.

**Changes to contracts within `packages/contracts-bedrock/contracts` are usually NOT considered backwards compatible and SHOULD be made against a release candidate branch**.
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a release candidate branch has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using the latest release candidate branch.
See below for info about release candidate branches.

### Release candidate branches

Branches marked `release/X.X.X` are **release candidate branches**.
Changes that are not backwards compatible and all changes to contracts within `packages/contracts-bedrock/contracts` MUST be directed towards a release candidate branch.
Release candidates are merged into `develop` and then into `master` once they've been fully deployed.
We may sometimes have more than one active `release/X.X.X` branch if we're in the middle of a deployment.
See table in the **Active Branches** section above to find the right branch to target.

## Releases

### Changesets

We use [changesets](https://github.com/changesets/changesets) to mark packages for new releases.
When merging commits to the `develop` branch you MUST include a changeset file if your change would require that a new version of a package be released.

To add a changeset, run the command `pnpm changeset` in the root of this monorepo.
You will be presented with a small prompt to select the packages to be released, the scope of the release (major, minor, or patch), and the reason for the release.
Comments within changeset files will be automatically included in the changelog of the package.

### Triggering Releases

Releases can be triggered using the following process:

1. Create a PR that merges the `develop` branch into the `master` branch.
2. Wait for the auto-generated `Version Packages` PR to be opened (may take several minutes).
3. Change the base branch of the auto-generated `Version Packages` PR from `master` to `develop` and merge into `develop`.
4. Create a second PR to merge the `develop` branch into the `master` branch.

After merging the second PR into the `master` branch, packages will be automatically released to their respective locations according to the set of changeset files in the `develop` branch at the start of the process.
Please carry this process out exactly as listed to avoid `develop` and `master` falling out of sync.

**NOTE**: PRs containing changeset files merged into `develop` during the release process can cause issues with changesets that can require manual intervention to fix.
It's strongly recommended to avoid merging PRs into develop during an active release.

## License

All other files within this repository are licensed under the [MIT License](https://github.com/ethereum-optimism/optimism/blob/master/LICENSE) unless stated otherwise.

```

```
