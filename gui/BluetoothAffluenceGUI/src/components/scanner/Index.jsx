import { useEffect, useMemo, useState } from "react";
import apiClient from "@/services/api";
import useApiResponse from "@/hooks/useApiResponse";
import { toast } from "react-toastify";

const ScannerCard = ({ scanner, onClick, key }) => {
  return (
    <div className="col-6 col-md-4 col-lg-4 col-xl-4 mb-4">
      <div className="card card-sm card-body text-center d-flex justify-content-center align-items-center"  onClick={() => onClick(scanner)}>
        <div className="text-muted mb-2"><b>Alias/Name: </b>{scanner.alias ? scanner.alias : (scanner.name ? scanner.name : "-")}</div>
        <div className="text-muted mb-2"><b>Address: </b>{scanner.address}</div>
        <hr className="w-100 mb-3 mt-1" />
        <a className="btn btn-primary btn-sm px-4" style={{width: 120}} href="#">View</a>
      </div>
    </div>
  );
};

const ScannerList = ({scanners}) =>{
  // Create a card foreach scanner
  return (
    <div className="row">
      {scanners.map((scanner, index) => (
        <ScannerCard key={index} scanner={scanner} />
      ))}
    </div>
  );
  
};

const ScannerIndex = () => {
  const [loading, setLoading] = useState(false);
  const handleFetchResponse = useApiResponse();
  const [scanners, setScanners] = useState([]);

  useEffect(() => {
    const controller = new AbortController();

    const fetchApi = () => {
      return apiClient.get("/scanner/index", {
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


  return loading ? "loading ..." : <ScannerList scanners={scanners}/>
};

export default ScannerIndex;
