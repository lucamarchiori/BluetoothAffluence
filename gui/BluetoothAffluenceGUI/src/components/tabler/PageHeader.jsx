const PageHeader = ({ children }) => {
  return (
    <div className="page-header">
      <div className="container-xl">
        {children}
      </div>
    </div>
  );
};

export default PageHeader;