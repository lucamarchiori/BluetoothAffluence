import React from "react";
import { Line } from "react-chartjs-2";
import { Chart as ChartJS, LinearScale, TimeScale, PointElement, LineElement, Title, Tooltip, Legend } from "chart.js";
import ZoomPlugin from 'chartjs-plugin-zoom';

const PlotDevices = ({ chartData = [] }) => {
  const scanTimes = chartData.map((entry) => entry.scanTime);
  const counts = chartData.map((entry) => entry.count);

  ChartJS.register(LinearScale, TimeScale, PointElement, LineElement, Title, Tooltip, Legend, ZoomPlugin);

  const options = {
    bezierCurve: true,
    scales: {
      x: {
        type: "time",
        time: {
          unit: "minute",
        },
        title: {
          display: true,
          text: "Scan Time",
        },
      },
      y: {
        title: {
          display: true,
          text: "Device Count",
        },
      },
    },
    responsive: true,
    plugins: {
      legend: {
        position: "top",
      },
      title: {
        display: true,
        text: "Occupancy rate chart",
      },
      zoom: {
        pan: {
          enabled: true,
          mode: "x",
        },
        zoom: {
          enabled: true,
          drag: true,
          mode: "xy",
        },
      },
    },
  };

  const data = {
    labels: scanTimes,
    datasets: [
      {
        label: "Device count",
        data: counts,
        borderColor: "rgb(255, 99, 132)",
        backgroundColor: "rgba(255, 99, 132, 0.5)",
        tension: 0.25,
      },
    ],
  };

  return (
    <>
      <Line data={data} options={options} />
    </>
  );
};

export default PlotDevices;
