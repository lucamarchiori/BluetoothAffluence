import { Link } from "react-router-dom";
import EmptyElement from "../../components/tabler/EmptyElement";

const NotFound = () => {
  return (
    <div className="container-thight py-4">
      <EmptyElement
        header="404"
        title="Oops... Hai trovato un pagina di errore"
        subtitle="La pagina che cercavi non l'abbiamo trovata, ci dispiace"
        action={
          <Link to="/" className="btn btn-primary">
            Riportami alla <span className="ms-1" lang="en">homepage</span>
          </Link>
        }
      />
    </div>
  );
};

export default NotFound;