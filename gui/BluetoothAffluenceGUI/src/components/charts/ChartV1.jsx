import Chart from "react-apexcharts";
export default function ChartV1({ height, title, categories, serie, loading, ts }) {


    var xaxis = {
        show: true,
        type: 'category',
        categories: categories ? categories : [],
        labels: {
            hideOverlappingLabels: true,
            trim: true,
            rotate: 10,
            offsetY: 10
        }
    }
    var tooltip = {}
    if (ts) {
        xaxis = {
            type: 'datetime',
            categories: categories ? categories : [],
            labels: {
                format: 'HH:mm',
            }
        }
        tooltip = {
            x: {
                format: 'HH:mm',
            }
        }
    }


    var chart = {
        options: {
            chart: {
                type: "area",
                fontFamily: 'inherit',
                toolbar: {
                    show: true,
                    tools: {
                        download: true,
                        selection: true,
                        zoom: true,
                        zoomin: true,
                        zoomout: true,
                        pan: true,
                        reset: true,
                    },
                },
            },
            fill: {
                type: 'solid',
                opacity: .25
            },
            grid: {
                padding: {
                    top: -20,
                    right: 0,
                    left: -4,
                    bottom: -4
                },
                strokeDashArray: 4,
            },
            xaxis: xaxis,
            stroke: {
                curve: "smooth",
                width: 2,
            },
            colors: ['#206bc4'],
            title: {
                text: title,
                align: 'left',
                margin: 10,
                offsetX: 0,
                offsetY: 0,
                floating: false,
                style: {
                    fontSize: '14px',
                    fontWeight: 'bold',
                    fontFamily: 'inherit',
                    color: '#263238'
                },
            },
            dataLabels: {
                enabled: false,
            },
            tooltip: tooltip,
            noData: {
                text: loading ? "Loading..." : "Dati non presenti",
                align: 'center',
                verticalAlign: 'middle',
                offsetX: 0,
                offsetY: 0,
                style: {
                    color: "#000000",
                    fontSize: '14px',
                    fontFamily: "Helvetica"
                }
            }
        },
        series: [
            {
                name: title,
                data: serie
            }
        ]
    };

    return (
        <>
            <Chart
                options={chart.options}
                series={chart.series}
                type="area"
                height={height}
            />
        </>
    )
}