import React from "react";
import AuthService from "../services/auth.service";

const Profile = () => {
  const currentUser = AuthService.getCurrentUser();

  return (
    <div className="container">
      <header className="jumbotron">
        <h3>
          <strong>{currentUser}</strong> Profile
        </h3>
      </header>
      <p>
        <strong>Token: </strong>
        {JSON.parse(localStorage.getItem("token")).Authorization}
      </p>
    </div>
  );
};

export default Profile;
