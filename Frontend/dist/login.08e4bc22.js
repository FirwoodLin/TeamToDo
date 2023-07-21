// modules are defined as an array
// [ module function, map of requires ]
//
// map of requires is short require name -> numeric require
//
// anything defined in a previous bundle is accessed via the
// orig method which is the require for previous bundles

(function (modules, entry, mainEntry, parcelRequireName, globalName) {
  /* eslint-disable no-undef */
  var globalObject =
    typeof globalThis !== 'undefined'
      ? globalThis
      : typeof self !== 'undefined'
      ? self
      : typeof window !== 'undefined'
      ? window
      : typeof global !== 'undefined'
      ? global
      : {};
  /* eslint-enable no-undef */

  // Save the require from previous bundle to this closure if any
  var previousRequire =
    typeof globalObject[parcelRequireName] === 'function' &&
    globalObject[parcelRequireName];

  var cache = previousRequire.cache || {};
  // Do not use `require` to prevent Webpack from trying to bundle this call
  var nodeRequire =
    typeof module !== 'undefined' &&
    typeof module.require === 'function' &&
    module.require.bind(module);

  function newRequire(name, jumped) {
    if (!cache[name]) {
      if (!modules[name]) {
        // if we cannot find the module within our internal map or
        // cache jump to the current global require ie. the last bundle
        // that was added to the page.
        var currentRequire =
          typeof globalObject[parcelRequireName] === 'function' &&
          globalObject[parcelRequireName];
        if (!jumped && currentRequire) {
          return currentRequire(name, true);
        }

        // If there are other bundles on this page the require from the
        // previous one is saved to 'previousRequire'. Repeat this as
        // many times as there are bundles until the module is found or
        // we exhaust the require chain.
        if (previousRequire) {
          return previousRequire(name, true);
        }

        // Try the node require function if it exists.
        if (nodeRequire && typeof name === 'string') {
          return nodeRequire(name);
        }

        var err = new Error("Cannot find module '" + name + "'");
        err.code = 'MODULE_NOT_FOUND';
        throw err;
      }

      localRequire.resolve = resolve;
      localRequire.cache = {};

      var module = (cache[name] = new newRequire.Module(name));

      modules[name][0].call(
        module.exports,
        localRequire,
        module,
        module.exports,
        this
      );
    }

    return cache[name].exports;

    function localRequire(x) {
      var res = localRequire.resolve(x);
      return res === false ? {} : newRequire(res);
    }

    function resolve(x) {
      var id = modules[name][1][x];
      return id != null ? id : x;
    }
  }

  function Module(moduleName) {
    this.id = moduleName;
    this.bundle = newRequire;
    this.exports = {};
  }

  newRequire.isParcelRequire = true;
  newRequire.Module = Module;
  newRequire.modules = modules;
  newRequire.cache = cache;
  newRequire.parent = previousRequire;
  newRequire.register = function (id, exports) {
    modules[id] = [
      function (require, module) {
        module.exports = exports;
      },
      {},
    ];
  };

  Object.defineProperty(newRequire, 'root', {
    get: function () {
      return globalObject[parcelRequireName];
    },
  });

  globalObject[parcelRequireName] = newRequire;

  for (var i = 0; i < entry.length; i++) {
    newRequire(entry[i]);
  }

  if (mainEntry) {
    // Expose entry point to Node, AMD or browser globals
    // Based on https://github.com/ForbesLindesay/umd/blob/master/template.js
    var mainExports = newRequire(mainEntry);

    // CommonJS
    if (typeof exports === 'object' && typeof module !== 'undefined') {
      module.exports = mainExports;

      // RequireJS
    } else if (typeof define === 'function' && define.amd) {
      define(function () {
        return mainExports;
      });

      // <script>
    } else if (globalName) {
      this[globalName] = mainExports;
    }
  }
})({"1i85t":[function(require,module,exports) {
var global = arguments[3];
var HMR_HOST = null;
var HMR_PORT = null;
var HMR_SECURE = false;
var HMR_ENV_HASH = "d6ea1d42532a7575";
module.bundle.HMR_BUNDLE_ID = "1f23704608e4bc22";
"use strict";
/* global HMR_HOST, HMR_PORT, HMR_ENV_HASH, HMR_SECURE, chrome, browser, __parcel__import__, __parcel__importScripts__, ServiceWorkerGlobalScope */ /*::
import type {
  HMRAsset,
  HMRMessage,
} from '@parcel/reporter-dev-server/src/HMRServer.js';
interface ParcelRequire {
  (string): mixed;
  cache: {|[string]: ParcelModule|};
  hotData: {|[string]: mixed|};
  Module: any;
  parent: ?ParcelRequire;
  isParcelRequire: true;
  modules: {|[string]: [Function, {|[string]: string|}]|};
  HMR_BUNDLE_ID: string;
  root: ParcelRequire;
}
interface ParcelModule {
  hot: {|
    data: mixed,
    accept(cb: (Function) => void): void,
    dispose(cb: (mixed) => void): void,
    // accept(deps: Array<string> | string, cb: (Function) => void): void,
    // decline(): void,
    _acceptCallbacks: Array<(Function) => void>,
    _disposeCallbacks: Array<(mixed) => void>,
  |};
}
interface ExtensionContext {
  runtime: {|
    reload(): void,
    getURL(url: string): string;
    getManifest(): {manifest_version: number, ...};
  |};
}
declare var module: {bundle: ParcelRequire, ...};
declare var HMR_HOST: string;
declare var HMR_PORT: string;
declare var HMR_ENV_HASH: string;
declare var HMR_SECURE: boolean;
declare var chrome: ExtensionContext;
declare var browser: ExtensionContext;
declare var __parcel__import__: (string) => Promise<void>;
declare var __parcel__importScripts__: (string) => Promise<void>;
declare var globalThis: typeof self;
declare var ServiceWorkerGlobalScope: Object;
*/ var OVERLAY_ID = "__parcel__error__overlay__";
var OldModule = module.bundle.Module;
function Module(moduleName) {
    OldModule.call(this, moduleName);
    this.hot = {
        data: module.bundle.hotData[moduleName],
        _acceptCallbacks: [],
        _disposeCallbacks: [],
        accept: function(fn) {
            this._acceptCallbacks.push(fn || function() {});
        },
        dispose: function(fn) {
            this._disposeCallbacks.push(fn);
        }
    };
    module.bundle.hotData[moduleName] = undefined;
}
module.bundle.Module = Module;
module.bundle.hotData = {};
var checkedAssets /*: {|[string]: boolean|} */ , assetsToDispose /*: Array<[ParcelRequire, string]> */ , assetsToAccept /*: Array<[ParcelRequire, string]> */ ;
function getHostname() {
    return HMR_HOST || (location.protocol.indexOf("http") === 0 ? location.hostname : "localhost");
}
function getPort() {
    return HMR_PORT || location.port;
}
// eslint-disable-next-line no-redeclare
var parent = module.bundle.parent;
if ((!parent || !parent.isParcelRequire) && typeof WebSocket !== "undefined") {
    var hostname = getHostname();
    var port = getPort();
    var protocol = HMR_SECURE || location.protocol == "https:" && !/localhost|127.0.0.1|0.0.0.0/.test(hostname) ? "wss" : "ws";
    var ws = new WebSocket(protocol + "://" + hostname + (port ? ":" + port : "") + "/");
    // Web extension context
    var extCtx = typeof chrome === "undefined" ? typeof browser === "undefined" ? null : browser : chrome;
    // Safari doesn't support sourceURL in error stacks.
    // eval may also be disabled via CSP, so do a quick check.
    var supportsSourceURL = false;
    try {
        (0, eval)('throw new Error("test"); //# sourceURL=test.js');
    } catch (err) {
        supportsSourceURL = err.stack.includes("test.js");
    }
    // $FlowFixMe
    ws.onmessage = async function(event /*: {data: string, ...} */ ) {
        checkedAssets = {} /*: {|[string]: boolean|} */ ;
        assetsToAccept = [];
        assetsToDispose = [];
        var data /*: HMRMessage */  = JSON.parse(event.data);
        if (data.type === "update") {
            // Remove error overlay if there is one
            if (typeof document !== "undefined") removeErrorOverlay();
            let assets = data.assets.filter((asset)=>asset.envHash === HMR_ENV_HASH);
            // Handle HMR Update
            let handled = assets.every((asset)=>{
                return asset.type === "css" || asset.type === "js" && hmrAcceptCheck(module.bundle.root, asset.id, asset.depsByBundle);
            });
            if (handled) {
                console.clear();
                // Dispatch custom event so other runtimes (e.g React Refresh) are aware.
                if (typeof window !== "undefined" && typeof CustomEvent !== "undefined") window.dispatchEvent(new CustomEvent("parcelhmraccept"));
                await hmrApplyUpdates(assets);
                // Dispose all old assets.
                let processedAssets = {} /*: {|[string]: boolean|} */ ;
                for(let i = 0; i < assetsToDispose.length; i++){
                    let id = assetsToDispose[i][1];
                    if (!processedAssets[id]) {
                        hmrDispose(assetsToDispose[i][0], id);
                        processedAssets[id] = true;
                    }
                }
                // Run accept callbacks. This will also re-execute other disposed assets in topological order.
                processedAssets = {};
                for(let i = 0; i < assetsToAccept.length; i++){
                    let id = assetsToAccept[i][1];
                    if (!processedAssets[id]) {
                        hmrAccept(assetsToAccept[i][0], id);
                        processedAssets[id] = true;
                    }
                }
            } else fullReload();
        }
        if (data.type === "error") {
            // Log parcel errors to console
            for (let ansiDiagnostic of data.diagnostics.ansi){
                let stack = ansiDiagnostic.codeframe ? ansiDiagnostic.codeframe : ansiDiagnostic.stack;
                console.error("\uD83D\uDEA8 [parcel]: " + ansiDiagnostic.message + "\n" + stack + "\n\n" + ansiDiagnostic.hints.join("\n"));
            }
            if (typeof document !== "undefined") {
                // Render the fancy html overlay
                removeErrorOverlay();
                var overlay = createErrorOverlay(data.diagnostics.html);
                // $FlowFixMe
                document.body.appendChild(overlay);
            }
        }
    };
    ws.onerror = function(e) {
        console.error(e.message);
    };
    ws.onclose = function() {
        console.warn("[parcel] \uD83D\uDEA8 Connection to the HMR server was lost");
    };
}
function removeErrorOverlay() {
    var overlay = document.getElementById(OVERLAY_ID);
    if (overlay) {
        overlay.remove();
        console.log("[parcel] ✨ Error resolved");
    }
}
function createErrorOverlay(diagnostics) {
    var overlay = document.createElement("div");
    overlay.id = OVERLAY_ID;
    let errorHTML = '<div style="background: black; opacity: 0.85; font-size: 16px; color: white; position: fixed; height: 100%; width: 100%; top: 0px; left: 0px; padding: 30px; font-family: Menlo, Consolas, monospace; z-index: 9999;">';
    for (let diagnostic of diagnostics){
        let stack = diagnostic.frames.length ? diagnostic.frames.reduce((p, frame)=>{
            return `${p}
<a href="/__parcel_launch_editor?file=${encodeURIComponent(frame.location)}" style="text-decoration: underline; color: #888" onclick="fetch(this.href); return false">${frame.location}</a>
${frame.code}`;
        }, "") : diagnostic.stack;
        errorHTML += `
      <div>
        <div style="font-size: 18px; font-weight: bold; margin-top: 20px;">
          🚨 ${diagnostic.message}
        </div>
        <pre>${stack}</pre>
        <div>
          ${diagnostic.hints.map((hint)=>"<div>\uD83D\uDCA1 " + hint + "</div>").join("")}
        </div>
        ${diagnostic.documentation ? `<div>📝 <a style="color: violet" href="${diagnostic.documentation}" target="_blank">Learn more</a></div>` : ""}
      </div>
    `;
    }
    errorHTML += "</div>";
    overlay.innerHTML = errorHTML;
    return overlay;
}
function fullReload() {
    if ("reload" in location) location.reload();
    else if (extCtx && extCtx.runtime && extCtx.runtime.reload) extCtx.runtime.reload();
}
function getParents(bundle, id) /*: Array<[ParcelRequire, string]> */ {
    var modules = bundle.modules;
    if (!modules) return [];
    var parents = [];
    var k, d, dep;
    for(k in modules)for(d in modules[k][1]){
        dep = modules[k][1][d];
        if (dep === id || Array.isArray(dep) && dep[dep.length - 1] === id) parents.push([
            bundle,
            k
        ]);
    }
    if (bundle.parent) parents = parents.concat(getParents(bundle.parent, id));
    return parents;
}
function updateLink(link) {
    var href = link.getAttribute("href");
    if (!href) return;
    var newLink = link.cloneNode();
    newLink.onload = function() {
        if (link.parentNode !== null) // $FlowFixMe
        link.parentNode.removeChild(link);
    };
    newLink.setAttribute("href", // $FlowFixMe
    href.split("?")[0] + "?" + Date.now());
    // $FlowFixMe
    link.parentNode.insertBefore(newLink, link.nextSibling);
}
var cssTimeout = null;
function reloadCSS() {
    if (cssTimeout) return;
    cssTimeout = setTimeout(function() {
        var links = document.querySelectorAll('link[rel="stylesheet"]');
        for(var i = 0; i < links.length; i++){
            // $FlowFixMe[incompatible-type]
            var href /*: string */  = links[i].getAttribute("href");
            var hostname = getHostname();
            var servedFromHMRServer = hostname === "localhost" ? new RegExp("^(https?:\\/\\/(0.0.0.0|127.0.0.1)|localhost):" + getPort()).test(href) : href.indexOf(hostname + ":" + getPort());
            var absolute = /^https?:\/\//i.test(href) && href.indexOf(location.origin) !== 0 && !servedFromHMRServer;
            if (!absolute) updateLink(links[i]);
        }
        cssTimeout = null;
    }, 50);
}
function hmrDownload(asset) {
    if (asset.type === "js") {
        if (typeof document !== "undefined") {
            let script = document.createElement("script");
            script.src = asset.url + "?t=" + Date.now();
            if (asset.outputFormat === "esmodule") script.type = "module";
            return new Promise((resolve, reject)=>{
                var _document$head;
                script.onload = ()=>resolve(script);
                script.onerror = reject;
                (_document$head = document.head) === null || _document$head === void 0 || _document$head.appendChild(script);
            });
        } else if (typeof importScripts === "function") {
            // Worker scripts
            if (asset.outputFormat === "esmodule") return import(asset.url + "?t=" + Date.now());
            else return new Promise((resolve, reject)=>{
                try {
                    importScripts(asset.url + "?t=" + Date.now());
                    resolve();
                } catch (err) {
                    reject(err);
                }
            });
        }
    }
}
async function hmrApplyUpdates(assets) {
    global.parcelHotUpdate = Object.create(null);
    let scriptsToRemove;
    try {
        // If sourceURL comments aren't supported in eval, we need to load
        // the update from the dev server over HTTP so that stack traces
        // are correct in errors/logs. This is much slower than eval, so
        // we only do it if needed (currently just Safari).
        // https://bugs.webkit.org/show_bug.cgi?id=137297
        // This path is also taken if a CSP disallows eval.
        if (!supportsSourceURL) {
            let promises = assets.map((asset)=>{
                var _hmrDownload;
                return (_hmrDownload = hmrDownload(asset)) === null || _hmrDownload === void 0 ? void 0 : _hmrDownload.catch((err)=>{
                    // Web extension bugfix for Chromium
                    // https://bugs.chromium.org/p/chromium/issues/detail?id=1255412#c12
                    if (extCtx && extCtx.runtime && extCtx.runtime.getManifest().manifest_version == 3) {
                        if (typeof ServiceWorkerGlobalScope != "undefined" && global instanceof ServiceWorkerGlobalScope) {
                            extCtx.runtime.reload();
                            return;
                        }
                        asset.url = extCtx.runtime.getURL("/__parcel_hmr_proxy__?url=" + encodeURIComponent(asset.url + "?t=" + Date.now()));
                        return hmrDownload(asset);
                    }
                    throw err;
                });
            });
            scriptsToRemove = await Promise.all(promises);
        }
        assets.forEach(function(asset) {
            hmrApply(module.bundle.root, asset);
        });
    } finally{
        delete global.parcelHotUpdate;
        if (scriptsToRemove) scriptsToRemove.forEach((script)=>{
            if (script) {
                var _document$head2;
                (_document$head2 = document.head) === null || _document$head2 === void 0 || _document$head2.removeChild(script);
            }
        });
    }
}
function hmrApply(bundle /*: ParcelRequire */ , asset /*:  HMRAsset */ ) {
    var modules = bundle.modules;
    if (!modules) return;
    if (asset.type === "css") reloadCSS();
    else if (asset.type === "js") {
        let deps = asset.depsByBundle[bundle.HMR_BUNDLE_ID];
        if (deps) {
            if (modules[asset.id]) {
                // Remove dependencies that are removed and will become orphaned.
                // This is necessary so that if the asset is added back again, the cache is gone, and we prevent a full page reload.
                let oldDeps = modules[asset.id][1];
                for(let dep in oldDeps)if (!deps[dep] || deps[dep] !== oldDeps[dep]) {
                    let id = oldDeps[dep];
                    let parents = getParents(module.bundle.root, id);
                    if (parents.length === 1) hmrDelete(module.bundle.root, id);
                }
            }
            if (supportsSourceURL) // Global eval. We would use `new Function` here but browser
            // support for source maps is better with eval.
            (0, eval)(asset.output);
            // $FlowFixMe
            let fn = global.parcelHotUpdate[asset.id];
            modules[asset.id] = [
                fn,
                deps
            ];
        } else if (bundle.parent) hmrApply(bundle.parent, asset);
    }
}
function hmrDelete(bundle, id) {
    let modules = bundle.modules;
    if (!modules) return;
    if (modules[id]) {
        // Collect dependencies that will become orphaned when this module is deleted.
        let deps = modules[id][1];
        let orphans = [];
        for(let dep in deps){
            let parents = getParents(module.bundle.root, deps[dep]);
            if (parents.length === 1) orphans.push(deps[dep]);
        }
        // Delete the module. This must be done before deleting dependencies in case of circular dependencies.
        delete modules[id];
        delete bundle.cache[id];
        // Now delete the orphans.
        orphans.forEach((id)=>{
            hmrDelete(module.bundle.root, id);
        });
    } else if (bundle.parent) hmrDelete(bundle.parent, id);
}
function hmrAcceptCheck(bundle /*: ParcelRequire */ , id /*: string */ , depsByBundle /*: ?{ [string]: { [string]: string } }*/ ) {
    if (hmrAcceptCheckOne(bundle, id, depsByBundle)) return true;
    // Traverse parents breadth first. All possible ancestries must accept the HMR update, or we'll reload.
    let parents = getParents(module.bundle.root, id);
    let accepted = false;
    while(parents.length > 0){
        let v = parents.shift();
        let a = hmrAcceptCheckOne(v[0], v[1], null);
        if (a) // If this parent accepts, stop traversing upward, but still consider siblings.
        accepted = true;
        else {
            // Otherwise, queue the parents in the next level upward.
            let p = getParents(module.bundle.root, v[1]);
            if (p.length === 0) {
                // If there are no parents, then we've reached an entry without accepting. Reload.
                accepted = false;
                break;
            }
            parents.push(...p);
        }
    }
    return accepted;
}
function hmrAcceptCheckOne(bundle /*: ParcelRequire */ , id /*: string */ , depsByBundle /*: ?{ [string]: { [string]: string } }*/ ) {
    var modules = bundle.modules;
    if (!modules) return;
    if (depsByBundle && !depsByBundle[bundle.HMR_BUNDLE_ID]) {
        // If we reached the root bundle without finding where the asset should go,
        // there's nothing to do. Mark as "accepted" so we don't reload the page.
        if (!bundle.parent) return true;
        return hmrAcceptCheck(bundle.parent, id, depsByBundle);
    }
    if (checkedAssets[id]) return true;
    checkedAssets[id] = true;
    var cached = bundle.cache[id];
    assetsToDispose.push([
        bundle,
        id
    ]);
    if (!cached || cached.hot && cached.hot._acceptCallbacks.length) {
        assetsToAccept.push([
            bundle,
            id
        ]);
        return true;
    }
}
function hmrDispose(bundle /*: ParcelRequire */ , id /*: string */ ) {
    var cached = bundle.cache[id];
    bundle.hotData[id] = {};
    if (cached && cached.hot) cached.hot.data = bundle.hotData[id];
    if (cached && cached.hot && cached.hot._disposeCallbacks.length) cached.hot._disposeCallbacks.forEach(function(cb) {
        cb(bundle.hotData[id]);
    });
    delete bundle.cache[id];
}
function hmrAccept(bundle /*: ParcelRequire */ , id /*: string */ ) {
    // Execute the module.
    bundle(id);
    // Run the accept callbacks in the new version of the module.
    var cached = bundle.cache[id];
    if (cached && cached.hot && cached.hot._acceptCallbacks.length) cached.hot._acceptCallbacks.forEach(function(cb) {
        var assetsToAlsoAccept = cb(function() {
            return getParents(module.bundle.root, id);
        });
        if (assetsToAlsoAccept && assetsToAccept.length) {
            assetsToAlsoAccept.forEach(function(a) {
                hmrDispose(a[0], a[1]);
            });
            // $FlowFixMe[method-unbinding]
            assetsToAccept.push.apply(assetsToAccept, assetsToAlsoAccept);
        }
    });
}

},{}],"8mNOi":[function(require,module,exports) {
//global variable
var parcelHelpers = require("@parcel/transformer-js/src/esmodule-helpers.js");
parcelHelpers.defineInteropFlag(exports);
//export
parcelHelpers.export(exports, "token", ()=>token);
parcelHelpers.export(exports, "currentGroupID", ()=>currentGroupID);
parcelHelpers.export(exports, "currentUserID", ()=>currentUserID);
parcelHelpers.export(exports, "currentUserName", ()=>currentUserName);
parcelHelpers.export(exports, "currentUserEmail", ()=>currentUserEmail);
parcelHelpers.export(exports, "currentUserAvatar", ()=>currentUserAvatar);
parcelHelpers.export(exports, "UserID", ()=>UserID);
parcelHelpers.export(exports, "getGroupIDs", ()=>getGroupIDs);
parcelHelpers.export(exports, "getGroupNames", ()=>getGroupNames);
parcelHelpers.export(exports, "getGroupMemberNames", ()=>getGroupMemberNames);
parcelHelpers.export(exports, "getGroupMemberIDs", ()=>getGroupMemberIDs);
parcelHelpers.export(exports, "getGroupMemberAvatars", ()=>getGroupMemberAvatars);
parcelHelpers.export(exports, "getTaskNames", ()=>getTaskNames);
parcelHelpers.export(exports, "getTaskIDs", ()=>getTaskIDs);
parcelHelpers.export(exports, "getTaskDescriptions", ()=>getTaskDescriptions);
parcelHelpers.export(exports, "getTaskStatuses", ()=>getTaskStatuses);
parcelHelpers.export(exports, "getTaskDeadlines", ()=>getTaskDeadlines);
parcelHelpers.export(exports, "getTaskStartAts", ()=>getTaskStartAts);
parcelHelpers.export(exports, "convertDateTimeFormat", ()=>convertDateTimeFormat);
parcelHelpers.export(exports, "formatDateTimeLocal", ()=>formatDateTimeLocal);
parcelHelpers.export(exports, "updateSelectOptions", ()=>updateSelectOptions);
parcelHelpers.export(exports, "updateGroupMembersList", ()=>updateGroupMembersList);
parcelHelpers.export(exports, "updateTaskList", ()=>updateTaskList);
parcelHelpers.export(exports, "getUserRole", ()=>getUserRole);
parcelHelpers.export(exports, "formatDateTimeLocalToClient", ()=>formatDateTimeLocalToClient);
var token;
var currentGroupID = initCurrentGroupID();
var currentUserID, currentUserName, currentUserEmail, currentUserAvatar;
var UserID;
//global function
// 查询所有群组ID
function getGroupIDs() {
    fetch("http://localhost:8080/api/groups", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + token
        }
    }).then((response)=>{
        if (!response.ok) throw new Error("HTTP error " + response.status);
        return response.json();
    }).then((json)=>{
        if (json.success) {
            let groupIDs = json.data.groups.map((group)=>group.groupID);
            return groupIDs;
        } else throw new Error(json.hint);
    }).catch((error)=>{
        console.error("An error occurred:", error);
    });
}
// 查询所有群组名
function getGroupNames() {
    fetch("http://localhost:8080/api/groups", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + token
        }
    }).then((response)=>{
        if (!response.ok) throw new Error("HTTP error " + response.status);
        return response.json();
    }).then((json)=>{
        if (json.success) {
            let groupNames = json.data.groups.map((group)=>group.groupName);
            return groupNames;
        } else throw new Error(json.hint);
    }).catch((error)=>{
        console.error("An error occurred:", error);
    });
}
// 查询群组所有成员名
function getGroupMemberNames(groupID) {
    fetch(`http://localhost:8080/api/groups/${groupID}/members`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var userNames = data.data.map((item)=>item.User.userName);
            return userNames;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
// 查询群组所有成员ID
function getGroupMemberIDs(groupID) {
    fetch(`http://localhost:8080/api/groups/${groupID}/members`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var userIDs1 = data.data.map((item)=>item.User.userID);
            return userIDs1;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
// 查询群组所有成员Avatar
function getGroupMemberAvatars(groupID) {
    fetch(`http://localhost:8080/api/groups/${groupID}/members`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var userAvatars = data.data.map((item)=>item.User.userAvatar);
            return userIDs;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
//可以重载,依据传递的参数数量不同,执行不同的操作,后面的其他信息也是同理
function getTaskNames(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskNames = data.data.tasks.map((task)=>task.taskName);
            return taskNames;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function getTaskIDs(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskIDs = data.data.tasks.map((task)=>task.taskID);
            return taskIDs;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function getTaskDescriptions(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskDescriptions = data.data.tasks.map((task)=>task.description);
            return taskDescriptions;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function getTaskStatuses(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskStatuses = data.data.tasks.map((task)=>task.taskStatus);
            return taskStatuses;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function getTaskDeadlines(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskDeadlines1 = data.data.tasks.map((task)=>task.deadline);
            taskDeadlines1 = taskDeadlines1.map((deadline)=>convertDateTimeFormat(deadline));
            return taskDeadlines1;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function getTaskStartAts(groupID, userID) {
    fetch(`http://localhost:8080/api/tasks?groupID=${groupID}&userID=${userID}`, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then((response)=>{
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        return response.json();
    }).then((data)=>{
        if (data.success) {
            var taskStartAts = data.data.tasks.map((task)=>task.startAt);
            taskStartAts = taskStartAts.map((startAt)=>convertDateTimeFormat(startAt));
            return taskDeadlines;
        } else console.log("请求未成功: " + data.hint);
    }).catch((e)=>{
        console.log("请求出错: " + e);
    });
}
function convertDateTimeFormat(inputDateTime) {
    var date = new Date(inputDateTime);
    var formattedDate = date.toLocaleString("en-GB", {
        hour: "2-digit",
        minute: "2-digit",
        year: "numeric",
        month: "2-digit",
        day: "2-digit"
    });
    var parts = formattedDate.split(",");
    var outputDateTime = parts[1].trim() + ", " + parts[0].trim();
    return outputDateTime;
}
function formatDateTimeLocal(inputDateTimeLocal) {
    // Create a Date object from the input string
    var date = new Date(inputDateTimeLocal);
    // Format the date and time parts
    var year = date.getFullYear();
    var month = (date.getMonth() + 1).toString().padStart(2, "0");
    var day = date.getDate().toString().padStart(2, "0");
    var hours = date.getHours().toString().padStart(2, "0");
    var minutes = date.getMinutes().toString().padStart(2, "0");
    var seconds = date.getSeconds().toString().padStart(2, "0");
    var milliseconds = date.getMilliseconds().toString().padStart(3, "0");
    // Format the timezone offset
    var offset = -date.getTimezoneOffset();
    var offsetSign = offset >= 0 ? "+" : "-";
    offset = Math.abs(offset);
    var offsetHours = Math.floor(offset / 60).toString().padStart(2, "0");
    var offsetMinutes = (offset % 60).toString().padStart(2, "0");
    var outputDateTime = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${milliseconds}${offsetSign}${offsetHours}:${offsetMinutes}`;
    return outputDateTime;
}
function formatDateTimeLocalToClient(dateTimeStr) {
    var date = new Date(dateTimeStr);
    var day = ("0" + date.getDate()).slice(-2);
    var month = ("0" + (date.getMonth() + 1)).slice(-2);
    var year = date.getFullYear();
    var hours = ("0" + date.getHours()).slice(-2);
    var minutes = ("0" + date.getMinutes()).slice(-2);
    var formattedDate = hours + ":" + minutes + ", " + day + "/" + month + "/" + year;
    return formattedDate;
}
function updateSelectOptions() {
    let groupNames = getGroupNames();
    let groupIDs = getGroupIDs();
    let select = document.getElementById("teams");
    select.innerHTML = "";
    for(let i = 0; i < groupNames.length; i++){
        let option = document.createElement("option");
        option.value = groupIDs[i];
        option.text = groupNames[i];
        select.appendChild(option);
    }
}
function updateGroupMembersList(groupID) {
    // 调用相应的函数获取新的成员名称，ID 和头像
    let memberNames = getGroupMemberNames(groupID);
    let memberIDs = getGroupMemberIDs(groupID);
    let memberAvatars = getGroupMemberAvatars(groupID);
    let list = document.querySelector(".instance-parent");
    list.innerHTML = "";
    for(let i = 0; i < memberNames.length; i++){
        let listItem = document.createElement("li");
        listItem.className = "user-instance";
        let userInfo = document.createElement("div");
        userInfo.className = "user-info";
        let avatar = document.createElement("img");
        avatar.className = "user-avatar";
        avatar.src = memberAvatars[i];
        avatar.alt = "用户头像";
        let name = document.createElement("div");
        name.className = "user-name";
        name.textContent = memberNames[i];
        name.value = memberIDs[i];
        let deleteUser = document.createElement("div");
        deleteUser.className = "delete-user";
        userInfo.appendChild(avatar);
        userInfo.appendChild(name);
        userInfo.appendChild(deleteUser);
        listItem.appendChild(userInfo);
        list.appendChild(listItem);
    }
}
function updateTaskList(groupID, userID) {
    let taskNames = getTaskNames(groupID, userID);
    let taskIDs = getTaskIDs(groupID, userID);
    let taskStatuses = getTaskStatuses(groupID, userID);
    let taskDeadlines1 = getTaskDeadlines(groupID, userID);
    let taskStartAts = getTaskStartAts(groupID, userID);
    let taskList = document.querySelector(".task-list");
    while(taskList.firstChild)taskList.removeChild(taskList.firstChild);
    for(let i = 0; i < taskNames.length; i++){
        let li = document.createElement("li");
        li.className = "task";
        let div = document.createElement("div");
        div.className = "delete-task";
        li.appendChild(div);
        let p1 = document.createElement("p");
        p1.className = "item1";
        p1.textContent = taskNames[i];
        p1.value = taskIDs[i]; // Set taskID as value
        li.appendChild(p1);
        let p2 = document.createElement("p");
        p2.className = "item2";
        p2.textContent = taskStartAts[i];
        li.appendChild(p2);
        let p3 = document.createElement("p");
        p3.className = "item3";
        p3.textContent = taskDeadlines1[i];
        li.appendChild(p3);
        let p4 = document.createElement("p");
        p4.className = "item4";
        p4.textContent = taskStatuses[i] ? "已完成" : "未完成";
        p4.value = taskStatuses[i]; // Set taskStatus as value
        li.appendChild(p4);
        taskList.appendChild(li);
    }
}
function initCurrentGroupID() {
    let firstGroupID = getGroupIDs()[0];
    return firstGroupID;
}
// 查询成员在某群组的身份
async function getUserRole(groupID) {
    try {
        var response = await fetch(`http://localhost:8080/api/groups/${groupID}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        var data = await response.json();
        if (response.ok) {
            var userRole = data.data.role;
            console.log("用户的身份是: " + userRole);
            return userRole;
        } else alert("查询失败: " + data.hint);
    } catch (error) {
        console.error("Error querying user role:" + error);
    }
}

},{"@parcel/transformer-js/src/esmodule-helpers.js":"gkKU3"}],"gkKU3":[function(require,module,exports) {
exports.interopDefault = function(a) {
    return a && a.__esModule ? a : {
        default: a
    };
};
exports.defineInteropFlag = function(a) {
    Object.defineProperty(a, "__esModule", {
        value: true
    });
};
exports.exportAll = function(source, dest) {
    Object.keys(source).forEach(function(key) {
        if (key === "default" || key === "__esModule" || dest.hasOwnProperty(key)) return;
        Object.defineProperty(dest, key, {
            enumerable: true,
            get: function() {
                return source[key];
            }
        });
    });
    return dest;
};
exports.export = function(dest, destName, get) {
    Object.defineProperty(dest, destName, {
        enumerable: true,
        get: get
    });
};

},{}]},["1i85t","8mNOi"], "8mNOi", "parcelRequire2914")

//# sourceMappingURL=login.08e4bc22.js.map
