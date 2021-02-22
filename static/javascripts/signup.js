const tooltipList = {};

function createTooltip(elem, message) {
  const type = elem.getAttribute("name");
  if (tooltipList[type] === undefined) {
    // Create tooltip element (popper.js required)
    const $tooltip = document.createElement("div", {"role": "tooltip", "data-popper-placement": "bottom"});
    const $tooltipContent = document.createElement("span");
    const $tooltipArrow = document.createElement("div", {"data-popper-arrow": ""});
    // Set class name
    $tooltip.className = "tooltip-err";
    $tooltipContent.className = "tooltip-content";
    $tooltipArrow.className = "arrow";
    // Set content and combine component
    $tooltipContent.appendChild(document.createTextNode(message));
    $tooltip.appendChild($tooltipContent);
    $tooltip.appendChild($tooltipArrow);
    $tooltip.setAttribute("data-show", "");
    // Add created tooltip element
    document.getElementById("form-signup").appendChild($tooltip);

    // Add tooltip element and object in tooltipList
    tooltipList[type] = {
      element: $tooltip,
      object: Popper.createPopper(elem, $tooltip, {
        modifiers: [{
          name: "offset",
          options: {
            offset: [-100, 8]
          }
        }]
      })
    };
  }
}

function destroyTooltip(type) {
  if (tooltipList[type] !== undefined) {
    tooltipList[type].element.remove();
    tooltipList[type].object.destroy();
    tooltipList[type] = undefined;
  }
}

function verifyInputValue(elem) {
  let unsatisfactory = false;
  const blankPattern = /[\s]/g;
  if (blankPattern.test(elem.value)) {
    unsatisfactory = true;
    message = "Can't contain spaces.";
  } else {
    const type = elem.getAttribute("name");
    if (type === "name") {
      unsatisfactory = elem.value.length > 15;
      message = "Username is not available.";
    } else if (type === "email") {
      const emailRegex = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/i;
      unsatisfactory = (elem.value.length > 0 && !emailRegex.test(elem.value));
      message = "This email is not valid.";
    } else if (type === "password") {
      unsatisfactory = (elem.value.length !== 0 && elem.value.length < 8);
      message = "Password must be at least 8 digits long.";
    }
  }

  if (unsatisfactory) {
    createTooltip(elem, message);
    elem.classList.add("invalid");
    elem.setAttribute("data-valid", false);
  } else {
    destroyTooltip(elem.getAttribute("name"));
    elem.classList.remove("invalid");
    if (elem.value.length === 0) {
      elem.setAttribute("data-valid", false);
    } else {
      elem.setAttribute("data-valid", true);
    }
  }
}

// CryptoJS(by google) required
// crypto-core.js and crypto-sha256.js
function createHash(value) {
  return CryptoJS.SHA256(value, {outputLength: 40}).toString(CryptoJS.enc.Base64);
}