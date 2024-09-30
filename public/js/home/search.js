const input = document.getElementById('searchInput');

input.addEventListener('input', function() {
    const searchTerm = this.value.trim();
    clearHighlights(document.body);
    if (searchTerm) {
        highlightText(document.body, searchTerm);
    }
});

function clearHighlights(node) {
    const highlights = node.querySelectorAll('.highlight');
    highlights.forEach(highlight => {
        const parent = highlight.parentNode;
        parent.replaceChild(document.createTextNode(highlight.textContent), highlight);
    });
}

function highlightText(node, searchTerm) {
    if (node.nodeType === Node.TEXT_NODE) {
        const text = node.nodeValue;
        const regex = new RegExp(`(${searchTerm})`, 'gi'); // Поиск всех вхождений
        let match;
        let lastIndex = 0;
        const fragments = [];

        while ((match = regex.exec(text)) !== null) {
            // Добавляем текст до совпадения
            if (lastIndex < match.index) {
                fragments.push(document.createTextNode(text.slice(lastIndex, match.index)));
            }
            // Добавляем выделенное совпадение
            const highlightSpan = document.createElement('span');
            highlightSpan.className = 'highlight';
            highlightSpan.textContent = match[0];
            fragments.push(highlightSpan);
            lastIndex = match.index + match[0].length;
        }

        // Добавляем оставшийся текст после последнего совпадения
        if (lastIndex < text.length) {
            fragments.push(document.createTextNode(text.slice(lastIndex)));
        }

        const parent = node.parentNode;
        // Заменяем текстовый узел на фрагменты
        fragments.forEach(fragment => parent.insertBefore(fragment, node));
        parent.removeChild(node);
    } else if (node.nodeType === Node.ELEMENT_NODE) {
        const childNodes = Array.from(node.childNodes);
        childNodes.forEach(child => highlightText(child, searchTerm));
    }
}