const apiBaseUrl = "http://localhost:3001/api";

async function uniFetch(url, options) {
  options = options || {};
  const method = options.method || 'GET';
  const body = JSON.stringify(options.body) || undefined;
  // let it throw network (and CORS etc.) errors
  const response = await fetch(apiBaseUrl + url, {
    method,
    body,
    headers: {
      'Content-Type': 'application/json'
    }
  });
  // request completed but response code not 200
  if (!response.ok) {
    let errMsg = null;
    try {
      // try to get error data
      const { data } = await response.json();
      errMsg = data;
    } catch (e) {
      // on error, do thing and fallback to normal route
    }
    if (errMsg) throw errMsg;
  }
  const json = await response.json();
  return json["data"];
}

export { apiBaseUrl, uniFetch };
