async function handleLogin() {
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    if (!email || !password) {
        alert("Email dan password harus diisi!");
        return;
    }

    try {
        const res = await fetch(`${BASE_URL}/auth/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await res.json();
        if (res.ok) {
            localStorage.setItem('jwt_token', data.token);
            window.location.href = 'index.html'; 
        } else {
            alert("Gagal Login: " + data.error);
        }
    } catch (err) {
        alert("Gagal terhubung ke server backend.");
    }
}

function handleGoogleLogin() {
    window.location.href = `${BASE_URL}/auth/google/login`;
}