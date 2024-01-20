import { Link } from "react-router-dom";

const Navbar = () => {

  return (
    <>
      <footer className="footer footer-transparent d-print-none border-top pb-1">
          <div className="container-xl" >
            <div className="row text-center align-items-center">
              <div className="col-12 col-lg-auto mt-3 mt-lg-0">
                <ul className="list-inline list-inline-dots mb-0">
                  <li className="list-inline-item">
                    Luca Marchiori - Università degli Studi di Padova - Wireless Networks for Mobile Applications A.A. 2023-2024
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </footer>
    </>
  );
};

export default Navbar;
