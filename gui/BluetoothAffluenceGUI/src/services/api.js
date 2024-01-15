import axios from 'axios';

const BASE_URL = "http://127.0.0.1:4000";

const apiClient = axios.create({
  baseURL: BASE_URL,
  withCredentials: false,
  headers: { 'Accept': 'application/json' }
});
export default apiClient;