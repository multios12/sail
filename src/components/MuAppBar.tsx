import 'bulma/css/bulma.css';
import UploadButton from "./MuUploadButton"

export default () => {
  return (
    <header className="navbar is-dark">
      <div className="navbar-brand">
        <span className="navbar-item">sail</span>
      </div>
      <div className="navbar-end">
        <UploadButton />
      </div>
    </header>
  )
}