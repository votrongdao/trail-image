'use strict';

const BaseElement = require('./elements/pdf-element.js');
const PhotoPage = require('./pages/photo-page.js');
const CoverPage = require('./pages/cover-page.js');
const CopyrightPage = require('./pages/copyright-page.js');
const IndexPage = require('./pages/index-page.js');

/**
 * @extends {BaseElement}
 * @see https://github.com/devongovett/pdfkit
 */
class PrintBook extends BaseElement {

	constructor() {
		super();

		/** @type {PhotoPage[]} */
		this.pages = [];
		/**
		 * Current page index while rendering
		 * @type {Number}
		 * @private
		 */
		this._renderIndex = 0;
	}

	/**
	 * @param {PhotoPage|CopyrightPage|CoverPage|IndexPage} p
	 * @return {PhotoPage|CopyrightPage|CoverPage|IndexPage}
	 */
	add(p) { this.pages.push(p); return p; }

	/**
	 * @param {Post} post
	 * @return {PrintBook}
	 */
	static fromPost(post) {
		let b = new PrintBook();
		// page count
		let c = 1;
		let indexPage = new IndexPage();

		b.add(CoverPage.fromPost(post));
		//b.add(new CopyrightPage());
		//
		//for (let p of post.photos) {
		//	b.add(PhotoPage.fromPhoto(p, c++));
		//	indexPage.addWord(p.tags, c);
		//}
		//b.add(indexPage);

		return b;
	}

	/**
	 * Render the book
	 * @param {PDFStyle} style
	 * @param {PDFDocument} pdf
	 * @param {function} [callback]
	 */
	render(style, pdf, callback) {
		this.pages[this._renderIndex].render(style, pdf, ()=> {
			this._renderIndex++;

			if (this._renderIndex >= this.pages.length) {
				callback();
			} else {
				this.render(style, pdf, callback);
			}
		});
	}
}

module.exports = PrintBook;