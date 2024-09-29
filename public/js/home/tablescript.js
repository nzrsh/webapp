const queueElement = document.getElementById('queue');
const addButton = document.getElementById('addButton');

//Получить продукты
document.addEventListener('DOMContentLoaded', fetchProducts);
//Получить данные пользователя
document.addEventListener('DOMContentLoaded', fetchUser);



//Отрисовать продукт
function productManager(items) {
    items.forEach(item => {
        const div = document.createElement('div');
        div.id = item.id;
        div.className = 'item';
        div.name = item.type;
        const img = document.createElement('img');
        img.src = 'api/img/' + item.id;
        img.width = 300;
        img.height = 200
        div.appendChild(img);
        div.textContent = item.name;
        const p = document.createElement('p')
        p.textContent = item.price
        div.appendChild(p)
        queueElement.appendChild(div);
    });
}

//Кнопка "Войти"
function DisplayLoginbutton() {
    const container = document.getElementById('button-container');
    container.innerHTML = `
        <button id="loginButton">Войти</button>
    `;
    document.getElementById('loginButton').onclick = () => {
        window.location.href = '/login'; // Перенаправляем на страницу логина
    };
}


function displayLogoutButton() {
    const container = document.getElementById('button-container');
    container.innerHTML = `
        <button id="logoutButton">Выйти</button>
    `;
    document.getElementById('logoutButton').onclick = async () => {
        const response = await fetch('/auth/logout', {
            method: 'POST',
        });

        if (response.ok) {
            alert('Вы успешно вышли.');
            window.location.reload()
        } else {
            alert('Ошибка при выходе.');
        }
    };
}


//Запрос получения продуктов
async function fetchProducts() {
    try {
        const response = await fetch('/api/products');
        if (!response.ok) {
            throw new Error('Сеть ответила с ошибкой: ' + response.status);
        }
        const products = await response.json(); // Преобразуем ответ в JSON
        productManager(products); // Выводим полученные продукты в консоль
    } catch (error) {
        console.error('Произошла ошибка:', error);
    }
}

//Запрос получения данных о пользователе
async function fetchUser() {
    const userLabel = document.getElementById("user-label")
    const StatusUnauthorized = false;
    try {
        const response = await fetch('/auth/me', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        });

        if (!response.ok) {
            if (response.status == 401) {
                userLabel.textContent = "Гость";
                const StatusUnauthorized = true;
                DisplayLoginbutton();
            } else {
                throw new Error('Ошибка при получении данных с сервера');
            }
        }

        const data = await response.json();
        console.log('Полученные данные:', data);
        userLabel.textContent = data.login;
        displayLogoutButton();
    } catch (error) {
        if (StatusUnauthorized) {
            console.error('Ошибка:', error);
        }
    }
}


/*document.getElementById('addButton').addEventListener('click', function () {
    const inputContainer = document.getElementById('inputContainer');
    inputContainer.style.display = inputContainer.style.display === 'none' ? 'block' : 'none';
});*/
