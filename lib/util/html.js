"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const is_1 = require("../is");
const regex_1 = require("../regex");
const config_1 = require("../config");
const text_1 = require("./text");
function story(text) {
    if (!is_1.default.empty(text)) {
        if (regex_1.default.poetry.all.test(text)) {
            text = text.replace(regex_1.default.poetry.delimiter, '');
            if (regex_1.default.haiku.all.test(text)) {
                text = formatHaiku(text, regex_1.default.haiku.all);
            }
            else {
                text = '<p class="poem">' + text
                    .replace(regex_1.default.lineBreak, '<br/>')
                    .replace(regex_1.default.poetry.indent, '<span class="tab"></span>') + '</p>';
            }
        }
        else if (regex_1.default.haiku.any.test(text)) {
            text = formatHaiku(text, regex_1.default.haiku.any);
        }
        else {
            text = caption(text);
        }
    }
    return text;
}
exports.story = story;
function fixMalformedLink(text) {
    let index = 0;
    text = text.replace(regex_1.default.tag.truncatedLink, (match, missedPart, i) => {
        index = i;
        return missedPart + '</a>';
    });
    if (index > 0) {
        const protocol = /https?:\/\//;
        const oldLink = text.substring(text.lastIndexOf('<a', index), text.indexOf('</a>', index) + 4);
        const newLink = oldLink.replace(regex_1.default.tag.link, (match, url, name) => {
            if (!protocol.test(name)) {
                name = 'http://' + name;
            }
            return text_1.format('<a href="{0}">{1}</a>', name, decodeURI(name.replace(protocol, '')));
        });
        text = text.replace(oldLink, newLink);
    }
    else {
        text = text.replace(regex_1.default.tag.ellipsisLink, '<a href="$1$2$3">$2$3</a>');
    }
    return text;
}
exports.fixMalformedLink = fixMalformedLink;
exports.shortenLinkText = (text) => text.replace(regex_1.default.tag.linkToUrl, (match, protocol, url) => {
    const parts = url.split('/');
    const domain = parts[0].replace('www.', '');
    let lastPart = /\/$/.test(url) ? parts.length - 2 : parts.length - 1;
    if (lastPart > 0 && /^[\?#]/.test(parts[lastPart])) {
        lastPart--;
    }
    let middle = '/';
    const page = parts[lastPart]
        .replace(regex_1.default.queryString, '')
        .replace(regex_1.default.tag.anchor, '')
        .replace(regex_1.default.fileExt, '');
    if (lastPart > 1) {
        middle = '/&hellip;/';
    }
    if (protocol === undefined) {
        protocol = 'http://';
    }
    return '<a href="' + protocol + url + '">' + domain + middle + decodeURIComponent(page) + '</a>';
});
exports.fraction = (text) => text.replace(/(\d+)\/(\d+)/, '<sup>$1</sup>&frasl;<sub>$2</sub>');
exports.typography = (text) => is_1.default.empty(text) ? '' : text
    .replace(regex_1.default.quote.rightSingle, '$1&rsquo;')
    .replace(regex_1.default.quote.leftSingle, '&lsquo;$1')
    .replace(regex_1.default.quote.rightDouble, '$1&rdquo;')
    .replace(regex_1.default.quote.leftDouble, '&ldquo;$2')
    .replace(regex_1.default.tag.encodedLink, (match, tag, name) => tag.replace(regex_1.default.quote.html, '"') + name + '</a>');
function photoTagList(list) {
    let links = '';
    const link = '<a href="/photo-tag/{0}" rel="tag">{1}</a>';
    if (is_1.default.array(list)) {
        list
            .sort()
            .forEach(t => { links += text_1.format(link, t.toLowerCase().replace(/\W/g, ''), t) + ' '; });
    }
    return links;
}
exports.photoTagList = photoTagList;
exports.iconTag = (name) => `<i class="material-icons ${name}">${name}</i>`;
function postCategoryIcon(title) {
    const map = config_1.default.style.icon.category;
    if (is_1.default.value(map)) {
        for (const name in map) {
            if (name == title) {
                return exports.iconTag(map[name]);
            }
        }
        if (map.default) {
            return exports.iconTag(map.default);
        }
    }
    return '';
}
exports.postCategoryIcon = postCategoryIcon;
function postModeIcon(categories) {
    const icons = config_1.default.style.icon;
    const map = icons.post;
    if (!is_1.default.array(categories)) {
        categories = Object.keys(categories);
    }
    if (is_1.default.value(map)) {
        const iconName = Object.keys(map).find(iconName => {
            const re = map[iconName];
            return categories.find(c => re.test(c)) !== undefined;
        });
        if (is_1.default.value(iconName)) {
            return iconName;
        }
        else if (icons.postDefault) {
            return icons.postDefault;
        }
    }
    return '';
}
exports.postModeIcon = postModeIcon;
function linkPattern(url) {
    return '<a href="' + url + '$1" target="_blank">$1</a>';
}
exports.linkPattern = linkPattern;
function formatNotes(notes) {
    const start = (/^\s*\*/g.test(notes)) ? ' start="0"' : '';
    notes = '<ol class="footnotes"' + start + '><li><span>'
        + notes
            .replace(/[⁰¹²³⁴⁵⁶⁷⁸⁹]+\s*/g, '')
            .replace(/[\r\n]+/g, '</span></li><li><span>')
            .replace(/<\/span><\/li><li><span>\s*$/, '')
        + '</span></li></ol>';
    return notes.replace(/<li><span>\s*\*\s*/gi, '<li class="credit">' + exports.iconTag('star') + '<span>');
}
exports.formatNotes = formatNotes;
function formatHaiku(text, regex) {
    const match = regex.exec(text);
    return '<p class="haiku">'
        + match[1] + '<br/>'
        + match[2] + '<br/>'
        + match[3] + exports.iconTag('spa') + '</p>'
        + caption(text.replace(match[0], ''));
}
exports.formatHaiku = formatHaiku;
function formatPoem(text) {
    return '<blockquote class="poem"><p>' + text
        .replace(regex_1.default.trailingWhiteSpace, '')
        .replace(regex_1.default.lineBreak, '<br/>')
        .replace(/(<br\/>){2,}/gi, '</p><p>')
        .replace(regex_1.default.poetry.indent, '<span class="tab"></span>')
        .replace(regex_1.default.footnote.number, '$1<sup>$2</sup>')
        + '</p></blockquote>';
}
exports.formatPoem = formatPoem;
function caption(text) {
    if (!is_1.default.empty(text)) {
        const ph = '[POEM]';
        let footnotes = '';
        let poem = '';
        text = fixMalformedLink(text);
        text = exports.shortenLinkText(text);
        text = exports.typography(text);
        text = text
            .replace(regex_1.default.footnote.text, (match, prefix, body) => {
            footnotes = formatNotes(body);
            return '';
        })
            .replace(regex_1.default.poetry.any, (match, space, body) => {
            poem = formatPoem(body);
            return ph;
        })
            .replace(regex_1.default.quote.block, (match, newLines, body) => '[Q]' + body.replace(regex_1.default.quote.curly, '') + '[/Q]');
        text = '<p>' + text + '</p>';
        text = text
            .replace(regex_1.default.newLine, '</p><p>')
            .replace(regex_1.default.tag.emptyParagraph, '')
            .replace(regex_1.default.quip, (match, tag, body) => '<p class="quip">' + body)
            .replace(regex_1.default.footnote.number, '$1<sup>$2</sup>')
            .replace(/\[\/Q][\r\n\s]*([^<]+)/g, '[/Q]<p class="first">$1')
            .replace(/(<p>)?\[Q]/g, '<blockquote><p>')
            .replace(/\[\/Q](<\/p>)?/g, '</p></blockquote>');
        if (poem.length > 0) {
            text = text
                .replace(ph, '</p>' + poem + '<p class="first">')
                .replace(regex_1.default.tag.emptyParagraph, '');
        }
        return text + footnotes;
    }
    return '';
}
exports.caption = caption;
exports.logMessage = (r, fieldName) => {
    if (is_1.default.defined(r, fieldName) && is_1.default.value(r[fieldName])) {
        r[fieldName] = r[fieldName]
            .replace(/(\d{10,11})/, linkPattern(config_1.default.log.photoUrl))
            .replace(regex_1.default.log.path, '<a href="$1" target="_blank">$1</a>$2')
            .replace(regex_1.default.ipAddress, linkPattern(config_1.default.log.ipLookupUrl));
    }
    else {
        r[fieldName] = '[no message]';
    }
    return r[fieldName];
};
exports.characterEntities = (text) => text.replace(/[\u00A0-\u2666<>\&]/g, c => '&' + (htmlEntity[c.charCodeAt(0)] || '#' + c.charCodeAt(0)) + ';');
const htmlEntity = {
    34: 'quot',
    38: 'amp',
    39: 'apos',
    60: 'lt',
    62: 'gt',
    160: 'nbsp',
    161: 'iexcl',
    162: 'cent',
    163: 'pound',
    164: 'curren',
    165: 'yen',
    166: 'brvbar',
    167: 'sect',
    168: 'uml',
    169: 'copy',
    170: 'ordf',
    171: 'laquo',
    172: 'not',
    173: 'shy',
    174: 'reg',
    175: 'macr',
    176: 'deg',
    177: 'plusmn',
    178: 'sup2',
    179: 'sup3',
    180: 'acute',
    181: 'micro',
    182: 'para',
    183: 'middot',
    184: 'cedil',
    185: 'sup1',
    186: 'ordm',
    187: 'raquo',
    188: 'frac14',
    189: 'frac12',
    190: 'frac34',
    191: 'iquest',
    192: 'Agrave',
    193: 'Aacute',
    194: 'Acirc',
    195: 'Atilde',
    196: 'Auml',
    197: 'Aring',
    198: 'AElig',
    199: 'Ccedil',
    200: 'Egrave',
    201: 'Eacute',
    202: 'Ecirc',
    203: 'Euml',
    204: 'Igrave',
    205: 'Iacute',
    206: 'Icirc',
    207: 'Iuml',
    208: 'ETH',
    209: 'Ntilde',
    210: 'Ograve',
    211: 'Oacute',
    212: 'Ocirc',
    213: 'Otilde',
    214: 'Ouml',
    215: 'times',
    216: 'Oslash',
    217: 'Ugrave',
    218: 'Uacute',
    219: 'Ucirc',
    220: 'Uuml',
    221: 'Yacute',
    222: 'THORN',
    223: 'szlig',
    224: 'agrave',
    225: 'aacute',
    226: 'acirc',
    227: 'atilde',
    228: 'auml',
    229: 'aring',
    230: 'aelig',
    231: 'ccedil',
    232: 'egrave',
    233: 'eacute',
    234: 'ecirc',
    235: 'euml',
    236: 'igrave',
    237: 'iacute',
    238: 'icirc',
    239: 'iuml',
    240: 'eth',
    241: 'ntilde',
    242: 'ograve',
    243: 'oacute',
    244: 'ocirc',
    245: 'otilde',
    246: 'ouml',
    247: 'divide',
    248: 'oslash',
    249: 'ugrave',
    250: 'uacute',
    251: 'ucirc',
    252: 'uuml',
    253: 'yacute',
    254: 'thorn',
    255: 'yuml',
    402: 'fnof',
    913: 'Alpha',
    914: 'Beta',
    915: 'Gamma',
    916: 'Delta',
    917: 'Epsilon',
    918: 'Zeta',
    919: 'Eta',
    920: 'Theta',
    921: 'Iota',
    922: 'Kappa',
    923: 'Lambda',
    924: 'Mu',
    925: 'Nu',
    926: 'Xi',
    927: 'Omicron',
    928: 'Pi',
    929: 'Rho',
    931: 'Sigma',
    932: 'Tau',
    933: 'Upsilon',
    934: 'Phi',
    935: 'Chi',
    936: 'Psi',
    937: 'Omega',
    945: 'alpha',
    946: 'beta',
    947: 'gamma',
    948: 'delta',
    949: 'epsilon',
    950: 'zeta',
    951: 'eta',
    952: 'theta',
    953: 'iota',
    954: 'kappa',
    955: 'lambda',
    956: 'mu',
    957: 'nu',
    958: 'xi',
    959: 'omicron',
    960: 'pi',
    961: 'rho',
    962: 'sigmaf',
    963: 'sigma',
    964: 'tau',
    965: 'upsilon',
    966: 'phi',
    967: 'chi',
    968: 'psi',
    969: 'omega',
    977: 'thetasym',
    978: 'upsih',
    982: 'piv',
    8226: 'bull',
    8230: 'hellip',
    8242: 'prime',
    8243: 'Prime',
    8254: 'oline',
    8260: 'frasl',
    8472: 'weierp',
    8465: 'image',
    8476: 'real',
    8482: 'trade',
    8501: 'alefsym',
    8592: 'larr',
    8593: 'uarr',
    8594: 'rarr',
    8595: 'darr',
    8596: 'harr',
    8629: 'crarr',
    8656: 'lArr',
    8657: 'uArr',
    8658: 'rArr',
    8659: 'dArr',
    8660: 'hArr',
    8704: 'forall',
    8706: 'part',
    8707: 'exist',
    8709: 'empty',
    8711: 'nabla',
    8712: 'isin',
    8713: 'notin',
    8715: 'ni',
    8719: 'prod',
    8721: 'sum',
    8722: 'minus',
    8727: 'lowast',
    8730: 'radic',
    8733: 'prop',
    8734: 'infin',
    8736: 'ang',
    8743: 'and',
    8744: 'or',
    8745: 'cap',
    8746: 'cup',
    8747: 'int',
    8756: 'there4',
    8764: 'sim',
    8773: 'cong',
    8776: 'asymp',
    8800: 'ne',
    8801: 'equiv',
    8804: 'le',
    8805: 'ge',
    8834: 'sub',
    8835: 'sup',
    8836: 'nsub',
    8838: 'sube',
    8839: 'supe',
    8853: 'oplus',
    8855: 'otimes',
    8869: 'perp',
    8901: 'sdot',
    8968: 'lceil',
    8969: 'rceil',
    8970: 'lfloor',
    8971: 'rfloor',
    9001: 'lang',
    9002: 'rang',
    9674: 'loz',
    9824: 'spades',
    9827: 'clubs',
    9829: 'hearts',
    9830: 'diams',
    338: 'OElig',
    339: 'oelig',
    352: 'Scaron',
    353: 'scaron',
    376: 'Yuml',
    710: 'circ',
    732: 'tilde',
    8194: 'ensp',
    8195: 'emsp',
    8201: 'thinsp',
    8204: 'zwnj',
    8205: 'zwj',
    8206: 'lrm',
    8207: 'rlm',
    8211: 'ndash',
    8212: 'mdash',
    8216: 'lsquo',
    8217: 'rsquo',
    8218: 'sbquo',
    8220: 'ldquo',
    8221: 'rdquo',
    8222: 'bdquo',
    8224: 'dagger',
    8225: 'Dagger',
    8240: 'permil',
    8249: 'lsaquo',
    8250: 'rsaquo',
    8364: 'euro'
};
//# sourceMappingURL=html.js.map