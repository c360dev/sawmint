# sawmint

----

**Saw**tooth + Tender**mint** research and early execution repository

## What is Sawmint?

Sawmint is an attempt to create software that uses Tendermint consensus for Sawtooth applications.  No attempt is made to reproduce SGX functionality at this stage, instead focusing solely on replacing the default PoET consensus with Tendermint Consensus.  

Sawtooth claims to support [unpluggable consensus](https://github.com/danintel/sawtooth-faq/blob/master/consensus.rst).  Sawmint is an attempt at plugging Tendermint into Sawtooth.  

## How about other efforts to do this?  What's different?

Basically, the difference for us is that we're building on a consensus kit that wasn't pre-baked for Sawtooth so we're going to need to graft some things into place along our journey, whereas the RAFT / PBFT implementations are already built as Sawtooth microservices.  

[raft](https://github.com/hyperledger/sawtooth-raft)

[pbft](https://github.com/hyperledger/sawtooth-pbft)

In our case, we're trying to bring the sawtooth microservices to Tendermint, and then run a Sawtooth app on top of all that.  

## Design

Currently, Sawmint is a Sawtooth-style REST API grafted onto a Tendermint core "node" implementation, with some example transaction processors.  While the REST API is fully featured, it is not connected to tendermint consensus yet.  For that,   

![Sawtooth Diagram](https://github.com/faddat/sawmint/raw/master/sawmint.png)


## Current State:

The Sawtooth REST API has been grafted (via swagger/openapi) onto the standard tendermint executable.   Additionally, protocol buffers files for batches and transactions were generated from .proto to Go.  

The next step is to make the REST API or transaction processor actually do something that makes sense to tendermint and results in some of the following:

* Genesis on a Tendermint chain with a Sawtooth genesis.batch
* Transactions on a tendermint chain

Once we know the two pieces of software are actually talking to one another, we can begin to improve the sawtooth features supported on sawmint.  

## Running Sawmint Binary
![alt text](https://github.com/c360dev/sawmint/blob/master/inaction.png)

## Transaction Processor Support

* Settings TP
* Identity TP
* PoET validator registry TP
* Suspect that this TP will prove either difficult or impossible to implement.  
* IntegerKey TP

Thus far, work towards TP support has involved making Go files out of the appropraite .proto files.  

## To do 
- implement staking mechanism APIs
- once transaction format is concrete, implement transaction processor parser (1-2 days work)
- dynamic peering (subject to sawtooth validator modification, kyc etc.)
- define application specific 
- it's now time to connect the generated code to middleware that will translate the sawtooth calls to tendermint calls.    

## Sawtooth Application Support

Sawmint supports Sawtooth applications by providing them with interfaces that are equivalent to what they'd find in Sawtooth.  
