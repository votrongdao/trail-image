"use strict";

var Enum = require('./enum.js');
var Format = require('./format.js');

/** @type {Number} */
exports.timestamp = new Date().getTime();
/** @type {Boolean} */
exports.isProduction = false;
/** @type {String} */
exports.logFile = './temp/trail-image.log';
/** @type {String} */
exports.emailRecipient = Format.decodeBase64(process.env.SMTP_RECIPIENT);
/** @type {url} */
exports.redis = null;

/**
 * @enum {String}
 */
exports.reCaptcha =
{
	privateKey: process.env.RECAPTCHA_PRIVATE,
	publicKey: process.env.RECAPTCHA_PUBLIC
};

/**
 * @enum {String|Boolean|Object}
 * @see https://devcenter.heroku.com/articles/config-vars
 */
exports.flickr =
{
	key: process.env.FLICKR_KEY,
	userID: '60950751@N04',
	secret: process.env.FLICKR_SECRET,
	token: process.env.FLICKR_TOKEN,
	tokenSecret: process.env.FLICKR_TOKEN_SECRET,
	favoriteSet: '72157631638576162',
	poemSet: '72157632729508554',
	defaultCollection: '72157630885395608',
	/** @enum {String} */
	url:
	{
		requestToken: 'http://www.flickr.com/services/oauth/request_token',
		authorize: 'http://www.flickr.com/services/oauth/authorize',
		accessToken: 'http://www.flickr.com/services/oauth/access_token'
	},
	useCache: true
};

/**
 * @enum {String}
 * @see https://github.com/ciaranj/node-oauth
 */
exports.oauth =
{
	version: '1.0A',
	encryption: 'HMAC-SHA1'
};

exports.cacheDuration = Enum.time.day * 2;
exports.retryDelay = Enum.time.second * 30;

/**
 * @enum {string|boolean}
 * @see https://developers.facebook.com/docs/reference/plugins/like/
 * @see https://developers.facebook.com/apps/110860435668134/summary
 */
exports.facebook =
{
	appID: '110860435668134',
	pageID: '241863632579825',
	siteID: '578261855525416',
	adminID: '1332883594',
	enabled: true,
	authorURL: 'https://www.facebook.com/jason.e.abbott'
};

/**
 * @see http://code.google.com/apis/console/#project:1033232213688
 * @see http://developers.google.com/maps/documentation/staticmaps/
 * @type {string}
 */
exports.google =
{
	apiKey: process.env.GOOGLE_KEY,
	projectID: '1033232213688',
	analyticsID: '22180727',        // shown as 'UA-22180727-1
	searchEngineID: process.env.GOOGLE_SEARCH_ID,
	blogID: '118459106898417641',
	userID: Format.decodeBase64(process.env.SMTP_LOGIN),
	password: Format.decodeBase64(process.env.SMTP_PASSWORD),
	maxMarkers: 70      // max URL is 2048; 160 is used for base URL; each coordinates needs about 26 characters
};

/**
 * @type {string}
 * @const
 */
exports.domain = 'trailimage.com';

/**
 * @type {String}
 * @const
 */
exports.title = 'Trail Image';

/**
 * @type {String}
 * @const
 */
exports.description = 'Stories, image and videos of small adventure trips in and around the state of Idaho';