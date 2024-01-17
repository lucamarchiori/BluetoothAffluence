import React from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Line } from "react-chartjs-2";
import ScannerIndex from "@/components/scanner/Index"
ChartJS.register(ArcElement, Tooltip, Legend);

import Page from "@/components/tabler/Page";

const Homepage = () => {
  return (
    <Page title="">
      <main className="overflow-x-hidden">

        <div className="container-xl text-center pt-5">
          <p className="fs-1 fw-light lh-lg pt-5 pb-5">
          Occupancy rate bluethooth scanner
          </p>
          <ScannerIndex></ScannerIndex>

        </div>
      </main>
    </Page>
  );
};

export default Homepage;
