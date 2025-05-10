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