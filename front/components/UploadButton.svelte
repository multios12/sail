<script lang="ts">
  import { pop } from "svelte-spa-router";
  /** メッセージ */
  let message = "";
  let isDialogOpen = false;
  let isUploading = false;

  /** ファイル選択イベント */
  const fileChange = async () => {
    const e = document.querySelector("#fileInput") as HTMLInputElement;
    const file = e.files?.item(0);
    if (!file) {
      message = "select file is null";
      return;
    }
    isUploading = true;
    const data = new FormData();
    data.append("file", file);
    fetch("/api/salary/files", { method: "post", body: data })
      .then((r) => {
        if (r.status === 200) {
          isDialogOpen = false;
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
      .finally(() => {
        isUploading = false;
      });
  };

  /** モーダルトグルイベント */
  const toggleClick = () => {
    isDialogOpen = !isDialogOpen;
  };
</script>

<div class="upload-button">
  <button class="upload-trigger" onclick={toggleClick}>
    <span class="material-icons" aria-hidden="true"> file_upload </span>
    upload
  </button>

  {#if isDialogOpen}
    <div class="upload-modal" id="dialog" role="dialog" aria-modal="true">
      <button
        class="dialog-backdrop"
        aria-label="close"
        onclick={toggleClick}
      ></button>
      <div class="dialog-panel">
        <header class="dialog-head">
          <p class="dialog-title">明細書 アップロード</p>
          <button class="close-button" aria-label="close" onclick={toggleClick}></button>
        </header>
        <section class="dialog-body">
          {#if message != ""}
            <div class="error-message">{message}</div>
          {/if}
          <label class="file-picker">
            <input
              id="fileInput"
              class="picker-input"
              type="file"
              name="resume"
              onchange={fileChange}
            />
            <span class="picker-target">
              <span class="material-icons" aria-hidden="true"> file_upload </span>
              <span>Choose a file...</span>
            </span>
            {#if isUploading}
              <progress class="upload-progress" max="100"></progress>
            {/if}
          </label>
        </section>
      </div>
    </div>
  {/if}
</div>

<style>
  .upload-button {
    display: inline-flex;
    align-items: center;
  }

  .upload-trigger {
    display: inline-flex;
    align-items: center;
    gap: 0.35rem;
    min-height: 2.5rem;
    border: 1px solid #0ea5e9;
    border-radius: 4px;
    padding: 0.5rem 0.8rem;
    background: #0ea5e9;
    color: #fff;
    cursor: pointer;
    font: inherit;
    line-height: 1;
  }

  .upload-trigger:hover {
    background: #0284c7;
    border-color: #0284c7;
  }

  .upload-modal {
    position: fixed;
    inset: 0;
    z-index: 40;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 1rem;
  }

  .dialog-backdrop {
    position: absolute;
    inset: 0;
    border: 0;
    background: rgb(17 24 39 / 70%);
    cursor: pointer;
  }

  .dialog-panel {
    position: relative;
    z-index: 1;
    width: min(100%, 640px);
    overflow: hidden;
    border-radius: 6px;
    background: #fff;
    color: #111827;
    box-shadow: 0 24px 60px rgb(15 23 42 / 35%);
  }

  .dialog-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 1rem;
    border-bottom: 1px solid #e5e7eb;
    padding: 1rem 1.25rem;
    background: #f9fafb;
  }

  .dialog-title {
    margin: 0;
    font-size: 1.15rem;
    font-weight: 600;
  }

  .close-button {
    position: relative;
    width: 2rem;
    height: 2rem;
    border: 0;
    border-radius: 999px;
    background: transparent;
    cursor: pointer;
  }

  .close-button::before,
  .close-button::after {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 1rem;
    height: 2px;
    border-radius: 999px;
    background: #4b5563;
    content: "";
  }

  .close-button::before {
    transform: translate(-50%, -50%) rotate(45deg);
  }

  .close-button::after {
    transform: translate(-50%, -50%) rotate(-45deg);
  }

  .close-button:hover {
    background: #e5e7eb;
  }

  .dialog-body {
    padding: 1.25rem;
  }

  .error-message {
    margin-bottom: 1rem;
    border-radius: 4px;
    padding: 0.75rem 1rem;
    background: #fee2e2;
    color: #991b1b;
  }

  .file-picker {
    display: block;
    width: 100%;
  }

  .picker-input {
    position: absolute;
    width: 1px;
    height: 1px;
    overflow: hidden;
    clip: rect(0 0 0 0);
    white-space: nowrap;
  }

  .picker-target {
    display: flex;
    min-height: 8rem;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    border: 2px dashed #94a3b8;
    border-radius: 6px;
    padding: 1.5rem;
    background: #f8fafc;
    color: #334155;
    cursor: pointer;
    text-align: center;
  }

  .picker-target:hover {
    border-color: #0ea5e9;
    background: #f0f9ff;
  }

  .upload-progress {
    width: 100%;
    height: 0.5rem;
    margin-top: 1rem;
    accent-color: #0ea5e9;
  }
</style>
