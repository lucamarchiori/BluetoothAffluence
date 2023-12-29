import { Link } from "react-router-dom";
import EmptyElement from "../../components/tabler/EmptyElement";

const Unauthorized = () => {
  return (
    <div className="container-thight py-4">
      <EmptyElement
        header="401"
        title="Non sei autorizzato ad accedere a questo contenuto"
        subtitle="Se vedi questa pagina vuol dire che non hai l'autorizzazione al contenuto che cerchi."
        action={
          <Link to="/" className="btn btn-primary">Riportami alla <span lang="en" className="ms-1">homepage</span></Link>
        }
      />
    </div>
  );
};

export default Unauthorized;