const queueElement = document.getElementById('queue');
const addButton = document.getElementById('addButton');

addButton.addEventListener('click', () => {
    queueElement.innerHTML = ''; // Очищаем текущий вывод
    queue.forEach(item => {
        const div = document.createElement('div');
        div.className = 'item';
        div.textContent = item;
        queueElement.appendChild(div);
    });
});