import EmptyResultImage from "@/assets/tabler/img/undraw_quitting_time_dm8t.svg?react";

const EmptyElement = ({ header, title, subtitle, action, hasImage = false }) => {
  return (
    <div className="empty">
      {hasImage && <div className="empty-image">
        <EmptyResultImage width={256} height={128} />
      </div>}
      {header && <div className="empty-header">{header}</div>}
      {title && <p className="empty-title">{title}</p>}
      {subtitle && <p className="empty-subtitle text-muted">{subtitle}</p>}
      {action && <div className="empty-action">{action}</div>}
    </div>
  );
};

export default EmptyElement;