import "./MuHome.css";
import { useState } from "react";
import { BalanceYear } from "../models";
import MuHomeMonthTr from "./MuHomeMonthTr";

/** 年表示カードProps */
type YearCardProps = {
  Year: string
  Balance: BalanceYear
}

/** 年表示カードコンポーネント */
const MuHomeYearCard = ({ Year: year, Balance: model }: YearCardProps) => {
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
              {model.Balances.map(v => <MuHomeMonthTr key={v.Month} Value={v} EditMonth={editMonth} SetEditMonth={setEditMonth} />)}
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

export default MuHomeYearCard