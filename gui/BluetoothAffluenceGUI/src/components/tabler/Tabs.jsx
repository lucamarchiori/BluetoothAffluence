import { useState } from "react";

export const TabPanel = ({ children, value, currentId, className = "" }) => {
  const classNames = className ? `tab-pane ${className}` : `tab-pane`;
  return (
    <div className={value === currentId ? `${classNames} active show` : classNames} role="tabpanel">
      {children}
    </div>
  );
};

export const TabSwitcher = ({ tabsData, className = "", overflow = false, disabled = false }) => {
  const [selectedTabId, setSelectedTabId] = useState(0);

  return (
    <div className={className ? "card " + className : "card"}>
      <div className="card-header">
        <ul className="nav nav-tabs card-header-tabs nav-fill" role="tablist">
          {tabsData.map((tab) => (
            <li key={tab.id} className="nav-item" role="presentation">
              <button
                className={disabled ? "nav-link disabled" : tab.id === selectedTabId ? "nav-link active" : "nav-link"}
                aria-selected={tab.id === selectedTabId}
                role="tab"
                onClick={() => setSelectedTabId(tab.id)}
              >
                {tab.header}
              </button>
            </li>
          ))}
        </ul>
      </div>
      <div className={overflow ? "card-body p-0 overflow-auto" : "card-body p-0"}>
        <div className={overflow ? "tab-content h-100 overflow-auto" : "tab-content"}>
          {tabsData.map((tab) => tab.renderPanel(tab.id, selectedTabId))}
        </div>
      </div>
    </div>
  );
};