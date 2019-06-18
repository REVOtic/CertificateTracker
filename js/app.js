const IPFS = IpfsApi;
const ipfs = new IPFS({ host: 'ipfs.infura.io', port: 5001, protocol: 'https' });

let fileIPFSHash = "";

// Inject our version (1.0.0-beta.55) of web3.js into the DApp.
window.addEventListener('load', async () => {
    // Modern dapp browsers...
    if (window.ethereum) {
        window.web3 = new Web3(ethereum);
        try {
            // Request account access if needed
            await ethereum.enable();
            // Acccounts now exposed
            console.log('Using Web3 Version:', web3.version);
            if(web3._provider.isMetaMask) {
                // Check if Metamask is available
                console.log('Got MetaMask! Hurray');
                startApp();
            }
        } catch (error) {
            // User denied account access...
            console.log('Please Use Metamask extension to transact in ETH.');
            document.getElementById('meta-mask-required').innerHTML = 'You need <a href="https://metamask.io/">MetaMask</a> browser plugin to run this example'
        }
    }
})

function checkNetwork() {
    web3.eth.net.getNetworkType((err, netId) => {
        switch (netId) {
            case "main":    console.log('The Mainnet');
                            break
            case "ropsten": console.log('Ropsten Test Network');
                            break
            case "kovan":   console.log('Kovan Test Network');
                            break
            case "rinkeby": console.log('Rinkeby Test Network');
                            break
            default:        console.log('This is an Unknown Network');
        }
        if (netId != "rinkeby") {
            alert("Please connect to Rinkeby Testnet to Continue!");
        } else {
            console.log("Connected to Rinkeby Testnet!");
        }
    });
}

function initDApp() {
    window.CertificateTrackerContract = new web3.eth.Contract(CertificateTrackerContractABI, CertificateTrackerContractAddress, {
        from: walletAddress,
        gasPrice: '20000000000'                 // default gas price in wei, 20 gwei in this case
    });
}

function getETHBalance() {
    return new Promise(resolve => {
        web3.eth.getBalance(walletAddress, (error, result) => {
            if (!error) {
                console.log(web3.utils.fromWei(result, "ether"));
                resolve(web3.utils.fromWei(result, "ether"));
            } else {
                resolve(error);
            }
        });
    });
}

function getCoinbase() {
    return new Promise(resolve => {
        web3.eth.getCoinbase((error, result) => {
            if (!error) {
                console.log("Coinbase: "+result);
                resolve(result);
            } else {
                resolve(error);
            }
        });
    });
}

async function fetchAccountDetails() {
    // Fetch the Account Details
    window.walletAddress = await getCoinbase();
    document.getElementById('accountAddress').innerHTML = walletAddress;
    let walletETHBalance = await getETHBalance();
    document.getElementById('etherBalance').innerHTML = walletETHBalance;
}

async function startApp() {
    // Check For correct Network
    await checkNetwork();
    await fetchAccountDetails();
    await initDApp();
}

function chooseFile() {
    // Select the File and Convert to ReadableStream
    const file = event.target.files[0];
    let instanceOfFileReader = new window.FileReader();
    instanceOfFileReader.readAsArrayBuffer(file);

    instanceOfFileReader.onloadend = async (error, result) => {
        // console.log(error, result);
        // file is converted to a buffer to prepare for uploading to IPFS
        var Buffer = buffer.Buffer;
        const fileBuffer = Buffer.from(instanceOfFileReader.result);
        await ipfs.files.add(fileBuffer, (error, ipfsHash) => {
            if(!error) {
                fileIPFSHash = ipfsHash[0].hash;
                console.log("Upload to IPFS Succssful. Hash: " +fileIPFSHash);
            }
            else {
                console.log("Error in Uploading to IPFS: " +error);
            }
        });
    }
}

function issueCertificateOnETH(ownerName, ownerAddress, ownerEmail, ownerCertificate) {
    return new Promise(resolve => {
        CertificateTrackerContract.methods.setCertificateDetails(ownerAddress, ownerName, ownerEmail, ownerCertificate).send((error, result) => {
            if (!error) {
                console.log(result);
                resolve(result);
            } else {
                resolve(error);
            }
        });
    });
}

async function issueCertificate() {
    let ownerName = document.getElementById('ownerName').value;
    let ownerAddress = document.getElementById('ownerAddress').value;
    let ownerEmail = document.getElementById('ownerEmail').value;
    let ownerCertificate = fileIPFSHash;
    console.log(ownerName, ownerAddress, ownerEmail, ownerCertificate);
    let submitResponseTxHash = await issueCertificateOnETH(ownerName, ownerAddress, ownerEmail, ownerCertificate);
    console.log("Transaction Hash: "+submitResponseTxHash);
    let submitResponseURL = "https://rinkeby.etherscan.io/tx/" + submitResponseTxHash;
    document.getElementById('issueCertificateResponse').innerHTML = 'Check Tx at: <a>'+submitResponseURL+'</a>';
}