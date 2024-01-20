

import { useEffect, useMemo, useState } from "react";
import apiClient from "@/services/api";
import useApiResponse from "@/hooks/useApiResponse";
import { toast } from "react-toastify";

const Show = ({scannerId}) => {
  const [loading, setLoading] = useState(false);
  const handleFetchResponse = useApiResponse();
  const [data, setData] = useState([]);

  useEffect(() => {
    const controller = new AbortController();

    const fetchApi = () => {
      return apiClient.get("/scanner/show", {
        signal: controller.signal,
      });
    };

    const fetchData = async () => {
      setLoading(true);
      try {
        const [lastResponse] = await Promise.all([fetchApi()]);
        setScanners(lastResponse.data.data.scanners);
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

    return () => controller.abort();
  }, [apiClient, handleFetchResponse]);


  return loading ? "loading ..." : "ciao"
};

export default Show;
