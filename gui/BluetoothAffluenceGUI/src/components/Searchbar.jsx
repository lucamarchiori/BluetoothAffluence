const Searchbar = ({ labelText, placeholder, searchText, handleChange, handleSubmit, disabledSearchButton = false }) => {
  return (
    <form className="row" onSubmit={handleSubmit}>
      <div className="col-auto pb-3">
        <label htmlFor="search" className="visually-hidden">{labelText}</label>
        <input
          id="search"
          type="text"
          className="form-control"
          style={{ width: "250px" }}
          placeholder={placeholder}
          value={searchText}
          onChange={handleChange}
        />
      </div>
      <div className="col-auto pb-3">
        <button className="btn btn-primary" disabled={disabledSearchButton}>Cerca</button>
      </div>
    </form>
  );
};

export default Searchbar;