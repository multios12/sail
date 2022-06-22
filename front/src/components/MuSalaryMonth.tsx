import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { SalaryMonthModel } from "../models";

/** 汎用Props */
type Props<T> = {
  item: T
}

/** 給与収入（月ごと）コンポーネント */
const MuSalaryMonth = () => {
  const item: SalaryMonthModel = { Month: "202010", Counts: [], Times: [], Salarys: [], Costs: [], Totals: [], Expense: 0, Expenses: [], Images: [] }
  const { year, month } = useParams();
  const [model, setModel] = useState(item);

  const refreshClick = (e: any) => {
    const t = e.target as HTMLButtonElement

    t.classList.add("is-loading")
    document.querySelector(".card-content")?.classList.add("is-hidden")
    const url = `./api/salary/${year}/${month}`
    fetch(url, { method: "post" }).then(r => r.json()).then(r => setModel(r))
      .finally(() => {
        t.classList.remove("is-loading")
        document.querySelector(".card-content")?.classList.remove("is-hidden")
      })
  }

  useEffect(() => {
    const url = `./api/salary/${year}/${month}`
    fetch(url).then(r => r.json()).then(r => setModel(r))
  }, [year, month]);

  return (
    <div className="card px-10">
      <div className="card-header">
        <div className="card-header-title">{model.Title}
          <div className="control px-5"><button className="button is-info" onClick={refreshClick}><span className="material-icons"> refresh </span></button></div>
        </div>
      </div>
      <div className="card-content">
        <table className="table is-fullwidth">
          <tbody>
            <tr className="has-background-success-light">
              <td>
                <p>支給</p>
                <p>{model.Totals && model.Totals.length > 0 ? model.Totals[0].Value.toLocaleString() : 0}</p>
              </td>
              <td>
                <div className="columns">
                  {model.Salarys?.map(item => {
                    return <div className="column">
                      <article className="tile is-child box">
                        <p >{item.Name}</p>
                        <p >{item.Value.toLocaleString()}</p>
                      </article>
                    </div>
                  })}
                </div>
              </td>
            </tr>
            <tr className="has-background-danger-light">
              <td className="is-one-fifth">
                <p>控除</p>
                <p>{model.Totals && model.Totals.length > 0 ? model.Totals[1].Value.toLocaleString() : 0}</p>
              </td>
              <td>
                <div className="columns">
                  {model.Costs?.map(item => {
                    return <div className="column">
                      <article className="tile is-child box">
                        <p >{item.Name}</p>
                        <p >{item.Value.toLocaleString()}</p>
                      </article>
                    </div>
                  })}
                </div>
              </td>
            </tr>
            <tr className="has-background-info-light">
              <td className="is-one-fifth">
                <p>差引</p>
                <p>{model.Totals && model.Totals.length > 0 ? model.Totals[2].Value.toLocaleString() : 0}</p>
              </td>
              <td>
              </td>
            </tr>

          </tbody>
        </table>
        {model.Images?.map(i => <img key={i} src={`./api/salary/${year}/${month}/images/${i}`} alt="salary" />)}
      </div>
    </div>
  )
}

export default MuSalaryMonth