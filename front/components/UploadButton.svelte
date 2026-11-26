<script lang="ts">
  export let onImported: () => void = () => {};

  const labels = {
    imported: "取り込み完了",
    replaced: "差し替え完了",
    already_registered: "同じ年月・種別の明細が登録済みです",
    invalid_filename: "ファイル名が規則に合っていません",
    conversion_failed: "変換に失敗しました。PDFの内容とパスワードを確認してください",
    save_failed: "保存に失敗しました",
  };
  type Result = { filename: string; status: keyof typeof labels };

  let message = "";
  let files: File[] = [];
  let results: Result[] = [];
  let replacement: number | null = null;
  let isDialogOpen = false;
  let isUploading = false;

  async function upload(selected: File[], replaceIndex?: number) {
    isUploading = true;
    message = "";
    const data = new FormData();
    selected.forEach((file) => data.append("file", file));
    if (replaceIndex !== undefined) data.append("replace", "true");

    try {
      const response = await fetch("./api/salary/files", { method: "POST", body: data });
      if (!response.ok) {
        const detail = (await response.text()).trim();
        message = detail || ([502, 503, 504].includes(response.status)
          ? "sailのAPIに接続できません。開発環境では Go API（3001番）を起動してください。"
          : `アップロードに失敗しました（HTTP ${response.status}）`);
        return;
      }

      const received: Result[] = await response.json();
      if (replaceIndex === undefined) {
        results = received;
      } else {
        results = results.map((result, i) => i === replaceIndex ? received[0] : result);
        if (received[0].status === "replaced" || received[0].status === "imported") replacement = null;
      }
      if (received.some(({ status }) => status === "imported" || status === "replaced")) {
        onImported();
      }
    } catch {
      message = "アップロード結果を取得できませんでした";
    } finally {
      isUploading = false;
    }
  }

  function fileChange(event: Event) {
    const input = event.currentTarget as HTMLInputElement;
    files = Array.from(input.files ?? []);
    input.value = "";
    results = [];
    replacement = null;
    if (files.length) void upload(files);
  }

  function toggleClick() {
    if (isUploading) return;
    if (isDialogOpen) {
      files = [];
      results = [];
      replacement = null;
      message = "";
    }
    isDialogOpen = !isDialogOpen;
  }
</script>

<div class="upload-button">
  <button class="upload-trigger" onclick={toggleClick}>
    <span class="material-icons" aria-hidden="true"> file_upload </span>
    upload
  </button>

  {#if isDialogOpen}
    <div class="upload-modal" id="dialog" role="dialog" aria-modal="true" aria-labelledby="upload-title">
      <button
        class="dialog-backdrop"
        aria-label="閉じる" disabled={isUploading}
        onclick={toggleClick}
      ></button>
      <div class="dialog-panel">
        <header class="dialog-head">
          <p class="dialog-title" id="upload-title">明細書 アップロード</p>
          <button class="close-button" aria-label="閉じる" disabled={isUploading} onclick={toggleClick}></button>
        </header>
        <section class="dialog-body">
          {#if message != ""}
            <div class="error-message" role="alert">{message}</div>
          {/if}
          <label class="file-picker">
            <input
              id="fileInput"
              class="picker-input"
              type="file"
              accept=".pdf,application/pdf"
              multiple
              disabled={isUploading}
              onchange={fileChange}
            />
            <span class="picker-target">
              <span class="material-icons" aria-hidden="true"> file_upload </span>
              <span>PDFを選択（複数可）</span>
            </span>
            {#if isUploading}
              <span role="status">PDFを取り込み中です…</span>
              <progress class="upload-progress" max="100" aria-label="取り込み中"></progress>
            {/if}
          </label>
          <p class="filename-hint">ファイル名の例　2026年11月給与_明細.pdf</p>
          {#if results.length}
            <ul class="upload-results" aria-live="polite">
              {#each results as result, i}
                <li>
                  <strong>{result.filename}</strong>
                  <div>{labels[result.status]}</div>
                  {#if replacement === i}
                    <p>このPDFで現在の明細を上書きしますか？</p>
                    <button type="button" class="button is-small is-info" disabled={isUploading}
                      onclick={() => upload([files[i]], i)}>差し替える</button>
                    <button type="button" class="button is-small" disabled={isUploading}
                      onclick={() => (replacement = null)}>キャンセル</button>
                  {:else if result.status === "already_registered"}
                    <button type="button" class="button is-small" disabled={isUploading}
                      onclick={() => (replacement = i)}>差し替えて再取り込み</button>
                  {/if}
                </li>
              {/each}
            </ul>
          {/if}
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
    max-height: 90vh;
    overflow-y: auto;
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
  .filename-hint {
    margin-top: 0.5rem;
    color: #475569;
    font-size: 0.875rem;
  }

  .upload-results li {
    border-top: 1px solid #e5e7eb;
    margin-top: 0.75rem;
    padding-top: 0.75rem;
    overflow-wrap: anywhere;
  }

  .upload-results button {
    margin-top: 0.35rem;
  }
</style>
