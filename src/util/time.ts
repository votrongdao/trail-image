import config from '../config';
import { month, weekday } from '../constants';
import is from '../is';
import { leadingZeros, parseNumber } from './number';
import { format } from './text';

/**
 * Return AM or PM
 */
export const hourOfDay = (h:number) => (h > 12) ? 'PM ' + (h - 12) : 'AM ' + h;

/**
 * Format date as Month Day, Year (March 15, 1973)
 */
export const toDateString = (d:Date) => month[d.getMonth()] + ' ' + d.getDate() + ', ' + d.getFullYear();

export function toLogTime(text:string):string {
   const d = new Date(text);
   //var logOffset = d.getTimezoneOffset();
   //var localOffset = (new Date()).getTimezoneOffset();

   // just be dumb for now
   if (config.isProduction) { d.setHours(d.getHours() - 6); }

   return format('{0}/{1} {2} {3}:{4}:{5}.{6}',
      d.getMonth() + 1,
      d.getDate(),
      weekday[d.getDay()],
      hourOfDay(d.getHours()),
      leadingZeros(d.getMinutes(), 2),
      leadingZeros(d.getSeconds(), 2),
      leadingZeros(d.getMilliseconds(), 3)
   );
}

/**
 * Whether daylight savings applies to date
 *
 * http://javascript.about.com/library/bldst.htm
 */
export function inDaylightSavings(date = new Date()):boolean {
   const jan = new Date(date.getFullYear(), 0, 1);
   const jul = new Date(date.getFullYear(), 6, 1);
   const nonDstOffset = Math.max(jan.getTimezoneOffset(), jul.getTimezoneOffset());

   return date.getTimezoneOffset() < nonDstOffset;
}

export const timeZoneOffset = (date = new Date()) => config.timeZone + (inDaylightSavings(date) ? 1 : 0);

/**
 * Convert text to date object. Date constructor uses local time which we
 * need to defeat since local time will be different on host servers. Example:
 *
 *    2012-06-17 17:34:33
 */
export function parseDate(text:string):Date {
   const parts = text.split(' ');
   const date = parts[0].split('-').map(d => parseInt(d));
   const time = parts[1].split(':').map(d => parseInt(d));
   // convert local date to UTC time by adding offset
   const h = time[0] - config.timeZone;
   // date constructor automatically converts to local time
   const d = new Date(Date.UTC(date[0], date[1] - 1, date[2], h, time[1], time[2]));
   if (inDaylightSavings(d)) { d.setHours(d.getHours() - 1); }
   return d;
}

/**
 * Timestamps are created on hosted servers so time zone isn't known
 */
export function fromTimeStamp(timestamp:Date|number|string):Date {
   if (is.date(timestamp)) {
      return timestamp;
   } else if (is.text(timestamp)) {
      timestamp = parseNumber(timestamp);
   }
   return new Date(timestamp * 1000);
}

/**
 * Example 2013-10-02T11:55Z
 *
 * http://en.wikipedia.org/wiki/ISO_8601
 * https://developers.facebook.com/docs/reference/opengraph/object-type/article/
 */
export const iso8601time = (timestamp:number|Date) => fromTimeStamp(timestamp).toISOString();

/**
 * Convert decimal hours to hours:minutes
 */
export function hoursAndMinutes(hours:number):string {
   const h = Math.floor(hours);
   const m = hours - h;

   return h + ':' + leadingZeros(Math.round(60 * m), 2);
}