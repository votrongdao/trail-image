'use strict';

const CacheBase = require('./cache-base');
const is = require('../is.js');
const db = {};

/**
 * Hash kept in memory
 * Match Redis behavior
 * @extends {EventEmitter}
 * @extends {CacheBase}
 */
class MemoryCache extends CacheBase {
	constructor() {
		super();
		log = require('../config.js').provider.log;
		this.connected = true;
	}

	select(key, callback) {
		if (contains(key)) {
			callback(errorType.none, db[key]);
		} else {
			callback(errorType.notFound);
		}
	}

	selectMember(key, memberKey, callback) {
		if (contains(key, memberKey)) {
			callback(errorType.none, db[key][memberKey]);
		} else {
			callback(errorType.notFound);
		}
	}

	selectAll(key, callback) {
		// returning a field that's a hash is no different than returning a primitive type
		this.select(key, callback);
	}

	/**
	 * Insert value into hash
	 * @param {String} key
	 * @param {String} value
	 * @param {function(Object, (String|Number)} [callback] Method generated by responder()
	 */
	add(key, value, callback) {
		db[key] = value;
		if (is.callable(callback)) { callback(errorType.none, true); }
	}

	/**
	 * Insert value into hash
	 * @param {String} key
	 * @param {String} memberKey
	 * @param {String} value
	 * @param {function(Object, (String|Number)} [callback] Method generated by responder()
	 * @see http://redis.io/commands/hset
	 */
	addMember(key, memberKey, value, callback) {
		// match redis behavior
		let created = true;

		if (contains(key)) {
			created = false;
		} else {
			db[key] = {};
		}
		db[key][memberKey] = value;

		callback(errorType.none, created);
	}

	addAll(key, hash, callback) {
		db[key] = hash;
		if (is.callable(callback)) { callback(errorType.none, true); }
	}

	exists(key, callback) {
		callback(errorType.none, contains(key));
	}

	memberExists(key, memberKey, callback) {
		callback(errorType.none, contains(key, memberKey));
	}

	/**
	 * Keys matching pattern
	 * @param {String} pattern
	 * @param {function(Number, [String[]])} callback
	 */
	keys(pattern, callback) {
		let re = ruleToPattern(pattern);
		let keys = Object.keys(db).filter(k => re.test(k));
		let error = (keys.length > 0) ? errorType.none : errorType.notFound;

		callback(error, keys);
	}

	memberKeys(key, callback) {
		let fields = [];
		let error = errorType.notFound;

		if (is.defined(db, key)) {
			fields = Object.keys(db[key]);
			if (fields.length > 0) { error = errorType.none; }
		}
		callback(error, fields);
	}

	/**
	 * @param {String|String[]} key One or more item keys
	 * @param {function} callback
	 */
	remove(key, callback) {
		let removed = 0;

		if (is.array(key)) {
			for (let k of key) {
				if (contains(k)) {
					removed++;
					delete db[k];
				}
			}
		} else if (contains(key)) {
			removed = 1;
			delete db[key];
		}
		callback(errorType.none, removed);
	}

	removeMember(key, memberKey, callback) {
		let removed = 0;

		if (contains(key, memberKey)) {
			removed = 1;
			delete db[key][memberKey];
		}
		callback(errorType.none, removed);
	}

	/**
	 * Whether cache access caused an error
	 * @param {String|String[]} key
	 * @param {Object|Number|String} err
	 * @return {Boolean}
	 */
	hasError(key, err) {
		if (is.value(err) && err != errorType.none) {
			if (is.array(key)) { key = key.toString(); }
			let message = 'an unknown error';

			switch (err) {
				case errorType.notFound: message = 'key not found'; break;
			}
			log.error('Cache operation with key "%s" resulted in %s', key, message);
			return true;
		}
		return false;
	}

	/**
	 * Objects are stored natively, not serialized
	 * Will convert undefined to null
	 * @param {Object|String} value
	 */
	parseObject(value) { return (is.value(value)) ? value : null; }

	/**
	 * Objects are stored natively so no need to normalize
	 * @param {Object|String|Array} value
	 */
	normalize(value) { return value; }

	/**
	 * Native boolean
	 * @param {Number|Boolean} value
	 * @returns {Boolean}
	 */
	isTrue(value) { return value; }

	/**
	 * Native boolean
	 * @param {String|Boolean} value
	 * @returns {Boolean}
	 */
	isOkay(value) { return value; }
}

module.exports = MemoryCache;

// - Private static members ---------------------------------------------------

/**
 * @type {LogBase}
 */
let log = null;

const errorType = {
	none: 0,
	notFound: 1
};

/**
 * Whether key and memberKey are in the memory hash
 * @param {String} key
 * @param {String} [memberKey]
 * @returns {Boolean}
 */
function contains(key, memberKey) {
	let defined = is.defined(db, key);

	return (defined && memberKey !== undefined)
		? is.defined(db[key], memberKey)
		: defined;
}

/**
 * Convert wildcards to regex pattern
 * @param {String} rule
 * @return {RegExp}
 */
function ruleToPattern(rule) {
	rule = rule.replace('*', '.*');
	rule = rule.replace('?', '.?');
	rule = '^' + rule + '$';
	return new RegExp(rule);
}