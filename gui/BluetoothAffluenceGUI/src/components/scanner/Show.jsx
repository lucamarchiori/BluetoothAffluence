import { useEffect, useMemo, useState, useRef } from "react";
import apiClient from "@/services/api";
import useApiResponse from "@/hooks/useApiResponse";
import { toast } from "react-toastify";
import React from "react";
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from "chart.js";
import { Line } from "react-chartjs-2";
import PlotDevices from "../charts/PlotDevices";
import 'chartjs-adapter-moment';

export const options = {
  responsive: true,
  plugins: {
    legend: {
      position: "top",
    },
    title: {
      display: true,
      text: "Chart.js Line Chart",
    },
  },
};

const Show = ({ scannerId }) => {
  const [loading, setLoading] = useState(false);
  const handleFetchResponse = useApiResponse();
  const [data, setData] = useState([]);

  useEffect(() => {
    const controller = new AbortController();

    const fetchApi = () => {
      return apiClient.get("/scanner/count-scan-devices", {
        signal: controller.signal,
      });
    };

    const fetchData = async () => {
      setLoading(true);
      try {
        const [lastResponse] = await Promise.all([fetchApi()]);
        setData(lastResponse.data.data.count);
      } catch (error) {
        if (!controller.signal.aborted) {
          console.log("Error");
          console.log(error);
          toast.error(error.config.baseURL);
        }
      } finally {
        if (!controller.signal.aborted) {
          setLoading(false);
        }
      }
    };

    fetchData();

    return () => {
      controller.abort();
    }

  }, [apiClient, handleFetchResponse]);

  const chartCanvasRef = useRef(null);
  const chartInstanceRef = useRef(null);
  console.log(data)
  return loading ? "loading ..." : <PlotDevices chartData={data} /> ;
};

export default Show;
