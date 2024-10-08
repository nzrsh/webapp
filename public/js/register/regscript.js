function goToHome(){
    window.location.href = '/';
}
function goToLog(){
    document.location.href = "/login";
}

async function register(event) {
    event.preventDefault(); // Переместили сюда

    const login = document.getElementById("login").value;
    const password = document.getElementById("password").value;
    const repeatPassword = document.getElementById("repeatPassword").value;


    // Проверка на совпадение паролей
    if (password !== repeatPassword) {
        alert("Пароли не совпадают.");
        return;
    }

    const creds = {
        login: login,
        password: password
    };
    
    console.log("Данные: ", creds);

    // Исправлено название переменной
    const response = await fetch("/auth/register", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(creds) // Исправлено формирование JSON
    });

    // Исправлено: используем 'response' вместо 'responce'
    if (response.status === 201) {
        localStorage.login = creds.login;
        console.log(localStorage.login);
        window.location.href = "/"
    } else if (response.status === 401) {
        const message = await response.text();
        alert(message); // Отображаем сообщение об ошибке
    } else if (response.status === 400) {
        const message = await response.text();
        alert(message); // Отображаем сообщение об ошибке
    }else {
        alert('Ошибка регистрации. Попробуйте позже.');
    }

}
