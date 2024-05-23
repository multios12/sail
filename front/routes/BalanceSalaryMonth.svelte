<script lang="ts">
  import type { SalaryMonthModel } from "../models/balanceModels.js";
  import { onMount } from "svelte";

  export let params: { year: string | undefined; month: string | undefined } = {
    year: undefined,
    month: undefined,
  };

  let model: SalaryMonthModel = {
    Month: "202010",
    Counts: [],
    Times: [],
    Salarys: [],
    Costs: [],
    Totals: [],
    Expense: 0,
    Expenses: [],
    Images: [],
  };

  const refreshClick = (e: any) => {
    const t = e.target as HTMLButtonElement;

    t.classList.add("is-loading");
    document.querySelector(".card-content")?.classList.add("is-hidden");
    const url = `./api/salary/${params.year}/${params.month}`;
    fetch(url, { method: "post" })
      .then((r) => r.json())
      .then((r) => (model = r))
      .finally(() => {
        t.classList.remove("is-loading");
        document.querySelector(".card-content")?.classList.remove("is-hidden");
      });
  };

  onMount(async () => {
    const url = `./api/salary/${params.year}/${params.month}`;
    console.log(url);
    const r = await fetch(url);
    model = await r.json();
  });
</script>

<div class="card px-10">
  <div class="card-header">
    <div class="card-header-title">
      {model.Title}
      <div class="control px-5">
        <button class="button is-info" on:click={refreshClick}
          ><span class="material-icons"> refresh </span></button
        >
      </div>
    </div>
  </div>
  <div class="card-content">
    <table class="table is-fullwidth">
      <tbody>
        <tr class="has-background-success-light">
          <td>
            <p>支給</p>
            <p>
              {model.Totals && model.Totals.length > 0
                ? model.Totals[0].Value.toLocaleString()
                : 0}
            </p>
          </td>
          <td>
            <div class="columns">
              {#each model.Salarys as item}
                <div class="column">
                  <article class="tile is-child box">
                    <p>{item.Name}</p>
                    <p>{item.Value.toLocaleString()}</p>
                  </article>
                </div>
              {/each}
            </div>
          </td>
        </tr>
        <tr class="has-background-danger-light">
          <td class="is-one-fifth">
            <p>控除</p>
            <p>
              {model.Totals && model.Totals.length > 0
                ? model.Totals[1].Value.toLocaleString()
                : 0}
            </p>
          </td>
          <td>
            <div class="columns">
              {#each model.Costs as item}
                <div class="column">
                  <article class="tile is-child box">
                    <p>{item.Name}</p>
                    <p>{item.Value.toLocaleString()}</p>
                  </article>
                </div>
              {/each}
            </div>
          </td>
        </tr>
        <tr class="has-background-info-light">
          <td class="is-one-fifth">
            <p>差引</p>
            <p>
              {model.Totals && model.Totals.length > 0
                ? model.Totals[2].Value.toLocaleString()
                : 0}
            </p>
          </td>
          <td />
        </tr>
      </tbody>
    </table>
    {#each model.Images as i}
      <img
        src={`./api/salary/${params.year}/${params.month}/images/${i}`}
        alt="salary"
      />
    {/each}
  </div>
</div>
