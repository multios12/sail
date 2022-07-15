<script lang="ts">
  import { onMount } from "svelte";
  type lineType = { Day: string; Memo: string };
  type listType = { WritedMonths: string[]; Lines: lineType[] };
  let date: Date | null = new Date();
  let memo = "";
  let items: listType = { WritedMonths: [], Lines: [] };
  let selectMonth = "2022-02";

  /** 送信 クリックイベント */
  const sendButtonClick = async () => {
    if (date != null) {
      let url = formatDate(date);
      url = `./api/${url}`;

      await fetch(url, {
        method: "post",
        body: JSON.stringify({ Memo: memo }),
      });
      showLines();
    }
  };

  /** リストクリックイベント */
  const listClick = (e: any, l: lineType) => {
    // const y = Number(l.Day.substring(0, 4))
    // const m = Number(l.Day.substring(5, 7)) - 1
    // const d = Number(l.Day.substring(8, 10))

    // const day = new Date(y, m, d)
    setDate2(l.Day);
    memo = l.Memo;
  };

  const showLines = async () => {
    const url = `./api/${formatMonth(new Date())}`;
    const r = await fetch(url).then((r) => r.json());
    items = r as listType;
  };

  const setDate2 = (value: string) => {
    let e = value === "" ? null : new Date(value);

    date = e;
    if (e == null) {
      e = new Date();
    }

    var d = e.getFullYear() + "-" + (e.getMonth() + 1);
    var a = items?.WritedMonths?.filter((v, i) => v === d);

    if (a.length === 0 && items.WritedMonths.length > 0) {
      d = items.WritedMonths[0];
      selectMonth = d;
    } else {
      selectMonth = d;
    }
  };

  const formatMonth = (dt: Date) => {
    return `${dt.getFullYear()}/${("00" + (dt.getMonth() + 1)).slice(-2)}`;
  };

  const formatDate = (dt: Date) => {
    return `${dt.getFullYear()}/${("00" + (dt.getMonth() + 1)).slice(-2)}/${(
      "00" + dt.getDate()
    ).slice(-2)}`;
  };

  onMount(showLines);
</script>

<nav class="navbar is-dark" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="./">la-diary</a>

    <button class="navbar-burger" aria-label="menu" aria-expanded="false">
      <span aria-hidden="true" />
      <span aria-hidden="true" />
      <span aria-hidden="true" />
    </button>
  </div>

  <div id="menu" class="navbar-menu">
    <div class="navbar-start" />

    <div class="navbar-end" />
  </div>
</nav>
<main>
  <div class="card px-10">
    <div class="card-header">
      <input type="date" />
      <input type="text" bind:value={memo} />
      <button class="button" onClick={sendButtonClick}>send</button>
    </div>
    <div class="card-content">
      <div class="select">
        <select value={selectMonth}>
          {#each items.WritedMonths as v}
            <option value={v}>{v}</option>
          {/each}
        </select>
      </div>
      <table>
        <tbody>
          {#each items.Lines as v}
            <tr on:click={(e) => listClick(e, v)}><td>{v.Day} {v.Memo}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</main>

<style>
  @import url("bulma/css/bulma.css");
</style>
