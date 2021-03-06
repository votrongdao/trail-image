const res = require('../mocks/response.mock');
const req = require('../mocks/request.mock');
const mocha = require('mocha');
const { prepare } = require('./index.test');
const { expect } = require('chai');

describe('RSS', ()=> {
   before(done => prepare(done));
   beforeEach(() => { res.reset(); req.reset(); });

   it('generates valid RSS 2.0 XML', ()=> {
      const Feed = require('feed');
      const nl = '\n';
      const tab = '    ';
      const updated = new Date();
      const authorName = 'Test Person';
      const title = 'Feed Title';
      const description = 'Feed Description';
      const url = 'http://www.domain.com';
      const image = 'http://www.domain.com/img/logo.png';
      const author = { name: authorName, link: 'https://www.facebook.com/test.person' };
      const copyright = 'Copyright © ' + updated.getFullYear() + ' ' + authorName + '. All rights reserved';
      const feed = new Feed({
         title: title,
         description: description,
         link: url,
         image: image,
         copyright: copyright,
         author: author,
         updated: updated
      });
      const source = feed.rss2();
      const target = '<?xml version="1.0" encoding="utf-8"?>' + nl
         + '<rss version="2.0">' + nl
         + tab + '<channel>' + nl
         + tab + tab + '<title>' + title + '</title>' + nl
         + tab + tab + '<link>' + url + '</link>' + nl
         + tab + tab + '<description>' + description + '</description>' + nl
         + tab + tab + '<lastBuildDate>' + updated.toUTCString() + '</lastBuildDate>' + nl
         + tab + tab + '<docs>http://blogs.law.harvard.edu/tech/rss</docs>' + nl
         + tab + tab + '<generator>Feed for Node.js</generator>' + nl
         + tab + tab + '<image>' + nl
         + tab + tab + tab + '<title>' + title + '</title>' + nl
         + tab + tab + tab + '<url>' + image + '</url>' + nl
         + tab + tab + tab + '<link>' + url + '</link>' + nl
         + tab + tab + '</image>' + nl
         + tab + tab + '<copyright>' + copyright + '</copyright>' + nl
         + tab + '</channel>' + nl
         + '</rss>';

      expect(source).equals(target);
   });
});