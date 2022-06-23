<script lang="ts">
  import type { BalanceYear } from "../models";
  import HomeMonthTr from "./HomeMonthTr.svelte";
  import { link } from "svelte-spa-router";

  export let year: string;
  export let model: BalanceYear;

  // 現在編集している行
  let editMonth = "";
</script>

{#if year === model.Year}
  <div class="card px-10">
    <div class="card-header">
      <div class="card-header-title">{model.Year}年</div>
    </div>
    <div class="card-content">
      <table class="table is-striped is-hoverable">
        <thead>
          <tr>
            <th />
            <th>給与総額</th>
            <th><a href="/salary/{year}" use:link>振込額</a></th>
            <th><a href="/cost/{year}" use:link>固定支出額</a></th>
            <th>貯蓄額</th>
            <th>メモ</th>
            <th />
          </tr>
        </thead>
        <tbody>
          {#each model.Balances as model}
            <HomeMonthTr {model} {editMonth} />
          {/each}
        </tbody>
      </table>
    </div>
  </div>
{:else}
  <div class="card">
    <div class="card-header">
      <div class="card-header-title">
        <a href="/{year}" use:link>{year}年</a>
      </div>
    </div>
  </div>
{/if}
