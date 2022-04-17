import axios from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { BalanceItem, BalanceYear } from "../models";

/** 汎用Props */
type Props<T> = {
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
              </tr>
            </thead>
            <tbody>
              {model.Balances.map(v => <MuMonthTr key={v.Month} Value={v} />)}
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
const MuMonthTr = ({ Value }: Props<BalanceItem>) => {
  const year = Value.Month.toString().substring(0, 4)
  const month = Value.Month.toString().substring(4)
  return (
    <tr key={Value.Month}>
      <td>{year}年{month}月</td>
      <td>{Value.Salary.toLocaleString()}</td>
      <td>{Value.Paid.toLocaleString()}</td>
      <td className={Value.IsNotCost ? 'has-background-danger-light' : ''}>{Value.Cost.toLocaleString()}</td>
      <td>{Value.Saving.toLocaleString()}</td>
      <td>{Value.Memo}</td>
    </tr>
  )
}

export default MuHome