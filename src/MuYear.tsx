import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { DetailItem, DetailModel, YearModel } from "./models";

type Props<T> = {
  Value: T
}

type MonthProps = {
  year: string
  model: YearModel
}

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
      {model.EnableYears.map(year => <MuListMonth key={year} year={year} model={model} />)}
    </div>
  )
}

const MuListMonth = ({ year, model }: MonthProps) => {
  if (year == model.Year) {
    return (
      <div key={model.Year} className="card px-10">
        <div className="card-header">
          <div className="card-header-title">{model.Year}年</div>
        </div>
        <div className="card-content">
          <nav className="level">
            {model.Totals?.map(v => <MuTotalItem key={v.Name} Value={v} />)}
          </nav>
          <table className="table is-striped is-hoverable">
            <thead>
              <tr>
                <th></th>
                <th>支給額</th>
                <th>差引支給額</th>
                <th>経費支給額</th>
                <th>出勤日数</th>
              </tr>
            </thead>
            <tbody>
              {model.Details?.map(v => <MuListItem key={v.Month} Value={v} />)}
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
const MuTotalItem = ({Value}: Props<DetailItem>) => {
  return (
    <article className="tile is-child box">
      <p className="is-size-6">{Value.Name}</p>
      <p className="is-size-4">{Value.Value.toLocaleString()}</p>
    </article>
  )
}

// リストアイテムコンポーネント
const MuListItem = (props: Props<DetailModel>) => {
  const model = props.Value
  return (
    <tr key={model.Month} className={model.IsError ? 'has-background-danger-light' : ''}>
      <td><a href={`#/${model.Month.substring(0, 4)}/${model.Month.substring(4)}`}>{model.Title}</a></td>
      <td>{model.Totals ? model.Totals[0].Value.toLocaleString() : 0}</td>
      <td>{model.Totals ? model.Totals[2].Value.toLocaleString() : 0}</td>
      <td>{model.Month.length == 6 ? model.Expense.toLocaleString() : ""}</td>
      <td>{model.Month.length == 6 ? model.Counts ? model.Counts[0].Value : 0 : ""}</td>
    </tr>
  )
}

export default MuYear