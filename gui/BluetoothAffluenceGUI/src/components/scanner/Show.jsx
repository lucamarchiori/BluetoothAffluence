import { useEffect, useMemo, useState, useRef } from "react";
import apiClient from "@/services/api";
import useApiResponse from "@/hooks/useApiResponse";
import { toast } from "react-toastify";
import React from "react";
import PlotDevicesV2 from "../charts/PlotDevicesV2";

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
        setData(lastResponse.data);
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

  console.log(data)
  return loading ? "loading ..." : (
    <>
      {(data && data.data && data.data.count.count) ? <PlotDevicesV2 chartData={data.data.count.count} ts={true} title={"Device count"} /> : ""}
      {(data && data.data && data.data.countMovingAvg.count) ? <PlotDevicesV2 chartData={data.data.countMovingAvg.count} ts={true} title={"Device count (moving AVG)"} /> : ""}
      {(data && data.data && data.data.timeAvg.count) ? <PlotDevicesV2 chartData={data.data.timeAvg.count} title={"Device count (time AVG)"} /> : ""}
    
    </>
  );
};

export default Show;
