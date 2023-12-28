import React from "react";

export default class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    // Update state so the next render will show the fallback UI.
    return { hasError: true };
  }

  componentDidCatch(error, info) {
    console.error(error, info);
  }

  render() {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      return (
        <div className="row text-center">
          <div className="col-12 col-md-6 col-lg-4 mx-auto">
            <div className="alert alert-danger mt-6" role="alert">
              <h4 className="alert-heading">
                Something went wrong!
              </h4>
              <hr />
              <p className="mb-0">
                Please contact the website administrators
              </p>
            </div>
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}