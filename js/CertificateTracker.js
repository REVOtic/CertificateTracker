var CertificateTrackerContractAddress = "0xf51C5146EDad18cb5a636BF0d563F3f5F287BE6d";

var CertificateTrackerContractABI = [
	{
		"constant": false,
		"inputs": [
			{
				"name": "_ownerAddress",
				"type": "address"
			},
			{
				"name": "_name",
				"type": "string"
			},
			{
				"name": "_email",
				"type": "string"
			},
			{
				"name": "_certificate",
				"type": "string"
			}
		],
		"name": "setCertificateDetails",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "issuerAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "ownerAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "name",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "certificate",
				"type": "string"
			}
		],
		"name": "CertificateIssued",
		"type": "event"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "_issuerAddress",
				"type": "address"
			},
			{
				"name": "_ownerAddress",
				"type": "address"
			}
		],
		"name": "getCertificateHash",
		"outputs": [
			{
				"name": "",
				"type": "string"
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
		"name": "getIssuedCertificateDetails",
		"outputs": [
			{
				"name": "",
				"type": "address"
			},
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "string"
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
				"name": "_certificate",
				"type": "string"
			}
		],
		"name": "getOwnerDetailsByCertificateHash",
		"outputs": [
			{
				"name": "",
				"type": "address"
			},
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]