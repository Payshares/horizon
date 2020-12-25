---
title: Horizon
---

Horizon is the server for the client facing API for the Payshares ecosystem.  It acts as the interface between [payshares-core](https://www.payshares.org/developers/learn/payshares-core) and applications that want to access the Payshares network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Payshares ecosystem](https://www.payshares.org/developers/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://www.payshares.org/developers/js-payshares-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon-testnet.payshares.org/](https://horizon-testnet.payshares.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/payshares/js-payshares-sdk)
- [Java](https://github.com/payshares/java-payshares-sdk)
- [Go](https://github.com/payshares/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/payshares/ruby-payshares-sdk)
- [Python](https://github.com/PaysharesCN/py-payshares-base)
