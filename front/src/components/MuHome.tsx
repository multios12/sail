import "./MuHome.css";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { BalanceYear } from "../models";
import MuHomeYearCard from "./MuHomeYearCard";
/** 総合表示コンポーネント */
const MuHome = () => {
  const v: BalanceYear = { Year: new Date().getFullYear().toString(), Balances: [], EnableYears: [] }
  const { year } = useParams();
  const [model, setModel] = useState(v);

  useEffect(() => {
    const url = `./api/${year ?? (new Date()).getFullYear()}`
    fetch(url).then(r => r.json()).then(r => setModel(r))
  }, [year]);
  return (
    <div>
      {model.EnableYears.map(year => <MuHomeYearCard key={year} Year={year} Balance={model} />)}
    </div>
  )
}

export default MuHome