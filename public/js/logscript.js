function goToReg(){
    document.location.href = "/register";;
}

async function login(event) {
    event.preventDefault();
    console.log("+");

    const login = document.getElementById("login").value;
    const password = document.getElementById("password").value;


    const creds = {
        login: login,
        password: password
    };
    
    console.log("Данные: ", creds);

    // Исправлено название переменной
    const response = await fetch("/auth/login", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(creds) // Исправлено формирование JSON
    });

    // Исправлено: используем 'response' вместо 'responce'
    if (response.status === 201) {
        console.log('response')
        localStorage.getItem('response')
        window.location.href = "/"
    } else if (response.status === 401) {
        const message = await response.text();
        alert(message); 
    } else {
        alert('Ошибка Авторизации. Попробуйте позже.');
    }

}