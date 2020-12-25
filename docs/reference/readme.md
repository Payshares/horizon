---
title: Overview
---

Horizon is an API server for the Payshares ecosystem.  It acts as the interface between [payshares-core](https://github.com/payshares/payshares-core) and applications that want to access the Payshares network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Payshares ecosystem](https://www.payshares.org/developers/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Payshares.org developer Scott Fleckenstein:

[![Horizon: API webserver for the Payshares network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the Payshares network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the Payshares network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Payshares SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.payshares.org/developers/js-payshares-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon-testnet.payshares.org/](https://horizon-testnet.payshares.org/) and one that is connected to the public Payshares network:
[https://horizon.payshares.org/](https://horizon.payshares.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/payshares/js-payshares-sdk)
- [Java](https://github.com/payshares/java-payshares-sdk)
- [Go](https://github.com/payshares/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/payshares/ruby-payshares-sdk)
- [Python](https://github.com/PaysharesCN/py-payshares-base)
- [C#](https://github.com/QuantozTechnology/csharp-payshares-base)
