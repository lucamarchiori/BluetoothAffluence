import React from "react";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  TimeScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';

const PlotDevices = (props) => {
  var dataExample = [
    { scan_time: "2021-08-01 00:00:00", numDevices: 10 },
    { scan_time: "2021-08-01 00:01:00", numDevices: 11 },
    { scan_time: "2021-08-01 00:02:00", numDevices: 12 },
    { scan_time: "2021-08-01 00:03:00", numDevices: 13 },
    { scan_time: "2021-08-01 00:04:00", numDevices: 14 },
    { scan_time: "2021-08-01 00:05:00", numDevices: 15 },
    { scan_time: "2021-08-01 00:06:00", numDevices: 16 },
    { scan_time: "2021-08-01 00:07:00", numDevices: 17 },
    { scan_time: "2021-08-01 00:08:00", numDevices: 18 },
    { scan_time: "2021-08-01 00:09:00", numDevices: 19 },
    { scan_time: "2021-08-01 00:10:00", numDevices: 16 },
    { scan_time: "2021-08-01 00:11:00", numDevices: 17 },
    { scan_time: "2021-08-01 00:12:00", numDevices: 13 },
    { scan_time: "2021-08-01 00:13:00", numDevices: 14 },
    { scan_time: "2021-08-01 00:14:00", numDevices: 8 },
    { scan_time: "2021-08-01 00:15:00", numDevices: 9 },
  ];

  const scanTimes = dataExample.map((entry) => entry.scan_time);
  const counts = dataExample.map((entry) => entry.numDevices);

  console.log(scanTimes, counts)

  ChartJS.register(
    LinearScale,
    TimeScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
  );

  const options = {
    scales: {
      x: {
        type: 'time',
        time: {
          unit: 'minute',
        },
        title: {
          display: true,
          text: 'Scan Time',
        },
      },
      y: {
        title: {
          display: true,
          text: 'Device Count',
        },
      },
    },
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'Occupancy rate chart'
      },
    },
  };

  const data = {
    labels: scanTimes,
    datasets: [
      {
        label: 'Device count',
        data: counts,
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
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
