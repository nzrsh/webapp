document.getElementById('queue').onclick = function(event) {
    if (event.target.classList.contains('item')) {
        console.log('DIV нажат');
        getElementById('overlay').style.display = 'block';
        var inf = document.getElementById('info');
        inf.style.display = 'block';
    }

};