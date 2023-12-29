import { Outlet } from "react-router-dom";

import Navbar from "@/components/tabler/Navbar";

const App = () => {
  return (
    <>
      <Navbar />
      <div className="page-wrapper">
        <Outlet />
      </div>
    </>
  );
};

export default App;
