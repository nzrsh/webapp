function goToHome(){
    window.location.href = '/';
}

document.getElementById("addProductForm").addEventListener("submit", async function(event) {
    event.preventDefault();
    
    const formData = new FormData(this); // Создаем FormData из формы

    try {
        const response = await fetch('/api/products', {
            method: 'POST',
            body: formData // Отправляем FormData
        });

        if (!response.ok) {
            if (response.status == 400) {
                alert("Введите корректные данные!");
                return;
            } else {
            throw new Error('Ошибка сети: ' + response.status);
            }
        }

        const result = await response.json();
        console.log('Продукт создан:', result);
        alert('Продукт успешно создан!');
        document.getElementById("type").value = '';
        document.getElementById("name").value = '';
        document.getElementById("price").value = '';
        document.getElementById("image").value = '';
    } catch (error) {
        console.error('Произошла ошибка:', error);
        alert('Произошла ошибка при создании продукта.');
    }
});