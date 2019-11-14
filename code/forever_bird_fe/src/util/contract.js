export const CONTRACT_HASH = '0x010a6bdc5d54dd97a5e2f57425229392d5ab0a75';

export const ABI = [
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newCEOAddress",
          "type": "address"
        }
      ],
      "name": "setCEO",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_adminAddress",
          "type": "address"
        },
        {
          "name": "_newType",
          "type": "uint8"
        }
      ],
      "name": "updateAdmin",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "unpause",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "kill",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_newAdminAddress",
          "type": "address"
        },
        {
          "name": "_type",
          "type": "uint8"
        }
      ],
      "name": "addAdmin",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "paused",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_adminAddress",
          "type": "address"
        }
      ],
      "name": "delAdmin",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "revoked",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_adminAddress",
          "type": "address"
        }
      ],
      "name": "getAdmin",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "pause",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "totalPets",
      "outputs": [
        {
          "name": "",
          "type": "uint64"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "param",
          "type": "string"
        }
      ],
      "name": "FunParams",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "sign",
          "type": "bytes32"
        }
      ],
      "name": "FunSignParams",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "buyerAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "petId",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "currentTime",
          "type": "uint64"
        }
      ],
      "name": "create",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "buyerAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "sellerAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "petId",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "price",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "fee",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "currentTime",
          "type": "uint64"
        }
      ],
      "name": "Trade",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "petId",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "fruitId",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "currentTime",
          "type": "uint64"
        }
      ],
      "name": "feedFruitEvent",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "sender",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "challengerId",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "resisterId",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "isWin",
          "type": "bool"
        },
        {
          "indexed": false,
          "name": "challengerRewardExp",
          "type": "uint32"
        },
        {
          "indexed": false,
          "name": "resisterRewardExp",
          "type": "uint32"
        },
        {
          "indexed": false,
          "name": "winnerRewardCoin",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "attack1",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "attack2",
          "type": "uint64"
        },
        {
          "indexed": false,
          "name": "currentTime",
          "type": "uint64"
        }
      ],
      "name": "PKEvent",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "errCode",
          "type": "int256"
        }
      ],
      "name": "ErrorLog",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "ceoAddress",
          "type": "address"
        }
      ],
      "name": "InitContract",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "originAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "newAddress",
          "type": "address"
        }
      ],
      "name": "AppointNewCEO",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "adminAddress",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "paused",
          "type": "bool"
        }
      ],
      "name": "ContractPaused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "ceoAddr",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "revoked",
          "type": "bool"
        }
      ],
      "name": "ContractRevoked",
      "type": "event"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_petId",
          "type": "uint64"
        }
      ],
      "name": "getPetInfo",
      "outputs": [
        {
          "name": "genes",
          "type": "bytes32"
        },
        {
          "name": "birthTime",
          "type": "uint64"
        },
        {
          "name": "owner",
          "type": "address"
        },
        {
          "name": "coin",
          "type": "uint64"
        },
        {
          "name": "exp",
          "type": "uint32"
        },
        {
          "name": "power",
          "type": "uint32"
        },
        {
          "name": "speed",
          "type": "uint32"
        },
        {
          "name": "level",
          "type": "uint32"
        },
        {
          "name": "eatFruitTime",
          "type": "uint64"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_ownerAddress",
          "type": "address"
        }
      ],
      "name": "getPetCoin",
      "outputs": [
        {
          "name": "",
          "type": "uint64"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_petId",
          "type": "uint64"
        }
      ],
      "name": "getPetCoin",
      "outputs": [
        {
          "name": "coin",
          "type": "uint64"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_ownerAddress",
          "type": "address"
        }
      ],
      "name": "getOwnedPetsInfo",
      "outputs": [
        {
          "name": "",
          "type": "uint256[]"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "tokenId",
          "type": "uint64"
        },
        {
          "name": "price",
          "type": "uint256"
        },
        {
          "name": "fee",
          "type": "uint256"
        },
        {
          "name": "sign",
          "type": "bytes32"
        }
      ],
      "name": "trade",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "catchBird",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "challengerId",
          "type": "uint64"
        },
        {
          "name": "resisterId",
          "type": "uint64"
        },
        {
          "name": "sign",
          "type": "bytes32"
        }
      ],
      "name": "pk",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "tokenId",
          "type": "uint64"
        },
        {
          "name": "fruitId",
          "type": "uint8"
        },
        {
          "name": "sign",
          "type": "bytes32"
        }
      ],
      "name": "feedFruit",
      "outputs": [],
      "payable": true,
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getCEO",
      "outputs": [
        {
          "name": "ceo",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    }
  ]