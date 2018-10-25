# sawmint

----

**Saw**tooth + Tender**mint** research and early execution repository

## What is Sawmint?

Sawmint is an attempt to create software that uses Tendermint consensus for Sawtooth applications.  No attempt is made to reproduce SGX functionality at this stage, instead focusing solely on replacing the default PoET consensus with Tendermint Consensus.  

Sawtooth claims to support [unpluggable consensus](https://github.com/danintel/sawtooth-faq/blob/master/consensus.rst).  Sawmint is an attempt at plugging Tendermint into Sawtooth.  

## How about other efforts to do this?

[raft](https://github.com/hyperledger/sawtooth-raft)

[pbft](https://github.com/hyperledger/sawtooth-pbft)

##

Sawtooth Validator configuration:

https://sawtooth.hyperledger.org/docs/core/nightly/master/sysadmin_guide/configuring_sawtooth/validator_configuration_file.html

## Basic Interfaces/diagram


#### Validator Component:

messages between validators use zero mq.  The question is if any application layer comms also go over 0mq.  
http://zeromq.org/

## Running Sawmint Binary
![alt text](https://github.com/c360dev/sawmint/blob/master/inaction.png)

## To Do 
- implement staking mechanism APIs 
- once transaction format is concrete, implement transaction processor parser (1-2 days work)
- dynamic peering (subject to sawtooth validator modification, kyc etc.)
