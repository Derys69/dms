async function handleRegister() {
    const name = document.getElementById('reg-name').value;
    const email = document.getElementById('reg-email').value;
    const password = document.getElementById('reg-password').value;

    if (!name || !email || !password) {
        alert("Harap isi semua kolom pendaftaran!");
        return;
    }

    try {
        const res = await fetch(`${BASE_URL}/auth/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password })
        });

        const data = await res.json();

        if (res.ok) {
            alert("Berhasil: " + data.message);
            window.location.href = 'login.html'; 
        } else {
            alert("Gagal Daftar: " + data.error);
        }
    } catch (err) {
        alert("Tidak dapat terhubung ke server backend.");
    }
}