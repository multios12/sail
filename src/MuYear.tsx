import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { DetailItem, DetailModel, YearModel } from "./models";

// リストコンポーネント
const MuYear = () => {
  const { year } = useParams();
  const [model, setModel] = useState({ Year: '2022', Details: [], Totals: [], EnableYears: [] });

  useEffect(() => {
    const url = `./api/${year ? year : (new Date).getFullYear()}`
    axios.get(url).then(r => {
      setModel(r.data)
    })
  }, [year]);
  return (
    <div>
      {model.EnableYears.map(year => MuListMonth(year, model))}
    </div>
  )
}

const MuListMonth = (year: string, model: YearModel) => {
  if (year == model.Year) {
    return (
      <div className="card px-10">
        <div className="card-header">
          <div className="card-header-title">{model.Year}年</div>
        </div>
        <div className="card-content">
          <nav className="level">
            {model.Totals?.map(i => MuTotalItem(i))}
          </nav>
          <table className="table is-striped is-hoverable">
            <thead>
              <th></th>
              <th>支給額</th>
              <th>差引支給額</th>
              <th>出勤日数</th>
            </thead>
            <tbody>
              {model.Details?.map(v => MuListItem(v))}
            </tbody>
          </table>
        </div>
      </div>
    )
  } else {
    return (
      <div className="card">
        <div className="card-header">
          <div className="card-header-title"><a href={"#/" + year}>{year}年</a></div>
        </div>
      </div>
    )
  }
}

// 合計表示コンポーネント
const MuTotalItem = (item: DetailItem) => {
  return (
    <div className="level-item has-text-centered">
      <div>
        <p className="heading">{item.Name}</p>
        <p className="title">{item.Value.toLocaleString()}</p>
      </div>
    </div>)
}

// リストアイテムコンポーネント
const MuListItem = (model: DetailModel) => {
  return (
    <tr className={model.IsError ? 'has-background-danger-light' : ''}>
      <td><a href={`#/${model.Month.substring(0, 4)}/${model.Month.substring(4)}`}>{model.Title}</a></td>
      <td>{model.Totals ? model.Totals[0].Value.toLocaleString(): 0}</td>
      <td>{model.Totals ? model.Totals[2].Value.toLocaleString() : 0}</td>
      <td>{model.Counts ? model.Counts[0].Value : 0}</td>
    </tr>
  )
}

export default MuYear