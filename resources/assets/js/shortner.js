var snackbarTimer

function getShortenURL(longURL, alias) {
    var body = {
        alias,
        long_url: longURL
    }

    return fetch('/api', {
        method: 'post',
        body: JSON.stringify(body)
    }).then(function(response) {
        return response.json();
    }).then(function(data) {
        if (data.alias) {
            return [window.location.origin, data.alias].join('/');
        }

        if (data.message) {
            throw new Error(data.message);
        }

        throw new Error('Error while generation shorten URL');
    });
}

function showSnackbar(text, style) {
    var x = document.getElementById('snackbar');

    x.innerHTML = text;
    x.className = 'show ' + style;

    clearTimeout(snackbarTimer)
    snackbarTimer = setTimeout(function(){ x.className = ''; }, 3000);
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
    var alias = document.getElementById('alias');

    if (!url.value) {
        return showSnackbar("URL is required", 'error');
    }

    var regex = new RegExp("^[A-Za-z0-9_-]+$");
    if(alias.value && !regex.test(alias.value)) {
        return showSnackbar("Not allowed alias", 'error');
    }

    getShortenURL(url.value, alias.value)
        .then(function(shortenURL) {
            showResult(shortenURL);
            showSnackbar("URL created!", "success")
            url.value = '';
            alias.value = '';
        })
        .catch(function(error) {
            showSnackbar(error.message, 'error');
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
