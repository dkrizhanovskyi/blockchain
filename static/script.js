document.getElementById('load-blockchain').addEventListener('click', loadBlockchain);
document.getElementById('load-nodes').addEventListener('click', loadNodes);
document.getElementById('send-transaction-form').addEventListener('submit', sendTransaction);
document.getElementById('mine-block-form').addEventListener('submit', mineBlock);
document.getElementById('register-form').addEventListener('submit', registerUser);
document.getElementById('login-form').addEventListener('submit', loginUser);

function loadBlockchain() {
    fetch('/blockchain')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            const blockchainDiv = document.getElementById('blockchain');
            blockchainDiv.innerHTML = '';
            data.forEach(block => {
                const blockDiv = document.createElement('div');
                blockDiv.classList.add('block');
                blockDiv.innerHTML = `
                    <p><strong>Index:</strong> ${block.Index}</p>
                    <p><strong>Timestamp:</strong> ${block.Timestamp}</p>
                    <p><strong>Transactions:</strong> ${JSON.stringify(block.Transactions)}</p>
                    <p><strong>Hash:</strong> ${block.Hash}</p>
                    <p><strong>Previous Hash:</strong> ${block.PrevHash}</p>
                `;
                blockchainDiv.appendChild(blockDiv);
            });
        })
        .catch(error => {
            console.error('Error loading blockchain:', error);
            showNotification('Error loading blockchain', true);
        });
}

function loadNodes() {
    fetch('/nodes')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            const nodesDiv = document.getElementById('nodes');
            nodesDiv.innerHTML = '';
            data.forEach(node => {
                const nodeDiv = document.createElement('div');
                nodeDiv.classList.add('node');
                nodeDiv.innerHTML = `
                    <p><strong>Node Address:</strong> ${node.Address}</p>
                `;
                nodesDiv.appendChild(nodeDiv);
            });
        })
        .catch(error => {
            console.error('Error loading nodes:', error);
            showNotification('Error loading nodes', true);
        });
}

function sendTransaction(event) {
    event.preventDefault();

    const sender = document.getElementById('sender').value;
    const recipient = document.getElementById('recipient').value;
    const amount = parseInt(document.getElementById('amount').value);

    const transaction = { Sender: sender, Recipient: recipient, Amount: amount };

    fetch('/send', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(transaction),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Transaction sent:', data);
        showNotification('Transaction sent successfully', false);
    })
    .catch(error => {
        console.error('Error sending transaction:', error);
        showNotification('Error sending transaction', true);
    });
}

function mineBlock(event) {
    event.preventDefault();

    const recipient = document.getElementById('recipient-address').value;
    const reward = parseInt(document.getElementById('mining-reward').value);

    const mineRequest = { recipient: recipient, amount: reward };

    fetch('/mine', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(mineRequest),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Block mined:', data);
        showNotification('Block mined successfully', false);
    })
    .catch(error => {
        console.error('Error mining block:', error);
        showNotification('Error mining block', true);
    });
}

function registerUser(event) {
    event.preventDefault();

    const username = document.getElementById('register-username').value;
    const password = document.getElementById('register-password').value;

    const registerRequest = { username: username, password: password };

    fetch('/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(registerRequest),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('User registered:', data);
        showNotification('User registered successfully', false);
    })
    .catch(error => {
        console.error('Error registering user:', error);
        showNotification('Error registering user', true);
    });
}

function loginUser(event) {
    event.preventDefault();

    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;

    const loginRequest = { username: username, password: password };

    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(loginRequest),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        console.log('Login successful:', data);
        showNotification('Login successful', false);
    })
    .catch(error => {
        console.error('Error logging in:', error);
        showNotification('Error logging in', true);
    });
}

function showNotification(message, isError) {
    const notification = document.getElementById('notification');
    notification.textContent = message;
    notification.className = isError ? 'notification error' : 'notification';
    notification.style.display = 'block';
    setTimeout(() => {
        notification.style.display = 'none';
    }, 3000);
}

setInterval(loadBlockchain, 10000); // Автоматическое обновление каждые 10 секунд
