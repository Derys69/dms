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
                            <h4 class="font-bold text-lg">${doc.Title || doc.title}</h4>
                            <p class="text-sm text-gray-500">Departemen: ${doc.Department || doc.department}</p>
                        </div>
                        <button onclick="deleteDoc(${doc.ID || doc.id})" class="text-red-500 font-bold hover:text-red-700 bg-red-100 px-3 py-1 rounded">Hapus</button>
                    </div>
                `;
            });
        } else {
            list.innerHTML = "<p class='text-gray-500 text-center py-4'>Belum ada dokumen yang tersedia.</p>";
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
    if (!id) {
        alert("ID Dokumen tidak valid");
        return;
    }

    const token = localStorage.getItem('jwt_token');
    if (!confirm("Apakah Anda yakin ingin menghapus dokumen ini?")) return;

    try {
        const res = await fetch(`${BASE_URL}/documents/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': `Bearer ${token}` }
        });
        
        if (!res.ok) {
            const data = await res.json();
            alert("Ditolak: " + data.error);
        } else {
            loadDocuments();
        }
    } catch (err) {
        alert("Gagal menghapus dokumen, periksa koneksi server.");
    }
}