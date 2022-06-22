import "./MuHome.css";
import { MouseEvent, useEffect, useState } from "react";
import { BalanceItem } from "../models";

/** 月ごと表示Props */
type MonthViewProps<T> = {
  EditMonth: string
  SetEditMonth: React.Dispatch<React.SetStateAction<string>>
  Value: T
}

/** 月ごと表示テーブル行コンポーネント */
const MuHomeMonthTr = (props: MonthViewProps<BalanceItem>) => {
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
    fetch(url, { method: "post", body: JSON.stringify(props.Value) })
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
      <td className="MuNumber has-text-right">{props.Value.Salary.toLocaleString()}</td>
      <td className="MuNumber has-text-right">{props.Value.Paid.toLocaleString()}</td>
      <td className={(props.Value.IsNotCost ? 'has-background-danger-light' : '') + ' has-text-right'}>{props.Value.Cost.toLocaleString()}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{saving.toLocaleString()}</td>
      <td>{memo}</td>
      <td><button id={`edit${props.Value.Month}`} className="button is-info is-small is-inverted material-icons" onClick={() => { props.SetEditMonth(props.Value.Month) }}>edit</button></td>
    </tr>
  )
}

export default MuHomeMonthTr