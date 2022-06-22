import './MuCostYear.css';
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { BalanceYear } from "../models";
import MuCostYearMonthTr from "./MuCostYearMonthTr"
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
    fetch(url).then(r => r.json()).then(r => setSumCost(r))
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
              {sumCost.Balances.map(v => <MuCostYearMonthTr key={v.Month} Value={v} EditMonth={editMonth} SetEditMonth={setEditMonth} />)}
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
export default MuSalaryYear
