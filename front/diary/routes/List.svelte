<script lang="ts">
  import { onMount } from "svelte";
  type lineType = { Day: string; Memo: string };
  type listType = { WritedMonths: string[]; Lines: lineType[] };
  let edit: lineType = { Day: "", Memo: "" };
  let model: listType = { WritedMonths: [], Lines: [] };
  let isDayEdit: boolean;
  let selectMonth = "2022-02";

  /** 送信 クリックイベント */
  const sendClick = async () => {
    if (edit.Day == null) {
      return;
    }

    let url = edit.Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    const init = {
      method: "post",
      body: JSON.stringify({ Memo: edit.Memo }),
    };
    await fetch(url, init);
    showLines();
  };

  /** リストクリックイベント */
  const listClick = (e: any, l: lineType) => {
    isDayEdit = false;
    edit.Day = l.Day;
    edit.Memo = l.Memo;
  };

  const clearClick = () => {
    edit.Day = "";
    edit.Memo = "";
  };

  const addClick = () => {
    isDayEdit = true;
    const dt = new Date();
    edit.Day = `${dt.getFullYear()}-`;
    edit.Day += ("00" + (dt.getMonth() + 1)).slice(-2);
    edit.Day += `-${("00" + dt.getDate()).slice(-2)}`;
    edit.Memo = "";
  };

  const showLines = async () => {
    const dt = new Date();
    let url = `${dt.getFullYear()}/${("00" + (dt.getMonth() + 1)).slice(-2)}`;
    url = `./api/diary/${url}`;
    const r = await fetch(url).then((r) => r.json());
    model = r as listType;
  };

  onMount(() => {
    const dt = new Date();
    selectMonth = `${dt.getFullYear()}-${("00" + (dt.getMonth() + 1)).slice(
      -2
    )}`;
    showLines();
  });
</script>

<div class="box">
  {#if edit.Day != ""}
    <div class="card px-10">
      <div class="card-content">
        <div class="columns">
          <div class="column">
            <input
              type="date"
              class="input"
              bind:value={edit.Day}
              disabled={!isDayEdit}
            />
          </div>
          <div class="column">
            <input type="text" class="input" bind:value={edit.Memo} />
          </div>
          <div class="column">
            <button class="button is-primary" on:click={sendClick}>send</button>
            <button class="button" on:click={clearClick}>clear</button>
          </div>
        </div>
      </div>
    </div>
  {/if}
  <div class="card px-10">
    <div class="card-content">
      <div class="columns">
        <div class="column">
          <div class="select">
            <select class="select" value={selectMonth} on:change={showLines}>
              {#each model.WritedMonths as v}
                <option value={v}>{v}</option>
              {/each}
            </select>
          </div>
        </div>
        <div class="column">
          {#if edit.Day == ""}
            <button class="button is-primary" on:click={addClick}>add</button>
          {/if}
        </div>
      </div>
      <table class="table is-hoverable is-fullwidth">
        <tbody>
          {#each model.Lines as v}
            <tr on:click={(e) => listClick(e, v)}>
              <td>{v.Day} {v.Memo}</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</div>
