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
        p.textContent = 'Тип продукта: ' + event.target.name;
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