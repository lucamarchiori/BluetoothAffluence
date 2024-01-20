import { Outlet } from "react-router-dom";
import Navbar from "@/components/tabler/Navbar";
import Footer from "@/components/tabler/Footer";

const App = () => {
  return (
    <>
      <Navbar />
      <div className="page-wrapper">
        <Outlet />
      </div>
      <Footer />
    </>
  );
};

export default App;
