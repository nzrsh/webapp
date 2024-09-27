const login = document.getElementById('login');
const password = document.getElementById('password');
const repeatPassword = document.getElementById('repeatPassword');
const error = document.getElementById('error');

repeatPassword.addEventListener('input', function() {
    if (password.value !== repeatPassword.value) {
        error.style.display = 'block';
        error.textContent = 'Пароли не совпадают';
    } else {
        error.style.display = 'none';
    }
});

function regUser(){
    // Проверка, что пароли совпадают перед отправкой формы
    if (password.value !== repeatPassword.value) {
        error.style.display = 'block';
        error.textContent = 'Пароли не совпадают';
        return; // Прекращаем выполнение, если пароли не совпадают
    }

    const Credentials = {
        login: login.value,
        password: password.value
    };

        const response = fetch('/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(Credentials)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка сети: ' + response.status);
            }
            return response.json();
        })
        .then(result => {
            console.log(result);
            alert('Пользователь успешно зарегистрирован!');
        })
        .catch(error => {
            console.error("Ошибка с сервера:", error.message);
        });
};
