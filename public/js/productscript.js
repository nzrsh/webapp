document.getElementById('queue').onclick = function(event) {
    if (event.target.classList.contains('item')) {
        console.log('DIV нажат');
        var inf = document.getElementById('info');
        inf.style.display = 'block';
    }

};