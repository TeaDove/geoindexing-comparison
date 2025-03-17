function initFieldset(){
    fetch("/tasks", {method: "GET"})
        .then((response) => response.json())
        .then((json) => console.log(`runs-resume ${json}`));
}

function resumeClick(el) {
    fetch("/runs/resume", {
        method: "POST",
        body: JSON.stringify({}),
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        }
    })
        .then((response) => response.json())
        .then((json) => console.log(`runs-resume ${json}`));
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
        .then((json) => console.log(`runs-reset ${json}`));
}

function setupPlotWS(chart){
    const wsSocket = new WebSocket("/plots/ws");

    wsSocket.onopen = (_) => {
        console.info("plot.ws.socket.open")
    };
    wsSocket.onclose = (_) => {
        console.info("plot.ws.socket.closed")
    };

    wsSocket.onmessage = (event) => {
        console.debug(`plot.ws.event.received ${event.data}`)
        const drawTO = JSON.parse(event.data);
        if (drawTO.legendToPoints !== undefined){
            for (const label in drawTO.legendToPoints) {
                chart.data.datasets.forEach((dataset, idx) => {
                    if (dataset.label === label){
                        chart.data.datasets[idx].data.push(...drawTO.legendToPoints[label]);
                        chart.update();
                    }
                })
            }
        }
    };
}

function handleHover(evt, item, legend) {
    legend.chart.data.datasets[0].backgroundColor.forEach((color, index, colors) => {
        colors[index] = index === item.index || color.length === 9 ? color : color + '4D';
    });
    legend.chart.update();
}

function handleLeave(evt, item, legend) {
    legend.chart.data.datasets[0].backgroundColor.forEach((color, index, colors) => {
        colors[index] = color.length === 9 ? color.slice(0, -2) : color;
    });
    legend.chart.update();
}

function setupChart() {
    const ctx = document.getElementById('speedChart');

    const data = {
        datasets: [{
            label: 'First Dataset',
            data: [{
                x: 0,
                y: 0,
            }, {
                x: 1,
                y: 1,
            }, {
                x: 2,
                y: 2,
            }],
            backgroundColor: 'rgb(255, 99, 132)'
        },
            {
                label: 'Second Dataset',
                data: [{
                    x: 1,
                    y: 0,
                }, {
                    x: 2,
                    y: 1,
                }, {
                    x: 3,
                    y: 1,
                }],
                backgroundColor: 'rgb(84,157,81)'
            }],
    };
    return new Chart(ctx, {
        type: 'scatter',
        data: data,
        options: {
            showLine: true,
            responsive: true,
            plugins: {
                legend: {
                    position: 'top',
                    onHover: handleHover,
                    onLeave: handleLeave
                },
                title: {
                    display: true,
                    text: 'Время решения задачи'
                }
            },
        }
    })
}

function main(){
    initFieldset();
    const chart = setupChart();
    setupPlotWS(chart);
}

main()