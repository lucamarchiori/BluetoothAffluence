import { useMatch, useResolvedPath, Link } from "react-router-dom";

const NavbarLink = ({ children, to, end = true, ...props }) => {
  const resolved = useResolvedPath(to);
  const match = useMatch({ path: resolved.pathname, end: end });

  return (
    <li className={match ? "nav-item active" : "nav-item"}>
      <Link to={to} className="nav-link" {...props}>
        {children}
      </Link>
    </li>
  );
};

export default NavbarLink;