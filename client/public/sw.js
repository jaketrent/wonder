console.log('sw.js...')

const CACHE_NAME = 'wonder-cache-v1'

const urlsToCache = [
  '/',
  '/global.css',
  '/bundle.css',
  '/bundle.js',
  '/palm.png',
  'https://unpkg.com/@pluralsight/ps-design-system-normalize',

  '/api/v1/users',
  '/api/v1/user-wonders',
]

self.addEventListener('install', function install(event) {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then(function cacheUrls(cache) {
        console.log('Opened cache');
        return cache.addAll(urlsToCache);
      })
  );
});

self.addEventListener('fetch', function(event) {
  event.respondWith(
    caches.match(event.request)
      .then(function handleCacheMatch(cachedRes) {
        if (cachedRes) {
          return cachedRes;
        }

        return fetch(event.request).then(
          function(res) {
            const isInvalidRes = !res || res.status !== 200 || res.type !== 'basic' || event.request.method !== 'GET'
            if(isInvalidRes) {
              return res;
            }

            const clonedResForCacheConsumption = res.clone();

            caches.open(CACHE_NAME)
              .then(function cacheNewResponse(cache) {
                cache.put(event.request, clonedResForCacheConsumption);
              });

            return res;
          }
        );
      })
    );
});

// TODO: pickup: get the now-cached api responses to update with live server data
// https://serviceworke.rs/strategy-cache-update-and-refresh_service-worker_doc.html0
