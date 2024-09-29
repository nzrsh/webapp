mainInf = document.getElementById('info').innerHTML;

document.getElementById('queue').onclick = function(event) {
    if (event.target.classList.contains('item')) {
        document.getElementById('overlay').style.display = 'block';
        console.log('DIV нажат');
        var inf = document.getElementById('info');
        inf.style.display = 'block';
        const sourceDiv = document.getElementById(event.target.id);
        const targetDiv = document.getElementById('info');
        inf.innerHTML = sourceDiv.innerHTML+ inf.innerHTML;
    }

};

function closeInfo(){
    document.getElementById('info').innerHTML = mainInf;
    document.getElementById('overlay').style.display = 'none';
    document.getElementById('info').style.display = 'none';
}

async function deleteProduct() {
    const productId = event.target.dataset.productId;
    const response = await fetch(`/api/products/${productId}`, {
        method: 'DELETE',
    });
    
}