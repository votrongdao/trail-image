// translate Flickr, Google and Redis responses into standard objects

const is = require('../is');
const log = require('../logger');
const map = require('../map');
const config = require('../config');
const library = require('./library');
const category = require('./category');
const photoSize = require('./photo-size');
const exif = require('./exif');
const cache = require('./cache');

const BLOG_JSON_KEY = 'blog-map';

// can be replaced with injection
let flickr = require('./flickr');
let google = require('./google');

/**
 * @param {boolean} [emptyIfLoaded] Whether to reset the library before loading
 * @returns {Promise.<Library>} Resolve with list of changed post keys
 */
function buildLibrary(emptyIfLoaded = true) {
   // store existing post keys to compute changes
   const hadPostKeys = library.postKeys();
   if (emptyIfLoaded && library.loaded) { library.empty(); }
   // reset changed keys to none
   library.changedKeys = [];

   return Promise
      .all([flickr.getCollections(), flickr.getAllPhotoTags()])
      .then(([collections, tags]) => {
         // parse collections and photo tags
         library.tags = is.value(tags) ? parsePhotoTags(tags) : {};
         collections.forEach(c => category.make(c, true));
         correlatePosts();
         library.loaded = true;
         log.infoIcon('photo_library', 'Loaded %d photo posts from Flickr: beginning detail retrieval', library.posts.length);
         // retrieve additional post info without waiting for it to finish
         Promise
            .all(library.posts.map(p => getPostInfo(p)))
            .then(() => {
               library.postInfoLoaded = true;
               log.info('Finished loading post details');
            });

         return Promise.resolve();
      })
      .then(()=> {
         // attach Flickr lookup methods to the library so it doesn't have
         // to require factory or Flickr modules (avoid circular dependencies)
         library.getPostWithPhoto = getPostWithPhoto;
         library.getEXIF = getEXIF;
         library.getPhotosWithTags = getPhotosWithTags;
         library.load = buildLibrary;
         return Promise.resolve();
      })
      .then(()=> {
         // find changed post and category keys so their caches can be invalidated
         if (hadPostKeys.length > 0) {
            let changedKeys = [];
            library.posts
               .filter(p => hadPostKeys.indexOf(p.key) == -1)
               .forEach(p => {
                  log.info('Found new post "%s"', p.title);
                  // all post categories will need to be refreshed
                  changedKeys = changedKeys.concat(Object.keys(p.categories));
                  // update adjecent posts to correct next/previous links
                  if (is.value(p.next)) { changedKeys.push(p.next.key); }
                  if (is.value(p.previous)) { changedKeys.push(p.previous.key); }
               });
            library.changedKeys = changedKeys;
         }
         return library;
      });
}

/**
 * Match posts that are part of a series
 */
function correlatePosts() {
   let p = library.posts[0];
   let parts = [];

   while (p != null && p.previous != null) {
      if (p.subTitle !== null) {
         parts.push(p);

         while (p.previous != null && p.previous.title == p.title) {
            p = p.previous;
            parts.unshift(p);
         }

         if (parts.length > 1) {
            parts[0].makeSeriesStart();

            for (let i = 0; i < parts.length; i++) {
               parts[i].part = i + 1;
               parts[i].totalParts = parts.length;
               parts[i].isPartial = true;

               if (i > 0) { parts[i].previousIsPart = true; }
               if (i < parts.length - 1) { parts[i].nextIsPart = true; }
            }
         } else {
            p.ungroup();
         }
         parts = [];
      }
      p = p.previous;
   }
}

/**
 * @this {Library}
 * @param {Photo|string} photo
 * @returns {Promise}
 */
function getPostWithPhoto(photo) {
   const id = (typeof photo == is.type.STRING) ? photo : photo.id;
   return flickr.getPhotoContext(id).then(sets => (is.value(sets))
      ? this.posts.find(p => p.id == sets[0].id)
      : null
   );
}

/**
 * Load photos for post and calculate summaries
 * @param {Post} p
 * @returns {Promise.<Photo[]>}
 */
const getPostPhotos = p => p.photosLoaded
   ? Promise.resolve(p.photos)
   : flickr.getSetPhotos(p.id).then(p.updatePhotos);

/**
 * Add information to existing post object
 * @param {Post} p
 * @returns {Promise.<Post>}
 */
