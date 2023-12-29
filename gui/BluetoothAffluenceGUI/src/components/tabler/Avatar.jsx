const avatarVariants = new Map([
  ["default", ""],
  ["xl", "avatar-xl"],
  ["lg", "avatar-lg"],
  ["md", "avatar-md"],
  ["sm", "avatar-sm"],
  ["xs", "avatar-xs"],
])

const Avatar = ({ children, url = null, variant = "default", className = "" }) => {
  const style = url ? {
    backgroundImage: `url(${url})`
  } : null;

  const classNames = ["avatar", avatarVariants.get(variant), className].filter((cn) => cn != "").join(" ");
  return ( 
    <span className={classNames} style={style} aria-hidden>
      {!url && children}
    </span>
  );
}

export default Avatar;