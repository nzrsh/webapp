const input = document.getElementById('searchInput');

input.addEventListener('input', function() {
    const searchTerm = this.value.trim();
    clearHighlights(document.body); // Очистка предыдущих выделений
    if (searchTerm) {
        highlightText(document.body, searchTerm); // Поиск и выделение нового текста
    }
});

function clearHighlights(node) {
    // Ищем все элементы с классом "highlight" и заменяем их на текст
    const highlights = node.querySelectorAll('.highlight');
    highlights.forEach(highlight => {
        const parent = highlight.parentNode;
        // Создаем текстовый узел с оригинальным текстом
        parent.replaceChild(document.createTextNode(highlight.textContent), highlight);
        parent.normalize(); // Объединяем смежные текстовые узлы для чистоты DOM
    });
}

function highlightText(node, searchTerm) {
    if (node.nodeType === Node.TEXT_NODE) {
        const text = node.nodeValue;
        const regex = new RegExp(`(${searchTerm})`, 'gi'); // Регулярное выражение для поиска

        // Проверяем, есть ли совпадения в тексте
        if (regex.test(text)) {
            const parent = node.parentNode;
            const matches = text.split(regex); // Разбиваем текст на части: совпадения и не-совпадения

            // Создаем документ-фрагмент, чтобы быстро заменить текст узла
            const fragment = document.createDocumentFragment();

            matches.forEach(part => {
                if (regex.test(part)) {
                    // Совпадающая часть: создаем элемент span с выделением
                    const highlightSpan = document.createElement('span');
                    highlightSpan.className = 'highlight';
                    highlightSpan.style.backgroundColor = 'yellow';
                    highlightSpan.textContent = part;
                    fragment.appendChild(highlightSpan);
                } else {
                    // Обычный текст, не совпадающий с поиском
                    fragment.appendChild(document.createTextNode(part));
                }
            });

            // Заменяем старый текстовый узел на новый фрагмент
            parent.replaceChild(fragment, node);
        }
    } else if (node.nodeType === Node.ELEMENT_NODE && node.tagName !== 'SCRIPT' && node.tagName !== 'STYLE') {
        // Рекурсивно обходим дочерние узлы, кроме скриптов и стилей
        node.childNodes.forEach(child => highlightText(child, searchTerm));
    }
}