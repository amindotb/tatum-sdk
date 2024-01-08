<p align="center">
  <a href="https://tatum.io/">
    <img src="https://assets-global.website-files.com/62624e283b503f3e68275638/62624e283b503fde012757c1_Light.svg" alt="Logo" width="200" height="100">
  </a>
</p>

<h3 align="center">
**Unofficial**
Tatum
SDK</h3>

<p align="center">
  Welcome to Tatum SDK - Golang Library for Simplifying Blockchain Development.<br>
  <a href="https://docs.tatum.io/"><strong>Documentation</strong></a>
  <br>
  <br>
</p>

<div align="center">

<a href="">[![GitHub license](https://img.shields.io/npm/dm/@tatumio/tatum)](https://img.shields.io/npm/dm/@tatumio/tatum)</a>
<a href="">[![npm version](https://img.shields.io/npm/v/@tatumio/tatum.svg?style=flat-square)](https://www.npmjs.com/package/@tatumio/tatum)</a>
<a href="">[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/tatumio/tatum-js/blob/master/LICENSE.txt)</a>
<a href="">[![Build](https://img.shields.io/github/actions/workflow/status/tatumio/tatum-js/build.yml?branch=master)](https://img.shields.io/github/actions/workflow/status/tatumio/tatum-js/build.yml?branch=master)</a>

</div>
<hr>

## 🚀 Tatum SDK
A powerful, feature-rich Golang library that streamlines the development of blockchain applications.

🔍 **Designed For Developers**
If you're looking to integrate blockchain functionalities into your projects, [Tatum SDK](https://docs.tatum.io/sdk/get-started-with-tatum-sdk) is for you.


| Documentation                                                                                             |
|-----------------------------------------------------------------------------------------------------------|
| **EVM Blockchains**                                                                                       |
| [Ethereum RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/ethereum-rpc-documentation)                 |
| [Polygon RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/polygon-rpc-documentation)                   |
| [Flare RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/flare-rpc-documentation)                       |
| [Haqq RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/haqq-rpc-documentation)                         |
| [Optimism RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/optimism-rpc-documentation)                 |
| [Horizen EON RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/horizen-eon-rpc-documentation)           |
| [Arbitrum One RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/arbitrum-rpc-documentation)             |
| [Chiliz RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/chiliz-rpc-documentation)                     |
| [Ethereum Classic RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/ethereum-classic-rpc-documentation) |
| [Klaytn RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/klaytn-rpc-documentation)                     |
| [Avalanche RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/avalanche-rpc-documentation)               |
| [Celo RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/celo-rpc-documentation)                         |
| [XinFin RPC](https://docs.tatum.io/docs/rpc/evm-blockchains/xinfin-rpc-documentation)                     |
| **UTXO Blockchains**                                                                                      |
| [Bitcoin RPC](https://docs.tatum.io/docs/rpc/utxo-blockchains/bitcoin-rpc-documentation)                  |
| [Litecoin RPC](https://docs.tatum.io/docs/rpc/utxo-blockchains/litecoin-rpc-documentation)                |
| [Dogecoin RPC](https://docs.tatum.io/docs/rpc/utxo-blockchains/dogecoin-rpc-documentation)                |
| [ZCash RPC](https://docs.tatum.io/docs/rpc/utxo-blockchains/zcash-rpc-documentation)                      |
| [Bitcoin Cash RPC](https://docs.tatum.io/docs/rpc/utxo-blockchains/bitcion-cash-rpc-documentation)        |
| **Other Blockchains**                                                                                     |
| [Solana RPC](https://docs.tatum.io/docs/rpc/solana-rpc-documentation)                                     |
| [XPR RPC](https://docs.tatum.io/docs/rpc/xrp-rpc-documentation)                                           |
| [Tron RPC](https://docs.tatum.io/docs/rpc/tron-rpc-documentation)                                         |
| [Eos RPC](https://docs.tatum.io/docs/rpc/eos-rpc-documentation)                                           |
| [Tezos RPC](https://docs.tatum.io/docs/rpc/tezos-rpc-documentation)                                       |
| [Agorand RPC](https://docs.tatum.io/docs/rpc/algo-rpc-documentation)                                      |
| [Cardano RPC](https://docs.tatum.io/docs/rpc/cardano-rpc-documentation)                                   |
| [Stellar RPC](https://docs.tatum.io/docs/rpc/stellar-rpc-documentation)                                   |


## Installation

To install TatumSDK, simply run the following command in your terminal or command prompt:

### Install using 
```console
go get github.com/amindotb/tatum-sdk
```

## Getting started

### Basic Usage

Here's a brief overview of how to utilize TatumSDK with an v4 API key.

### Initialization

Start by importing the TatumSDK library and initializing Ethereum client as follows:

```go
const apiKey = "YOU_API_KEY_HERE"
var api = NewAPI(client.WithAPIKey(apiKey))
```

### 🔔 Create Notifications

Effortlessly monitor wallet activities. Set up real-time notifications for events like:

| Documentation |
| ----- |
| [Start monitoring of the address](https://docs.tatum.io/docs/notifications/notification-workflow/start-monitoring-of-the-address) |
| [Stop monitoring of the address](https://docs.tatum.io/docs/notifications/notification-workflow/stop-monitoring-of-the-address) |
| [Get all sent notifications](https://docs.tatum.io/docs/notifications/notification-workflow/get-all-sent-notifications) |
| [Get all existing monitoring subscriptions](https://docs.tatum.io/docs/notifications/notification-workflow/get-all-existing-monitoring-subscriptions) |

#### Start monitoring of the address 
```go
var subscription = subscription.SubscriptionInput{
    Type: subscription.TYPE_ADDRESS_EVENT,
    Attribute: subscription.SubscriptionAttribute{
        Address: "0x8e12113ded05113d1a59bad8c37c295633ff6a97",
        Chain:   constants.CHAIN_ETH,
        Url:     "https://your.website/webhook",
    },
}
resp, err := api.subscription.Subscribe(subscription)
```

#### Stop monitoring of the address
```go
var subscriptionId = "SUBSCRIPTION_ID_HERE"
resp, err := api.subscription.Unsubscribe(subscriptionId)
```

#### Get all sent notifications
```go
var pagination = client.Pagination{
    Offset:   0,
    PageSize: 10,
}
var direction = subscription.DIRECTION_DESC
var filterFailed = false
resp, err := api.subscription.GetAllExecutedWebhooks(pagination, direction, filterFailed)
```

#### Get all existing monitoring subscriptions
```go
var pagination = client.Pagination{
    Offset:   0,
    PageSize: 10,
}
var address = "0x8e12113ded05113d1a59bad8c37c295633ff6a97"
resp, err := api.subscription.GetAll(pagination, address)
```


### 👛 Access Wallet Information

Through a single interface, obtain crucial wallet details such as balances, transaction history, and other pertinent information.

| Documentation |
| ----- |
| [Get all assets the wallet holds](https://docs.tatum.io/docs/wallet-address-operations/get-all-assets-the-wallet-holds) |
| [Get all transactions on the wallet](https://docs.tatum.io/docs/wallet-address-operations/get-all-transactions-on-the-wallet) |

#### Get all assets the wallet holds
```go
var pagination = client.Pagination{
    Offset:   0,
    PageSize: 10,
}
var addresses = []string{"0x8e12113ded05113d1a59bad8c37c295633ff6a97"}
resp, err := api.address.GetBalance(pagination, addresses, constants.CHAIN_ETHEREUM_MAINNET)
```

#### Get all transactions on the wallet
```go
var pagination = client.Pagination{
    Offset:   0,
    PageSize: 10,
}
var addresses = []string{"0x8e12113ded05113d1a59bad8c37c295633ff6a97"}
resp, err := api.address.GetTransactions(pagination, addresses, constants.CHAIN_ETHEREUM_MAINNET)
```