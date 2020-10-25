function getShortenURL(longURL) {
    var body = {
        long_url: longURL
    }

    return fetch('/api', {
        method: 'post',
        body: JSON.stringify(body)
    }).then(function(response) {
        return response.json();
    }).then(function(data) {
        return window.location.href + data.alias;
    });
}

function showSnackbar(text) {
    var x = document.getElementById('snackbar');

    x.innerHTML = text;
    x.className = 'show';

    setTimeout(function(){ x.className = x.className.replace('show', ''); }, 3000);
}

function copy(url) {
    var dummy = document.createElement("textarea");
    document.body.appendChild(dummy);
    dummy.value = url;
    dummy.select();
    document.execCommand("copy");
    document.body.removeChild(dummy);
    showSnackbar("URL copied!");
}

function showResult(url) {
    var results = document.getElementById('results');
    
    var div = document.createElement('div');
    var button = document.createElement('button');
    var a = document.createElement('a');

    div.id = 'result';

    a.href = url;
    a.innerHTML = url;
    a.target = '_blank';
    a.className = 'shorten-url';
    div.appendChild(a);

    button.innerHTML = 'copy';
    button.className = 'copy-button';
    button.addEventListener('click', () => copy(url));
    div.appendChild(button);

    results.prepend(div);

    setTimeout(function() {
        div.className = 'show';
    }, 10);
}

function processForm(e) {
    e.preventDefault();
    var url = document.getElementById('url');

    getShortenURL(url.value)
        .then(function(shortenURL) {
            showResult(shortenURL);
            url.value = '';
        });
}

document.addEventListener('DOMContentLoaded', (event) => {
    var form = document.getElementById('form');
    if (form.attachEvent) {
        form.attachEvent("submit", processForm);
    } else {
        form.addEventListener("submit", processForm);
    }
})
