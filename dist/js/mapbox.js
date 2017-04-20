"use strict";$(function(){function o(o,t,e){o?(j.keyNav=function(o){switch((window.event?window.event:o).keyCode){case 27:j.mapInteraction(o);break;case 37:e();break;case 39:t()}},document.addEventListener("keydown",j.keyNav)):document.removeEventListener("keydown",j.keyNav)}function t(o,t){var n=w.getZoom(),a=1.5*n/Math.pow(2,1.3*n)*10,r=[o.lng-a,o.lat-a],i=[o.lng+a,o.lat+a],c=z.features.filter(function(o){var t=o.geometry.coordinates;return t[0]>=r[0]&&t[1]>=r[1]&&t[0]<=i[0]&&t[1]<=i[1]}).map(function(t){return t.properties.distance=e(o,t.geometry.coordinates),t});return c.sort(function(o,t){return o.properties.distance-t.properties.distance}),c.slice(0,t)}function e(o,t){var e=o.lng,n=o.lat,a=t[0],r=t[1];return Math.sqrt(Math.pow(a-e,2)+Math.pow(r-n,2))}function n(){var o=w.getCenter(),t="/map?lat="+o.lat+"&lon="+o.lng+"&zoom="+w.getZoom();window.history.replaceState(null,null,t)}function a(o){o!=l&&(l=o,l?(g.on.show(),g.off.hide()):(g.on.hide(),g.off.show()),w.once("data",function(o){"style"==o.dataType&&s()}),w.setStyle("mapbox://styles/"+(l?d.imagery:d.basic)))}function r(){w.getZoom()>p.zoom&&!C?(C=!0,m.click(function(){w.easeTo(p)}).removeClass("disabled")):w.getZoom()<=p.zoom&&C&&(C=!1,m.off("click").addClass("disabled"))}function i(o){return void 0===o&&(o=""),function(){k.style.cursor=o}}function c(o){var t=o.split("/"),e=t[t.length-1].split("_");window.location="/"+e[0]}function s(){w.addSource("photos",{type:"geojson",data:z,cluster:!0,clusterMaxZoom:18,clusterRadius:30}),w.addLayer({id:"cluster",type:"circle",source:"photos",filter:["has","point_count"],paint:{"circle-color":"#524948","circle-radius":{property:"point_count",type:"interval",stops:[[0,13],[50,16],[100,19]]},"circle-opacity":b,"circle-stroke-width":3,"circle-stroke-color":"#ccc"}}),w.addLayer({id:"cluster-count",type:"symbol",source:"photos",filter:["has","point_count"],layout:{"text-field":"{point_count_abbreviated}","text-font":["Open Sans Bold","Arial Unicode MS Bold"],"text-size":14},paint:{"text-color":"#fff"}}),w.addLayer({id:"photo",type:"circle",source:"photos",filter:["!has","point_count"],paint:{"circle-color":"#f00","circle-radius":7,"circle-stroke-width":4,"circle-stroke-color":"#fdd","circle-opacity":b}}),w.on("mouseenter","cluster",i("pointer")).on("mouseleave","cluster",i()).on("mouseenter","photo",i("pointer")).on("mouseleave","photo",i()).on("move",j.mapInteraction).on("click",j.mapInteraction).on("zoomend",j.zoomEnd).on("moveend",n).on("mousedown","cluster",j.clusterClick).on("mousedown","photo",j.photoClick)}var l=!1,p={zoom:6.5,center:[-116.0987,44.7]},d={basic:"jabbott7/cj1k069f0000p2slh5775akgj",hybrid:"jabbott7/cj1mcsd7t000h2rpitaiafuq0",imagery:"jabbott7/cj1mcsd7t000h2rpitaiafuq0",outdoors:"mapbox/outdoors-v10"},u=$("#photo-count"),f=$("#photo-preview"),h=$("#toggle-satellite"),m=$("#zoom-out"),g={on:$("nav .glyphicon-check"),off:$("nav .glyphicon-unchecked")},y=function(){for(var o=window.location.search.split(/[&\?]/g),t={},e=0;e<o.length;e++){var n=o[e].split("=");2==n.length&&(t[n[0]]=parseFloat(n[1]))}return t.hasOwnProperty("lat")&&t.hasOwnProperty("lon")&&(t.center=[t.lon,t.lat]),t}(),v=new mapboxgl.NavigationControl,w=new mapboxgl.Map({container:"map-canvas",style:"mapbox://styles/"+d.outdoors,center:y.center||p.center,zoom:y.zoom||p.zoom,maxZoom:18,dragRotate:!1,keyboard:!1}),k=w.getCanvasContainer(),b=.6,C=!1,z=null,x={icon:function(o,t){var e=$("<span>").addClass("glyphicon glyphicon-"+o);return void 0!==t&&e.click(t),e},coordinate:function(o){var t=Math.pow(10,5),e=function(o){return Math.round(o*t)/t};return e(o[1])+", "+e(o[0])},photo:function(o){var t=o.properties,e="Click or tap to enlarge";return $("<figure>").append($("<img>").attr("src",t.url).attr("title",e).attr("alt",e).click(function(){c(t.url)})).append($("<figcaption>").html(this.coordinate(o.geometry.coordinates)))}},j={zoomEnd:function(){n(),r()},keyNav:null,mapInteraction:function(){f.hide(),o(!1)},photoClick:function(o){f.empty().css({top:o.point.y+15,left:o.point.x}).append(x.photo(o.features[0])).show()},clusterClick:function(e){var n=e.features[0].properties,a=w.getZoom(),r=function(){w.easeTo({center:e.lngLat,zoom:18-a<2?18:a+2})};if(n.point_count>10&&a<16)r();else{var i=t(e.lngLat,n.point_count);if(0==i.length)r();else{var c=1,s=$("<div>").addClass("photo-list"),l=$("<div>").addClass("markers"),p=function(o){c+=o,c>i.length?c=1:c<1&&(c=i.length),$("figure",s).hide(),$("span",l).removeClass("selected"),$("figure:nth-child("+c+")",s).show(),$("span:nth-child("+c+")",l).addClass("selected")},d=function(){p(-1)},u=function(){p(1)};o(!0,u,d);for(var h=0;h<i.length;h++)s.append(x.photo(i[h])),l.append(x.icon("map-marker"));$("span:first-child",l).addClass("selected"),f.empty().css({top:e.point.y+15,left:e.point.x}).append($("<nav>").append(x.icon("arrow-left",d)).append(l).append(x.icon("arrow-right",u))).append(s).show()}}}};y.center&&r(),w.addControl(v,"top-right").on("load",function(){h.click(function(){a(!l)}),$.getJSON("/geo.json",function(o){z=o,u.html(z.features.length+" photos").show(),s(),w.resize()})})});
//# sourceMappingURL=/js/maps/mapbox.js.map
