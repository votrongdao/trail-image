'use strict';

const is = require('../../is.js');
const BaseElement = require('./pdf-element.js');
const Align = BaseElement.Align;

/**
 * @extends {PDFElement}
 */
class TextElement extends BaseElement {
	/**
	 * @param {String|Number} text
	 * @param {String} [style] pdf-style.json rule
	 */
	constructor(text, style) {
		super((style === undefined) ? 'defaultText' : style);
		/**
		 * @type {String}
		 */
		this.font = null;
		/**
		 * @type {Number}
		 */
		this.fontSize = 12;
		/**
		 * @type {String}
		 */
		this.text = text;
		/**
		 * @type {Number}
		 */
		this.indent = 0;
	}

	/**
	 * Whether caption is empty
	 * @returns {Boolean}
	 */
	get empty() { return is.empty(this.text); }

	/**
	 * Height of string wrapped to width
	 * Calculator deals with PDF points (72 DPI) or pixels
	 * @param {PDFDocument} pdf
	 */
	calculateHeight(pdf) {
		this.heightPixels = pdf
			.font(this.font).fontSize(this.fontSize)
			.heightOfString(this.text, { width: this.widthPixels });
	}

	/**
	 * @param {PDFStyle} style
	 * @param {PDFDocument} pdf
	 * @param {function} callback
	 */
	render(style, pdf, callback) {
		this.updateLayout(style, pdf);

		let p = this.layoutPixels;
		let options = {};
		let addTop = isNaN(this.margin.top) ? 0 : this.inchesToPixels(this.margin.top);
		let addBottom = isNaN(this.margin.bottom) ? 0 : this.inchesToPixels(this.margin.bottom);

		if (!isNaN(p.width)) { options.width = p.width; }

		switch (this.alignContent) {
			case Align.Center: options.align = 'center'; break;
			case Align.Justify: options.align = 'justify'; break;
			case Align.Right: options.align = 'right'; break;
		}

		pdf.font(this.font).fontSize(this.fontSize).fillColor(this.color, this.opacity);
		pdf.y += addTop;

		if (isNaN(p.left) && isNaN(p.top)) {
			// relative position
			pdf.text(this.text, options);
		} else if (isNaN(p.top)) {
			pdf.text(this.text, p.left, pdf.y, options);
		} else if (isNaN(p.left)) {
			pdf.text(this.text, pdf.x, p.top, options);
		} else {
			// absolute position
			pdf.text(this.text, p.left, p.top, options);
		}
		pdf.y += addBottom;

		callback();
	}
}

module.exports = TextElement;