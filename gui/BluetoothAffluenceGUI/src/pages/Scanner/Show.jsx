import Page from "@/components/tabler/Page";
import PageHeader from "@/components/tabler/PageHeader";
import PageBody from "@/components/tabler/PageBody";
import Show from "@/components/scanner/Show";
import { useParams } from "react-router-dom";

const ShowPage = () => {
    const { scannerId } = useParams();
    return (
        <Page>
          <PageHeader>
            <h1 className="page-title">Show occupancy rate of device {scannerId}</h1>
          </PageHeader>
          <PageBody>
            <Show scannerId={scannerId}/>
          </PageBody>
        </Page>
      );
};

export default ShowPage;
