import './MuCostYear.css';
import React, { MouseEvent, useEffect, useState } from "react";
import { BalanceItem } from "../models";
/** 月ごと表示Props */
type MonthViewProps<T> = {
  EditMonth: string
  SetEditMonth: React.Dispatch<React.SetStateAction<string>>
  Value: T
}

/** 月ごと表示テーブル行コンポーネント */
const MuCostYearMonthTr = (props: MonthViewProps<BalanceItem>) => {
  const [Water, setWater] = useState<number>(props.Value.CostWater)
  const [Electric, setElectric] = useState<number>(props.Value.CostElectric)
  const [Gas, setGas] = useState<number>(props.Value.CostGas)
  const [Mobile, setMobile] = useState<number>(props.Value.CostMobile)
  const [Line, setLine] = useState<number>(props.Value.CostLine)
  const [Tax, setTax] = useState<number>(props.Value.CostTax)

  useEffect(() => {
    if (props.EditMonth === "" || props.EditMonth === props.Value.Month) {
      document.querySelector(`#edit${props.Value.Month}`)?.classList.remove("is-hidden")
    } else {
      document.querySelector(`#edit${props.Value.Month}`)?.classList.add("is-hidden")
    }

  }, [props.Value.Month, props.EditMonth]);

  const saveClick = (e: MouseEvent<HTMLButtonElement>) => {
    const url = `./api/${props.Value.Month.toString().substring(0, 4)}/${props.Value.Month.toString().substring(4)}`
    const d = { Month: props.Value.Month, CostWater: Water, CostElectric: Electric, CostGas: Gas, CostMobile: Mobile, CostLine: Line, CostTax: Tax }
    fetch(url, { method: "post", body: JSON.stringify(d) })
    props.SetEditMonth("")
  }
  return (props.EditMonth === props.Value.Month ?
    <tr key={props.Value.Month}>
      <td>{`${props.Value.Month.substring(0, 4)}年${props.Value.Month.substring(4)}月`}</td>
      <td className="MuNumber px-0 has-text-right">{Water + Electric + Gas + Mobile + Line + Tax}</td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Water} onChange={e => setWater(Number(e.target.value))} /></td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Electric} onChange={e => setElectric(Number(e.target.value))} /></td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Gas} onChange={e => setGas(Number(e.target.value))} /></td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Mobile} onChange={e => setMobile(Number(e.target.value))} /></td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Line} onChange={e => setLine(Number(e.target.value))} /></td>
      <td className="MuNumber p-0 has-text-right"><input type="number" className="input px-0 has-text-right" value={Tax} onChange={e => setTax(Number(e.target.value))} /></td>
      <td>{<button className="button is-primary is-small material-icons" onClick={saveClick}>save</button>}</td>
    </tr>
    : <tr>
      <td>{`${props.Value.Month.substring(0, 4)}年${props.Value.Month.substring(4)}月`}</td>
      <td className="MuNumber px-0 has-text-right">{Water + Electric + Gas + Mobile + Line + Tax}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Water}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Electric}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Gas}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Mobile}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Line}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Tax}</td>
      <td><button id={"edit" + props.Value.Month} className="button is-info is-small is-inverted material-icons" onClick={() => { props.SetEditMonth(props.Value.Month) }}>edit</button></td>
    </tr>
  )
}
export default MuCostYearMonthTr
