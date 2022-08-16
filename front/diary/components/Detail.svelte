<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import TagsInput from "./TagsInput.svelte";
  import type { detailType, lineType, listType } from "../components/models";

  const dispatch = createEventDispatcher();
  export let Day: string | null;
  let Outline = "";
  let Detail = "";
  let Tags = [];
  let isDayEdit: boolean;

  $: {
    if (Day === undefined || Day === null) {
    } else if (Day === "") {
      isDayEdit = true;
      const dt = new Date();
      Day = `${dt.getFullYear()}-`;
      Day += ("00" + (dt.getMonth() + 1)).slice(-2);
      Day += `-${("00" + dt.getDate()).slice(-2)}`;
      Outline = "";
      Tags = [];
      Detail = "";
    } else {
      isDayEdit = false;
      let url = Day.replaceAll("-", "/");
      url = `./api/diary/${url}`;
      fetch(url, { method: "get" })
        .then((r) => r.json())
        .then((r) => {
          const s = r as detailType;
          Outline = s.Outline;
          Detail = s.Detail;
          Tags = s.Tags;
        });
    }
  }

  /** 送信 クリックイベント */
  const sendClick = async () => {
    if (Day == null) {
      return;
    }

    let url = Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    const init = {
      method: "post",
      body: JSON.stringify({ Tags, Outline, Detail }),
    };
    await fetch(url, init);
    Day = null;
    dispatch("update");
  };

  /** キャンセル クリックイベント */
  const cancelClick = () => (Day = null);

  /** 削除 クリックイベント */
  const deleteClick = async () => {
    let url = Day.replaceAll("-", "/");
    url = `./api/diary/${url}`;
    await fetch(url, { method: "delete" });
    Day = null;
    dispatch("update");
  };
</script>

<div class="modal" class:is-active={Day != null}>
  <div class="modal-background" />
  <div class="modal-card">
    <header class="modal-card-head sp-panel-heading">
      <input type="date" class="input" bind:value={Day} disabled={!isDayEdit} />
      {#if !isDayEdit}
        <button
          class="button is-inverted is-small has-text-danger"
          on:click={deleteClick}
        >
          <i class="material-icons">delete</i>
        </button>
      {/if}
    </header>
    <section class="modal-card-body">
      <!-- Card Content -->
      <input
        type="text"
        placeholder="outline"
        class="input"
        bind:value={Outline}
      />
      <div class="field">
        <div class="control">
          <TagsInput bind:items={Tags} />
        </div>
        <textarea class="textarea" placeholder="detail" bind:value={Detail} />
      </div>
    </section>
    <footer class="modal-card-foot">
      <button class="button is-primary" on:click={sendClick}>send</button>
      <button class="button" on:click={cancelClick}>cancel</button>
    </footer>
  </div>
</div>
