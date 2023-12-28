import { Route, createBrowserRouter, createRoutesFromElements } from "react-router-dom";

import App from "@/App";
import Homepage from "@/pages/Homepage";
import ErrorRouting from "@/components/ErrorRouting";


export const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<App />} errorElement={<ErrorRouting />}>
      <Route index element={<Homepage />} />
    </Route>
  )
);
