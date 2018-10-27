# Notes
This file consists of notes that aren't yet in any order.  It's got cribs so we can DRY and links to the most important information on the topic.  But it's certainly not yet in any kind of order.  

## Sawtooth is a set of systemd services by default:

* Rest API
* Transaction processors
* Validator Registry

## Creating a sawtooth genesis block

To create a sawtooth genesis block you make keys, and then you turn those keys into a genesis block with a set of commands.  Then you're left with a genesis.batch file AND genesis batch is printed to /var/lib/sawtooth/genesis.batch after it is generated. 

## How to actually deploy the generated batch file as described here: 

https://sawtooth.hyperledger.org/docs/core/nightly/master/sysadmin_guide/setting_up_sawtooth_poet-sim.html
https://sawtooth.hyperledger.org/docs/core/releases/1.0/architecture/journal.html

## Where Sawtooth should meet tendermint and vice versa

So the idea here is to take Sawtooth's notion of an "app" and Pair it with tendermint as the consensus engine.  

**Sawtooth's composition: 40% rust and 60% Python -@danintel from sawtooth chat**

[sawtooth consensus faq](https://github.com/danintel/sawtooth-faq/blob/master/consensus.rst#id21)

Following this guide "how do I change sawtooth consensus" it would seem that we need to prepare a drop-in module that has all of the same interfaces that the current consensus bits have.  

## Batches

"Batches are the atomic unit of state change"

[transaction batching guide](https://sawtooth.hyperledger.org/docs/core/releases/1.0/architecture/transactions_and_batches.html)

Here, it would seem to be necessary to modify Tendermint to support such batches.  

## Sawtooth-pbft

Sawtooth-pbft is composed of a set of docker containers that can launch a pbft sawtooth test network.  I'm not sure that anything has been built on it yet, but I'm trying to find out.  It seems like the consensus mechanism for sawtooth-pbft is written in Rust, but I'm just judging by the fact that it uses the cargo compiler.  I still haven't made an exact determination of how sawtooth components written in different languages / run in different containers talk to one another (though clearly they use protocol buffers.)


Sawtooth-pbft does not compile:

```
error[E0053]: method `start` has an incompatible type for trait
   --> src/engine.rs:41:5
    |
41  | /     fn start(
42  | |         &mut self,
43  | |         updates: Receiver<Update>,
44  | |         mut service: Box<Service>,
...   |
113 | |         }
114 | |     }
    | |_____^ expected enum `std::result::Result`, found ()
    |
    = note: expected type `fn(&mut engine::PbftEngine, std::sync::mpsc::Receiver<sawtooth_sdk::consensus::engine::Update>, std::boxed::Box<(dyn sawtooth_sdk::consensus::service::Service + 'static)>, sawtooth_sdk::consensus::engine::StartupState) -> std::result::Result<(), sawtooth_sdk::consensus::engine::Error>`
               found type `fn(&mut engine::PbftEngine, std::sync::mpsc::Receiver<sawtooth_sdk::consensus::engine::Update>, std::boxed::Box<(dyn sawtooth_sdk::consensus::service::Service + 'static)>, sawtooth_sdk::consensus::engine::StartupState)`

error: aborting due to previous error

For more information about this error, try `rustc --explain E0053`.
error: Could not compile `sawtooth-pbft`.

To learn more, run the command again with --verbose.
```

Maybe sawtooth-pbft requires a higher version of Ubuntu like 18.04 despite all sawtooth reqs stating 16.04?  (Guessing this from the version mismatch in sawtooth-pbft's docker compose file)

## sawtooth-raft
```
Sawtooth Raft is a consensus engine for Hyperledger Sawtooth based on the crash fault tolerant consensus algorithm Raft. Specifically, it is built on top of the Rust implementation of Raft used by TiKV, raft-rs.

Currently, Sawtooth Raft is in the prototype phase of development and there is additional work to be done in order to make it production worthy.
```

Sawtooth-raft requires multiple machines to run, there is no containerized single-machine configuration like pbft has.  

**need to re-run these tests with more than a single server**

## applying lessons from sawtooth-raft and sawtooth-pbft


## Tendermint

Here are the big difficulties I see for moving tendermint over to this design pattern:

* TM transactions are not batched and there's no inbuilt processing for batched transactions.  There may not be any support whatsoever for the `sawtooth batch` type commands.  This is of course surmountable but it's significant.  
* The tendermint light client features a native REST API, but it's incomplete and it's not going to be compatible (conceptually) with Sawtooth's native REST API. 
* If 0mq serves anything other than consensus, then we'll need to bring 0mq into the tendermint application architecture.  
* 


## Full container Loadout for sawtooth:
This command starts a validator with the following components attached to it:

REST API (available on host port 8008)
IntKey transaction processor (Python implementation)
Settings transaction processor
XO transaction processor (Python implementation)
Shell (for running Sawtooth commands)

## Sawtooth Validator configuration:

https://sawtooth.hyperledger.org/docs/core/nightly/master/sysadmin_guide/configuring_sawtooth/validator_configuration_file.html

#### Validator Component:

messages between validators use zero mq.  The question is if any application layer comms also go over 0mq.  
http://zeromq.org/




