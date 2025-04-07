function initFieldset() {
    fetch("/tasks", { method: "GET" })
        .then((response) => response.json())
        .then(tasks => {
            const legend = document.querySelector("fieldset#tasks");
            tasks.forEach(task => {
                const el = document.createElement("div");
                el.innerHTML = `
                    <input type="checkbox" id="index-${task.info.shortName}" name="${task.info.shortName}" />
                    <label for="index-${task.info.shortName}">${task.info.longName}</label>
            `;
                legend.appendChild(el);
            })
        });

    fetch("/indexes", { method: "GET" })
        .then((response) => response.json())
        .then(indexes => {
            const legend = document.querySelector("fieldset#indexes");
            indexes.forEach(index => {
                const el = document.createElement("div");
                el.innerHTML = `
                    <input type="checkbox" id="index-${index.info.shortName}" name="${index.info.shortName}" />
                    <label for="index-${index.info.shortName}">${index.info.longName}</label>
            `;
                legend.appendChild(el);
            })
        });
}

function resumeClick(el) {
    const selectedTasks = Array.from(document.querySelectorAll('fieldset#tasks input:checked'))
        .map(checkbox => checkbox.name);

    const selectedIndexes = Array.from(document.querySelectorAll('fieldset#indexes input:checked'))
        .map(checkbox => checkbox.name);

    const pointsStart = parseInt(document.getElementById('pointsStart').value);
    const pointsEnd = parseInt(document.getElementById('pointsEnd').value);
    const pointsStep = parseInt(document.getElementById('pointsStep').value);

    fetch("/runs/resume", {
        method: "POST",
        body: JSON.stringify({
            "tasks": selectedTasks,
            "indexes": selectedIndexes,
            "start": pointsStart,
            "end": pointsEnd,
            "step": pointsStep
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((response) => response.json())
        .then((json) => console.log(`runs-resume ${JSON.stringify(json)}`));
}

function resetClick(el) {
    fetch("/runs/reset", {
        method: "POST",
        body: JSON.stringify({}),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((response) => response.json())
        .then((json) => console.log(`runs-reset ${JSON.stringify(json)}`));

    window.location.reload();
}


function main() {
    initFieldset();
}

main()