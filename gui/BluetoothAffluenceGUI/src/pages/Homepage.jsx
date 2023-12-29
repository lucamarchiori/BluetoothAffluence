import React from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Line } from "react-chartjs-2";

ChartJS.register(ArcElement, Tooltip, Legend);

import Page from "@/components/tabler/Page";

const Homepage = () => {
  const blur = "blur(20px)";
  return (
    <Page title="Rust2dakar">
      <main className="overflow-x-hidden">

        <div className="container-xl text-center pt-5">
          <p className="fs-1 fw-light lh-lg pt-5 pb-5">
          Lorem ipsum
          </p>
      <Line />

        </div>
         
      </main>
    </Page>
  );
};

export default Homepage;
