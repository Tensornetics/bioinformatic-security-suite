// Import necessary modules
const crypto = window.crypto;
const subtle = crypto.subtle;

// Define the biometric auth check
async function biometricAuth() {
    const authCredential = await navigator.credentials.create({
        publicKey: {
            rpId: 'localhost',
            user: {
                id: new Uint8Array([1, 2, 3]),
                name: 'John Doe',
                displayName: 'John',
            },
            challenge: new Uint8Array([4, 5, 6]),
            pubKeyCredParams: [{
                type: 'public-key',
                alg: -7,
            }],
        },
    });

    // Handle authCredential response
    if (!authCredential) {
        console.log('Biometric authentication failed or denied.');
        return;
    }

    console.log('Biometric authentication successful!');
    console.log(authCredential);
    // Do something with the authCredential response
    // like sending it to the server for verification
}

// Register the event listener
document.querySelector('#biometric-button').addEventListener('click', biometricAuth);
