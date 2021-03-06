import is from '../is';
import {
   format,
   rot13,
   IPv6,
   slug,
   titleCase,
   capitalize,
   topDomain,
   encodeBase64,
   decodeBase64 } from './text';
import {
   typography,
   characterEntities,
   caption,
   fixMalformedLink,
   fraction,
   photoTagList,
   shortenLinkText,
   story,
   iconTag,
   postModeIcon,
   postCategoryIcon,
   logMessage } from './html';
import {
   fromTimeStamp,
   hoursAndMinutes,
   hourOfDay,
   inDaylightSavings,
   iso8601time,
   parseDate,
   timeZoneOffset,
   toDateString,
   toLogTime } from './time';
import {
   sayNumber,
   leadingZeros,
   parseNumber,
   maybeNumber } from './number';

export default {
   format,
   logMessage,
   topDomain,
   IPv6,
   slug,
   capitalize,
   titleCase,

   date: {
      fromTimeStamp,
      hoursAndMinutes,
      hourOfDay,
      inDaylightSavings,
      iso8601time,
      parse: parseDate,
      timeZoneOffset,
      toString: toDateString,
      toLogTime

   },

   encode: {
      rot13,
      toBase64: encodeBase64,
      fromBase64: decodeBase64,
      characterEntities
   },

   html: {
      typography,
      caption,
      fixMalformedLink,
      fraction,
      photoTagList,
      shortenLinkText,
      story
   },

   icon: {
      tag: iconTag,
      mode: postModeIcon,
      category: postCategoryIcon
   },

   number: {
      say: sayNumber,
      pad: leadingZeros,
      parse: parseNumber,
      maybe: maybeNumber
   },

   /**
    * Shuffle an array
    *
    * http://sroucheray.org/blog/2009/11/array-sort-should-not-be-used-to-shuffle-an-array/
    */
   shuffle<T>(a:T[]):T[] {
      if (!is.array(a) || a.length === 0) { return null; }

      let i = a.length;

      while (--i) {
         const j = Math.floor(Math.random() * (i + 1));
         const temp = a[i];
         a[i] = a[j];
         a[j] = temp;
      }
      return a;
   }
};

