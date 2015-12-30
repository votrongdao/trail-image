'use strict';

const TI = require('../');
const ActionSchema = TI.LinkData.Action;
const qi = 'query-input';

/**
 * @extends TI.LinkData.Action
 * @extends TI.LinkData.Thing
 * @alias TI.LinkData.SearchAction
 * @see http://schema.org/SearchAction
 * @see https://developers.google.com/structured-data/slsb-overview
 */
class SearchActionSchema extends ActionSchema {
	/**
	 * @param {String} type
	 */
	constructor(type) {
		if (type === undefined) { type = TI.LinkData.Type.searchAction; }

		super(type);

		/** @type String */
		this.query = null;
		/** @type String */
		this[qi] = null;
	}

	/** @param {String} v */
	set queryInput(v) { this[qi] = v; }
	/** @return {String} */
	get queryInput() { return this[qi]; }
}

module.exports = SearchActionSchema;