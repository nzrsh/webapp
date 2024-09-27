const queueElement = document.getElementById('queue');
const addButton = document.getElementById('addButton');
fetchProducts();

function productManager(items) {
    items.forEach(item => {
        const div = document.createElement('div');
        div.className = 'item';
        div.name = item.type;
        div.textContent = item.name;
        const p = document.createElement('p')
        p.textContent = item.price
        div.appendChild(p)
        queueElement.appendChild(div);
    });
}

function goToReg(){
    window.location.href = '/register';
}

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
document.getElementById('addButton').addEventListener('click', function() {
    const inputContainer = document.getElementById('inputContainer');
    inputContainer.style.display = inputContainer.style.display === 'none' ? 'block' : 'none';
});
