// Получение списка файлов с сервера

function goToHome(){
    window.location.href = '/';
}
async function fetchFiles() {
    const response = await fetch('/storage/files', {
        method: 'GET'
    });

    if (response.ok) {
        const files = await response.json();
        displayFiles(files);
    } else {
        alert('Ошибка загрузки файлов');
    }
}

// Отображение списка файлов в таблице
function displayFiles(files) {
    const fileTableBody = document.getElementById('fileTable').querySelector('tbody');
    fileTableBody.innerHTML = ''; // Очищаем таблицу

    files.forEach(file => {
        const row = document.createElement('tr');

        const nameCell = document.createElement('td');
        
        // Создаём контейнер для изображения и заголовка
        const fileContainer = document.createElement('div');
        fileContainer.style.textAlign = 'center'; // Выравнивание по центру

        if (file.isImage) {
            const img = document.createElement('img');
            img.src = `/storage/files/image?filename=${encodeURIComponent(file.name)}`; 
            img.alt = file.name;
            img.width = 50;
            fileContainer.appendChild(img); // Добавляем изображение в контейнер

            // Создаём заголовок
            const title = document.createElement('div');
            title.textContent = file.name; // Устанавливаем текст заголовка
            title.style.marginTop = '5px'; // Отступ сверху для заголовка
            fileContainer.appendChild(title); // Добавляем заголовок в контейнер
        } else {
            // Если это не изображение, просто добавляем имя файла
            fileContainer.textContent = file.name; 
        }

        nameCell.appendChild(fileContainer); // Добавляем контейнер в ячейку
        row.appendChild(nameCell);

        const sizeCell = document.createElement('td');
        sizeCell.textContent = `${(file.size / 1024).toFixed(2)} KB`;

        const modTimeCell = document.createElement('td');
        modTimeCell.textContent = new Date(file.modTime).toLocaleString();

        const actionsCell = document.createElement('td');

        // Кнопка для удаления файла
        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Удалить';
        deleteButton.onclick = () => deleteFile(file.name);
        actionsCell.appendChild(deleteButton);

        // Кнопка для переименования файла
        const renameButton = document.createElement('button');
        renameButton.textContent = 'Переименовать';
        renameButton.onclick = () => renameFile(file.name);
        actionsCell.appendChild(renameButton);

        row.appendChild(nameCell);
        row.appendChild(sizeCell);
        row.appendChild(modTimeCell);
        row.appendChild(actionsCell);

        fileTableBody.appendChild(row);
    });
}



// Загрузка файлов
document.getElementById('uploadForm').addEventListener('submit', async (event) => {
    event.preventDefault();

    const files = document.getElementById('fileInput').files;
    if (files.length === 0) {
        alert('Выберите хотя бы один файл');
        return;
    }

    const formData = new FormData();
    for (let file of files) {
        formData.append('file', file);
    }

    const response = await fetch('/storage/files/upload', {
        method: 'POST',
        body: formData // Токен из куки будет отправлен автоматически
    });

    if (response.ok) {
        fetchFiles(); // Обновить список файлов
    } else {
        alert('Ошибка загрузки файлов');
    }
});

// Удаление файла
async function deleteFile(filename) {
    const response = await fetch(`/storage/files/${filename}`, {
        method: 'DELETE'
    });

    if (response.ok) {
        fetchFiles(); // Обновить список файлов
    } else {
        alert('Ошибка удаления файла');
    }
}

// Переименование файла
async function renameFile(oldName) {
    const newName = prompt('Введите новое имя файла', oldName);
    if (!newName) return;

    const response = await fetch(`/storage/files/${oldName}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ newName })
    });

    if (response.ok) {
        fetchFiles(); // Обновить список файлов
    } else {
        alert('Ошибка переименования файла');
    }
}

// Загрузка списка файлов при загрузке страницы
document.addEventListener('DOMContentLoaded', fetchFiles);
