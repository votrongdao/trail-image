var setting = require('../settings.js');
/** @see https://npmjs.org/package/uglify-js */
var uglify = require('uglify-js');
/** @type {String} */
var key = 'menu';

/**
 * Default route action
 */
exports.view = function(req, res)
{
	var library = require('../models/library.js');

	res.setHeader('Vary', 'Accept-Encoding');
	res.fromCache(key, 'application/javascript', function(cacher)
	{
		cacher('menu-script', {'library': library, 'setting': setting, 'layout': null }, function(text)
		{
			return uglify.minify(text, {fromString: true}).code;
		});
	});
};