// CryptoJS(by google) required
// crypto-core.js and crypto-sha256.js
function createHash(value) {
  return CryptoJS.SHA256(value, {outputLength: 512}).toString(CryptoJS.enc.Base64);
}

function blankCheck(value) {
  const pattern = /^\s+|\s+$/g;
  if (value.replace(pattern, "") === "") {
    return true;
  } else {
    return false;
  }
}