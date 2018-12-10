const app = require('express');
const request = require('request');
const Datastore = require('nedb');
require('dotenv').load();

let auth = {
	app_id: typeof(process.env.APP_ID) == 'string' ? process.env.APP_ID : null,
	app_key: typeof(process.env.APP_KEY) == 'string' ? process.env.APP_KEY : null,
};

let req_info = {
	num: 54, 	// Bus line number
	id: 208		// Stop id
};

function format_url (info) {
	return `https://api.tmb.cat/v1/ibus/lines/${info.num}/stops/${info.id}`;
}

function get_data (info) {
	if (auth.app_id == null || auth.app_key == null) {
		console.log('Auth error. Put your APP_ID and APP_KEY into a .env file.')
		process.exit(1);
	}
	
	if (info.num == null || info.id == null) {
		console.log('Auth error. Put your APP_ID and APP_KEY into a .env file.')
		process.exit(1);
	}

	
	request.get({ uri: url, qs: auth }, 
		(err, res, body) => {
			if (err) return err;
			
			let data = JSON.parse(body).data.ibus[0];
			let time_s = data['t-in-s'];
			let time_m = data['t-in-min'];
			console.log(`Time left for ${req_info.num} @ stop ${req_info.id}: ${time_m} min ${time_s - 60 * time_m} s.`)
		});
}

get_data();
