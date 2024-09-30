mainInf = document.getElementById('info').innerHTML;
var idInfo;

document.getElementById('queue').onclick = function(event) {
    if (event.target.classList.contains('item')) {
        document.getElementById('overlay').style.display = 'block';
        idInfo = event.target.id;
        console.log(idInfo);
        var inf = document.getElementById('info');
        inf.style.display = 'block';
        const sourceDiv = document.getElementById(event.target.id);
        const p = document.createElement('p');
        p.textContent = event.target.name;
        sourceDiv.appendChild(p);
        inf.innerHTML = sourceDiv.innerHTML+ inf.innerHTML;
    }

};



function closeInfo(){
    idInfo = 'none';
    console.log(idInfo);
    document.getElementById('info').innerHTML = mainInf;
    document.getElementById('overlay').style.display = 'none';
    document.getElementById('info').style.display = 'none';
}



// Убедитесь, что idInfo правильно определён перед вызовом этой функции
async function sendData(event) {
    event.preventDefault(); // Остановка стандартного поведения формы

    const formData = new FormData(updateForm); // Получаем данные формы
    console.log(formData)

    try {
        const response = await fetch(`/api/products/${idInfo}`, {
            method: 'POST',
            body: formData // Отправляем FormData
        });

        if (!response.ok) {
            if (response.status === 400) {
                alert("Введите корректные данные!");
                return;
            } else {
                throw new Error('Ошибка сети: ' + response.status);
            }
        }

        alert('Продукт успешно обновлён!');

        // Сброс формы
        document.getElementById("update-name").value = '';
        document.getElementById("update-price").value = '';
        document.getElementById("update-type").value = '';
        document.getElementById("update-image").value = '';
        window.location.reload();
        return
    } catch (error) {
        console.error('Произошла ошибка:', error);
        alert('Произошла ошибка при обновлении продукта.');
    }
};

async function deleteProduct() {
    const productId = idInfo;
    try {
        const response = await fetch(`/api/products/${productId}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            if (response.status == 400) {
                const message = await response.text();
                alert('Ошибка удаления: ' + message);
                return;
            } 
            if (response.status == 500) {
                alert('Внутренняя ошибка сервера');
                return;
            }
        }
        closeInfo();
        alert('Продукт успешно удален!');
        window.location.reload();
    } catch (error) {
            throw new Error('Ошибка при получении данных с сервера');
    }
  
}