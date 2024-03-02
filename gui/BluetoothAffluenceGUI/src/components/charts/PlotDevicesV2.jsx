import ChartV1 from "./ChartV1";
export default function PlotDevicesV2({ chartData = [], loading, title, ts }) {
    const scanTimes = chartData.map((entry) => entry.scanTime);
    const counts = chartData.map((entry) => entry.count);

    const loadingChart = (
        <>
            <div style={{ height: 200 }}>
                <center>
                    <div style={{ height: "50px", width: "50px" }} className="spinner-border mt-5"></div>
                </center>
            </div>
        </>
    )

    const emptyChart = (
        <>
            <div>
                <center>
                    <div className="mt-5 mb-5">Nessun dato da visualizzare</div>
                </center>
            </div>
        </>
    )

    return (
        <>
            <div className="row row-cards mb-3">
                <div className="col-12 col-md-12 col-xl-12">
                    <div className="card">
                        <div className="card-body px-1 pt-1 pb-0">
                            {loading ? (loadingChart) : (scanTimes && counts ?
                                <ChartV1
                                    height={200}
                                    title={title}
                                    serie={counts}
                                    loading={loading}
                                    categories={scanTimes}
                                    ts = {ts}
                                />
                                : emptyChart)}
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}