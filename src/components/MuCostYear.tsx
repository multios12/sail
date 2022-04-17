import './MuCostYear.css';
import axios from "axios";
import { MouseEvent, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { CostModel, SumCostModel } from "../models";

/** 月ごと表示Props */
type MonthViewProps<T> = {
  onChange: Function
  Value: T
}

/** 年表示カードProps */
type YearCardProps = {
  Year: string
  SumCost: SumCostModel
}

/** 支出（年集計）コンポーネント */
export default () => {
  const { year } = useParams();
  const [sumCost, setSumCost] = useState<SumCostModel>({ Year: new Date().getFullYear().toString(), EnableYears: [], Costs: [] });

  useEffect(() => {
    const url = `./api/cost/${year ?? (new Date).getFullYear()}`
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
              </tr>
            </thead>
            <tbody>
              {sumCost.Costs.map(v => <MuMonthViewTr key={v.Month} Value={v} onChange={()=> {console.log("")}} />)}
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
const MuMonthViewTr = ({onChange, Value }: MonthViewProps<CostModel>) => {
  const [isEditable, setIsEditable] = useState<boolean>(false)
  const [Water, setWater] = useState<number>(Value.Water)
  const [Electric, setElectric] = useState<number>(Value.Electric)
  const [Gas, setGas] = useState<number>(Value.Gas)
  const [Mobile, setMobile] = useState<number>(Value.Mobile)
  const [Line, setLine] = useState<number>(Value.Line)

  const saveClick = (e: MouseEvent<HTMLButtonElement>) => {
    const url = `./api/cost/${Value.Month.toString().substring(0, 4)}/${Value.Month.toString().substring(4)}`
    axios.post(url, { Month: Value.Month, Water, Electric, Gas, Mobile, Line })
    setIsEditable(!isEditable)
  }
  return (isEditable ?
    <tr key={Value.Month}>
      <td>{`${Value.Month.toString().substring(0, 4)}年${Value.Month.toString().substring(4)}月`}</td>
      <td className="MuNumber px-0 has-text-right">{Water + Electric + Gas + Mobile + Line}</td>
      <td className="MuNumber p-0 has-text-right">{<input type="number" className="input px-0 has-text-right" value={Water} onChange={e => setWater(Number(e.target.value))} />}</td>
      <td className="MuNumber p-0 has-text-right">{<input type="number" className="input px-0 has-text-right" value={Electric} onChange={e => setElectric(Number(e.target.value))} />}</td>
      <td className="MuNumber p-0 has-text-right">{<input type="number" className="input px-0 has-text-right" value={Gas} onChange={e => setGas(Number(e.target.value))} />}</td>
      <td className="MuNumber p-0 has-text-right">{<input type="number" className="input px-0 has-text-right" value={Mobile} onChange={e => setMobile(Number(e.target.value))} />}</td>
      <td className="MuNumber p-0 has-text-right">{<input type="number" className="input px-0 has-text-right" value={Line} onChange={e => setLine(Number(e.target.value))} />}</td>
      <td>{<button className="button is-primary is-small material-icons" onClick={saveClick}>save</button>}</td>
    </tr>
    : <tr onClick={() => { setIsEditable(!isEditable) }}>
      <td>{`${Value.Month.toString().substring(0, 4)}年${Value.Month.toString().substring(4)}月`}</td>
      <td className="MuNumber px-0 has-text-right">{Water + Electric + Gas + Mobile + Line}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Water}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Electric}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Gas}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Mobile}</td>
      <td className="MuNumber px-0 has-text-right pr-4">{Line}</td>
      <td><button className="button is-info is-small is-inverted material-icons" onClick={() => {onChange(); setIsEditable(!isEditable) }}>edit</button></td>
    </tr>
  )
}
