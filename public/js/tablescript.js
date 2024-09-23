const queue = [1,2,12,3,23,23,2,3,213,2,3,23,2,13,23,23];
const queueElement = document.getElementById('queue');
const addButton = document.getElementById('addButton');

addButton.addEventListener('click', () => {
    queue.forEach(item => {
        const div = document.createElement('div');
        div.className = 'item';
        div.textContent = 'sadasd';
        queueElement.appendChild(div);
        console.log('addButton clicked');
    });
});