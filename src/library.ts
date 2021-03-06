import { Library, Post, Category, Photo, EXIF } from './types/';
import is from './is';

/**
 * Singleton collection of photos grouped into "posts" (called a "set" or
 * "album" in most providers) that are in turn assigned categories. Additional
 * library methods are added by the factory.
 */
export default {
   /**
    * Root categories indexed by their name
    */
   categories: {} as {[key:string]:Category},
   /**
    * Flat reference to all posts for simplified lookup
    */
   posts: [] as Post[],
   /**
    * All photo tags in hash[slug] = full name format
    */
   tags: {} as {[key:string]:string},
   loaded: false,
   postInfoLoaded: false,
   /**
    * Track keys of posts and categories that change on library reload
    * (can be used for cache invalidation)
    */
   changedKeys: [] as string[],

   // empty library object before reloading from source
   empty(this:Library) {
      this.loaded = false;
      this.postInfoLoaded = false;
      this.posts = [];
      this.categories = {};
      this.tags = {};
      return this;
   },

   /**
    * Add post to library and link with adjacent posts
    */
   addPost(this:Library, p:Post) {
      // exit if post with same ID is already present
      if (this.posts.filter(e => e.id === p.id).length > 0) { return; }
      this.posts.push(p);
      if (p.chronological && this.posts.length > 1) {
         const next = this.posts[this.posts.length - 2];
         if (next.chronological) {
            p.next = next;
            next.previous = p;
         }
      }
   },

   /**
    * Array of all post keys
    */
   postKeys(this:Library):string[] { return this.posts.map(p => p.key); },

   /**
    * Array of all category keys
    */
   categoryKeys(this:Library, filterList:string[] = []):string[] {
      const keys:string[] = [];

      if (filterList.length > 0) {
         // get keys only for named categories
         if (!is.array(filterList)) { filterList = [filterList]; }
         for (const filterName of filterList) {
            for (const name in this.categories) {
               const category = this.categories[name];
               const subcat = category.getSubcategory(filterName);

               if (name == filterName) {
                  keys.push(category.key);
               } else if (is.value(subcat)) {
                  keys.push(subcat.key);
               }
            }
         }
      } else {
         // get keys for all categories
         for (const name in this.categories) {
            const category = this.categories[name];
            keys.push(category.key);
            category.subcategories.forEach(s => { keys.push(s.key); });
         }
      }
      return keys;
   },

   /**
    * Find category with given key
    */
   categoryWithKey(this:Library, key:string):Category {
      const rootKey = (key.includes('/')) ? key.split('/')[0] : key;

      for (const name in this.categories) {
         const cat = this.categories[name];
         if (cat.key == rootKey) {
            return (key != rootKey) ? cat.getSubcategory(key) : cat;
         }
      }
      return null;
   },

   getPhotos(this:Library):Promise<Photo[]> {
      return Promise
         .all(this.posts.map(p => p.getPhotos()))
         .then(photos => photos.reduce((all, p) => all.concat(p), [] as Photo[]));
   },

   /**
    * Find post with given ID
    */
   postWithID(this:Library, id:string):Post { return this.posts.find(p => p.id == id); },

   /**
    * Find post with given slug
    */
   postWithKey(this:Library, key:string, partKey:string = null):Post {
      if (is.value(partKey)) { key += '/' + partKey; }
      return this.posts.find(p => p.hasKey(key));
   },

   /**
    * Unload particular posts to force refresh from source
    */
   unload(this:Library, keys:string|string[]) {
      if (!is.array(keys)) { keys = [keys]; }
      for (const k of keys) {
         const p = this.postWithKey(k);
         // removing post details will force it to reload on next access
         if (is.value(p)) { p.empty(); }
      }
   },

   /**
    * Remove posts (primarily for testing)
    */
   remove(keys:string|string[]) {
      if (!is.array(keys)) { keys = [keys]; }
      for (const k of keys) {
         const p = this.postWithKey(k);
         if (is.value(p)) {
            this.posts.splice(this.posts.indexOf(p), 1);
            for (const key in this.categories) { this.categories[key].removePost(p); }
         }
      }
   },

   /**
    * Get unique list of tags used on photos in the post and update photo tags
    * to use full names.
    */
   photoTagList(this:Library, photos:Photo[]):string {
      // all photo tags in the post
      const postTags = [];

      for (const p of photos) {
         // tag slugs to remove from photo
         const toRemove = [];

         for (let i = 0; i < p.tags.length; i++) {
            const tagSlug = p.tags[i];
            // lookup full tag name from its slug
            const tagName = this.tags[tagSlug];

            if (is.value(tagName)) {
               // replace tag slug in photo with tag name
               p.tags[i] = tagName;
               if (postTags.indexOf(tagName) == -1) { postTags.push(tagName); }
            } else {
               // remove tag slug from list
               // this can happen if a photo has tags intentionally excluded from the library
               toRemove.push(tagSlug);
            }
         }

         for (const tagSlug of toRemove) {
            const index = p.tags.indexOf(tagSlug);
            if (index >= 0) { p.tags.splice(index, 1); }
         }
      }
      return (postTags.length > 0) ? postTags.join(', ') : null;
   },

   load():Promise<Library> { return null; },
   getEXIF():Promise<EXIF> { return null; },
   getPostWithPhoto():Promise<Post> { return null; },
   getPhotosWithTags():Promise<Photo[]> { return null; }
} as Library;