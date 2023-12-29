import logo from "../assets/images/FloidLogo.png";

const PageLoading = ({ message, className }) => {
  return (
    <div className="page page-center">
      <div className="container container-slim py-4">
        <div className="text-center">
          <div className="mb-3">
            <img
              src={logo}
              width="110"
              height="32"
              alt="Securways"
              className="navbar-brand-image"
            />
          </div>
          <div className="text-muted mb-3">{message}</div>
          <div className="progress progress-sm">
            <div className="progress-bar progress-bar-indeterminate"></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PageLoading;