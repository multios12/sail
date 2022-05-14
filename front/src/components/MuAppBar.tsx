import 'bulma/css/bulma.css';
import UploadButton from "./MuUploadButton"

const MuAppBar = () => {

  const burgerClick = () => {
    document.querySelector("#menu")?.classList.toggle('is-active');
    document.querySelector('.navbar-burger')?.classList.toggle('is-active');
  }

  return (
    <nav className="navbar is-dark" role="navigation" aria-label="main navigation">
      <div className="navbar-brand">
        <a className="navbar-item" href="./">sail</a>
        <a className="navbar-item" href="./#/salary"><i className="material-icons">attach_money</i></a>
        <a className="navbar-item" href="./#/cost"><i className="material-icons">payment</i></a>

        <button className="navbar-burger" onClick={burgerClick} aria-label="menu" aria-expanded="false">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </button>
      </div>

      <div id="menu" className="navbar-menu">
        <div className="navbar-start">
        </div>

        <div className="navbar-end">
          <UploadButton />
        </div>
      </div>

    </nav>
  )
}
export default MuAppBar
