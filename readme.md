## Design and implement “Word of Wisdom” tcp server.
#### - TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
#### - The choice of the POW algorithm should be explained.
#### - After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
#### - Docker file should be provided both for the server and for the client that solves the POW challenge

## PoW
#### PoW algorithm based on finding a nonce that, when appended to a given challenge string and hashed, produces a hash with a specific number of leading zeroes. The difficulty of this PoW is determined by the number of leading zeroes required in the hash.
- Simplicity and Effectiveness: This algorithm is straightforward to implement and understand. It provides a clear problem (finding the nonce) that requires computational effort to solve, which can be easily adjusted by changing the difficulty level (number of leading zeroes required in the hash).

- Security: SHA-256, the hashing algorithm used, is a cryptographic hash function that is widely regarded as secure. It produces a unique, fixed-size hash (256 bits), making it computationally infeasible to reverse-engineer the nonce or to find a different nonce that produces the same hash.

- Rate Limiting and Anti-Spam: In applications like rate limiting or anti-spam measures, PoW can deter abusive behavior. Since solving the PoW requires computational resources, it can make it costly for an attacker to perform large-scale spam or abuse.

- Fairness in Distribution: In scenarios like cryptocurrency mining, PoW ensures that the chance of solving the hash puzzle is proportional to the computational power expended. This creates a fair competition among participants.

- Network Security in Cryptocurrencies: In blockchain and cryptocurrencies, PoW is used to secure the network. Miners solve complex hash puzzles to validate transactions and add new blocks to the blockchain, making it extremely difficult for a single entity to control or alter the blockchain.

- Preventing Double-Spending: In digital currency systems, PoW helps to prevent double-spending by ensuring that the consensus on the state of the blockchain is maintained through computational work.


## Docker
#### docker-compose up --build