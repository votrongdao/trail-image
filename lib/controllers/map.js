"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const is_1 = require("../is");
const logger_1 = require("../logger");
const node_fetch_1 = require("node-fetch");
const config_1 = require("../config");
const kml_1 = require("../map/kml");
const geojson_1 = require("../map/geojson");
const template_1 = require("../template");
const library_1 = require("../library");
const _1 = require("../factory/");
const google_1 = require("../providers/google");
const compress = require("zlib");
const constants_1 = require("../constants");
let google = google_1.default;
function view(post, req, res) {
    if (is_1.default.value(post)) {
        const key = post.isPartial ? post.seriesKey : post.key;
        const photoID = req.params[constants_1.route.PHOTO_ID];
        post.getPhotos().then(() => {
            res.render(template_1.default.page.MAPBOX, {
                layout: template_1.default.layout.NONE,
                title: post.name() + ' Map',
                description: post.description,
                post,
                key,
                photoID: is_1.default.numeric(photoID) ? photoID : 0,
                config: config_1.default
            });
        });
    }
    else {
        res.notFound();
    }
}
function post(req, res) {
    view(library_1.default.postWithKey(req.params[constants_1.route.POST_KEY]), req, res);
}
function series(req, res) {
    view(library_1.default.postWithKey(req.params[constants_1.route.SERIES_KEY], req.params[constants_1.route.PART_KEY]), req, res);
}
function blog(req, res) {
    res.render(template_1.default.page.MAPBOX, {
        layout: template_1.default.layout.NONE,
        title: config_1.default.site.title + ' Map',
        config: config_1.default
    });
}
function photoJSON(req, res) {
    _1.default.map.photos()
        .then(item => { res.sendCompressed(constants_1.mimeType.JSON, item); })
        .catch(err => {
        logger_1.default.error(err);
        res.notFound();
    });
}
function trackJSON(req, res) {
    _1.default.map.track(req.params[constants_1.route.POST_KEY])
        .then(item => { res.sendCompressed(constants_1.mimeType.JSON, item); })
        .catch(err => {
        logger_1.default.error(err);
        res.notFound();
    });
}
function source(req, res) {
    const key = req.params[constants_1.route.MAP_SOURCE];
    if (!is_1.default.text(key)) {
        return res.notFound();
    }
    const s = config_1.default.map.source[key.replace('.json', '')];
    if (!is_1.default.value(s)) {
        return res.notFound();
    }
    const parser = fetchKMZ(s.provider);
    node_fetch_1.default(s.url, { headers: { 'User-Agent': 'node.js' } }).then(reply => {
        if (reply.status == constants_1.httpStatus.OK) {
            parser(reply)
                .then(JSON.stringify)
                .then(geoText => {
                compress.gzip(Buffer.from(geoText), (err, buffer) => {
                    if (is_1.default.value(err)) {
                        res.internalError(err);
                    }
                    else {
                        res.setHeader(constants_1.header.content.ENCODING, constants_1.encoding.GZIP);
                        res.setHeader(constants_1.header.CACHE_CONTROL, 'max-age=86400, public');
                        res.setHeader(constants_1.header.content.TYPE, constants_1.mimeType.JSON + ';charset=utf-8');
                        res.setHeader(constants_1.header.content.DISPOSITION, `attachment; filename=${key}`);
                        res.write(buffer);
                        res.end();
                    }
                });
            })
                .catch(err => {
                res.internalError(err);
            });
        }
        else {
            res.end(reply.status);
        }
    });
}
const fetchKMZ = (sourceName) => (res) => res.buffer().then(kml_1.default.fromKMZ).then(geojson_1.default.featuresFromKML(sourceName));
function gpx(req, res) {
    const post = config_1.default.map.allowDownload ? library_1.default.postWithKey(req.params[constants_1.route.POST_KEY]) : null;
    if (is_1.default.value(post)) {
        google.drive.loadGPX(post, res)
            .then(() => { res.end(); })
            .catch(res.notFound);
    }
    else {
        res.notFound();
    }
}
exports.default = {
    gpx,
    post,
    series,
    blog,
    json: {
        blog: photoJSON,
        post: trackJSON
    },
    source,
    inject: {
        set google(g) { google = g; }
    }
};
//# sourceMappingURL=map.js.map