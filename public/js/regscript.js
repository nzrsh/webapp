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

document.getElementById('reg-button').addEventListener('click', async () => {
    // Проверка, что пароли совпадают перед отправкой формы
    if (password.value !== repeatPassword.value) {
        error.style.display = 'block';
        error.textContent = 'Пароли не совпадают';
        return; // Прекращаем выполнение, если пароли не совпадают
    }

    const newProduct = {
        login: login.value,
        password: password.value
    };

    try {
        const response = await fetch('/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newProduct)
        });

        if (!response.ok) {
            throw new Error('Ошибка сети: ' + response.status);
        }

        const result = await response.json();
        console.log(result);
        console.log('Пользователь создан:', result);

        alert('Пользователь успешно зарегистрирован!');
        window.location.href = '/login';
    } catch (error) {
        console.error("Ошибка с сервера:", error.message);
    }
});
