

document.getElementById('createProductButton').addEventListener('click', async () => {
    const type = document.getElementById('productType').value;
    const name = document.getElementById('productName').value;
    const price = parseFloat(document.getElementById('productPrice').value);

    if (!type || !name || isNaN(price)) {
        alert('Пожалуйста, заполните все поля корректно.');
        return;
    }

    const newProduct = {
        type:type,
        name: name,
        price: price
    };

    try {
        const response = await fetch('/api/products', {
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
        console.log('Продукт создан:', result);
        alert('Продукт успешно создан!');
    } catch (error) {
        console.log('Произошла ошибка:', error);
        alert('Произошла ошибка при создании продукта.');
    }
});
