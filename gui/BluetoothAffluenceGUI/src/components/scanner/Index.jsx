import { useEffect, useMemo, useState } from "react";
import apiClient from "@/services/api";
import useApiResponse from "@/hooks/useApiResponse";
import { toast } from "react-toastify";

const ScannerCard = ({ scanner, onClick, key }) => {
  return (
    <div className="col-6 col-md-4 col-lg-3 col-xl-2 mb-4">
      <div className="card card-sm card-body text-center" onClick={() => onClick(scanner)}>
        <div className="text-muted mb-2">{scanner.alias}</div>
        <div className="h2 mb-0">{scanner.count}</div>
      </div>
    </div>
  );
};

const ScannerList = ({scanners}) =>{
  // Create a card foreach scanner
  return (
    <div>
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

        handleFetchResponse(lastResponse, ["scanners"], () => {
          const { scanners } = lastResponse.scanners;
          setScanners(scanners)
         
        });

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


  console.log(scanners)
  return loading ? "loading ..." : <ScannerList scanners={scanners}/>
};

export default ScannerIndex;
