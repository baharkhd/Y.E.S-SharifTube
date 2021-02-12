import { useState } from "react";
import constants from "../constants.js";

function useToken() {
  const getToken = () => {
    const tokenString = localStorage.getItem(constants.AUTH_TOKEN);
    return tokenString;
  };

  const [token, setToken] = useState(getToken());

  const saveToken = userToken => {
    if (userToken == undefined) {
      localStorage.removeItem(constants.AUTH_TOKEN);
    } else {
      localStorage.setItem(constants.AUTH_TOKEN, JSON.stringify(userToken));
      setToken(userToken);
    }
  };

  return {
    setToken: saveToken,
    token
  };
}

export default useToken;
