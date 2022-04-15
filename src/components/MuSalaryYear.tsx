import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { DetailItem, SalaryMonthModel, SalaryModel } from "../models";

/** 汎用Props */
type Props<T> = {
  Value: T
}

/** 年表示カードProps */
type YearCardProps = {
  Year: string
  Model: SalaryModel
}

/** 給与収入（年集計）コンポーネント */
export default () => {
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
      {model.EnableYears.map(year => <MuYearCard key={year} Year={year} Model={model} />)}
    </div>
  )
}

/** 年表示カードコンポーネント */
const MuYearCard = ({ Year: year, Model: model }: YearCardProps) => {
  if (year == model.Year) {
    return (
      <div key={model.Year} className="card px-10">
        <div className="card-header">
          <div className="card-header-title">{model.Year}年</div>
        </div>
        <div className="card-content">
          <nav className="level">
            {model.Totals?.map(v => <MuTotalTile key={v.Name} Value={v} />)}
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
              {model.Details?.map(v => <MuMonthTr key={v.Month} Value={v} />)}
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

/** 合計表示タイルコンポーネント */
const MuTotalTile = ({Value}: Props<DetailItem>) => {
  return (
    <article className="tile is-child box">
      <p className="is-size-6">{Value.Name}</p>
      <p className="is-size-4">{Value.Value.toLocaleString()}</p>
    </article>
  )
}

/** 月ごと表示テーブル行コンポーネント */
const MuMonthTr = ({Value}: Props<SalaryMonthModel>) => {
  return (
    <tr key={Value.Month} className={Value.IsError ? 'has-background-danger-light' : ''}>
      <td><a href={`#/${Value.Month.substring(0, 4)}/${Value.Month.substring(4)}`}>{Value.Title}</a></td>
      <td>{Value.Totals ? Value.Totals[0].Value.toLocaleString() : 0}</td>
      <td>{Value.Totals ? Value.Totals[2].Value.toLocaleString() : 0}</td>
      <td>{Value.Month.length == 6 ? Value.Expense.toLocaleString() : ""}</td>
      <td>{Value.Month.length == 6 ? Value.Counts ? Value.Counts[0].Value : 0 : ""}</td>
    </tr>
  )
}