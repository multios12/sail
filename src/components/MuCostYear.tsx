import './MuCostYear.css';
import axios from "axios";
import React, { MouseEvent, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { BalanceItem, BalanceYear } from "../models";

/** 月ごと表示Props */
type MonthViewProps<T> = {
  EditMonth: string
  SetEditMonth: React.Dispatch<React.SetStateAction<string>>
  Value: T
}

/** 年表示カードProps */
type YearCardProps = {
  Year: string
  SumCost: BalanceYear
}

/** 支出（年集計）コンポーネント */
const MuSalaryYear = () => {
  const { year } = useParams();
  const [sumCost, setSumCost] = useState<BalanceYear>({ Year: new Date().getFullYear().toString(), EnableYears: [], Balances: [] });

  useEffect(() => {
    const url = `./api/${year ?? (new Date()).getFullYear()}`
    axios.get(url).then(r => {
      setSumCost(r.data)
    })
  }, [year]);
  return (
    <div>
      {sumCost.EnableYears.map(y => <MuYearCard key={y} Year={y} SumCost={sumCost} />)}
    </div>
  )
}

/** 年表示カードコンポーネント */
const MuYearCard = ({ Year: year, SumCost: sumCost }: YearCardProps) => {
  // 現在編集している行
  const [editMonth, setEditMonth] = useState<string>("")

  if (year === sumCost.Year) {
    return (
      <div key={sumCost.Year} className="card px-10">
        <div className="card-header">
          <div className="card-header-title">{sumCost.Year}年支出</div>
        </div>
        <div className="card-content">
          <table className="table is-striped is-hoverable">
            <thead>
              <tr>
                <th></th>
                <th>合計</th>
                <th>水道</th>
                <th>電気</th>
                <th>ガス</th>
                <th>携帯</th>
                <th>通信</th>
                <th>納税</th>
              </tr>
            </thead>
            <tbody>
              {sumCost.Balances.map(v => <MuMonthViewTr key={v.Month} Value={v} EditMonth={editMonth} SetEditMonth={setEditMonth} />)}
            </tbody>
          </table>
        </div>
      </div>
    )
  } else {
    return (
      <div className="card">
        <div className="card-header">
          <div className="card-header-title"><a href={"#/cost/" + year}>{year}年</a></div>
        </div>
      </div>
    )
  }
}

/** 月ごと表示テーブル行コンポーネント */
const MuMonthViewTr = (props: MonthViewProps<BalanceItem>) => {
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
    axios.post(url, { Month: props.Value.Month, CostWater: Water, CostElectric: Electric, CostGas: Gas, CostMobile: Mobile, CostLine: Line, CostTax: Tax })
    props.SetEditMonth("")
  }
  return (props.EditMonth === props.Value.Month ?
    <tr key={props.Value.Month}>
      <td>{`${props.Value.Month.substring(0, 4)}年${props.Value.Month.substring(4)}月`}</td>
      <td className="MuNumber px-0 has-text-right">{Water + Electric + Gas + Mobile + Line}</td>
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
export default MuSalaryYear
