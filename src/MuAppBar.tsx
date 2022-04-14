import 'bulma/css/bulma.css';
import UploadDialog from "./MuUploadButton"

export default function MuAppBar() {
  return (
    <header className="navbar is-dark">
      <div className="navbar-brand">
        <span className="navbar-item">sail</span>
      </div>
      <div className="navbar-end">
        <UploadDialog />
      </div>
    </header>
  )
}