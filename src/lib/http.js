'use strict';

function doGet(url) {
  return new Promise(function (resolve, reject) {
    const r = new XMLHttpRequest();
    r.onreadystatechange = function () {
      if (r.readyState === 4) {
        if (r.status > 204) {
          return reject(new Error(`wrong status code ${r.status}`));
        }
        let body;
        try {
          body = JSON.parse(r.responseText);
        } catch (err) {
          body = r.responseText;
        }
        resolve({
          statusCode: r.status,
          body
        });
      }
    };
    r.open('GET', url, true);
    r.send(null);
  });
}

export default doGet;
