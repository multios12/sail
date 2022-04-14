import axios from "axios";
import { MouseEventHandler, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { DetailItem, DetailModel } from "./models"
type Props<T> = {
  item: T
}
const MuMonth = () => {
  const item: DetailModel = { Month: '202201', Counts: [], Times: [], Salarys: [], Costs: [], Totals: [], Expense: 0, Expenses: [], Images: [] }
  const { year, month } = useParams();
  const [model, setModel] = useState(item);

  const refreshClick = (e: any) => {
    const t = e.target as HTMLButtonElement

    t.classList.add("is-loading")
    document.querySelector(".card-content")?.classList.add("is-hidden")
    const url = `./api/${year}/${month}`
    axios.put(url).then(r => {
      const url = `./api/${year}/${month}`
      axios.get(url).then(r => {
        setModel(r.data)
      }).finally(() => {
        t.classList.remove("is-loading")
        document.querySelector(".card-content")?.classList.remove("is-hidden")
      })
    }).catch(() => {
      t.classList.remove("is-loading")
      document.querySelector(".card-content")?.classList.remove("is-hidden")
    })
  }

  useEffect(() => {
    const url = `./api/${year}/${month}`
    axios.get(url).then(r => {
      setModel(r.data)
    })
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
                  {model.Salarys?.map(i => <MuDetailItem key={i.Name} item={i} />)}
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
                  {model.Costs?.map(i => <MuDetailItem key={i.Name} item={i} />)}
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
        {model.Images?.map(i => <img key={i} src={`./api/${year}/${month}/images/${i}`} />)}
      </div>
    </div>
  )
}

// 詳細項目コンポーネント
const MuDetailItem = ({ item }: Props<DetailItem>) => {
  return (
    <div className="column">
      <article className="tile is-child box">
        <p >{item.Name}</p>
        <p >{item.Value.toLocaleString()}</p>
      </article>
    </div>)
}

export default MuMonth