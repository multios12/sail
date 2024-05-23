<script lang="ts">
  import { pop } from "svelte-spa-router";
  /** メッセージ */
  let message = "";

  /** ファイル選択イベント */
  const fileChange = async () => {
    const e = document.querySelector("#fileInput") as HTMLInputElement;
    const file = e.files?.item(0);
    document.querySelector("#progress")?.classList.remove("is-hidden");
    if (!file) {
      message = "select file is null";
      return;
    }
    const data = new FormData();
    data.append("file", file);
    const r = fetch("./api/salary/files", { method: "post", body: data })
      .then((r) => {
        if (r.status === 200) {
          document.querySelector("#dialog")?.classList.remove("is-active");
          pop();
        } else {
          r.text().then((text) => {
            message = text;
          });
        }
      })
      .catch((r) => {
        console.log(r);
        message = "ファイルを保存できませんでした";
      })
      .finally(() =>
        document.querySelector("#progress")?.classList.add("is-hidden"),
      );
  };

  /** モーダルトグルイベント */
  const toggleClick = () => {
    document.querySelector("#dialog")?.classList.toggle("is-active");
  };
</script>

<div class="navbar-item field is-grouped">
  <p class="control nabbar-item">
    <button class="button is-info" on:click={toggleClick}
      ><span class="material-icons"> file_upload </span>upload</button
    >
  </p>

  <div class="modal" id="dialog">
    <div class="modal-background" />
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">明細書 アップロード</p>
        <button class="delete" aria-label="close" on:click={toggleClick} />
      </header>
      <section class="modal-card-body">
        {#if message != ""}
          <div class="notification is-danger">{message}</div>
        {/if}
        <div class="file has-name is-boxed is-fullwidth">
          <label class="file-label">
            <input
              id="fileInput"
              class="file-input"
              type="file"
              name="resume"
              on:change={fileChange}
            />
            <span class="file-cta">
              <span class="file-icon">
                <span class="material-icons"> file_upload </span>
              </span>
              <span class="file-label"> Choose a file… </span>
              <progress
                id="progress"
                class="progress is-small is-primary is-hidden"
                max="100"
              />
            </span>
          </label>
        </div>
      </section>
    </div>
  </div>
</div>
