
const BASE_URL = "http://localhost:3005/api/v1"; // URL harus sesuai dengan backend Go

window.onload = () => {
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    
    if (token) {
        localStorage.setItem('jwt_token', token);
        window.history.replaceState({}, document.title, window.location.pathname);
    }
    
    checkAuth();
};

function checkAuth() {
    const token = localStorage.getItem('jwt_token');
    const path = window.location.pathname;

    if (!token) {
        if (path.includes('index.html') || path.endsWith('/UI/')) {
            window.location.href = 'login.html';
        }
    } else {
        if (path.includes('login.html') || path.includes('register.html')) {
            window.location.href = 'index.html';
        } else if (path.includes('index.html')) {
            if (typeof loadDocuments === "function") {
                loadDocuments();
            }
        }
    }
}

// Fungsi Logout
function handleLogout() {
    localStorage.removeItem('jwt_token');
    window.location.href = 'login.html';
}