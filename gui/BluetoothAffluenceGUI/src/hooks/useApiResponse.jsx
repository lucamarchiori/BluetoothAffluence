import { useCallback } from "react";
import { toast } from "react-toastify";

const useApiResponse = () => {
  const handleFetchResponse = useCallback((response, properties, onSuccess) => {
    if (response.status === 200 && response.data) {
      const { data } = response;
      const hasProperties = properties.every((property) => data.hasOwnProperty(property));
      if (data.status === 200 && hasProperties) {
        onSuccess();
      }
      if (data.message && data.message.error) {
        toast.dismiss();
        toast.error(data.message.error);
      }
    }
  }, []);

  return handleFetchResponse;
};

export default useApiResponse;