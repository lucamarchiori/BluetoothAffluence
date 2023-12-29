const InfoList = ({ listItemsInfo = [], loading = false, className = "", itemClassName = "mb-3" }) => {
  const classNames = [
    "list-unstyled",
    className,
    loading ? "placeholder-glow" : "",
  ];



  return (
    <ul className={classNames.filter((cn) => cn != "").join(" ")}>
      {loading ? (
        listItemsInfo.map((_, index) => (
          <li key={index} className={itemClassName}>
            <div className={`placeholder placeholder-xs col-${(index + 8) - (Math.floor((index + 1) / 3) * 3)}`}></div>
          </li>
        ))
      ) : (
        listItemsInfo.map((item, index) => (
          <li key={index} className={itemClassName}>
            <span>{item.label}:</span>
            <span className="fw-bold ms-1">{item.content}</span>
          </li>
        ))
      )}
    </ul>
  )
};

export default InfoList;