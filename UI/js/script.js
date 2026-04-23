const BASE_URL = "http://localhost:3005/api/v1";

window.onload = () => {
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    
    if (token) {
        localStorage.setItem('jwt_token', token);
        window.history.replaceState({}, document.title, window.location.pathname);
    }
    checkAuth();
};

function toggleAuth(isLogin) {
    document.getElementById('login-form').classList.toggle('hidden', !isLogin);
    document.getElementById('register-form').classList.toggle('hidden', isLogin);
    document.getElementById('auth-title').innerText = isLogin ? "DMS Login" : "DMS Register";
}

function checkAuth() {
    const token = localStorage.getItem('jwt_token');
    if (token) {
        document.getElementById('auth-section').classList.add('hidden');
        document.getElementById('dashboard-section').classList.remove('hidden');
        loadDocuments();
    }
}

async function handleLogin() {
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    try {
        const res = await fetch(`${BASE_URL}/auth/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await res.json();
        if (res.ok) {
            localStorage.setItem('jwt_token', data.token);
            checkAuth();
        } else {
            alert(data.error);
        }
    } catch (err) {
        alert("Gagal terhubung ke server.");
    }
}

async function handleRegister() {
    const name = document.getElementById('reg-name').value;
    const email = document.getElementById('reg-email').value;
    const password = document.getElementById('reg-password').value;

    if (!name || !email || !password) {
        alert("Harap isi semua kolom!");
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
            alert(data.message);
            toggleAuth(true); 
        } else {
            alert("Gagal Daftar: " + data.error);
        }
    } catch (err) {
        alert("Tidak dapat terhubung ke server backend.");
    }
}

function handleGoogleLogin() {
    window.location.href = `${BASE_URL}/auth/google/login`;
}

function handleLogout() {
    localStorage.removeItem('jwt_token');
    window.location.href = "/UI/index.html";
}

async function loadDocuments() {
    const token = localStorage.getItem('jwt_token');
    try {
        const res = await fetch(`${BASE_URL}/documents/`, {
            headers: { 'Authorization': `Bearer ${token}` }
        });
        
        const data = await res.json();
        const list = document.getElementById('doc-list');
        list.innerHTML = "";
        
        if (data.data && data.data.length > 0) {
            data.data.forEach(doc => {
                list.innerHTML += `
                    <div class="border-b pb-2 flex justify-between items-center">
                        <div>
                            <h4 class="font-bold">${doc.title}</h4>
                            <p class="text-sm text-gray-500">Dept: ${doc.department}</p>
                        </div>
                        <button onclick="deleteDoc(${doc.id})" class="text-red-500 font-bold hover:text-red-700">Hapus</button>
                    </div>
                `;
            });
        } else {
            list.innerHTML = "<p class='text-gray-500'>Belum ada dokumen.</p>";
        }
    } catch (err) {
        console.error("Gagal meload dokumen", err);
    }
}

async function createDoc() {
    const titleInput = document.getElementById('doc-title');
    const contentInput = document.getElementById('doc-content');
    const token = localStorage.getItem('jwt_token');

    if (!titleInput.value || !contentInput.value) {
        alert("Judul dan isi dokumen tidak boleh kosong!");
        return;
    }

    try {
        const res = await fetch(`${BASE_URL}/documents/`, {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({ 
                title: titleInput.value, 
                content: contentInput.value 
            })
        });

        const result = await res.json();

        if (res.ok) {
            alert("Mantap! Dokumen tersimpan.");
            titleInput.value = "";
            contentInput.value = "";
            loadDocuments();
        } else {
            alert("Gagal: " + (result.error || "Akses ditolak"));
        }
    } catch (err) {
        alert("Server tidak merespon. Pastikan backend Go sedang berjalan.");
    }
}

async function deleteDoc(id) {
    const token = localStorage.getItem('jwt_token');
    if (!confirm("Yakin ingin menghapus dokumen ini?")) return;

    try {
        const res = await fetch(`${BASE_URL}/documents/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': `Bearer ${token}` }
        });
        
        if (!res.ok) {
            const data = await res.json();
            alert("Error: " + data.error);
        } else {
            loadDocuments();
        }
    } catch (err) {
        alert("Gagal menghapus, cek koneksi server.");
    }
}