const getPostInfo = p => p.infoLoaded
   ? Promise.resolve(p)
   : flickr.getSetInfo(p.id).then(p.updateInfo);

/**
 * All photos with given tags
 * @param {string|string[]} tags
 * @returns {Promise.<Photo[]>}
 */
const getPhotosWithTags = tags => flickr.photoSearch(tags)
   .then(photos => photos.map(json => ({
      id: json.id,
      size: { thumb: photoSize.make(json, config.flickr.photoSize.search[0]) }
   })));

/**
 * Convert tags to hash of phrases keyed to their "clean" abbreviation
 * @param {Flickr.Tag[]} rawTags
 * @returns {object}
 */
function parsePhotoTags(rawTags) {
   const exclusions = is.array(config.flickr.excludeTags) ? config.flickr.excludeTags : [];
   return rawTags.reduce((tags, t) => {
      const text = t.raw[0]._content;
      // ensure not machine or exluded tag
      if (text.indexOf('=') == -1 && exclusions.indexOf(text) == -1) { tags[t.clean] = text; }
      return tags;
   }, {});
}

/**
 * @this {Photo}
 * @returns {Promise}
 */
function getEXIF() { return flickr.getExif(this.id).then(exif.make); }

/**
 * Load all map information (track and photo features) for a post
 * @param {string} postKey
 * @returns {Promise.<ViewCacheItem>}
 * @see http://geojsonlint.com/
 */
const mapForPost = postKey => config.cache.maps
   ? cache.map.getItem(postKey).then(item => is.cacheItem(item) ? item : loadMapForPost(postKey))
   : loadMapForPost(postKey);

/**
 * Load map photos for all posts.
 * @returns {Promise.<ViewCacheItem>}
 */
const mapForBlog = () => config.cache.maps
   ? cache.map.getItem(BLOG_JSON_KEY).then(item => is.cacheItem(item) ? item : loadMap())
   : loadMap();

/**
 * Get photo GeoJSON (not tracks) for all posts.
 * @returns {Promise.<ViewCacheItem>}
 */
const loadMap = () => Promise.resolve(map.features())
   .then(geo => mapPhotoFeatures(geo))
   .then(geo => cache.map.add(BLOG_JSON_KEY, geo));

/**
 * Get GeoJSON for single post. If post has no track then GPX will only include
 * photo markers.
 * @param {string} postKey
 * @returns {Promise.<ViewCacheItem>}
 */
function loadMapForPost(postKey) {
   const post = library.postWithKey(postKey);

   if (!is.value(post)) { throw new ReferenceError(`Post ${postKey} not found in library`); }

   const noGPX = Promise.resolve(map.features());
   const getFeatures = (post.triedTrack && !post.hasTrack)
      ? noGPX
      : google.drive.loadGPX(post)
         .then(map.featuresFromGPX)
         .catch(() => noGPX);

   return getFeatures
      .then(geo => mapPostPhotoFeatures(post, geo))
      .then(geo => cache.map.add(postKey, geo));
}

/**
 * Append photo GeoFeatures to GeoJSON
 * @param {GeoJSON.FeatureCollection} [geo]
 * @returns {Promise.<object>} GeoJSON
 */
const mapPhotoFeatures = geo => new Promise(resolve => { addPhotoFeatures(geo, resolve); });

/**
 * Append photo GeoFeatures to GeoJSON
 * @param {Post} post
 * @param {GeoJSON.FeatureCollection} [geo]
 * @returns {Promise.<object>} GeoJSON
 */
const mapPostPhotoFeatures = (post, geo) => new Promise(resolve => {
   // move to the first post in a series
   if (post.isPartial) { while (!post.isSeriesStart) { post = post.previous; } }
   addPostPhotoFeatures(post, geo, resolve);
});

/**
 * Add GeoJSON feature information for all photos in library.
 * @param {GeoJSON.FeatureCollection} geo
 * @param {function} resolve
 */
function addPhotoFeatures(geo, resolve) {
   library.getPhotos().then(photos => {
      geo.features = geo.features.concat(photos
         .filter(p => p.latitude > 0)
         .map(p => map.pointFromPhoto(p)));

      resolve(geo);
   });
}

module.exports = {
   buildLibrary,
   map: {
      forPost: mapForPost,
      forBlog: mapForBlog
   },
   // inject different data providers
   inject: {
      set flickr(f) { flickr = f; },
      set google(g) { google = g; }
   }
};