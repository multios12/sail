import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { DetailItem, DetailModel } from "./models"

const MuMonth = () => {
  const { year, month } = useParams();
  const [model, setModel] = useState({ Month: '202201', IsError: false, Counts: [], Times: [], Salarys: [], Costs: [], Totals: [] });

  useEffect(() => {
    const url = `./api/${year}/${month}`
    axios.get(url).then(r => {
      setModel(r.data)
    })
  }, [year, month]);

  return (
    <div className="card px-10">
      <div className="card-header">
        <div className="card-header-title">{year}年{month}月</div>
      </div>
      <div className="card-content">
        <div className="level">
          {model.Totals?.map(i => MuDetailItem(i))}
        </div>

        <div className="level">
          {model.Salarys?.map(i => MuDetailItem(i))}
        </div>

        <div className="level">
          {model.Costs?.map(i => MuDetailItem(i))}
        </div>

        <img src={`./api/${year}/${month}/detailImage`} />

      </div>
    </div>
  )
}

// 詳細項目コンポーネント
const MuDetailItem = (item: DetailItem) => {
  return (
    <div className="level-item has-text-centered">
      <div>
        <p className="heading">{item.Name}</p>
        <p className="title">{item.Value}</p>
      </div>
    </div>)
}

export default MuMonth