async function applyConfiguration() {
    const url = window.location.pathname;

    const specificConfigRes = await fetch(`http://127.0.0.1:8080/api/specific/123e4567-e89b-12d3-a456-526614174000?url=${url}`);
    const specificConfig = await specificConfigRes.json();
    
    const matchedConfigs = specificConfig.matchedConfigs || [];
    
    let allActions = [];

    for (const configId of matchedConfigs) {
        const configRes = await fetch(`http://127.0.0.1:8080/api/configuration/${configId}`);
        const config = await configRes.json();
        allActions = allActions.concat(config.actions);
    }

    allActions = prioritizeActions(allActions);

    allActions.forEach(applyAction);
}

function prioritizeActions(actions) {
    const priorityMap = {
        "remove": 3,  // Highest priority: Remove elements first
        "replace": 2, // Replace elements second
        "alter": 1,   // Alter text third
        "insert": 0   // Insert elements last
    };

    return actions.sort((a, b) => priorityMap[a.type] - priorityMap[b.type]);
}

function applyAction(action) {
    switch(action.type) {
        case "insert":
            insertElement(action);
            break;
        case "alter":
            alterText(action);
            break;
        case "remove":
            removeElement(action);
            break;
        case "replace":
            replaceElement(action);
            break;
    }
}

function insertElement(action) {
    const target = document.querySelector(action.target);
    if (!target) return;

    const newElement = document.createElement('div');
    newElement.innerHTML = action.element;

    if (action.position === "before") {
        target.before(newElement);
    } else {
        target.after(newElement);
    }
}

function alterText(action) {
    document.body.innerHTML = document.body.innerHTML.replace(new RegExp(action.oldValue, "g"), action.newValue);
}

function removeElement(action) {
    document.querySelectorAll(action.selector).forEach(el => el.remove());
}

function replaceElement(action) {
    document.querySelectorAll(action.selector).forEach(el => {
        const newEl = document.createElement('div');
        newEl.innerHTML = action.newElement;
        el.replaceWith(newEl);
    });
}

// Run when the page loads
window.onload = applyConfiguration;
