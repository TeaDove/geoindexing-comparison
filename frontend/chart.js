// function setupPlotWS(chart) {
//     const wsSocket = new WebSocket("/plots/ws");

//     wsSocket.onopen = (_) => {
//         console.info("plot.ws.socket.open")
//     };
//     wsSocket.onclose = (_) => {
//         console.info("plot.ws.socket.closed")
//     };

//     wsSocket.onmessage = (event) => {
//         console.debug(`plot.ws.event.received len=${event.data.length}`)
//         const points = JSON.parse(event.data);
//         points.forEach(point => {
//             const chart = getOrCreateChart(point.chart);
//             for (const dataset of chart.data.datasets) {
//                 if (dataset.label == point.dataset) {
//                     dataset.data.push({ x: point.x, y: point.y });
//                     return
//                 }
//             }

//             chart.data.datasets.push({
//                 label: point.dataset,
//                 data: [{ x: point.x, y: point.y }],
//             });
//         });
//     };
// }
//     const chart = getOrCreateChart(drawTO.chartName);
//     if (drawTO.legendToPoints !== undefined) {
//         for (const label in drawTO.legendToPoints) {
//             chart.data.datasets.forEach((dataset, idx) => {
//                 if (dataset.label === label) {
//                     chart.data.datasets[idx].data.push(...drawTO.legendToPoints[label]);
//                     chart.update();
//                 }
//             })
//         }
//     }
// };


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

const charts = {};

function getOrCreateChart(chartName) {
    if (charts[chartName] !== undefined) {
        return charts[chartName];
    }

    var chartElement = document.getElementById(chartName);
    if (chartElement === null) {
        chartElement = document.createElement("canvas");
        chartElement.id = chartName;
        document.getElementById("chartsDiv").appendChild(chartElement);
    }

    const chart = new Chart(chartElement, {
        type: 'scatter',
        data: { datasets: [] },
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
                    text: chartName,
                }
            },
        }
    })

    charts[chartName] = chart;
    return chart;
}

function main() {
    setupPlotWS();
}

main()