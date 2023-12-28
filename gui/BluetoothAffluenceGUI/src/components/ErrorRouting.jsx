import { isRouteErrorResponse, useRouteError, Link } from "react-router-dom";

import NotFound from "../pages/errors/NotFound";
import Unauthorized from "../pages/errors/Unauthorized";
import Page from "./tabler/Page";

const ErrorRouting = () => {
  const error = useRouteError();

  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      return (
        <Page title="Errore 404 | Securways">
          <NotFound />
        </Page>
      );
    }

    if (error.status === 401) {
      return (
        <Page title="Errore 401 | Securways">
          <Unauthorized />
        </Page>
      );
    }
  }

  return (
    <Page title="Pagina di errore | Securways">
      <div className="row text-center">
        <div className="col-12 col-md-6 col-lg-4 mx-auto">
          <div className="alert alert-danger mt-6" role="alert">
            <h4 className="alert-heading" lang="en">
              Something went wrong!
            </h4>
            <hr />
            <p className="mb-0" lang="en">
              Please contact the website administrators
            </p>
            <Link to="/" lang="en">Return to the homepage</Link>
          </div>
        </div>
      </div>
    </Page>
  );
};

export default ErrorRouting;