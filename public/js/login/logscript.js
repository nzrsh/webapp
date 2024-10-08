function goToHome(){
    window.location.href = '/';
}
function goToReg(){
    document.location.href = "/register";;
}

async function auth(event) {
    event.preventDefault();

    const username = document.getElementById("login").value; 
    const password = document.getElementById("password").value;

    const creds = {
        login: username,
        password: password
    };
    
    console.log("Данные: ", creds);

    const response = await fetch("/auth/login", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(creds)
    });

    if (response.status === 200) {
        window.location.href = "/";
    } else if (response.status === 400) {
        const message = await response.text();
        alert(message);
    } else if (response.status === 401) {
        const message = await response.text();
        alert(message);  
    } else {
        alert('Ошибка Авторизации. Попробуйте позже.');
    }
}
