(function(n,d){function p(){var a=g.elements;return"string"==typeof a?a.split(" "):a}function h(a){var b=q[a[r]];b||(b={},k++,a[r]=k,q[k]=b);return b}function t(a,b,c){b||(b=d);if(f)return b.createElement(a);c||(c=h(b));b=c.cache[a]?c.cache[a].cloneNode():v.test(a)?(c.cache[a]=c.createElem(a)).cloneNode():c.createElem(a);return b.canHaveChildren&&!w.test(a)?c.frag.appendChild(b):b}function x(a,b){b.cache||(b.cache={},b.createElem=a.createElement,b.createFrag=a.createDocumentFragment,b.frag=b.createFrag());
a.createElement=function(c){return g.shivMethods?t(c,a,b):b.createElem(c)};a.createDocumentFragment=Function("h,f","return function(){var n=f.cloneNode(),c=n.createElement;h.shivMethods&&("+p().join().replace(/[\w\-]+/g,function(a){b.createElem(a);b.frag.createElement(a);return'c("'+a+'")'})+");return n}")(g,b.frag)}function u(a){a||(a=d);var b=h(a);if(g.shivCSS&&!l&&!b.hasCSS){var c=a;var e=c.createElement("p");c=c.getElementsByTagName("head")[0]||c.documentElement;e.innerHTML="x<style>article,aside,dialog,figcaption,figure,footer,header,hgroup,main,nav,section{display:block}mark{background:#FF0;color:#000}template{display:none}</style>";
e=c.insertBefore(e.lastChild,c.firstChild);b.hasCSS=!!e}f||x(a,b);return a}var m=n.html5||{},w=/^<|^(?:button|map|select|textarea|object|iframe|option|optgroup)$/i,v=/^(?:a|b|code|div|fieldset|h1|h2|h3|h4|h5|h6|i|label|li|ol|p|q|span|strong|style|table|tbody|td|th|tr|ul)$/i,l,r="_html5shiv",k=0,q={},f;(function(){try{var a=d.createElement("a");a.innerHTML="<xyz></xyz>";l="hidden"in a;var b;if(!(b=1==a.childNodes.length)){d.createElement("a");var c=d.createDocumentFragment();b="undefined"==typeof c.cloneNode||
"undefined"==typeof c.createDocumentFragment||"undefined"==typeof c.createElement}f=b}catch(e){f=l=!0}})();var g={elements:m.elements||"abbr article aside audio bdi canvas data datalist details dialog figcaption figure footer header hgroup main mark meter nav output progress section summary template time video",version:"3.7.0",shivCSS:!1!==m.shivCSS,supportsUnknownElements:f,shivMethods:!1!==m.shivMethods,type:"default",shivDocument:u,createElement:t,createDocumentFragment:function(a,b){a||(a=d);
if(f)return a.createDocumentFragment();b=b||h(a);a=b.frag.cloneNode();b=0;for(var c=p(),e=c.length;b<e;b++)a.createElement(c[b]);return a}};n.html5=g;u(d)})(this,document);
