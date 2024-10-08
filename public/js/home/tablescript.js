const queueElement = document.getElementById('queue');
const addButton = document.getElementById('addButton');

document.addEventListener('DOMContentLoaded', fetchProducts);
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
        img.width = 150;
        img.height = 150;
        div.appendChild(img); // Добавляем изображение

        const p = document.createElement('p');
        p.textContent = item.name; // Добавляем имя продукта
        div.appendChild(p);

        const price = document.createElement('p');
        price.textContent = item.price; // Добавляем цену
        div.appendChild(price);

        queueElement.appendChild(div); // Добавляем весь div в контейнер
    });
}

//Кнопка "Войти"
function DisplayLoginbutton() {
    const container = document.getElementById('button-container');
    container.innerHTML = `
        <button id="loginButton">Войти</button>
    `;
    document.getElementById('loginButton').onclick = () => {
        goTo('/login') // Перенаправляем на страницу логина
    };
}


function displayLogoutButton() {
    const container = document.getElementById('button-container');
    container.innerHTML = `
        <button id="logoutButton">Выйти</button>
        <button id="storageButton">Хранилище</button>
        <button id="addButton">Добавление</button>

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
    document.getElementById('storageButton').onclick = async () => {
        goTo('/storage');
    };
    document.getElementById('addButton').onclick = async () => {
        goTo('/add')
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
