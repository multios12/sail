import axios from 'axios';
import 'bulma/css/bulma.css';
import { useState } from 'react';

export default function MuUploadButton() {
  const [message, setMessage] = useState('')

  /** ファイル選択イベント */
  const fileChange = () => {
    const e = document.querySelector("#fileInput") as HTMLInputElement
    const file = e.files?.item(0)
    document.querySelector("#progress")?.classList.remove("is-hidden")
    if (!file) {
      setMessage("select file is null")
      return
    }
    const data = new FormData()
    data.append("file", file)
    axios.post("api/files", data).then((r) => {
      document.querySelector("#dialog")?.classList.remove('is-active');
      window.location.href = "./"
    }).catch(r => setMessage(r))
      .finally(() => document.querySelector("#progress")?.classList.add("is-hidden"))
  }

  /** モーダル表示イベント */
  const showClick = () => {
    document.querySelector("#dialog")?.classList.add('is-active');
  }

  /** モーダル非表示イベント */
  const closeClick = () => {
    document.querySelector("#dialog")?.classList.remove('is-active');
  }

  return (
    <div className="navbar-item field is-grouped">
      <p className='control nabbar-item '>
        <a role="button" className="button is-info" onClick={showClick}><span className="material-icons"> file_upload </span>upload</a>
      </p>

      <div className="modal" id="dialog">
        <div className="modal-background"></div>
        <div className="modal-card">
          <header className="modal-card-head">
            <p className="modal-card-title">明細書 アップロード</p>
            <button className="delete" aria-label="close" onClick={closeClick}></button>
          </header>
          <section className="modal-card-body">
            {message ?? <div className="notification is-danger">{message}</div>}
            <div className="file has-name is-boxed is-fullwidth">
              <label className="file-label">
                <input id="fileInput" className="file-input" type="file" name="resume" onChange={fileChange} />
                <span className="file-cta">
                  <span className="file-icon">
                    <span className="material-icons"> file_upload </span>
                  </span>
                  <span className="file-label">
                    Choose a file…
                  </span>
                  <progress id="progress" className="progress is-small is-primary is-hidden" max="100" />
                </span>
              </label>
            </div>
          </section>
        </div>
      </div>
    </div>
  )
}
