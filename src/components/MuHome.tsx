import "./MuHome.css";
import axios from "axios";
import { MouseEvent, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { BalanceItem, BalanceYear } from "../models";

/** 汎用Props */
type Props<T> = {
  Value: T
}

/** 月ごと表示Props */
type MonthViewProps<T> = {
  EditMonth: string
  SetEditMonth: React.Dispatch<React.SetStateAction<string>>
  Value: T
}

/** 年表示カードProps */
type YearCardProps = {
  Year: string
  Balance: BalanceYear
}

/** 総合表示コンポーネント */
const MuHome= () => {
  const { year } = useParams();
  const [model, setModel] = useState({ Year: new Date().getFullYear().toString(), Balances: [], EnableYears: [] });

  useEffect(() => {
    const url = `./api/${year ?? (new Date()).getFullYear()}`
    axios.get(url).then(r => {
      setModel(r.data)
    })
  }, [year]);
  return (
    <div>
      {model.EnableYears.map(year => <MuYearCard key={year} Year={year} Balance={model} />)}
    </div>
  )
}

/** 年表示カードコンポーネント */
const MuYearCard = ({ Year: year, Balance: model }: YearCardProps) => {
  // 現在編集している行
  const [editMonth, setEditMonth] = useState<string>("")

  if (year === model.Year) {
    return (
      <div key={model.Year} className="card px-10">
        <div className="card-header">
          <div className="card-header-title">{model.Year}年</div>
        </div>
        <div className="card-content">
          <table className="table is-striped is-hoverable">
            <thead>
              <tr>
                <th></th>
                <th>給与総額</th>
                <th><a href={`#/salary/${year}`}>振込額</a></th>
                <th><a href={`#/cost/${year}`}>固定支出額</a></th>
                <th>貯蓄額</th>
                <th>メモ</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {model.Balances.map(v => <MuMonthTr key={v.Month} Value={v} EditMonth={editMonth} SetEditMonth={setEditMonth} />)}
            </tbody>
          </table>
        </div>
      </div>
    )
  } else {
    return (
      <div className="card">
        <div className="card-header">
          <div className="card-header-title"><a href={`#/${year}`}>{year}年</a></div>
        </div>
      </div>
    )
  }
}

/** 月ごと表示テーブル行コンポーネント */
const MuMonthTr = (props: MonthViewProps<BalanceItem>) => {
  const [saving, SetSaving] = useState(props.Value.Saving)
  const [memo, SetMemo] = useState(props.Value.Memo)

  useEffect(() => {
    if (props.EditMonth === "" || props.EditMonth === props.Value.Month) {
      document.querySelector(`#edit${props.Value.Month}`)?.classList.remove("is-hidden")
    } else {
      document.querySelector(`#edit${props.Value.Month}`)?.classList.add("is-hidden")
    }

  }, [props.Value.Month, props.EditMonth]);

  const saveClick = (e: MouseEvent<HTMLButtonElement>) => {
    props.Value.Saving = saving
    props.Value.Memo = memo
    const url = `./api/${props.Value.Month.toString().substring(0, 4)}/${props.Value.Month.toString().substring(4)}`
    axios.post(url, props.Value)
    props.SetEditMonth("")
  }

  return (props.EditMonth === props.Value.Month ?
    <tr key={props.Value.Month}>
      <td className="MuNumber p-0 has-text-right">{props.Value.Month.substring(0, 4)}年{props.Value.Month.substring(4)}月</td>
      <td className="MuNumber p-0 has-text-right">{props.Value.Salary.toLocaleString()}</td>
      <td className="MuNumber p-0 has-text-right">{props.Value.Paid.toLocaleString()}</td>
      <td className={props.Value.IsNotCost ? 'has-background-danger-light' : ''}>{props.Value.Cost.toLocaleString()}</td>
      <td className="MuNumber p-0 has-text-right">
        <input type="number" className="input px-0 has-text-right" value={saving} onChange={e => SetSaving(Number(e.target.value))} />
      </td>
      <td className="MuNumber p-0 has-text-right">
        <input type="text" className="input px-0 has-text-right" value={memo} onChange={e => SetMemo(e.target.value)} />
      </td>
      <td>{<button className="button is-primary is-small material-icons" onClick={saveClick}>save</button>}</td>
    </tr>
    :
    <tr key={props.Value.Month}>
      <td>{props.Value.Month.substring(0, 4)}年{props.Value.Month.substring(4)}月</td>
      <td>{props.Value.Salary.toLocaleString()}</td>
      <td>{props.Value.Paid.toLocaleString()}</td>
      <td className={props.Value.IsNotCost ? 'has-background-danger-light' : ''}>{props.Value.Cost.toLocaleString()}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{saving.toLocaleString()}</td>
      <td>{memo}</td>
      <td><button id={"edit" + props.Value.Month} className="button is-info is-small is-inverted material-icons" onClick={() => { props.SetEditMonth(props.Value.Month) }}>edit</button></td>
    </tr>
  )
}

export default MuHome