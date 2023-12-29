import { Link } from "react-router-dom";

const Navbar = () => {

  return (
    <>
      <header className="navbar navbar-expand-md navbar-light d-print-none">
        <div className="container-xl">
          <h1 className="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
            <Link to="/">
              <img
                width="110"
                height="32"
                alt="Bluetooth affluence"
                className="navbar-brand-image py-1"
              />
            </Link>
          </h1>
          <div className="navbar-nav flex-row order-md-last">
            <div className="nav-item d-none d-md-flex me-3">
            </div>
            <div className="d-none d-md-flex">
            </div>
          </div>
        </div>
      </header>
      <header className="navbar-expand-md">
      </header>
    </>
  );
};

export default Navbar;
