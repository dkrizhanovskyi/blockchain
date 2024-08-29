document.getElementById('get-blockchain').addEventListener('click', getBlockchain);
document.getElementById('add-block').addEventListener('click', addBlock);
document.getElementById('validate-blockchain').addEventListener('click', validateBlockchain);

function getBlockchain() {
    fetch('/getblockchain')
        .then(response => response.json())
        .then(data => {
            displayOutput(JSON.stringify(data, null, 2));
        })
        .catch(error => {
            displayOutput('Error fetching blockchain: ' + error);
        });
}

function addBlock() {
    const blockData = prompt('Enter data for the new block:');
    if (blockData) {
        fetch('/addblock', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ data: blockData })
        })
        .then(response => response.json())
        .then(data => {
            displayOutput(data.message);
            getBlockchain(); // Refresh the blockchain view
        })
        .catch(error => {
            displayOutput('Error adding block: ' + error);
        });
    }
}

function validateBlockchain() {
    fetch('/validate')
        .then(response => response.json())
        .then(data => {
            displayOutput(data.message);
        })
        .catch(error => {
            displayOutput('Error validating blockchain: ' + error);
        });
}

function displayOutput(message) {
    const output = document.getElementById('output');
    output.value = message;
}
