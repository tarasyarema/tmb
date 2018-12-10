const express = require('express');
const request = require('request');
const app = express();
require('dotenv').load();

const auth = {
    app_id: typeof(process.env.APP_ID) == 'string' ? process.env.APP_ID : null,
    app_key: typeof(process.env.APP_KEY) == 'string' ? process.env.APP_KEY : null,
};

function format_url (info) {
    if (typeof(info.num) != 'number' || typeof(info.id) != 'number') {
        return null;
    } else {
        return `https://api.tmb.cat/v1/ibus/lines/${info.num}/stops/${info.id}`;
    }
}

app.get('/', (req, res) => {
    console.log('GET /');
    console.log('  ▷  Redirected to github page.')
    res.redirect('https://github.com/tarasyarema/tmb');
});

app.get('/api', (req, res) => {
    console.log('GET /api');
    console.log('  ▷  Redirected to github page.')
    res.redirect('https://github.com/tarasyarema/tmb');
});

app.get('/api/:id/:num', (req, res) => {    
    let info = {
        num: parseInt(req.params.num),
        id: parseInt(req.params.id),
    };
   
    console.log(`GET /api/${info.id}/${info.num}`);

    if (auth.app_id == null || auth.app_key == null) {
        console.log('Auth error. Put your APP_ID and APP_KEY into a .env file.')
        process.exit(1);
    }

    let url = format_url(info);

    if (typeof(url) == 'null') {
        console.log('Info error.')
        process.exit(1);
    }

    request.get({ uri: url, qs: auth }, 
        (error, response, body) => {
            if (error) {
                console.log(error);
                return error;
            }

            let data = JSON.parse(body).data.ibus[0];
            let time_s = data['t-in-s'];
            let time_m = data['t-in-min'];

            let time_left = {
                min: time_m,
                seg: time_s - 60 * time_m
            }; 

            console.log('  ▷ JSON response sent.')
            res.json(time_left);
    });
});

app.listen(process.env.PORT);
